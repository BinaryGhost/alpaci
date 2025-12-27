package errors

type LexingError int

const (
	UnclosedComment LexingError = iota
	InvalidComment
	UnclosedString
	UnknownCharacter
)

type ParsingError int
