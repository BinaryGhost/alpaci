package evaluators

import (
	"strconv"
	"strings"

	"github.com/BinaryGhost/gospel/evaluators/lexing"
)

type Type int

const (
	Bool Type = iota + 1
	String
	Float64
	// Possibly more...
)

type Atom struct {
	Val      any
	RealType Type
	Type     lexing.TokType
	Column   int
}

// Currently only returns 1 as a real value, but would in theory have to access
// an environment or similiar
func MakeAtom(tl *lexing.TokenList) Atom {
	atom, err := tl.Current()
	if err != nil {
		panic(err)
	}
	tl.Next()

	switch atom.Type {
	case lexing.Ident:
		// TODO
		// Is 1 for now, but val would be the variable name
		// variable access not implemented yet
		return Atom{
			Val:      1.0,
			Type:     atom.Type,
			Column:   atom.Column,
			RealType: Float64,
		}

	case lexing.False_k, lexing.True_k:
		return Atom{
			Val:      atom.Type == lexing.True_k,
			Type:     atom.Type,
			Column:   atom.Column,
			RealType: Bool,
		}

	case lexing.String:
		return Atom{
			Val:      atom.Value,
			Type:     atom.Type,
			Column:   atom.Column,
			RealType: String,
		}

	default:
		panic("")
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
		actual_number, _ = strconv.ParseFloat(before_point, 64)
		return Atom{
			Val:      actual_number,
			Type:     first_tok.Type,
			Column:   first_tok.Column,
			RealType: Float64,
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
		Val:      actual_number,
		Type:     third_tok.Type,
		Column:   third_tok.Column,
		RealType: Float64,
	}

}

func AccessVar() {}
