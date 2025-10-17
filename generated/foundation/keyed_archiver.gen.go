// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [KeyedArchiver] class.
var KeyedArchiverClass objc.Class

func init() {
	KeyedArchiverClass = objc.GetClass("NSKeyedArchiver")
}

type KeyedArchiver struct {
	objc.ID
}

func KeyedArchiverFrom(ptr unsafe.Pointer) KeyedArchiver {
	return KeyedArchiver{
		ID: objc.ID(ptr),
	}
}


// Encodes a given   value and associates it with a key. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSKeyedArchiver/encode(_:forKey:)-1mkfl
func (k_ KeyedArchiver) EncodeDoubleForKey(value float64, key string) {
	sel := objc.RegisterName("encodeDouble:forKey:")
	k_.ID.Send(sel, value, key)
}

