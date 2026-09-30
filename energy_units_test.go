package units

import (
	"testing"
)

var energyConvTests = []conversionTest{
	{from: "joule", to: "calorie", val: "0.2390057"},
	{from: "joule", to: "electronvolt", val: "6241495961752113152"},
	{from: "joule", to: "megawatt-hour", val: "0.0000000002777778"},
	{from: "joule", to: "gigawatt-hour", val: "0.0000000000002777778"},
	{from: "joule", to: "kilowatt-hour", val: "0.0000002777778"},
	{from: "joule", to: "megajoule", val: "0.000001"},
	{from: "joule", to: "gigajoule", val: "0.000000001"},
	{from: "joule", to: "terajoule", val: "0.000000000001"},
	{from: "joule", to: "petajoule", val: "0.000000000000001"},
	{from: "joule", to: "exajoule", val: "0.000000000000000001"},
	{from: "joule", to: "zettajoule", val: "0.000000000000000000001"},
	{from: "joule", to: "yottajoule", val: "0.000000000000000000000001"},
	{from: "calorie", to: "joule", val: "4.184"},
	{from: "calorie", to: "electronvolt", val: "26114419103970840576"},
	{from: "calorie", to: "megawatt-hour", val: "0.000000001162222"},
	{from: "calorie", to: "gigawatt-hour", val: "0.000000000001162222"},
	{from: "calorie", to: "kilowatt-hour", val: "0.000001162222"},
	{from: "calorie", to: "megajoule", val: "0.000004184"},
	{from: "calorie", to: "gigajoule", val: "0.000000004184"},
	{from: "calorie", to: "terajoule", val: "0.000000000004184"},
	{from: "calorie", to: "petajoule", val: "0.000000000000004184"},
	{from: "calorie", to: "exajoule", val: "0.000000000000000004184"},
	{from: "calorie", to: "zettajoule", val: "0.000000000000000000004184"},
	{from: "calorie", to: "yottajoule", val: "0.000000000000000000000004184"},
	{from: "electronvolt", to: "joule", val: "0.000000000000000000160218"},
	{from: "electronvolt", to: "calorie", val: "0.000000000000000000038293"},
	{from: "electronvolt", to: "megawatt-hour", val: "0.000000000000000000000000000044505"},
	{from: "electronvolt", to: "gigawatt-hour", val: "0.000000000000000000000000000000044505"},
	{from: "electronvolt", to: "kilowatt-hour", val: "0.000000000000000000000000044505"},
	{from: "electronvolt", to: "megajoule", val: "0.000000000000000000000000160218"},
	{from: "electronvolt", to: "gigajoule", val: "0.000000000000000000000000000160218"},
	{from: "electronvolt", to: "terajoule", val: "0.000000000000000000000000000000160218"},
	{from: "electronvolt", to: "petajoule", val: "0.000000000000000000000000000000000160218"},
	{from: "electronvolt", to: "exajoule", val: "0.000000000000000000000000000000000000160218"},
	{from: "electronvolt", to: "zettajoule", val: "0.000000000000000000000000000000000000000160218"},
	{from: "electronvolt", to: "yottajoule", val: "0.000000000000000000000000000000000000000000160218"},
	{from: "megawatt-hour", to: "joule", val: "3600000000"},
}

func TestEnergy(t *testing.T) {
	testConversions(t, energyConvTests)
}

func TestWattHourPrefixesUseJoules(t *testing.T) {
	cases := []struct {
		from   Unit
		joules float64
	}{
		{KiloWattHour, 3.6e6},
		{MegaWattHour, 3.6e9},
		{GigaWattHour, 3.6e12},
		{TeraWattHour, 3.6e15},
		{PetaWattHour, 3.6e18},
	}
	for _, tc := range cases {
		got := MustConvertFloat(1, tc.from, Joule).Float()
		if got != tc.joules {
			t.Errorf("%s: got %v, want %v joules", tc.from.Name, got, tc.joules)
		}
	}

	if _, err := ConvertFloat(1, KiloWattHour, Watt); err == nil {
		t.Fatal("kilowatt-hour converted to watt")
	}
}
