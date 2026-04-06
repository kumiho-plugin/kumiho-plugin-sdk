package errors

import "fmt"

// ErrorCode represents a standardized plugin error code.
type ErrorCode string

const (
	// ── General ──────────────────────────────────────────────────────────────

	// ErrCodeUnknown is an unclassified error.
	ErrCodeUnknown ErrorCode = "unknown"

	// ErrCodeUnsupported means the plugin does not support the requested operation.
	ErrCodeUnsupported ErrorCode = "unsupported"

	// ── Provider / API ───────────────────────────────────────────────────────

	// ErrCodeNotFound means the requested resource was not found at the provider.
	ErrCodeNotFound ErrorCode = "not_found"

	// ErrCodeRateLimited means the provider API rate limit was exceeded.
	ErrCodeRateLimited ErrorCode = "rate_limited"

	// ErrCodeQuotaExceeded means the provider account quota has been exhausted.
	ErrCodeQuotaExceeded ErrorCode = "quota_exceeded"

	// ErrCodeUnauthorized means the API key is missing, invalid, or expired.
	ErrCodeUnauthorized ErrorCode = "unauthorized"

	// ErrCodeTimeout means the provider API did not respond within the allowed time.
	ErrCodeTimeout ErrorCode = "timeout"

	// ErrCodeInvalidRequest means the request parameters were malformed or incomplete.
	ErrCodeInvalidRequest ErrorCode = "invalid_request"

	// ErrCodeProviderError means the provider API returned an unexpected error.
	ErrCodeProviderError ErrorCode = "provider_error"

	// ── Install / Verification ───────────────────────────────────────────────

	// ErrCodeInstallFailed means the plugin artifact could not be installed.
	ErrCodeInstallFailed ErrorCode = "install_failed"

	// ErrCodeChecksumMismatch means the downloaded artifact checksum did not match.
	ErrCodeChecksumMismatch ErrorCode = "checksum_mismatch"

	// ErrCodeSignatureInvalid means the artifact signature verification failed.
	ErrCodeSignatureInvalid ErrorCode = "signature_invalid"

	// ErrCodeArtifactNotFound means no artifact was found for the current platform.
	ErrCodeArtifactNotFound ErrorCode = "artifact_not_found"

	// ── Compatibility ────────────────────────────────────────────────────────

	// ErrCodeIncompatibleVersion means the plugin is not compatible with the running core version.
	ErrCodeIncompatibleVersion ErrorCode = "incompatible_version"

	// ErrCodeUnsupportedPlatform means the plugin does not support the current OS/architecture.
	ErrCodeUnsupportedPlatform ErrorCode = "unsupported_platform"

	// ErrCodeSDKVersionMismatch means the plugin was built against an incompatible SDK version.
	ErrCodeSDKVersionMismatch ErrorCode = "sdk_version_mismatch"

	// ── Config ───────────────────────────────────────────────────────────────

	// ErrCodeConfigInvalid means the plugin configuration failed validation.
	ErrCodeConfigInvalid ErrorCode = "config_invalid"

	// ErrCodeConfigMissingRequired means a required config field is not set.
	ErrCodeConfigMissingRequired ErrorCode = "config_missing_required"

	// ── Migration ────────────────────────────────────────────────────────────

	// ErrCodeMigrationFailed means the config schema migration could not be applied.
	ErrCodeMigrationFailed ErrorCode = "migration_failed"

	// ErrCodeMigrationNotSupported means no migration path exists between the two schema versions.
	ErrCodeMigrationNotSupported ErrorCode = "migration_not_supported"

	// ── Runtime / Lifecycle ──────────────────────────────────────────────────

	// ErrCodeHealthCheckFailed means the plugin failed its healthcheck.
	ErrCodeHealthCheckFailed ErrorCode = "healthcheck_failed"

	// ErrCodePluginCrashed means the plugin process exited unexpectedly.
	ErrCodePluginCrashed ErrorCode = "plugin_crashed"

	// ErrCodePluginNotReady means the plugin has not yet completed initialization.
	ErrCodePluginNotReady ErrorCode = "plugin_not_ready"
)

// PluginError is the standard error type returned in plugin response payloads.
type PluginError struct {
	Code      ErrorCode `json:"code"`
	Message   string    `json:"message"`
	Retryable bool      `json:"retryable,omitempty"`
}

func (e *PluginError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// New creates a non-retryable PluginError.
func New(code ErrorCode, message string) *PluginError {
	return &PluginError{Code: code, Message: message}
}

// NewRetryable creates a retryable PluginError.
func NewRetryable(code ErrorCode, message string) *PluginError {
	return &PluginError{Code: code, Message: message, Retryable: true}
}

// Newf creates a non-retryable PluginError with a formatted message.
func Newf(code ErrorCode, format string, args ...any) *PluginError {
	return &PluginError{Code: code, Message: fmt.Sprintf(format, args...)}
}

// Parse reconstructs a PluginError from Error() output.
func Parse(message string) (*PluginError, bool) {
	if len(message) < 4 || message[0] != '[' {
		return nil, false
	}

	end := -1
	for i := 1; i < len(message); i++ {
		if message[i] == ']' {
			end = i
			break
		}
	}
	if end <= 1 || len(message) <= end+2 || message[end+1] != ' ' {
		return nil, false
	}

	return &PluginError{
		Code:    ErrorCode(message[1:end]),
		Message: message[end+2:],
	}, true
}
