package gsock

import "time"

// ServiceMeta contains organizational metadata about a service
// This information helps identify and categorize services in a distributed system
type ServiceMeta struct {
	OrgName string `json:"orgName"` // Organization name that owns the service (e.g. "acme-corp")
	AppName string `json:"appName"` // Application name the service belongs to (e.g. "billing")
	Version string `json:"version"` // Service version in SemVer format (e.g. "1.2.3")
}

// ServiceInfo contains complete service registration information
// Used for service discovery, health checking, and monitoring
type ServiceInfo struct {
	ServerName     string       `json:"name"`           // Unique service identifier (e.g. "billing-service-v1")
	ServerPort     string       `json:"port"`           // Network port the service listens on (e.g. "8080")
	ServerType     string       `json:"type"`           // Protocol type ("http", "grpc", "websocket", etc.)
	HealthPath     string       `json:"healthPath"`     // Endpoint path for health checks (e.g. "/health")
	HealthTimeout  int          `json:"healthTimeout"`  // Health check timeout in seconds (e.g. 5)
	MetaData       *ServiceMeta `json:"metaData"`       // Organizational metadata
	Status         int          `json:"status"`         // Current status (0=down, 1=up, 2=degraded)
	LastActiveTime time.Time    `json:"lastActiveTime"` // Last time service was active
	Pid            int          `json:"pid"`            // Process ID of running service
}

// Service status constants
const (
	ServiceStatusDown     = 0 // Service is not responding
	ServiceStatusUp       = 1 // Service is healthy
	ServiceStatusDegraded = 2 // Service is running but with issues
)

// DefaultHealthTimeout is the recommended default health check timeout
const DefaultHealthTimeout = 5 // seconds
