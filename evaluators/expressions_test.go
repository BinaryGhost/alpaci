package evaluators

import (
	"github.com/BinaryGhost/alpaci/lexing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func generateTokes(input string) lexing.TokenList {
	runes := []rune(input)
	inp := lexing.Input(runes)

	return inp.CreateTokens()
}

func TestNumbers()  {}
func TestStrings()  {}
func TestBooleans() {}
func TestAnything() {}
