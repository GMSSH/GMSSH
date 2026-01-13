package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// getCurrentAbPath returns the absolute path of the current working directory.
// It handles different execution scenarios:
// - Regular compiled executable
// - 'go run' execution
// - Running from temporary directories
func getCurrentAbPath() string {
	// First try getting path from executable location
	dir := getCurrentAbPathByExecutable()

	// If we're in a temp directory (like with 'go run'), fall back to caller-based path
	if strings.Contains(dir, getTmpDir()) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// getTmpDir returns the system's temporary directory path.
// It checks both TEMP and TMP environment variables for compatibility.
// Returns the resolved path after evaluating any symbolic links.
func getTmpDir() string {
	// Try TEMP first, fall back to TMP if TEMP not set
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}

	// Resolve any symbolic links in the path
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// getCurrentAbPathByExecutable returns the absolute path of the directory
// containing the current executable. This works for compiled binaries.
// Panics if unable to determine executable path.
func getCurrentAbPathByExecutable() string {
	// Get path of current executable
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err) // Fatal if we can't determine executable path
	}

	// Get directory of executable and resolve any symlinks
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// getCurrentAbPathByCaller returns the absolute path of the current file
// using runtime caller information. This works when running with 'go run'.
// Returns empty string if unable to determine path.
func getCurrentAbPathByCaller() string {
	var abPath string

	// Get caller information (skip=0 means current function)
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		// Extract directory from full filename
		abPath = path.Dir(filename)
	}
	return abPath
}
