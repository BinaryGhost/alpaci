package evaluators

import "github.com/BinaryGhost/alpaci/lexing"

type Operator struct {
	Val    string
	Type   lexing.TokType
	Column int
}

const EMPTY = -99
const NO_OP = 0
const ASSIGN = 0.1

type ExpressionKind int

const (
	Infix ExpressionKind = iota
	Postix
	Prefix
)

func GetPrefixBindingPower(op lexing.TokType) (int, int) {
	switch op {
	case lexing.Plus_a, lexing.Minus_a:
		return EMPTY, 5
	default:
		panic("Bad operator")
	}
}

func GetInfixBindingPower(op lexing.TokType) (int, int) {
	switch op {
	case lexing.Plus_a, lexing.Minus_a:
		return 1, 2
	case lexing.Div_a, lexing.Mult_a:
		return 3, 4
	default:
		panic("Bad operator")
	}
}

func IsInfixOp(op lexing.TokType) bool {
	switch op {
	case lexing.Div_a, lexing.Mult_a, lexing.Mod_a, lexing.Plus_a, lexing.Minus_a:
		return true
	default:
		return false
	}
}
