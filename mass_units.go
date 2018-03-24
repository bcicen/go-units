package units

var (
	Mass = NewQuantity("mass")

	// metric
	Gram      = Mass.NewUnit("gram", "g")
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
	Grain  = Mass.NewUnit("grain", "gr", BI)
	Drachm = Mass.NewUnit("drachm", "dr", BI)
	Ounce  = Mass.NewUnit("ounce", "oz", BI)
	Pound  = Mass.NewUnit("pound", "lb", BI)
	Stone  = Mass.NewUnit("stone", "st", BI)
	Ton    = Mass.NewUnit("ton", "t", BI)
	Slug   = Mass.NewUnit("slug", "", BI)
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
