package capability

// Capability represents a feature that a plugin can provide.
// Plugins declare which capabilities they support in their manifest.
type Capability string

const (
	// ── Metadata ─────────────────────────────────────────────────────────────

	// MetadataSearch searches for metadata candidates based on local file information.
	MetadataSearch Capability = "metadata.search"

	// MetadataFetch fetches full metadata details for a selected candidate.
	MetadataFetch Capability = "metadata.fetch"

	// ── Cover ─────────────────────────────────────────────────────────────────

	// CoverSearch searches for cover image candidates.
	CoverSearch Capability = "cover.search"

	// CoverFetch fetches a specific cover image.
	CoverFetch Capability = "cover.fetch"

	// ── Identifier ───────────────────────────────────────────────────────────

	// IdentifierLookup performs a direct lookup by identifier (ISBN, ASIN, AniList ID, etc.).
	IdentifierLookup Capability = "identifier.lookup"

	// ── Extended (future) ────────────────────────────────────────────────────

	// SeriesMatch matches a work to a series entry and returns series metadata.
	SeriesMatch Capability = "series.match"

	// PersonMatch matches an author/artist name to a person record.
	PersonMatch Capability = "person.match"
)
