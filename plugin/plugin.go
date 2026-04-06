// Package plugin defines the plugin contract that all Kumiho plugins must satisfy,
// regardless of runtime type.
//
// # BinaryRuntime plugins
//
// Implement the [MetadataPlugin] or [TranslationPlugin] interface. The core uses
// HashiCorp go-plugin to manage the child process lifecycle and RPC communication.
//
// # ServiceRuntime plugins (Docker / HTTP)
//
// Metadata plugins expose the following HTTP endpoints:
//
//	POST /search    → accepts types.SearchRequest, returns types.SearchResponse
//	POST /fetch     → accepts types.FetchRequest,  returns types.FetchResponse
//	GET  /health    → returns healthcheck.Response
//	GET  /manifest  → returns manifest.Manifest
//
// Translation plugins expose the following HTTP endpoints:
//
//	POST /translate → accepts types.TranslateRequest, returns types.TranslateResponse
//	POST /detect    → accepts types.DetectRequest,    returns types.DetectResponse
//	GET  /health    → returns healthcheck.Response
//	GET  /manifest  → returns manifest.Manifest
//
// HTTP status codes:
//   - 200 OK for successful responses (even if the response body contains an error field)
//   - 400 Bad Request for malformed JSON input
//   - 405 Method Not Allowed for wrong HTTP method
//   - 500 Internal Server Error for unexpected plugin-side panics
//
// The plugin process must be ready to serve requests within [StartupTimeout].
package plugin

import (
	"context"
	"time"

	"github.com/kumiho-plugin/kumiho-plugin-sdk/healthcheck"
	"github.com/kumiho-plugin/kumiho-plugin-sdk/manifest"
	"github.com/kumiho-plugin/kumiho-plugin-sdk/types"
)

// StartupTimeout is the maximum time the core waits for a plugin to become ready
// after the process or container starts.
const StartupTimeout = 15 * time.Second

// MetadataPlugin is the Go interface implemented by BinaryRuntime metadata plugins.
// ServiceRuntime plugins expose equivalent HTTP endpoints (see package doc).
type MetadataPlugin interface {
	// Search returns metadata candidates for the given request.
	// Plugins should return a non-nil SearchResponse even on error;
	// set SearchResponse.Error instead of returning a Go error when possible.
	Search(ctx context.Context, req *types.SearchRequest) (*types.SearchResponse, error)

	// Fetch returns full metadata for the candidate identified by req.
	// Plugins should return a non-nil FetchResponse even on error.
	Fetch(ctx context.Context, req *types.FetchRequest) (*types.FetchResponse, error)

	// Healthcheck returns the plugin's current operational status.
	Healthcheck(ctx context.Context) (*healthcheck.Response, error)

	// Manifest returns the static plugin declaration.
	Manifest() *manifest.Manifest
}

// TranslationPlugin is the Go interface implemented by BinaryRuntime translation plugins.
// ServiceRuntime plugins expose equivalent HTTP endpoints (see package doc).
type TranslationPlugin interface {
	// Translate translates the given text to the target language.
	// Plugins should return a non-nil TranslateResponse even on error;
	// set TranslateResponse.Error instead of returning a Go error when possible.
	Translate(ctx context.Context, req *types.TranslateRequest) (*types.TranslateResponse, error)

	// Detect detects the language of the given text.
	// Plugins should return a non-nil DetectResponse even on error.
	Detect(ctx context.Context, req *types.DetectRequest) (*types.DetectResponse, error)

	// Healthcheck returns the plugin's current operational status.
	Healthcheck(ctx context.Context) (*healthcheck.Response, error)

	// Manifest returns the static plugin declaration.
	Manifest() *manifest.Manifest
}
