package imocker

const mockTemplate = `
package {{$.Package}}

type Mock{{$.Name}} struct {
	{{range $methodName, $method := .Methods}}
	Test{{$methodName}} func({{range $paramName, $param := $method.NamedParams}}{{$paramName}} {{$param}}{{end}}) ({{range $j, $ret := .UnNamedReturns}}{{$ret}}{{end}})
	{{end}}
}

{{range $methodName, $method := .Methods}}
func (m *Mock{{$.Name}}) {{$methodName}}({{range $paramName, $param := $method.NamedParams}}{{$paramName}} {{$param}}{{end}}) ({{range $j, $ret := .UnNamedReturns}}{{$ret}}{{end}}) {
	return m.Test{{$methodName}}({{range $paramName, $param := $method.NamedParams}}{{$paramName}}{{end}})
}
{{end}}
`
