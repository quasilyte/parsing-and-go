package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/alecthomas/participle/v2"
)

// Note: using something like `@@ "|" @@` will cause infinite recursion.
// Note: this "grammar" produces inefficient AST, so it's impractical for big
// inputs, but OK for phpdoc types

type PrimaryExpr struct {
	TypeName *TypeNameExpr `@@`
	Parens   *UnionExpr    `| "(" @@ ")"`
}

type TypeNameExpr struct {
	Primitive *string        `@("null"|"false"|"int"|"float"|"string"|"bool")`
	Class     *ClassNameExpr `| @@`
}

type ClassNameExpr struct {
	Part string         `@Ident`
	Next *ClassNameExpr `("\\" @@)?`
}

type PrefixExpr struct {
	Ops   []*PrefixOp  `@@*`
	Right *PrimaryExpr `@@`
}

type PrefixOp struct {
	Tok string `@("?")`
}

type PostfixExpr struct {
	Left *PrefixExpr  `@@`
	Ops  []*PostfixOp `@@*`
}

type PostfixOp struct {
	Tok string `@("[" "]" | "?")`
}

type IntersectionExpr struct {
	Left  *PostfixExpr      `@@`
	Right *IntersectionExpr `("&" @@)?`
}

type UnionExpr struct {
	Left  *IntersectionExpr `@@`
	Right *UnionExpr        `("|" @@)?`
}

func main() {
	flag.Parse()

	parser := participle.MustBuild(&UnionExpr{})
	var result UnionExpr
	if err := parser.ParseString("", flag.Args()[0], &result); err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(result); err != nil {
		panic(err)
	}

	var conv converter
	ast := conv.Convert(&result)
	fmt.Printf("%#v\n", ast)
	fmt.Printf("%s\n", ast)
}
