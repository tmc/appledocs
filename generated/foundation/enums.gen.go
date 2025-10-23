// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Enum types and constants
// DataBase64DecodingOptions - Options to modify the decoding algorithm used to decode Base64 encoded data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64DecodingOptions
type DataBase64DecodingOptions uint

const (
	DataBase64DecodingIgnoreUnknownCharacters DataBase64DecodingOptions = 1
)

// DataBase64EncodingOptions - Options for methods used to Base64 encode data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/Base64EncodingOptions
type DataBase64EncodingOptions uint

const (
	DataBase64Encoding64CharacterLineLength DataBase64EncodingOptions = 1
	DataBase64Encoding76CharacterLineLength DataBase64EncodingOptions = 2
	DataBase64EncodingEndLineWithCarriageReturn DataBase64EncodingOptions = 16
	DataBase64EncodingEndLineWithLineFeed DataBase64EncodingOptions = 32
)

// DataCompressionAlgorithm - An algorithm that indicates how to compress or decompress data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm
type DataCompressionAlgorithm uint

const (
	DataCompressionAlgorithmLZFSE DataCompressionAlgorithm = 0
	DataCompressionAlgorithmLZ4 DataCompressionAlgorithm = 1
	DataCompressionAlgorithmLZMA DataCompressionAlgorithm = 2
	DataCompressionAlgorithmZlib DataCompressionAlgorithm = 3
)

// DataReadingOptions - Options for methods used to read data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions
type DataReadingOptions uint

const (
	DataReadingMappedIfSafe DataReadingOptions = 1
	DataReadingUncached DataReadingOptions = 2
	DataReadingMappedAlways DataReadingOptions = 3
	DataReadingMapped DataReadingOptions = 4
	MappedRead DataReadingOptions = 5
	UncachedRead DataReadingOptions = 6
)

// DataSearchOptions - Options for method used to search data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/SearchOptions
type DataSearchOptions uint

const (
	DataSearchBackwards DataSearchOptions = 1
	DataSearchAnchored DataSearchOptions = 2
)

// DataWritingOptions - Options for methods used to write data objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions
type DataWritingOptions uint

const (
	DataWritingAtomic DataWritingOptions = 1
	DataWritingWithoutOverwriting DataWritingOptions = 2
	DataWritingFileProtectionNone DataWritingOptions = 3
	DataWritingFileProtectionComplete DataWritingOptions = 4
	DataWritingFileProtectionCompleteUnlessOpen DataWritingOptions = 5
	DataWritingFileProtectionCompleteUntilFirstUserAuthentication DataWritingOptions = 6
	DataWritingFileProtectionCompleteWhenUserInactive DataWritingOptions = 7
	DataWritingFileProtectionMask DataWritingOptions = 8
	AtomicWrite DataWritingOptions = 9
)

// StringCompareOptions - These values represent the options available to many of the string classes’ search and comparison methods.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions
type StringCompareOptions uint

const (
	CaseInsensitiveSearch StringCompareOptions = 1
	LiteralSearch StringCompareOptions = 2
	BackwardsSearch StringCompareOptions = 4
	AnchoredSearch StringCompareOptions = 8
	NumericSearch StringCompareOptions = 64
	DiacriticInsensitiveSearch StringCompareOptions = 65
	WidthInsensitiveSearch StringCompareOptions = 66
	ForcedOrderingSearch StringCompareOptions = 67
	RegularExpressionSearch StringCompareOptions = 68
)

// StringDrawingOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions
type StringDrawingOptions uint

// StringEncodingConversionOptions - Options for converting string encodings.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EncodingConversionOptions
type StringEncodingConversionOptions uint

const (
	StringEncodingConversionAllowLossy StringEncodingConversionOptions = 1
	StringEncodingConversionExternalRepresentation StringEncodingConversionOptions = 2
)

// StringEnumerationOptions - Constants to specify kinds of substrings and styles of enumeration.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions
type StringEnumerationOptions uint

const (
	StringEnumerationByLines StringEnumerationOptions = 0
	StringEnumerationByParagraphs StringEnumerationOptions = 1
	StringEnumerationByComposedCharacterSequences StringEnumerationOptions = 2
	StringEnumerationByWords StringEnumerationOptions = 3
	StringEnumerationBySentences StringEnumerationOptions = 4
	StringEnumerationByCaretPositions StringEnumerationOptions = 5
	StringEnumerationByDeletionClusters StringEnumerationOptions = 6
	StringEnumerationReverse StringEnumerationOptions = 256
	StringEnumerationSubstringNotRequired StringEnumerationOptions = 512
	StringEnumerationLocalized StringEnumerationOptions = 1024
)


