package units

import (
	"fmt"

	"github.com/bcicen/xiny/log"
)

var (
	// Shorthand for pre-defined unit systems
	BI = UnitOptionSystem("imperial")
	SI = UnitOptionSystem("metric")
	US = UnitOptionSystem("us")
)

type Unit struct {
	Name     string
	Symbol   string
	Quantity string
	plural   string // either "none", "auto", or a specific plural name
	aliases  []string
	system   string
}

// NewUnit registers a new Unit within the package, returning the newly created Unit
func NewUnit(name, symbol, quantity string, opts ...UnitOption) Unit {
	if _, ok := UnitMap[name]; ok {
		panic(fmt.Errorf("duplicate unit name: %s", name))
	}

	u := Unit{
		Name:     name,
		Symbol:   symbol,
		Quantity: quantity,
		plural:   "auto",
	}

	for _, opt := range opts {
		u = opt(u)
	}

	UnitMap[name] = u
	log.Debugf("loaded unit %s", name)
	return u
}

// Returns all names and symbols this unit may be referred to
func (u Unit) Names() []string {
	names := []string{u.Name}
	if u.Symbol != "" {
		names = append(names, u.Symbol)
	}
	if u.plural != "none" && u.plural != "auto" {
		names = append(names, u.PluralName())
	}
	return append(names, u.aliases...)
}

// Return the system of units this Unit belongs to, if any
func (u Unit) System() string { return u.system }

// Return the plural name for this unit
func (u Unit) PluralName() string {
	switch u.plural {
	case "none":
		return u.Name
	case "auto":
		return fmt.Sprintf("%ss", u.Name)
	default: // custom plural name
		return u.plural
	}
}

// Return a Value for this Unit
func (u Unit) MakeValue(v float64) Value { return Value{v, u} }

// Option that may be passed to NewUnit
type UnitOption func(Unit) Unit

// Either "none", "auto", or a custom plural unit name
// "none" - labels will use the unmodified unit name in a plural context
// "auto" - labels for this unit will be created with a plural suffix when appropriate (default)
func UnitOptionPlural(s string) UnitOption {
	return func(u Unit) Unit {
		u.plural = s
		return u
	}
}

// Additional names, spellings, or symbols that this unit may be referred to as
func UnitOptionAliases(a ...string) UnitOption {
	return func(u Unit) Unit {
		for _, s := range a {
			u.aliases = append(u.aliases, s)
		}
		return u
	}
}

// Set a system of units for which this Unit belongs
func UnitOptionSystem(s string) UnitOption {
	return func(u Unit) Unit {
		u.system = s
		return u
	}
}
