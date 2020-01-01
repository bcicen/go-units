package units

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	valuate "github.com/Knetic/govaluate"
	"github.com/bcicen/bfstree"
	//"github.com/bcicen/xiny/log"
)

var (
	convs []Conversion
	tree  = bfstree.New()
)

type ConversionFn func(float64) float64

type Conversion struct {
	from    Unit
	to      Unit
	Fn      ConversionFn
	Formula string
}

// String representation of conversion formula
func (c Conversion) String() string { return c.Formula }

// Conversion implements bfstree.Edge interface
func (c Conversion) To() string   { return c.to.Name }
func (c Conversion) From() string { return c.from.Name }

// Register a conversion formula and the inverse, given a ratio of
// from Unit in to Unit
func NewRatioConversion(from, to Unit, ratio float64) {
	ratioStr := fmt.Sprintf("%.62f", ratio)
	NewConversion(from, to, "x * " + ratioStr)
	NewConversion(to, from, "x / " + ratioStr)
}

// NewConversion registers a new conversion formula from one Unit to another
func NewConversion(from, to Unit, formula string) {
	expr, err := valuate.NewEvaluableExpression(formula)
	if err != nil {
		panic(err)
	}

	// create conversion function
	fn := func(x float64) float64 {
		params := make(map[string]interface{})
		params["x"] = x

		res, err := expr.Evaluate(params)
		if err != nil {
			panic(err)
		}
		return res.(float64)
	}

	c := Conversion{from, to, fn, fmtFormula(formula)}
	convs = append(convs, c)
	tree.AddEdge(c)
}

var fmtFormulaRe = regexp.MustCompile("(-?[0-9.]+)")

// Replace float in formula string with scientific notation where necessary
func fmtFormula(s string) string {
	for _, match := range fmtFormulaRe.FindAllString(s, -1) {
		f, err := strconv.ParseFloat(match, 64)
		if err != nil {
			return s
		}
		s = strings.Replace(s, match, fmt.Sprintf("%g", f), 1)
	}
	return s
}

// ResolveConversion resolves a path of one or more Conversions between two units
func ResolveConversion(from, to Unit) (cpath []Conversion, err error) {
	path, err := tree.FindPath(from.Name, to.Name)
	if err != nil {
		return cpath, errors.New("failed to resolve conversion: " + err.Error())
	}

	for _, edge := range path.Edges() {
		conv, err := lookupConv(edge.From(), edge.To())
		if err != nil {
			return cpath, err
		}
		cpath = append(cpath, conv)
	}

	return cpath, nil
}

// find conversion function between two units
func lookupConv(from, to string) (c Conversion, err error) {
	for _, c := range convs {
		if c.From() == from && c.To() == to {
			return c, nil
		}
	}
	return c, errors.New("conversion not found")
}
