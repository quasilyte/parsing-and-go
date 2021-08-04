%{
package main

import (
    "github.com/quasilyte/parsing-in-go/phpdoc"
    "strings"
)

%}

%union{
    tok   rune
    expr  phpdoc.Type
    text  string
}

%token <tok> T_NULL
%token <tok> T_FALSE
%token <tok> T_INT
%token <tok> T_FLOAT
%token <tok> T_STRING
%token <tok> T_BOOL
%token <text> T_NAME

%right '|'
%right '&'
%left OPTIONAL
%right '['
%left '?'

%type <expr> type_expr
%type <expr> primitive_type

%start start

%%

start
    : type_expr { Phpdoclex.(*PhpdocLex).result = $1 }
    ;

type_expr
    : primitive_type { $$ = $1 }
    | T_NAME { $$ = &phpdoc.TypeName{Parts: strings.Split($1, `\`)} }
    | '(' type_expr ')' { $$ = $2 }
    | '?' type_expr { $$ = &phpdoc.NullableType{Elem: $2} }
    | type_expr '[' ']' { $$ = &phpdoc.ArrayType{Elem: $1} }
    | type_expr '?' %prec OPTIONAL { $$ = &phpdoc.OptionalKeyType{Elem: $1} }
    | type_expr '&' type_expr { $$ = &phpdoc.IntersectionType{X: $1, Y: $3} }
    | type_expr '|' type_expr { $$ = &phpdoc.UnionType{X: $1, Y: $3} }
    ;

primitive_type
    : T_NULL { $$ = &phpdoc.PrimitiveTypeName{Name: "null"} }
    | T_FALSE { $$ = &phpdoc.PrimitiveTypeName{Name: "false"} }
    | T_INT { $$ = &phpdoc.PrimitiveTypeName{Name: "int"} }
    | T_FLOAT { $$ = &phpdoc.PrimitiveTypeName{Name: "float"} }
    | T_STRING { $$ = &phpdoc.PrimitiveTypeName{Name: "string"} }
    | T_BOOL { $$ = &phpdoc.PrimitiveTypeName{Name: "bool"} }
    ;

%%
