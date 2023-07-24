package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// aggregate all unit names, aliases, etc
func aggrNames() (a []string) {
	for _, u := range All() {
		a = append(a, u.Names()...)
	}
	return a
}

// aggregate units by quantity
func aggrByQuantity() map[string][]Unit {
	m := make(map[string][]Unit)

	for _, u := range All() {
		if _, ok := m[u.Quantity]; !ok {
			m[u.Quantity] = []Unit{}
		}
		m[u.Quantity] = append(m[u.Quantity], u)
	}

	return m
}

func TestUnitLookup(t *testing.T) {
	for _, name := range aggrNames() {
		u, err := Find(name)
		if err != nil {
			t.Errorf(err.Error())
			continue
		}
		t.Logf("found unit by name: %s (%s)", name, u.Name)
	}
}

func TestUnitNameOverlap(t *testing.T) {
	nameMap := make(map[string]Unit)

	var total, failed int
	for _, u := range nameMap {
		for _, name := range u.Names() {
			if existing, ok := nameMap[name]; ok {
				t.Errorf("overlap in unit names: %s, %s (%s)", u.Name, existing.Name, name)
				failed++
			} else {
				nameMap[name] = u
			}
			total++
		}
	}
	t.Logf("tested %d unit names, %d overlaps", total, failed)
}

func TestSimilarSymbolLookup(t *testing.T) {
	// The casing of gb could match both gigabyte (GB) and gigabit (Gb)
	symbol := "gb"

	// Run the same assertion multiple times to ensure there is no inconsistency to the results based on random map ordering
	for range make([]bool, 100) {
		u, err := Find(symbol)
		if err != nil {
			t.Errorf("failed to find unit for symbol %s", symbol)
		}
		
		assert.Equal(t, GigaBit, u)
	}
}

// ensure all units within the same quantity resolve
// a conversion path
func TestPathResolve(t *testing.T) {
	for qname, qunits := range aggrByQuantity() {
		t.Logf("testing conversion paths for quantity: %s", qname)
		for _, u1 := range qunits {
			v1 := NewValue(1.0, u1)
			for _, u2 := range qunits {
				if u1.Name == u2.Name {
					continue
				}
				_, err := v1.Convert(u2)
				if err != nil {
					t.Errorf("failed to resolve path: %s -> %s", u1.Name, u2.Name)
				}
			}
		}
	}
}
