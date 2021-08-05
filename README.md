# parsing-in-go

> This repository provides examples I referred to in "Parsing & Go" talk

Parsing phpdoc type expressions using several approaches:

* [participle](/participle) example uses [alecthomas/participle](https://github.com/alecthomas/participle) with default lexer
* [yacc](/yacc) example uses [goyacc](https://pkg.go.dev/golang.org/x/tools/cmd/goyacc) with [text/scanner](https://pkg.go.dev/text/scanner) for lexing
* [yacc_ragel](/yacc_ragel) example uses [goyacc](https://pkg.go.dev/golang.org/x/tools/cmd/goyacc) with a lexer generated with [Ragel](https://github.com/adrian-thurston/ragel)
* [handwritten](/handwritten) example shows a manually created parser from [NoVerify](https://github.com/VKCOM/noverify) static analyzer

These example parsers understand these types (and their combinations):

| Type | Example |
|---|---|
| primitive type | `int`, `float`, `null`, ... |
| type name | `Foo`, `Foo\Bar` |
| nullable type | `?T` |
| optional key type | `T?` |
| array type | `T[]` |
| union type | `X\|Y` |
| intersection type | `X&Y` |

Every parser package is a `main` that can parse a command-line argument:

```bash
go run ./participle 'int|Foo\Bar'

go run ./yacc 'int|Foo\Bar'

go run ./yacc_ragel 'int|Foo\Bar'

go run ./handwritten 'int|Foo\Bar'
```

Some useful `make` commands:

```bash
# build all parsers (requires goyacc and ragel installed)
make all

# run all tests
make test

# run all benchmarks
make bench
```
