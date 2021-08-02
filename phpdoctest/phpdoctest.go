package phpdoctest

import (
	"testing"

	"github.com/quasilyte/parsing-in-go/phpdoc"
)

type parser interface {
	Parse(s string) (phpdoc.Type, error)
}

func Run(t *testing.T, p parser) {
	tests := []struct {
		input  string
		expect string
	}{
		{`int`, `int`},
		{`float`, `float`},
		{`string`, `string`},
		{`bool`, `bool`},

		{`A`, `A`},
		{`A\B`, `A\B`},
		{`A\B\C`, `A\B\C`},
		{`A \ B \ C`, `A\B\C`},

		{`(int)`, `int`},
		{`((int))`, `int`},

		{`?int`, `?(int)`},
		{`??int`, `?(?(int))`},
		{`?x[]`, `(?(x))[]`},
		{`?x[][]`, `((?(x))[])[]`},
		{`?(int[])[]`, `(?((int)[]))[]`},

		{`int?`, `(int)?`},
		{`(int[])[]?`, `(((int)[])[])?`},

		{`int[][]`, `((int)[])[]`},
		{`(int)[][]`, `((int)[])[]`},
		{`(int[])[]`, `((int)[])[]`},

		{`int|null[]`, `(int)|((null)[])`},

		{`int|float[]`, `(int)|((float)[])`},
		{`(int|float)[]`, `((int)|(float))[]`},

		{`int&float[]`, `(int)&((float)[])`},
		{`(int&float)[]`, `((int)&(float))[]`},

		{`x|y|z`, `(x)|((y)|(z))`},
		{`x&y&z`, `(x)&((y)&(z))`},
		{`x&y|z`, `((x)&(y))|(z)`},
		{`x|y&z`, `(x)|((y)&(z))`},
		{`x&(y|z)`, `(x)&((y)|(z))`},
		{`x&(y|z)[]`, `(x)&(((y)|(z))[])`},
		{`x&(y|z)[][]`, `(x)&((((y)|(z))[])[])`},

		{`?int|Reader&Writer|false`, `(?(int))|(((Reader)&(Writer))|(false))`},
	}

	for _, test := range tests {
		typ, err := p.Parse(` ` + test.input + ` `)
		if err != nil {
			t.Errorf("parse `%s` error: %v", test.input, err)
			continue
		}
		if typ.String() != test.expect {
			t.Errorf("parse `%s`:\nhave `%s`\nwant `%s`",
				test.input, typ.String(), test.expect)
		}
	}
}