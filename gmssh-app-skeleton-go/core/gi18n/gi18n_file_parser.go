package gi18n

import (
	"github.com/DemonZack/simplejrpc-go/container/gmap"
	"gopkg.in/ini.v1"
)

// IParser defines the interface for internationalization file parsers
// All parsers must implement GetContent() to return parsed data
type IParser interface {
	// GetContent returns parsed content as a StrAnyMap (string to any value mapping)
	// Returns error if parsing fails
	GetContent() (*gmap.StrAnyMap, error)
}

// FileType represents supported internationalization file formats
type FileType int

const (
	IniFile  FileType = iota // INI configuration file format
	JsonFile                 // JSON file format
	TomlFile                 // TOML file format
)

// String returns the string representation of the FileType
func (f FileType) String() string {
	switch f {
	case IniFile:
		return "ini"
	case JsonFile:
		return "json"
	case TomlFile:
		return "toml"
	}
	return "" // Default empty string for unknown types
}

// I18nIniParserAdapter implements IParser for INI file format
type I18nIniParserAdapter struct {
	file string // Path to the INI file
}

// NewI18nIniParserAdapter creates a new INI parser adapter instance
func NewI18nIniParserAdapter(file string) *I18nIniParserAdapter {
	return &I18nIniParserAdapter{file: file}
}

// getDefaultSection extracts the DEFAULT section from INI file
// Returns content as StrAnyMap where keys are INI keys and values are string values
func (f *I18nIniParserAdapter) getDefaultSection(cfg *ini.File) (*gmap.StrAnyMap, error) {
	defaultSection := "DEFAULT"
	sections, err := cfg.SectionsByName(defaultSection)
	if err != nil {
		return nil, err
	}

	sectionMap := &gmap.StrAnyMap{}
	for _, section := range sections {
		for _, val := range section.Keys() {
			sectionMap.Set(val.Name(), val.Value())
		}
	}

	return sectionMap, nil
}

// GetContent implements IParser interface for INI files
// Loads the INI file and returns parsed DEFAULT section content
func (f *I18nIniParserAdapter) GetContent() (out *gmap.StrAnyMap, err error) {
	cfg, err1 := ini.Load(f.file)
	if err1 != nil {
		err = err1
		return
	}
	return f.getDefaultSection(cfg)
}

// I18nJsonParserAdapter implements IParser for JSON file format
type I18nJsonParserAdapter struct {
	file string // Path to the JSON file
}

// NewI18nJsonParserAdapter creates a new JSON parser adapter instance
func NewI18nJsonParserAdapter(file string) *I18nJsonParserAdapter {
	return &I18nJsonParserAdapter{file: file}
}

// GetContent implements IParser interface for JSON files
// Currently just a stub implementation
func (f *I18nJsonParserAdapter) GetContent() (out *gmap.StrAnyMap, err error) {
	return
}

// I18nTomlParserAdapter implements IParser for TOML file format
type I18nTomlParserAdapter struct {
	file string // Path to the TOML file
}

// NewTomlJsonParserAdapter creates a new TOML parser adapter instance
// Note: Function name seems to have a typo (Json vs Toml)
func NewTomlJsonParserAdapter(file string) *I18nTomlParserAdapter {
	return &I18nTomlParserAdapter{file: file}
}

// GetContent implements IParser interface for TOML files
// Currently just a stub implementation
func (f *I18nTomlParserAdapter) GetContent() (out *gmap.StrAnyMap, err error) {
	return
}

// CreateI18nParser is a factory function that creates appropriate parser based on file type
// Defaults to INI parser if file type is unknown
func CreateI18nParser(fileType FileType, path string) IParser {
	switch fileType {
	case IniFile:
		return NewI18nIniParserAdapter(path)
	case JsonFile:
		return NewI18nJsonParserAdapter(path)
	case TomlFile:
		return NewTomlJsonParserAdapter(path)
	}
	return NewI18nIniParserAdapter(path) // Default to INI parser
}
