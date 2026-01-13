package gvalid

// validErrorHandler processes validation errors and returns the most appropriate error.
// It handles two main cases:
// 1. ValidationErrors (multiple validation errors) - returns the first error
// 2. Other error types - returns the error as-is
//
// Parameters:
//   - err: The error to process (can be nil, ValidationErrors, or other error type)
//
// Returns:
//   - error: The processed error, or nil if input was nil
//
// Example usage:
//
//	if err := validErrorHandler(validationErr); err != nil {
//	    return err
//	}
func validErrorHandler(err error) error {
	if err != nil {
		// Check if this is a ValidationErrors type (multiple validation errors)
		if verrs, ok := err.(ValidationErrors); ok {
			// Return just the first validation error
			// Note: In some cases you might want to handle all errors
			for _, verr := range verrs {
				return verr
			}
		}
		// For non-validation errors, return as-is
		return err
	}
	// No error case
	return nil
}
