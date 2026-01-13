package gvalid

import "strings"

// Visit implements the Visitor interface to validate a struct field based on its tags.
// It processes validation rules specified in the field's tag and applies corresponding validators.
//
// The method handles:
// - Tag parsing and rule extraction
// - Validator lookup and invocation
// - Custom error message handling
// - Options passing to validators
//
// Tag Format:
//
//	`tagName:"rule1:opt1,opt2|rule2#custom error|rule3:opt"`
//
// Where:
// - ruleN: Validation rule name (e.g., "required", "minlen")
// - #custom error: Optional custom error message
// - :opt1,opt2: Rule-specific options
func (v *ValidatorVisitor) Visit(field *FieldInfo) error {
	// Skip if no tag name specified
	if field.TagName == "" {
		return nil
	}

	// Get the validation tag content
	tag := field.Field.Tag.Get(field.TagName)
	if tag == "" {
		return nil
	}

	// Split into individual rules
	rules := strings.Split(tag, "|")
	for _, rule := range rules {
		var (
			ruleName    string   // The validation rule name (e.g., "required")
			customLabel string   // Custom error message (after #)
			options     []string // Rule options (after :)
		)

		// Extract custom error message if present
		if strings.Contains(rule, "#") {
			parts := strings.SplitN(rule, "#", 2)
			rule, customLabel = parts[0], parts[1]
		}

		// Split rule name from its options
		parts := strings.SplitN(rule, ":", 2)
		ruleName = parts[0]
		if len(parts) > 1 {
			options = strings.Split(parts[1], ",")
		}

		// Find and execute the validator
		if validator, exists := v.validators[ruleName]; exists {
			// Set custom message if provided
			if customLabel != "" {
				validator.Message(customLabel)
			}

			// Pass options to validator
			field.TagOptions = options

			// Execute validation
			if err := validator.Validate(field, field.Value); err != nil {
				return err
			}
		}
	}

	return nil
}
