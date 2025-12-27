package lexing

import (
	"errors"
	"unicode"
)

type Input []rune
type TokenList []Token

func (input *Input) peekNext(index int, expected rune) bool {
	return index+1 < len(*input) && (*input)[index+1] == expected
}

func (tlist *TokenList) appendSymbol(ttyp TokType, value string, index *int) {
	*tlist = append(*tlist, Token{Type: ttyp, Value: value, Column: *index})
	*index += len(value)
}

func (tl *TokenList) Current() (Token, error) {
	if len(*tl) == 0 {
		return Token{}, errors.New("Tried to see current Token, but TokenList is already empty")
	}
	return (*tl)[0], nil
}

func (tl *TokenList) Next() error {
	if len(*tl) == 0 {
		return errors.New("Tried to advance, but TokenList is already empty")
	}

	*tl = (*tl)[1:]
	return nil
}

func (input *Input) CreateTokens() TokenList {
	tl := TokenList{}
	ilen := len(*input)
	i := 0

	for i < ilen {
		tok := (*input)[i]

		switch {
		case tok == '#':
			if !input.peekNext(i, '.') {
				panic("Invalid comments")
			}
			i++

			for i < ilen && (*input)[i] != '#' {
				i++
			}

			if i >= ilen {
				panic("Comment was never closed")
			}
			i++

		case tok == ' ', tok == '\t', tok == '\r':
			i++
			continue

		case tok == '*':
			if input.peekNext(i, '*') {
				if input.peekNext(i+1, '=') {
					tl.appendSymbol(Pow_Eq_a, "**=", &i)
				} else {
					tl.appendSymbol(Pow_a, "**", &i)
				}
			} else if input.peekNext(i, '=') {
				tl.appendSymbol(Mult_Eq_a, "*=", &i)
			} else {
				tl.appendSymbol(Mult_a, "*", &i)
			}

		case tok == '/':
			if input.peekNext(i, '/') {
				if input.peekNext(i+1, '=') {
					tl.appendSymbol(Divflat_Eq_a, "//=", &i)
				} else {
					tl.appendSymbol(Divflat_a, "//", &i)
				}
			} else if input.peekNext(i, '=') {
				tl.appendSymbol(Div_Eq_a, "/=", &i)
			} else {
				tl.appendSymbol(Div_a, "/", &i)
			}

		case tok == '+':
			if input.peekNext(i, '+') {
				tl.appendSymbol(Inc_a, "++", &i)
			} else if input.peekNext(i, '=') {
				tl.appendSymbol(Plus_Eq_a, "+=", &i)
			} else {
				tl.appendSymbol(Plus_a, "+", &i)
			}

		case tok == '-':
			if input.peekNext(i, '-') {
				tl.appendSymbol(Decr_a, "--", &i)
			} else if input.peekNext(i, '>') {
				tl.appendSymbol(Arrow, "->", &i)
			} else if input.peekNext(i, '=') {
				tl.appendSymbol(Minus_Eq_a, "-=", &i)
			} else {
				tl.appendSymbol(Minus_a, "-", &i)
			}

		case tok == '%':
			if input.peekNext(i, '=') {
				tl.appendSymbol(Mod_Eq_a, "%=", &i)
			} else {
				tl.appendSymbol(Mod_a, "%", &i)
			}

		case tok == '(':
			tl.appendSymbol(Lparenth, "(", &i)

		case tok == ')':
			tl.appendSymbol(Rparenth, ")", &i)

		case tok == '.':
			tl.appendSymbol(Point, ".", &i)
		case tok == '!':
			if input.peekNext(i, '=') {
				tl.appendSymbol(Neq_l, "!=", &i)
			} else {
				tl.appendSymbol(Bang_l, "!", &i)
			}
		case tok == '|':
			if input.peekNext(i, '|') {
				tl.appendSymbol(Or_l, "||", &i)
			} else {
				tl.appendSymbol(Pipe, "|", &i)
			}
		case tok == '&':
			if input.peekNext(i, '&') {
				tl.appendSymbol(And_l, "&&", &i)
			} else {
				tl.appendSymbol(And, "&&", &i)
			}
		case tok == '<':
			if input.peekNext(i, '=') {
				tl.appendSymbol(LthanEq_l, "<=", &i)
			} else {
				tl.appendSymbol(Lthan_l, "<", &i)
			}
		case tok == '>':
			if input.peekNext(i, '=') {
				tl.appendSymbol(GthanEq_l, ">=", &i)
			} else {
				tl.appendSymbol(Gthan_l, ">", &i)
			}
		case tok == '=':
			if input.peekNext(i, '=') {
				tl.appendSymbol(Eq_l, "==", &i)
			} else {
				tl.appendSymbol(Eq, "=", &i)
			}
		case tok == ',':
			tl.appendSymbol(Comma, ",", &i)
		case tok == '~':
			tl.appendSymbol(Tilde, "~", &i)
		case tok == ':':
			tl.appendSymbol(Colon, ":", &i)
		case tok == ';':
			tl.appendSymbol(Semicolon, ";", &i)
		case tok == '[':
			tl.appendSymbol(Lbrack, "[", &i)
		case tok == ']':
			tl.appendSymbol(Rbrack, "]", &i)
		case tok == '{':
			tl.appendSymbol(Lbrace, "{", &i)
		case tok == '}':
			tl.appendSymbol(Rbrace, "}", &i)
		case tok == '$':
			tl.appendSymbol(Dollar, "$", &i)
		case tok == '\\':
			tl.appendSymbol(Slash, "\\", &i)
		case tok == '?':
			tl.appendSymbol(Qmark, "?", &i)

		// TODO: Unimplemented -> Implement this
		case tok == '\'', tok == '"', tok == '`':
			continue

		case unicode.IsDigit(tok):
			start := i
			for i < ilen && (unicode.IsDigit((*input)[i]) || (*input)[i] == '_') {
				i++
			}
			number := string((*input)[start:i])
			tl = append(tl, Token{Type: Number, Value: number, Column: start})

		case unicode.IsLetter(tok):
			start := i
			for i < ilen && (unicode.IsLetter((*input)[i]) || unicode.IsDigit((*input)[i])) {
				i++
			}

			ident := string((*input)[start:i])
			switch ident {
			case "and":
				tl = append(tl, Token{Type: And_l, Value: "and", Column: start})
			case "or":
				tl = append(tl, Token{Type: Or_l, Value: "or", Column: start})
			case "forc":
				tl = append(tl, Token{Type: Forc_k, Value: ident, Column: start})
			case "for":
				tl = append(tl, Token{Type: For_k, Value: ident, Column: start})
			case "in":
				tl = append(tl, Token{Type: In_k, Value: ident, Column: start})
			case "switch":
				tl = append(tl, Token{Type: Switch_k, Value: ident, Column: start})
			case "while":
				tl = append(tl, Token{Type: While_k, Value: ident, Column: start})
			case "if":
				tl = append(tl, Token{Type: If_k, Value: ident, Column: start})
			case "elif":
				tl = append(tl, Token{Type: Elif_k, Value: ident, Column: start})
			case "else":
				tl = append(tl, Token{Type: Else_k, Value: ident, Column: start})
			case "case":
				tl = append(tl, Token{Type: Case_k, Value: ident, Column: start})
			case "default":
				tl = append(tl, Token{Type: Default_k, Value: ident, Column: start})
			case "bool":
				tl = append(tl, Token{Type: Bool_k, Value: ident, Column: start})
			case "num":
				tl = append(tl, Token{Type: Num_k, Value: ident, Column: start})
			case "str":
				tl = append(tl, Token{Type: Str_k, Value: ident, Column: start})
			case "false":
				tl = append(tl, Token{Type: False_k, Value: ident, Column: start})
			case "true":
				tl = append(tl, Token{Type: True_k, Value: ident, Column: start})
			case "end":
				tl = append(tl, Token{Type: End_k, Value: ident, Column: start})
			default:
				tl = append(tl, Token{Type: Ident, Value: ident, Column: start})
			}

		default:
			panic("Unknown symbol found")
		}
	}

	tl = append(tl, Token{Type: EOF, Value: "", Column: i})
	return tl
}
