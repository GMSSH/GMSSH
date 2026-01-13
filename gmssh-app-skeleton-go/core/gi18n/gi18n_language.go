package gi18n

import "strings"

// Language represents a supported language in the internationalization system
type Language int

// Enumeration of supported languages
const (
	English            Language = iota // English language (default)
	Chinese                            // Generic Chinese (fallback for Chinese variants)
	SimplifiedChinese                  // Chinese simplified characters (zh-CN)
	TraditionalChinese                 // Chinese traditional characters (zh-TW)
)

// NewLanguage creates a Language from a string identifier
// Supported inputs:
//   - "en" for English
//   - "zh" for Chinese (generic)
//   - "zh-CN" for Simplified Chinese
//   - "zh-TW" for Traditional Chinese
//
// Returns English as default if unknown language code is provided
func NewLanguage(lang string) Language {
	switch strings.ToLower(lang) { // Case-insensitive comparison
	case "en":
		return English
	case "zh":
		return Chinese
	case "zh-cn":
		return SimplifiedChinese
	case "zh-tw":
		return TraditionalChinese
	default:
		return English // Default fallback
	}
}

// String returns the standardized language code string representation
// Returns empty string for unknown language values
func (l Language) String() string {
	switch l {
	case English:
		return "en" // ISO 639-1 code for English
	case Chinese:
		return "zh" // ISO 639-1 code for Chinese
	case SimplifiedChinese:
		return "zh-CN" // ISO 639-1 + country code
	case TraditionalChinese:
		return "zh-TW" // ISO 639-1 + country code
	default:
		return "" // Unknown language
	}
}
