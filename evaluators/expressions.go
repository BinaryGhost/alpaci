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
		if e.Kind == Prefix {
			return UnaryPlus(right)
		}
		return Plus(left, right)
	case lexing.Minus_a:
		if e.Kind == Prefix {
			return UnaryMinus(right)
		}
		return Minus(left, right)
	case lexing.Mult_a:
		return Multiply(left, right)
	case lexing.Div_a:
		return Divide(left, right)
	case lexing.Divflat_a:
		return DivideFlat(left, right)
	case lexing.Pow_a:
		return Power(left, right)
	case lexing.Mod_a:
		return Modulo(left, right)
	case lexing.Inc_a:
		return Increment(left)
	case lexing.Decr_a:
		return Decrement(left)
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
		return Not(right)
	case lexing.Neq_l:
		return NotEquals(left, right)
	case lexing.Eq_l:
		return Equals(left, right)
	default:
		panic("Unknown operator '" + e.Operator.Val + "'")
	}
}

func ParseExpression(tl *lexing.TokenList, min_bp float32) Expression {
	var lhs Expression

	first_tok, err := tl.Current()
	if err != nil {
		panic(err)
	}
	// fmt.Printf("debug: tok -> %s, prec -> %d \n", lexing.TokenToString(first_tok), min_bp)

	switch first_tok.Type {
	case lexing.EOF:
		break
	case lexing.Ident:
		lhs = Expression{Atom: MakeIdentAtom(tl)}
	case lexing.Number:
		lhs = Expression{Atom: MakeNumberAtom(tl)}
	case lexing.Lparenth:
		tl.Next()

		lhs = ParseExpression(tl, 0.0)
		if tok, err := tl.Current(); err != nil || tok.Type != lexing.Rparenth {
			panic("Expected ')'")
		}

		tl.Next()
	case lexing.Plus_a, lexing.Minus_a, lexing.Bang_l:
		_, r_bp := GetPrefixBindingPower(first_tok)
		rhs := ParseExpression(tl, r_bp)

		op := Operator{
			Column: first_tok.Column,
			Type:   first_tok.Type,
			Val:    first_tok.Value,
		}

		lhs = Expression{
			Kind:     Prefix,
			Operator: op,
			Right:    &rhs,
		}
		tl.Next()
	default:
		panic("Bad token")
	}

	for {
		op_tok, err := tl.Current()
		if err != nil {
			panic(err)
		}
		// fmt.Printf("debug-op: '%s' \n", lexing.TokTypeAsString(op_tok.Type))

		if op_tok.Type == lexing.EOF || op_tok.Type == lexing.Rparenth {
			break
		}

		if l_bp, _ := GetPostfixBindingPower(op_tok.Type); IsPostfixOp(op_tok.Type) {
			if l_bp < min_bp {
				break
			}
			tl.Next()

			lhs = Expression{
				Kind: Postfix,
				Operator: Operator{
					Column: op_tok.Column,
					Type:   op_tok.Type,
					Val:    op_tok.Value,
				},
				Left: &lhs,
			}
			continue
		}

		l_bp, r_bp := GetInfixBindingPower(op_tok)
		if l_bp < min_bp {
			break
		}

		tl.Next()

		rhs := ParseExpression(tl, r_bp)

		left := lhs
		new_lhs := Expression{
			Kind: Infix,
			Operator: Operator{
				Column: op_tok.Column,
				Type:   op_tok.Type,
				Val:    op_tok.Value,
			},
			Left:  &left,
			Right: &rhs,
		}

		lhs = new_lhs
	}

	return lhs
}

// So Go says: "Only the top-level value gets the Stringer treatment. Everything else gets the default struct formatting."
// How nice of you, I spend 1-1.5 days doubting if my tree is correct or not

func (e *Expression) String() string {
	switch e.Kind {
	case Infix:
		return fmt.Sprintf("(%s %s %s)", e.Operator.Val, e.Left.String(), e.Right.String())
	case Prefix:
		return fmt.Sprintf("(%s %s)", e.Operator.Val, e.Right.String())
	case Postfix:
		return fmt.Sprintf("(%s %s)", e.Left.String(), e.Operator.Val)
	default: // __base__ — this is the atom/leaf case
		return fmt.Sprintf("%v", e.Atom.Val)
	}
}
