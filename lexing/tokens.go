package lexing

import "fmt"

type Token struct {
	Type   TokType
	Value  string
	Column int
}

func TokenToString(t Token) string {
	return fmt.Sprintf("Token {Value: '%s', Type: %s, Column: %d}", t.Value, TokTypeAsString(t.Type), t.Column)
}

func TokTypeAsString(tt TokType) string {
	switch tt {
	case Mult_a:
		return "Mult_a"
	case Div_a:
		return "Div_a"
	case Plus_a:
		return "Plus_a"
	case Pow_a:
		return "Pow_a"
	case Divflat_a:
		return "Divflat_a"
	case Minus_a:
		return "Minus_a"
	case Mod_a:
		return "Mod_a"
	case Mult_Eq_a:
		return "Mult_Eq_a"
	case Div_Eq_a:
		return "Div_Eq_a"
	case Plus_Eq_a:
		return "Plus_Eq_a"
	case Pow_Eq_a:
		return "Pow_Eq_a"
	case Divflat_Eq_a:
		return "Divflat_Eq_a"
	case Minus_Eq_a:
		return "Minus_Eq_a"
	case Mod_Eq_a:
		return "Mod_Eq_a"
	case Inc_a:
		return "Inc_a"
	case Decr_a:
		return "Decr_a"
	case Gthan_l:
		return "Gthan_l"
	case Lthan_l:
		return "Lthan_l"
	case GthanEq_l:
		return "GthanEq_l"
	case LthanEq_l:
		return "LthanEq_l"
	case Bang_l:
		return "Bang_l"
	case Neq_l:
		return "Neq_l"
	case Eq_l:
		return "Eq_l"
	case And_l:
		return "And_l"
	case Or_l:
		return "Or_l"
	case Eq:
		return "Eq"
	case Arrow:
		return "Arrow"
	//
	case Pipe:
		return "Pipe"
	case And:
		return "And"
	case Colon:
		return "Colon"
	case Comma:
		return "Comma"
	case Number:
		return "Number"
	case Ident:
		return "Ident"
	case Dollar:
		return "Dollar"
	case Semicolon:
		return "Semicolon"
	case Rparenth:
		return "Rparenth"
	case Lparenth:
		return "Lparenth"
	case Rbrack:
		return "Rbrack"
	case Rbrace:
		return "Rbrace"
	case Lbrack:
		return "Lbrack"
	case Lbrace:
		return "Lbrace"
	case Tilde:
		return "Tilde"
	case Qmark:
		return "Qmark"
	case Slash:
		return "Slash"
	case Point:
		return "Point"
	//
	case In_k:
		return "In_k"
	case Num_k:
		return "Num_k"
	case Str_k:
		return "Str_k"
	case Bool_k:
		return "Bool_k"
	case False_k:
		return "False_k"
	case True_k:
		return "True_k"
	case If_k:
		return "If_k"
	case Elif_k:
		return "Elif_k"
	case Else_k:
		return "Else_k"
	case For_k:
		return "For_k"
	case Forc_k:
		return "Forc_k"
	case While_k:
		return "While_k"
	case Switch_k:
		return "Switch_k"
	case Case_k:
		return "Case_k"
	case Default_k:
		return "Default_k"
	case End_k:
		return "End_k"
	//
	case EOF:
		return "EOF"
	default:
		panic("Could not turn this type into a token")
	}
}

type TokType int

const (
	Mult_a TokType = iota
	Div_a
	Plus_a
	Pow_a
	Divflat_a
	Minus_a
	Mod_a
	Mult_Eq_a
	Div_Eq_a
	Plus_Eq_a
	Pow_Eq_a
	Divflat_Eq_a
	Minus_Eq_a
	Mod_Eq_a
	Inc_a
	Decr_a
	Gthan_l
	Lthan_l
	GthanEq_l
	LthanEq_l
	Bang_l
	Neq_l
	Eq_l
	And_l
	Or_l
	Eq
	Arrow
	//
	Pipe
	And
	Colon
	Comma
	Number
	Ident
	Dollar
	Semicolon
	Rparenth
	Lparenth
	Rbrack
	Rbrace
	Lbrack
	Lbrace
	Tilde
	Qmark
	Slash
	Point
	//
	In_k
	Num_k
	Str_k
	Bool_k
	False_k
	True_k
	If_k
	Elif_k
	Else_k
	For_k
	Forc_k
	While_k
	Switch_k
	Case_k
	Default_k
	End_k
	//
	EOF
)
