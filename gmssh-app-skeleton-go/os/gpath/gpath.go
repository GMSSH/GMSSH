package gpath

import "path/filepath"

var (
	// GmCfgPath stores the base configuration directory path
	// Defaults to current directory
	GmCfgPath = filepath.Join(filepath.Dir(""), ".")
	// Alternative: GmCfgPath = filepath.Join(filepath.Dir(getCurrentAbPath()), "..")
)
