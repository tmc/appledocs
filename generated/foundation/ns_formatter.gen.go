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
	IsPartialStringValidNewEditingStringErrorDescription(partialString string /* primitive/slice/pointer */, newString string /* primitive/slice/pointer */, error_ string /* primitive/slice/pointer */) bool /* primitive/slice/pointer */
	StringForObjectValue(obj objectivec.IObject) String /* not a class type */
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



// Returns a Boolean value that indicates whether a partial string is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/isPartialStringValid(_:newEditingString:errorDescription:)
func (f_ Formatter) IsPartialStringValidNewEditingStringErrorDescription(partialString string /* primitive/slice/pointer */, newString string /* primitive/slice/pointer */, error_ string /* primitive/slice/pointer */) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isPartialStringValid:newEditingString:errorDescription:"), objc.String(partialString), objc.String(newString), objc.String(error_))
	return rv
}


// The default implementation of this method raises an exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/string(for:)
func (f_ Formatter) StringForObjectValue(obj objectivec.IObject) String /* not a class type */ {
	rv := objc.Send[String](f_.ID, objc.Sel("stringForObjectValue:"), obj)
	return rv
}



