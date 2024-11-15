// antha/AnthaStandardLibrary/Packages/Platereader/Platereader.go: Part of the Antha language
// Copyright (C) 2015 The Antha authors. All rights reserved.
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
//
// For more information relating to the software or licensing issues please
// contact license@jkmathew.org or write to the Antha team c/o
// Synthace Ltd. The London Bioscience Innovation Centre
// 2 Royal College St, London NW1 0NH UK

//Package platereader contains functions for manipulating absorbance readings and platereader data.
package platereader

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/jkmathew/antha/antha/anthalib/wtype"
	"github.com/jkmathew/antha/antha/anthalib/wunit"
)

func ReadAbsorbance(plate *wtype.Plate, solution *wtype.Liquid, wavelength float64) (abs wtype.Absorbance) {
	abs.Reading = 0.0 // obviously placeholder
	abs.Wavelength = wavelength
	// add calculation to work out pathlength from volume and well geometry abs.Pathlength

	return abs
}

func Blankcorrect(blank wtype.Absorbance, sample wtype.Absorbance) (blankcorrected wtype.Absorbance, err error) {

	if sample.Wavelength == blank.Wavelength &&
		sample.Pathlength == blank.Pathlength &&
		sample.Reader == blank.Reader {
		blankcorrected.Reading = sample.Reading - blank.Reading

		blankcorrected.Annotations = append(blankcorrected.Annotations, sample.Annotations...)

		blankcorrected.Annotations = append(blankcorrected.Annotations, "Blank Corrected")
	} else {
		err = fmt.Errorf("Cannot pathlength correct as Absorbance readings %+v and %+v are incompatible due to either wavelength, pathlength or reader differences", sample, blank)
	}
	return
}

// EstimatePathLength estimate the height of liquid of the given volume in the plates welltype - the length of the light path
func EstimatePathLength(plate *wtype.Plate, volume wunit.Volume) (wunit.Length, error) {

	if plate.Welltype.Bottom != wtype.FlatWellBottom {
		return wunit.ZeroLength(), errors.Errorf("Can't estimate pathlength for non flat bottom type: %s", plate.Welltype.Bottom)
	}

	// get the maximum cross-sectional area in mm^2
	var maxWellAreaMM2 float64
	if area, err := plate.Welltype.CalculateMaxCrossSectionArea(); err != nil {
		return wunit.ZeroLength(), errors.WithMessage(err, "while calculating max cross-section")
	} else if area, err := area.InStringUnit("mm^2"); err != nil {
		return wunit.ZeroLength(), errors.WithMessage(err, "while converting max cross-section")
	} else {
		maxWellAreaMM2 = area.RawValue()
	}

	// get the geometric well volume in ul
	var wellVolumeUL float64
	if vol, err := plate.Welltype.CalculateMaxVolume(); err != nil {
		return wunit.ZeroLength(), errors.WithMessage(err, "while calculating well volume")
	} else if vol, err := vol.InStringUnit("ul"); err != nil {
		return wunit.ZeroLength(), errors.WithMessage(err, "while converting well volume")
	} else {
		wellVolumeUL = vol.RawValue()
	}

	// get the volume argument in ul
	var volumeUL float64
	if vol, err := volume.InStringUnit("ul"); err != nil {
		return wunit.ZeroLength(), errors.WithMessage(err, "while converting volume argument")
	} else {
		volumeUL = vol.RawValue()
	}

	ratio := volumeUL / wellVolumeUL

	wellheightinmm := wellVolumeUL / maxWellAreaMM2

	return wunit.NewLength(wellheightinmm*ratio, "mm"), nil

}

func PathlengthCorrect(pathlength wunit.Length, reading wtype.Absorbance) (pathlengthcorrected wtype.Absorbance) {

	referencepathlength := wunit.NewLength(10, "mm")

	pathlengthcorrected.Reading = reading.Reading * referencepathlength.RawValue() / pathlength.RawValue()
	return
}

// based on Beer Lambert law A = ε l c
/*
Limitations of the Beer-Lambert law

The linearity of the Beer-Lambert law is limited by chemical and instrumental factors. Causes of nonlinearity include:
deviations in absorptivity coefficients at high concentrations (>0.01M) due to electrostatic interactions between molecules in close proximity
scattering of light due to particulates in the sample
fluoresecence or phosphorescence of the sample
changes in refractive index at high analyte concentration
shifts in chemical equilibria as a function of concentration
non-monochromatic radiation, deviations can be minimized by using a relatively flat part of the absorption spectrum such as the maximum of an absorption band
stray light
*/
func Concentration(pathlengthcorrected wtype.Absorbance, molarabsorbtivityatwavelengthLpermolpercm float64) (conc wunit.Concentration) {

	A := pathlengthcorrected
	l := 1                                         // 1cm if pathlengthcorrected add logic to use pathlength of absorbance reading input
	ε := molarabsorbtivityatwavelengthLpermolpercm // L/Mol/cm

	concfloat := A.Reading / (float64(l) * ε) // Mol/L
	// fmt.Println("concfloat", concfloat)
	conc = wunit.NewConcentration(concfloat, "M/l")
	// fmt.Println("concfloat", conc)
	return
}

//example

/*
func OD(Platetype wtype.LHPLate,wellvolume wtype.Volume,reading wtype.Absorbance) (od wtype.Absorbance){
volumetopathlengthconversionfactor := 0.0533//WellCrosssectionalArea
OD = (Blankcorrected_absorbance * 10/(total_volume*volumetopathlengthconversionfactor)// 0.0533 could be written as function of labware and liquid volume (or measureed height)
}

DCW = OD * ODtoDCWconversionfactor

*/
/*
type Absorbance struct {
	Reading    float64
	Wavelength float64
	Pathlength *wtype.Length
	Status     *[]string
	Reader     *string
}
*/
