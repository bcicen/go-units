package units

var (
	Speed = UnitOptionQuantity("speed")

	MetersPerSecond   = NewUnit("meter per second", "m/s", Speed, SI, UnitOptionPlural("meters per second"))
	MetersPerHour     = NewUnit("meter per hour", "m/h", Speed, SI, UnitOptionPlural("meters per hour"))
	KilometersPerHour = NewUnit("kilometer per hour", "km/h", Speed, SI, UnitOptionPlural("kilometers per hour"))
	MilesPerHour      = NewUnit("mile per hour", "mph", Speed, BI, UnitOptionPlural("miles per hour"))
	Knot              = NewUnit("knot", "kn", Speed, BI, UnitOptionAliases("nautical miles per hour"))
)

func init() {
	NewRatioConversion(MetersPerSecond, KilometersPerHour, 3.6)
	NewRatioConversion(MetersPerSecond, MetersPerHour, 3600)
	NewRatioConversion(MilesPerHour, KilometersPerHour, 1.609344)
	NewRatioConversion(Knot, KilometersPerHour, 1.852)
}
