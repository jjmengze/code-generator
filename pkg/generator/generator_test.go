package generator

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
	"text/template"
)

func TestTemplateWriter_Do(t *testing.T) {
	const templateText = `asldjalksjdlkajsdlk,$.$`

	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"title": strings.Title,
	}

	type fields struct {
		w       io.Writer
		left    string
		right   string
		funcMap template.FuncMap
		err     error
	}
	type args struct {
		format string
		args   interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TemplateWriter
	}{
		{
			name: "test",
			fields: fields{
				w:       os.Stdout,
				left:    "$",
				right:   "$",
				funcMap: funcMap,
			},
			args: args{
				//"func GetOpenAPIDefinitions(ref $.ReferenceCallback|raw$) map[string]$.OpenAPIDefinition|raw$ {\n",
				templateText,
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			},
			want: nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TemplateWriter{
				w:       tt.fields.w,
				left:    tt.fields.left,
				right:   tt.fields.right,
				funcMap: tt.fields.funcMap,
				err:     tt.fields.err,
			}
			if got := s.Do(tt.args.format, tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
