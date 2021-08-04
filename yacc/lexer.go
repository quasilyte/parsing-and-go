package main

import (
	"fmt"
	"strings"
	"text/scanner"
	"unicode"

	"github.com/quasilyte/parsing-in-go/phpdoc"
)

type PhpdocLex struct {
	s      scanner.Scanner
	result phpdoc.Type
}

func NewLexer() *PhpdocLex {
	lexer := &PhpdocLex{}
	lexer.s.IsIdentRune = func(ch rune, i int) bool {
		return unicode.IsLetter(ch) ||
			(unicode.IsDigit(ch) && i > 0) ||
			ch == '\\'
	}
	return lexer
}

func (l *PhpdocLex) Init(s string) {
	l.s.Init(strings.NewReader(s))
}

var nameToToken = map[string]rune{
	"int":    T_INT,
	"float":  T_FLOAT,
	"null":   T_NULL,
	"string": T_STRING,
	"false":  T_FALSE,
	"bool":   T_BOOL,
}

func (l *PhpdocLex) nextToken(sym *PhpdocSymType) rune {
	tok := l.s.Scan()
	if tok == scanner.Ident {
		text := l.s.TokenText()
		tok, ok := nameToToken[text]
		if ok {
			return tok
		}
		sym.text = text
		return T_NAME
	}
	return tok
}

func (l *PhpdocLex) Lex(sym *PhpdocSymType) int {
	tok := l.nextToken(sym)
	sym.tok = tok
	return int(tok)
}

func (l *PhpdocLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
