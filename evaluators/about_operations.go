package evaluators

import "math"

func Plus(a, b any) any {
	switch a_val := a.(type) {
	case int:
		if b_val, ok := b.(int); ok {
			return a_val + b_val
		}
		if b_val, ok := b.(float64); ok {
			return float64(a_val) + b_val
		}

		panic("")
	case float64:
		if b_val, ok := b.(float64); ok {
			return a_val + b_val
		}
		if b_val, ok := b.(int); ok {
			return a_val + float64(b_val)
		}

		panic("")
	case string:
		if b_val, ok := b.(string); ok {
			return a_val + b_val
		}

		panic("")
	default:
		panic("")
	}
}

func Minus(a, b any) any {
	switch a_val := a.(type) {
	case int:
		if b_val, ok := b.(int); ok {
			return a_val - b_val
		}
		if b_val, ok := b.(float64); ok {
			return float64(a_val) - b_val
		}

		panic("")
	case float64:
		if b_val, ok := b.(float64); ok {
			return a_val - b_val
		}
		if b_val, ok := b.(int); ok {
			return a_val - float64(b_val)
		}

		panic("")
	default:
		panic("")
	}
}

func Multiply(a, b any) any {
	switch a_val := a.(type) {
	case int:
		if b_val, ok := b.(int); ok {
			return a_val * b_val
		}
		if b_val, ok := b.(float64); ok {
			return float64(a_val) * b_val
		}

		panic("")
	case float64:
		if b_val, ok := b.(float64); ok {
			return a_val * b_val
		}
		if b_val, ok := b.(int); ok {
			return a_val * float64(b_val)
		}

		panic("")
	default:
		panic("")
	}
}

func Divide(a, b any) float64 {
	if b == 0 || b == 0.0 {
		panic("")
	}

	switch a_val := a.(type) {
	case int:
		if b_val, ok := b.(int); ok {
			return float64(a_val) / float64(b_val)
		}
		if b_val, ok := b.(float64); ok {
			return float64(a_val) / b_val
		}

		panic("")
	case float64:
		if b_val, ok := b.(float64); ok {
			return a_val / b_val
		}
		if b_val, ok := b.(int); ok {
			return a_val / float64(b_val)
		}

		panic("")
	default:
		panic("")
	}
}

func DivideFlat(a, b any) float64 {
	if b == 0 || b == 0.0 {
		panic("")
	}

	switch a_val := a.(type) {
	case int:
		if b_val, ok := b.(int); ok {
			return math.Round(float64(a_val) / float64(b_val))
		}
		if b_val, ok := b.(float64); ok {
			return math.Round(float64(a_val) / b_val)
		}

		panic("")
	case float64:
		if b_val, ok := b.(float64); ok {
			return math.Round(a_val / b_val)
		}
		if b_val, ok := b.(int); ok {
			return math.Round(a_val / float64(b_val))
		}

		panic("")
	default:
		panic("")
	}
}

func Modulo(a, b any) {}

func Power(a, b any) {}

func Increment(a any) {}
func Decrement(a any) {}

func And(a, b any)               {}
func Or(a, b any)                {}
func GreaterThan(a, b any)       {}
func LesserThan(a, b any)        {}
func Equals(a, b any)            {}
func GreaterThanEquals(a, b any) {}
func LesserThanEquals(a, b any)  {}
func Not(a any)                  {}
func NotEquals(a, b any)         {}
