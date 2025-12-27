package evaluators

import "github.com/BinaryGhost/alpaci/lexing"

type Expression struct {
	Kind     ExpressionKind
	Operator Operator
	Left     *Expression
	Right    *Expression
	Atom     Atom
}

func (e *Expression) Eval() {}

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
