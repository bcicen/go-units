package units

var (
	Volume = UnitOptionQuantity("volume")

	// metric
	Liter      = NewUnit("liter", "l", Volume, UnitOptionAliases("litre"))
	ExaLiter   = Exa(Liter)
	PetaLiter  = Peta(Liter)
	TeraLiter  = Tera(Liter)
	GigaLiter  = Giga(Liter)
	MegaLiter  = Mega(Liter)
	KiloLiter  = Kilo(Liter)
	HectoLiter = Hecto(Liter)
	DecaLiter  = Deca(Liter)
	DeciLiter  = Deci(Liter)
	CentiLiter = Centi(Liter)
	MilliLiter = Milli(Liter)
	MicroLiter = Micro(Liter)
	NanoLiter  = Nano(Liter)
	PicoLiter  = Pico(Liter)
	FemtoLiter = Femto(Liter)
	AttoLiter  = Atto(Liter)

	// imperial
	Quart      = NewUnit("quart", "qt", Volume, BI)
	Pint       = NewUnit("pint", "pt", Volume, BI)
	Gallon     = NewUnit("gallon", "gal", Volume, BI)
	FluidOunce = NewUnit("fluid ounce", "fl oz", Volume, BI)

	// US
	FluidQuart          = NewUnit("fluid quart", "", Volume, US)
	FluidPint           = NewUnit("fluid pint", "", Volume, US)
	FluidGallon         = NewUnit("fluid gallon", "", Volume, US)
	CustomaryFluidOunce = NewUnit("customary fluid ounce", "", Volume, US)
)

func init() {
	NewRatioConv(Quart, Liter, 1.1365225)
	NewRatioConv(Pint, Liter, 0.56826125)
	NewRatioConv(Gallon, Liter, 4.54609)
	NewRatioConv(FluidOunce, MilliLiter, 28.4130625)

	NewRatioConv(FluidQuart, Liter, 0.946352946)
	NewRatioConv(FluidPint, Liter, 0.473176473)
	NewRatioConv(FluidGallon, Liter, 3.785411784)
	NewRatioConv(CustomaryFluidOunce, MilliLiter, 29.5735295625)
}
