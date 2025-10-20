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
}

// An occurrence of textual content found during the analysis of a block of text, such as when matching a regular expression.
//
// On both iOS and macOS, instances of are returned by the class and the class to indicate the discovery of content. In those cases, what is found may be a match for a regular expression or a date, address, phone number, and so on. In macOS, instances of are returned by the object to describe the results of spelling, grammar, or text-substitution actions.
//
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/addressCheckingResult(range:components:)
func (tc _TextCheckingResultClass) AddressCheckingResultWithRangeComponents(range_ Range, components unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("addressCheckingResultWithRange:components:"), range_, components)
	return rv
}

// Creates and returns a text checking result with the specified phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/phoneNumberCheckingResult(range:phoneNumber:)
func (tc _TextCheckingResultClass) PhoneNumberCheckingResultWithRangePhoneNumber(range_ Range, phoneNumber string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("phoneNumberCheckingResultWithRange:phoneNumber:"), range_, objc.String(phoneNumber))
	return rv
}

// Creates and returns a type checking result with the specified regular expression data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/regularExpressionCheckingResult(ranges:count:regularExpression:)
func (tc _TextCheckingResultClass) RegularExpressionCheckingResultWithRangesCountRegularExpression(ranges unsafe.Pointer, count uint, regularExpression unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("regularExpressionCheckingResultWithRanges:count:regularExpression:"), ranges, count, regularExpression)
	return rv
}

// Creates and returns a text checking result with the range of a misspelled word.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/spellCheckingResult(range:)
func (tc _TextCheckingResultClass) SpellCheckingResultWithRange(range_ Range) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("spellCheckingResultWithRange:"), range_)
	return rv
}

// The date component of a type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/date
func (t_ TextCheckingResult) Date() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("date"))
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
func (t_ TextCheckingResult) PhoneNumber() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("phoneNumber"))
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
func (t_ TextCheckingResult) RegularExpression() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("regularExpression"))
	return rv
}

// Returns the text checking result type that the receiver represents.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/resultType
func (t_ TextCheckingResult) ResultType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("resultType"))
	return rv
}



