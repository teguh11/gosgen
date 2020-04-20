package tmpl

// StructTmpl for define Struct Name
type StructTmpl struct {
	StructName string
	FieldName  []StructField
}

// StructField for define Struct data
type StructField struct {
	Name     string
	DataType string
}

// StructTemplate is function for generating struct based on StructTmpl and StructField
func StructTemplate() []byte {
	return []byte(`package models

	type {{.StructName}} struct{ {{range .FieldName}}
		{{.Name}} {{.DataType}}{{end}}
	}
`)
}
