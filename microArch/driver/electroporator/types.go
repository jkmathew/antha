package electroporator

// TODO identify cuvette compatibility
type EPProperties struct {
	MinVoltage     int
	MaxVoltage     int
	MinCapacitance int
	MaxCapacitance int
	MinResistance  int
	MaxResistance  int
}

type EPStatus struct {
	Ready        bool
	Error        bool
	ErrorMessage string
	ErrorType    int
}
