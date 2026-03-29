package types

import pluginerrors "github.com/kumiho-plugin/kumiho-plugin-sdk/errors"

// SearchRequest is the input sent to a plugin's metadata.search capability.
// Fields are optional; plugins must handle partial inputs gracefully.
type SearchRequest struct {
	// Local file information
	LocalTitle    string      `json:"local_title,omitempty"`
	Filename      string      `json:"filename,omitempty"`
	SeriesName    string      `json:"series_name,omitempty"`
	VolumeNumber  *int        `json:"volume_number,omitempty"`
	ChapterNumber *float64    `json:"chapter_number,omitempty"`
	Language      Language    `json:"language,omitempty"`
	AuthorHint    string      `json:"author_hint,omitempty"`
	ContentType   ContentType `json:"content_type,omitempty"`

	// Known identifiers for direct lookup (e.g. "isbn", "asin", "anilist_id").
	// If present, the plugin should prioritize identifier-based lookup over text search.
	Identifiers map[string]string `json:"identifiers,omitempty"`

	// Maximum number of candidates to return (0 = provider default).
	Limit int `json:"limit,omitempty"`
}

// SearchCandidate is a single search result returned by the plugin.
type SearchCandidate struct {
	// Source identifies this candidate at the provider.
	// Use Source.ID and Source.Name as FetchRequest input.
	Source SourceRef `json:"source"`

	Title         string      `json:"title"`
	OriginalTitle string      `json:"original_title,omitempty"`
	Authors       []string    `json:"authors,omitempty"`
	Description   string      `json:"description,omitempty"`
	ContentType   ContentType `json:"content_type,omitempty"`
	Year          *int        `json:"year,omitempty"`
	CoverURL      string      `json:"cover_url,omitempty"`

	// Score is the provider's relevance score, normalized to [0.0, 1.0].
	Score float64 `json:"score"`

	// Confidence is the plugin's own confidence in this candidate, normalized to [0.0, 1.0].
	// Plugins should factor in identifier matches, title similarity, author match, etc.
	Confidence float64 `json:"confidence"`

	// Reason is a human-readable explanation for the confidence value.
	Reason string `json:"reason,omitempty"`
}

// SearchResponse is the output of a plugin's metadata.search capability.
type SearchResponse struct {
	Candidates []SearchCandidate        `json:"candidates"`
	Error      *pluginerrors.PluginError `json:"error,omitempty"`
}
