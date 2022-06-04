package platereading

import (
	"time"
)

// these are taken from a specific type of platereader... most likely need
// generalising

const (
	PR_READY = iota
	PR_BUSY
	PR_RUNNING
	PR_PAUSING
	PR_ERROR
	PR_HWERROR
)

type WaveSet struct {
	Min int // minimum
	Max int // maximum
	Inc int // increment
}

type PRMeasurement struct {
	EWavelength int       //	excitation wavelength
	RWavelength int       //	emission wavelength
	Reading     int       // 	value read
	Xoff        int       //	position - x, relative to well centre
	Yoff        int       //	position - y, relative to well centre
	Zoff        int       // 	position - z, relative to well centre
	Timestamp   time.Time // instant measurement was taken
	Temp        int       //   temperature
	O2          int       // o2 conc when measurement was taken
	CO2         int       // co2 conc when measurement was taken
}

type PRMeasurementSet []PRMeasurement

type PROutput struct {
	Readings []PRMeasurementSet
}

type PRProperties struct {
	ApertureType string
	AbsOptionIn  bool    // can we measure absorbance?
	LumiOptionIn bool    // luminescence?
	TRFOptionIn  bool    // time resolved fluorescence
	InubIn       bool    // incubator?
	IncTempMin   int     // temperature range - minimum
	IncTempMax   int     // temperature range - maximum
	AbsLimits    WaveSet // absorbance range measurable
	LumLimits    WaveSet // luminescence range measurable
	FexLimits    WaveSet // fluorescence excitation range
	FemLimits    WaveSet // fluorescence emission range
	Valcount     int     // number of data blocks that can be stored
}

type PRState struct {
	// encapsulates state information
	Status                     int    // current device status
	DeviceBusy                 bool   // is device busy?
	DeviceWaitingForEndOfCycle bool   // is device waiting?
	DevicePausing              bool   // is device paused pending continue?
	DevicePausingTime          bool   // is device paused for a time period?
	DeviceError                bool   // hardware error?
	DeviceWarning              bool   // warning generated??
	QuitCodes                  []int  // Array for error codes on quit
	Error                      string // last error / warning message
	SoftNum                    string // software version number
	EPROMNum                   string // reader firmware version
	BoardNum                   string // main board version / measurement board version
	PlateOut                   bool   // plate carrier outside instrument?
	ReagOpen                   bool   // instrument lid open?
	MeasPlateInserted          bool   // is there a plate in the machine?
	MeasPlateValid             bool   // sometimes the above is not valid
	Temp1                      int    // temp of top plate
	Temp2                      int    // temp of bottom plate
	T1notreached               bool   // reached first target temp?
	T2notreached               bool   // reached second target temp?
	O2Conc                     int    // oxygen concentration
	CO2Conc                    int    // co2 concentration
	TestDur                    int    // test run duration
	IntTime                    int    // kinetic interval time
	MeasureData                bool   // is there measurement data available?
	ActRow                     int    // last well row whose results are available
	ActCol                     int    //  ""   ""  col   " ...
	ActCycle                   int    // last kinetic cycle results available
	ActRowRet                  int    // last well row whose results were read
	ActColRet                  int    // ditto column
	ActCycleRet                int    // ditto kinetic cycle
	OffsetData                 bool   // offset determination data available?
	TimeData                   bool   // test run /cycle time available?
	GainData                   bool   // gain and focus values available?
	FocalHeight                int    // last optimal focal height determined
	FocusRaw                   int    // raw result obtained during focal height adjustment
	Gain1                      int    // gain value determined for channel A
	Gain1Raw                   int    // raw gain absolute
	Gain1Percent               int    // raw result gain relative
	KFactor                    int    // k factor determined
	ActRowGain                 int    // row used for gain adjustment
	ActColGain                 int    // col used for gain adjustment
	MotorEnabled               bool   // motors enabled?
	StackerStatus              bool   // stacker status, if attached
	StackerKindOfResponse      string // last response from stacker
}
