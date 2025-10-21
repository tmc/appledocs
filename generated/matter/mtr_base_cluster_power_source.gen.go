// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterPowerSource] class.
var (
	MTRBaseClusterPowerSourceClass     _MTRBaseClusterPowerSourceClass
	MTRBaseClusterPowerSourceClassOnce sync.Once
)

func getMTRBaseClusterPowerSourceClass() _MTRBaseClusterPowerSourceClass {
	MTRBaseClusterPowerSourceClassOnce.Do(func() {
		MTRBaseClusterPowerSourceClass = _MTRBaseClusterPowerSourceClass{objc.GetClass("MTRBaseClusterPowerSource")}
	})
	return MTRBaseClusterPowerSourceClass
}

type _MTRBaseClusterPowerSourceClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterPowerSource] class.
type IMTRBaseClusterPowerSource interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSource
type MTRBaseClusterPowerSource struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterPowerSourceFrom constructs a [MTRBaseClusterPowerSource] from an unsafe.Pointer.
func MTRBaseClusterPowerSourceFrom(ptr unsafe.Pointer) MTRBaseClusterPowerSource {
	return MTRBaseClusterPowerSource{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterPowerSourceClass) Alloc() MTRBaseClusterPowerSource {
	rv := objc.Send[MTRBaseClusterPowerSource](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterPowerSourceClass) New() MTRBaseClusterPowerSource {
	rv := objc.Send[MTRBaseClusterPowerSource](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterPowerSource) Init() MTRBaseClusterPowerSource {
	rv := objc.Send[MTRBaseClusterPowerSource](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterPowerSource) Autorelease() MTRBaseClusterPowerSource {
	rv := objc.Send[MTRBaseClusterPowerSource](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterPowerSource creates a new MTRBaseClusterPowerSource instance.
func NewMTRBaseClusterPowerSource() MTRBaseClusterPowerSource {
	return getMTRBaseClusterPowerSourceClass().New()
}




