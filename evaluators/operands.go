package evaluators

import "github.com/BinaryGhost/alpaci/lexing"

type Operator struct {
	Type lexing.TokType
}

type BindingPowers = [2]int

const EMPTY = -99
const NO_OP = 0
const ASSIGN = 0.1

type ExpressionKind int

const (
	Infix ExpressionKind = iota
	Postix
	Prefix
)

func (op *Operator) GetBindingpower(kind ExpressionKind) BindingPowers {
	return BindingPowers{1, 2}
}
