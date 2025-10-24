// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

/* debug [enums.gen.go]: Generating 7 enums for NaturalLanguage */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum NLContextualEmbeddingAssetsResult (3 cases) */
// NLContextualEmbeddingAssetsResult - The status of an asset request.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/AssetsResult
type NLContextualEmbeddingAssetsResult uint

const (
	// NLContextualEmbeddingAssetsResultAvailable - A result that indicates assets are available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/AssetsResult/available
	NLContextualEmbeddingAssetsResultAvailable NLContextualEmbeddingAssetsResult = 0
	// NLContextualEmbeddingAssetsResultError - A result that indicates the framework encounters an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/AssetsResult/error
	NLContextualEmbeddingAssetsResultError NLContextualEmbeddingAssetsResult = 0
	// NLContextualEmbeddingAssetsResultNotAvailable - A result that indicates assets aren’t available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/AssetsResult/notAvailable
	NLContextualEmbeddingAssetsResultNotAvailable NLContextualEmbeddingAssetsResult = 0
)

/* debug [enums.gen.go]: Processing enum NLModelType (2 cases) */
// NLModelType - The different types of a natural language model.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/ModelType
type NLModelType uint

const (
	// NLModelTypeClassifier - A classifier model type that tags text at the phrase, sentence, paragraph, or higher level.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/ModelType/classifier
	NLModelTypeClassifier NLModelType = 0
	// NLModelTypeSequence - A sequence model type that tags text at the token level.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLModel/ModelType/sequence
	NLModelTypeSequence NLModelType = 0
)

/* debug [enums.gen.go]: Processing enum NLTaggerAssetsResult (3 cases) */
// NLTaggerAssetsResult - The response to an asset request.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/AssetsResult
type NLTaggerAssetsResult uint

const (
	// NLTaggerAssetsResultAvailable - The asset is now available and loaded onto the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/AssetsResult/available
	NLTaggerAssetsResultAvailable NLTaggerAssetsResult = 0
	// NLTaggerAssetsResultError - The framework couldn’t load the asset due to an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/AssetsResult/error
	NLTaggerAssetsResultError NLTaggerAssetsResult = 0
	// NLTaggerAssetsResultNotAvailable - The asset is unavailable on the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/AssetsResult/notAvailable
	NLTaggerAssetsResultNotAvailable NLTaggerAssetsResult = 0
)

/* debug [enums.gen.go]: Processing enum NLTaggerOptions (6 cases) */
// NLTaggerOptions - Constants for linguistic tagger enumeration specifying which tokens to omit and whether to join names.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options
type NLTaggerOptions uint

const (
	// NLTaggerJoinContractions - Contractions will be returned as one token.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options/joinContractions
	NLTaggerJoinContractions NLTaggerOptions = 0
	// NLTaggerJoinNames - Typically, multiple-word names will be returned as multiple tokens, following the standard tokenization practice of the tagger.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options/joinNames
	NLTaggerJoinNames NLTaggerOptions = 0
	// NLTaggerOmitOther - Omit tokens of type   (non-linguistic items, such as symbols).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options/omitOther
	NLTaggerOmitOther NLTaggerOptions = 0
	// NLTaggerOmitPunctuation - Omit tokens of type   (all punctuation).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options/omitPunctuation
	NLTaggerOmitPunctuation NLTaggerOptions = 0
	// NLTaggerOmitWhitespace - Omit tokens of type   (whitespace of all sorts).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options/omitWhitespace
	NLTaggerOmitWhitespace NLTaggerOptions = 0
	// NLTaggerOmitWords - Omit tokens of type   (items considered to be words).
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options/omitWords
	NLTaggerOmitWords NLTaggerOptions = 0
)

/* debug [enums.gen.go]: Processing enum NLTokenizerAttributes (3 cases) */
// NLTokenizerAttributes - Hints about the contents of the string for the tokenizer.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/Attributes
type NLTokenizerAttributes uint

const (
	// NLTokenizerAttributeEmoji - The string contains emoji.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/Attributes/emoji
	NLTokenizerAttributeEmoji NLTokenizerAttributes = 0
	// NLTokenizerAttributeNumeric - The string contains numbers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/Attributes/numeric
	NLTokenizerAttributeNumeric NLTokenizerAttributes = 0
	// NLTokenizerAttributeSymbolic - The string contains symbols.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/Attributes/symbolic
	NLTokenizerAttributeSymbolic NLTokenizerAttributes = 0
)

/* debug [enums.gen.go]: Processing enum NLDistanceType (1 cases) */
// NLDistanceType - The means of calculating a distance between two locations in a text embedding.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLDistanceType
type NLDistanceType uint

const (
	// NLDistanceTypeCosine - A method of calculating distance by using cosine similarity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLDistanceType/cosine
	NLDistanceTypeCosine NLDistanceType = 0
)

/* debug [enums.gen.go]: Processing enum NLTokenUnit (4 cases) */
// NLTokenUnit - Constants representing linguistic units.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit
type NLTokenUnit uint

const (
	// NLTokenUnitDocument - The document in its entirety.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit/document
	NLTokenUnitDocument NLTokenUnit = 0
	// NLTokenUnitParagraph - An individual paragraph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit/paragraph
	NLTokenUnitParagraph NLTokenUnit = 0
	// NLTokenUnitSentence - An individual sentence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit/sentence
	NLTokenUnitSentence NLTokenUnit = 0
	// NLTokenUnitWord - An individual word.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit/word
	NLTokenUnitWord NLTokenUnit = 0
)


