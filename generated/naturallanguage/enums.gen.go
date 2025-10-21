// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

// Enum types and constants
// NLContextualEmbeddingAssetsResult - The status of an asset request.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/AssetsResult
type ContextualEmbeddingAssetsResult uint

const (
// ContextualEmbeddingAssetsResultAvailable - A result that indicates assets are available.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/AssetsResult/available
ContextualEmbeddingAssetsResultAvailable ContextualEmbeddingAssetsResult = 0
// ContextualEmbeddingAssetsResultError - A result that indicates the framework encounters an error.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/AssetsResult/error
ContextualEmbeddingAssetsResultError ContextualEmbeddingAssetsResult = 0
// ContextualEmbeddingAssetsResultNotAvailable - A result that indicates assets aren’t available.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/AssetsResult/notAvailable
ContextualEmbeddingAssetsResultNotAvailable ContextualEmbeddingAssetsResult = 0
)

// NLDistanceType - The means of calculating a distance between two locations in a text embedding.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLDistanceType
type DistanceType uint

// NLModelType - The different types of a natural language model.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/ModelType
type ModelType uint

// NLTaggerAssetsResult - The response to an asset request.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/AssetsResult
type TaggerAssetsResult uint

const (
// TaggerAssetsResultAvailable - The asset is now available and loaded onto the device.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/AssetsResult/available
TaggerAssetsResultAvailable TaggerAssetsResult = 0
// TaggerAssetsResultError - The framework couldn’t load the asset due to an error.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/AssetsResult/error
TaggerAssetsResultError TaggerAssetsResult = 0
// TaggerAssetsResultNotAvailable - The asset is unavailable on the device.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/AssetsResult/notAvailable
TaggerAssetsResultNotAvailable TaggerAssetsResult = 0
)

// NLTaggerOptions - Constants for linguistic tagger enumeration specifying which tokens to omit and whether to join names.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options
type TaggerOptions uint

const (
// TaggerJoinContractions - Contractions will be returned as one token.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options/joinContractions
TaggerJoinContractions TaggerOptions = 0
// TaggerJoinNames - Typically, multiple-word names will be returned as multiple tokens, following the standard tokenization practice of the tagger.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options/joinNames
TaggerJoinNames TaggerOptions = 0
// TaggerOmitOther - Omit tokens of type   (non-linguistic items, such as symbols).
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options/omitOther
TaggerOmitOther TaggerOptions = 0
// TaggerOmitPunctuation - Omit tokens of type   (all punctuation).
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options/omitPunctuation
TaggerOmitPunctuation TaggerOptions = 0
// TaggerOmitWhitespace - Omit tokens of type   (whitespace of all sorts).
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options/omitWhitespace
TaggerOmitWhitespace TaggerOptions = 0
// TaggerOmitWords - Omit tokens of type   (items considered to be words).
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options/omitWords
TaggerOmitWords TaggerOptions = 0
)

// NLTokenUnit - Constants representing linguistic units.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit
type TokenUnit uint

const (
// TokenUnitDocument - The document in its entirety.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit/document
TokenUnitDocument TokenUnit = 0
// TokenUnitParagraph - An individual paragraph.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit/paragraph
TokenUnitParagraph TokenUnit = 0
// TokenUnitSentence - An individual sentence.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit/sentence
TokenUnitSentence TokenUnit = 0
// TokenUnitWord - An individual word.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit/word
TokenUnitWord TokenUnit = 0
)

// NLTokenizerAttributes - Hints about the contents of the string for the tokenizer.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/Attributes
type TokenizerAttributes uint

const (
// TokenizerAttributeEmoji - The string contains emoji.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/Attributes/emoji
TokenizerAttributeEmoji TokenizerAttributes = 0
// TokenizerAttributeNumeric - The string contains numbers.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/Attributes/numeric
TokenizerAttributeNumeric TokenizerAttributes = 0
// TokenizerAttributeSymbolic - The string contains symbols.
//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/Attributes/symbolic
TokenizerAttributeSymbolic TokenizerAttributes = 0
)


