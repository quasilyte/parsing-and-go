package main

import (
	"testing"

	"github.com/alecthomas/participle/v2"
	"github.com/quasilyte/parsing-in-go/phpdoc"
	"github.com/quasilyte/parsing-in-go/phpdoctest"
)

type parserWrapper struct {
	parser *participle.Parser
}

func (p *parserWrapper) Parse(s string) (phpdoc.Type, error) {
	var result UnionExpr
	if err := p.parser.ParseString("", s, &result); err != nil {
		return nil, err
	}
	var conv converter
	return conv.Convert(&result), nil
}

func TestMain(t *testing.T) {
	parser := &parserWrapper{
		parser: participle.MustBuild(&UnionExpr{}),
	}
	phpdoctest.Run(t, parser)
}
