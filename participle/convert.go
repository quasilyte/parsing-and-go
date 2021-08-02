package main

import (
	"github.com/quasilyte/parsing-in-go/phpdoc"
)

type converter struct{}

func (conv *converter) Convert(expr *UnionExpr) phpdoc.Type {
	return conv.convertUnion(expr)
}

func (conv *converter) convertPrefix(expr *PrefixExpr) phpdoc.Type {
	result := conv.convertPrimary(expr.Right)
	if len(expr.Ops) != 0 {
		for _, prefix := range expr.Ops {
			switch prefix.Tok {
			case "?":
				result = &phpdoc.NullableType{Elem: result}
			}
		}
	}
	// if expr.Postfix != nil {

	// }
	return result
}

func (conv *converter) convertPostfix(expr *PostfixExpr) phpdoc.Type {
	result := conv.convertPrefix(expr.Left)
	if len(expr.Ops) != 0 {
		for _, postfix := range expr.Ops {
			switch postfix.Tok {
			case "[]":
				result = &phpdoc.ArrayType{Elem: result}
			case "?":
				result = &phpdoc.OptionalKeyType{Elem: result}
			}
		}
	}
	return result
}

func (conv *converter) convertUnion(expr *UnionExpr) phpdoc.Type {
	left := conv.convertIntersection(expr.Left)
	if expr.Right != nil {
		right := conv.convertUnion(expr.Right)
		return &phpdoc.UnionType{X: left, Y: right}
	}
	return left
}

func (conv *converter) convertIntersection(expr *IntersectionExpr) phpdoc.Type {
	left := conv.convertPostfix(expr.Left)
	if expr.Right != nil {
		right := conv.convertIntersection(expr.Right)
		return &phpdoc.IntersectionType{X: left, Y: right}
	}
	return left
}

func (conv *converter) convertPrimary(expr *PrimaryExpr) phpdoc.Type {
	if expr.TypeName != nil {
		if expr.TypeName.Primitive != nil {
			return &phpdoc.PrimitiveTypeName{Name: *expr.TypeName.Primitive}
		}
		return conv.convertClassName(expr.TypeName.Class, nil)
	}
	if expr.Parens != nil {
		return conv.Convert(expr.Parens)
	}
	return nil
}

func (conv *converter) convertClassName(name *ClassNameExpr, parts []string) phpdoc.Type {
	parts = append(parts, name.Part)
	if name.Next == nil {
		return &phpdoc.TypeName{Parts: parts}
	}
	return conv.convertClassName(name.Next, parts)
}
