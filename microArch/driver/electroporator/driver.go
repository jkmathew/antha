package electroporator

type ElectroporationDriver interface {
	Pulse(Voltage, Capacitance, Resistance int)
}
