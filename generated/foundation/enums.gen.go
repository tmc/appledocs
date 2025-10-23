// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Enum types and constants
// DecodingFailurePolicy - Policies describing the action the coder should take when encountering decode failures.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/DecodingFailurePolicy-swift.enum
type DecodingFailurePolicy uint

const (
	// DecodingFailurePolicyRaiseException - A failure policy that directs the coder to raise an exception.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/DecodingFailurePolicy-swift.enum/raiseException
	DecodingFailurePolicyRaiseException DecodingFailurePolicy = 0
	// DecodingFailurePolicySetErrorAndReturn - A failure policy that directs the coder to capture the failure as an error object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/DecodingFailurePolicy-swift.enum/setErrorAndReturn
	DecodingFailurePolicySetErrorAndReturn DecodingFailurePolicy = 1
)

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
	// AnchoredSearch - Search is limited to start (or end, if  ) of source string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/anchored
	AnchoredSearch StringCompareOptions = 8
	// BackwardsSearch - Search from end of source string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/backwards
	BackwardsSearch StringCompareOptions = 4
	// CaseInsensitiveSearch - A case-insensitive search.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/caseInsensitive
	CaseInsensitiveSearch StringCompareOptions = 1
	// DiacriticInsensitiveSearch - Search ignores diacritic marks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/diacriticInsensitive
	DiacriticInsensitiveSearch StringCompareOptions = 65
	// ForcedOrderingSearch - Comparisons are forced to return either   or   if the strings are equivalent but not strictly equal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/forcedOrdering
	ForcedOrderingSearch StringCompareOptions = 67
	// LiteralSearch - Exact character-by-character equivalence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/literal
	LiteralSearch StringCompareOptions = 2
	// NumericSearch - Numbers within strings are compared using numeric value, that is,   <   <  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/numeric
	NumericSearch StringCompareOptions = 64
	// RegularExpressionSearch - The search string is treated as an ICU-compatible regular expression. If set, no other options can apply except   and  . You can use this option only with the  … methods and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/regularExpression
	RegularExpressionSearch StringCompareOptions = 68
	// WidthInsensitiveSearch - Search ignores width differences in characters that have full-width and half-width forms, as occurs in East Asian character sets.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions/widthInsensitive
	WidthInsensitiveSearch StringCompareOptions = 66
)

// StringDrawingOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions
type StringDrawingOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/disableScreenFontSubstitution
	disableScreenFontSubstitution StringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/oneShot
	oneShot StringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/optionsResolvesNaturalAlignmentWithBaseWritingDirection
	optionsResolvesNaturalAlignmentWithBaseWritingDirection StringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/truncatesLastVisibleLine
	truncatesLastVisibleLine StringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/usesDeviceMetrics
	usesDeviceMetrics StringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/usesFontLeading
	usesFontLeading StringDrawingOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions/usesLineFragmentOrigin
	usesLineFragmentOrigin StringDrawingOptions = 0
)

// StringEncodingConversionOptions - Options for converting string encodings.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EncodingConversionOptions
type StringEncodingConversionOptions uint

const (
	// StringEncodingConversionAllowLossy - Allows lossy conversion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EncodingConversionOptions/allowLossy
	StringEncodingConversionAllowLossy StringEncodingConversionOptions = 1
	// StringEncodingConversionExternalRepresentation - Specifies an external representation (with a byte-order mark, if necessary, to indicate endianness).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EncodingConversionOptions/externalRepresentation
	StringEncodingConversionExternalRepresentation StringEncodingConversionOptions = 2
)

// StringEnumerationOptions - Constants to specify kinds of substrings and styles of enumeration.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions
type StringEnumerationOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byCaretPositions
	StringEnumerationByCaretPositions StringEnumerationOptions = 5
	// StringEnumerationByComposedCharacterSequences - Enumerates by composed character sequences. Equivalent to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byComposedCharacterSequences
	StringEnumerationByComposedCharacterSequences StringEnumerationOptions = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byDeletionClusters
	StringEnumerationByDeletionClusters StringEnumerationOptions = 6
	// StringEnumerationByLines - Enumerates by lines. Equivalent to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byLines
	StringEnumerationByLines StringEnumerationOptions = 0
	// StringEnumerationByParagraphs - Enumerates by paragraphs. Equivalent to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byParagraphs
	StringEnumerationByParagraphs StringEnumerationOptions = 1
	// StringEnumerationBySentences - Enumerates by sentences.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/bySentences
	StringEnumerationBySentences StringEnumerationOptions = 4
	// StringEnumerationByWords - Enumerates by words.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/byWords
	StringEnumerationByWords StringEnumerationOptions = 3
	// StringEnumerationLocalized - Causes the enumeration to occur using the current locale. This does not make a difference in line, paragraph, or composed character sequence enumeration, but it may for words or sentences.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/localized
	StringEnumerationLocalized StringEnumerationOptions = 1024
	// StringEnumerationReverse - Causes enumeration to occur from the end of the specified range to the start.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/reverse
	StringEnumerationReverse StringEnumerationOptions = 256
	// StringEnumerationSubstringNotRequired - A way to indicate that the block does not need substring, in which case   will be passed. This is simply a performance shortcut.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions/substringNotRequired
	StringEnumerationSubstringNotRequired StringEnumerationOptions = 512
)


