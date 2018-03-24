package units

var (
	Mass = UnitOptionQuantity("mass")

	// metric
	Gram      = NewUnit("gram", "g", Mass)
	ExaGram   = Exa(Gram)
	PetaGram  = Peta(Gram)
	TeraGram  = Tera(Gram)
	GigaGram  = Giga(Gram)
	MegaGram  = Mega(Gram)
	KiloGram  = Kilo(Gram)
	HectoGram = Hecto(Gram)
	DecaGram  = Deca(Gram)
	DeciGram  = Deci(Gram)
	CentiGram = Centi(Gram)
	MilliGram = Milli(Gram)
	MicroGram = Micro(Gram)
	NanoGram  = Nano(Gram)
	PicoGram  = Pico(Gram)
	FemtoGram = Femto(Gram)
	AttoGram  = Atto(Gram)

	// imperial
	Grain  = NewUnit("grain", "gr", Mass, BI)
	Drachm = NewUnit("drachm", "dr", Mass, BI)
	Ounce  = NewUnit("ounce", "oz", Mass, BI)
	Pound  = NewUnit("pound", "lb", Mass, BI)
	Stone  = NewUnit("stone", "st", Mass, BI)
	Ton    = NewUnit("ton", "t", Mass, BI)
	Slug   = NewUnit("slug", "", Mass, BI)
)

func init() {
	NewRatioConv(Grain, Gram, 0.06479891)
	NewRatioConv(Drachm, Gram, 1.7718451953125)
	NewRatioConv(Ounce, Gram, 28.349523125)
	NewRatioConv(Pound, Gram, 453.59237)
	NewRatioConv(Stone, Gram, 6350.29318)
	NewRatioConv(Ton, Gram, 1016046.9088)
	NewRatioConv(Slug, Gram, 14593.90294)
}
