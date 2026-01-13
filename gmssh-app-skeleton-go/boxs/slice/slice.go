package slice

import (
	"strconv"
	"strings"
)

// JoinIntSlice concatenates the elements of an integer slice into a single string,
// with elements separated by spaces. This is more efficient than repeated string
// concatenation for large slices.
//
// Parameters:
//   - array: the slice of integers to be joined
//
// Returns:
//   - string: the resulting string with space-separated integer values
func JoinIntSlice(array []int) string {
	// Use strings.Builder for efficient string concatenation
	var builder strings.Builder

	// Iterate through each element in the slice
	for i, item := range array {
		// Add a space separator before all elements except the first one
		if i > 0 {
			builder.WriteByte(' ')
		}

		// Convert the integer to string and append to the builder
		builder.WriteString(strconv.Itoa(item))
	}

	// Return the final concatenated string
	return builder.String()
}
