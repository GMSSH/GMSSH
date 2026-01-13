package gvalid

// RequiredValidator validates that a field contains a non-empty value.
// It implements the Validator interface for mandatory field validation.
//
// The validator checks that:
// - The value is not nil/zero/empty
// - For strings: not empty string
// - For numbers: not zero
// - For pointers: not nil
// - For slices/maps: not empty
//
// This should typically be used as the first validator in a validation chain.
type RequiredValidator struct {
	ValidatorErrorMessage // Embedded for custom error message support
}

// Validate checks if the field contains a non-empty value.
//
// Parameters:
//   - field: Contains field metadata (name, type, etc.)
//   - value: The value to validate
//
// Returns:
//   - error: ValidationError if value is empty, nil if valid
//
// Validation Rules:
//  1. Nil values fail validation
//  2. Zero/empty values fail validation
//  3. All other values pass validation
//
// Example Usage:
//
//	type User struct {
//	    Username string `validate:"required"`  // Must be non-empty
//	    Age      int    `validate:"required"`  // Must be non-zero
//	}
func (v *RequiredValidator) Validate(field *FieldInfo, value any) error {
	if isEmpty(value) {
		return v.NewValidationError(
			field.Field.Name,
			"is required",
		).WithMessage()
	}
	return nil
}
