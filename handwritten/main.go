package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	parser := NewTypeParser()
	result := parser.Parse(flag.Args()[0])

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(result); err != nil {
		panic(err)
	}

	var conv converter
	ast := conv.Convert(result.Expr)
	fmt.Printf("%#v\n", ast)
	fmt.Printf("%s\n", ast)
}
