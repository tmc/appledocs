// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFormatter */


/* debug [class_header]: Header for NSFormatter */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Formatter */
// An interface definition for the [Formatter] class.
type IFormatter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Formatter */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Formatter */
	// methods:
	AttributedStringForObjectValueWithDefaultAttributes(obj objc.IObject, attrs IDictionary) IAttributedString
	EditingStringForObjectValue(obj objc.IObject) IString
	GetObjectValueForStringErrorDescription(obj objectivec.IObject, string_ IString, error_ IString) bool
	IsPartialStringValidNewEditingStringErrorDescription(partialString IString, newString IString, error_ IString) bool
	IsPartialStringValidProposedSelectedRangeOriginalStringOriginalSelectedRangeErrorDescription(partialStringPtr IString, proposedSelRangePtr RangePointer, origString IString, origSelRange objc.IObject /* cross-framework: Range */, error_ IString) bool
	StringForObjectValue(obj objc.IObject) IString
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Formatter */
// Alloc allocates a new instance without initialization.
func (fc _FormatterClass) Alloc() Formatter {
	rv := objc.Send[Formatter](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Formatter */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Formatter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Formatter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Formatter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Formatter */

// The default implementation returns to indicate that the formatter object does not provide an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/attributedString(for:withDefaultAttributes:)
func (f_ Formatter) AttributedStringForObjectValueWithDefaultAttributes(obj objc.IObject, attrs IDictionary) IAttributedString {
	rv := objc.Send[AttributedString](f_.ID, objc.Sel("attributedStringForObjectValue:withDefaultAttributes:"), obj, attrs)
	return rv
}/* debug [instance_methods/method]: AttributedStringForObjectValueWithDefaultAttributes */


// The default implementation of this method invokes .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/editingString(for:)
func (f_ Formatter) EditingStringForObjectValue(obj objc.IObject) IString {
	rv := objc.Send[String](f_.ID, objc.Sel("editingStringForObjectValue:"), obj)
	return rv
}/* debug [instance_methods/method]: EditingStringForObjectValue */


// The default implementation of this method raises an exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/getObjectValue(_:for:errorDescription:)
func (f_ Formatter) GetObjectValueForStringErrorDescription(obj objectivec.IObject, string_ IString, error_ IString) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("getObjectValue:forString:errorDescription:"), obj, string_, error_)
	return rv
}/* debug [instance_methods/method]: GetObjectValueForStringErrorDescription */


// Returns a Boolean value that indicates whether a partial string is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/isPartialStringValid(_:newEditingString:errorDescription:)
func (f_ Formatter) IsPartialStringValidNewEditingStringErrorDescription(partialString IString, newString IString, error_ IString) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isPartialStringValid:newEditingString:errorDescription:"), partialString, newString, error_)
	return rv
}/* debug [instance_methods/method]: IsPartialStringValidNewEditingStringErrorDescription */


// This method should be implemented in subclasses that want to validate user changes to a string in a field, where the user changes are not necessarily at the end of the string, and preserve the selection (or set a different one, such as selecting the erroneous part of the string the user has typed).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/isPartialStringValid(_:proposedSelectedRange:originalString:originalSelectedRange:errorDescription:)
func (f_ Formatter) IsPartialStringValidProposedSelectedRangeOriginalStringOriginalSelectedRangeErrorDescription(partialStringPtr IString, proposedSelRangePtr RangePointer, origString IString, origSelRange objc.IObject /* cross-framework: Range */, error_ IString) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isPartialStringValid:proposedSelectedRange:originalString:originalSelectedRange:errorDescription:"), partialStringPtr, proposedSelRangePtr, origString, origSelRange, error_)
	return rv
}/* debug [instance_methods/method]: IsPartialStringValidProposedSelectedRangeOriginalStringOriginalSelectedRangeErrorDescription */


// The default implementation of this method raises an exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/string(for:)
func (f_ Formatter) StringForObjectValue(obj objc.IObject) IString {
	rv := objc.Send[String](f_.ID, objc.Sel("stringForObjectValue:"), obj)
	return rv
}/* debug [instance_methods/method]: StringForObjectValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Formatter */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFormatter */



