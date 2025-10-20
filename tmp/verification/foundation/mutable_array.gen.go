// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var mutableArrayClass _MutableArrayClass

func init() {
	mutableArrayClass = _MutableArrayClass{objc.GetClass("NSMutableArray")}
}

type _MutableArrayClass struct {
	class objc.Class
}

type MutableArray struct {
	objc.ID
}

func MutableArrayFrom(ptr unsafe.Pointer) MutableArray {
	return MutableArray{
		ID: objc.ID(ptr),
	}
}


// Sorts the receiver using a given array of sort descriptors. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(using:)-4eh07
func (m_ MutableArray) SortUsingDescriptors(sortDescriptors unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingDescriptors:"), sortDescriptors)
}


