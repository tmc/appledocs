// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Formatter] class.
var formatterClass = _FormatterClass{objc.GetClass("NSFormatter")}

type _FormatterClass struct {
	class objc.Class
}

// An interface definition for the [Formatter] class.
type IFormatter interface {
	objectivec.IObject
	StringForObjectValue(obj objc.ID) unsafe.Pointer
}

// An abstract class that declares an interface for objects that create, interpret, and validate the textual representation of values. [Full Topic]
//
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

// The default implementation of this method raises an exception. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/string(for:)
func (f_ Formatter) StringForObjectValue(obj objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("stringForObjectValue:"), obj)
	return rv
}


