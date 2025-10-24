// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Enum types and constants
// ComparisonResult - Constants that indicate sort order.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ComparisonResult
type ComparisonResult int

const (
	OrderedAscending ComparisonResult = -1
	OrderedSame ComparisonResult = 0
	OrderedDescending ComparisonResult = 1
)

// EnumerationOptions - Options for block enumeration operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions
type EnumerationOptions uint

const (
	EnumerationConcurrent EnumerationOptions = 1
	EnumerationReverse EnumerationOptions = 2
)

// GrammaticalCase enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase
type GrammaticalCase uint

const (
	GrammaticalCaseNotSet GrammaticalCase = 0
	GrammaticalCaseNominative GrammaticalCase = 1
	GrammaticalCaseAccusative GrammaticalCase = 2
	GrammaticalCaseDative GrammaticalCase = 3
	GrammaticalCaseGenitive GrammaticalCase = 4
	GrammaticalCasePrepositional GrammaticalCase = 5
	GrammaticalCaseAblative GrammaticalCase = 6
	GrammaticalCaseAdessive GrammaticalCase = 7
	GrammaticalCaseAllative GrammaticalCase = 8
	GrammaticalCaseElative GrammaticalCase = 9
	GrammaticalCaseIllative GrammaticalCase = 10
	GrammaticalCaseEssive GrammaticalCase = 11
	GrammaticalCaseInessive GrammaticalCase = 12
	GrammaticalCaseLocative GrammaticalCase = 13
	GrammaticalCaseTranslative GrammaticalCase = 14
)

// GrammaticalDefiniteness enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDefiniteness
type GrammaticalDefiniteness uint

const (
	GrammaticalDefinitenessNotSet GrammaticalDefiniteness = 0
	GrammaticalDefinitenessIndefinite GrammaticalDefiniteness = 1
	GrammaticalDefinitenessDefinite GrammaticalDefiniteness = 2
)

// GrammaticalDetermination enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalDetermination
type GrammaticalDetermination uint

const (
	GrammaticalDeterminationNotSet GrammaticalDetermination = 0
	GrammaticalDeterminationIndependent GrammaticalDetermination = 1
	GrammaticalDeterminationDependent GrammaticalDetermination = 2
)

// GrammaticalPerson enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson
type GrammaticalPerson uint

const (
	GrammaticalPersonNotSet GrammaticalPerson = 0
	GrammaticalPersonFirst GrammaticalPerson = 1
	GrammaticalPersonSecond GrammaticalPerson = 2
	GrammaticalPersonThird GrammaticalPerson = 3
)

// GrammaticalPronounType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType
type GrammaticalPronounType uint

const (
	GrammaticalPronounTypeNotSet GrammaticalPronounType = 0
	GrammaticalPronounTypePersonal GrammaticalPronounType = 1
	GrammaticalPronounTypeReflexive GrammaticalPronounType = 2
	GrammaticalPronounTypePossessive GrammaticalPronounType = 3
)

// KeyValueChange - The kinds of changes that can be observed.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueChange
type KeyValueChange uint

const (
	KeyValueChangeSetting KeyValueChange = 1
	KeyValueChangeInsertion KeyValueChange = 2
	KeyValueChangeRemoval KeyValueChange = 3
	KeyValueChangeReplacement KeyValueChange = 4
)

// KeyValueObservingOptions - The values that can be returned in a change dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueObservingOptions
type KeyValueObservingOptions uint

const (
	KeyValueObservingOptionNew KeyValueObservingOptions = 1
	KeyValueObservingOptionOld KeyValueObservingOptions = 2
	KeyValueObservingOptionInitial KeyValueObservingOptions = 3
	KeyValueObservingOptionPrior KeyValueObservingOptions = 4
)

// KeyValueSetMutationKind enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSetMutationKind
type KeyValueSetMutationKind uint

const (
	KeyValueUnionSetMutation KeyValueSetMutationKind = 1
	KeyValueMinusSetMutation KeyValueSetMutationKind = 2
	KeyValueIntersectSetMutation KeyValueSetMutationKind = 3
	KeyValueSetSetMutation KeyValueSetMutationKind = 4
)

// LocaleLanguageDirection - The directions that a language may take across a page of text.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/LanguageDirection
type LocaleLanguageDirection uint

// PresentationIntentKind - An enumeration of intended display styles for blocks of text like paragraphs, lists, and code blocks.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind
type PresentationIntentKind int

const (
	PresentationIntentKindParagraph PresentationIntentKind = 0
	PresentationIntentKindHeader PresentationIntentKind = 1
	PresentationIntentKindOrderedList PresentationIntentKind = 2
	PresentationIntentKindUnorderedList PresentationIntentKind = 3
	PresentationIntentKindListItem PresentationIntentKind = 4
	PresentationIntentKindCodeBlock PresentationIntentKind = 5
	PresentationIntentKindBlockQuote PresentationIntentKind = 6
	PresentationIntentKindThematicBreak PresentationIntentKind = 7
	PresentationIntentKindTable PresentationIntentKind = 8
	PresentationIntentKindTableHeaderRow PresentationIntentKind = 9
	PresentationIntentKindTableRow PresentationIntentKind = 10
	PresentationIntentKindTableCell PresentationIntentKind = 11
)

// MatchingFlags - Set by the Block as the matching progresses, completes, or fails. Used by the method 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingFlags
type MatchingFlags uint

const (
	MatchingProgress MatchingFlags = 1
	MatchingCompleted MatchingFlags = 2
	MatchingHitEnd MatchingFlags = 4
	MatchingRequiredEnd MatchingFlags = 8
	MatchingInternalError MatchingFlags = 16
)

// MatchingOptions - The matching options constants specify the reporting, completion and matching rules to the expression matching methods. These constants are used by all methods that search for, or replace values, using a regular expression.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions
type MatchingOptions uint

const (
	// MatchingReportCompletion - Call the Block once after the completion of any matching. This option has no effect for methods other than  . See   for a description of the constant in context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions/reportCompletion
	MatchingReportCompletion MatchingOptions = 2
	// MatchingReportProgress - Call the Block periodically during long-running match operations. This option has no effect for methods other than  . See   for a description of the constant in context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions/reportProgress
	MatchingReportProgress MatchingOptions = 1
)

// RegularExpressionOptions - These constants define the regular expression options. These constants are used by the property 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct
type RegularExpressionOptions uint

const (
	// RegularExpressionAnchorsMatchLines - Allow   and   to match the start and end of lines.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct/anchorsMatchLines
	RegularExpressionAnchorsMatchLines RegularExpressionOptions = 16
	// RegularExpressionCaseInsensitive - Match letters in the pattern independent of case.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct/caseInsensitive
	RegularExpressionCaseInsensitive RegularExpressionOptions = 1
	// RegularExpressionDotMatchesLineSeparators - Allow   to match any character, including line separators.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct/dotMatchesLineSeparators
	RegularExpressionDotMatchesLineSeparators RegularExpressionOptions = 8
	// RegularExpressionUseUnicodeWordBoundaries - Use Unicode   to specify word boundaries (otherwise, traditional regular expression word boundaries are used).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct/useUnicodeWordBoundaries
	RegularExpressionUseUnicodeWordBoundaries RegularExpressionOptions = 64
)

// SortOptions - Options for block sorting operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortOptions
type SortOptions uint

const (
	SortConcurrent SortOptions = 1
	SortStable SortOptions = 16
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

// TextCheckingType - These constants specify the type of checking the methods should do. They are returned by 
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType
type TextCheckingType uint

const (
	// TextCheckingTypeRegularExpression - Matches a regular expression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType/regularExpression
	TextCheckingTypeRegularExpression TextCheckingType = 513
)

// XPCConnectionOptions - Options that you can pass to a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/Options
type XPCConnectionOptions uint

const (
	XPCConnectionPrivileged XPCConnectionOptions = 4096
)


