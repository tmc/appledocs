// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [HKQueryAnchor] class.
var (
	HKQueryAnchorClass     _HKQueryAnchorClass
	HKQueryAnchorClassOnce sync.Once
)

func getHKQueryAnchorClass() _HKQueryAnchorClass {
	HKQueryAnchorClassOnce.Do(func() {
		HKQueryAnchorClass = _HKQueryAnchorClass{objc.GetClass("HKQueryAnchor")}
	})
	return HKQueryAnchorClass
}

type _HKQueryAnchorClass struct {
	class objc.Class
}

// An interface definition for the [HKQueryAnchor] class.
type IHKQueryAnchor interface {
	objectivec.IObject
}

// An object used to identify all the samples previously returned by an anchored object query.
//
// The system returns objects in both the anchored object query’s results handler and it’s update handler. Use the anchors to query for samples added or deleted after the result or update.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQueryAnchor
type HKQueryAnchor struct {
	objectivec.Object
}

// HKQueryAnchorFrom constructs a [HKQueryAnchor] from an unsafe.Pointer.
//
// An object used to identify all the samples previously returned by an anchored object query.
func HKQueryAnchorFrom(ptr unsafe.Pointer) HKQueryAnchor {
	return HKQueryAnchor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKQueryAnchorClass) Alloc() HKQueryAnchor {
	rv := objc.Send[HKQueryAnchor](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKQueryAnchorClass) New() HKQueryAnchor {
	rv := objc.Send[HKQueryAnchor](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKQueryAnchor) Init() HKQueryAnchor {
	rv := objc.Send[HKQueryAnchor](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKQueryAnchor) Autorelease() HKQueryAnchor {
	rv := objc.Send[HKQueryAnchor](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKQueryAnchor creates a new HKQueryAnchor instance.
func NewHKQueryAnchor() HKQueryAnchor {
	return getHKQueryAnchorClass().New()
}




