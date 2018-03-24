package units

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	valuate "github.com/Knetic/govaluate"
	"github.com/bcicen/bfstree"
	"github.com/bcicen/xiny/log"
)

var (
	convs []conversion
	tree  = bfstree.New()
)

type conversionFn func(float64) float64

type conversion struct {
	from    Unit
	to      Unit
	Fn      conversionFn
	Formula string
}

// String representation of conversion formula
func (c conversion) String() string { return c.Formula }

// Conversion implements bfstree.Edge interface
func (c conversion) To() string   { return c.to.Name }
func (c conversion) From() string { return c.from.Name }

// Register a conversion formula and the inverse, given a ratio of
// from Unit in to Unit
func NewRatioConv(from, to Unit, ratio float64) {
	ratioStr := fmt.Sprintf("%.62f", ratio)
	NewConversion(from, to, fmt.Sprintf("x * %s", ratioStr))
	NewConversion(to, from, fmt.Sprintf("x / %s", ratioStr))
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

	c := conversion{from, to, fn, fmtFormula(formula)}
	convs = append(convs, c)
	tree.AddEdge(c)
}

// Replace float in formula string with scientific notation where necessary
func fmtFormula(s string) string {
	re := regexp.MustCompile("(-?[0-9.]+)")
	for _, match := range re.FindAllString(s, -1) {
		f, err := strconv.ParseFloat(match, 64)
		if err != nil {
			return s
		}
		s = strings.Replace(s, match, fmt.Sprintf("%g", f), 1)
	}
	return s
}

// Resolve a path of one or more conversions between two units
func resolveConv(from, to Unit) (fns []conversionFn, err error) {
	path, err := tree.FindPath(from.Name, to.Name)
	if err != nil {
		return fns, err
	}

	formula := ""
	for _, edge := range path.Edges() {
		conv, err := lookupConv(edge.From(), edge.To())
		if err != nil {
			return fns, err
		}
		if formula != "" {
			formula = fmt.Sprintf("(%s)", strings.Replace(conv.Formula, "x", formula, 1))
		} else {
			formula = fmt.Sprintf("(%s)", conv.Formula)
		}
		log.Debugf("%s -> %s: %s", edge.From(), edge.To(), conv.Formula)
		fns = append(fns, conv.Fn)
	}
	log.Infof("%s -> %s: %s", from.Name, to.Name, formula)

	return fns, nil
}

// find conversion function between two units
func lookupConv(from, to string) (c conversion, err error) {
	for _, c := range convs {
		if c.From() == from && c.To() == to {
			return c, nil
		}
	}
	return c, fmt.Errorf("conversion not found")
}
