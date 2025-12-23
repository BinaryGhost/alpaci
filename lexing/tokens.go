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
	Inc_a
	Decr_a
	Rparenth
	Lparenth
	Eq
	Gthan_l
	Lthan_l
	GthanEq_l
	LthanEq_l
	Neq_l
	Pipe
	And
	Eq_l
	And_l
	Or_l
	Arrow
	Colon
	Comma
	Bang_l
	Number
	Ident
	Dollar
	Semicolon
	Rbrack
	Rbrace
	Lbrack
	Lbrace
	Tilde
	Qmark
	Slash
	Point
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
	EOF
)
