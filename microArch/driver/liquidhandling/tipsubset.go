package liquidhandling

import (
	"github.com/jkmathew/antha/antha/anthalib/wtype"
)

type TipSubset struct {
	Mask    []bool
	TipType string
	Channel *wtype.LHChannelParameter
}
