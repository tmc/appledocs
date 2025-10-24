// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterPowerSource] class.
var (
	MTRClusterPowerSourceClass     _MTRClusterPowerSourceClass
	MTRClusterPowerSourceClassOnce sync.Once
)

func getMTRClusterPowerSourceClass() _MTRClusterPowerSourceClass {
	MTRClusterPowerSourceClassOnce.Do(func() {
		MTRClusterPowerSourceClass = _MTRClusterPowerSourceClass{objc.GetClass("MTRClusterPowerSource")}
	})
	return MTRClusterPowerSourceClass
}

type _MTRClusterPowerSourceClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterPowerSource] class.
type IMTRClusterPowerSource interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPowerSource
type MTRClusterPowerSource struct {
	MTRGenericCluster
}

// MTRClusterPowerSourceFrom constructs a [MTRClusterPowerSource] from an unsafe.Pointer.
func MTRClusterPowerSourceFrom(ptr unsafe.Pointer) MTRClusterPowerSource {
	return MTRClusterPowerSource{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterPowerSourceClass) Alloc() MTRClusterPowerSource {
	rv := objc.Send[MTRClusterPowerSource](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterPowerSourceClass) New() MTRClusterPowerSource {
	rv := objc.Send[MTRClusterPowerSource](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterPowerSource) Init() MTRClusterPowerSource {
	rv := objc.Send[MTRClusterPowerSource](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterPowerSource) Autorelease() MTRClusterPowerSource {
	rv := objc.Send[MTRClusterPowerSource](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterPowerSource creates a new MTRClusterPowerSource instance.
func NewMTRClusterPowerSource() MTRClusterPowerSource {
	return getMTRClusterPowerSourceClass().New()
}
