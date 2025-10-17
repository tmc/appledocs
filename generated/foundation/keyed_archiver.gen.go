// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [KeyedArchiver] class.
var keyedArchiverClass = _KeyedArchiverClass{objc.GetClass("NSKeyedArchiver")}

type _KeyedArchiverClass struct {
	class objc.Class
}

// An encoder that stores an object’s data to an archive referenced by keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver

type KeyedArchiver struct {
	Coder
}

// KeyedArchiverFrom constructs a [KeyedArchiver] from an unsafe.Pointer.
//
// An encoder that stores an object’s data to an archive referenced by keys.
func KeyedArchiverFrom(ptr unsafe.Pointer) KeyedArchiver {
	return KeyedArchiver{
		Coder: CoderFrom(ptr),
	}
}

// Encodes a given value and associates it with a key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encode(_:forKey:)-1mkfl
func (k_ KeyedArchiver) EncodeDoubleForKey(value float64, key string) {
	objc.Send[objc.ID](k_.ID, objc.Sel("encodeDouble:forKey:"), value, key)
}


