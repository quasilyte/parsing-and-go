//+build ignore
// +build ignore

package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/alecthomas/participle/v2"
)

// TODO: read http://www.craftinginterpreters.com/parsing-expressions.html

// int
// Foo
// A\B
// A|B
// int[]
// ?int
// ?int[]
// (A|B)[]
// ?(A|B)[]
// A|B[]

// Note: using something like `@@ "|" @@` will cause infinite recursion.

type TypeExpr struct {
	Left  *TypeExprTerm `@@`
	Right *TypeExprOp   `@@*`
}

type TypeExprTerm struct {
	TypeName *TypeNameExpr `@@`
	Nullable *TypeExpr     `| "?" @@`
	Parens   *TypeExpr     `| "(" @@ ")"`
}

type TypeExprOp struct {
	PostfixOp *PostfixOp `@@`
	InfixOp   *InfixOp   `| @@`
}

type PostfixOp struct {
	Op   string     `@("[" "]" | "?")`
	Next *PostfixOp "@@?"
}

type IntersectionExpr struct {
}

type InfixOp struct {
	IntersectionExpr *TypeExpr `("&" @@)`
	UnionExpr        *TypeExpr `| ("|" @@)`

	// Op   string    `@("|" | "&")`
	// Expr *TypeExpr `@@`
}

type TypeNameExpr struct {
	Primitive *string   `@("int"|"float"|"string"|"bool")`
	Name      *NameExpr `| @@`
}

type NameExpr struct {
	Part string    `@Ident`
	Next *NameExpr `("\\" @@)?`
}

type ArrayTypeExpr struct {
	Elem *TypeExpr `@@ "[" "]"`
}

func main() {
	flag.Parse()

	parser := participle.MustBuild(&TypeExpr{})
	var result TypeExpr
	if err := parser.ParseString("", flag.Args()[0], &result); err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(result); err != nil {
		panic(err)
	}

	// var conv converter
	// ast := conv.Convert(&result)
	// fmt.Printf("%#v\n", ast)
	// fmt.Printf("%s\n", ast)
}
