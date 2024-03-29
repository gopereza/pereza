package gen

import (
	"github.com/gopereza/pereza/core"
	"github.com/gopereza/pereza/core/boolstub"
	"github.com/gopereza/pereza/core/common"
	"github.com/gopereza/pereza/core/complexstub"
	"github.com/gopereza/pereza/core/intstub"
	"github.com/gopereza/pereza/core/stringstub"
	"io"
	"reflect"
)

type Generator struct {
	packagePath string
	packageName string
	hashString  string

	types []reflect.Type
}

func NewGenerator(packagePath, packageName, filename string) *Generator {
	ret := &Generator{
		packagePath: packagePath,
		packageName: packageName,
		hashString:  unique(filename),

		types: make([]reflect.Type, 0, 1),
	}

	return ret
}

func (g *Generator) Add(obj interface{}) {
	t := reflect.TypeOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	g.types = append(g.types, t)
}

// Run runs the generator and outputs generated code to out.
func (g *Generator) Run(out io.Writer) error {
	out.Write(g.header())

	body := make([]byte, 0)
	imports := make([]string, 0)

	for _, t := range g.types {
		typeBody, typeImports := g.genStructEncoder(t)

		body = append(body, typeBody...)
		imports = append(imports, typeImports...)
	}

	if len(imports) > 0 {
		result := common.AppendImports(nil, imports)

		out.Write(result)
		out.Write([]byte{'\n', '\n'})
	}

	out.Write(body)

	return nil
}

// header prints package declaration and imports.
func (g *Generator) header() []byte {
	const header = `// Code generated by pereza for marshaling/unmarshaling. DO NOT EDIT.

package `

	var result []byte

	result = append(result, header...)
	result = append(result, g.packageName...)
	result = append(result, '\n', '\n')

	return result
}

func (g *Generator) genStructEncoder(t reflect.Type) ([]byte, []string) {
	length := t.NumField()

	if length == 0 {
		return common.EmptyResultStub(t.Name()), nil
	}

	switch length {
	case 1:
		field := t.Field(0)

		jsonName, standard := core.StandardStructureField(field)

		if standard {
			kind := field.Type.Kind()

			switch kind {
			case reflect.Bool:
				return boolstub.OneFieldStub(t.Name(), field.Name, jsonName), nil
			case reflect.String:
				return stringstub.StringResultStub(t.Name(), field.Name, jsonName), nil
			case reflect.Int,
				reflect.Int8,
				reflect.Int16,
				reflect.Int32,
				reflect.Int64,
				reflect.Uint,
				reflect.Uint8,
				reflect.Uint16,
				reflect.Uint32,
				reflect.Uint64:

				return intstub.IntResultStubByType(t.Name(), field.Name, jsonName, kind), []string{intstub.IntImport}
			}
		}
	default:
		fieldsNames, jsonNames, standard := core.MultiBoolStandardStructure(t)

		if standard {
			if core.MatchAllBooleanFields(t) {
				if len(fieldsNames) > core.MultiBoolMaxProperties {
					return boolstub.LargeFieldStub(t.Name(), fieldsNames, jsonNames), nil
				}

				return boolstub.CombinatorBoolResultStub(t.Name(), fieldsNames, jsonNames), nil
			}

			return complexstub.StandardStub(t, fieldsNames, jsonNames)
		}
	}

	return common.EmptyResultStub(t.Name()), nil
}
