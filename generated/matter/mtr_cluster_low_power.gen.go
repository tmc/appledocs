// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterLowPower] class.
var (
	MTRClusterLowPowerClass     _MTRClusterLowPowerClass
	MTRClusterLowPowerClassOnce sync.Once
)

func getMTRClusterLowPowerClass() _MTRClusterLowPowerClass {
	MTRClusterLowPowerClassOnce.Do(func() {
		MTRClusterLowPowerClass = _MTRClusterLowPowerClass{objc.GetClass("MTRClusterLowPower")}
	})
	return MTRClusterLowPowerClass
}

type _MTRClusterLowPowerClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterLowPower] class.
type IMTRClusterLowPower interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLowPower
type MTRClusterLowPower struct {
	MTRGenericCluster
}

// MTRClusterLowPowerFrom constructs a [MTRClusterLowPower] from an unsafe.Pointer.
func MTRClusterLowPowerFrom(ptr unsafe.Pointer) MTRClusterLowPower {
	return MTRClusterLowPower{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterLowPowerClass) Alloc() MTRClusterLowPower {
	rv := objc.Send[MTRClusterLowPower](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterLowPowerClass) New() MTRClusterLowPower {
	rv := objc.Send[MTRClusterLowPower](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterLowPower) Init() MTRClusterLowPower {
	rv := objc.Send[MTRClusterLowPower](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterLowPower) Autorelease() MTRClusterLowPower {
	rv := objc.Send[MTRClusterLowPower](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterLowPower creates a new MTRClusterLowPower instance.
func NewMTRClusterLowPower() MTRClusterLowPower {
	return getMTRClusterLowPowerClass().New()
}
