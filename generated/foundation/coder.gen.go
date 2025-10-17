// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Coder] class.
var CoderClass objc.Class

func init() {
	CoderClass = objc.GetClass("NSCoder")
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
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSCoder/decodeObjectOfClass:forKey:
func (c_ Coder) DecodeObjectOfClassForKey(aClass objc.Class, key string) objc.ID {
	sel := objc.RegisterName("decodeObjectOfClass:forKey:")
	ret := c_.ID.Send(sel, aClass, key)
	return ret
}

