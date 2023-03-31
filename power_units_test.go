package units

import (
	"testing"
)

var powerConvTests = []conversionTest{
	{from: "watt", to: "kilowatt", val: "0.001"},
	{from: "watt", to: "megawatt", val: "0.000001"},
	{from: "watt", to: "gigawatt", val: "0.000000001"},
	{from: "watt", to: "terawatt", val: "0.000000000001"},
	{from: "watt", to: "petawatt", val: "0.000000000000001"},
	{from: "watt", to: "exawatt", val: "0.000000000000000001"},
	{from: "watt", to: "zettawatt", val: "0.000000000000000000001"},
	{from: "watt", to: "yottawatt", val: "0.000000000000000000000001"},
}

func TestPower(t *testing.T) {
	testConversions(t, powerConvTests)
}
