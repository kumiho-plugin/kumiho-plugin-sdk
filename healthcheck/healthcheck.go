package healthcheck

import (
	"context"
	"time"
)

// Status represents the operational status of a plugin.
type Status string

const (
	// StatusOK means the plugin is fully operational.
	StatusOK Status = "ok"

	// StatusDegraded means the plugin is running but with reduced capability
	// (e.g. rate limited, partial API access, elevated error rate).
	StatusDegraded Status = "degraded"

	// StatusDown means the plugin cannot serve requests.
	StatusDown Status = "down"
)

// Response is the payload returned by a plugin's healthcheck endpoint.
type Response struct {
	Status  Status `json:"status"`
	Version string `json:"version"`
	Message string `json:"message,omitempty"`
}

// Checker is implemented by plugins that support healthchecks.
// BinaryRuntime plugins implement this interface directly.
// ServiceRuntime plugins expose GET /health returning a JSON Response.
type Checker interface {
	Healthcheck(ctx context.Context) (*Response, error)
}

// Convention constants.
// These are the expected defaults; the core may override them via configuration.
const (
	// DefaultTimeout is the maximum time the core waits for a healthcheck response.
	DefaultTimeout = 5 * time.Second

	// DefaultInterval is how often the core polls an active plugin's health.
	DefaultInterval = 30 * time.Second

	// DefaultMaxRetries is the number of consecutive healthcheck failures before
	// the core transitions the plugin to the unhealthy state.
	DefaultMaxRetries = 3

	// DefaultRetryBackoff is the initial backoff between retries (exponential).
	DefaultRetryBackoff = 2 * time.Second
)
