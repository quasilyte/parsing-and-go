yacc_parser:
	@cd yacc && goyacc -p Phpdoc phpdoc.y && rm y.output

yacc_ragel_parser:
	@cd yacc_ragel && goyacc -p Phpdoc phpdoc.y && rm y.output
	@cd yacc_ragel && ragel -Z -G2 phpdoc.rl -o lexer.go

all: yacc_parser yacc_ragel_parser

test:
	@go test -v -race -count 2 ./participle
	@go test -v -race -count 2 ./yacc
	@go test -v -race -count 2 ./yacc_ragel
	@go test -v -race -count 2 ./handwritten

bench:
	@go test -bench=. -benchmem ./participle
	@go test -bench=. -benchmem ./yacc
	@go test -bench=. -benchmem ./yacc_ragel
	@go test -bench=. -benchmem ./handwritten

.PHONY: yacc_parser yacc_ragel_parser test bench all
