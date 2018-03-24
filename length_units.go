package units

var (
	Length = UnitOptionQuantity("length")

	// metric
	Meter      = NewUnit("meter", "m", Length, SI, UnitOptionAliases("metre"))
	ExaMeter   = Exa(Meter)
	PetaMeter  = Peta(Meter)
	TeraMeter  = Tera(Meter)
	GigaMeter  = Giga(Meter)
	MegaMeter  = Mega(Meter)
	KiloMeter  = Kilo(Meter)
	HectoMeter = Hecto(Meter)
	DecaMeter  = Deca(Meter)
	DeciMeter  = Deci(Meter)
	CentiMeter = Centi(Meter)
	MilliMeter = Milli(Meter)
	MicroMeter = Micro(Meter)
	NanoMeter  = Nano(Meter)
	PicoMeter  = Pico(Meter)
	FemtoMeter = Femto(Meter)
	AttoMeter  = Atto(Meter)

	Inch    = NewUnit("inch", "in", Length, BI, UnitOptionPlural("inches"))
	Foot    = NewUnit("foot", "ft", Length, BI, UnitOptionPlural("feet"))
	Yard    = NewUnit("yard", "yd", Length, BI)
	Mile    = NewUnit("mile", "", Length, BI)
	League  = NewUnit("league", "lea", Length, BI)
	Furlong = NewUnit("furlong", "fur", Length, BI)
)

func init() {
	NewRatioConv(Inch, Meter, 0.0254)
	NewRatioConv(Foot, Meter, 0.3048)
	NewRatioConv(Yard, Meter, 0.9144)
	NewRatioConv(Mile, Meter, 1609.344)
	NewRatioConv(League, Meter, 4828.032)
	NewRatioConv(Furlong, Meter, 201.168)
}
