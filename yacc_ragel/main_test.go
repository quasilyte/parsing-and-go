package main

import (
	"testing"

	"github.com/quasilyte/parsing-in-go/phpdoc"
	"github.com/quasilyte/parsing-in-go/phpdoctest"
)

type parserWrapper struct {
	parser PhpdocParserImpl
}

func (p *parserWrapper) Parse(s string) (phpdoc.Type, error) {
	lexer := NewLexer()
	lexer.Init(s)
	p.parser.Parse(lexer)
	return lexer.result, nil
}

func TestMain(t *testing.T) {
	parser := &parserWrapper{}
	phpdoctest.Run(t, parser)
}

func BenchmarkParser(b *testing.B) {
	parser := &parserWrapper{}
	phpdoctest.RunBenchmark(b, parser)
}
