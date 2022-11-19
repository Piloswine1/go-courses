package converter

import "strings"

// Simple converter, consisits of translation method
type Converter interface {
	IntToRoman(num int) string
}

type converter struct {
	values  []int    // integer values, e.g. 1, 2, 3...
	symbols []string // corresponding Roman symbol, e.g., I, V, X
}

// Main converting method:
// Going thrue values, since Roman numbers postioned,
// we can subtract int number and add corresponding Roman symbol
func (c *converter) IntToRoman(num int) string {
	var r strings.Builder
	for i, v := range c.values {
		for num >= v {
			r.WriteString(c.symbols[i])
			num -= v
		}
	}

	return r.String()
}

// Constructing new converter
func NewConverter(values []int, symbols []string) Converter {
	return &converter{
		values:  values,
		symbols: symbols,
	}
}
