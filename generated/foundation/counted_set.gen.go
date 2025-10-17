// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CountedSet] class.
var countedSetClass = _CountedSetClass{objc.GetClass("NSCountedSet")}

type _CountedSetClass struct {
	class objc.Class
}

// An interface definition for the [CountedSet] class.
type ICountedSet interface {
	IMutableSet
}

// A mutable, unordered collection of distinct objects that may appear more than once in the collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet

type CountedSet struct {
	MutableSet
}

// CountedSetFrom constructs a [CountedSet] from an unsafe.Pointer.
//
// A mutable, unordered collection of distinct objects that may appear more than once in the collection.
func CountedSetFrom(ptr unsafe.Pointer) CountedSet {
	return CountedSet{
		MutableSet: MutableSetFrom(ptr),
	}
}



