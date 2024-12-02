package schema

type OperationType string

const (
	QueryOperation OperationType = "query"
	MutationOperation OperationType = "mutation"
	SubscriptionOperation OperationType = "subscription"
)

type TypeDefinition struct {
	Name []byte
	Fields []*FieldDefinition
	tokens Tokens
}

type FieldType struct {
	Name []byte
	Nullable bool
	IsList bool
	ListType *FieldType
}

type FieldDefinition struct {
	Name []byte
	Arguments []*ArgumentDefinition
	Type *FieldType
}

type OperationDefinition struct {
	OperationType OperationType
	Name []byte
	Fields []*FieldDefinition
	Type []byte
}

type ArgumentDefinition struct {
	Name []byte
	Type []byte
	DefaultValue any
}

type EnumDefinition struct {
	Name []byte
	Values [][]byte
}

type UnionDefinition struct {
	Name []byte
	Types [][]byte
}

type InterfaceDefinition struct {
	Name []byte
	Fields []*FieldDefinition
}

type DirectiveDefinition struct {
	Name []byte
	Arguments []*ArgumentDefinition
	Locations [][]byte
}

type InputDefinition struct {
	Name []byte
	Fields []*FieldDefinition
}

type Schema struct {
	tokens Tokens
	Operations []*OperationDefinition
	Types []*TypeDefinition
	Enums []*EnumDefinition
	Unions []*UnionDefinition
	Interfaces []*InterfaceDefinition
	Directives []*DirectiveDefinition
	Inputs []*InputDefinition
}