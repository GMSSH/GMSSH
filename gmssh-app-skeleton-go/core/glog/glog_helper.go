package glog

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// parseRotateExpire parses a duration string and converts it to hours
// Supported formats:
//   - "Nd" for days (e.g., "7d" = 168 hours)
//   - "Nh" for hours (e.g., "24h")
//
// Defaults to 24 hours (1 day) for invalid/unrecognized formats
func parseRotateExpire(expire string) int {
	// Handle day format (e.g., "7d")
	if strings.HasSuffix(expire, "d") {
		days := strings.TrimSuffix(expire, "d")
		return 24 * parseInt(days, 1) // Convert days to hours, default 1 day
	}
	// Handle hour format (e.g., "48h")
	if strings.HasSuffix(expire, "h") {
		hours := strings.TrimSuffix(expire, "h")
		return parseInt(hours, 24) // Default 24 hours
	}
	// Default fallback
	return 24 // 1 day in hours
}

// parseInt safely converts a string to integer with fallback default
// Returns defaultValue if parsing fails
func parseInt(s string, defaultValue int) int {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil {
		return defaultValue
	}
	return n
}

// replaceDatePlaceholders replaces date templates in filenames with current date
// Currently supports:
//   - {Y-m-d} → "2006-01-02" format
//
// Returns original string if no placeholders found
func replaceDatePlaceholders(filename string) string {
	now := time.Now()
	// Replace {Y-m-d} with current date
	if strings.Contains(filename, "{Y-m-d}") {
		return strings.ReplaceAll(filename, "{Y-m-d}", now.Format("2006-01-02"))
	}
	return filename
}

// mapToStruct populates a struct from map values using reflection
// Note: This is a simple implementation that doesn't handle:
//   - Nested structures
//   - Type conversion
//   - Complex field types
func mapToStruct(m map[string]any, s any) error {
	structValue := reflect.ValueOf(s).Elem()
	structType := structValue.Type()

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldType := structType.Field(i)
		fieldName := fieldType.Name

		if value, ok := m[fieldName]; ok {
			field.Set(reflect.ValueOf(value))
		}
	}

	return nil
}

// map2Struct converts a map to struct using JSON serialization
// More robust than mapToStruct as it handles:
//   - JSON tags
//   - Type conversions
//   - Nested structures
//
// But has higher overhead due to JSON marshaling
func map2Struct(m map[string]any, s any) error {
	// Convert map to JSON bytes
	jsonData, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal map to JSON: %w", err)
	}

	// Unmarshal JSON into target struct
	err = json.Unmarshal(jsonData, s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON to struct: %w", err)
	}
	return nil
}
