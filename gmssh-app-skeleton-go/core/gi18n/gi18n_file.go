package gi18n

import (
	"context"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	array "github.com/DemonZack/simplejrpc-go/container/garray"
	"github.com/DemonZack/simplejrpc-go/os/gpath"
)

// Package-level constants and variables
var (
	// DefaultLanguage sets the fallback language if detection fails
	defaultLanguage = English

	// Supported file formats for internationalization
	supportedFileTypes = []FileType{IniFile}

	// Default folder name containing language files
	defaultLanguageFolder = "i18n"

	// Directories to search for language files
	resourceTryFolders = []string{
		"i18n/", // Primary location for i18n files
	}

	// Custom path for language files from environment variable
	specialConfigPath = ""
)

// fileDesc describes a language file with its path and type
type fileDesc struct {
	path     string   // Full path to the language file
	fileType FileType // Format of the file (INI, JSON, etc.)
}

// I18nFileAdapter handles loading and managing internationalization files
type I18nFileAdapter struct {
	fileCheck   bool                    // Flag indicating if files were found
	localePath  string                  // Base path for locale files
	searchPaths *array.AnyArray[string] // All discovered file paths
	rootPath    string                  // Application root path
	localesPath map[Language]fileDesc   // Mapping of languages to their files
}

// NewI18nFileAdapter creates a new internationalization file adapter
// Optional paths parameter can override default search locations
func NewI18nFileAdapter(paths ...string) *I18nFileAdapter {
	// Check for custom path in environment variable
	configPath := os.Getenv("I18N_PATH")
	if configPath != "" {
		specialConfigPath = configPath
	}

	var localePath string

	// Use provided path or default location
	if len(paths) > 0 {
		localePath = paths[0]
		specialConfigPath = localePath
	} else {
		localePath = defaultLanguageFolder
	}

	// Initialize adapter
	i18nAdapter := &I18nFileAdapter{
		localePath:  localePath,
		rootPath:    gpath.GmCfgPath,
		searchPaths: array.NewDefaultArray[string](),
		localesPath: make(map[Language]fileDesc),
	}

	// Scan filesystem for language files
	i18nAdapter.scanPath()
	return i18nAdapter
}

// extractBaseName parses a file path to extract language and filename
// Returns the detected language and original path
func (i *I18nFileAdapter) extractBaseName(path string) (lang Language, fName string) {
	fs, err := os.Stat(path)
	if err != nil {
		return defaultLanguage, path
	}

	// Split filename by dots to extract language code
	parts := strings.Split(fs.Name(), ".")
	return NewLanguage(strings.Join(parts[:len(parts)-1], "")), path
}

// walk recursively scans a directory for supported language files
func (i *I18nFileAdapter) walk(path string) {
	filepath.Walk(path, func(childPath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if file matches supported types
		for _, fileSuffix := range supportedFileTypes {
			if strings.HasSuffix(childPath, fileSuffix.String()) {
				// Add to search paths
				i.searchPaths.Append(childPath)

				// Extract language and store file info
				key, val := i.extractBaseName(childPath)
				i.localesPath[key] = fileDesc{
					fileType: fileSuffix,
					path:     val,
				}
				i.fileCheck = true
			}
		}
		return nil
	})
}

// scanPath searches for language files in predefined locations
func (i *I18nFileAdapter) scanPath() {
	// Check special config path first
	if specialConfigPath != "" {
		i.searchPaths.Append(specialConfigPath)
		_, err := os.Stat(specialConfigPath)
		if err == nil {
			i.walk(specialConfigPath)
			return
		}
	}

	// Check standard locations
	folders := make([]string, 0)
	folders = append(folders, resourceTryFolders...)
	absPath, err := filepath.Abs(i.rootPath)
	if err != nil {
		return
	}

	for _, path := range folders {
		filePath := filepath.Join(absPath, path)
		i.walk(filePath)
	}
}

// doGetFilePath returns the last found file path if available
func (i *I18nFileAdapter) doGetFilePath() (filePath string) {
	if i.fileCheck {
		filePath = i.searchPaths.Last()
	}
	return
}

// Available checks if any language files were found
func (i *I18nFileAdapter) Available(ctx context.Context, fileName ...string) bool {
	return i.fileCheck
}

// getContent reads the content of the language file
func (i *I18nFileAdapter) getContent() ([]byte, error) {
	filePath := i.doGetFilePath()

	bData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return bData, nil
}

// Data loads and parses the language file into a string map
func (a *I18nFileAdapter) Data(ctx context.Context) (data map[string]string, err error) {
	bData, err := a.getContent()
	if err != nil {
		return
	}

	data = make(map[string]string)
	err = json.Unmarshal(bData, &data)
	return
}
