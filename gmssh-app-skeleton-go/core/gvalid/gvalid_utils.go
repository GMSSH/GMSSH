package gvalid

import (
	"reflect"
	"strconv"
)

// parseInt converts a string to an integer with error handling.
// It uses strconv.ParseInt internally with base 10 and 64-bit precision.
//
// Parameters:
//   - s: The string to convert to an integer
//
// Returns:
//   - int: The converted integer value
//   - error: Any conversion error that occurred
//
// Example:
//
//	val, err := parseInt("42") // val = 42, err = nil
func parseInt(s string) (val int, err error) {
	val64, err := strconv.ParseInt(s, 10, 64)
	return int(val64), err
}

// parseFloat converts a string to a float64 with error handling.
// It uses strconv.ParseFloat internally with 64-bit precision.
//
// Parameters:
//   - s: The string to convert to a float
//
// Returns:
//   - float64: The converted floating-point value
//   - error: Any conversion error that occurred
//
// Example:
//
//	val, err := parseFloat("3.14") // val = 3.14, err = nil
func parseFloat(s string) (val float64, err error) {
	return strconv.ParseFloat(s, 64)
}

// toFloat attempts to convert an interface{} value to float64.
// This is useful for handling unknown numeric types in validation.
//
// Parameters:
//   - v: The value to convert (expected to be float64 compatible)
//
// Returns:
//   - float64: The converted value if successful
//   - bool: True if conversion was successful, false otherwise
//
// Example:
//
//	val, ok := toFloat(3.14) // val = 3.14, ok = true
func toFloat(v any) (val float64, ok bool) {
	val, ok = v.(float64)
	return val, ok
}

// isEmpty checks if a value is empty/nil/zero.
// It handles various types including pointers, collections, and primitives.
//
// Parameters:
//   - i: The value to check
//
// Returns:
//   - bool: True if the value is empty/nil/zero, false otherwise
//
// Supported types:
//   - nil values
//   - Pointers, maps, slices, channels, functions, interfaces
//   - Strings (empty string)
//   - Numeric types (zero values)
//   - Booleans (false)
//   - Structs (zero-valued structs)
//
// Example:
//
//	isEmpty("") // true
//	isEmpty(0)  // true
//	isEmpty(struct{}{}) // true if struct has zero values
func isEmpty(i any) bool {
	if i == nil {
		return true
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func, reflect.Interface:
		return v.IsNil()
	case reflect.String:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Struct:
		return reflect.DeepEqual(i, reflect.Zero(v.Type()).Interface())
	default:
		return false
	}
}
