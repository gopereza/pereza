package boolstub

import "github.com/gopereza/pereza/core/common"

const (
	multiBoolJSONResultHeader    = "return []byte(`{"
	multiBoolJSONResultFooter    = "}`), nil"
	multiBoolJSONResultFixedSize = len(multiBoolJSONResultHeader) + len(multiBoolJSONResultFooter)
)

type DumpGenerator struct {
	jsonNames []string
	buffer    []byte
	last      int
}

func NewDumpGenerator(jsonNames []string) *DumpGenerator {
	length := len(jsonNames)

	commaCount := length - 1
	jsonNameLength := common.StringSliceSize(jsonNames)

	maxCapacity := multiBoolJSONResultFixedSize + jsonNameLength + length*wrapFalse + commaCount

	return &DumpGenerator{
		jsonNames: jsonNames,
		buffer:    make([]byte, 0, maxCapacity),
		last:      len(jsonNames) - 1,
	}
}

func (g *DumpGenerator) Generate(values []bool) []byte {
	result := g.buffer[:0]

	result = append(result, multiBoolJSONResultHeader...)

	for i := 0; i < g.last; i++ {
		result = AppendBool(result, g.jsonNames[i], values[i])

		result = append(result, ',')
	}

	result = AppendBool(result, g.jsonNames[g.last], values[g.last])

	result = append(result, multiBoolJSONResultFooter...)

	return result
}

func AppendBool(source []byte, jsonName string, value bool) []byte {
	result := append(source, '"')
	result = append(result, jsonName...)
	result = append(result, '"', ':')

	if value {
		result = append(result, "true"...)
	} else {
		result = append(result, "false"...)
	}

	return result
}
