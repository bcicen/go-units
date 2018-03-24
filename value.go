package units

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var DefaultFmtOptions = FmtOptions{false, 6}

type FmtOptions struct {
	Short     bool // if false, use unit shortname or symbol
	Precision int  // maximum meaningful precision to truncate value
}

type Value struct {
	val  float64
	unit Unit
}

func (v Value) Float() float64 { return v.val }
func (v Value) String() string { return v.Fmt(DefaultFmtOptions) }

func (v Value) Fmt(opts FmtOptions) string {
	var label string

	if opts.Short {
		label = v.unit.Symbol
	} else {
		label = v.unit.Name
		// make label plural if needed
		if v.val > 1.0 {
			label = v.unit.PluralName()
		}
	}

	prec := opts.Precision
	// expand precision if needed to present meaningful value
	if v.val < 1 {
		prec = int((math.Log10(v.val)-0.5)*-1) + prec
	}

	vstr := strconv.FormatFloat(v.val, 'f', prec, 64)
	vstr = trimTrailing(vstr)

	return fmt.Sprintf("%s %s", vstr, label)
}

// Convert this Value to another Unit, returning the new Value
func (v Value) Convert(to Unit) (Value, error) {
	// allow converting to same unit
	if v.unit.Name == to.Name {
		return v, nil
	}

	return Convert(v.val, v.unit, to)
}

// Trim trailing zeros from string
func trimTrailing(s string) string {
	s = strings.TrimRight(s, "0")
	if s[len(s)-1] == '.' {
		s = strings.TrimSuffix(s, ".")
	}
	return s
}