// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterLowPower] class.
var (
	MTRBaseClusterLowPowerClass     _MTRBaseClusterLowPowerClass
	MTRBaseClusterLowPowerClassOnce sync.Once
)

func getMTRBaseClusterLowPowerClass() _MTRBaseClusterLowPowerClass {
	MTRBaseClusterLowPowerClassOnce.Do(func() {
		MTRBaseClusterLowPowerClass = _MTRBaseClusterLowPowerClass{objc.GetClass("MTRBaseClusterLowPower")}
	})
	return MTRBaseClusterLowPowerClass
}

type _MTRBaseClusterLowPowerClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterLowPower] class.
type IMTRBaseClusterLowPower interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLowPower
type MTRBaseClusterLowPower struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterLowPowerFrom constructs a [MTRBaseClusterLowPower] from an unsafe.Pointer.
func MTRBaseClusterLowPowerFrom(ptr unsafe.Pointer) MTRBaseClusterLowPower {
	return MTRBaseClusterLowPower{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterLowPowerClass) Alloc() MTRBaseClusterLowPower {
	rv := objc.Send[MTRBaseClusterLowPower](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterLowPowerClass) New() MTRBaseClusterLowPower {
	rv := objc.Send[MTRBaseClusterLowPower](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterLowPower) Init() MTRBaseClusterLowPower {
	rv := objc.Send[MTRBaseClusterLowPower](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterLowPower) Autorelease() MTRBaseClusterLowPower {
	rv := objc.Send[MTRBaseClusterLowPower](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterLowPower creates a new MTRBaseClusterLowPower instance.
func NewMTRBaseClusterLowPower() MTRBaseClusterLowPower {
	return getMTRBaseClusterLowPowerClass().New()
}
