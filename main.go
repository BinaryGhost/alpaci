package main

import (
	"fmt"
	"github.com/BinaryGhost/alpaci/evaluators"
	"github.com/BinaryGhost/alpaci/lexing"
)

func main() {

	// A lil bit peeking
	str := []rune("'hi' + 'hi'")
	inp := lexing.Input(str)

	tl := inp.CreateTokens()

	expr := evaluators.WrapForExpression(&tl)

	fmt.Println(expr.String())
	fmt.Println(evaluators.Eval(&expr))

}
