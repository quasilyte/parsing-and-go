package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()

	lexer := NewLexer()
	lexer.s.Init(strings.NewReader(flag.Args()[0]))
	parser := PhpdocParserImpl{}
	status := parser.Parse(lexer)
	fmt.Println(status)
	ast := lexer.result
	fmt.Printf("%#v\n", ast)
	fmt.Printf("%s\n", ast)
}
