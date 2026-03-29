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

// LocalizedString stores per-locale UI text for plugins.
type LocalizedString map[string]string

// ConfigField describes a single user-configurable field for a plugin.
type ConfigField struct {
	Key             string          `json:"key"`
	Type            FieldType       `json:"type"`
	Label           string          `json:"label"`
	LabelI18n       LocalizedString `json:"label_i18n,omitempty"`
	Description     string          `json:"description,omitempty"`
	DescriptionI18n LocalizedString `json:"description_i18n,omitempty"`
	Required        bool            `json:"required,omitempty"`
	Default         any             `json:"default,omitempty"`
	EnvKey          string          `json:"env_key,omitempty"`
	Placeholder     string          `json:"placeholder,omitempty"`
	PlaceholderI18n LocalizedString `json:"placeholder_i18n,omitempty"`
	AutoComplete    string          `json:"auto_complete,omitempty"`

	// Options is used only for FieldTypeSelect.
	Options []string `json:"options,omitempty"`
}

// Schema is the versioned configuration schema declared by a plugin.
type Schema struct {
	Version string        `json:"version"`
	Fields  []ConfigField `json:"fields"`
}

// AuthActionType defines the type of an interactive authentication flow.
type AuthActionType string

const (
	// AuthActionTypePasswordGrant exchanges username/password for tokens.
	AuthActionTypePasswordGrant AuthActionType = "password_grant"
)

// AuthAction describes a declarative authentication workflow rendered by the core.
type AuthAction struct {
	ID                           string                     `json:"id"`
	Type                         AuthActionType             `json:"type"`
	Title                        string                     `json:"title"`
	TitleI18n                    LocalizedString            `json:"title_i18n,omitempty"`
	Description                  string                     `json:"description,omitempty"`
	DescriptionI18n              LocalizedString            `json:"description_i18n,omitempty"`
	ButtonLabel                  string                     `json:"button_label,omitempty"`
	ButtonLabelI18n              LocalizedString            `json:"button_label_i18n,omitempty"`
	RepeatLabel                  string                     `json:"repeat_label,omitempty"`
	RepeatLabelI18n              LocalizedString            `json:"repeat_label_i18n,omitempty"`
	DeleteLabel                  string                     `json:"delete_label,omitempty"`
	DeleteLabelI18n              LocalizedString            `json:"delete_label_i18n,omitempty"`
	Fields                       []ConfigField              `json:"fields"`
	Endpoint                     string                     `json:"endpoint,omitempty"`
	Params                       map[string]string          `json:"params,omitempty"`
	StoreMappings                map[string]string          `json:"store_mappings,omitempty"`
	RequiredMessage              string                     `json:"required_message,omitempty"`
	RequiredMessageI18n          LocalizedString            `json:"required_message_i18n,omitempty"`
	SuccessMessage               string                     `json:"success_message,omitempty"`
	SuccessMessageI18n           LocalizedString            `json:"success_message_i18n,omitempty"`
	SuccessReactivateMessage     string                     `json:"success_reactivate_message,omitempty"`
	SuccessReactivateMessageI18n LocalizedString            `json:"success_reactivate_message_i18n,omitempty"`
	ErrorMessage                 string                     `json:"error_message,omitempty"`
	ErrorMessageI18n             LocalizedString            `json:"error_message_i18n,omitempty"`
	ErrorMessages                map[string]string          `json:"error_messages,omitempty"`
	ErrorMessagesI18n            map[string]LocalizedString `json:"error_messages_i18n,omitempty"`
	DeleteMessage                string                     `json:"delete_message,omitempty"`
	DeleteMessageI18n            LocalizedString            `json:"delete_message_i18n,omitempty"`
	DeleteReactivateMessage      string                     `json:"delete_reactivate_message,omitempty"`
	DeleteReactivateMessageI18n  LocalizedString            `json:"delete_reactivate_message_i18n,omitempty"`
	DeleteErrorMessage           string                     `json:"delete_error_message,omitempty"`
	DeleteErrorMessageI18n       LocalizedString            `json:"delete_error_message_i18n,omitempty"`
	DeleteErrorMessages          map[string]string          `json:"delete_error_messages,omitempty"`
	DeleteErrorMessagesI18n      map[string]LocalizedString `json:"delete_error_messages_i18n,omitempty"`
}

// AuthSchema declares one or more authentication workflows for a plugin.
type AuthSchema struct {
	Actions []AuthAction `json:"actions,omitempty"`
}
