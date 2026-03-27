package config

// OpType defines the type of a declarative migration operation.
type OpType string

const (
	// OpTypeRenameField renames an existing field key.
	OpTypeRenameField OpType = "rename_field"

	// OpTypeAddField adds a new field with a default value.
	// No-op if the field already exists.
	OpTypeAddField OpType = "add_field"

	// OpTypeRemoveField removes a field that no longer exists in the new schema.
	OpTypeRemoveField OpType = "remove_field"

	// OpTypeMoveSecret moves a secret value from one key to another.
	OpTypeMoveSecret OpType = "move_secret"
)

// MigrationOp is a single declarative migration step.
type MigrationOp struct {
	Type OpType `json:"type"`

	// OldKey is the source field key (used by rename, remove, move_secret).
	OldKey string `json:"old_key,omitempty"`

	// NewKey is the destination field key (used by rename, add, move_secret).
	NewKey string `json:"new_key,omitempty"`

	// Default is the value to assign when adding a new field.
	Default any `json:"default,omitempty"`
}

// MigrationDescriptor describes how to migrate plugin config from one schema version to another.
// Plugins provide this as a static declaration; the core executes it.
type MigrationDescriptor struct {
	FromVersion string        `json:"from_version"`
	ToVersion   string        `json:"to_version"`
	Ops         []MigrationOp `json:"ops"`

	// ManualReviewRequired signals that the migration cannot be automated.
	// The core will mark the plugin config as requiring user review.
	ManualReviewRequired bool `json:"manual_review_required,omitempty"`
}
