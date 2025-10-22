// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextCheckingResult] class.
var (
	TextCheckingResultClass     _TextCheckingResultClass
	TextCheckingResultClassOnce sync.Once
)

func getTextCheckingResultClass() _TextCheckingResultClass {
	TextCheckingResultClassOnce.Do(func() {
		TextCheckingResultClass = _TextCheckingResultClass{objc.GetClass("NSTextCheckingResult")}
	})
	return TextCheckingResultClass
}

type _TextCheckingResultClass struct {
	class objc.Class
}

// An interface definition for the [TextCheckingResult] class.
type ITextCheckingResult interface {
	objectivec.IObject
	Date() NSDate
	Duration() TimeInterval
	PhoneNumber() string
	Range() Range
	RegularExpression() NSRegularExpression
	ResultType() TextCheckingType
	NSNotFound() int
	AddressComponents() string
	SetAddressComponents(value string)
	AlternativeStrings() string
	SetAlternativeStrings(value string)
	Components() string
	SetComponents(value string)
	GrammarDetails() string
	SetGrammarDetails(value string)
	NumberOfRanges() int
	SetNumberOfRanges(value int)
	Orthography() NSOrthography
	SetOrthography(value IOrthography)
	ReplacementString() string
	SetReplacementString(value string)
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
	Url() URL
	SetUrl(value IURL)
}

// An occurrence of textual content found during the analysis of a block of text, such as when matching a regular expression.
//
// On both iOS and macOS, instances of are returned by the class and the class to indicate the discovery of content. In those cases, what is found may be a match for a regular expression or a date, address, phone number, and so on. In macOS, instances of are returned by the object to describe the results of spelling, grammar, or text-substitution actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult
type TextCheckingResult struct {
	objectivec.Object
}

// TextCheckingResultFrom constructs a [TextCheckingResult] from an unsafe.Pointer.
//
// An occurrence of textual content found during the analysis of a block of text, such as when matching a regular expression.
func TextCheckingResultFrom(ptr unsafe.Pointer) TextCheckingResult {
	return TextCheckingResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextCheckingResultClass) Alloc() TextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextCheckingResultClass) New() TextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextCheckingResult) Init() TextCheckingResult {
	rv := objc.Send[TextCheckingResult](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextCheckingResult) Autorelease() TextCheckingResult {
	rv := objc.Send[TextCheckingResult](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextCheckingResult creates a new TextCheckingResult instance.
func NewTextCheckingResult() TextCheckingResult {
	return getTextCheckingResultClass().New()
}



// Creates and returns a text checking result with the specified address components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/addressCheckingResult(range:components:)

func (tc _TextCheckingResultClass) AddressCheckingResultWithRangeComponents(range_ IRange, components unsafe.Pointer) TextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("addressCheckingResultWithRange:components:"), range_, components)
	return rv
}


// Creates and returns a text checking result with the specified phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/phoneNumberCheckingResult(range:phoneNumber:)

func (tc _TextCheckingResultClass) PhoneNumberCheckingResultWithRangePhoneNumber(range_ IRange, phoneNumber string) TextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("phoneNumberCheckingResultWithRange:phoneNumber:"), range_, objc.String(phoneNumber))
	return rv
}


// Creates and returns a type checking result with the specified regular expression data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/regularExpressionCheckingResult(ranges:count:regularExpression:)

func (tc _TextCheckingResultClass) RegularExpressionCheckingResultWithRangesCountRegularExpression(ranges IRangePointer, count uint, regularExpression IRegularExpression) TextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("regularExpressionCheckingResultWithRanges:count:regularExpression:"), ranges, count, regularExpression)
	return rv
}


// Creates and returns a text checking result with the range of a misspelled word.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/spellCheckingResult(range:)

func (tc _TextCheckingResultClass) SpellCheckingResultWithRange(range_ IRange) TextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("spellCheckingResultWithRange:"), range_)
	return rv
}

// The date component of a type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/date
func (t_ TextCheckingResult) Date() NSDate {
	rv := objc.Send[NSDate](t_.ID, objc.Sel("date"))
	return rv
}

// The duration component of a type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/duration
func (t_ TextCheckingResult) Duration() TimeInterval {
	rv := objc.Send[TimeInterval](t_.ID, objc.Sel("duration"))
	return rv
}

// The phone number of a type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/phoneNumber
func (t_ TextCheckingResult) PhoneNumber() string {
	rv := objc.Send[string](t_.ID, objc.Sel("phoneNumber"))
	return rv
}

// Returns the range of the result that the receiver represents.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/range
func (t_ TextCheckingResult) Range() Range {
	rv := objc.Send[Range](t_.ID, objc.Sel("range"))
	return rv
}

// The regular expression of a type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/regularExpression
func (t_ TextCheckingResult) RegularExpression() NSRegularExpression {
	rv := objc.Send[NSRegularExpression](t_.ID, objc.Sel("regularExpression"))
	return rv
}

// Returns the text checking result type that the receiver represents.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/resultType
func (t_ TextCheckingResult) ResultType() TextCheckingType {
	rv := objc.Send[TextCheckingType](t_.ID, objc.Sel("resultType"))
	return rv
}

// A value indicating that a requested item couldn’t be found or doesn’t exist.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotfound-4qp9h
func (t_ TextCheckingResult) NSNotFound() int {
	rv := objc.Send[int](t_.ID, objc.Sel("NSNotFound"))
	return rv
}

// The address dictionary of a type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/addresscomponents
func (t_ TextCheckingResult) AddressComponents() string {
	rv := objc.Send[string](t_.ID, objc.Sel("addressComponents"))
	return rv
}


// SetAddressComponents sets the value of the addressComponents property.
// The address dictionary of a type checking result.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/addresscomponents
func (t_ TextCheckingResult) SetAddressComponents(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAddressComponents:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/alternativestrings
func (t_ TextCheckingResult) AlternativeStrings() string {
	rv := objc.Send[string](t_.ID, objc.Sel("alternativeStrings"))
	return rv
}


// SetAlternativeStrings sets the value of the alternativeStrings property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/alternativestrings
func (t_ TextCheckingResult) SetAlternativeStrings(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlternativeStrings:"), objc.String(value))
}

// A dictionary containing the components of a type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/components
func (t_ TextCheckingResult) Components() string {
	rv := objc.Send[string](t_.ID, objc.Sel("components"))
	return rv
}


// SetComponents sets the value of the components property.
// A dictionary containing the components of a type checking result.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/components
func (t_ TextCheckingResult) SetComponents(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setComponents:"), objc.String(value))
}

// The details of a located grammatical type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/grammardetails
func (t_ TextCheckingResult) GrammarDetails() string {
	rv := objc.Send[string](t_.ID, objc.Sel("grammarDetails"))
	return rv
}


// SetGrammarDetails sets the value of the grammarDetails property.
// The details of a located grammatical type checking result.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/grammardetails
func (t_ TextCheckingResult) SetGrammarDetails(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGrammarDetails:"), objc.String(value))
}

// Returns the number of ranges.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/numberofranges
func (t_ TextCheckingResult) NumberOfRanges() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfRanges"))
	return rv
}


// SetNumberOfRanges sets the value of the numberOfRanges property.
// Returns the number of ranges.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/numberofranges
func (t_ TextCheckingResult) SetNumberOfRanges(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNumberOfRanges:"), value)
}

// The detected orthography of a type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/orthography
func (t_ TextCheckingResult) Orthography() NSOrthography {
	rv := objc.Send[NSOrthography](t_.ID, objc.Sel("orthography"))
	return rv
}


// SetOrthography sets the value of the orthography property.
// The detected orthography of a type checking result.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/orthography
func (t_ TextCheckingResult) SetOrthography(value IOrthography) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setOrthography:"), value)
}

// A replacement string from one of a number of replacement checking results.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/replacementstring
func (t_ TextCheckingResult) ReplacementString() string {
	rv := objc.Send[string](t_.ID, objc.Sel("replacementString"))
	return rv
}


// SetReplacementString sets the value of the replacementString property.
// A replacement string from one of a number of replacement checking results.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/replacementstring
func (t_ TextCheckingResult) SetReplacementString(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setReplacementString:"), objc.String(value))
}

// The time zone component of a type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/timezone
func (t_ TextCheckingResult) TimeZone() TimeZone {
	rv := objc.Send[TimeZone](t_.ID, objc.Sel("timeZone"))
	return rv
}


// SetTimeZone sets the value of the timeZone property.
// The time zone component of a type checking result.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/timezone
func (t_ TextCheckingResult) SetTimeZone(value ITimeZone) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTimeZone:"), value)
}

// The URL of a type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/url
func (t_ TextCheckingResult) Url() URL {
	rv := objc.Send[URL](t_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The URL of a type checking result.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/url
func (t_ TextCheckingResult) SetUrl(value IURL) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUrl:"), value)
}



