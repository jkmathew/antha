package liquidhandling

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/pkg/errors"
	"github.com/qri-io/jsonschema"

	"github.com/jkmathew/antha/antha/anthalib/wtype"
	"github.com/jkmathew/antha/antha/anthalib/wunit"
	driver "github.com/jkmathew/antha/microArch/driver/liquidhandling"
	simulator "github.com/jkmathew/antha/microArch/simulator/liquidhandling"
	"github.com/jkmathew/antha/utils"
)

const (
	LayoutSummaryVersion  = "1.0"
	ActionsSummaryVersion = "1.0"
)

//go:generate go-bindata -o ./schemas.go -pkg liquidhandling -prefix schemas/ ./schemas/

func validateJSON(schemaName string, jsonToValidate []byte) error {

	schema := &jsonschema.RootSchema{}
	if bs, err := Asset(schemaName); err != nil {
		panic(errors.WithMessage(err, fmt.Sprintf(`unable to load json schema "%s"`, schemaName)))
	} else if err := json.Unmarshal(bs, schema); err != nil {
		// the provided schema is invalid so we can't ever hope to generate a valid summary
		panic(errors.WithMessage(err, fmt.Sprintf(`invalid json schema "%s"`, schemaName)))
	}

	if errs, _ := schema.ValidateBytes(jsonToValidate); len(errs) > 0 {
		// the default Error() on this type is pretty sparse, including more detail to help with debugging
		e := make(utils.ErrorSlice, 0, len(errs))
		for _, err := range errs {
			e = append(e, errors.Errorf(`rule "%s" broken: failed to set property "%s": %s`, err.RulePath, err.PropertyPath, err.Error()))
		}
		return e.Pack()
	}
	return nil
}

// SummarizeLayout produce a description of the positions and states of objects on deck
// before and after the entire liquidhandling operation
// The returned JSON is validated against the schema found in ./schemas/layout.schema.json,
// which is the cannonical description of the format for communicating layout to the front end.
// initialState and finalState are the robot states before and after the operation,
// initialToFinalIDs maps object ids in the inisial state to the final state
// errors are returned if the json cannot be constructed or the result fails to validate
func SummarizeLayout(initialState, finalState *driver.LHProperties, initialToFinalIDs map[string]string) ([]byte, error) {
	ls := &layoutSummary{
		Before: newDeckSummary(initialState),
		After:  newDeckSummary(finalState),
		IDMap:  initialToFinalIDs,
	}

	if bs, err := json.Marshal(ls); err != nil {
		return nil, err
	} else if err := validateJSON("layout.schema.json", bs); err != nil {
		return nil, errors.WithMessage(err, "generated an invalid layout summary")
	} else {
		return bs, nil
	}
}

// SummarizeActions return a description of all the steps which take place during the liquidhandling operation
// The returned JSON is validated against the schema found in ./schemas/actions.schema.json, which
// is the cannonical description of the format for communication liquidhandling actions to the frontend
// initialState: the initial state of the robot, used to track state updates
// itree: the instruction tree generated during the Plan(...) stage
// errors are returned if the json cannot be constructed or the result fails to validate
func SummarizeActions(initialState *driver.LHProperties, itree *driver.ITree) ([]byte, error) {

	// nb. The physical simulator is used here to track the volumes and constituents of wells.
	// This is because the instructions themselves to not contain all the information required
	// (namely the sub-components), and in some situations contain wildly incorrect volumes.
	//
	// The physical simulator also has some drawbacks, namely that at present it does not
	// account for carry volume (though the functions below assume that it may).
	// It also does not present a full list of changes with each instruction, instead the identity
	// of the changed wells are inferred from the instructions.

	timer := initialState.GetTimer()
	var cumulativeTime time.Duration

	// create the simulator
	settings := simulator.DefaultSimulatorSettings()
	settings.EnablePipetteSpeedWarning(simulator.WarnNever)
	settings.EnableAutoChannelWarning(simulator.WarnNever)
	settings.EnableLiquidTypeWarning(simulator.WarnNever)
	settings.EnableTipboxCollision(false)
	vlh, err := simulator.NewVirtualLiquidHandler(initialState, settings)
	if err != nil {
		return nil, err
	}

	// initialize and setup the vlh
	vlh.Initialize()
	if err := vlh.Simulate(initialState.GetSetupInstructions()); err != nil {
		return nil, err
	}

	// we care about recording transfer block and message instructions
	acts := itree.Refine(driver.TFB, driver.MSG)

	actions := make(actionsSummary, 0, len(acts))
	for _, act := range acts {
		var timeForAct time.Duration
		if timer != nil {
			// find cumulative time estimate for all sub instructions
			for _, leaf := range act.Leaves() {
				timeForAct += timer.TimeFor(leaf)
			}
		}
		switch act.Instruction().Type() {
		case driver.MSG:
			// record messages as a prompt action
			if action, err := newPromptAction(vlh, act, cumulativeTime); err != nil {
				return nil, err
			} else {
				action.TimeEstimate = timeForAct.Seconds()
				cumulativeTime = time.Duration(action.CumulativeTimeEstimate * 1e9)
				actions = append(actions, action)
			}
		case driver.TFB:
			// record transfer block instructions as transfer actions
			if action, err := newTransferAction(vlh, act, timer, cumulativeTime); err != nil {
				return nil, err
			} else {
				action.TimeEstimate = timeForAct.Seconds()
				cumulativeTime = time.Duration(action.CumulativeTimeEstimate * 1e9)
				actions = append(actions, action)
			}
		default:
			// output anything else to the simulator to keep the state up to date
			if err := vlh.Simulate(act.Leaves()); err != nil {
				return nil, err
			}
			// keep cumulativeTime up to date for any instructions not included in actions
			cumulativeTime += timeForAct
		}
	}

	if bs, err := json.Marshal(actions); err != nil {
		return nil, err
	} else if err := validateJSON("actions.schema.json", bs); err != nil {
		return bs, errors.WithMessage(err, "generated an invalid action summary")
	} else {
		return bs, nil
	}
}

// layoutSummary summarize the layout of the deck before and after the liquidhandling step
type layoutSummary struct {
	Before *deckSummary      `json:"before"`  // the layout before the liquidhandling takes place
	After  *deckSummary      `json:"after"`   // the layout after the liquidhandling takes place
	IDMap  map[string]string `json:"new_ids"` // maps from ids in "before" to ids in "after"
}

func (ls *layoutSummary) MarshalJSON() ([]byte, error) {
	type LayoutSummaryAlias layoutSummary
	return json.Marshal(struct {
		*LayoutSummaryAlias
		Version string `json:"version"`
	}{
		LayoutSummaryAlias: (*LayoutSummaryAlias)(ls),
		Version:            LayoutSummaryVersion,
	})
}

// deckSummary summarize the layout of the deck
type deckSummary struct {
	Positions map[string]*deckPosition `json:"positions"` // map from position name to object description
}

// newDeckSummary create the deck layout from the properties file
func newDeckSummary(props *driver.LHProperties) *deckSummary {
	positions := make(map[string]*deckPosition, len(props.Positions))
	for posName, pos := range props.Positions {
		if objID, ok := props.PosLookup[posName]; ok {
			positions[posName] = &deckPosition{
				Position: newCoordinates3D(pos.Location),
				Size:     newCoordinates2D(pos.Size),
				Item:     newItemSummary(props.PlateLookup[objID].(wtype.LHObject)),
			}
		} else {
			positions[posName] = &deckPosition{
				Position: newCoordinates3D(pos.Location),
				Size:     newCoordinates2D(pos.Size),
			}
		}
	}

	return &deckSummary{Positions: positions}
}

// deckPosition a slot on the deck of a robot
type deckPosition struct {
	Position coordinates  `json:"position"`
	Size     coordinates  `json:"size"`
	Item     *itemSummary `json:"item,omitempty"`
}

// coordinates x,y,z coordinates with appropriate JSON struct tags
type coordinates struct {
	X float64 `json:"x_mm"`
	Y float64 `json:"y_mm"`
	Z float64 `json:"z_mm,omitempty"`
}

// newCoordinates3D
func newCoordinates3D(coord wtype.Coordinates3D) coordinates {
	return coordinates{X: coord.X, Y: coord.Y, Z: coord.Z}
}

// newCoordinates2D
func newCoordinates2D(coord wtype.Coordinates2D) coordinates {
	return coordinates{X: coord.X, Y: coord.Y}
}

// wellType the shape of the wells
type wellType string

const (
	roundWell  wellType = "cylinder"
	squareWell wellType = "cuboid"
)

type wellCoords struct {
	Row    int `json:"row"`
	Column int `json:"col"`
}

// itemSummary summarize an item on the deck and its initial contents
type itemSummary struct {
	ID             string                         `json:"id"`   // ID of the object to be referenced in actions
	Name           string                         `json:"name"` // display name for the item
	Type           string                         `json:"type"` // the type, e.g. DWST96
	Manufacturer   string                         `json:"manufacturer"`
	Kind           string                         `json:"kind"` // "plate", "tipbox", "tipwaste", etc
	Description    string                         `json:"description"`
	Rows           int                            `json:"rows"`
	Columns        int                            `json:"columns"`
	Dimensions     coordinates                    `json:"dimensions"`
	WellDimensions coordinates                    `json:"well_dimensions"`
	WellOffset     coordinates                    `json:"well_offset"`
	WellStart      coordinates                    `json:"well_start"`
	WellType       wellType                       `json:"well_type"`
	Contents       map[int]map[int]*liquidSummary `json:"contents,omitempty"` // Contents[column][row], omit means empty
	MissingTips    []*wellCoords                  `json:"missing_tips,omitempty"`
	ResidualVolume *measurementSummary            `json:"residual_volume,omitempty"`
}

// newItemSummary build an item summary from the object itself
func newItemSummary(obj wtype.LHObject) *itemSummary {

	switch o := obj.(type) {
	case *wtype.Plate:
		contents := make(map[int]map[int]*liquidSummary, o.NCols())
		for i, col := range o.Cols {
			c := make(map[int]*liquidSummary, o.NRows())
			for j, well := range col {
				if !well.IsEmpty() {
					c[j] = newLiquidSummary(well.Contents())
				}
			}
			if len(c) > 0 {
				contents[i] = c
			}
		}

		shape := squareWell
		if o.Welltype.Shape().Type.IsRound() {
			shape = roundWell
		}

		return &itemSummary{
			ID:             o.ID,
			Name:           o.PlateName,
			Type:           o.Type,
			Manufacturer:   o.Mnfr,
			Kind:           "plate",
			Description:    fmt.Sprintf("Plate with %dx%d wells", o.NCols(), o.NRows()),
			Dimensions:     newCoordinates3D(o.GetSize()),
			WellDimensions: newCoordinates3D(o.Welltype.GetSize()),
			WellStart:      coordinates{X: o.WellXStart, Y: o.WellYStart, Z: o.WellZStart},
			WellOffset:     coordinates{X: o.WellXOffset, Y: o.WellYOffset},
			WellType:       shape,
			Rows:           o.NRows(),
			Columns:        o.NCols(),
			Contents:       contents,
			ResidualVolume: newMeasurementSummary(o.Welltype.ResidualVolume()),
		}
	case *wtype.LHTipbox:
		missingTips := make([]*wellCoords, 0, o.NCols()*o.NRows())
		for rowNum, row := range o.Tips {
			for colNum, tip := range row {
				if tip == nil {
					missingTips = append(missingTips, &wellCoords{Row: rowNum, Column: colNum})
				}
			}
		}

		return &itemSummary{
			ID:             o.ID,
			Name:           o.Boxname,
			Type:           o.Type,
			Manufacturer:   o.Mnfr,
			Kind:           "tipbox",
			Description:    fmt.Sprintf("Tipbox containing \"%s\" tips from %s", o.Tiptype.Type, o.Mnfr),
			Dimensions:     newCoordinates3D(o.GetSize()),
			WellDimensions: newCoordinates3D(o.AsWell.GetSize()),
			WellStart:      coordinates{X: o.TipXStart, Y: o.TipYStart, Z: o.TipZStart},
			WellOffset:     coordinates{X: o.TipXOffset, Y: o.TipYOffset},
			WellType:       roundWell,
			Rows:           o.NRows(),
			Columns:        o.NCols(),
			MissingTips:    missingTips,
		}
	case *wtype.LHTipwaste:
		shape := squareWell
		if o.AsWell.Shape().Type.IsRound() {
			shape = roundWell
		}

		return &itemSummary{
			ID:             o.ID,
			Name:           o.Name,
			Type:           o.Type,
			Manufacturer:   o.Mnfr,
			Kind:           "tipwaste",
			Description:    fmt.Sprintf("Tipwaste of type \"%s\" by %s", o.Type, o.Mnfr),
			Dimensions:     newCoordinates3D(o.GetSize()),
			WellDimensions: newCoordinates3D(o.AsWell.GetSize()),
			WellStart:      coordinates{X: o.WellXStart, Y: o.WellYStart, Z: o.WellZStart},
			WellType:       shape,
			Rows:           1,
			Columns:        1,
		}
	}
	panic(fmt.Sprintf("Unknown object of type %T", obj))
}

// measurementSummary summarize a measurement
type measurementSummary struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// newMeasurementSummary convert from wunit to Json-friendly representation
func newMeasurementSummary(v wunit.Volume) *measurementSummary {
	return &measurementSummary{
		Value: v.RawValue(),
		Unit:  v.Unit().PrefixedSymbol(),
	}
}

// height records the height during pipetting
type height struct {
	measurementSummary
	Reference wtype.WellReference
}

func (h height) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		measurementSummary
		Reference string `json:"reference"`
	}{
		measurementSummary: h.measurementSummary,
		Reference:          h.Reference.String(),
	})
}

// liquidSummary summarize a liquid (nee LHComponent)
type liquidSummary struct {
	Name        string              `json:"name"`
	TotalVolume *measurementSummary `json:"total_volume"`
	Components  []subComponent      `json:"components,omitempty"` // what other liquids make up this liquid
}

// newLiquidSummary create a liquid summary
func newLiquidSummary(l *wtype.Liquid) *liquidSummary {
	return &liquidSummary{
		Name:        l.MeaningfulName(),
		TotalVolume: &measurementSummary{Value: l.Vol, Unit: l.Vunit},
		Components:  newSubComponents(l.SubComponents),
	}
}

// subComponent summarize a sub-component of a liquid
// n.b: this is perhaps not the most useful format long-term, since users probably
// prefer to see subcomponent volume rather than concentration (in arbitrary units)
// but that's what we currently store in the wtype.Liquid object
type subComponent struct {
	Name          string  `json:"name"` // display name of this part of the component
	Concentration float64 `json:"concentration"`
	Unit          string  `json:"unit"`
}

func newSubComponents(cl wtype.ComponentList) []subComponent {
	names := make([]string, 0, len(cl.Components))
	for name := range cl.Components {
		names = append(names, name)
	}
	sort.Strings(names)

	r := make([]subComponent, 0, len(cl.Components))
	for _, name := range names {
		r = append(r, subComponent{
			Name:          name,
			Concentration: cl.Components[name].RawValue(),
			Unit:          cl.Components[name].Unit().PrefixedSymbol(),
		})
	}
	return r
}

// action an operation that is carried out by the liquidhandler (or possibly the user) during
// the liquidhandling task which is to be displayed to the user during the mix summary
//
// actions should be designed such that they store any updates in robot state, e.g. if
// an action changes the liquid in a well, that change should be explicitly stored so the
// front end is not required to do any maths on the fly
type action interface {
	json.Marshaler
	isAction()
}

type wellLocation struct {
	DeckItemID string `json:"deck_item_id"`
	Row        int    `json:"row"`
	Column     int    `json:"col"`
}

type actionsSummary []action

func (as actionsSummary) MarshalJSON() ([]byte, error) {
	type ActionsSummaryAlias actionsSummary
	return json.Marshal(struct {
		Actions ActionsSummaryAlias `json:"actions"`
		Version string              `json:"version"`
	}{
		Actions: ActionsSummaryAlias(as),
		Version: ActionsSummaryVersion,
	})
}

// contentUpdate
type contentUpdate struct {
	Location   wellLocation   `json:"loc"`
	NewContent *liquidSummary `json:"new_content"`
}

func newContentUpdate(well *wtype.LHWell) *contentUpdate {
	return &contentUpdate{
		Location: wellLocation{
			DeckItemID: wtype.IDOf(well.GetParent()),
			Row:        well.Crds.Y,
			Column:     well.Crds.X,
		},
		NewContent: newLiquidSummary(well.Contents()),
	}
}

type mixSummary struct {
	Cycles   int                 `json:"cycles"`
	Volume   *measurementSummary `json:"volume"`
	Height   *height             `json:"height"`
	FlowRate *measurementSummary `json:"flow_rate"`
	LLF      bool                `json:"liquid_level_follow"`
}

type blowoutSummary struct {
	Volume   *measurementSummary `json:"volume"`
	Height   *height             `json:"height"`
	FlowRate *measurementSummary `json:"flow_rate"`
}

type pipettingOptions struct {
	Height   *height             `json:"height"`
	FlowRate *measurementSummary `json:"flow_rate"`
	LLF      bool                `json:"liquid_level_follow"`
	Mixing   *mixSummary         `json:"mixing,omitempty"`
	Blowout  *blowoutSummary     `json:"blowout,omitempty"`
	TouchOff *bool               `json:"touchoff,omitempty"` // this is ptr-to-bool so json will omit the key if value is nil, otherwise set to true or false
}

// transferSummary summarizes a single one-to-one transfer, giving the updated contents of both locations
// as well as other details of the transfer
type transferSummary struct {
	From              *contentUpdate      `json:"from"`   // the source from which liquid is taken, and the new contents
	To                []*contentUpdate    `json:"to"`     // the destination(s) in which liquid is placed, and their new contents. Multi-dispenses are represented as a slice of destinations
	Volume            *measurementSummary `json:"volume"` // the volume of liquid transfered
	Wasted            *measurementSummary `json:"wasted"` // the volume lost as "carry volume" during the transfer
	Policy            string              `json:"policy"` // the liquid type for this transfer
	AspirateBehaviour *pipettingOptions   `json:"asp"`
	DispenseBehaviour *pipettingOptions   `json:"dsp"`
	Head              int                 `json:"head"`
}

// parallelTransfer a slice of one or more Transfers which happen simultaneously
type parallelTransfer struct {
	Channels               map[int]*transferSummary `json:"channels"`
	TimeEstimate           float64                  `json:"time_estimate"`
	CumulativeTimeEstimate float64                  `json:"cumulative_time_estimate"`
}

// newParallelTransfer create a parallelTransfer from the ChannelTransferInstruction held by the ITree node
// inspects the lower level instructions generated by the CTI to determine the precise details
func newParallelTransfer(vlh *simulator.VirtualLiquidHandler, tree *driver.ITree, timeEstimate, cumulativeTimeEstimate time.Duration) (*parallelTransfer, error) {
	cti := tree.Instruction().(*driver.ChannelTransferInstruction)

	// fetch a well from the virtual liquidhandler
	getWell := func(address, wellcoords string) *wtype.LHWell {
		plate := vlh.GetObjectAt(address).(*wtype.LHPlate)
		well, _ := plate.WellAtString(wellcoords)
		return well
	}

	// channels are not necessarily contiguous, find which ones are used and the minimum number of channels to represent the transfer
	channels := cti.Channels()
	numChannels := 0
	for _, ch := range channels {
		if ch+1 > numChannels {
			numChannels = ch + 1
		}
	}

	// fetch each of the source and destination wells
	// a fair amount of the complexity here is dealing with the "trough" case, when these wells are actually the same wells
	sourceWells := make([]*wtype.LHWell, numChannels)
	destWells := make([]*wtype.LHWell, numChannels)
	for _, ch := range channels {
		sourceWells[ch] = getWell(cti.PltFrom[ch], cti.WellFrom[ch])
		destWells[ch] = getWell(cti.PltTo[ch], cti.WellTo[ch])
	}

	// map from source well to channel index
	channelsInWell := make(map[*wtype.LHWell][]int, len(sourceWells))
	for ch, well := range sourceWells {
		if well != nil {
			channelsInWell[well] = append(channelsInWell[well], ch)
		}
	}

	// what volumes do we expect to be left in the source wells
	expectedVolume := make(map[*wtype.LHWell]wunit.Volume, len(channels))
	for well, channels := range channelsInWell {
		ev := well.CurrentVolume()
		for _, ch := range channels {
			ev = wunit.SubtractVolumes(ev, cti.Volume[ch])
		}
	}

	// simulate the transfer
	instructions := tree.Leaves()
	if err := vlh.Simulate(instructions); err != nil {
		return nil, err
	}

	// helpers to create *bools (leaving as nil omits the key)
	newTrue := func() *bool {
		b := true
		return &b
	}
	newFalse := func() *bool {
		b := false
		return &b
	}

	// pick out the relevant details from the lowest level instructions
	aspOptions := make([]pipettingOptions, numChannels)
	dspOptions := make([]pipettingOptions, numChannels)
	pipetteSpeed := make([]*measurementSummary, numChannels)
	lastHeight := make([]*height, numChannels)
	seenDispense := false

	for _, ins := range instructions {
		ins.Visit(driver.RobotInstructionBaseVisitor{
			HandleSetPipetteSpeed: func(sps *driver.SetPipetteSpeedInstruction) {
				if sps.Channel < 0 {
					for _, ch := range channels {
						// yes, we really do use "ml/min" as the unit for pipette speed
						pipetteSpeed[ch] = &measurementSummary{Value: sps.Speed, Unit: "ml/min"}
					}
				} else {
					pipetteSpeed[sps.Channel] = &measurementSummary{Value: sps.Speed, Unit: "ml/min"}
				}
			},
			HandleAspirate: func(asp *driver.AspirateInstruction) {
				for _, ch := range channels {
					aspOptions[ch].FlowRate = &(*pipetteSpeed[ch])
					aspOptions[ch].Height = &(*lastHeight[ch])
					aspOptions[ch].LLF = asp.LLF[ch]
				}
			},
			HandleDispense: func(dsp *driver.DispenseInstruction) {
				seenDispense = true
				for _, ch := range channels {
					dspOptions[ch].FlowRate = &(*pipetteSpeed[ch])
					dspOptions[ch].Height = &(*lastHeight[ch])
					dspOptions[ch].LLF = dsp.LLF[ch]

					// last move wasn't a touchoff
					dspOptions[ch].TouchOff = newFalse()
				}
			},
			HandleBlowout: func(blo *driver.BlowoutInstruction) {
				for _, ch := range channels {
					dspOptions[ch].Blowout = &blowoutSummary{
						Volume:   newMeasurementSummary(blo.Volume[ch]),
						Height:   &(*lastHeight[ch]),
						FlowRate: &(*pipetteSpeed[ch]),
					}
					// last move wasn't a touchoff
					dspOptions[ch].TouchOff = newFalse()
				}
			},
			HandleMix: func(mix *driver.MixInstruction) {
				for _, ch := range channels {
					ms := &mixSummary{
						Volume:   newMeasurementSummary(mix.Volume[ch]),
						Cycles:   mix.Cycles[ch],
						FlowRate: &(*pipetteSpeed[ch]),
						Height:   &(*lastHeight[ch]),
					}
					// nb. LLF not specified in driver.MixInstruction
					if !seenDispense {
						aspOptions[ch].Mixing = ms
					} else {
						dspOptions[ch].Mixing = ms
					}

					// last move wasn't a touchoff
					dspOptions[ch].TouchOff = newFalse()
				}
			},
			HandleMove: func(move *driver.MoveInstruction) {
				for _, ch := range channels {
					lastHeight[ch] = &height{
						measurementSummary: measurementSummary{Value: move.OffsetZ[ch], Unit: "mm"},
						Reference:          wtype.WellReference(move.Reference[ch]),
					}

					// set a touchoff if this is the last move
					dspOptions[ch].TouchOff = newTrue()
				}

			},
		})
	}

	// now build each transfer
	transfers := make(map[int]*transferSummary, len(channels))
	for _, ch := range channels {
		// divide total missing volume equally between each channel that was in the source well
		missing := wunit.SubtractVolumes(expectedVolume[sourceWells[ch]], sourceWells[ch].CurrentVolume())
		missing.DivideBy(float64(len(channelsInWell[sourceWells[ch]])))

		tfs := &transferSummary{
			From:              newContentUpdate(sourceWells[ch]),
			To:                []*contentUpdate{newContentUpdate(destWells[ch])}, // multi dispense will require multiple entries here
			Volume:            newMeasurementSummary(wunit.CopyVolume(cti.Volume[ch])),
			Wasted:            newMeasurementSummary(missing),
			Policy:            cti.What[ch],
			AspirateBehaviour: &aspOptions[ch],
			DispenseBehaviour: &dspOptions[ch],
			Head:              cti.Prms[ch].Head,
		}

		transfers[ch] = tfs
	}

	return &parallelTransfer{
		Channels:               transfers,
		TimeEstimate:           timeEstimate.Seconds(),
		CumulativeTimeEstimate: cumulativeTimeEstimate.Seconds(),
	}, nil
}

func (pt *parallelTransfer) MarshalJSON() ([]byte, error) {
	// alias the type so as not to invoke this function in a loop
	type Alias parallelTransfer
	return json.Marshal(struct {
		Kind string `json:"kind"`
		*Alias
	}{
		Kind:  "parallel_transfer",
		Alias: (*Alias)(pt),
	})
}

func (*parallelTransfer) isTransferChild() {}

type transferChild interface {
	isTransferChild()
}

// transferAction represents all the transfers carried out by a TransferInstruction as a slice of parallelTransfers which occur in serial.
// Ultimately, this represents all the instructions which were sorted to a single link in the IChain, i.e. the results of high level LHInstructions
// which _could_ all be executed together given a sufficiently flexible device
type transferAction struct {
	Children               []transferChild `json:"children"`
	TimeEstimate           float64         `json:"time_estimate"`
	CumulativeTimeEstimate float64         `json:"cumulative_time_estimate"`
}

// newTransferAction create a new transfer action from the act, which is assumed to have generated ChannelTransferInstructions
// and outputs all leaves of the act to the simulator
func newTransferAction(vlh *simulator.VirtualLiquidHandler, act *driver.ITree, timer driver.LHTimer, cumulativeTimeEstimate time.Duration) (*transferAction, error) {

	instructions := act.Refine(driver.CTI)

	children := make([]transferChild, 0, len(act.Children()))

	// keep track of the last time each location was updated so we can update the final name later
	lastUpdate := make(map[wellLocation]*liquidSummary)

	for _, ins := range instructions {
		var timeEstimate time.Duration

		for _, leaf := range ins.Leaves() {
			timeEstimate += timer.TimeFor(leaf)
		}

		cumulativeTimeEstimate += timeEstimate

		switch ins.Instruction().Type() {
		case driver.CTI:
			if pt, err := newParallelTransfer(vlh, ins, timeEstimate, cumulativeTimeEstimate); err != nil {
				return nil, err
			} else {
				for _, ch := range pt.Channels {
					for _, cUpdate := range ch.To {
						lastUpdate[cUpdate.Location] = cUpdate.NewContent
					}
				}
				children = append(children, pt)
			}
		case driver.LOD:
			load := ins.Instruction().(*driver.LoadTipsInstruction)
			children = append(children, newTipAction(vlh, loadTipAction, load.Multi, load.Head, load.Pos, load.Well, timeEstimate, cumulativeTimeEstimate))

			if err := load.OutputTo(vlh); err != nil {
				return nil, err
			}

		case driver.ULD:
			unload := ins.Instruction().(*driver.UnloadTipsInstruction)
			children = append(children, newTipAction(vlh, unloadTipAction, unload.Multi, unload.Head, unload.Pos, unload.Well, timeEstimate, cumulativeTimeEstimate))

			if err := unload.OutputTo(vlh); err != nil {
				return nil, err
			}

		default:
			if err := vlh.Simulate(ins.Leaves()); err != nil {
				return nil, err
			}
		}
	}

	// the transferBlockInstruction includes the original mix instructions and their outputs with user defined names set,
	// here we use those to set the final names of the liquids
	tfb := act.Instruction().(*driver.TransferBlockInstruction)

	for _, ins := range tfb.Inss {
		for _, output := range ins.Outputs {
			pl := output.PlateLocation()
			if content, ok := lastUpdate[wellLocation{DeckItemID: pl.ID, Row: pl.Coords.Y, Column: pl.Coords.X}]; ok {
				content.Name = output.MeaningfulName()
				// update the name within the virtual liquidhandler so that it persists if the component is aspirated from later
				if well := vlh.GetWellAt(pl); well != nil {
					well.Contents().SetName(output.MeaningfulName())
				}
			}
		}
	}

	return &transferAction{Children: children, CumulativeTimeEstimate: cumulativeTimeEstimate.Seconds()}, nil
}

func (*transferAction) isAction() {}

func (ta *transferAction) MarshalJSON() ([]byte, error) {
	// alias the type so as not to invoke this function in a loop
	type Alias transferAction
	return json.Marshal(struct {
		Kind string `json:"kind"`
		*Alias
	}{
		Kind:  "transfer",
		Alias: (*Alias)(ta),
	})
}

type tipActionType string

const (
	loadTipAction   tipActionType = "load"
	unloadTipAction tipActionType = "unload"
)

// tipAction a load or unload tips action
type tipAction struct {
	Kind                   tipActionType         `json:"kind"`
	Head                   int                   `json:"head"`
	Channels               map[int]*wellLocation `json:"channels"`
	TimeEstimate           float64               `json:"time_estimate"`
	CumulativeTimeEstimate float64               `json:"cumulative_time_estimate"`
}

func newTipAction(vlh *simulator.VirtualLiquidHandler, kind tipActionType, multi, head int, positions, wellcoords []string, timeEstimate, cumulativeTimeEstimate time.Duration) *tipAction {
	tipSources := make(map[int]*wellLocation, multi)
	for ch := 0; ch < len(positions); ch++ {
		// ignore channels where no tip is loaded/unloaded
		if wc := wtype.MakeWellCoords(wellcoords[ch]); positions[ch] != "" && !wc.IsZero() {
			tipSources[ch] = &wellLocation{
				DeckItemID: wtype.IDOf(vlh.GetObjectAt(positions[ch])),
				Row:        wc.Y,
				Column:     wc.X,
			}
		}
	}

	return &tipAction{
		Kind:                   kind,
		Head:                   head,
		Channels:               tipSources,
		TimeEstimate:           timeEstimate.Seconds(),
		CumulativeTimeEstimate: cumulativeTimeEstimate.Seconds(),
	}
}

func (*tipAction) isTransferChild() {}

type promptAction struct {
	DurationSeconds        float64 `json:"duration_seconds,omitempty"`
	CumulativeTimeEstimate float64 `json:"cumulative_time_estimate"`
	TimeEstimate           float64 `json:"time_estimate"`
	Message                string  `json:"message"`
}

func newPromptAction(vlh *simulator.VirtualLiquidHandler, act *driver.ITree, cumulativeTime time.Duration) (*promptAction, error) {
	msg := act.Instruction().(*driver.MessageInstruction)

	if err := msg.OutputTo(vlh); err != nil {
		return nil, err
	}
	return &promptAction{
		Message:                msg.Message,
		CumulativeTimeEstimate: (msg.WaitTime + cumulativeTime).Seconds(),
		TimeEstimate:           msg.WaitTime.Seconds(),
	}, nil
}

func (*promptAction) isAction() {}

func (ma *promptAction) MarshalJSON() ([]byte, error) {
	// alias the type so as not to invoke this function in a loop
	type Alias promptAction
	return json.Marshal(struct {
		Kind string `json:"kind"`
		*Alias
	}{
		Kind:  "prompt",
		Alias: (*Alias)(ma),
	})
}
