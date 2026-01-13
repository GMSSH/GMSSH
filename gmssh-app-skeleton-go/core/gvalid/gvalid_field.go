package gvalid

import (
	"reflect"
)

// Validator is the core validation interface that defines how validation rules
// should be implemented. Any custom validator must implement this interface.
type Validator interface {
	// Validate performs the actual validation logic on a field
	// Parameters:
	//   - field: Metadata about the field being validated
	//   - value: The actual value to validate
	// Returns:
	//   - error: Validation error if validation fails, nil otherwise
	Validate(field *FieldInfo, value any) error

	// Message sets a custom error message for the validator
	// Parameters:
	//   - message: The custom error message template
	Message(message string)

	// NewValidationError creates a standardized validation error
	// Parameters:
	//   - field: The name of the field being validated
	//   - format: Error message format string
	//   - args: Arguments for the format string
	// Returns:
	//   - *ValidationError: A new validation error instance
	NewValidationError(field, format string, args ...any) *ValidationError
}

// FieldInfo contains metadata about a struct field being validated.
// This information is used by validators to make context-aware decisions.
type FieldInfo struct {
	Kind       reflect.Kind        // The reflect.Kind of the field
	Field      reflect.StructField // The complete struct field information
	Value      any                 // The actual field value to validate
	Parent     reflect.Value       // The parent struct containing this field
	TagName    string              // The validation tag name (e.g., "required")
	TagOptions []string            // Any options specified in the validation tag
}

// Visitor defines the interface for visiting and validating struct fields.
// This follows the visitor design pattern for traversing and validating structs.
type Visitor interface {
	// Visit processes a single field for validation
	// Parameters:
	//   - field: The field to validate
	// Returns:
	//   - error: Validation error if validation fails, nil otherwise
	Visit(field *FieldInfo) error

	// RegisterValidator adds a new validator for a specific tag
	// Parameters:
	//   - tag: The struct tag that triggers this validator
	//   - validator: The validator implementation
	RegisterValidator(tag string, validator Validator)
}

// ValidatorVisitor implements the Visitor interface and manages a collection
// of validators that can be applied to struct fields.
type ValidatorVisitor struct {
	validators map[string]Validator // Maps tag names to validator implementations
}

// NewValidatorVisitor creates a new ValidatorVisitor instance with an empty
// validator registry. Validators must be registered before use.
func NewValidatorVisitor() *ValidatorVisitor {
	return &ValidatorVisitor{
		validators: make(map[string]Validator),
	}
}

// RegisterValidator adds a validator to the visitor's registry.
// This validator will be used when encountering fields with the specified tag.
// Parameters:
//   - tag: The struct tag name (e.g., "required", "email")
//   - validator: The validator implementation for this tag
func (v *ValidatorVisitor) RegisterValidator(tag string, validator Validator) {
	v.validators[tag] = validator
}
