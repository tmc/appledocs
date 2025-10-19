// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OrderedCollectionDifference] class.
var (
	orderedCollectionDifferenceClass     _OrderedCollectionDifferenceClass
	orderedCollectionDifferenceClassOnce sync.Once
)

func getOrderedCollectionDifferenceClass() _OrderedCollectionDifferenceClass {
	orderedCollectionDifferenceClassOnce.Do(func() {
		orderedCollectionDifferenceClass = _OrderedCollectionDifferenceClass{objc.GetClass("NSOrderedCollectionDifference")}
	})
	return orderedCollectionDifferenceClass
}

type _OrderedCollectionDifferenceClass struct {
	class objc.Class
}

// An interface definition for the [OrderedCollectionDifference] class.
type IOrderedCollectionDifference interface {
	objectivec.IObject
}

// An object representing the difference between two ordered collections.
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

// Alloc allocates a new instance without initialization.
func (oc _OrderedCollectionDifferenceClass) Alloc() OrderedCollectionDifference {
	rv := objc.Send[OrderedCollectionDifference](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OrderedCollectionDifferenceClass) New() OrderedCollectionDifference {
	rv := objc.Send[OrderedCollectionDifference](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OrderedCollectionDifference) Init() OrderedCollectionDifference {
	rv := objc.Send[OrderedCollectionDifference](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OrderedCollectionDifference) Autorelease() OrderedCollectionDifference {
	rv := objc.Send[OrderedCollectionDifference](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOrderedCollectionDifference creates a new OrderedCollectionDifference instance.
func NewOrderedCollectionDifference() OrderedCollectionDifference {
	return getOrderedCollectionDifferenceClass().New()
}




