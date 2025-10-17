// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Coder] class.
var coderClass = _CoderClass{objc.GetClass("NSCoder")}

type _CoderClass struct {
	class objc.Class
}

// An abstract class that serves as the basis for objects that enable archiving and distribution of other objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder

type Coder struct {
	objectivec.Object
}

// CoderFrom constructs a [Coder] from an unsafe.Pointer.
//
// An abstract class that serves as the basis for objects that enable archiving and distribution of other objects.
func CoderFrom(ptr unsafe.Pointer) Coder {
	return Coder{objectivec.Object{objc.ID(ptr)}}
}

// Decodes an object for the key, restricted to the specified class. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeObjectOfClass:forKey:
func (c_ Coder) DecodeObjectOfClassForKey(aClass objc.Class, key string) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("decodeObjectOfClass:forKey:"), aClass, key)
	return rv
}


