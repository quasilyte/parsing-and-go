package phpdoctest

import (
	"testing"

	"github.com/quasilyte/parsing-in-go/phpdoc"
)

type parser interface {
	Parse(s string) (phpdoc.Type, error)
}

func RunBenchmark(b *testing.B, p parser) {
	tests := []struct {
		label string
		input string
	}{
		{`simple`, `int`},
		{`normal`, `Foo|Bar|null`},
		{`complex`, `(?a|c|false)&d`},
		{`array`, `int[][]`},
		{`classname`, `A\B\C\D`},
		{`whitespace`, `   (    int   )  `},
	}

	for _, test := range tests {
		input := test.input
		b.Run(test.label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := p.Parse(input)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
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

		{`   (    int   )  `, `int`},

		{`A`, `A`},
		{`A\B`, `A\B`},
		{`A\B\C`, `A\B\C`},

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
		if typ == nil {
			t.Errorf("parse `%s` returned type is nil!", test.input)
			continue
		}
		if typ.String() != test.expect {
			t.Errorf("parse `%s`:\nhave `%s`\nwant `%s`",
				test.input, typ.String(), test.expect)
			continue
		}
		typ2, err := p.Parse(typ.String())
		if err != nil {
			t.Errorf("re-parse `%s` error: %v", typ.String(), err)
			continue
		}
		if typ.String() != typ2.String() {
			t.Errorf("re-parsed representation mismatch:\nA: `%s`\nB: `%s`",
				typ.String(), typ2.String())
			continue
		}

	}
}
