package shell

// slexer -> 's'hell lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type ShellInput []rune

type ShellTokenList []ShellToken

func (si *ShellInput) peek(index int, expected rune) bool {
	return index+1 < len(*si) && (*si)[index+1] == expected
}

func AppendSpecialChar(lst *ShellTokenList, stok ShellToken, index *int) {
	*lst = append(*lst, stok)
	*index++
}

func (si *ShellInput) MakeShellTokens() {
	var all_tokens ShellTokenList
	slen := len(*si)
	i := 0

	for i < slen {
		stok := (*si)[i]

		switch stok {
		case '(':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: "(",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Lparenth,
			}, &i)
		case ')':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: ")",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Rparenth,
			}, &i)
		case '[':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: "[",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Lbracket,
			}, &i)
		case ']':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: "]",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Rbracket,
			}, &i)
		case '{':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: "{",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Lbrace,
			}, &i)
		case '}':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: "}",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Rbrace,
			}, &i)
		case '\\':
			/* TODO:
			Implement ManualNewline -> \\n

			Example
				cmd1 -flag \
					--word

			*/
			continue
		case ':':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: ":",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Colon,
			}, &i)
		case '&':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: ":",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            And,
			}, &i)
		case '>':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: ":",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Gthan,
			}, &i)
		case '<':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: ":",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Lthan,
			}, &i)
		case '|':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: ":",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Pipe,
			}, &i)
		case '!':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: ":",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Bang,
			}, &i)
		case '=':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: ":",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Eq,
			}, &i)
		case '?':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: ":",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Qmark,
			}, &i)
		case '.':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: ":",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Point,
			}, &i)
		case ';':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: ";",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Semicolon,
			}, &i)
		case ',':
			AppendSpecialChar(&all_tokens, ShellToken{
				Representation: ";",
				IsGospelCmd:    false,
				IsSomething:    false,
				Spos:           i,
				Typ:            Comma,
			}, &i)

		case '#':
			if !si.peek(i, '.') {
				panic("Forgot a point after '#' ?")
			}
			i++

			for i < slen && (*si)[i] != '#' {
				i++
			}

			if i >= slen {
				panic("Comment was never closed")
			}
			i++
		case ' ', '\t', '\r':
			i++

		case '\'', '"', '`':
			i++

			start_pos := i
			has_cr := false

			for i < slen && (*si)[i] != stok {
				i++

				cur_char := (*si)[i]

				if cur_char < 0 {
					panic("String not properly terminated")
				}

				if cur_char == '\r' {
					has_cr = true
				}

				if cur_char == '\\' && i+1 < slen && (*si)[i+1] == stok {
					i++
				}
			}

			if i >= slen {
				error := fmt.Sprintf("String(%s ... ->%s<-) was never closed", string(stok), string(stok))
				panic(error)
			}

			extracted_string := string((*si)[start_pos:i])
			var raw_string string

			if has_cr {
				raw_string = strings.ReplaceAll(extracted_string, "\r", "")
			} else {
				raw_string = extracted_string
			}

			// TODO: Handle later
			// tl = append(tl, Token{Type: String, Value: raw_string, Column: start_str})
			fmt.Println("raw_string: " + raw_string)
			i++

		case '$':
			all_tokens = append(all_tokens, ShellToken{
				Representation: "$",
				IsGospelCmd:    true,
				IsSomething:    false,
				Spos:           i,
				Typ:            VariableCmd,
			})

			i++

			start := i
			if !unicode.IsLetter((*si)[i]) {
				panic(fmt.Sprintf("Variable names have to start with letters, not '%s'", string((*si)[i])))
			}

			for i < slen && (unicode.IsLetter((*si)[i]) || unicode.IsDigit((*si)[i])) {
				i++
			}

			ident := string((*si)[start:i])

			ls := append([]string{}, Keywords...)
			ls = append(ls, Types...)
			if ok, kw := IsExcluded(ident, ls); ok {
				panic(fmt.Sprintf("Can not use keyword '%s' as an identifier", kw))
			}

			all_tokens = append(all_tokens, ShellToken{
				Representation: ident,
				IsGospelCmd:    false,
				IsSomething:    true,
				Spos:           start,
				Typ:            PureIdentifier,
			})

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			start := i
			point_counter := 0
			seems_pure := true

			for i < slen && (unicode.IsDigit((*si)[i]) || (*si)[i] == '.' || unicode.IsLetter((*si)[i])) {
				if (*si)[i] == '.' {
					point_counter++
				}

				if unicode.IsLetter((*si)[i]) {
					seems_pure = false
				}

				i++
			}

			number := (*si)[start:i]
			if point_counter > 1 || !seems_pure {
				all_tokens = append(all_tokens, ShellToken{
					Representation: string(number),
					IsGospelCmd:    false,
					IsSomething:    true,
					Spos:           start,
					Typ:            ImpureIdentifier,
				})
			} else {
				all_tokens = append(all_tokens, ShellToken{
					Representation: string(number),
					IsGospelCmd:    false,
					IsSomething:    true,
					Spos:           start,
					Typ:            Number,
				})
			}

		default:
			start := i
			seems_pure := true

			for i < slen {
				if excluded, r := IsExcluded((*si)[i], SpecialChar); excluded && unicode.IsSpace(r) {
					break
				}

				if !unicode.IsLetter((*si)[i]) {
					seems_pure = false
				}

				i++
			}

			something := string(*si)[start:i]

			kws := append([]string{}, Keywords...)
			tps := append([]string{}, Types...)

			if ok, kw := IsExcluded(something, kws); ok {
				all_tokens = append(all_tokens, ShellToken{
					Representation: kw,
					IsGospelCmd:    true,
					IsSomething:    false,
					Spos:           start,
					Typ:            keyword,
				})
			} else if ok, tp := IsExcluded(something, tps); ok {
				all_tokens = append(all_tokens, ShellToken{
					Representation: tp,
					IsGospelCmd:    true,
					IsSomething:    false,
					Spos:           start,
					Typ:            Type,
				})
			} else if something == "false" || something == "true" {
				all_tokens = append(all_tokens, ShellToken{
					Representation: something,
					IsGospelCmd:    false,
					IsSomething:    true,
					Spos:           start,
					Typ:            BooleanValue,
				})
			} else if seems_pure {
				all_tokens = append(all_tokens, ShellToken{
					Representation: something,
					IsGospelCmd:    false,
					IsSomething:    true,
					Spos:           start,
					Typ:            PureIdentifier,
				})
			} else {
				all_tokens = append(all_tokens, ShellToken{
					Representation: something,
					IsGospelCmd:    false,
					IsSomething:    true,
					Spos:           start,
					Typ:            ImpureIdentifier,
				})
			}
		}
	}

}
