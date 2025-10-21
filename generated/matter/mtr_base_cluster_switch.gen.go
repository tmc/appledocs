// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterSwitch] class.
var (
	MTRBaseClusterSwitchClass     _MTRBaseClusterSwitchClass
	MTRBaseClusterSwitchClassOnce sync.Once
)

func getMTRBaseClusterSwitchClass() _MTRBaseClusterSwitchClass {
	MTRBaseClusterSwitchClassOnce.Do(func() {
		MTRBaseClusterSwitchClass = _MTRBaseClusterSwitchClass{objc.GetClass("MTRBaseClusterSwitch")}
	})
	return MTRBaseClusterSwitchClass
}

type _MTRBaseClusterSwitchClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterSwitch] class.
type IMTRBaseClusterSwitch interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSwitch
type MTRBaseClusterSwitch struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterSwitchFrom constructs a [MTRBaseClusterSwitch] from an unsafe.Pointer.
func MTRBaseClusterSwitchFrom(ptr unsafe.Pointer) MTRBaseClusterSwitch {
	return MTRBaseClusterSwitch{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterSwitchClass) Alloc() MTRBaseClusterSwitch {
	rv := objc.Send[MTRBaseClusterSwitch](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterSwitchClass) New() MTRBaseClusterSwitch {
	rv := objc.Send[MTRBaseClusterSwitch](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterSwitch) Init() MTRBaseClusterSwitch {
	rv := objc.Send[MTRBaseClusterSwitch](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterSwitch) Autorelease() MTRBaseClusterSwitch {
	rv := objc.Send[MTRBaseClusterSwitch](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterSwitch creates a new MTRBaseClusterSwitch instance.
func NewMTRBaseClusterSwitch() MTRBaseClusterSwitch {
	return getMTRBaseClusterSwitchClass().New()
}




