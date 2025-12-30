package evaluators

import (
	"github.com/BinaryGhost/alpaci/lexing"
)

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
	__base__ ExpressionKind = iota
	Infix
	Postfix
	Prefix
)

func GetPrefixBindingPower(op lexing.Token) (int, int) {
	switch op.Type {
	case lexing.Plus_a, lexing.Minus_a:
		return EMPTY, 5
	default:
		panic("Bad prefix-operator for " + lexing.TokenToString(op))
	}
}

func GetInfixBindingPower(op lexing.Token) (int, int) {
	switch op.Type {
	case lexing.Plus_a, lexing.Minus_a:
		return 1, 2
	case lexing.Div_a, lexing.Mult_a:
		return 3, 4
	default:
		panic("Bad infix-operator for " + lexing.TokenToString(op))
	}
}

func GetPostfixBindingPower(op lexing.TokType) (int, int) {
	switch op {
	case lexing.Decr_a, lexing.Inc_a:
		return 7, EMPTY
	default:
		return EMPTY, EMPTY
	}
}

func IsPostfixOp(op lexing.TokType) bool {
	return op == lexing.Inc_a || op == lexing.Decr_a
}

// func IsInfixOp(op lexing.TokType) bool {
// 	switch op {
// 	case lexing.Div_a, lexing.Mult_a, lexing.Mod_a, lexing.Plus_a, lexing.Minus_a:
// 		return true
// 	default:
// 		return false
// 	}
// }
