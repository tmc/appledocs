// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OrderedCollectionDifference] class.
var OrderedCollectionDifferenceClass objc.Class

func init() {
	OrderedCollectionDifferenceClass = objc.GetClass("NSOrderedCollectionDifference")
}

type OrderedCollectionDifference struct {
	objc.ID
}

func OrderedCollectionDifferenceFrom(ptr unsafe.Pointer) OrderedCollectionDifference {
	return OrderedCollectionDifference{
		ID: objc.ID(ptr),
	}
}




