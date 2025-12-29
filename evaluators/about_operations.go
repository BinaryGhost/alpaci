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

func UnaryMinus(a any) any {
	if a_val, ok := a.(int); ok {
		return a_val * -1
	}

	if a_val, ok := a.(float64); ok {
		return a_val * -1
	}

	panic("")
}

func UnaryPlus(a any) any {
	if a_val, ok := a.(int); ok {
		return a_val * 1
	}

	if a_val, ok := a.(float64); ok {
		return a_val * 1
	}

	panic("")
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

func Divide(a, b any) any {
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

func DivideFlat(a, b any) any {
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

func Modulo(a, b any) any {
	if b == 0 || b == 0.0 {
		panic("")
	}

	switch a_val := a.(type) {
	case int:
		if b_val, ok := b.(int); ok {
			return a_val % b_val
		}

		panic("")
	default:
		panic("")
	}
}

func Power(a, b any) any {
	switch a_val := a.(type) {
	case int:
		if b_val, ok := b.(int); ok {
			return math.Pow(float64(a_val), float64(b_val))
		}
		if b_val, ok := b.(float64); ok {
			return math.Pow(float64(a_val), b_val)
		}

		panic("")
	case float64:
		if b_val, ok := b.(float64); ok {
			return math.Pow(a_val, b_val)
		}
		if b_val, ok := b.(int); ok {
			return math.Pow(a_val, float64(b_val))
		}

		panic("")
	default:
		panic("")
	}
}

func Increment(a any) any {
	if a_val, ok := a.(int); ok {
		return a_val + 1
	}

	if a_val, ok := a.(float64); ok {
		return a_val + 1
	}

	panic("")
}

func Decrement(a any) any {
	if a_val, ok := a.(int); ok {
		return a_val - 1
	}

	if a_val, ok := a.(float64); ok {
		return a_val - 1
	}

	panic("")
}

func tryMakeFalsy(val any) bool {
	if val == "" {
		return false
	} else if val == 0 || val == 0.0 {
		return false
	} else if vl, ok := val.(bool); ok {
		return vl
	} else {
		return true
	}
}

func And(a, b any) any {
	return tryMakeFalsy(a) && tryMakeFalsy(b)
}

func Or(a, b any) any {
	return tryMakeFalsy(a) || tryMakeFalsy(b)
}

func GreaterThan(a, b any) any {
	switch a_val := a.(type) {
	case int:
		if b_val, ok := b.(int); ok {
			return a_val > b_val
		}
		if b_val, ok := b.(float64); ok {
			return float64(a_val) > b_val
		}

		panic("")
	case float64:
		if b_val, ok := b.(float64); ok {
			return a_val > b_val
		}
		if b_val, ok := b.(int); ok {
			return a_val > float64(b_val)
		}

		panic("")
	default:
		panic("")
	}
}
func LesserThan(a, b any) any {
	switch a_val := a.(type) {
	case int:
		if b_val, ok := b.(int); ok {
			return a_val < b_val
		}
		if b_val, ok := b.(float64); ok {
			return float64(a_val) < b_val
		}

		panic("")
	case float64:
		if b_val, ok := b.(float64); ok {
			return a_val < b_val
		}
		if b_val, ok := b.(int); ok {
			return a_val < float64(b_val)
		}

		panic("")
	default:
		panic("")
	}
}

func Equals(a, b any) any {
	switch a_val := a.(type) {
	case int:
		if b_val, ok := b.(int); ok {
			return a_val == b_val
		}
		if b_val, ok := b.(float64); ok {
			return float64(a_val) == b_val
		}

		return a == b
	case float64:
		if b_val, ok := b.(float64); ok {
			return a_val == b_val
		}
		if b_val, ok := b.(int); ok {
			return a_val == float64(b_val)
		}

		return a == b
	default:
		return a == b
	}
}
func GreaterThanEquals(a, b any) any {
	switch a_val := a.(type) {
	case int:
		if b_val, ok := b.(int); ok {
			return a_val >= b_val
		}
		if b_val, ok := b.(float64); ok {
			return float64(a_val) >= b_val
		}

		panic("")
	case float64:
		if b_val, ok := b.(float64); ok {
			return a_val >= b_val
		}
		if b_val, ok := b.(int); ok {
			return a_val >= float64(b_val)
		}

		panic("")
	default:
		panic("")
	}
}
func LesserThanEquals(a, b any) any {
	switch a_val := a.(type) {
	case int:
		if b_val, ok := b.(int); ok {
			return a_val <= b_val
		}
		if b_val, ok := b.(float64); ok {
			return float64(a_val) <= b_val
		}

		panic("")
	case float64:
		if b_val, ok := b.(float64); ok {
			return a_val <= b_val
		}
		if b_val, ok := b.(int); ok {
			return a_val <= float64(b_val)
		}

		panic("")
	default:
		panic("")
	}
}

func Not(a any) any {
	if a_val, ok := a.(bool); ok {
		return !a_val
	}
	panic("")
}

func NotEquals(a, b any) any {
	switch a_val := a.(type) {
	case int:
		if b_val, ok := b.(int); ok {
			return a_val != b_val
		}
		if b_val, ok := b.(float64); ok {
			return float64(a_val) != b_val
		}

		return a != b
	case float64:
		if b_val, ok := b.(float64); ok {
			return a_val != b_val
		}
		if b_val, ok := b.(int); ok {
			return a_val != float64(b_val)
		}

		return a != b
	default:
		return a != b
	}
}
