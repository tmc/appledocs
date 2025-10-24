// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterOccupancySensing] class.
var (
	MTRBaseClusterOccupancySensingClass     _MTRBaseClusterOccupancySensingClass
	MTRBaseClusterOccupancySensingClassOnce sync.Once
)

func getMTRBaseClusterOccupancySensingClass() _MTRBaseClusterOccupancySensingClass {
	MTRBaseClusterOccupancySensingClassOnce.Do(func() {
		MTRBaseClusterOccupancySensingClass = _MTRBaseClusterOccupancySensingClass{objc.GetClass("MTRBaseClusterOccupancySensing")}
	})
	return MTRBaseClusterOccupancySensingClass
}

type _MTRBaseClusterOccupancySensingClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterOccupancySensing] class.
type IMTRBaseClusterOccupancySensing interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOccupancySensing
type MTRBaseClusterOccupancySensing struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterOccupancySensingFrom constructs a [MTRBaseClusterOccupancySensing] from an unsafe.Pointer.
func MTRBaseClusterOccupancySensingFrom(ptr unsafe.Pointer) MTRBaseClusterOccupancySensing {
	return MTRBaseClusterOccupancySensing{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterOccupancySensingClass) Alloc() MTRBaseClusterOccupancySensing {
	rv := objc.Send[MTRBaseClusterOccupancySensing](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterOccupancySensingClass) New() MTRBaseClusterOccupancySensing {
	rv := objc.Send[MTRBaseClusterOccupancySensing](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterOccupancySensing) Init() MTRBaseClusterOccupancySensing {
	rv := objc.Send[MTRBaseClusterOccupancySensing](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterOccupancySensing) Autorelease() MTRBaseClusterOccupancySensing {
	rv := objc.Send[MTRBaseClusterOccupancySensing](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterOccupancySensing creates a new MTRBaseClusterOccupancySensing instance.
func NewMTRBaseClusterOccupancySensing() MTRBaseClusterOccupancySensing {
	return getMTRBaseClusterOccupancySensingClass().New()
}




