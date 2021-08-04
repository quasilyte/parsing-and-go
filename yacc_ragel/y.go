// Code generated by goyacc -p Phpdoc phpdoc.y. DO NOT EDIT.

//line phpdoc.y:2
package main

import __yyfmt__ "fmt"

//line phpdoc.y:2

import (
	"github.com/quasilyte/parsing-in-go/phpdoc"
	"strings"
)

//line phpdoc.y:11
type PhpdocSymType struct {
	yys  int
	tok  rune
	expr phpdoc.Type
	text string
}

const T_NULL = 57346
const T_FALSE = 57347
const T_INT = 57348
const T_FLOAT = 57349
const T_STRING = 57350
const T_BOOL = 57351
const T_NAME = 57352
const OPTIONAL = 57353

var PhpdocToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"T_NULL",
	"T_FALSE",
	"T_INT",
	"T_FLOAT",
	"T_STRING",
	"T_BOOL",
	"T_NAME",
	"'|'",
	"'&'",
	"OPTIONAL",
	"'['",
	"'?'",
	"'('",
	"')'",
	"']'",
}
var PhpdocStatenames = [...]string{}

const PhpdocEofCode = 1
const PhpdocErrCode = 2
const PhpdocInitialStackSize = 16

//line phpdoc.y:62

//line yacctab:1
var PhpdocExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const PhpdocPrivate = 57344

const PhpdocLast = 37

var PhpdocAct = [...]int{

	7, 8, 9, 10, 11, 12, 4, 19, 1, 3,
	0, 6, 5, 16, 15, 2, 13, 14, 0, 22,
	0, 17, 18, 16, 15, 0, 13, 14, 0, 0,
	0, 20, 21, 15, 0, 13, 14,
}
var PhpdocPact = [...]int{

	-4, -1000, 12, -1000, -1000, -4, -4, -1000, -1000, -1000,
	-1000, -1000, -1000, -11, -1000, -4, -4, 2, -1000, -1000,
	21, 12, -1000,
}
var PhpdocPgo = [...]int{

	0, 15, 9, 8,
}
var PhpdocR1 = [...]int{

	0, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 2, 2,
}
var PhpdocR2 = [...]int{

	0, 1, 1, 1, 3, 2, 3, 2, 3, 3,
	1, 1, 1, 1, 1, 1,
}
var PhpdocChk = [...]int{

	-1000, -3, -1, -2, 10, 16, 15, 4, 5, 6,
	7, 8, 9, 14, 15, 12, 11, -1, -1, 18,
	-1, -1, 17,
}
var PhpdocDef = [...]int{

	0, -2, 1, 2, 3, 0, 0, 10, 11, 12,
	13, 14, 15, 0, 7, 0, 0, 0, 5, 6,
	8, 9, 4,
}
var PhpdocTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 12, 3,
	16, 17, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 15, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 14, 3, 18, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 11,
}
var PhpdocTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 13,
}
var PhpdocTok3 = [...]int{
	0,
}

var PhpdocErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	PhpdocDebug        = 0
	PhpdocErrorVerbose = false
)

type PhpdocLexer interface {
	Lex(lval *PhpdocSymType) int
	Error(s string)
}

type PhpdocParser interface {
	Parse(PhpdocLexer) int
	Lookahead() int
}

type PhpdocParserImpl struct {
	lval  PhpdocSymType
	stack [PhpdocInitialStackSize]PhpdocSymType
	char  int
}

func (p *PhpdocParserImpl) Lookahead() int {
	return p.char
}

func PhpdocNewParser() PhpdocParser {
	return &PhpdocParserImpl{}
}

const PhpdocFlag = -1000

func PhpdocTokname(c int) string {
	if c >= 1 && c-1 < len(PhpdocToknames) {
		if PhpdocToknames[c-1] != "" {
			return PhpdocToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func PhpdocStatname(s int) string {
	if s >= 0 && s < len(PhpdocStatenames) {
		if PhpdocStatenames[s] != "" {
			return PhpdocStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func PhpdocErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !PhpdocErrorVerbose {
		return "syntax error"
	}

	for _, e := range PhpdocErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + PhpdocTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := PhpdocPact[state]
	for tok := TOKSTART; tok-1 < len(PhpdocToknames); tok++ {
		if n := base + tok; n >= 0 && n < PhpdocLast && PhpdocChk[PhpdocAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if PhpdocDef[state] == -2 {
		i := 0
		for PhpdocExca[i] != -1 || PhpdocExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; PhpdocExca[i] >= 0; i += 2 {
			tok := PhpdocExca[i]
			if tok < TOKSTART || PhpdocExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if PhpdocExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += PhpdocTokname(tok)
	}
	return res
}

func Phpdoclex1(lex PhpdocLexer, lval *PhpdocSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = PhpdocTok1[0]
		goto out
	}
	if char < len(PhpdocTok1) {
		token = PhpdocTok1[char]
		goto out
	}
	if char >= PhpdocPrivate {
		if char < PhpdocPrivate+len(PhpdocTok2) {
			token = PhpdocTok2[char-PhpdocPrivate]
			goto out
		}
	}
	for i := 0; i < len(PhpdocTok3); i += 2 {
		token = PhpdocTok3[i+0]
		if token == char {
			token = PhpdocTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = PhpdocTok2[1] /* unknown char */
	}
	if PhpdocDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", PhpdocTokname(token), uint(char))
	}
	return char, token
}

func PhpdocParse(Phpdoclex PhpdocLexer) int {
	return PhpdocNewParser().Parse(Phpdoclex)
}

func (Phpdocrcvr *PhpdocParserImpl) Parse(Phpdoclex PhpdocLexer) int {
	var Phpdocn int
	var PhpdocVAL PhpdocSymType
	var PhpdocDollar []PhpdocSymType
	_ = PhpdocDollar // silence set and not used
	PhpdocS := Phpdocrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Phpdocstate := 0
	Phpdocrcvr.char = -1
	Phpdoctoken := -1 // Phpdocrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Phpdocstate = -1
		Phpdocrcvr.char = -1
		Phpdoctoken = -1
	}()
	Phpdocp := -1
	goto Phpdocstack

ret0:
	return 0

ret1:
	return 1

Phpdocstack:
	/* put a state and value onto the stack */
	if PhpdocDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", PhpdocTokname(Phpdoctoken), PhpdocStatname(Phpdocstate))
	}

	Phpdocp++
	if Phpdocp >= len(PhpdocS) {
		nyys := make([]PhpdocSymType, len(PhpdocS)*2)
		copy(nyys, PhpdocS)
		PhpdocS = nyys
	}
	PhpdocS[Phpdocp] = PhpdocVAL
	PhpdocS[Phpdocp].yys = Phpdocstate

Phpdocnewstate:
	Phpdocn = PhpdocPact[Phpdocstate]
	if Phpdocn <= PhpdocFlag {
		goto Phpdocdefault /* simple state */
	}
	if Phpdocrcvr.char < 0 {
		Phpdocrcvr.char, Phpdoctoken = Phpdoclex1(Phpdoclex, &Phpdocrcvr.lval)
	}
	Phpdocn += Phpdoctoken
	if Phpdocn < 0 || Phpdocn >= PhpdocLast {
		goto Phpdocdefault
	}
	Phpdocn = PhpdocAct[Phpdocn]
	if PhpdocChk[Phpdocn] == Phpdoctoken { /* valid shift */
		Phpdocrcvr.char = -1
		Phpdoctoken = -1
		PhpdocVAL = Phpdocrcvr.lval
		Phpdocstate = Phpdocn
		if Errflag > 0 {
			Errflag--
		}
		goto Phpdocstack
	}

Phpdocdefault:
	/* default state action */
	Phpdocn = PhpdocDef[Phpdocstate]
	if Phpdocn == -2 {
		if Phpdocrcvr.char < 0 {
			Phpdocrcvr.char, Phpdoctoken = Phpdoclex1(Phpdoclex, &Phpdocrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if PhpdocExca[xi+0] == -1 && PhpdocExca[xi+1] == Phpdocstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Phpdocn = PhpdocExca[xi+0]
			if Phpdocn < 0 || Phpdocn == Phpdoctoken {
				break
			}
		}
		Phpdocn = PhpdocExca[xi+1]
		if Phpdocn < 0 {
			goto ret0
		}
	}
	if Phpdocn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Phpdoclex.Error(PhpdocErrorMessage(Phpdocstate, Phpdoctoken))
			Nerrs++
			if PhpdocDebug >= 1 {
				__yyfmt__.Printf("%s", PhpdocStatname(Phpdocstate))
				__yyfmt__.Printf(" saw %s\n", PhpdocTokname(Phpdoctoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Phpdocp >= 0 {
				Phpdocn = PhpdocPact[PhpdocS[Phpdocp].yys] + PhpdocErrCode
				if Phpdocn >= 0 && Phpdocn < PhpdocLast {
					Phpdocstate = PhpdocAct[Phpdocn] /* simulate a shift of "error" */
					if PhpdocChk[Phpdocstate] == PhpdocErrCode {
						goto Phpdocstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if PhpdocDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", PhpdocS[Phpdocp].yys)
				}
				Phpdocp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if PhpdocDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", PhpdocTokname(Phpdoctoken))
			}
			if Phpdoctoken == PhpdocEofCode {
				goto ret1
			}
			Phpdocrcvr.char = -1
			Phpdoctoken = -1
			goto Phpdocnewstate /* try again in the same state */
		}
	}

	/* reduction by production Phpdocn */
	if PhpdocDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Phpdocn, PhpdocStatname(Phpdocstate))
	}

	Phpdocnt := Phpdocn
	Phpdocpt := Phpdocp
	_ = Phpdocpt // guard against "declared and not used"

	Phpdocp -= PhpdocR2[Phpdocn]
	// Phpdocp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Phpdocp+1 >= len(PhpdocS) {
		nyys := make([]PhpdocSymType, len(PhpdocS)*2)
		copy(nyys, PhpdocS)
		PhpdocS = nyys
	}
	PhpdocVAL = PhpdocS[Phpdocp+1]

	/* consult goto table to find next state */
	Phpdocn = PhpdocR1[Phpdocn]
	Phpdocg := PhpdocPgo[Phpdocn]
	Phpdocj := Phpdocg + PhpdocS[Phpdocp].yys + 1

	if Phpdocj >= PhpdocLast {
		Phpdocstate = PhpdocAct[Phpdocg]
	} else {
		Phpdocstate = PhpdocAct[Phpdocj]
		if PhpdocChk[Phpdocstate] != -Phpdocn {
			Phpdocstate = PhpdocAct[Phpdocg]
		}
	}
	// dummy call; replaced with literal code
	switch Phpdocnt {

	case 1:
		PhpdocDollar = PhpdocS[Phpdocpt-1 : Phpdocpt+1]
//line phpdoc.y:39
		{
			Phpdoclex.(*PhpdocLex).result = PhpdocDollar[1].expr
		}
	case 2:
		PhpdocDollar = PhpdocS[Phpdocpt-1 : Phpdocpt+1]
//line phpdoc.y:43
		{
			PhpdocVAL.expr = PhpdocDollar[1].expr
		}
	case 3:
		PhpdocDollar = PhpdocS[Phpdocpt-1 : Phpdocpt+1]
//line phpdoc.y:44
		{
			PhpdocVAL.expr = &phpdoc.TypeName{Parts: strings.Split(PhpdocDollar[1].text, `\`)}
		}
	case 4:
		PhpdocDollar = PhpdocS[Phpdocpt-3 : Phpdocpt+1]
//line phpdoc.y:45
		{
			PhpdocVAL.expr = PhpdocDollar[2].expr
		}
	case 5:
		PhpdocDollar = PhpdocS[Phpdocpt-2 : Phpdocpt+1]
//line phpdoc.y:46
		{
			PhpdocVAL.expr = &phpdoc.NullableType{Elem: PhpdocDollar[2].expr}
		}
	case 6:
		PhpdocDollar = PhpdocS[Phpdocpt-3 : Phpdocpt+1]
//line phpdoc.y:47
		{
			PhpdocVAL.expr = &phpdoc.ArrayType{Elem: PhpdocDollar[1].expr}
		}
	case 7:
		PhpdocDollar = PhpdocS[Phpdocpt-2 : Phpdocpt+1]
//line phpdoc.y:48
		{
			PhpdocVAL.expr = &phpdoc.OptionalKeyType{Elem: PhpdocDollar[1].expr}
		}
	case 8:
		PhpdocDollar = PhpdocS[Phpdocpt-3 : Phpdocpt+1]
//line phpdoc.y:49
		{
			PhpdocVAL.expr = &phpdoc.IntersectionType{X: PhpdocDollar[1].expr, Y: PhpdocDollar[3].expr}
		}
	case 9:
		PhpdocDollar = PhpdocS[Phpdocpt-3 : Phpdocpt+1]
//line phpdoc.y:50
		{
			PhpdocVAL.expr = &phpdoc.UnionType{X: PhpdocDollar[1].expr, Y: PhpdocDollar[3].expr}
		}
	case 10:
		PhpdocDollar = PhpdocS[Phpdocpt-1 : Phpdocpt+1]
//line phpdoc.y:54
		{
			PhpdocVAL.expr = &phpdoc.PrimitiveTypeName{Name: "null"}
		}
	case 11:
		PhpdocDollar = PhpdocS[Phpdocpt-1 : Phpdocpt+1]
//line phpdoc.y:55
		{
			PhpdocVAL.expr = &phpdoc.PrimitiveTypeName{Name: "false"}
		}
	case 12:
		PhpdocDollar = PhpdocS[Phpdocpt-1 : Phpdocpt+1]
//line phpdoc.y:56
		{
			PhpdocVAL.expr = &phpdoc.PrimitiveTypeName{Name: "int"}
		}
	case 13:
		PhpdocDollar = PhpdocS[Phpdocpt-1 : Phpdocpt+1]
//line phpdoc.y:57
		{
			PhpdocVAL.expr = &phpdoc.PrimitiveTypeName{Name: "float"}
		}
	case 14:
		PhpdocDollar = PhpdocS[Phpdocpt-1 : Phpdocpt+1]
//line phpdoc.y:58
		{
			PhpdocVAL.expr = &phpdoc.PrimitiveTypeName{Name: "string"}
		}
	case 15:
		PhpdocDollar = PhpdocS[Phpdocpt-1 : Phpdocpt+1]
//line phpdoc.y:59
		{
			PhpdocVAL.expr = &phpdoc.PrimitiveTypeName{Name: "bool"}
		}
	}
	goto Phpdocstack /* stack new state and value */
}