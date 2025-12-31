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

func GetPrefixBindingPower(op lexing.Token) (float32, float32) {
	switch op.Type {
	case lexing.Plus_a, lexing.Minus_a, lexing.Bang_l:
		return EMPTY, 9.0
	default:
		panic("Bad prefix-operator for " + op.Value)
	}
}

func GetInfixBindingPower(op lexing.Token) (float32, float32) {
	switch op.Type {
	case lexing.Plus_a, lexing.Minus_a:
		return 1.0, 2.0
	case lexing.Div_a, lexing.Mult_a, lexing.Mod_a, lexing.Divflat_a:
		return 3.0, 4.0
	case lexing.Pow_a:
		return 6.0, 5.0
	case lexing.Gthan_l, lexing.Lthan_l, lexing.GthanEq_l, lexing.LthanEq_l, lexing.Eq_l, lexing.Neq_l:
		return 0.5, 0.4

	case lexing.And_l:
		return 0.3, 0.2
	case lexing.Or_l:
		return 0.1, 0.05
	default:
		panic("Bad infix-operator for " + op.Value)
	}
}

func GetPostfixBindingPower(op lexing.TokType) (float32, float32) {
	switch op {
	case lexing.Decr_a, lexing.Inc_a:
		return 10.0, EMPTY
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
