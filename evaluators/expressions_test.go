package evaluators

import (
	"github.com/BinaryGhost/alpaci/lexing"
	"github.com/stretchr/testify/assert"
	"testing"
)

type test_expression struct {
	input          string
	expected_value any
}

func evaluateExpression(input string) any {
	str := []rune(input)
	inp := lexing.Input(str)

	tl := inp.CreateTokens()

	expr := WrapForExpression(&tl)

	return Eval(&expr)
}

func TestNumbers(t *testing.T) {
	bunch_of_tests := []test_expression{
		{input: "1", expected_value: 1.0},
		{input: "-1", expected_value: -1.0},
		{input: "(1)-- // 8", expected_value: 0.0},
		{input: "-(2 ** 3)", expected_value: -8.0},
		{input: "((((1 - 2) * 2))) - 4", expected_value: -6.0},
		{input: "((2))", expected_value: 2.0},
		{input: "-2 + -3", expected_value: -5.0},
		{input: "2++ + -1-- -1", expected_value: 2.0},
		{input: "7 - + - + - 1", expected_value: 6.0},
		{input: "2 ** 2 * 2 + 2", expected_value: 10.0},
	}

	// bunch_of_panics := []test_expression{
	// 	{input: "--1", expected_value: 0},              // should panic
	// 	{input: "7 / 0", expected_value: 0},            // should panic
	// 	{input: "(2 * 9.7 (1 -2))", expected_value: 0}, // should panic
	// 	{input: "8 + 9)", expected_value: 0},           // should panic
	// 	{input: "(8 + 1", expected_value: 0},           // should panic
	// 	{input: "7 ^ 9", expected_value: 0},            // should panic
	// 	{input: "2 ** 2 ** 2", expected_value: 16},     // should panic
	// }

	t.Run("Look if numberic-expressions are correct", func(t *testing.T) {
		for _, test := range bunch_of_tests {
			assert.Equal(t, test.expected_value, evaluateExpression(test.input))
		}
	})
}
func TestStrings(t *testing.T)  {}
func TestBooleans(t *testing.T) {}
func TestAnything(t *testing.T) {}
