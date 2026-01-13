package gi18n

import (
	"fmt"
	"sync"

	"github.com/DemonZack/simplejrpc-go/container/gmap"
)

// I18nMessage stores and manages internationalized messages
// It provides thread-safe access to translations in multiple languages
type I18nMessage struct {
	language string                       // Currently active language code (e.g., "en", "zh-CN")
	mu       sync.RWMutex                 // Protects access to the message map
	lmu      sync.RWMutex                 // Protects access to the language field
	message  map[Language]*gmap.StrAnyMap // Map of languages to their translations
}

// NewI18nMessage creates a new I18nMessage instance
// Initializes an empty map for storing translations
func NewI18nMessage() *I18nMessage {
	return &I18nMessage{
		message: make(map[Language]*gmap.StrAnyMap),
	}
}

// SetLanguageContent updates translations for a specific language
// Sets both the language content and switches the active language
// Uses mutex to ensure thread safety
func (i *I18nMessage) SetLanguageContent(lang string, val *gmap.StrAnyMap) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.SetLanguage(lang)
	i.message[NewLanguage(lang)] = val
}

// SetLanguage changes the currently active language
// Uses a separate mutex to prevent contention with translation lookups
func (i *I18nMessage) SetLanguage(lang string) {
	i.lmu.Lock()
	defer i.lmu.Unlock()
	i.language = lang
}

// GetLanguage returns the currently active language code
// Defaults to English if no language is set
func (i *I18nMessage) GetLanguage() string {
	i.lmu.RLock()
	defer i.lmu.RUnlock()
	if i.language == "" {
		return English.String() // Default to English
	}
	return i.language
}

// Translate retrieves a localized message for the given key
// Returns empty string if language not found, or the key itself if translation missing
// Uses read lock for concurrent access
func (i *I18nMessage) Translate(key string) string {
	i.mu.RLock()
	defer i.mu.RUnlock()

	// Get translations for current language
	val, ok := i.message[NewLanguage(i.GetLanguage())]
	if !ok {
		return "" // Language not loaded
	}

	// Lookup translation
	result := val.GetString(key)
	if result == "" {
		return key // Fallback to key if translation missing
	}
	return result
}

// T is a shorthand alias for Translate
func (i *I18nMessage) T(key string) string {
	return i.Translate(key)
}

// TranslateFormat retrieves and formats a localized message
// Supports printf-style formatting (e.g., %s, %d)
func (i *I18nMessage) TranslateFormat(key string, values ...any) string {
	content := i.Translate(key)
	return fmt.Sprintf(content, values...)
}

// Tf is a shorthand alias for TranslateFormat
func (i *I18nMessage) Tf(key string, values ...any) string {
	return i.TranslateFormat(key, values...)
}
