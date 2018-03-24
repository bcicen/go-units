package units

var (
	Volume = NewQuantity("volume")

	// metric
	Liter      = Volume.NewUnit("liter", "l", UnitOptionAliases("litre"))
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
	Quart      = Volume.NewUnit("quart", "qt", BI)
	Pint       = Volume.NewUnit("pint", "pt", BI)
	Gallon     = Volume.NewUnit("gallon", "gal", BI)
	FluidOunce = Volume.NewUnit("fluid ounce", "fl oz", BI)

	// US
	FluidQuart          = Volume.NewUnit("fluid quart", "", US)
	FluidPint           = Volume.NewUnit("fluid pint", "", US)
	FluidGallon         = Volume.NewUnit("fluid gallon", "", US)
	CustomaryFluidOunce = Volume.NewUnit("customary fluid ounce", "", US)
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
