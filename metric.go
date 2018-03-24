package units

import (
	"fmt"
	"math"
)

type magnitude struct {
	Symbol string
	Prefix string
	Power  float64
}

var mags = map[string]magnitude{
	"exa":   magnitude{"E", "exa", 18.0},
	"peta":  magnitude{"P", "peta", 15.0},
	"tera":  magnitude{"T", "tera", 12.0},
	"giga":  magnitude{"G", "giga", 9.0},
	"mega":  magnitude{"M", "mega", 6.0},
	"kilo":  magnitude{"k", "kilo", 3.0},
	"hecto": magnitude{"h", "hecto", 2.0},
	"deca":  magnitude{"da", "deca", 1.0},
	"deci":  magnitude{"d", "deci", -1.0},
	"centi": magnitude{"c", "centi", -2.0},
	"milli": magnitude{"m", "milli", -3.0},
	"micro": magnitude{"μ", "micro", -6.0},
	"nano":  magnitude{"n", "nano", -9.0},
	"pico":  magnitude{"p", "pico", -12.0},
	"femto": magnitude{"f", "femto", -15.0},
	"atto":  magnitude{"a", "atto", -18.0},
}

// Magnitude prefix methods create and return a new Unit, while automatically registering
// conversions to and from the provided base Unit
func Exa(b Unit, o ...UnitOption) Unit   { return mags["exa"].makeUnit(b, o...) }
func Peta(b Unit, o ...UnitOption) Unit  { return mags["peta"].makeUnit(b, o...) }
func Tera(b Unit, o ...UnitOption) Unit  { return mags["tera"].makeUnit(b, o...) }
func Giga(b Unit, o ...UnitOption) Unit  { return mags["giga"].makeUnit(b, o...) }
func Mega(b Unit, o ...UnitOption) Unit  { return mags["mega"].makeUnit(b, o...) }
func Kilo(b Unit, o ...UnitOption) Unit  { return mags["kilo"].makeUnit(b, o...) }
func Hecto(b Unit, o ...UnitOption) Unit { return mags["hecto"].makeUnit(b, o...) }
func Deca(b Unit, o ...UnitOption) Unit  { return mags["deca"].makeUnit(b, o...) }
func Deci(b Unit, o ...UnitOption) Unit  { return mags["deci"].makeUnit(b, o...) }
func Centi(b Unit, o ...UnitOption) Unit { return mags["centi"].makeUnit(b, o...) }
func Milli(b Unit, o ...UnitOption) Unit { return mags["milli"].makeUnit(b, o...) }
func Micro(b Unit, o ...UnitOption) Unit { return mags["micro"].makeUnit(b, o...) }
func Nano(b Unit, o ...UnitOption) Unit  { return mags["nano"].makeUnit(b, o...) }
func Pico(b Unit, o ...UnitOption) Unit  { return mags["pico"].makeUnit(b, o...) }
func Femto(b Unit, o ...UnitOption) Unit { return mags["femto"].makeUnit(b, o...) }
func Atto(b Unit, o ...UnitOption) Unit  { return mags["atto"].makeUnit(b, o...) }

// Create magnitude unit and conversion given a base unit
func (mag magnitude) makeUnit(base Unit, addOpts ...UnitOption) Unit {
	name := fmt.Sprintf("%s%s", mag.Prefix, base.Name)
	symbol := fmt.Sprintf("%s%s", mag.Symbol, base.Symbol)

	// set system to metric by default
	opts := []UnitOption{SI}

	// create prefixed aliases if needed
	for _, alias := range base.aliases {
		magAlias := fmt.Sprintf("%s%s", mag.Prefix, alias)
		opts = append(opts, UnitOptionAliases(magAlias))
	}

	// append any supplmental options
	for _, opt := range addOpts {
		opts = append(opts, opt)
	}

	// append quantity name opt
	opts = append(opts, UnitOptionQuantity(base.Quantity))

	u := NewUnit(name, symbol, opts...)

	// only create conversions to and from base unit
	ratio := 1.0 * math.Pow(10.0, mag.Power)
	NewRatioConv(u, base, ratio)

	return u
}
