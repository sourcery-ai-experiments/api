package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"time"

	rApi "github.com/kloudlite/operator/pkg/operator"
	"github.com/sanity-io/litter"
	apiExtensionsV1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Parser interface {
	GenerateGraphQLSchema(structName string, name string, t reflect.Type) error
	LoadStruct(name string, data any) error

	PrintTypes(w io.Writer)
	PrintCommonTypes(w io.Writer)

	DebugSchema(w io.Writer)
	DumpSchema(dir string) error
	WithPagination(types []string)
}

type GraphqlType string

const (
	Type  GraphqlType = "type"
	Input GraphqlType = "input"
	Enum  GraphqlType = "enum"
)

var scalarMappings = map[reflect.Type]string{
	reflect.TypeOf(metav1.Time{}):      "Date",
	reflect.TypeOf(&metav1.Time{}):     "Date",
	reflect.TypeOf(time.Time{}):        "Date",
	reflect.TypeOf(&time.Time{}):       "Date",
	reflect.TypeOf(json.RawMessage{}):  "Any",
	reflect.TypeOf(&json.RawMessage{}): "Any",
}

var kindMap = map[reflect.Kind]string{
	reflect.Int:   "Int",
	reflect.Int8:  "Int",
	reflect.Int16: "Int",
	reflect.Int32: "Int",
	reflect.Int64: "Int",

	reflect.Uint:   "Int",
	reflect.Uint8:  "Int",
	reflect.Uint16: "Int",
	reflect.Uint32: "Int",
	reflect.Uint64: "Int",

	reflect.Float32: "Float",
	reflect.Float64: "Float",

	reflect.Bool:      "Boolean",
	reflect.Interface: "Any",

	reflect.String: "String",
}

type Struct struct {
	Types  map[string][]string
	Inputs map[string][]string
	Enums  map[string][]string
}

func newStruct() *Struct {
	return &Struct{
		Types:  map[string][]string{},
		Inputs: map[string][]string{},
		Enums:  map[string][]string{},
	}
}

type Field struct {
	ParentName  string
	Name        string
	PkgPath     string
	Type        reflect.Type
	StructName  string
	Fields      *[]string
	InputFields *[]string

	Parser *parser

	JsonTag
	GraphqlTag
}

const (
	commonLabel = "common-types"
)

type parser struct {
	structs map[string]*Struct
	//schemaCli    k8s.ExtendedK8sClient
	schemaCli SchemaClient
}

type JsonTag struct {
	Value     string
	OmitEmpty bool
	Inline    bool
}

func fixPackagePath(pkgPath string) string {
	pkgPath = strings.ReplaceAll(pkgPath, ".", "_")
	pkgPath = strings.ReplaceAll(pkgPath, "/", "__")
	pkgPath = strings.ReplaceAll(pkgPath, "-", "___")

	return pkgPath
}

func parseJsonTag(field reflect.StructField) JsonTag {
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return JsonTag{Value: field.Name, OmitEmpty: false, Inline: false}
	}

	var jt JsonTag
	sp := strings.Split(jsonTag, ",")
	jt.Value = sp[0]

	if jt.Value == "" {
		jt.Value = field.Name
	}

	for i := 1; i < len(sp); i++ {
		if sp[i] == "omitempty" {
			jt.OmitEmpty = true
		}
		if sp[i] == "inline" {
			jt.Inline = true
		}
	}

	return jt
}

type schemaFormat string

type GraphqlTag struct {
	Uri          *string
	Enum         []string
	Ignore       bool
	NoInput      bool
	OnlyInput    bool
	DefaultValue any
}

func parseGraphqlTag(field reflect.StructField) (GraphqlTag, error) {
	tag := field.Tag.Get("graphql")
	if tag == "" {
		return GraphqlTag{}, nil
	}

	var gt GraphqlTag
	sp := strings.Split(tag, ",")
	for i := range sp {
		kv := strings.Split(sp[i], "=")

		switch kv[0] {
		case "uri":
			{
				if len(kv) != 2 {
					return GraphqlTag{}, fmt.Errorf("invalid graphql tag %s, must be of form key=value", tag)
				}
				gt.Uri = &kv[1]
			}
		case "enum":
			{
				if len(kv) != 2 {
					return GraphqlTag{}, fmt.Errorf("invalid graphql tag %s, must be of form key=value", tag)
				}
				enumVals := strings.Split(kv[1], ";")

				gt.Enum = make([]string, 0, len(enumVals))
				for k := range enumVals {
					if enumVals[k] != "" {
						gt.Enum = append(gt.Enum, enumVals[k])
					}
				}
			}
		case "noinput":
			{
				gt.NoInput = true
			}
		case "onlyinput":
			{
				gt.OnlyInput = true
			}

		case "ignore":
			{
				gt.Ignore = true
			}
		case "default":
			{
				if strings.HasPrefix(kv[1], "'") {
					return gt, fmt.Errorf("graphql string value can not start with single-quote, use double-quotes")
				}
				gt.DefaultValue = kv[1]
			}
		default:
			{
				return GraphqlTag{}, fmt.Errorf("unknown graphql tag %s", kv[0])
			}
		}
	}

	return gt, nil
}

func toFieldType(fieldType string, isRequired bool) string {
	if isRequired {
		return fieldType + "!"
	}
	return fieldType
}

func (s *Struct) mergeParser(other *Struct, overKey string) (fields []string, inputFields []string) {
	for k, v := range other.Types {
		if k == overKey {
			fields = append(fields, v...)
			continue
		}
		s.Types[k] = v
	}

	for k, v := range other.Inputs {
		if k == overKey+"In" {
			inputFields = append(inputFields, v...)
			continue
		}
		s.Inputs[k] = v
	}

	for k, v := range other.Enums {
		s.Enums[k] = v
	}

	return fields, inputFields
}

func (p *parser) GenerateGraphQLSchema(structName string, name string, t reflect.Type) error {
	var fields []string
	var inputFields []string

	if _, ok := p.structs[structName]; !ok {
		p.structs[structName] = newStruct()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if !field.IsExported() {
			continue
		}

		jt := parseJsonTag(field)
		if jt.Value == "-" {
			continue
		}

		gt, err := parseGraphqlTag(field)
		if err != nil {
			return err
		}

		if gt.Ignore {
			continue
		}

		fieldType := ""
		inputFieldType := ""

		if scalar, ok := scalarMappings[field.Type]; ok {
			fieldType = toFieldType(scalar, !jt.OmitEmpty)
			inputFieldType = toFieldType(scalar, !jt.OmitEmpty)
		}

		if field.Type.Kind() != reflect.String {
			if v, ok := kindMap[field.Type.Kind()]; ok {
				fieldType = toFieldType(v, !jt.OmitEmpty)
				inputFieldType = toFieldType(v, !jt.OmitEmpty)
			}
		}

		f := Field{
			ParentName:  name,
			Name:        field.Name,
			PkgPath:     field.Type.PkgPath(),
			Type:        field.Type,
			StructName:  structName,
			Fields:      &fields,
			InputFields: &inputFields,
			Parser:      p,
			JsonTag:     jt,
			GraphqlTag:  gt,
		}

		if fieldType == "" {
			switch field.Type.Kind() {
			case reflect.String:
				{
					fieldType, inputFieldType, _ = f.handleString()
				}
			case reflect.Struct:
				{
					fieldType, inputFieldType, _ = f.handleStruct()
				}
			case reflect.Slice:
				{
					fieldType, inputFieldType, _ = f.handleSlice()
				}
			case reflect.Ptr:
				{
					fieldType, inputFieldType, _ = f.handlePtr()
				}
			case reflect.Map:
				{
					fieldType, inputFieldType, _ = f.handleMap()
				}
			default:
				{
					fmt.Printf("default: name: %v (field-name: %v), type: %v, kind: %v\n", jt.Value, field.Name, field.Type, field.Type.Kind())
				}
			}
		}

		if fieldType != "" && !gt.OnlyInput {
			fields = append(fields, fmt.Sprintf("%s: %s", jt.Value, fieldType))
		}
		if inputFieldType != "" && !gt.NoInput {
			if gt.DefaultValue == nil {
				inputFields = append(inputFields, fmt.Sprintf("%s: %s", jt.Value, inputFieldType))
			} else {
				inputFields = append(inputFields, fmt.Sprintf("%s: %s = %v", jt.Value, inputFieldType, gt.DefaultValue))
			}
		}
	}

	if len(fields) > 0 {
		p.structs[structName].Types[name] = fields
	}

	if len(inputFields) > 0 {
		p.structs[structName].Inputs[name+"In"] = inputFields
	}
	return nil
}

func (p *parser) NavigateTree(s *Struct, name string, tree *apiExtensionsV1.JSONSchemaProps, depth ...int) error {
	currDepth := func() int {
		if len(depth) == 0 {
			return 1
		}
		return depth[0]
	}()

	m := map[string]bool{}
	for i := range tree.Required {
		m[tree.Required[i]] = true
	}

	typeName := genTypeName(name)

	fields := make([]string, 0, len(tree.Properties))
	inputFields := make([]string, 0, len(tree.Properties))

	for k, v := range tree.Properties {
		if currDepth == 1 {
			if k == "apiVersion" || k == "kind" {
				// fields = append(fields, genFieldEntry(k, "String", m[k]))
				fields = append(fields, genFieldEntry(k, "String", true))
				inputFields = append(inputFields, genFieldEntry(k, "String", m[k]))
				continue
			}
		}

		if v.Type == "array" {
			if v.Items.Schema != nil && v.Items.Schema.Type == "object" {
				fields = append(fields, genFieldEntry(k, fmt.Sprintf("[%s!]", typeName+genTypeName(k)), m[k]))
				inputFields = append(inputFields, genFieldEntry(k, fmt.Sprintf("[%sIn!]", typeName+genTypeName(k)), m[k]))
				if err := p.NavigateTree(s, typeName+genTypeName(k), v.Items.Schema, currDepth+1); err != nil {
					return err
				}
				continue
			}

			fields = append(fields, genFieldEntry(k, fmt.Sprintf("[%s!]", genTypeName(v.Items.Schema.Type)), m[k]))
			inputFields = append(inputFields, genFieldEntry(k, fmt.Sprintf("[%s!]", genTypeName(v.Items.Schema.Type)), m[k]))
			continue
		}

		if v.Type == "object" {
			if currDepth == 1 {
				// these types are common across all the types that will be generated
				if k == "metadata" {
					fields = append(fields, genFieldEntry(k, "Metadata! @goField(name: \"objectMeta\")", false))
					// fields = append(fields, genFieldEntry(k, "Metadata!", false))
					inputFields = append(inputFields, genFieldEntry(k, "MetadataIn!", false))

					metadata := struct {
						Name              string            `json:"name"`
						Namespace         string            `json:"namespace,omitempty"`
						Labels            map[string]string `json:"labels,omitempty"`
						Annotations       map[string]string `json:"annotations,omitempty"`
						Generation        int64             `json:"generation" graphql:"noinput"`
						CreationTimestamp metav1.Time       `json:"creationTimestamp" graphql:"noinput"`
						DeletionTimestamp *metav1.Time      `json:"deletionTimestamp,omitempty" graphql:"noinput"`
					}{}
					if err := p.GenerateGraphQLSchema(commonLabel, "Metadata", reflect.TypeOf(metadata)); err != nil {
						return err
					}
					continue
				}

				if k == "status" {
					pkgPath := fixPackagePath("github.com/kloudlite/operator/pkg/operator")

					gType := genTypeName(pkgPath + "_" + "Status")

					fields = append(fields, genFieldEntry(k, gType, m[k]))

					p2 := newParser(p.schemaCli)
					if err := p2.GenerateGraphQLSchema(commonLabel, gType, reflect.TypeOf(rApi.Status{})); err != nil {
						return err
					}

					for _, v := range p2.structs {
						for k, v2 := range v.Types {
							p.structs[commonLabel].Types[k] = v2
						}
						for k, v2 := range v.Enums {
							p.structs[commonLabel].Enums[k] = v2
						}
					}

					continue
				}
			}

			if len(v.Properties) == 0 {
				fields = append(fields, genFieldEntry(k, "Map", m[k]))
				inputFields = append(inputFields, genFieldEntry(k, "Map", m[k]))
				continue
			}

			fields = append(fields, genFieldEntry(k, typeName+genTypeName(k), m[k]))
			inputFields = append(inputFields, genFieldEntry(k, typeName+genTypeName(k)+"In", m[k]))
			if err := p.NavigateTree(s, typeName+genTypeName(k), &v, currDepth+1); err != nil {
				return err
			}
			continue
		}

		if v.Type == "string" {
			if len(v.Enum) > 0 {
				fqtn := typeName + genTypeName(k)
				fields = append(fields, genFieldEntry(k, fqtn, m[k]))
				inputFields = append(inputFields, genFieldEntry(k, fqtn, m[k]))

				enums := make([]string, len(v.Enum))
				for i := range v.Enum {
					vjson, _ := v.Enum[i].MarshalJSON()
					var v string
					if err := json.Unmarshal(vjson, &v); err != nil {
						return nil
					}
					enums[i] = v
				}
				s.Enums[fqtn] = enums
				continue
			}
		}

		fields = append(fields, genFieldEntry(k, gqlTypeMap(v.Type), m[k]))
		inputFields = append(inputFields, genFieldEntry(k, gqlTypeMap(v.Type), m[k]))
	}

	s.Types[typeName] = fields
	s.Inputs[typeName+"In"] = inputFields
	return nil
}

func (p *parser) GenerateFromJsonSchema(s *Struct, name string, schema *apiExtensionsV1.JSONSchemaProps) error {
	return p.NavigateTree(s, name, schema)
}

func (p *parser) LoadStruct(name string, data any) error {
	ty := reflect.TypeOf(data)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	return p.GenerateGraphQLSchema(name, name, ty)
}

func (s *Struct) WriteSchema(w io.Writer) {
	keys := make([]string, 0, len(s.Types))
	for k := range s.Types {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for i := range keys {
		io.WriteString(w, fmt.Sprintf("type %s @shareable {\n", keys[i]))
		sort.Slice(s.Types[keys[i]], func(p, q int) bool {
			return strings.ToLower(s.Types[keys[i]][p]) < strings.ToLower(s.Types[keys[i]][q])
		})
		io.WriteString(w, fmt.Sprintf("  %s\n", strings.Join(s.Types[keys[i]], "\n  ")))
		io.WriteString(w, "}\n\n")
	}

	keys = make([]string, 0, len(s.Inputs))
	for k := range s.Inputs {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for i := range keys {
		io.WriteString(w, fmt.Sprintf("input %s {\n", keys[i]))
		sort.Slice(s.Inputs[keys[i]], func(p, q int) bool {
			return strings.ToLower(s.Inputs[keys[i]][p]) < strings.ToLower(s.Inputs[keys[i]][q])
		})
		io.WriteString(w, fmt.Sprintf("  %s\n", strings.Join(s.Inputs[keys[i]], "\n  ")))
		io.WriteString(w, "}\n\n")
	}

	keys = make([]string, 0, len(s.Enums))
	for k := range s.Enums {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for i := range keys {
		io.WriteString(w, fmt.Sprintf("enum %s {\n", keys[i]))
		sort.Slice(s.Enums[keys[i]], func(p, q int) bool {
			return strings.ToLower(s.Enums[keys[i]][p]) < strings.ToLower(s.Enums[keys[i]][q])
		})
		io.WriteString(w, fmt.Sprintf("  %s\n", strings.Join(s.Enums[keys[i]], "\n  ")))
		io.WriteString(w, "}\n\n")
	}
}

func (p *parser) PrintTypes(w io.Writer) {
	keys := make([]string, 0, len(p.structs))
	for k := range p.structs {
		keys = append(keys, k)
	}

	for _, v := range keys {
		if v != commonLabel {
			p.structs[v].WriteSchema(w)
		}
	}
}

func (p *parser) PrintCommonTypes(w io.Writer) {
	if v, ok := p.structs[commonLabel]; ok {
		v.WriteSchema(w)
	}
}

func (p *parser) WithPagination(types []string) {
	for i := range types {
		k := types[i]
		v, ok := p.structs[types[i]]
		if !ok {
			continue
		}

		paginatedTypes := map[string][]string{
			fmt.Sprintf("%sPaginatedRecords", k): {
				"totalCount: Int!",
				fmt.Sprintf("edges: [%sEdge!]!", k),
				fmt.Sprintf("pageInfo: PageInfo!"),
			},
			fmt.Sprintf("%sEdge", k): {
				fmt.Sprintf("node: %v!", k),
				"cursor: String!",
			},
		}

		for i := range paginatedTypes {
			v.Types[i] = paginatedTypes[i]
		}

		if _, ok := p.structs[commonLabel]; !ok {
			p.structs[commonLabel] = newStruct()
		}
	}

	if len(types) > 0 {
		if p.structs[commonLabel] == nil {
			p.structs[commonLabel] = newStruct()
		}

		p.structs[commonLabel].Types["PageInfo"] = []string{
			"hasNextPage: Boolean",
			"hasPreviousPage: Boolean",
			"startCursor: String",
			"endCursor: String",
		}
	}
}

func (p *parser) DebugSchema(w io.Writer) {
	for k, v := range p.structs {
		io.WriteString(w, fmt.Sprintf("struct: %v\n", k))
		io.WriteString(w, litter.Sdump(v))
		io.WriteString(w, "\n")
	}
}

func (p *parser) DumpSchema(dir string) error {
	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err := os.MkdirAll(dir, 0o766); err != nil {
			return err
		}
	}

	for k, v := range p.structs {
		f, err := os.Create(filepath.Join(dir, strings.ToLower(k)+".graphqls"))
		if err != nil {
			return err
		}

		v.WriteSchema(f)
		f.Close()
	}
	return nil
}

type SchemaClient interface {
	GetK8sJsonSchema(name string) (*apiExtensionsV1.JSONSchemaProps, error)
	GetHttpJsonSchema(url string) (*apiExtensionsV1.JSONSchemaProps, error)
}

func newParser(schemaCli SchemaClient) *parser {
	return &parser{
		structs: map[string]*Struct{
			commonLabel: {
				Types:  map[string][]string{},
				Inputs: map[string][]string{},
				Enums:  map[string][]string{},
			},
		},
		schemaCli: schemaCli,
	}
}

func NewParser(cli SchemaClient) Parser {
	return newParser(cli)
}

func NewUnsafeParser(strucs map[string]*Struct, cli SchemaClient) Parser {
	return &parser{
		structs: strucs,
	}
}
