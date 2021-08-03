package main

import (
	"fmt"
	"strings"

	"github.com/quasilyte/parsing-in-go/phpdoc"
)

type converter struct {
}

func (conv *converter) convertUnion(dst *phpdoc.UnionType, args []TypeExpr) {
	dst.X = conv.Convert(args[0])
	args = args[1:]
	if len(args) >= 2 {
		nested := &phpdoc.UnionType{}
		conv.convertUnion(nested, args)
		dst.Y = nested
	} else if len(args) == 1 {
		dst.Y = conv.Convert(args[0])
	}
}

func (conv *converter) convertInter(dst *phpdoc.IntersectionType, args []TypeExpr) {
	dst.X = conv.Convert(args[0])
	args = args[1:]
	if len(args) >= 2 {
		nested := &phpdoc.IntersectionType{}
		conv.convertInter(nested, args)
		dst.Y = nested
	} else if len(args) == 1 {
		dst.Y = conv.Convert(args[0])
	}
}

func (conv *converter) Convert(expr TypeExpr) phpdoc.Type {
	switch expr.Kind {
	case ExprName:
		switch expr.Value {
		case `int`, `float`, `string`, `bool`, `false`, `null`:
			return &phpdoc.PrimitiveTypeName{Name: expr.Value}
		default:
			return &phpdoc.TypeName{Parts: strings.Split(expr.Value, `\`)}
		}
	case ExprParen:
		return conv.Convert(expr.Args[0])
	case ExprNullable:
		return &phpdoc.NullableType{Elem: conv.Convert(expr.Args[0])}
	case ExprArray:
		return &phpdoc.ArrayType{Elem: conv.Convert(expr.Args[0])}
	case ExprOptional:
		return &phpdoc.OptionalKeyType{Elem: conv.Convert(expr.Args[0])}
	case ExprUnion:
		union := &phpdoc.UnionType{}
		conv.convertUnion(union, expr.Args)
		return union
	case ExprInter:
		inter := &phpdoc.IntersectionType{}
		conv.convertInter(inter, expr.Args)
		return inter
	default:
		fmt.Println("unhandled " + expr.Value)
	}
	return nil
}
