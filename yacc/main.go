package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "int|?float"
	lexer := NewLexer()
	lexer.s.Init(strings.NewReader(input))
	parser := PhpdocParserImpl{}
	status := parser.Parse(lexer)
	fmt.Println(status)
	fmt.Println(lexer.result)
}
