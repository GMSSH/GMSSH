package mapping

import (
	"encoding/json"
	"reflect"
)

// StructToMap converts a struct to a map[string]any using JSON serialization.
// This approach ensures proper handling of JSON tags and nested structures.
//
// Parameters:
//   - obj: the struct to be converted (can be any type, but should be a struct for meaningful results)
//
// Returns:
//   - map[string]any: the resulting map where keys are struct field names and values are field values
//   - error: if JSON marshaling/unmarshaling fails
func StructToMap(obj any) (map[string]any, error) {
	// First marshal the struct to JSON bytes
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	// Then unmarshal the JSON into a map
	result := make(map[string]any)
	err = json.Unmarshal(data, &result)
	return result, err
}

// MapToStruct populates a struct with values from a map by matching field names.
// Note: This is a simple implementation that doesn't handle type conversion or nested structures.
//
// Parameters:
//   - m: the source map containing field values (keys should match struct field names)
//   - s: pointer to the target struct to be populated
//
// Returns:
//   - error: always returns nil in current implementation (should be enhanced for error handling)
func MapToStruct(m map[string]any, s any) error {
	// Get the reflect.Value of the struct (dereferencing the pointer)
	structValue := reflect.ValueOf(s).Elem()
	// Get the reflect.Type of the struct
	structType := structValue.Type()

	// Iterate through all fields of the struct
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)    // Get the field's reflect.Value
		fieldType := structType.Field(i) // Get the field's reflect.StructField
		fieldName := fieldType.Name      // Get the field name

		// If the map contains a value for this field name
		if value, ok := m[fieldName]; ok {
			// Set the field's value using reflection
			field.Set(reflect.ValueOf(value))
		}
	}

	return nil
}
