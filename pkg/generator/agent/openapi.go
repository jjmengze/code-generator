package agent

import (
	"code-genrtator/pkg/generator"
	"io"
)

type openAPIGen struct {
}

func newOpenAPIGen(sanitizedName string, targetPackage string) *openAPIGen {
	return &openAPIGen{}
}

//func argsFromType(t *types.Type) generator.Args {
//	return generator.Args{
//		"ReferenceCallback": types.Ref(openAPICommonPackagePath, "ReferenceCallback"),
//		"OpenAPIDefinition": types.Ref(openAPICommonPackagePath, "OpenAPIDefinition"),
//		"SpecSchemaType":    types.Ref(specPackagePath, "Schema"),
//	}
//}

func (g *openAPIGen) Init(w io.Writer) error {
	sw := generator.NewTemplateWriter(w, "$", "$")
	sw.Do("func GetOpenAPIDefinitions(ref $.ReferenceCallback|raw$) map[string]$.OpenAPIDefinition|raw$ {\n", nil)

	sw.Do("}\n", nil)
	sw.Do("}\n\n", nil)

	return nil
}
