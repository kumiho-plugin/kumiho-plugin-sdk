// Package service defines the HTTP contract for ServiceRuntime plugins.
//
// ServiceRuntime plugins (Docker containers or any HTTP server) must expose
// the endpoints listed below. All bodies are JSON (application/json).
// The plugin process must be ready within [plugin.StartupTimeout] after start.
//
// # Endpoint table
//
//	POST /search    Search for metadata candidates
//	POST /fetch     Fetch full metadata for a selected candidate
//	GET  /health    Return operational status
//	GET  /manifest  Return the static plugin declaration
//
// # HTTP status rules
//
//	200 OK                 — always used for search/fetch/health/manifest responses,
//	                         even when the body contains an error field
//	400 Bad Request        — malformed or missing JSON request body
//	405 Method Not Allowed — wrong HTTP method for the endpoint
//	500 Internal Error     — unexpected plugin-side panic or crash
//
// # Request/response types
//
//	POST /search   body: types.SearchRequest  → types.SearchResponse
//	POST /fetch    body: types.FetchRequest   → types.FetchResponse
//	GET  /health   body: —                   → healthcheck.Response
//	GET  /manifest body: —                   → manifest.Manifest
package service

// HTTP path constants for ServiceRuntime endpoints.
const (
	PathSearch   = "/search"
	PathFetch    = "/fetch"
	PathHealth   = "/health"
	PathManifest = "/manifest"
)

// HTTP header and content type constants.
const (
	HeaderContentType    = "Content-Type"
	HeaderAccept         = "Accept"
	ContentTypeJSON      = "application/json"
	ContentTypeJSONUTF8  = "application/json; charset=utf-8"
)

// HTTP status codes used by the contract.
const (
	StatusOK                 = 200
	StatusBadRequest         = 400
	StatusMethodNotAllowed   = 405
	StatusInternalError      = 500
)
