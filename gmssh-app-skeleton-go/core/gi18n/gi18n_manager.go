package gi18n

import "sync"

// Package-level global variables for managing i18n instances
var (
	i18nAdapter *I18nFileAdapter // Global adapter for i18n file operations
	i18nManager *I18nManager     // Global manager for i18n functionality
)

// I18nManager provides thread-safe internationalization management
type I18nManager struct {
	mu          sync.RWMutex // Mutex for concurrent access protection
	i18nMessage *I18nMessage // Underlying message storage and translation
}

// NewI18nManager creates a new I18nManager instance
// Initializes with a new I18nMessage container
func NewI18nManager() *I18nManager {
	i18nMessage := NewI18nMessage()
	return &I18nManager{
		i18nMessage: i18nMessage,
	}
}

// SetPath configures the i18n system with a new file path
// Loads English translations by default from the specified path
// Panics if unable to load the translation file
func (i *I18nManager) SetPath(path string) {
	i.mu.Lock()
	defer i.mu.Unlock()

	// Initialize adapter with new path
	i18nAdapter = NewI18nFileAdapter(path)

	// Load English translations by default
	localePath := i18nAdapter.localesPath[English]
	parser := CreateI18nParser(localePath.fileType, localePath.path)
	content, err := parser.GetContent()
	if err != nil {
		panic(err)
	}

	// Store loaded translations
	i.i18nMessage.SetLanguageContent(English.String(), content)
}

// SetLanguage switches the active translation language
// Panics if unable to load the specified language file
func (i *I18nManager) SetLanguage(lang string) {
	// Get file path for requested language
	localePath := i18nAdapter.localesPath[NewLanguage(lang)]

	// Create appropriate parser and load content
	parser := CreateI18nParser(localePath.fileType, localePath.path)
	content, err := parser.GetContent()
	if err != nil {
		panic(err)
	}

	// Store loaded translations
	i.i18nMessage.SetLanguageContent(lang, content)
}

// Translate returns the translation for the given key
// Returns the key itself if translation not found
func (i *I18nManager) Translate(key string) string {
	return i.i18nMessage.Translate(key)
}

// T is an alias for Translate - provides shorthand syntax
func (i *I18nManager) T(key string) string {
	return i.Translate(key)
}

// TranslateFormat returns a formatted translation with substituted values
// Supports printf-style formatting in translation strings
func (i *I18nManager) TranslateFormat(key string, values ...any) string {
	return i.i18nMessage.TranslateFormat(key, values...)
}

// Tf is an alias for TranslateFormat - provides shorthand syntax
func (i *I18nManager) Tf(key string, values ...any) string {
	return i.TranslateFormat(key, values...)
}
