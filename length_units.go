package units

var (
	Length = NewQuantity("length")

	// metric
	Meter      = Length.NewUnit("meter", "m", SI, UnitOptionAliases("metre"))
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

	Inch    = Length.NewUnit("inch", "in", BI, UnitOptionPlural("inches"))
	Foot    = Length.NewUnit("foot", "ft", BI, UnitOptionPlural("feet"))
	Yard    = Length.NewUnit("yard", "yd", BI)
	Mile    = Length.NewUnit("mile", "", BI)
	League  = Length.NewUnit("league", "lea", BI)
	Furlong = Length.NewUnit("furlong", "fur", BI)
)

func init() {
	NewRatioConv(Inch, Meter, 0.0254)
	NewRatioConv(Foot, Meter, 0.3048)
	NewRatioConv(Yard, Meter, 0.9144)
	NewRatioConv(Mile, Meter, 1609.344)
	NewRatioConv(League, Meter, 4828.032)
	NewRatioConv(Furlong, Meter, 201.168)
}
