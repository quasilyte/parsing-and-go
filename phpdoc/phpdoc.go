package phpdoc

import (
	"fmt"
	"strings"
)

type Type interface {
	String() string
}

type PrimitiveTypeName struct {
	Name string
}

type TypeName struct {
	Parts []string
}

type NullableType struct {
	Elem Type
}

type OptionalKeyType struct {
	Elem Type
}

type ArrayType struct {
	Elem Type
}

type UnionType struct {
	X Type
	Y Type
}

type IntersectionType struct {
	X Type
	Y Type
}

func (typ *PrimitiveTypeName) String() string { return typ.Name }
func (typ *TypeName) String() string          { return strings.Join(typ.Parts, `\`) }
func (typ *NullableType) String() string      { return "?(" + typ.Elem.String() + ")" }
func (typ *OptionalKeyType) String() string   { return "(" + typ.Elem.String() + ")?" }
func (typ *ArrayType) String() string         { return "(" + typ.Elem.String() + ")[]" }
func (typ *UnionType) String() string         { return fmt.Sprintf("(%s)|(%s)", typ.X, typ.Y) }
func (typ *IntersectionType) String() string  { return fmt.Sprintf("(%s)&(%s)", typ.X, typ.Y) }
