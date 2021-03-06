// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protoc-gen-swagger/options/openapiv2.proto

/*
Package options is a generated protocol buffer package.

It is generated from these files:
	protoc-gen-swagger/options/openapiv2.proto
	protoc-gen-swagger/options/annotations.proto

It has these top-level messages:
	Swagger
	Operation
	Info
	Contact
	ExternalDocumentation
	Schema
	JSONSchema
	Tag
*/
package options

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Swagger_SwaggerScheme int32

const (
	Swagger_UNKNOWN Swagger_SwaggerScheme = 0
	Swagger_HTTP    Swagger_SwaggerScheme = 1
	Swagger_HTTPS   Swagger_SwaggerScheme = 2
	Swagger_WS      Swagger_SwaggerScheme = 3
	Swagger_WSS     Swagger_SwaggerScheme = 4
)

var Swagger_SwaggerScheme_name = map[int32]string{
	0: "UNKNOWN",
	1: "HTTP",
	2: "HTTPS",
	3: "WS",
	4: "WSS",
}
var Swagger_SwaggerScheme_value = map[string]int32{
	"UNKNOWN": 0,
	"HTTP":    1,
	"HTTPS":   2,
	"WS":      3,
	"WSS":     4,
}

func (x Swagger_SwaggerScheme) String() string {
	return proto.EnumName(Swagger_SwaggerScheme_name, int32(x))
}
func (Swagger_SwaggerScheme) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type JSONSchema_JSONSchemaSimpleTypes int32

const (
	JSONSchema_UNKNOWN JSONSchema_JSONSchemaSimpleTypes = 0
	JSONSchema_ARRAY   JSONSchema_JSONSchemaSimpleTypes = 1
	JSONSchema_BOOLEAN JSONSchema_JSONSchemaSimpleTypes = 2
	JSONSchema_INTEGER JSONSchema_JSONSchemaSimpleTypes = 3
	JSONSchema_NULL    JSONSchema_JSONSchemaSimpleTypes = 4
	JSONSchema_NUMBER  JSONSchema_JSONSchemaSimpleTypes = 5
	JSONSchema_OBJECT  JSONSchema_JSONSchemaSimpleTypes = 6
	JSONSchema_STRING  JSONSchema_JSONSchemaSimpleTypes = 7
)

var JSONSchema_JSONSchemaSimpleTypes_name = map[int32]string{
	0: "UNKNOWN",
	1: "ARRAY",
	2: "BOOLEAN",
	3: "INTEGER",
	4: "NULL",
	5: "NUMBER",
	6: "OBJECT",
	7: "STRING",
}
var JSONSchema_JSONSchemaSimpleTypes_value = map[string]int32{
	"UNKNOWN": 0,
	"ARRAY":   1,
	"BOOLEAN": 2,
	"INTEGER": 3,
	"NULL":    4,
	"NUMBER":  5,
	"OBJECT":  6,
	"STRING":  7,
}

func (x JSONSchema_JSONSchemaSimpleTypes) String() string {
	return proto.EnumName(JSONSchema_JSONSchemaSimpleTypes_name, int32(x))
}
func (JSONSchema_JSONSchemaSimpleTypes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

// `Swagger` is a representation of OpenAPI v2 specification's Swagger object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#swaggerObject
//
// TODO(ivucica): document fields
type Swagger struct {
	Swagger      string                  `protobuf:"bytes,1,opt,name=swagger" json:"swagger,omitempty"`
	Info         *Info                   `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	Host         string                  `protobuf:"bytes,3,opt,name=host" json:"host,omitempty"`
	BasePath     string                  `protobuf:"bytes,4,opt,name=base_path,json=basePath" json:"base_path,omitempty"`
	Schemes      []Swagger_SwaggerScheme `protobuf:"varint,5,rep,packed,name=schemes,enum=grpc.gateway.protoc_gen_swagger.options.Swagger_SwaggerScheme" json:"schemes,omitempty"`
	Consumes     []string                `protobuf:"bytes,6,rep,name=consumes" json:"consumes,omitempty"`
	Produces     []string                `protobuf:"bytes,7,rep,name=produces" json:"produces,omitempty"`
	ExternalDocs *ExternalDocumentation  `protobuf:"bytes,14,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
}

func (m *Swagger) Reset()                    { *m = Swagger{} }
func (m *Swagger) String() string            { return proto.CompactTextString(m) }
func (*Swagger) ProtoMessage()               {}
func (*Swagger) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Swagger) GetSwagger() string {
	if m != nil {
		return m.Swagger
	}
	return ""
}

func (m *Swagger) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Swagger) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Swagger) GetBasePath() string {
	if m != nil {
		return m.BasePath
	}
	return ""
}

func (m *Swagger) GetSchemes() []Swagger_SwaggerScheme {
	if m != nil {
		return m.Schemes
	}
	return nil
}

func (m *Swagger) GetConsumes() []string {
	if m != nil {
		return m.Consumes
	}
	return nil
}

func (m *Swagger) GetProduces() []string {
	if m != nil {
		return m.Produces
	}
	return nil
}

func (m *Swagger) GetExternalDocs() *ExternalDocumentation {
	if m != nil {
		return m.ExternalDocs
	}
	return nil
}

// `Operation` is a representation of OpenAPI v2 specification's Operation object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#operationObject
//
// TODO(ivucica): document fields
type Operation struct {
	Tags         []string               `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty"`
	Summary      string                 `protobuf:"bytes,2,opt,name=summary" json:"summary,omitempty"`
	Description  string                 `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `protobuf:"bytes,4,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
	OperationId  string                 `protobuf:"bytes,5,opt,name=operation_id,json=operationId" json:"operation_id,omitempty"`
	Consumes     []string               `protobuf:"bytes,6,rep,name=consumes" json:"consumes,omitempty"`
	Produces     []string               `protobuf:"bytes,7,rep,name=produces" json:"produces,omitempty"`
	Schemes      []string               `protobuf:"bytes,10,rep,name=schemes" json:"schemes,omitempty"`
	Deprecated   bool                   `protobuf:"varint,11,opt,name=deprecated" json:"deprecated,omitempty"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Operation) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Operation) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *Operation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Operation) GetExternalDocs() *ExternalDocumentation {
	if m != nil {
		return m.ExternalDocs
	}
	return nil
}

func (m *Operation) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

func (m *Operation) GetConsumes() []string {
	if m != nil {
		return m.Consumes
	}
	return nil
}

func (m *Operation) GetProduces() []string {
	if m != nil {
		return m.Produces
	}
	return nil
}

func (m *Operation) GetSchemes() []string {
	if m != nil {
		return m.Schemes
	}
	return nil
}

func (m *Operation) GetDeprecated() bool {
	if m != nil {
		return m.Deprecated
	}
	return false
}

// `Info` is a representation of OpenAPI v2 specification's Info object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#infoObject
//
// TODO(ivucica): document fields
type Info struct {
	Title          string   `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Description    string   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	TermsOfService string   `protobuf:"bytes,3,opt,name=terms_of_service,json=termsOfService" json:"terms_of_service,omitempty"`
	Contact        *Contact `protobuf:"bytes,4,opt,name=contact" json:"contact,omitempty"`
	Version        string   `protobuf:"bytes,6,opt,name=version" json:"version,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Info) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Info) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Info) GetTermsOfService() string {
	if m != nil {
		return m.TermsOfService
	}
	return ""
}

func (m *Info) GetContact() *Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

func (m *Info) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// `Contact` is a representation of OpenAPI v2 specification's Contact object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#contactObject
//
// TODO(ivucica): document fields
type Contact struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Url   string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Contact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contact) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Contact) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// `ExternalDocumentation` is a representation of OpenAPI v2 specification's
// ExternalDocumentation object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#externalDocumentationObject
//
// TODO(ivucica): document fields
type ExternalDocumentation struct {
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	Url         string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *ExternalDocumentation) Reset()                    { *m = ExternalDocumentation{} }
func (m *ExternalDocumentation) String() string            { return proto.CompactTextString(m) }
func (*ExternalDocumentation) ProtoMessage()               {}
func (*ExternalDocumentation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ExternalDocumentation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ExternalDocumentation) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// `Schema` is a representation of OpenAPI v2 specification's Schema object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#schemaObject
//
// TODO(ivucica): document fields
type Schema struct {
	JsonSchema    *JSONSchema            `protobuf:"bytes,1,opt,name=json_schema,json=jsonSchema" json:"json_schema,omitempty"`
	Discriminator string                 `protobuf:"bytes,2,opt,name=discriminator" json:"discriminator,omitempty"`
	ReadOnly      bool                   `protobuf:"varint,3,opt,name=read_only,json=readOnly" json:"read_only,omitempty"`
	ExternalDocs  *ExternalDocumentation `protobuf:"bytes,5,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
	Example       *google_protobuf.Any   `protobuf:"bytes,6,opt,name=example" json:"example,omitempty"`
}

func (m *Schema) Reset()                    { *m = Schema{} }
func (m *Schema) String() string            { return proto.CompactTextString(m) }
func (*Schema) ProtoMessage()               {}
func (*Schema) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Schema) GetJsonSchema() *JSONSchema {
	if m != nil {
		return m.JsonSchema
	}
	return nil
}

func (m *Schema) GetDiscriminator() string {
	if m != nil {
		return m.Discriminator
	}
	return ""
}

func (m *Schema) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

func (m *Schema) GetExternalDocs() *ExternalDocumentation {
	if m != nil {
		return m.ExternalDocs
	}
	return nil
}

func (m *Schema) GetExample() *google_protobuf.Any {
	if m != nil {
		return m.Example
	}
	return nil
}

// `JSONSchema` represents properties from JSON Schema taken, and as used, in
// the OpenAPI v2 spec.
//
// This includes changes made by OpenAPI v2.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#schemaObject
//
// See also: https://cswr.github.io/JsonSchema/spec/basic_types/,
// https://github.com/json-schema-org/json-schema-spec/blob/master/schema.json
//
// TODO(ivucica): document fields
type JSONSchema struct {
	Title            string   `protobuf:"bytes,5,opt,name=title" json:"title,omitempty"`
	Description      string   `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Default          string   `protobuf:"bytes,7,opt,name=default" json:"default,omitempty"`
	MultipleOf       float64  `protobuf:"fixed64,10,opt,name=multiple_of,json=multipleOf" json:"multiple_of,omitempty"`
	Maximum          float64  `protobuf:"fixed64,11,opt,name=maximum" json:"maximum,omitempty"`
	ExclusiveMaximum bool     `protobuf:"varint,12,opt,name=exclusive_maximum,json=exclusiveMaximum" json:"exclusive_maximum,omitempty"`
	Minimum          float64  `protobuf:"fixed64,13,opt,name=minimum" json:"minimum,omitempty"`
	ExclusiveMinimum bool     `protobuf:"varint,14,opt,name=exclusive_minimum,json=exclusiveMinimum" json:"exclusive_minimum,omitempty"`
	MaxLength        uint64   `protobuf:"varint,15,opt,name=max_length,json=maxLength" json:"max_length,omitempty"`
	MinLength        uint64   `protobuf:"varint,16,opt,name=min_length,json=minLength" json:"min_length,omitempty"`
	Pattern          string   `protobuf:"bytes,17,opt,name=pattern" json:"pattern,omitempty"`
	MaxItems         uint64   `protobuf:"varint,20,opt,name=max_items,json=maxItems" json:"max_items,omitempty"`
	MinItems         uint64   `protobuf:"varint,21,opt,name=min_items,json=minItems" json:"min_items,omitempty"`
	UniqueItems      bool     `protobuf:"varint,22,opt,name=unique_items,json=uniqueItems" json:"unique_items,omitempty"`
	MaxProperties    uint64   `protobuf:"varint,24,opt,name=max_properties,json=maxProperties" json:"max_properties,omitempty"`
	MinProperties    uint64   `protobuf:"varint,25,opt,name=min_properties,json=minProperties" json:"min_properties,omitempty"`
	Required         []string `protobuf:"bytes,26,rep,name=required" json:"required,omitempty"`
	// Items in 'array' must be unique.
	Array []string                           `protobuf:"bytes,34,rep,name=array" json:"array,omitempty"`
	Type  []JSONSchema_JSONSchemaSimpleTypes `protobuf:"varint,35,rep,packed,name=type,enum=grpc.gateway.protoc_gen_swagger.options.JSONSchema_JSONSchemaSimpleTypes" json:"type,omitempty"`
}

func (m *JSONSchema) Reset()                    { *m = JSONSchema{} }
func (m *JSONSchema) String() string            { return proto.CompactTextString(m) }
func (*JSONSchema) ProtoMessage()               {}
func (*JSONSchema) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *JSONSchema) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *JSONSchema) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *JSONSchema) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *JSONSchema) GetMultipleOf() float64 {
	if m != nil {
		return m.MultipleOf
	}
	return 0
}

func (m *JSONSchema) GetMaximum() float64 {
	if m != nil {
		return m.Maximum
	}
	return 0
}

func (m *JSONSchema) GetExclusiveMaximum() bool {
	if m != nil {
		return m.ExclusiveMaximum
	}
	return false
}

func (m *JSONSchema) GetMinimum() float64 {
	if m != nil {
		return m.Minimum
	}
	return 0
}

func (m *JSONSchema) GetExclusiveMinimum() bool {
	if m != nil {
		return m.ExclusiveMinimum
	}
	return false
}

func (m *JSONSchema) GetMaxLength() uint64 {
	if m != nil {
		return m.MaxLength
	}
	return 0
}

func (m *JSONSchema) GetMinLength() uint64 {
	if m != nil {
		return m.MinLength
	}
	return 0
}

func (m *JSONSchema) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *JSONSchema) GetMaxItems() uint64 {
	if m != nil {
		return m.MaxItems
	}
	return 0
}

func (m *JSONSchema) GetMinItems() uint64 {
	if m != nil {
		return m.MinItems
	}
	return 0
}

func (m *JSONSchema) GetUniqueItems() bool {
	if m != nil {
		return m.UniqueItems
	}
	return false
}

func (m *JSONSchema) GetMaxProperties() uint64 {
	if m != nil {
		return m.MaxProperties
	}
	return 0
}

func (m *JSONSchema) GetMinProperties() uint64 {
	if m != nil {
		return m.MinProperties
	}
	return 0
}

func (m *JSONSchema) GetRequired() []string {
	if m != nil {
		return m.Required
	}
	return nil
}

func (m *JSONSchema) GetArray() []string {
	if m != nil {
		return m.Array
	}
	return nil
}

func (m *JSONSchema) GetType() []JSONSchema_JSONSchemaSimpleTypes {
	if m != nil {
		return m.Type
	}
	return nil
}

// `Tag` is a representation of OpenAPI v2 specification's Tag object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#tagObject
//
// TODO(ivucica): document fields
type Tag struct {
	// TODO(ivucica): Description should be extracted from comments on the proto
	// service object.
	Description  string                 `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `protobuf:"bytes,3,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Tag) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Tag) GetExternalDocs() *ExternalDocumentation {
	if m != nil {
		return m.ExternalDocs
	}
	return nil
}

func init() {
	proto.RegisterType((*Swagger)(nil), "grpc.gateway.protoc_gen_swagger.options.Swagger")
	proto.RegisterType((*Operation)(nil), "grpc.gateway.protoc_gen_swagger.options.Operation")
	proto.RegisterType((*Info)(nil), "grpc.gateway.protoc_gen_swagger.options.Info")
	proto.RegisterType((*Contact)(nil), "grpc.gateway.protoc_gen_swagger.options.Contact")
	proto.RegisterType((*ExternalDocumentation)(nil), "grpc.gateway.protoc_gen_swagger.options.ExternalDocumentation")
	proto.RegisterType((*Schema)(nil), "grpc.gateway.protoc_gen_swagger.options.Schema")
	proto.RegisterType((*JSONSchema)(nil), "grpc.gateway.protoc_gen_swagger.options.JSONSchema")
	proto.RegisterType((*Tag)(nil), "grpc.gateway.protoc_gen_swagger.options.Tag")
	proto.RegisterEnum("grpc.gateway.protoc_gen_swagger.options.Swagger_SwaggerScheme", Swagger_SwaggerScheme_name, Swagger_SwaggerScheme_value)
	proto.RegisterEnum("grpc.gateway.protoc_gen_swagger.options.JSONSchema_JSONSchemaSimpleTypes", JSONSchema_JSONSchemaSimpleTypes_name, JSONSchema_JSONSchemaSimpleTypes_value)
}

func init() { proto.RegisterFile("protoc-gen-swagger/options/openapiv2.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x8e, 0x1b, 0x35,
	0x14, 0x66, 0x12, 0x27, 0x71, 0x4e, 0x92, 0xc5, 0x9d, 0x6e, 0xc1, 0xdd, 0xfe, 0x90, 0x86, 0x22,
	0xa2, 0x56, 0x9b, 0xa0, 0xed, 0x25, 0x02, 0x69, 0xb7, 0x44, 0x65, 0xa7, 0xdb, 0xa4, 0x9a, 0xa4,
	0x2a, 0x20, 0xa1, 0xc8, 0x3b, 0x71, 0x66, 0x0d, 0x33, 0x9e, 0xe9, 0xfc, 0x6c, 0x93, 0xd7, 0xe0,
	0x79, 0x78, 0x0c, 0xde, 0x81, 0x37, 0xe0, 0x82, 0x2b, 0x64, 0x8f, 0x67, 0xff, 0x52, 0xd0, 0xaa,
	0x2a, 0x57, 0xf1, 0xf9, 0xbe, 0x73, 0x3e, 0xfb, 0xfc, 0x78, 0x1c, 0x78, 0x14, 0x27, 0x51, 0x16,
	0x79, 0xbb, 0x3e, 0x97, 0xbb, 0xe9, 0x5b, 0xe6, 0xfb, 0x3c, 0x19, 0x46, 0x71, 0x26, 0x22, 0x99,
	0x0e, 0xa3, 0x98, 0x4b, 0x16, 0x8b, 0xd3, 0xbd, 0x81, 0x76, 0xb2, 0xbf, 0xf4, 0x93, 0xd8, 0x1b,
	0xf8, 0x2c, 0xe3, 0x6f, 0xd9, 0xba, 0xc0, 0xbc, 0xb9, 0xcf, 0xe5, 0xdc, 0x04, 0x0e, 0x4c, 0xe0,
	0xce, 0x6d, 0x3f, 0x8a, 0xfc, 0x80, 0x0f, 0xb5, 0xcb, 0x71, 0xbe, 0x1c, 0x32, 0x69, 0xfc, 0x7b,
	0x7f, 0x55, 0xa1, 0x31, 0x2d, 0xdc, 0x6d, 0x0a, 0x0d, 0x13, 0x49, 0xad, 0xae, 0xd5, 0x6f, 0xba,
	0xa5, 0x69, 0xef, 0x03, 0x12, 0x72, 0x19, 0xd1, 0x4a, 0xd7, 0xea, 0xb7, 0xf6, 0x76, 0x07, 0xd7,
	0xdc, 0x78, 0x70, 0x28, 0x97, 0x91, 0xab, 0x43, 0x6d, 0x1b, 0xd0, 0x49, 0x94, 0x66, 0xb4, 0xaa,
	0x95, 0xf5, 0xda, 0xbe, 0x03, 0xcd, 0x63, 0x96, 0xf2, 0x79, 0xcc, 0xb2, 0x13, 0x8a, 0x34, 0x81,
	0x15, 0xf0, 0x92, 0x65, 0x27, 0xf6, 0x0f, 0xd0, 0x48, 0xbd, 0x13, 0x1e, 0xf2, 0x94, 0xd6, 0xba,
	0xd5, 0xfe, 0xd6, 0xde, 0xb7, 0xd7, 0xde, 0xd6, 0x24, 0x54, 0xfe, 0x4e, 0xb5, 0x8c, 0x5b, 0xca,
	0xd9, 0x3b, 0x80, 0xbd, 0x48, 0xa6, 0xb9, 0x92, 0xae, 0x77, 0xab, 0x6a, 0xd7, 0xd2, 0x56, 0x5c,
	0x9c, 0x44, 0x8b, 0xdc, 0xe3, 0x29, 0x6d, 0x14, 0x5c, 0x69, 0xdb, 0x1e, 0x74, 0xf8, 0x2a, 0xe3,
	0x89, 0x64, 0xc1, 0x7c, 0x11, 0x79, 0x29, 0xdd, 0xd2, 0xe5, 0xb8, 0xfe, 0xb9, 0x46, 0x26, 0xfa,
	0xbb, 0xc8, 0xcb, 0x43, 0x2e, 0x33, 0xa6, 0x60, 0xb7, 0xcd, 0xcf, 0xe1, 0xb4, 0x77, 0x00, 0x9d,
	0x4b, 0xc7, 0xb6, 0x5b, 0xd0, 0x78, 0x35, 0x7e, 0x3e, 0x9e, 0xbc, 0x1e, 0x93, 0x8f, 0x6c, 0x0c,
	0xe8, 0xfb, 0xd9, 0xec, 0x25, 0xb1, 0xec, 0x26, 0xd4, 0xd4, 0x6a, 0x4a, 0x2a, 0x76, 0x1d, 0x2a,
	0xaf, 0xa7, 0xa4, 0x6a, 0x37, 0xa0, 0xfa, 0x7a, 0x3a, 0x25, 0xc8, 0x41, 0x18, 0x93, 0xa6, 0x83,
	0x70, 0x93, 0x80, 0x83, 0x30, 0x90, 0x96, 0x83, 0x70, 0x8b, 0xb4, 0x1d, 0x84, 0xdb, 0xa4, 0xe3,
	0x20, 0xdc, 0x21, 0x5b, 0xbd, 0x3f, 0x2b, 0xd0, 0x9c, 0xc4, 0x3c, 0xd1, 0x67, 0x50, 0xdd, 0xc9,
	0x98, 0x9f, 0x52, 0x4b, 0xa7, 0xac, 0xd7, 0x7a, 0x1c, 0xf2, 0x30, 0x64, 0xc9, 0x5a, 0xf7, 0x5d,
	0x8d, 0x43, 0x61, 0xda, 0x5d, 0x68, 0x2d, 0x78, 0xea, 0x25, 0x42, 0xe7, 0x65, 0x5a, 0x7a, 0x11,
	0xda, 0x2c, 0x15, 0xfa, 0xf0, 0xa5, 0xb2, 0x1f, 0x40, 0x3b, 0x2a, 0x33, 0x98, 0x8b, 0x05, 0xad,
	0x15, 0xe7, 0x38, 0xc3, 0x0e, 0x17, 0xef, 0xdd, 0x6a, 0x7a, 0x3e, 0x7c, 0xa0, 0xa9, 0xb3, 0xe1,
	0xb9, 0x0f, 0xb0, 0xe0, 0x71, 0xc2, 0x3d, 0x96, 0xf1, 0x05, 0x6d, 0x75, 0xad, 0x3e, 0x76, 0x2f,
	0x20, 0x57, 0x6a, 0xdf, 0x26, 0x9d, 0xde, 0x1f, 0x16, 0x20, 0x75, 0x11, 0xec, 0x6d, 0xa8, 0x65,
	0x22, 0x0b, 0xb8, 0xb9, 0x5d, 0x85, 0x71, 0xb5, 0x98, 0x95, 0xcd, 0x62, 0xf6, 0x81, 0x64, 0x3c,
	0x09, 0xd3, 0x79, 0xb4, 0x9c, 0xa7, 0x3c, 0x39, 0x15, 0x1e, 0x37, 0x35, 0xdf, 0xd2, 0xf8, 0x64,
	0x39, 0x2d, 0x50, 0xdb, 0x81, 0x86, 0x17, 0xc9, 0x8c, 0x79, 0x99, 0x29, 0xf8, 0x57, 0xd7, 0x2e,
	0xf8, 0xd3, 0x22, 0xce, 0x2d, 0x05, 0x54, 0x09, 0x4e, 0x79, 0x92, 0xaa, 0x33, 0xd5, 0x8b, 0xf6,
	0x1b, 0xd3, 0x41, 0xb8, 0x46, 0xea, 0xbd, 0x11, 0x34, 0x4c, 0x8c, 0x9a, 0x1e, 0xc9, 0xc2, 0x32,
	0x2f, 0xbd, 0xb6, 0x09, 0x54, 0xf3, 0x24, 0x30, 0xe9, 0xa8, 0xa5, 0x4a, 0x9f, 0x87, 0x4c, 0x04,
	0xe6, 0xec, 0x85, 0xd1, 0x7b, 0x0e, 0xb7, 0xde, 0xd9, 0xeb, 0xab, 0x75, 0xb1, 0x36, 0xeb, 0xb2,
	0xb1, 0x45, 0xef, 0xf7, 0x0a, 0xd4, 0xf5, 0xb5, 0x61, 0xf6, 0x0c, 0x5a, 0xbf, 0xa4, 0x91, 0x9c,
	0xeb, 0xbe, 0x31, 0x1d, 0xde, 0xda, 0x7b, 0x72, 0xed, 0x72, 0x38, 0xd3, 0xc9, 0xb8, 0x50, 0x72,
	0x41, 0xe9, 0x18, 0xd5, 0x87, 0xd0, 0x59, 0x08, 0x75, 0x82, 0x50, 0x48, 0x96, 0x45, 0x89, 0xd9,
	0xfc, 0x32, 0xa8, 0xbe, 0x6b, 0x09, 0x67, 0x8b, 0x79, 0x24, 0x83, 0xb5, 0xce, 0x16, 0xbb, 0x58,
	0x01, 0x13, 0x19, 0xac, 0x37, 0xaf, 0x46, 0xed, 0x7f, 0xb8, 0x1a, 0x03, 0x68, 0xf0, 0x15, 0x0b,
	0xe3, 0x80, 0xeb, 0xe6, 0xb5, 0xf6, 0xb6, 0x07, 0xc5, 0x1b, 0x30, 0x28, 0xdf, 0x80, 0xc1, 0xbe,
	0x5c, 0xbb, 0xa5, 0x93, 0x83, 0x30, 0x22, 0xb5, 0xde, 0xdf, 0x75, 0x80, 0xf3, 0xc4, 0xcf, 0xe7,
	0xb5, 0xf6, 0x1f, 0xf3, 0x5a, 0xdf, 0xec, 0x0b, 0x85, 0xc6, 0x82, 0x2f, 0x59, 0x1e, 0x64, 0xb4,
	0x51, 0x4c, 0x8e, 0x31, 0xed, 0xcf, 0xa0, 0x15, 0xe6, 0x41, 0x26, 0xe2, 0x80, 0xcf, 0xa3, 0x25,
	0x85, 0xae, 0xd5, 0xb7, 0x5c, 0x28, 0xa1, 0xc9, 0x52, 0x85, 0x86, 0x6c, 0x25, 0xc2, 0x3c, 0xd4,
	0x57, 0xcb, 0x72, 0x4b, 0xd3, 0x7e, 0x0c, 0x37, 0xf8, 0xca, 0x0b, 0xf2, 0x54, 0x9c, 0xf2, 0x79,
	0xe9, 0xd3, 0xd6, 0xb5, 0x25, 0x67, 0xc4, 0x0b, 0xe3, 0xac, 0x64, 0x84, 0xd4, 0x2e, 0x1d, 0x23,
	0x53, 0x98, 0x57, 0x64, 0x8c, 0xcf, 0xd6, 0x55, 0x19, 0xe3, 0x7c, 0x0f, 0x20, 0x64, 0xab, 0x79,
	0xc0, 0xa5, 0x9f, 0x9d, 0xd0, 0x8f, 0xbb, 0x56, 0x1f, 0xb9, 0xcd, 0x90, 0xad, 0x8e, 0x34, 0xa0,
	0x69, 0x21, 0x4b, 0x9a, 0x18, 0x5a, 0x48, 0x43, 0x53, 0x68, 0xc4, 0x2c, 0x53, 0x4d, 0xa1, 0x37,
	0x8a, 0x32, 0x18, 0x53, 0xcd, 0x87, 0xd2, 0x15, 0x19, 0x0f, 0x53, 0xba, 0xad, 0xe3, 0x70, 0xc8,
	0x56, 0x87, 0xca, 0xd6, 0xa4, 0x90, 0x86, 0xbc, 0x65, 0x48, 0x21, 0x0b, 0xf2, 0x01, 0xb4, 0x73,
	0x29, 0xde, 0xe4, 0xdc, 0xf0, 0x9f, 0xe8, 0x93, 0xb7, 0x0a, 0xac, 0x70, 0xf9, 0x02, 0xb6, 0x94,
	0x78, 0x9c, 0xa8, 0xef, 0x60, 0x26, 0x78, 0x4a, 0xa9, 0x16, 0xe9, 0x84, 0x6c, 0xf5, 0xf2, 0x0c,
	0xd4, 0x6e, 0x42, 0x5e, 0x74, 0xbb, 0x6d, 0xdc, 0x84, 0xbc, 0xe0, 0xb6, 0x03, 0x38, 0xe1, 0x6f,
	0x72, 0x91, 0xf0, 0x05, 0xdd, 0x29, 0x3e, 0x92, 0xa5, 0xad, 0xe6, 0x83, 0x25, 0x09, 0x5b, 0xd3,
	0x9e, 0x26, 0x0a, 0xc3, 0xfe, 0x19, 0x50, 0xb6, 0x8e, 0x39, 0xfd, 0x5c, 0x3f, 0xda, 0x87, 0xef,
	0x71, 0xe3, 0x2e, 0x2c, 0xa7, 0x42, 0x8d, 0xe7, 0x6c, 0x1d, 0xf3, 0xd4, 0xd5, 0xb2, 0xbd, 0xb7,
	0x70, 0xeb, 0x9d, 0xf4, 0xe5, 0x77, 0xb2, 0x09, 0xb5, 0x7d, 0xd7, 0xdd, 0xff, 0x91, 0x58, 0x0a,
	0x3f, 0x98, 0x4c, 0x8e, 0x46, 0xfb, 0x63, 0x52, 0x51, 0xc6, 0xe1, 0x78, 0x36, 0x7a, 0x36, 0x72,
	0x49, 0x55, 0x3d, 0xa6, 0xe3, 0x57, 0x47, 0x47, 0x04, 0xd9, 0x00, 0xf5, 0xf1, 0xab, 0x17, 0x07,
	0x23, 0x97, 0xd4, 0xd4, 0x7a, 0x72, 0xe0, 0x8c, 0x9e, 0xce, 0x48, 0x5d, 0xad, 0xa7, 0x33, 0xf7,
	0x70, 0xfc, 0x8c, 0x34, 0x1c, 0x84, 0x2d, 0x52, 0x71, 0x10, 0xae, 0x90, 0xaa, 0x83, 0x70, 0x55,
	0x3f, 0xb3, 0x88, 0xd4, 0xae, 0x7c, 0xf0, 0x6d, 0x72, 0xd3, 0x41, 0xf8, 0x26, 0xd9, 0x76, 0x10,
	0xfe, 0x94, 0x50, 0x07, 0xe1, 0x3b, 0xe4, 0xae, 0x83, 0xf0, 0x5d, 0x72, 0xcf, 0x41, 0xf8, 0x1e,
	0xb9, 0xef, 0x20, 0x7c, 0x9f, 0xf4, 0x1c, 0x84, 0x1f, 0x92, 0x47, 0x0e, 0xc2, 0x8f, 0xc8, 0x63,
	0x07, 0xe1, 0xc7, 0x64, 0xd0, 0xfb, 0xcd, 0x82, 0xea, 0x8c, 0xf9, 0xd7, 0x78, 0x0f, 0x36, 0xbe,
	0x20, 0xd5, 0x0f, 0xff, 0x05, 0x29, 0xd2, 0x3d, 0xf8, 0xe6, 0xa7, 0xaf, 0x7d, 0x91, 0x9d, 0xe4,
	0xc7, 0x03, 0x2f, 0x0a, 0x87, 0xbf, 0x32, 0xb1, 0xbb, 0x10, 0xd2, 0x1f, 0xaa, 0x8d, 0x76, 0xcd,
	0x46, 0xc3, 0x7f, 0xff, 0xc7, 0x7a, 0x5c, 0xd7, 0xdc, 0x93, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x0b, 0x95, 0x40, 0x4c, 0xd6, 0x0a, 0x00, 0x00,
}
