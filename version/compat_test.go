package version_test

import (
	"testing"

	"github.com/kumiho-plugin/kumiho-plugin-sdk/version"
)

func TestIsCompatible(t *testing.T) {
	tests := []struct {
		name    string
		core    string
		min     string
		max     string
		want    bool
		wantErr bool
	}{
		// ── Basic range checks ──────────────────────────────────────────────
		{"exact match", "1.0.0", "1.0.0", "1.0.0", true, false},
		{"within range", "1.3.0", "1.0.0", "1.99.99", true, false},
		{"at min boundary", "1.0.0", "1.0.0", "2.0.0", true, false},
		{"at max boundary", "2.0.0", "1.0.0", "2.0.0", true, false},

		// ── Out of range ────────────────────────────────────────────────────
		{"below min", "0.9.0", "1.0.0", "", false, false},
		{"above max", "2.1.0", "", "2.0.0", false, false},
		{"below min patch", "1.0.0", "1.0.1", "2.0.0", false, false},
		{"above max patch", "1.0.2", "1.0.0", "1.0.1", false, false},

		// ── Unbounded ───────────────────────────────────────────────────────
		{"no bounds", "1.0.0", "", "", true, false},
		{"no min only", "5.0.0", "", "5.0.0", true, false},
		{"no max only", "0.1.0", "0.1.0", "", true, false},

		// ── v prefix ────────────────────────────────────────────────────────
		{"v prefix core", "v1.2.3", "1.0.0", "2.0.0", true, false},
		{"v prefix all", "v1.2.3", "v1.0.0", "v2.0.0", true, false},

		// ── Major version transitions ────────────────────────────────────────
		{"major bump rejected", "2.0.0", "1.0.0", "1.99.99", false, false},
		{"major bump allowed", "2.0.0", "1.0.0", "2.99.99", true, false},

		// ── Pre-release suffix stripping ─────────────────────────────────────
		{"pre-release suffix core", "1.2.0-rc1", "1.0.0", "2.0.0", true, false},
		{"pre-release suffix min", "1.2.0", "1.0.0-beta", "2.0.0", true, false},

		// ── Error cases ──────────────────────────────────────────────────────
		{"invalid core", "not-a-version", "1.0.0", "2.0.0", false, true},
		{"invalid min", "1.0.0", "bad", "2.0.0", false, true},
		{"invalid max", "1.0.0", "1.0.0", "bad", false, true},
		{"missing patch core", "1.0", "1.0.0", "2.0.0", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := version.IsCompatible(tt.core, tt.min, tt.max)
			if (err != nil) != tt.wantErr {
				t.Fatalf("IsCompatible(%q, %q, %q) error = %v, wantErr %v",
					tt.core, tt.min, tt.max, err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("IsCompatible(%q, %q, %q) = %v, want %v",
					tt.core, tt.min, tt.max, got, tt.want)
			}
		})
	}
}
