// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Formatter] class.
var (
	FormatterClass     _FormatterClass
	FormatterClassOnce sync.Once
)

func getFormatterClass() _FormatterClass {
	FormatterClassOnce.Do(func() {
		FormatterClass = _FormatterClass{objc.GetClass("NSFormatter")}
	})
	return FormatterClass
}

type _FormatterClass struct {
	class objc.Class
}

// An interface definition for the [Formatter] class.
type IFormatter interface {
	objectivec.IObject
	// properties:
	// methods:
	AttributedStringForObjectValueWithDefaultAttributes(obj objectivec.IObject, attrs IDictionary /* already interface */) IAttributedString
	EditingStringForObjectValue(obj objectivec.IObject) IString
	GetObjectValueForStringErrorDescription(obj unsafe.Pointer, string_ IString, error_ IString) bool /* primitive/slice/pointer. */
	IsPartialStringValidNewEditingStringErrorDescription(partialString IString, newString IString, error_ IString) bool /* primitive/slice/pointer. */
	IsPartialStringValidProposedSelectedRangeOriginalStringOriginalSelectedRangeErrorDescription(partialStringPtr IString, proposedSelRangePtr objc.IObject /* cross-framework RangePointer */, origString IString, origSelRange objc.IObject /* cross-framework Range */, error_ IString) bool /* primitive/slice/pointer. */
	StringForObjectValue(obj objectivec.IObject) IString
}

// An abstract class that declares an interface for objects that create, interpret, and validate the textual representation of values.
//
// The Foundation framework provides several concrete subclasses of , including , , , , , , and .


// An abstract class that declares an interface for objects that create, interpret, and validate the textual representation of values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter
type Formatter struct {
	objectivec.Object
}

// FormatterFrom constructs a [Formatter] from an unsafe.Pointer.
//
// An abstract class that declares an interface for objects that create, interpret, and validate the textual representation of values.
func FormatterFrom(ptr unsafe.Pointer) Formatter {
	return Formatter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FormatterClass) Alloc() Formatter {
	rv := objc.Send[Formatter](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FormatterClass) New() Formatter {
	rv := objc.Send[Formatter](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ Formatter) Init() Formatter {
	rv := objc.Send[Formatter](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ Formatter) Autorelease() Formatter {
	rv := objc.Send[Formatter](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFormatter creates a new Formatter instance.
func NewFormatter() Formatter {
	return getFormatterClass().New()
}



// The default implementation returns to indicate that the formatter object does not provide an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/attributedString(for:withDefaultAttributes:)
func (f_ Formatter) AttributedStringForObjectValueWithDefaultAttributes(obj objectivec.IObject, attrs IDictionary /* already interface */) IAttributedString {
	rv := objc.Send[AttributedString](f_.ID, objc.Sel("attributedStringForObjectValue:withDefaultAttributes:"), obj, attrs)
	return rv
}


// The default implementation of this method invokes .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/editingString(for:)
func (f_ Formatter) EditingStringForObjectValue(obj objectivec.IObject) IString {
	rv := objc.Send[String](f_.ID, objc.Sel("editingStringForObjectValue:"), obj)
	return rv
}


// The default implementation of this method raises an exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/getObjectValue(_:for:errorDescription:)
func (f_ Formatter) GetObjectValueForStringErrorDescription(obj unsafe.Pointer, string_ IString, error_ IString) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("getObjectValue:forString:errorDescription:"), obj, string_, error_)
	return rv
}


// Returns a Boolean value that indicates whether a partial string is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/isPartialStringValid(_:newEditingString:errorDescription:)
func (f_ Formatter) IsPartialStringValidNewEditingStringErrorDescription(partialString IString, newString IString, error_ IString) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isPartialStringValid:newEditingString:errorDescription:"), partialString, newString, error_)
	return rv
}


// This method should be implemented in subclasses that want to validate user changes to a string in a field, where the user changes are not necessarily at the end of the string, and preserve the selection (or set a different one, such as selecting the erroneous part of the string the user has typed).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/isPartialStringValid(_:proposedSelectedRange:originalString:originalSelectedRange:errorDescription:)
func (f_ Formatter) IsPartialStringValidProposedSelectedRangeOriginalStringOriginalSelectedRangeErrorDescription(partialStringPtr IString, proposedSelRangePtr objc.IObject /* cross-framework RangePointer */, origString IString, origSelRange objc.IObject /* cross-framework Range */, error_ IString) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isPartialStringValid:proposedSelectedRange:originalString:originalSelectedRange:errorDescription:"), partialStringPtr, proposedSelRangePtr, origString, origSelRange, error_)
	return rv
}


// The default implementation of this method raises an exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/string(for:)
func (f_ Formatter) StringForObjectValue(obj objectivec.IObject) IString {
	rv := objc.Send[String](f_.ID, objc.Sel("stringForObjectValue:"), obj)
	return rv
}



