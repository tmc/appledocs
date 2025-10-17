// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableArray] class.
var mutableArrayClass = _MutableArrayClass{objc.GetClass("NSMutableArray")}

type _MutableArrayClass struct {
	class objc.Class
}

// A dynamic ordered collection of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray

type MutableArray struct {
	Array
}

// MutableArrayFrom constructs a [MutableArray] from an unsafe.Pointer.
//
// A dynamic ordered collection of objects.
func MutableArrayFrom(ptr unsafe.Pointer) MutableArray {
	return MutableArray{
		Array: ArrayFrom(ptr),
	}
}

// Sorts the receiver using a given array of sort descriptors. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(using:)-4eh07
func (m_ MutableArray) SortUsingDescriptors(sortDescriptors unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingDescriptors:"), sortDescriptors)
}


