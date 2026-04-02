// Package version defines Kumiho plugin SDK versioning policy and
// provides helpers for core/plugin compatibility checks.
//
// # Versioning Policy
//
// Both the Kumiho core and plugins follow Semantic Versioning (semver):
//
//	vMAJOR.MINOR.PATCH
//
// Compatibility rules:
//   - A plugin declares MinCoreVersion and MaxCoreVersion in its manifest.
//   - The core checks its own version against these bounds at plugin activation.
//   - MAJOR version bumps in the core indicate breaking changes; plugins must
//     declare a compatible MaxCoreVersion.
//   - MINOR and PATCH bumps are backward-compatible within the same MAJOR.
//   - An empty MinCoreVersion means "no lower bound".
//   - An empty MaxCoreVersion means "no upper bound".
//
// # SDK Version
//
// Plugins should record the SDK version they were built against in their manifest.
// The core may reject plugins built against an incompatible SDK version.
package version

import (
	"fmt"
	"strconv"
	"strings"
)

// SDK is the current version of kumiho-plugin-sdk.
const SDK = "0.1.3"

// IsCompatible reports whether coreVersion satisfies the constraint
// minVersion <= coreVersion <= maxVersion.
//
// Versions must be in "MAJOR.MINOR.PATCH" or "vMAJOR.MINOR.PATCH" format.
// An empty minVersion or maxVersion means the bound is unconstrained.
func IsCompatible(coreVersion, minVersion, maxVersion string) (bool, error) {
	cv, err := parse(coreVersion)
	if err != nil {
		return false, fmt.Errorf("invalid core version %q: %w", coreVersion, err)
	}
	if minVersion != "" {
		min, err := parse(minVersion)
		if err != nil {
			return false, fmt.Errorf("invalid min_core_version %q: %w", minVersion, err)
		}
		if cv.less(min) {
			return false, nil
		}
	}
	if maxVersion != "" {
		max, err := parse(maxVersion)
		if err != nil {
			return false, fmt.Errorf("invalid max_core_version %q: %w", maxVersion, err)
		}
		if max.less(cv) {
			return false, nil
		}
	}
	return true, nil
}

// semVer holds a parsed semantic version.
type semVer struct {
	major, minor, patch int
}

func (a semVer) less(b semVer) bool {
	if a.major != b.major {
		return a.major < b.major
	}
	if a.minor != b.minor {
		return a.minor < b.minor
	}
	return a.patch < b.patch
}

func parse(v string) (semVer, error) {
	v = strings.TrimPrefix(v, "v")
	parts := strings.SplitN(v, ".", 3)
	if len(parts) != 3 {
		return semVer{}, fmt.Errorf("expected MAJOR.MINOR.PATCH, got %q", v)
	}
	nums := make([]int, 3)
	for i, p := range parts {
		// Strip any pre-release or build metadata suffix (e.g. "1-rc1" → 1)
		p = strings.FieldsFunc(p, func(r rune) bool {
			return r == '-' || r == '+'
		})[0]
		n, err := strconv.Atoi(p)
		if err != nil {
			return semVer{}, fmt.Errorf("non-numeric segment %q in version", p)
		}
		nums[i] = n
	}
	return semVer{nums[0], nums[1], nums[2]}, nil
}
