package main

import (
	"testing"

	"github.com/quasilyte/parsing-in-go/phpdoc"
	"github.com/quasilyte/parsing-in-go/phpdoctest"
)

type parserWrapper struct {
	parser *TypeParser
}

func (p *parserWrapper) Parse(s string) (phpdoc.Type, error) {
	typ := p.parser.Parse(s)
	var conv converter
	return conv.Convert(typ.Expr), nil
}

func TestMain(t *testing.T) {
	parser := &parserWrapper{
		parser: NewTypeParser(),
	}
	phpdoctest.Run(t, parser)
}

func BenchmarkParser(b *testing.B) {
	parser := &parserWrapper{
		parser: NewTypeParser(),
	}
	phpdoctest.RunBenchmark(b, parser)
}
