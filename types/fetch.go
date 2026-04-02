package types

import pluginerrors "github.com/kumiho-plugin/kumiho-plugin-sdk/errors"

// FetchRequest is the input sent to a plugin's metadata.fetch capability.
// It references a candidate returned by a prior SearchResponse.
type FetchRequest struct {
	Source SourceRef `json:"source"`
}

// CoverInfo holds metadata about a cover image.
type CoverInfo struct {
	URL    string `json:"url"`
	Width  *int   `json:"width,omitempty"`
	Height *int   `json:"height,omitempty"`
}

// MetadataCharacter is a normalized character entry returned by metadata.fetch.
type MetadataCharacter struct {
	ID          string            `json:"id,omitempty"`
	Name        string            `json:"name"`
	Role        string            `json:"role,omitempty"`
	Image       *CoverInfo        `json:"image,omitempty"`
	Identifiers map[string]string `json:"identifiers,omitempty"`
}

// MetadataResult is the full metadata payload returned by metadata.fetch.
type MetadataResult struct {
	// Source is the provider reference for this result.
	Source SourceRef `json:"source"`

	Title          string            `json:"title"`
	OriginalTitle  string            `json:"original_title,omitempty"`
	OriginalTitles map[string]string `json:"original_titles,omitempty"`
	Authors        []string          `json:"authors,omitempty"`
	Description    string            `json:"description,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	ContentType    ContentType       `json:"content_type,omitempty"`
	Language       Language          `json:"language,omitempty"`

	// PublicationDate is in YYYY-MM-DD or YYYY format.
	PublicationDate string `json:"publication_date,omitempty"`
	Publisher       string `json:"publisher,omitempty"`

	// Identifiers holds known external IDs for this work
	// (e.g. "isbn13", "asin", "anilist_id", "provider_id").
	Identifiers map[string]string `json:"identifiers,omitempty"`

	SeriesName   string `json:"series_name,omitempty"`
	VolumeNumber *int   `json:"volume_number,omitempty"`

	Cover      *CoverInfo          `json:"cover,omitempty"`
	Characters []MetadataCharacter `json:"characters,omitempty"`
}

// FetchResponse is the output of a plugin's metadata.fetch capability.
type FetchResponse struct {
	Result *MetadataResult           `json:"result,omitempty"`
	Error  *pluginerrors.PluginError `json:"error,omitempty"`
}
