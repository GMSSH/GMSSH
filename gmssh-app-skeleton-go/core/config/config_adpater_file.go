package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	array "github.com/DemonZack/simplejrpc-go/container/garray"
	"github.com/DemonZack/simplejrpc-go/os/gpath"
)

// FileAdapter provides read-only access to configuration files.
// It supports searching for config files in multiple directories
// and reading JSON configuration data.
// Note: This implementation does not support modification of config files.
type FileAdapter struct {
	defaultFileNameOrPath string                  // Default config file name or path
	searchPaths           *array.AnyArray[string] // List of paths searched for config files
	fileCheck             bool                    // Flag indicating if config file was found
	rootPath              string                  // Base directory for config file search
}

var (
	// supportedFileTypes lists all supported configuration file extensions
	supportedFileTypes = []string{"json"}

	// resourceTryFolders contains directories to search for config files (relative paths)
	resourceTryFolders = []string{
		"", "/", "config/", "config", "/config", "/config/", "/config/config",
	}

	// localSystemTryFolders contains higher priority directories for config files
	localSystemTryFolders = []string{"manifest/config"}

	// specialConfigPath stores custom config path from environment variable
	specialConfigPath = ""
)

// NewAdapterFile creates a new FileAdapter instance.
// It accepts an optional config file name/path parameter.
// If no parameter is provided, it uses DefaultConfigFileName.
// It initializes the config file search paths and checks availability.
func NewAdapterFile(fileNameOrPath ...string) (*FileAdapter, error) {
	// Check for custom config path in environment variable
	configPath := os.Getenv("CONFIG_PATH")
	if configPath != "" {
		specialConfigPath = configPath
	}

	var (
		err                error
		usedFileNameOrPath = DefaultConfigFileName
		rootPath           = gpath.GmCfgPath
	)

	// Use provided filename/path if available
	if len(fileNameOrPath) > 0 {
		usedFileNameOrPath = fileNameOrPath[0]
	} else {
		// TODO: Read custom config file from environment variables
		// return nil, errors.New("config file not found")
	}

	// Initialize FileAdapter instance
	fileAdapter := &FileAdapter{
		defaultFileNameOrPath: usedFileNameOrPath,
		searchPaths:           array.NewDefaultArray[string](),
		rootPath:              rootPath,
	}

	// Scan filesystem for config file
	fileAdapter.scanPath()

	return fileAdapter, err
}

// doGetFilePath returns the path of the found config file.
// Only returns a path if fileCheck is true (file was found).
func (a *FileAdapter) doGetFilePath() (filePath string) {
	if a.fileCheck {
		filePath = a.searchPaths.Last()
	}
	return
}

// scanPath searches for configuration files in predefined locations.
// It first checks the specialConfigPath if set, then falls back to
// standard search paths. Sets fileCheck flag if config file is found.
func (a *FileAdapter) scanPath() {
	// First check special config path if specified
	if specialConfigPath != "" {
		a.searchPaths.Append(specialConfigPath)
		_, err := os.Stat(specialConfigPath)
		if err == nil {
			// File exists, set flag and return
			a.fileCheck = true
			return
		}
	}

	// If special path not found, check standard locations
	folders := make([]string, 0, len(localSystemTryFolders)+len(resourceTryFolders))
	folders = append(folders, localSystemTryFolders...)
	folders = append(folders, resourceTryFolders...)

	// Get absolute path of root directory
	absPath, err := filepath.Abs(a.rootPath)
	if err != nil {
		return
	}

	// Check each folder for supported config file types
	for _, path := range folders {
		filePath := filepath.Join(absPath, path)
		for _, fileType := range supportedFileTypes {
			file := fmt.Sprintf("%s.%s", filePath, fileType)
			a.searchPaths.Append(file)
			_, err := os.Stat(file)
			if err != nil {
				continue
			}
			a.fileCheck = true
			return
		}
	}
}

// getContent reads and returns the content of the found config file.
// Returns error if file cannot be read.
func (a *FileAdapter) getContent() ([]byte, error) {
	filePath := a.doGetFilePath()

	bData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return bData, nil
}

// Available checks if a valid config file was found during initialization.
// The context and fileName parameters are currently unused.
func (a *FileAdapter) Available(ctx context.Context, fileName ...string) bool {
	return a.fileCheck
}

// Data reads and parses the config file into a map[string]any.
// Returns error if file cannot be read or parsed as JSON.
func (a *FileAdapter) Data(ctx context.Context) (data map[string]any, err error) {
	bData, err := a.getContent()
	if err != nil {
		return
	}

	data = make(map[string]any)
	err = json.Unmarshal(bData, &data)
	return
}
