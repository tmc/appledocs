// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterOccupancySensing] class.
var (
	MTRClusterOccupancySensingClass     _MTRClusterOccupancySensingClass
	MTRClusterOccupancySensingClassOnce sync.Once
)

func getMTRClusterOccupancySensingClass() _MTRClusterOccupancySensingClass {
	MTRClusterOccupancySensingClassOnce.Do(func() {
		MTRClusterOccupancySensingClass = _MTRClusterOccupancySensingClass{objc.GetClass("MTRClusterOccupancySensing")}
	})
	return MTRClusterOccupancySensingClass
}

type _MTRClusterOccupancySensingClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterOccupancySensing] class.
type IMTRClusterOccupancySensing interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOccupancySensing
type MTRClusterOccupancySensing struct {
	MTRGenericCluster
}

// MTRClusterOccupancySensingFrom constructs a [MTRClusterOccupancySensing] from an unsafe.Pointer.
func MTRClusterOccupancySensingFrom(ptr unsafe.Pointer) MTRClusterOccupancySensing {
	return MTRClusterOccupancySensing{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterOccupancySensingClass) Alloc() MTRClusterOccupancySensing {
	rv := objc.Send[MTRClusterOccupancySensing](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterOccupancySensingClass) New() MTRClusterOccupancySensing {
	rv := objc.Send[MTRClusterOccupancySensing](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterOccupancySensing) Init() MTRClusterOccupancySensing {
	rv := objc.Send[MTRClusterOccupancySensing](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterOccupancySensing) Autorelease() MTRClusterOccupancySensing {
	rv := objc.Send[MTRClusterOccupancySensing](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterOccupancySensing creates a new MTRClusterOccupancySensing instance.
func NewMTRClusterOccupancySensing() MTRClusterOccupancySensing {
	return getMTRClusterOccupancySensingClass().New()
}




