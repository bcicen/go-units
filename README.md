[![GoDoc](https://godoc.org/github.com/bcicen/go-units?status.svg)](https://godoc.org/github.com/bcicen/go-units)

# go-units

Go library for manipulating and converting between various units of measurement

## Usage

In the most basic usage, `go-units` may be used to convert a value from one unit to another:

```go
package main

import (
	"fmt"
	u "github.com/bcicen/go-units"
)

func main() {
	// convert a simple float from celsius to farenheit
	val, _ := u.ConvertFloat(25.5, u.Celsius, u.Farenheit)
	fmt.Println(val) // outputs "77.9 farenheit"

	// convert a units.Value instance
	val = u.Celsius.MakeValue(25.5)
	fmt.Println(val) // "25.5 celsius"

	val, _ = val.Convert(u.Farenheit)
	fmt.Println(val) // 77.9 farenheit

	val, _ = val.Convert(u.Kelvin)
	fmt.Println(val) // 298.65 kelvin
}
```

### Formatting

Aside from unit conversions, `go-units` may also be used for generating human-readable unit labels, plural names, and symbols:

```go
val := u.NewValue(2.0, u.Nibble)
fmt.Println(val)                     // 2 nibbles
fmt.Println(val.MustConvert(u.Byte)) // 1 byte

// value formatting options may also be specified:
opts := u.FmtOptions{
  Short:     true, // use unit symbol
  Precision: 3,
}

val = u.NewValue(15.456932, u.KiloMeter)
fmt.Println(val)           // 15.456932 kilometers
fmt.Println(val.Fmt(opts)) // 15.457 km
fmt.Println(val.Float())   // 15.456932
```

### Lookup

The package-level `Find()` method may be used to search for a unit by name, symbol, or alternative spelling:
```go
unit, err := u.Find("m")

unit, err := u.Find("meter")

unit, err := u.Find("metre")
```

### Custom Units

`go-units` comes with 260 unit names and symbols builtin; however, new units and conversions can be easily added:

```go
// register custom unit names
Ding := u.NewUnit("ding", "di")
Dong := u.NewUnit("dong", "do")

// there are 100 dongs in a ding
u.NewRatioConversion(Ding, Dong, 100.0)

val := u.NewValue(25.0, Ding)

fmt.Printf("%s = %s\n", val, val.MustConvert(Dong)) // "25 dings = 2500 dongs"

// conversions are automatically registered when using magnitude prefix helper methods
KiloDong := u.Kilo(Dong)
fmt.Println(u.MustConvertFloat(1000.0, Dong, KiloDong)) // "1 kilodong"
```
