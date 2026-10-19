package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPressureColumnSymbols(t *testing.T) {
	tests := []struct {
		symbol string
		unit   Unit
	}{
		{"mH2O", MH2O},
		{"mmH2O", MmH2O},
		{"cmH2O", CmH2O},
		{"mHg", MHg},
		{"mmHg", MmHg},
		{"cmHg", CmHg},
	}

	for _, tt := range tests {
		t.Run(tt.symbol, func(t *testing.T) {
			unit, err := Find(tt.symbol)
			if !assert.NoError(t, err) {
				return
			}
			assert.Equal(t, tt.unit, unit)
			assert.Equal(t, "2 "+tt.symbol, NewValue(2, tt.unit).Fmt(FmtOptions{
				Label:     true,
				Short:     true,
				Precision: 3,
			}))
		})
	}
}

func TestPressureColumnConversions(t *testing.T) {
	tests := []struct {
		unit    Unit
		pascals float64
	}{
		{MH2O, 9806.65},
		{MmH2O, 9.80665},
		{CmH2O, 98.0665},
		{MHg, 133322.368421},
		{MmHg, 133.322368421},
		{CmHg, 1333.22368421},
	}

	for _, tt := range tests {
		t.Run(tt.unit.Name, func(t *testing.T) {
			forward, err := ConvertFloat(1, tt.unit, Pascal)
			if assert.NoError(t, err) {
				assert.InEpsilon(t, tt.pascals, forward.Float(), 1e-12)
			}
			reverse, err := ConvertFloat(tt.pascals, Pascal, tt.unit)
			if assert.NoError(t, err) {
				assert.InDelta(t, 1, reverse.Float(), 1e-12)
			}
		})
	}
}

func TestMillimeterMercuryTorrConversion(t *testing.T) {
	tests := []struct {
		from Unit
		to   Unit
	}{
		{MmHg, Torr},
		{Torr, MmHg},
	}

	for _, tt := range tests {
		t.Run(tt.from.Name+" to "+tt.to.Name, func(t *testing.T) {
			value, err := ConvertFloat(760, tt.from, tt.to)
			if assert.NoError(t, err) {
				assert.InDelta(t, 760, value.Float(), 1e-9)
			}
		})
	}
}
