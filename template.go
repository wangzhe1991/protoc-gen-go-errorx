package main

import (
	"bytes"
	"text/template"
)

var errorsTemplate = `
var bizErrorCodeMap map[string]int = map[string]int{
{{ range .Errors }}
		"{{.Name}}_{{.Value}}":{{.BizErrorCode}},
{{- end }}
}
{{ range .Errors }}
var {{.CamelValue}} = gerr.New({{.HTTPCode}}, {{.BizErrorCode}},"{{.Name}}_{{.Value}}", "{{.Message}}")
{{- end }}

{{ range .Errors }}
// Is{{.CamelValue}} {{.Message}}
func Is{{.CamelValue}}(err error) bool {
	if err == nil {
		return false
	}
	e := gerr.FromError(err)
	return e.Reason == "{{.Name}}_{{.Value}}" && e.HttpCode == {{.HTTPCode}}
}

// Error{{.CamelValue}} {{.Message}}
func Error{{.CamelValue}}(format string, args ...interface{}) *gerr.Error {
	 return gerr.New({{.HTTPCode}}, {{.BizErrorCode}}, "{{.Name}}_{{.Value}}", fmt.Sprintf(format, args...))
}
{{- end }}

// BizErrorCode 获取业务code编码
func BizErrorCode(err error) int {
	if err == nil {
		return 0
	}
	e := gerr.FromError(err)
	return bizErrorCodeMap[e.Reason]
}
`

type errorInfo struct {
	Name         string
	Value        string
	HTTPCode     int
	BizErrorCode int
	CamelValue   string
	Message      string
}

type errorWrapper struct {
	Errors []*errorInfo
}

func (e *errorWrapper) generateTemp() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("errors").Parse(errorsTemplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, e); err != nil {
		panic(err)
	}
	return buf.String()
}
