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
	Date() IDate
	PhoneNumber() string
	NSNotFound() int
	AddressComponents() string
	SetAddressComponents(value string)
	AlternativeStrings() string
	SetAlternativeStrings(value string)
	Components() string
	SetComponents(value string)
	Duration() TimeInterval
	SetDuration(value TimeInterval)
	GrammarDetails() string
	SetGrammarDetails(value string)
	NumberOfRanges() int
	SetNumberOfRanges(value int)
	Orthography() IOrthography
	SetOrthography(value IOrthography)
	Range() Range
	SetRange(value Range)
	RegularExpression() IRegularExpression
	SetRegularExpression(value IRegularExpression)
	ReplacementString() string
	SetReplacementString(value string)
	ResultType() unsafe.Pointer
	SetResultType(value unsafe.Pointer)
	TimeZone() ITimeZone
	SetTimeZone(value ITimeZone)
	Url() IURL
	SetUrl(value IURL)
}

// An occurrence of textual content found during the analysis of a block of text, such as when matching a regular expression.
//
// On both iOS and macOS, instances of are returned by the class and the class to indicate the discovery of content. In those cases, what is found may be a match for a regular expression or a date, address, phone number, and so on. In macOS, instances of are returned by the object to describe the results of spelling, grammar, or text-substitution actions.


// An occurrence of textual content found during the analysis of a block of text, such as when matching a regular expression.
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
func (tc _TextCheckingResultClass) AddressCheckingResultWithRangeComponents(range_ Range, components IDictionary) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("addressCheckingResultWithRange:components:"), range_, components)
	return rv
}


// The date component of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/date
func (t_ TextCheckingResult) Date() IDate {
	rv := objc.Send[NSDate](t_.ID, objc.Sel("date"))
	return rv
}


// The phone number of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/phoneNumber
func (t_ TextCheckingResult) PhoneNumber() string {
	rv := objc.Send[string](t_.ID, objc.Sel("phoneNumber"))
	return rv
}


// A value indicating that a requested item couldn’t be found or doesn’t exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotfound-4qp9h
func (t_ TextCheckingResult) NSNotFound() int {
	rv := objc.Send[int](t_.ID, objc.Sel("NSNotFound"))
	return rv
}


// The address dictionary of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/addresscomponents
func (t_ TextCheckingResult) AddressComponents() string {
	rv := objc.Send[string](t_.ID, objc.Sel("addressComponents"))
	return rv
}


// The address dictionary of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/addresscomponents
func (t_ TextCheckingResult) SetAddressComponents(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAddressComponents:"), objc.String(value))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/alternativestrings
func (t_ TextCheckingResult) AlternativeStrings() string {
	rv := objc.Send[string](t_.ID, objc.Sel("alternativeStrings"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/alternativestrings
func (t_ TextCheckingResult) SetAlternativeStrings(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlternativeStrings:"), objc.String(value))
}


// A dictionary containing the components of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/components
func (t_ TextCheckingResult) Components() string {
	rv := objc.Send[string](t_.ID, objc.Sel("components"))
	return rv
}


// A dictionary containing the components of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/components
func (t_ TextCheckingResult) SetComponents(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setComponents:"), objc.String(value))
}


// The duration component of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/duration
func (t_ TextCheckingResult) Duration() TimeInterval {
	rv := objc.Send[TimeInterval](t_.ID, objc.Sel("duration"))
	return rv
}


// The duration component of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/duration
func (t_ TextCheckingResult) SetDuration(value TimeInterval) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDuration:"), value)
}


// The details of a located grammatical type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/grammardetails
func (t_ TextCheckingResult) GrammarDetails() string {
	rv := objc.Send[string](t_.ID, objc.Sel("grammarDetails"))
	return rv
}


// The details of a located grammatical type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/grammardetails
func (t_ TextCheckingResult) SetGrammarDetails(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGrammarDetails:"), objc.String(value))
}


// Returns the number of ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/numberofranges
func (t_ TextCheckingResult) NumberOfRanges() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfRanges"))
	return rv
}


// Returns the number of ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/numberofranges
func (t_ TextCheckingResult) SetNumberOfRanges(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNumberOfRanges:"), value)
}


// The detected orthography of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/orthography
func (t_ TextCheckingResult) Orthography() IOrthography {
	rv := objc.Send[NSOrthography](t_.ID, objc.Sel("orthography"))
	return rv
}


// The detected orthography of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/orthography
func (t_ TextCheckingResult) SetOrthography(value IOrthography) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setOrthography:"), value)
}


// Returns the range of the result that the receiver represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/range
func (t_ TextCheckingResult) Range() Range {
	rv := objc.Send[Range](t_.ID, objc.Sel("range"))
	return rv
}


// Returns the range of the result that the receiver represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/range
func (t_ TextCheckingResult) SetRange(value Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRange:"), value)
}


// The regular expression of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/regularexpression
func (t_ TextCheckingResult) RegularExpression() IRegularExpression {
	rv := objc.Send[NSRegularExpression](t_.ID, objc.Sel("regularExpression"))
	return rv
}


// The regular expression of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/regularexpression
func (t_ TextCheckingResult) SetRegularExpression(value IRegularExpression) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRegularExpression:"), value)
}


// A replacement string from one of a number of replacement checking results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/replacementstring
func (t_ TextCheckingResult) ReplacementString() string {
	rv := objc.Send[string](t_.ID, objc.Sel("replacementString"))
	return rv
}


// A replacement string from one of a number of replacement checking results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/replacementstring
func (t_ TextCheckingResult) SetReplacementString(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setReplacementString:"), objc.String(value))
}


// Returns the text checking result type that the receiver represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/resulttype
func (t_ TextCheckingResult) ResultType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("resultType"))
	return rv
}


// Returns the text checking result type that the receiver represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/resulttype
func (t_ TextCheckingResult) SetResultType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResultType:"), value)
}


// The time zone component of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/timezone
func (t_ TextCheckingResult) TimeZone() ITimeZone {
	rv := objc.Send[TimeZone](t_.ID, objc.Sel("timeZone"))
	return rv
}


// The time zone component of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/timezone
func (t_ TextCheckingResult) SetTimeZone(value ITimeZone) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTimeZone:"), value)
}


// The URL of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/url
func (t_ TextCheckingResult) Url() IURL {
	rv := objc.Send[URL](t_.ID, objc.Sel("url"))
	return rv
}


// The URL of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/url
func (t_ TextCheckingResult) SetUrl(value IURL) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUrl:"), value)
}



