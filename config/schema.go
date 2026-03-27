package config

// FieldType defines the data type of a config field.
type FieldType string

const (
	// FieldTypeString is a plain text field.
	FieldTypeString FieldType = "string"

	// FieldTypeSecret is a sensitive field (API key, token).
	// It is masked in the UI and encrypted at rest.
	FieldTypeSecret FieldType = "secret"

	// FieldTypeBoolean is a true/false toggle.
	FieldTypeBoolean FieldType = "boolean"

	// FieldTypeInteger is a whole number field.
	FieldTypeInteger FieldType = "integer"

	// FieldTypeSelect is an enum field with a fixed list of options.
	FieldTypeSelect FieldType = "select"
)

// ConfigField describes a single user-configurable field for a plugin.
type ConfigField struct {
	Key         string    `json:"key"`
	Type        FieldType `json:"type"`
	Label       string    `json:"label"`
	Description string    `json:"description,omitempty"`
	Required    bool      `json:"required,omitempty"`
	Default     any       `json:"default,omitempty"`

	// Options is used only for FieldTypeSelect.
	Options []string `json:"options,omitempty"`
}

// Schema is the versioned configuration schema declared by a plugin.
type Schema struct {
	Version string        `json:"version"`
	Fields  []ConfigField `json:"fields"`
}
