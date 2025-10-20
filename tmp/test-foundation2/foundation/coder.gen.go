// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var coderClass _CoderClass

func init() {
	coderClass = _CoderClass{objc.GetClass("NSCoder")}
}

type _CoderClass struct {
	class objc.Class
}

type Coder struct {
	objc.ID
}

func CoderFrom(ptr unsafe.Pointer) Coder {
	return Coder{
		ID: objc.ID(ptr),
	}
}


// Decodes an object for the key, restricted to the specified class. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObjectOfClass:forKey:
func (c_ Coder) DecodeObjectOfClassForKey(aClass objc.Class, key string) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObjectOfClass:forKey:"), aClass, key)
	return rv
}


