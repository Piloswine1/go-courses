package converter

import ("strings")

type Converter interface {
	IntToRoman(num int) string
}

type converter struct {
	values []int
	symbols []string
}

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

func NewConverter(values []int,symbols []string) Converter {
	return &converter{
		values: values,
		symbols: symbols,
	}
}
