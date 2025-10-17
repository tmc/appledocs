// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Formatter] class.
var FormatterClass objc.Class

func init() {
	FormatterClass = objc.GetClass("NSFormatter")
}

type Formatter struct {
	objc.ID
}

func FormatterFrom(ptr unsafe.Pointer) Formatter {
	return Formatter{
		ID: objc.ID(ptr),
	}
}


// The default implementation of this method raises an exception. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Formatter/string(for:)
func (f_ Formatter) StringForObjectValue(obj objc.ID) unsafe.Pointer {
	sel := objc.RegisterName("stringForObjectValue:")
	ret := f_.ID.Send(sel, obj)
	return unsafe.Pointer(ret)
}

