package types

// ContentType represents the type of content in a library.
type ContentType string

const (
	ContentTypeBook      ContentType = "book"
	ContentTypeComic     ContentType = "comic"
	ContentTypeNovel     ContentType = "novel"
	ContentTypeAudiobook ContentType = "audiobook"
)

// Language is an ISO 639-1 language code (e.g. "ko", "en", "ja").
type Language string

// SourceRef identifies a specific resource at a metadata provider.
// It is returned in SearchCandidate and MetadataResult, and used as input to FetchRequest.
type SourceRef struct {
	// ID is the provider-specific identifier for this resource.
	ID string `json:"id"`

	// Name is the provider name, e.g. "googlebooks", "anilist".
	Name string `json:"name"`

	// URL is the canonical web URL for this resource, if available.
	URL string `json:"url,omitempty"`
}
