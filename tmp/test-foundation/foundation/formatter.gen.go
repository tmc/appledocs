// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var FormatterClass _FormatterClass

func init() {
	FormatterClass = _FormatterClass{objc.GetClass("NSFormatter")}
}

type _FormatterClass struct {
	class objc.Class
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Formatter/string(for:)
func (f_ Formatter) StringForObjectValue(obj objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("stringForObjectValue:"), obj)
	return rv
}


