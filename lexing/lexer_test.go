package lexing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLexer(t *testing.T) {
	type test struct {
		input  *Input
		expect TokenList
	}

	test_number := test{
		input: &Input{'1', '2', '3'},
		expect: TokenList{
			Token{Type: Number, Value: "123", Column: 0},
			Token{Type: EOF, Value: "", Column: 3},
		},
	}

	test_number2 := test{
		input: &Input{'-', '-', '1', '2', '3'},
		expect: TokenList{
			Token{Type: Decr_a, Value: "--", Column: 0},
			Token{Type: Number, Value: "123", Column: 2},
			Token{Type: EOF, Value: "", Column: 5},
		},
	}

	test_comment1 := test{
		input: &Input{'#', ' ', 'h', 'e', 'l', 'l', 'o', ' ', '#'},
		expect: TokenList{
			Token{Type: EOF, Value: "", Column: 8},
		},
	}

	test_comment2 := test{
		input: &Input{'#', '.', ' ', 'h', 'e', 'l', 'l', 'o', ' '},
		expect: TokenList{
			Token{Type: EOF, Value: "", Column: 8},
		},
	}

	test_comment3 := test{
		input: &Input{'#', '.', ' ', 'h', 'e', 'l', 'l', 'o', ' ', '#'},
		expect: TokenList{
			Token{Type: EOF, Value: "", Column: 10},
		},
	}

	test_ident := test{
		input: &Input{'a', 'b', '9', ' ', 'i', 'f'},
		expect: TokenList{
			Token{Type: Ident, Value: "ab9", Column: 0},
			Token{Type: If_k, Value: "if", Column: 4},
			Token{Type: EOF, Value: "", Column: 6},
		},
	}

	t.Run("Validate lexer-expectations", func(t *testing.T) {
		assert.Equal(t, test_number.expect, test_number.input.CreateTokens(), "Lexing numbers failed")
		assert.Equal(t, test_number2.expect, test_number2.input.CreateTokens(), "Lexing decremental-numbers failed")

		//

		assert.PanicsWithValue(t, "Invalid comments", func() {
			test_comment1.input.CreateTokens()
		}, "Pure wrong-comment did not work as expected")
		assert.PanicsWithValue(t, "Comment was never closed", func() {
			test_comment2.input.CreateTokens()
		}, "Pure wrong-comment-no-closing did not work as expected")
		assert.Equal(t, test_comment3.expect, test_comment3.input.CreateTokens(), "Pure comment failed")

		//

		assert.Equal(t, test_ident.expect, test_ident.input.CreateTokens(), "Ident/Keyword failed")
	})
}
