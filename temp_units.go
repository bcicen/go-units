package units

var (
	Temp = NewQuantity("temperature")

	Celsius   = Temp.NewUnit("celsius", "C", UnitOptionPlural("none"), SI)
	Farenheit = Temp.NewUnit("farenheit", "F", UnitOptionPlural("none"), US)
	Kelvin    = Temp.NewUnit("kelvin", "K", UnitOptionPlural("none"), SI)
)

func init() {
	NewConversion(Celsius, Farenheit, "x * 1.8 + 32")
	NewConversion(Farenheit, Celsius, "(x - 32) / 1.8")
	NewConversion(Celsius, Kelvin, "x + 273.15")
	NewConversion(Kelvin, Celsius, "x - 273.15")
}
