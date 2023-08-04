package imocker

import "text/template"

var templateHelpers = template.FuncMap{
	"isLastElement": func(i int, len int) bool {
		return i+1 == len
	},
}

const mockTemplate = `
package {{$.Package}}

type Mock{{$.Name}} struct {
	{{range $methodName, $method := .Methods}}
	Test{{$methodName}} func({{range $paramIndex, $param := $method.NamedParams}}{{$param.Name}} {{$param.Type}}{{if not (isLastElement $paramIndex (len $method.NamedParams))}},{{end}}{{end}}) ({{if .UnNamedReturns}}{{range $retIndex, $ret := $method.UnNamedReturns}}{{$ret}}{{if not (isLastElement $retIndex (len $method.UnNamedReturns))}},{{end}}{{end}}{{else}}{{range $retIndex, $ret := .NamedReturns}}{{$ret.Name}} {{$ret.Type}}{{end}}{{end}})
	{{end}}
}

{{range $methodName, $method := .Methods}}
	func (m *Mock{{$.Name}}) {{$methodName}}({{range $paramIndex, $param := $method.NamedParams}}{{$param.Name}} {{$param.Type}}{{if not (isLastElement $paramIndex (len $method.NamedParams))}}, {{end}}{{end}}) ({{if .UnNamedReturns}}{{range $retIndex, $ret := $method.UnNamedReturns}}{{$ret}}{{if not (isLastElement $retIndex (len $method.UnNamedReturns))}},{{end}}{{end}}{{else}}{{range $retIndex, $ret := .NamedReturns}}{{$ret.Name}} {{$ret.Type}}{{if not (isLastElement $retIndex (len $method.NamedReturns))}},{{end}}{{end}}{{end}}) {
			return m.Test{{$methodName}}({{range $paramIndex, $param := $method.NamedParams}}{{$param.Name}}{{if not (isLastElement $paramIndex (len $method.NamedParams))}}, {{end}}{{end}})
	}
{{end}}
`
