// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextCheckingResult */


/* debug [class_header]: Header for NSTextCheckingResult */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextCheckingResult */
// An interface definition for the [TextCheckingResult] class.
type ITextCheckingResult interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextCheckingResult */
	// properties:
	AddressComponents() IDictionary
	AlternativeStrings() []string
	Components() IDictionary
	Date() IDate
	Duration() float64
	GrammarDetails() IDictionary
	NumberOfRanges() uint
	Orthography() IOrthography
	PhoneNumber() IString
	Range() objc.IObject /* cross-framework: Range */
	RegularExpression() IRegularExpression
	ReplacementString() IString
	ResultType() TextCheckingType
	TimeZone() ITimeZone
	URL() IURL
	NSNotFound() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextCheckingResult */
	// methods:
	ResultByAdjustingRangesWithOffset(offset int) ITextCheckingResult
	RangeAtIndex(idx uint) objc.IObject /* cross-framework: Range */
	RangeWithName(name IString) objc.IObject /* cross-framework: Range */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextCheckingResult */
// Alloc allocates a new instance without initialization.
func (tc _TextCheckingResultClass) Alloc() TextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextCheckingResult */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextCheckingResult *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextCheckingResult */

// Creates and returns a text checking result with the specified address components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/addressCheckingResult(range:components:)
func (tc _TextCheckingResultClass) AddressCheckingResultWithRangeComponents(range_ objc.IObject /* cross-framework: Range */, components IDictionary) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("addressCheckingResultWithRange:components:"), range_, components)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AddressCheckingResultWithRangeComponents) */


// Creates and returns a text checking result after detecting a possible correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/correctionCheckingResult(range:replacementString:)
func (tc _TextCheckingResultClass) CorrectionCheckingResultWithRangeReplacementString(range_ objc.IObject /* cross-framework: Range */, replacementString IString) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("correctionCheckingResultWithRange:replacementString:"), range_, replacementString)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CorrectionCheckingResultWithRangeReplacementString) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/correctionCheckingResult(range:replacementString:alternativeStrings:)
func (tc _TextCheckingResultClass) CorrectionCheckingResultWithRangeReplacementStringAlternativeStrings(range_ objc.IObject /* cross-framework: Range */, replacementString IString, alternativeStrings []string) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("correctionCheckingResultWithRange:replacementString:alternativeStrings:"), range_, replacementString, alternativeStrings)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CorrectionCheckingResultWithRangeReplacementStringAlternativeStrings) */


// Creates and returns a text checking result with the specified dash corrected replacement string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/dashCheckingResult(range:replacementString:)
func (tc _TextCheckingResultClass) DashCheckingResultWithRangeReplacementString(range_ objc.IObject /* cross-framework: Range */, replacementString IString) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("dashCheckingResultWithRange:replacementString:"), range_, replacementString)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DashCheckingResultWithRangeReplacementString) */


// Creates and returns a text checking result with the specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/dateCheckingResult(range:date:)
func (tc _TextCheckingResultClass) DateCheckingResultWithRangeDate(range_ objc.IObject /* cross-framework: Range */, date IDate) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("dateCheckingResultWithRange:date:"), range_, date)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DateCheckingResultWithRangeDate) */


// Creates and returns a text checking result with the specified date, time zone, and duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/dateCheckingResult(range:date:timeZone:duration:)
func (tc _TextCheckingResultClass) DateCheckingResultWithRangeDateTimeZoneDuration(range_ objc.IObject /* cross-framework: Range */, date IDate, timeZone ITimeZone, duration float64) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("dateCheckingResultWithRange:date:timeZone:duration:"), range_, date, timeZone, duration)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DateCheckingResultWithRangeDateTimeZoneDuration) */


// Creates and returns a text checking result with the specified array of grammatical errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/grammarCheckingResult(range:details:)
func (tc _TextCheckingResultClass) GrammarCheckingResultWithRangeDetails(range_ objc.IObject /* cross-framework: Range */, details IDictionary) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("grammarCheckingResultWithRange:details:"), range_, details)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GrammarCheckingResultWithRangeDetails) */


// Creates and returns a text checking result with the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/linkCheckingResult(range:url:)
func (tc _TextCheckingResultClass) LinkCheckingResultWithRangeURL(range_ objc.IObject /* cross-framework: Range */, url IURL) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("linkCheckingResultWithRange:URL:"), range_, url)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LinkCheckingResultWithRangeURL) */


// Creates and returns a text checking result with the specified orthography.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/orthographyCheckingResult(range:orthography:)
func (tc _TextCheckingResultClass) OrthographyCheckingResultWithRangeOrthography(range_ objc.IObject /* cross-framework: Range */, orthography IOrthography) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("orthographyCheckingResultWithRange:orthography:"), range_, orthography)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OrthographyCheckingResultWithRangeOrthography) */


// Creates and returns a text checking result with the specified phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/phoneNumberCheckingResult(range:phoneNumber:)
func (tc _TextCheckingResultClass) PhoneNumberCheckingResultWithRangePhoneNumber(range_ objc.IObject /* cross-framework: Range */, phoneNumber IString) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("phoneNumberCheckingResultWithRange:phoneNumber:"), range_, phoneNumber)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PhoneNumberCheckingResultWithRangePhoneNumber) */


// Creates and returns a text checking result with the specified quote-balanced replacement string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/quoteCheckingResult(range:replacementString:)
func (tc _TextCheckingResultClass) QuoteCheckingResultWithRangeReplacementString(range_ objc.IObject /* cross-framework: Range */, replacementString IString) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("quoteCheckingResultWithRange:replacementString:"), range_, replacementString)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=QuoteCheckingResultWithRangeReplacementString) */


// Creates and returns a type checking result with the specified regular expression data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/regularExpressionCheckingResult(ranges:count:regularExpression:)
func (tc _TextCheckingResultClass) RegularExpressionCheckingResultWithRangesCountRegularExpression(ranges RangePointer /* typedef */, count uint, regularExpression IRegularExpression) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("regularExpressionCheckingResultWithRanges:count:regularExpression:"), ranges, count, regularExpression)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RegularExpressionCheckingResultWithRangesCountRegularExpression) */


// Creates and returns a text checking result with the specified replacement string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/replacementCheckingResult(range:replacementString:)
func (tc _TextCheckingResultClass) ReplacementCheckingResultWithRangeReplacementString(range_ objc.IObject /* cross-framework: Range */, replacementString IString) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("replacementCheckingResultWithRange:replacementString:"), range_, replacementString)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReplacementCheckingResultWithRangeReplacementString) */


// Creates and returns a text checking result with the range of a misspelled word.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/spellCheckingResult(range:)
func (tc _TextCheckingResultClass) SpellCheckingResultWithRange(range_ objc.IObject /* cross-framework: Range */) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("spellCheckingResultWithRange:"), range_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SpellCheckingResultWithRange) */


// Creates and returns a text checking result with the specified transit information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/transitInformationCheckingResult(range:components:)
func (tc _TextCheckingResultClass) TransitInformationCheckingResultWithRangeComponents(range_ objc.IObject /* cross-framework: Range */, components IDictionary) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](objc.ID(tc.class), objc.Sel("transitInformationCheckingResultWithRange:components:"), range_, components)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TransitInformationCheckingResultWithRangeComponents) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextCheckingResult */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextCheckingResult */

// Returns a new text checking result after adjusting the ranges as specified by the offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/adjustingRanges(offset:)
func (t_ TextCheckingResult) ResultByAdjustingRangesWithOffset(offset int) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](t_.ID, objc.Sel("resultByAdjustingRangesWithOffset:"), offset)
	return rv
}/* debug [instance_methods/method]: ResultByAdjustingRangesWithOffset */


// Returns the result type that the range represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/range(at:)
func (t_ TextCheckingResult) RangeAtIndex(idx uint) objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("rangeAtIndex:"), idx)
	return rv
}/* debug [instance_methods/method]: RangeAtIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/range(withName:)
func (t_ TextCheckingResult) RangeWithName(name IString) objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("rangeWithName:"), name)
	return rv
}/* debug [instance_methods/method]: RangeWithName */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextCheckingResult */

// The address dictionary of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/addressComponents
func (t_ TextCheckingResult) AddressComponents() IDictionary {
	rv := objc.Send[Dictionary](t_.ID, objc.Sel("addressComponents"))
	return rv
}/* debug [instance_properties/getter]: addressComponents */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/alternativeStrings
func (t_ TextCheckingResult) AlternativeStrings() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("alternativeStrings"))
	return rv
}/* debug [instance_properties/getter]: alternativeStrings */


// A dictionary containing the components of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/components
func (t_ TextCheckingResult) Components() IDictionary {
	rv := objc.Send[Dictionary](t_.ID, objc.Sel("components"))
	return rv
}/* debug [instance_properties/getter]: components */


// The date component of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/date
func (t_ TextCheckingResult) Date() IDate {
	rv := objc.Send[Date](t_.ID, objc.Sel("date"))
	return rv
}/* debug [instance_properties/getter]: date */


// The duration component of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/duration
func (t_ TextCheckingResult) Duration() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The details of a located grammatical type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/grammarDetails
func (t_ TextCheckingResult) GrammarDetails() IDictionary {
	rv := objc.Send[Dictionary](t_.ID, objc.Sel("grammarDetails"))
	return rv
}/* debug [instance_properties/getter]: grammarDetails */


// Returns the number of ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/numberOfRanges
func (t_ TextCheckingResult) NumberOfRanges() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("numberOfRanges"))
	return rv
}/* debug [instance_properties/getter]: numberOfRanges */


// The detected orthography of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/orthography
func (t_ TextCheckingResult) Orthography() IOrthography {
	rv := objc.Send[Orthography](t_.ID, objc.Sel("orthography"))
	return rv
}/* debug [instance_properties/getter]: orthography */


// The phone number of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/phoneNumber
func (t_ TextCheckingResult) PhoneNumber() IString {
	rv := objc.Send[String](t_.ID, objc.Sel("phoneNumber"))
	return rv
}/* debug [instance_properties/getter]: phoneNumber */


// Returns the range of the result that the receiver represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/range
func (t_ TextCheckingResult) Range() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("range"))
	return rv
}/* debug [instance_properties/getter]: range */


// The regular expression of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/regularExpression
func (t_ TextCheckingResult) RegularExpression() IRegularExpression {
	rv := objc.Send[RegularExpression](t_.ID, objc.Sel("regularExpression"))
	return rv
}/* debug [instance_properties/getter]: regularExpression */


// A replacement string from one of a number of replacement checking results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/replacementString
func (t_ TextCheckingResult) ReplacementString() IString {
	rv := objc.Send[String](t_.ID, objc.Sel("replacementString"))
	return rv
}/* debug [instance_properties/getter]: replacementString */


// Returns the text checking result type that the receiver represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/resultType
func (t_ TextCheckingResult) ResultType() TextCheckingType {
	rv := objc.Send[TextCheckingType](t_.ID, objc.Sel("resultType"))
	return rv
}/* debug [instance_properties/getter]: resultType */


// The time zone component of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/timeZone
func (t_ TextCheckingResult) TimeZone() ITimeZone {
	rv := objc.Send[TimeZone](t_.ID, objc.Sel("timeZone"))
	return rv
}/* debug [instance_properties/getter]: timeZone */


// The URL of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/url
func (t_ TextCheckingResult) URL() IURL {
	rv := objc.Send[URL](t_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// A value indicating that a requested item couldn’t be found or doesn’t exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotfound-4qp9h
func (t_ TextCheckingResult) NSNotFound() int {
	rv := objc.Send[int](t_.ID, objc.Sel("NSNotFound"))
	return rv
}/* debug [instance_properties/getter]: NSNotFound */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextCheckingResult */



