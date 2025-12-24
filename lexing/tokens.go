package lexing

type Token struct {
	Type   TokType
	Value  string
	Column int
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
