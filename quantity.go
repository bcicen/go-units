package units

var (
	QuantityMap = make(map[string]*Quantity)
)

type Quantity struct {
	Name string
}

func NewQuantity(name string) *Quantity {
	q := &Quantity{Name: name}
	QuantityMap[name] = q
	return q
}

// Return all Units registered to this Quantity
func (q *Quantity) Units() (units []Unit) {
	for _, u := range UnitMap {
		if u.Quantity == q.Name {
			units = append(units, u)
		}
	}
	return units
}

// Create a new Unit within this Quantity and return it
func (q *Quantity) NewUnit(name, symbol string, opts ...UnitOption) Unit {
	return NewUnit(name, symbol, q.Name, opts...)
}
