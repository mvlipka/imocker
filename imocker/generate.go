package imocker

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"text/template"
)

// Mock describes an interface and its methods which will be used to generate a mock structure
type Mock struct {
	Package string
	Name    string
	Methods map[string]Method
}

// Method describes a method's parameters and return values
type Method struct {
	NamedParams    []NamedParam
	NamedReturns   []NamedParam
	UnNamedParams  []string
	UnNamedReturns []string
}

// NamedParam describes a parameter or return value's variable name and type definition as a string
type NamedParam struct {
	Name string
	Type string
}

// ParseMock reads a Go source file and builds a Mock object through parsing the file's AST
func ParseMock(reader io.Reader) ([]Mock, error) {
	src, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("error reading from Reader: %w", err)
	}

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error parsing source: %w", err)
	}

	mocks := make([]Mock, 0)
	packageName := f.Name.Name

	// Build up mock objects
	for name, obj := range f.Scope.Objects {

		// check if token is `type`
		switch declType := obj.Decl.(type) {
		case *ast.TypeSpec:

			// Check for an interface else continue
			switch typ := declType.Type.(type) {
			case *ast.InterfaceType:
				mock := Mock{
					Package: packageName,
					Name:    name,
					Methods: make(map[string]Method),
				}

				for _, method := range typ.Methods.List {

					// Check for a method else continue
					switch methodTyp := method.Type.(type) {
					case *ast.FuncType:
						mockMethod := Method{
							NamedParams:    make([]NamedParam, 0),
							NamedReturns:   make([]NamedParam, 0),
							UnNamedParams:  make([]string, 0),
							UnNamedReturns: make([]string, 0),
						}

						for _, param := range methodTyp.Params.List {

							// Check for parameters types else continue
							switch paramTyp := param.Type.(type) {
							case *ast.Ident:
								if len(param.Names) > 0 {
									mockMethod.NamedParams = append(mockMethod.NamedParams, NamedParam{
										Name: param.Names[0].Name,
										Type: paramTyp.Name,
									})
								} else {
									mockMethod.UnNamedParams = append(mockMethod.UnNamedParams, paramTyp.Name)
								}
								break
							default:
								continue
							}
						}

						for _, ret := range methodTyp.Results.List {

							// Check for return types else continue
							switch retTyp := ret.Type.(type) {
							case *ast.Ident:
								if len(ret.Names) > 0 {
									mockMethod.NamedReturns = append(mockMethod.NamedReturns, NamedParam{
										Name: ret.Names[0].Name,
										Type: retTyp.Name,
									})
								} else {
									mockMethod.UnNamedReturns = append(mockMethod.UnNamedReturns, retTyp.Name)
								}
								break
							default:
								continue
							}
						}

						// Add the method to the mock
						mock.Methods[method.Names[0].Name] = mockMethod
						break
					default:
						continue
					}
				}

				mocks = append(mocks, mock)
				break
			default:
				continue
			}
		default:
			continue
		}
	}

	return mocks, nil
}

// GenerateTemplate compiles the template and generates a mock
func GenerateTemplate(mock Mock) (string, error) {
	t := template.Must(template.New("mock").Funcs(templateHelpers).Parse(mockTemplate))

	buf := bytes.NewBufferString("")

	err := t.Execute(buf, mock)
	if err != nil {
		return "", fmt.Errorf("error executing mock template: %w", err)
	}

	return buf.String(), nil
}
