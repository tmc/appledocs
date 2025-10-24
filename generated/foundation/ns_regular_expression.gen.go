// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [RegularExpression] class.
var (
	RegularExpressionClass     _RegularExpressionClass
	RegularExpressionClassOnce sync.Once
)

func getRegularExpressionClass() _RegularExpressionClass {
	RegularExpressionClassOnce.Do(func() {
		RegularExpressionClass = _RegularExpressionClass{objc.GetClass("NSRegularExpression")}
	})
	return RegularExpressionClass
}

type _RegularExpressionClass struct {
	class objc.Class
}





// An interface definition for the [RegularExpression] class.
type IRegularExpression interface {
	objectivec.IObject
	

	// properties:
	NumberOfCaptureGroups() uint
	Options() RegularExpressionOptions
	Pattern() IString
	NSNotFound() int
	Range() objc.IObject /* cross-framework: Range */
	SetRange(value objc.IObject /* cross-framework: Range */)


	

	// methods:
	EnumerateMatchesInStringOptionsRangeUsingBlock(string_ IString, options MatchingOptions, range_ objc.IObject /* cross-framework: Range */, block unsafe.Pointer)
	FirstMatchInStringOptionsRange(string_ IString, options MatchingOptions, range_ objc.IObject /* cross-framework: Range */) ITextCheckingResult
	MatchesInStringOptionsRange(string_ IString, options MatchingOptions, range_ objc.IObject /* cross-framework: Range */) []TextCheckingResult
	NumberOfMatchesInStringOptionsRange(string_ IString, options MatchingOptions, range_ objc.IObject /* cross-framework: Range */) uint
	RangeOfFirstMatchInStringOptionsRange(string_ IString, options MatchingOptions, range_ objc.IObject /* cross-framework: Range */) objc.IObject /* cross-framework: Range */
	ReplaceMatchesInStringOptionsRangeWithTemplate(string_ IMutableString, options MatchingOptions, range_ objc.IObject /* cross-framework: Range */, templ IString) uint
	ReplacementStringForResultInStringOffsetTemplate(result ITextCheckingResult, string_ IString, offset int, templ IString) IString
	StringByReplacingMatchesInStringOptionsRangeWithTemplate(string_ IString, options MatchingOptions, range_ objc.IObject /* cross-framework: Range */, templ IString) IString


}





// Alloc allocates a new instance without initialization.
func (rc _RegularExpressionClass) Alloc() RegularExpression {
	rv := objc.Send[RegularExpression](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RegularExpressionClass) New() RegularExpression {
	rv := objc.Send[RegularExpression](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RegularExpression) Init() RegularExpression {
	rv := objc.Send[RegularExpression](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RegularExpression) Autorelease() RegularExpression {
	rv := objc.Send[RegularExpression](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRegularExpression creates a new RegularExpression instance.
func NewRegularExpression() RegularExpression {
	return getRegularExpressionClass().New()
}





// An immutable representation of a compiled regular expression that you apply to Unicode strings.
//
// The fundamental matching method for is a Block iterator method that allows clients to supply a Block object which will be invoked each time the regular expression matches a portion of the target string. There are additional convenience methods for returning all the matches as an array, the total number of matches, the first match, and the range of the first match. An individual match is represented by an instance of the class, which carries information about the overall matched range (via its property), and the range of each individual capture group (via the method). For basic objects, these match results will be of type , but subclasses may use other types.


// An immutable representation of a compiled regular expression that you apply to Unicode strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression
type RegularExpression struct {
	objectivec.Object
}

// RegularExpressionFrom constructs a [RegularExpression] from an unsafe.Pointer.
//
// An immutable representation of a compiled regular expression that you apply to Unicode strings.
func RegularExpressionFrom(ptr unsafe.Pointer) RegularExpression {
	return RegularExpression{objectivec.Object{objc.ID(ptr)}}
}






// Returns an initialized NSRegularExpression instance with the specified regular expression pattern and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/init(pattern:options:)
func NewRegularExpressionWithPatternOptionsError(pattern IString, options RegularExpressionOptions, error_ IError) RegularExpression {
	instance := getRegularExpressionClass().Alloc()
	rv := objc.Send[RegularExpression](instance.ID, objc.Sel("initWithPattern:options:error:"), pattern, options, error_)
	rv.Autorelease()
	return rv
}







// Returns a string by adding backslash escapes as necessary to protect any characters that would match as pattern metacharacters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/escapedPattern(for:)
func (rc _RegularExpressionClass) EscapedPatternForString(string_ IString) IString {
	rv := objc.Send[String](objc.ID(rc.class), objc.Sel("escapedPatternForString:"), string_)
	return rv
}


// Returns a template string by adding backslash escapes as necessary to protect any characters that would match as pattern metacharacters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/escapedTemplate(for:)
func (rc _RegularExpressionClass) EscapedTemplateForString(string_ IString) IString {
	rv := objc.Send[String](objc.ID(rc.class), objc.Sel("escapedTemplateForString:"), string_)
	return rv
}


// Creates an NSRegularExpression instance with the specified regular expression pattern and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/regularExpressionWithPattern:options:error:
func (rc _RegularExpressionClass) RegularExpressionWithPatternOptionsError(pattern IString, options RegularExpressionOptions, error_ IError) IRegularExpression {
	rv := objc.Send[RegularExpression](objc.ID(rc.class), objc.Sel("regularExpressionWithPattern:options:error:"), pattern, options, error_)
	return rv
}












// Enumerates the string allowing the Block to handle each regular expression match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/enumerateMatches(in:options:range:using:)
func (r_ RegularExpression) EnumerateMatchesInStringOptionsRangeUsingBlock(string_ IString, options MatchingOptions, range_ objc.IObject /* cross-framework: Range */, block unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("enumerateMatchesInString:options:range:usingBlock:"), string_, options, range_, block)
}


// Returns the first match of the regular expression within the specified range of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/firstMatch(in:options:range:)
func (r_ RegularExpression) FirstMatchInStringOptionsRange(string_ IString, options MatchingOptions, range_ objc.IObject /* cross-framework: Range */) ITextCheckingResult {
	rv := objc.Send[TextCheckingResult](r_.ID, objc.Sel("firstMatchInString:options:range:"), string_, options, range_)
	return rv
}


// Returns an array containing all the matches of the regular expression in the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/matches(in:options:range:)
func (r_ RegularExpression) MatchesInStringOptionsRange(string_ IString, options MatchingOptions, range_ objc.IObject /* cross-framework: Range */) []TextCheckingResult {
	rv := objc.Send[[]TextCheckingResult](r_.ID, objc.Sel("matchesInString:options:range:"), string_, options, range_)
	return rv
}


// Returns the number of matches of the regular expression within the specified range of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/numberOfMatches(in:options:range:)
func (r_ RegularExpression) NumberOfMatchesInStringOptionsRange(string_ IString, options MatchingOptions, range_ objc.IObject /* cross-framework: Range */) uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("numberOfMatchesInString:options:range:"), string_, options, range_)
	return rv
}


// Returns the range of the first match of the regular expression within the specified range of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/rangeOfFirstMatch(in:options:range:)
func (r_ RegularExpression) RangeOfFirstMatchInStringOptionsRange(string_ IString, options MatchingOptions, range_ objc.IObject /* cross-framework: Range */) objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("rangeOfFirstMatchInString:options:range:"), string_, options, range_)
	return rv
}


// Replaces regular expression matches within the mutable string using the template string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/replaceMatches(in:options:range:withTemplate:)
func (r_ RegularExpression) ReplaceMatchesInStringOptionsRangeWithTemplate(string_ IMutableString, options MatchingOptions, range_ objc.IObject /* cross-framework: Range */, templ IString) uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("replaceMatchesInString:options:range:withTemplate:"), string_, options, range_, templ)
	return rv
}


// Used to perform template substitution for a single result for clients implementing their own replace functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/replacementString(for:in:offset:template:)
func (r_ RegularExpression) ReplacementStringForResultInStringOffsetTemplate(result ITextCheckingResult, string_ IString, offset int, templ IString) IString {
	rv := objc.Send[String](r_.ID, objc.Sel("replacementStringForResult:inString:offset:template:"), result, string_, offset, templ)
	return rv
}


// Returns a new string containing matching regular expressions replaced with the template string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/stringByReplacingMatches(in:options:range:withTemplate:)
func (r_ RegularExpression) StringByReplacingMatchesInStringOptionsRangeWithTemplate(string_ IString, options MatchingOptions, range_ objc.IObject /* cross-framework: Range */, templ IString) IString {
	rv := objc.Send[String](r_.ID, objc.Sel("stringByReplacingMatchesInString:options:range:withTemplate:"), string_, options, range_, templ)
	return rv
}







// Returns the number of capture groups in the regular expression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/numberOfCaptureGroups
func (r_ RegularExpression) NumberOfCaptureGroups() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("numberOfCaptureGroups"))
	return rv
}


// Returns the options used when the regular expression option was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/options-swift.property
func (r_ RegularExpression) Options() RegularExpressionOptions {
	rv := objc.Send[RegularExpressionOptions](r_.ID, objc.Sel("options"))
	return rv
}


// Returns the regular expression pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/pattern
func (r_ RegularExpression) Pattern() IString {
	rv := objc.Send[String](r_.ID, objc.Sel("pattern"))
	return rv
}


// A value indicating that a requested item couldn’t be found or doesn’t exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotfound-4qp9h
func (r_ RegularExpression) NSNotFound() int {
	rv := objc.Send[int](r_.ID, objc.Sel("NSNotFound"))
	return rv
}


// Returns the range of the result that the receiver represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/range
func (r_ RegularExpression) Range() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("range"))
	return rv
}


// Returns the range of the result that the receiver represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/range
func (r_ RegularExpression) SetRange(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRange:"), value)
}







