package gvalid

// MinLengthValidator validates that a string field meets a minimum length requirement.
// It implements the Validator interface and can be used with struct validation.
//
// The validator checks that:
// - The value is a string
// - The string length is >= the specified minimum
// - Empty values are considered valid (use with Required validator if needed)
type MinLengthValidator struct {
	ValidatorErrorMessage // Embedded struct for custom error messages
}

// Validate checks if the field value meets the minimum length requirement.
//
// Parameters:
//   - field: Metadata about the field being validated (contains tag options)
//   - value: The actual value to validate (expected to be string or string-convertible)
//
// Returns:
//   - error: Returns a ValidationError if validation fails, nil otherwise
//
// Validation Rules:
//  1. Nil/empty values are considered valid (pass validation)
//  2. Non-string values are considered valid (pass validation)
//  3. Fails if string length is less than specified minimum
//
// Example Usage:
//
//	type User struct {
//	    Password string `validate:"minlen=8"`  // Requires minimum 8 characters
//	}
func (v *MinLengthValidator) Validate(field *FieldInfo, value any) error {
	// Skip validation for empty values (combine with Required validator if needed)
	if isEmpty(value) {
		return nil
	}

	// Only validate string types
	str, ok := value.(string)
	if !ok {
		return nil
	}

	// Parse minimum length from validation tag (first tag option)
	minLen, err := parseInt(field.TagOptions[0])
	if err != nil {
		// Return parsing error if min length isn't a valid number
		return err
	}

	// Perform length validation
	if len(str) < minLen {
		// Create formatted validation error
		return v.NewValidationError(
			field.Field.Name,
			"must be at least %d characters",
			minLen,
		).WithMessage()
	}

	return nil
}
