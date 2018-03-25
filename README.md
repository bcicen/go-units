# go-units

Go library for manipulating and converting between various units of measurement

## Usage

Most basic usage example:

```go
package main

import (
	"fmt"
	u "github.com/bcicen/go-units"
)

func main() {
	val, _ := u.ConvertFloat(25.5, u.Celsius, u.Farenheit)
	fmt.Println(val) // outputs "77.9 farenheit"
}
```

labels and plural names for units are also handled:

```go
fmt.Println(u.Bit.MakeValue(1.0)) // "1 bit"
fmt.Println(u.Bit.MakeValue(2.0)) // "2 bits"

// value formatting options may also be specified:
opts := u.FmtOptions{
  Short:     true, // use unit symbol
  Precision: 3,
}

val := u.KiloMeter.MakeValue(15.456932)
fmt.Println(val.Fmt(opts)) // "15.457 km"
fmt.Println(val.Float()) // "15.456932"
```

### Lookup

The package-level `Find()` method may be used to search for a unit by name, symbol, or alternative spelling:
```go
unit, _ := u.Find("m")

unit, _ := u.Find("meter")

unit, _ := u.Find("metre")
```

### Custom Units

`go-units` comes with 260 unit names and symbols builtin; however, new units and conversions can be easily added:

```go
// register custom unit names
Dong := u.NewUnit("dong", "do")
Ding := u.NewUnit("ding", "di")

// there are 100 dongs in a ding
u.NewRatioConv(Ding, Dong, 100.0)

val := Ding.MakeValue(25.0)
newVal, _ := val.Convert(Dong)
fmt.Printf("%s = %s", val, newVal) // "25 dings = 2500 dongs"
```
