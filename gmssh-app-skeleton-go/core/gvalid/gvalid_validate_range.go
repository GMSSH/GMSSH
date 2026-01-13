package gvalid

import "fmt"

// RangeValidator validates that a numeric value falls within a specified range.
// It implements the Validator interface for validating minimum and maximum bounds.
//
// The validator checks that:
// - The value is numeric (or convertible to float64)
// - The value is between the specified min and max (inclusive)
// - Non-numeric values are considered valid (should be combined with type validators)
type RangeValidator struct {
	ValidatorErrorMessage // Embedded for custom error message support
}

// Validate checks if the field value falls within the specified numeric range.
//
// Parameters:
//   - field: Contains field metadata including tag options with min/max values
//   - value: The value to validate (must be numeric/convertible to float64)
//
// Returns:
//   - error: ValidationError if value is outside range, nil if valid
//
// Validation Rules:
//  1. Non-numeric values pass validation (combine with numeric type validator)
//  2. Values within [min,max] range pass validation
//  3. Values outside range return formatted ValidationError
//
// Tag Format:
//
//	`validate:"range=MIN,MAX"`  // Where MIN and MAX are float values
//
// Example Usage:
//
//	type Product struct {
//	    Price float64 `validate:"range=0.99,999.99"` // Must be between $0.99-$999.99
//	}
func (v *RangeValidator) Validate(field *FieldInfo, value any) error {
	// Convert value to float64 if possible
	num, ok := toFloat(value)
	if !ok {
		// Skip validation for non-numeric values
		return nil
	}

	// Require both min and max values in tag
	if len(field.TagOptions) < 2 {
		return nil
	}

	// Parse minimum bound from first tag option
	min, err := parseFloat(field.TagOptions[0])
	if err != nil {
		return fmt.Errorf("invalid minimum range value: %w", err)
	}

	// Parse maximum bound from second tag option
	max, err := parseFloat(field.TagOptions[1])
	if err != nil {
		return fmt.Errorf("invalid maximum range value: %w", err)
	}

	// Validate value is within bounds
	if num < min || num > max {
		return v.NewValidationError(
			field.Field.Name,
			"must be between %v and %v", // Use %v for automatic float formatting
			min,
			max,
		).WithMessage()
	}

	return nil
}
