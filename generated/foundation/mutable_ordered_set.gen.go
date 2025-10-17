// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableOrderedSet] class.
var mutableOrderedSetClass = _MutableOrderedSetClass{objc.GetClass("NSMutableOrderedSet")}

type _MutableOrderedSetClass struct {
	class objc.Class
}

// A dynamic, ordered collection of unique objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet

type MutableOrderedSet struct {
	OrderedSet
}

// MutableOrderedSetFrom constructs a [MutableOrderedSet] from an unsafe.Pointer.
//
// A dynamic, ordered collection of unique objects.
func MutableOrderedSetFrom(ptr unsafe.Pointer) MutableOrderedSet {
	return MutableOrderedSet{
		OrderedSet: OrderedSetFrom(ptr),
	}
}



