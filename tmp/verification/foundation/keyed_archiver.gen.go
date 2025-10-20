// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var keyedArchiverClass _KeyedArchiverClass

func init() {
	keyedArchiverClass = _KeyedArchiverClass{objc.GetClass("NSKeyedArchiver")}
}

type _KeyedArchiverClass struct {
	class objc.Class
}

type KeyedArchiver struct {
	objc.ID
}

func KeyedArchiverFrom(ptr unsafe.Pointer) KeyedArchiver {
	return KeyedArchiver{
		ID: objc.ID(ptr),
	}
}


// Encodes a given value and associates it with a key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encode(_:forKey:)-1mkfl
func (k_ KeyedArchiver) EncodeDoubleForKey(value float64, key string) {
	objc.Send[objc.ID](k_.ID, objc.Sel("encodeDouble:forKey:"), value, key)
}


