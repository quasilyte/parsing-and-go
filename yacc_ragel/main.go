package main

import (
	"fmt"
)

func main() {
	lexer := NewLexer()
	lexer.Init("int[]")
	parser := PhpdocParserImpl{}
	status := parser.Parse(lexer)
	fmt.Println(status)
	fmt.Println(lexer.result)
}
