
//line phpdoc.rl:1
package main


//line phpdoc.rl:4

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

    
//line phpdoc.rl:52


    
//line lexer.go:43
	{
	cs = lexer_start
	ts = 0
	te = 0
	act = 0
	}

//line phpdoc.rl:55
    
//line lexer.go:53
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	}
	goto st_out
tr0:
//line phpdoc.rl:50
te = p+1
{ tok = int(data[ts]); {p++; cs = 0; goto _out } }
	goto st0
tr1:
//line phpdoc.rl:42
te = p+1

	goto st0
tr8:
//line NONE:1
	switch act {
	case 2:
	{p = (te) - 1
 tok = T_INT; {p++; cs = 0; goto _out } }
	case 3:
	{p = (te) - 1
 tok = T_FLOAT; {p++; cs = 0; goto _out } }
	case 4:
	{p = (te) - 1
 tok = T_NULL; {p++; cs = 0; goto _out } }
	case 5:
	{p = (te) - 1
 tok = T_STRING; {p++; cs = 0; goto _out } }
	case 6:
	{p = (te) - 1
 tok = T_FALSE; {p++; cs = 0; goto _out } }
	case 7:
	{p = (te) - 1
 tok = T_BOOL; {p++; cs = 0; goto _out } }
	case 8:
	{p = (te) - 1
 tok = T_NAME; {p++; cs = 0; goto _out } }
	}
	
	goto st0
tr9:
//line phpdoc.rl:49
te = p
p--
{ tok = T_NAME; {p++; cs = 0; goto _out } }
	goto st0
	st0:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
//line NONE:1
ts = p

//line lexer.go:159
		switch data[p] {
		case 9:
			goto tr1
		case 32:
			goto tr1
		case 96:
			goto tr0
		case 98:
			goto st2
		case 102:
			goto st5
		case 105:
			goto st12
		case 110:
			goto st14
		case 115:
			goto st17
		}
		switch {
		case data[p] < 91:
			if data[p] <= 64 {
				goto tr0
			}
		case data[p] > 94:
			if 123 <= data[p] && data[p] <= 127 {
				goto tr0
			}
		default:
			goto tr0
		}
		goto tr2
tr2:
//line NONE:1
te = p+1

//line phpdoc.rl:49
act = 8;
	goto st1
tr12:
//line NONE:1
te = p+1

//line phpdoc.rl:48
act = 7;
	goto st1
tr17:
//line NONE:1
te = p+1

//line phpdoc.rl:47
act = 6;
	goto st1
tr20:
//line NONE:1
te = p+1

//line phpdoc.rl:44
act = 3;
	goto st1
tr22:
//line NONE:1
te = p+1

//line phpdoc.rl:43
act = 2;
	goto st1
tr25:
//line NONE:1
te = p+1

//line phpdoc.rl:45
act = 4;
	goto st1
tr30:
//line NONE:1
te = p+1

//line phpdoc.rl:46
act = 5;
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line lexer.go:245
		switch data[p] {
		case 91:
			goto tr8
		case 96:
			goto tr8
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr8
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr8
				}
			case data[p] >= 93:
				goto tr8
			}
		default:
			goto tr8
		}
		goto tr2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 111:
			goto st3
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 111:
			goto st4
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 108:
			goto tr12
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 97:
			goto st6
		case 108:
			goto st9
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 108:
			goto st7
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 115:
			goto st8
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 101:
			goto tr17
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 111:
			goto st10
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 97:
			goto st11
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 116:
			goto tr20
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 110:
			goto st13
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 116:
			goto tr22
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 117:
			goto st15
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 108:
			goto st16
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 108:
			goto tr25
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 116:
			goto st18
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 114:
			goto st19
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 105:
			goto st20
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 110:
			goto st21
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 91:
			goto tr9
		case 96:
			goto tr9
		case 103:
			goto tr30
		}
		switch {
		case data[p] < 58:
			if data[p] <= 47 {
				goto tr9
			}
		case data[p] > 64:
			switch {
			case data[p] > 94:
				if 123 <= data[p] && data[p] <= 127 {
					goto tr9
				}
			case data[p] >= 93:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr2
	st_out:
	_test_eof0: cs = 0; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 1:
			goto tr8
		case 2:
			goto tr9
		case 3:
			goto tr9
		case 4:
			goto tr9
		case 5:
			goto tr9
		case 6:
			goto tr9
		case 7:
			goto tr9
		case 8:
			goto tr9
		case 9:
			goto tr9
		case 10:
			goto tr9
		case 11:
			goto tr9
		case 12:
			goto tr9
		case 13:
			goto tr9
		case 14:
			goto tr9
		case 15:
			goto tr9
		case 16:
			goto tr9
		case 17:
			goto tr9
		case 18:
			goto tr9
		case 19:
			goto tr9
		case 20:
			goto tr9
		case 21:
			goto tr9
		}
	}

	_out: {}
	}

//line phpdoc.rl:56

    l.pos = p
    if tok == T_NAME {
        lval.text = data[ts:te]
    }

	return tok
}

func (l *PhpdocLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}


//line lexer.go:982
const lexer_start int = 0
const lexer_first_final int = 0
const lexer_error int = -1

const lexer_en_main int = 0


//line phpdoc.rl:70
