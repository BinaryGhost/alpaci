package evaluators

import (
	"fmt"
	"github.com/BinaryGhost/alpaci/lexing"
)

type Expression struct {
	Kind     ExpressionKind
	Operator Operator
	Left     *Expression
	Right    *Expression
	Atom     Atom
}

func operate_on(a any, b any) any {
	switch v1 := a.(type) {
	case string:
		if v2, ok := b.(string); ok {
			return v1 + v2
		}
	case int:
		if v2, ok := b.(int); ok {
			return v1 + v2 // Add integers
		} else if v2, ok := b.(float64); ok {
			return float64(v1) + v2

		}
	case float64:
		if v2, ok := b.(float64); ok {
			return v1 + v2 // Add floats
		} else if v2, ok := b.(int); ok {
			return v1 + float64(v2)
		}
	}

	return nil
}

func Eval(e *Expression) any {
	// Since an atom can also be an expression, but with no kind
	if e.Kind == 0 {
		return e.Atom.Val
	}

	left := Eval(e.Left)
	right := Eval(e.Right)

	switch e.Operator.Type {
	case lexing.Plus_a:
		return Plus(left, right)
	case lexing.Mult_a:
		return Multiply(left, right)
	case lexing.Div_a:
		return Divide(left, right)
	case lexing.Divflat_a:
		return DivideFlat(left, right)
	default:
		panic("Unknown operator '" + e.Operator.Val + "'")
	}
}

func ParseExpression(tl *lexing.TokenList, min_bp int) Expression {
	var lhs Expression
	first_tok, err := tl.Current()
	if err != nil {
		panic(err)
	}

	switch first_tok.Type {
	case lexing.Ident:
		lhs = Expression{Atom: MakeIdentAtom(tl)}
	case lexing.Number:
		lhs = Expression{Atom: MakeNumberAtom(tl)}
	case lexing.Lparenth:
		tl.Next()

		lhs = ParseExpression(tl, 0)
		if tok, err := tl.Current(); err != nil || tok.Type != lexing.Rparenth {
			panic("Expected ')'")
		}

		tl.Next()
	default:
		panic("Bad token")
	}

	for {
		var op Operator
		op_tok, err := tl.Current()
		if err != nil {
			panic(err)
		}

		if !IsInfixOp(op_tok.Type) {
			break
		}

		l_bp, r_bp := GetInfixBindingPower(op_tok.Type)
		if l_bp < min_bp {
			break
		}

		op.Column = op_tok.Column
		op.Type = op_tok.Type
		op.Val = op_tok.Value

		tl.Next()

		rhs := ParseExpression(tl, r_bp)

		lhs = Expression{
			Kind:     Infix,
			Operator: op,
			Left:     &lhs,
			Right:    &rhs,
		}
	}

	return lhs
}

func (e *Expression) String() string { return "" }
