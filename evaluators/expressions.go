package evaluators

import "github.com/BinaryGhost/alpaci/lexing"

type Expression struct {
	operator Operator
	operands []Atom
	children []Expression
}

func ParseExpression(tl *lexing.TokenList) {
	// get 1st token (= first)
	// is it an operator or atom?
	//
	// peek to next token
	// 	first is operator -> eof -> error
	// 	first is operator -> operator -> error
	// 	first is operator -> atom -> no error if (+/-)
	// consume token
	// next Token
	// ---------------

}

func (e *Expression) String() string { return "" }
