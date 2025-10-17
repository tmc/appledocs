// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OrderedCollectionDifference] class.
var orderedCollectionDifferenceClass = _OrderedCollectionDifferenceClass{objc.GetClass("NSOrderedCollectionDifference")}

type _OrderedCollectionDifferenceClass struct {
	class objc.Class
}

// An interface definition for the [OrderedCollectionDifference] class.
type IOrderedCollectionDifference interface {
	objectivec.IObject
}

// An object representing the difference between two ordered collections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference

type OrderedCollectionDifference struct {
	objectivec.Object
}

// OrderedCollectionDifferenceFrom constructs a [OrderedCollectionDifference] from an unsafe.Pointer.
//
// An object representing the difference between two ordered collections.
func OrderedCollectionDifferenceFrom(ptr unsafe.Pointer) OrderedCollectionDifference {
	return OrderedCollectionDifference{objectivec.Object{objc.ID(ptr)}}
}



