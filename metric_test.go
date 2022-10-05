package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var magNames = []string{
	"exa",
	"peta",
	"tera",
	"giga",
	"mega",
	"kilo",
	"hecto",
	"deca",
	"deci",
	"centi",
	"milli",
	"micro",
	"nano",
	"pico",
	"femto",
	"atto",
}

type magFn func(Unit, ...UnitOption) Unit

func TestMagnitudes(t *testing.T) {
	u := NewUnit("dong", "₫")
	for i, mfn := range []magFn{Exa, Peta, Tera, Giga, Mega, Kilo, Hecto, Deca, Deci, Centi, Milli, Micro, Nano, Pico, Femto, Atto} {
		mu := mfn(u)
		t.Logf("created mag unit: %s (%s)", mu.Name, mu.Symbol)
		assert.Equal(t, mu.Name, magNames[i]+"dong")
	}
}
