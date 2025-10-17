// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [KeyedUnarchiver] class.
var keyedUnarchiverClass = _KeyedUnarchiverClass{objc.GetClass("NSKeyedUnarchiver")}

type _KeyedUnarchiverClass struct {
	class objc.Class
}

// A decoder that restores data from an archive referenced by keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver

type KeyedUnarchiver struct {
	Coder
}

// KeyedUnarchiverFrom constructs a [KeyedUnarchiver] from an unsafe.Pointer.
//
// A decoder that restores data from an archive referenced by keys.
func KeyedUnarchiverFrom(ptr unsafe.Pointer) KeyedUnarchiver {
	return KeyedUnarchiver{
		Coder: CoderFrom(ptr),
	}
}



