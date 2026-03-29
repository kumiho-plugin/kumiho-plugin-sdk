package manifest

import (
	"github.com/kumiho-plugin/kumiho-plugin-sdk/capability"
	sdkconfig "github.com/kumiho-plugin/kumiho-plugin-sdk/config"
)

// TrustLevel indicates how much the core trusts a plugin.
type TrustLevel string

const (
	// TrustLevelOfficial is an officially maintained Kumiho plugin.
	TrustLevelOfficial TrustLevel = "official"

	// TrustLevelVerifiedThirdParty is a third-party plugin that has been signature-verified.
	TrustLevelVerifiedThirdParty TrustLevel = "verified_third_party"

	// TrustLevelUnverified is an unverified plugin. The user installs at their own risk.
	TrustLevelUnverified TrustLevel = "unverified"
)

// RuntimeType indicates how the plugin is executed.
type RuntimeType string

const (
	// RuntimeTypeBinary means the plugin runs as a child process managed by the core.
	// Official Go plugins use HashiCorp go-plugin under this runtime.
	RuntimeTypeBinary RuntimeType = "binary"

	// RuntimeTypeService means the plugin runs as an endpoint-based service
	// (e.g. Docker container). The core communicates via HTTP JSON or gRPC.
	RuntimeTypeService RuntimeType = "service"
)

// Platform identifies an OS/runtime target for a plugin artifact.
type Platform string

const (
	PlatformLinuxDocker   Platform = "linux/docker"
	PlatformLinuxBinary   Platform = "linux/binary"
	PlatformWindowsBinary Platform = "windows/binary"
	PlatformMacOSBinary   Platform = "macos/binary"
)

// Artifact describes a single downloadable plugin binary or image.
type Artifact struct {
	Platform  Platform `json:"platform"`
	URL       string   `json:"url"`
	Checksum  string   `json:"checksum"`            // format: "sha256:<hex>"
	Signature string   `json:"signature,omitempty"` // detached signature for verified plugins
}

// IconSet describes optional visual assets for plugin UIs.
// The core should prefer SVG when available and fall back to PNG otherwise.
type IconSet struct {
	SVG string `json:"svg,omitempty"`
	PNG string `json:"png,omitempty"`
}

// Manifest is the static metadata declaration of a plugin.
// It is stored in the plugin registry and validated by the core on install.
type Manifest struct {
	// Identity
	ID          string   `json:"id"`   // e.g. "kumiho-plugin-metadata-kitsu"
	Name        string   `json:"name"` // e.g. "Kitsu Manga"
	Description string   `json:"description"`
	Version     string   `json:"version"` // semver, e.g. "1.2.0"
	Author      string   `json:"author"`
	Homepage    string   `json:"homepage,omitempty"`
	Repository  string   `json:"repository,omitempty"`
	License     string   `json:"license,omitempty"`
	Icons       *IconSet `json:"icons,omitempty"`

	// Trust & runtime
	TrustLevel         TrustLevel  `json:"trust_level"`
	RuntimeType        RuntimeType `json:"runtime_type"`
	SupportedPlatforms []Platform  `json:"supported_platforms"`

	// Capabilities this plugin provides
	Capabilities []capability.Capability `json:"capabilities"`

	// Permissions this plugin requires (informational; shown to user on install)
	Permissions []string `json:"permissions,omitempty"`

	// Core version compatibility
	MinCoreVersion string `json:"min_core_version,omitempty"`
	MaxCoreVersion string `json:"max_core_version,omitempty"`

	// SDKVersion is the version of kumiho-plugin-sdk the plugin was built against.
	// The core uses this to detect incompatible SDK versions at activation.
	SDKVersion string `json:"sdk_version,omitempty"`

	// Config schema version declared by this plugin version
	ConfigSchemaVersion string `json:"config_schema_version"`
	ConfigSchema        *sdkconfig.Schema     `json:"config_schema,omitempty"`
	Auth                *sdkconfig.AuthSchema `json:"auth,omitempty"`

	// Downloadable artifacts, one per supported platform
	Artifacts []Artifact `json:"artifacts"`
}
