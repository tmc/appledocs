// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MutableArray] class.
var MutableArrayClass objc.Class

func init() {
	MutableArrayClass = objc.GetClass("NSMutableArray")
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
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSMutableArray/sort(using:)-4eh07
func (m_ MutableArray) SortUsingDescriptors(sortDescriptors unsafe.Pointer) {
	sel := objc.RegisterName("sortUsingDescriptors:")
	m_.ID.Send(sel, sortDescriptors)
}

