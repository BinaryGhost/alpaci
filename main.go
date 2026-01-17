package main

import (
	"fmt"
	"github.com/BinaryGhost/gospel/evaluators"
	elex "github.com/BinaryGhost/gospel/evaluators/lexing"
)

func main() {

	// A lil bit peeking
	str := []rune("!(false and true)")
	inp := elex.Input(str)

	tl := inp.CreateTokens()

	expr := evaluators.WrapForExpression(&tl)

	fmt.Println(expr.String())
	fmt.Println(evaluators.Eval(&expr))

}
