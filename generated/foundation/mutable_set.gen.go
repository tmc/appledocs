// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableSet] class.
var mutableSetClass = _MutableSetClass{objc.GetClass("NSMutableSet")}

type _MutableSetClass struct {
	class objc.Class
}

// A dynamic unordered collection of unique objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet

type MutableSet struct {
	Set
}

// MutableSetFrom constructs a [MutableSet] from an unsafe.Pointer.
//
// A dynamic unordered collection of unique objects.
func MutableSetFrom(ptr unsafe.Pointer) MutableSet {
	return MutableSet{
		Set: SetFrom(ptr),
	}
}



