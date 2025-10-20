// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var OrderedCollectionDifferenceClass _OrderedCollectionDifferenceClass

func init() {
	OrderedCollectionDifferenceClass = _OrderedCollectionDifferenceClass{objc.GetClass("NSOrderedCollectionDifference")}
}

type _OrderedCollectionDifferenceClass struct {
	class objc.Class
}

type OrderedCollectionDifference struct {
	objc.ID
}

func OrderedCollectionDifferenceFrom(ptr unsafe.Pointer) OrderedCollectionDifference {
	return OrderedCollectionDifference{
		ID: objc.ID(ptr),
	}
}




