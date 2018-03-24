package units

var (
	Time = UnitOptionQuantity("time")

	Second      = NewUnit("second", "s", Time)
	ExaSecond   = Exa(Second)
	PetaSecond  = Peta(Second)
	TeraSecond  = Tera(Second)
	GigaSecond  = Giga(Second)
	MegaSecond  = Mega(Second)
	KiloSecond  = Kilo(Second)
	HectoSecond = Hecto(Second)
	DecaSecond  = Deca(Second)
	DeciSecond  = Deci(Second)
	CentiSecond = Centi(Second)
	MilliSecond = Milli(Second)
	MicroSecond = Micro(Second)
	NanoSecond  = Nano(Second)
	PicoSecond  = Pico(Second)
	FemtoSecond = Femto(Second)
	AttoSecond  = Atto(Second)

	Minute = NewUnit("minute", "min", Time)
	Hour   = NewUnit("hour", "hr", Time)
	Day    = NewUnit("day", "d", Time)
	Month  = NewUnit("month", "", Time)
	Year   = NewUnit("year", "yr", Time)

	Decade     = NewUnit("decade", "", Time)
	Century    = NewUnit("century", "", Time)
	Millennium = NewUnit("millennium", "", Time)

	// more esoteric time units
	PlanckTime = NewUnit("planck time", "𝑡ₚ", Time)
	Fortnight  = NewUnit("fortnight", "", Time)
	Score      = NewUnit("score", "", Time)
)

func init() {
	NewRatioConv(Minute, Second, 60.0)
	NewRatioConv(Hour, Second, 3600.0)
	NewRatioConv(Day, Hour, 24.0)
	NewRatioConv(Month, Day, 30.0)
	NewRatioConv(Year, Day, 365.25)

	NewRatioConv(Decade, Year, 10.0)
	NewRatioConv(Century, Year, 100.0)
	NewRatioConv(Millennium, Year, 1000.0)

	NewRatioConv(PlanckTime, Second, 5.39e-44)
	NewRatioConv(Fortnight, Day, 14)
	NewRatioConv(Score, Year, 20.0)
}
