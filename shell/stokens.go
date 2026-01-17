package shell

type Kind int

const (
	ImpureIdentifier Kind = iota + 1
	PureIdentifier        // Does contain only letters
	String
	Number
	Boolean
	Array
	Map
	Function
	//
	Colon
	Rparenth
	Lparenth
	Rbracket
	Lbracket
	Rbrace
	Lbrace
	And
	Gthan
	Lthan
	Pipe
	Bang
	Eq
	Qmark
	Point
	Semicolon
	Comma
	//
	ManualNewline
	VariableCmd
	keyword
	Type
	BooleanValue
)

type ShellToken struct {
	Representation string
	IsGospelCmd    bool
	IsSomething    bool // To denote, that it might be used for a command
	Spos           int  // start positioning
	Typ            Kind
}

var Keywords = []string{
	"end", "return", "for", "is", "do", "to", "from", "while", "if", "elif", "else", "switch", "case", "default",
}

var Types = []string{
	"str", "num", "bool", "array", "map",
}

var SpecialChar = []rune{
	':', '(', ')', '[', ']', '{', '}', '&', '>', '<', '|', '!', '=', '?', '.', ';', ',',
}

type char interface {
	~rune | ~string
}

func IsExcluded[T char](s T, excluded []T) (bool, T) {
	for _, word := range excluded {
		if s == word {
			return true, word
		}
	}
	return false, s
}
