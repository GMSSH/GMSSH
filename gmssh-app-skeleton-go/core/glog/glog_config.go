package glog

// LogConfig defines the configuration structure for logging system
type LogConfig struct {
	Path                 string `json:"path"`                 // Directory path for log files
	File                 string `json:"file"`                 // Base filename for logs
	Level                string `json:"level"`                // Logging level (e.g., "debug", "info", "warn", "error")
	Stdout               bool   `json:"stdout"`               // Whether to output logs to stdout
	StStatus             int    `json:"StStatus"`             // Status code for structured logging
	RotateBackupLimit    int    `json:"rotateBackupLimit"`    // Maximum number of rotated log files to keep
	WriterColorEnable    bool   `json:"writerColorEnable"`    // Enable colored output in console
	RotateBackupCompress int    `json:"RotateBackupCompress"` // Compression level for rotated logs (0=disabled)
	RotateExpire         string `json:"rotateExpire"`         // Expiration duration for rotated logs (e.g., "7d")
	Flag                 int    `json:"Flag"`                 // Additional logging flags (bitmask)
}

// LoadConfig creates a LogConfig from a map of configuration data
// Parameters:
//   - cData: Map containing configuration values (typically from JSON)
//
// Returns:
//   - *LogConfig: Pointer to initialized configuration
//   - error: Any error that occurred during conversion
//
// Example:
//
//	config, err := LoadConfig(map[string]any{
//	  "path":  "/var/log",
//	  "level": "info",
//	})
func LoadConfig(cData map[string]any) (*LogConfig, error) {
	var config LogConfig
	err := map2Struct(cData, &config) // Convert map to struct
	if err != nil {
		return nil, err
	}

	return &config, nil
}
