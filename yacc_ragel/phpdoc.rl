package main

%%machine lexer;

import (
    "github.com/quasilyte/parsing-in-go/phpdoc"
    "fmt"
)

type PhpdocLex struct {
    pos    int
    src    string
    result phpdoc.Type
}

func NewLexer() *PhpdocLex {
    return &PhpdocLex{}
}

func (l *PhpdocLex) Init(s string) {
    l.pos = 0
    l.src = s
}

func (l *PhpdocLex) Lex(lval *PhpdocSymType) int {
    tok := 0

    data := l.src
    cs, p, pe := 0, l.pos, len(data)
    eof := pe
    var ts, te int
    var act int

    %%{
        whitespace = [\t ];

        ident_first = [a-zA-Z_] | (0x0080..0x00FF);
        ident_rest  = ident_first | [0-9] | [\\];
        ident       = ident_first (ident_rest)*;

        main := |*
            whitespace => {};
            'int' => { tok = T_INT; fbreak; };
            'float' => { tok = T_FLOAT; fbreak; };
            'null' => { tok = T_NULL; fbreak; };
            'string' => { tok = T_STRING; fbreak; };
            'false' => { tok = T_FALSE; fbreak; };
            'bool' => { tok = T_BOOL; fbreak; };
            ident => { tok = T_NAME; fbreak; };
            any => { tok = int(data[ts]); fbreak; };
        *|;
    }%%

    %%write init;
    %%write exec;

    l.pos = p
    if tok == T_NAME {
        lval.text = data[ts:te]
    }

	return tok
}

func (l *PhpdocLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

%%write data;
