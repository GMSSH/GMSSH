package gvalid

import "reflect"

// StructWalker is responsible for traversing struct fields and applying validation
// using the visitor pattern. It handles both top-level and nested struct fields.
type StructWalker struct {
	visitor Visitor // The visitor implementation that processes each field
	tagName string  // The struct tag name used to identify validation rules (e.g., "validate")
}

// NewStructWalker creates a new StructWalker instance with the specified visitor
// and tag name. The tag name identifies which struct tags contain validation rules.
//
// Parameters:
//   - visitor: The visitor implementation that will process each field
//   - tagName: The struct tag name containing validation rules (e.g., "validate")
func NewStructWalker(visitor Visitor, tagName string) *StructWalker {
	return &StructWalker{
		visitor: visitor,
		tagName: tagName,
	}
}

// RegisterValidator adds a new validator for a specific tag. This allows
// dynamically extending the validation rules available to the walker.
//
// Parameters:
//   - tagName: The tag that will trigger this validator (e.g., "required")
//   - validator: The validator implementation for this tag
func (s *StructWalker) RegisterValidator(tagName string, validator Validator) {
	s.visitor.RegisterValidator(tagName, validator)
}

// Walk traverses a struct (or pointer to struct) and applies validation
// to each field. It handles validation errors through the validErrorHandler.
//
// Parameters:
//   - s: The struct (or pointer to struct) to validate
//
// Returns:
//   - error: Validation errors if any, nil otherwise
func (w *StructWalker) Walk(s any) error {
	val := reflect.ValueOf(s)

	// Dereference if pointer
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Only structs can be walked
	if val.Kind() != reflect.Struct {
		return nil
	}

	// Process struct and handle any validation errors
	return validErrorHandler(w.walkStruct(val))
}

// walkStruct recursively processes a struct's fields, including nested structs
func (w *StructWalker) walkStruct(val reflect.Value) error {
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		// Handle embedded/nested structs (anonymous fields)
		if field.Type.Kind() == reflect.Struct && field.Anonymous {
			if err := w.walkStruct(fieldValue); err != nil {
				return err
			}
			continue
		}

		// Create field metadata for validators
		fieldInfo := &FieldInfo{
			Kind:    field.Type.Kind(),      // Field type (string, int, etc.)
			Field:   field,                  // Complete field metadata
			Value:   fieldValue.Interface(), // Actual field value
			Parent:  val,                    // Parent struct value
			TagName: w.tagName,              // Tag name for validation rules
		}

		// Delegate field processing to visitor
		if err := w.visitor.Visit(fieldInfo); err != nil {
			return err
		}
	}

	return nil
}
