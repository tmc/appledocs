// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableIndexSet] class.
var mutableIndexSetClass = _MutableIndexSetClass{objc.GetClass("NSMutableIndexSet")}

type _MutableIndexSetClass struct {
	class objc.Class
}

// An interface definition for the [MutableIndexSet] class.
type IMutableIndexSet interface {
	IIndexSet
}

// A mutable collection of unique integer values that represent indexes in another collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet

type MutableIndexSet struct {
	IndexSet
}

// MutableIndexSetFrom constructs a [MutableIndexSet] from an unsafe.Pointer.
//
// A mutable collection of unique integer values that represent indexes in another collection.
func MutableIndexSetFrom(ptr unsafe.Pointer) MutableIndexSet {
	return MutableIndexSet{
		IndexSet: IndexSetFrom(ptr),
	}
}



