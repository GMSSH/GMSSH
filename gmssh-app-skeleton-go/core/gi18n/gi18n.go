package gi18n

// SetPath configures the directory path where internationalization files are stored.
// This initializes the i18n system by loading translation files from the specified path.
// The path should point to a directory containing language files (e.g., en.json, zh-CN.json).
// Note: This must be called before any translation operations.
func SetPath(path string) {
	Instance().SetPath(path)
}

// SetLanguage changes the active translation language.
// The language parameter should be a standard language code (e.g., "en", "zh-CN").
// If the specified language is not available, the system may fall back to default language.
// Note: The language files must be loaded via SetPath before calling this.
func SetLanguage(language string) {
	Instance().SetLanguage(language)
}

// T is a shorthand alias for Translate, providing convenient access to translations.
// Returns the translated string for the given content in the currently active language.
// If no translation is found, returns the original content string.
// Example: T("welcome_message") returns "Welcome" when English is active
func T(content string) string {
	return Instance().T(content)
}

// Tf is a shorthand alias for TranslateFormat, providing formatted translations.
// Combines translation with string formatting (like fmt.Sprintf).
// Returns the translated and formatted string in the active language.
// Example: Tf("welcome_user", "John") could return "Welcome, John"
func Tf(format string, values ...any) string {
	return Instance().TranslateFormat(format, values...)
}

// TranslateFormat retrieves a translated string and formats it with provided values.
// First looks up the translation for the format string in the active language,
// then applies standard fmt.Sprintf formatting with the provided values.
// Returns the formatted string or original format if translation not found.
// Example: TranslateFormat("items_remaining", 5) could return "5 items remaining"
func TranslateFormat(format string, values ...any) string {
	return Instance().TranslateFormat(format, values...)
}

// Translate retrieves the localized version of the given content string.
// Looks up the translation in the currently active language's dictionary.
// Returns the original content string if no translation is available.
// This is the primary method for simple string translations.
func Translate(content string) string {
	return Instance().Translate(content)
}
