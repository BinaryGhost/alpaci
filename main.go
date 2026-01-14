package main

import (
	"fmt"
	"github.com/BinaryGhost/gospel/evaluators"
	"github.com/BinaryGhost/gospel/lexing"
)

func main() {

	// A lil bit peeking
	str := []rune("!(false and true)")
	inp := lexing.Input(str)

	tl := inp.CreateTokens()

	expr := evaluators.WrapForExpression(&tl)

	fmt.Println(expr.String())
	fmt.Println(evaluators.Eval(&expr))

}
