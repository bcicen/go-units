package units

import (
	"fmt"
	"sort"
	"strings"
)

var (
	UnitMap = make(map[string]Unit)
)

// Return sorted list of all Unit names and symbols
func Names() (a []string) {
	for _, u := range UnitMap {
		a = append(a, u.Names()...)
	}
	sort.Strings(a)
	return a
}

// Find Unit matching name or symbol provided
func Find(s string) (Unit, error) {

	// first try case-sensitive match
	for _, u := range UnitMap {
		if matchUnit(s, u, true) {
			return u, nil
		}
	}

	// then case-insensitive
	for _, u := range UnitMap {
		if matchUnit(s, u, false) {
			return u, nil
		}
	}

	// finally, try stripping plural suffix
	if strings.HasSuffix(s, "s") || strings.HasSuffix(s, "S") {
		s = strings.TrimSuffix(s, "s")
		s = strings.TrimSuffix(s, "S")
		for _, u := range UnitMap {
			if matchUnit(s, u, false) {
				return u, nil
			}
		}
	}

	return Unit{}, fmt.Errorf("unit \"%s\" not found", s)
}

func matchUnit(s string, u Unit, matchCase bool) bool {
	for _, name := range u.Names() {
		if matchCase {
			if name == s {
				return true
			}
		} else {
			if strings.EqualFold(s, name) {
				return true
			}
		}
	}

	return false
}
