package types

// TranslateRequest represents a request to translate text.
type TranslateRequest struct {
	// Text to translate.
	Text []string `json:"text"`
	// TargetLang is the language code to translate into (e.g., "ko", "en-US").
	TargetLang string `json:"target_lang"`
}

// TranslationResult represents the result of translating a single text string.
type TranslationResult struct {
	// Translated text.
	Text string `json:"text"`
	// DetectedSourceLanguage is the language code detected from the source text.
	DetectedSourceLanguage string `json:"detected_source_language,omitempty"`
}

// TranslateResponse represents the result of translating text.
type TranslateResponse struct {
	// Translations contains the translated texts corresponding to the input.
	Translations []TranslationResult `json:"translations"`
	// Error contains an error message if the translation failed.
	Error string `json:"error,omitempty"`
}

// DetectRequest represents a request to detect the language of text.
type DetectRequest struct {
	// Text to detect.
	Text string `json:"text"`
}

// DetectResponse represents the result of language detection.
type DetectResponse struct {
	// Language is the detected language code.
	Language string `json:"language"`
	// Confidence is the confidence score of the detection, from 0.0 to 1.0.
	// Providers may leave this unset when they do not expose confidence data.
	Confidence float64 `json:"confidence,omitempty"`
	// Error contains an error message if the detection failed.
	Error string `json:"error,omitempty"`
}
