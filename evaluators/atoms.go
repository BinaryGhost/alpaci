package evaluators

import (
	"github.com/BinaryGhost/alpaci/lexing"
	"strconv"
	"strings"
)

type Atom struct {
	Val    any
	Type   lexing.TokType
	Column int
}

// Currently only returns 1 as a real value, but would in theory have to access
// an environment or similiar
func MakeIdentAtom(tl *lexing.TokenList) Atom {
	ident, err := tl.Current()
	if err != nil {
		panic(err)
	}

	// TODO
	// Is 1 for now, but val would be the variable name
	// variable access not implemented yet
	return Atom{
		Val:    1,
		Type:   ident.Type,
		Column: ident.Column,
	}
}

// If it can, it will return a float as a number-atom, else an integer-atom
//
// Panic, if the TokenList is not long enough to fit a float or integer
func MakeNumberAtom(tl *lexing.TokenList) Atom {
	var actual_number any
	first_tok, err := tl.Current()
	if err != nil {
		panic(err)
	}

	var before_point string
	if strings.Contains(first_tok.Value, "_") {
		before_point = strings.ReplaceAll(first_tok.Value, "_", "")
	} else {
		before_point = first_tok.Value
	}
	tl.Next()

	second_tok, err := tl.Current()
	if err != nil {
		panic(err)
	}

	if second_tok.Type != lexing.Point {
		actual_number, _ = strconv.ParseInt(before_point, 10, 64)
		return Atom{
			Val:    actual_number,
			Type:   second_tok.Type,
			Column: second_tok.Column,
		}
	}
	tl.Next()

	var after_point string
	third_tok, err := tl.Current()
	if err != nil {
		panic(err)
	}

	if third_tok.Type != lexing.Number {
		panic("Bad floating-point number: No number given after '.', only [" + third_tok.Value + "]")
	}

	if strings.Contains(third_tok.Value, "_") {
		after_point = strings.ReplaceAll(third_tok.Value, "_", "")
	} else {
		after_point = third_tok.Value
	}
	tl.Next()

	actual_number, _ = strconv.ParseFloat(after_point, 64)
	return Atom{
		Val:    actual_number,
		Type:   third_tok.Type,
		Column: third_tok.Column,
	}

}

func AccessVar() {}
