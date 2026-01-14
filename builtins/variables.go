package builtins

type scopes int

const (
	outside scopes = iota + 1
	global
	local
)

type types int

const (
	_array types = iota + 1
	_map
	_number
	_string
	_bool
	_access
)

type VariableValue struct {
	Attr  types
	Value any
	Scope scopes
}

type Lookup map[string]VariableValue

func BreatheLifeInto(l *Lookup, v_name string, info VariableValue) {}

func ChangeInto(l *Lookup, v_name string, info VariableValue) {}

func AskInto(l *Lookup, v_name string) {}

func BanishInto(l *Lookup, v_name string) {}
