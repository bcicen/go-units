package units

var (
	Temp = UnitOptionQuantity("temperature")

	Celsius   = NewUnit("celsius", "C", Temp, UnitOptionPlural("none"), SI)
	Farenheit = NewUnit("farenheit", "F", Temp, UnitOptionPlural("none"), US)
	Kelvin    = NewUnit("kelvin", "K", Temp, UnitOptionPlural("none"), SI)
)

func init() {
	NewConversion(Celsius, Farenheit, "x * 1.8 + 32")
	NewConversion(Farenheit, Celsius, "(x - 32) / 1.8")
	NewConversion(Celsius, Kelvin, "x + 273.15")
	NewConversion(Kelvin, Celsius, "x - 273.15")
}
