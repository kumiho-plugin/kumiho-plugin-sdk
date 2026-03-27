package state

// State represents the lifecycle state of a plugin managed by the core.
// State transitions are unidirectional except for explicit user actions
// (enable/disable) and recovery flows (unhealthy → active after fix).
type State string

const (
	// NotInstalled means the plugin is known (e.g. in the registry) but not yet installed.
	NotInstalled State = "not_installed"

	// Downloading means the core is downloading the plugin artifact.
	Downloading State = "downloading"

	// Installed means the artifact is on disk but the plugin has not been registered yet.
	// Waiting for core restart (initial implementation) or hot registration.
	Installed State = "installed"

	// Registered means the plugin has been loaded and its manifest validated by the core.
	// The plugin is ready to be activated.
	Registered State = "registered"

	// ActivationPending means activation has been requested but not yet completed
	// (e.g. waiting for restart or async init).
	ActivationPending State = "activation_pending"

	// Active means the plugin is running and passing healthchecks.
	Active State = "active"

	// Disabled means the plugin has been explicitly deactivated by the user.
	// The artifact remains installed and can be re-enabled.
	Disabled State = "disabled"

	// Unhealthy means the plugin is running but failing healthchecks.
	// The core will retry; persistent failure leads to Error.
	Unhealthy State = "unhealthy"

	// Error means the plugin encountered a fatal error and cannot be used.
	// User action (re-install or update) is required.
	Error State = "error"

	// Incompatible means the plugin is installed but not compatible with the current core version.
	// No action is taken until the core or plugin is updated.
	Incompatible State = "incompatible"
)

// IsRunning reports whether the plugin is expected to be serving requests.
func IsRunning(s State) bool {
	return s == Active || s == Unhealthy
}

// IsInstalled reports whether the plugin artifact is present on disk.
func IsInstalled(s State) bool {
	switch s {
	case Installed, Registered, ActivationPending, Active, Disabled, Unhealthy, Error, Incompatible:
		return true
	}
	return false
}
