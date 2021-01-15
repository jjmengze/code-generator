package generator

import (
	"fmt"
	"io"
	"runtime"
	"text/template"
)

type TemplateWriter struct {
	w io.Writer
	// Left & right delimiters. text/template defaults to "{{" and "}}"
	// which is totally unusable for go code based templates.
	left, right string
	funcMap     template.FuncMap
	err         error
}

func NewTemplateWriter(w io.Writer, left, right string) *TemplateWriter {
	sw := &TemplateWriter{
		w:       w,
		left:    left,
		right:   right,
		funcMap: template.FuncMap{},
	}

	return sw
}

func (s *TemplateWriter) Do(format string, args interface{}) *TemplateWriter {
	if s.err != nil {
		return s
	}
	// Name the template by source file:line so it can be found when
	// there's an error.
	_, file, line, _ := runtime.Caller(1)
	tmpl, err := template.
		New(fmt.Sprintf("%s:%d", file, line)).
		Delims(s.left, s.right).
		Funcs(s.funcMap).
		Parse(format)
	if err != nil {
		s.err = err
		return s
	}
	err = tmpl.Execute(s.w, args)
	if err != nil {
		s.err = err
	}
	return s
}