package cond

// If simulates a ternary conditional operator (condition ? a : b) found in other languages.
// It takes a boolean condition and two values of the same type T.
// If the condition is true, it returns the first value (a), otherwise it returns the second value (b).
//
// Generic type parameter:
//   - T: the type of the values to return (can be any type)
//
// Parameters:
//   - condition: the boolean expression to evaluate
//   - a: value to return if condition is true
//   - b: value to return if condition is false
//
// Returns:
//   - The selected value based on the condition (either a or b)
func If[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}
