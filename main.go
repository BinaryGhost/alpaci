package main

import (
	"fmt"
	"github.com/BinaryGhost/alpaci/evaluators"
	"github.com/BinaryGhost/alpaci/lexing"
)

func main() {

	// A lil bit peeking
	str := []rune("i++")
	inp := lexing.Input(str)

	tl := inp.CreateTokens()

	var zero int8
	expr := evaluators.ParseExpression(&tl, 0.0, &zero)

	fmt.Println(expr.String())
	fmt.Println(evaluators.Eval(&expr))

}
