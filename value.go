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
	Val  float64
	Unit Unit
}

// Convert this Value to another Unit, returning the new Value
func (v Value) Convert(to Unit) (newVal Value, err error) {
	// allow converting to same unit
	if v.Unit.Name == to.Name {
		return v, nil
	}

	fns, err := resolveConv(v.Unit, to)
	if err != nil {
		return newVal, err
	}

	fVal := v.Val
	for _, fn := range fns {
		fVal = fn(fVal)
	}

	return Value{fVal, to}, nil
}

func (v Value) String() string { return v.Fmt(DefaultFmtOptions) }

func (v Value) Fmt(opts FmtOptions) string {
	var label string

	if opts.Short {
		label = v.Unit.Symbol
	} else {
		label = v.Unit.Name
		// make label plural if needed
		if v.Val > 1.0 {
			label = v.Unit.PluralName()
		}
	}

	prec := opts.Precision
	// expand precision if needed to present meaningful value
	if v.Val < 1 {
		prec = int((math.Log10(v.Val)-0.5)*-1) + prec
	}

	vstr := strconv.FormatFloat(v.Val, 'f', prec, 64)
	vstr = trimTrailing(vstr)

	return fmt.Sprintf("%s %s", vstr, label)
}

// Trim trailing zeros from string
func trimTrailing(s string) string {
	s = strings.TrimRight(s, "0")
	if s[len(s)-1] == '.' {
		s = strings.TrimSuffix(s, ".")
	}
	return s
}
