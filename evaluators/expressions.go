package evaluators

import (
	"github.com/BinaryGhost/alpaci/lexing"
)

type Expression struct {
	Kind     ExpressionKind
	Operator Operator
	Left     *Expression
	Right    *Expression
	Atom     Atom
}

func Eval(e *Expression) any {
	// Since an atom can also be an expression, but with no kind
	if e.Kind == 0 {
		// TODO: Handle booleans and strings and identifiers
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
	case lexing.Pow_a:
		// return Power(left, right)
	case lexing.Mod_a:
		// return MOdulo(left, right)
	case lexing.Inc_a:
		// return Increment(...)
	case lexing.Decr_a:
		// return Decrement(...)
	//
	case lexing.Gthan_l:
		return GreaterThan(left, right)
	case lexing.Lthan_l:
		return LesserThan(left, right)
	case lexing.GthanEq_l:
		return GreaterThanEquals(left, right)
	case lexing.LthanEq_l:
		return LesserThanEquals(left, right)
	case lexing.Bang_l:
		// return Bang(...)
	case lexing.Neq_l:
		return NotEquals(left, right)
	case lexing.Eq_l:
		return Equals(left, right)

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
