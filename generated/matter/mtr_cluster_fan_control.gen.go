// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterFanControl] class.
var (
	MTRClusterFanControlClass     _MTRClusterFanControlClass
	MTRClusterFanControlClassOnce sync.Once
)

func getMTRClusterFanControlClass() _MTRClusterFanControlClass {
	MTRClusterFanControlClassOnce.Do(func() {
		MTRClusterFanControlClass = _MTRClusterFanControlClass{objc.GetClass("MTRClusterFanControl")}
	})
	return MTRClusterFanControlClass
}

type _MTRClusterFanControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterFanControl] class.
type IMTRClusterFanControl interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterFanControl
type MTRClusterFanControl struct {
	MTRGenericCluster
}

// MTRClusterFanControlFrom constructs a [MTRClusterFanControl] from an unsafe.Pointer.
func MTRClusterFanControlFrom(ptr unsafe.Pointer) MTRClusterFanControl {
	return MTRClusterFanControl{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterFanControlClass) Alloc() MTRClusterFanControl {
	rv := objc.Send[MTRClusterFanControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterFanControlClass) New() MTRClusterFanControl {
	rv := objc.Send[MTRClusterFanControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterFanControl) Init() MTRClusterFanControl {
	rv := objc.Send[MTRClusterFanControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterFanControl) Autorelease() MTRClusterFanControl {
	rv := objc.Send[MTRClusterFanControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterFanControl creates a new MTRClusterFanControl instance.
func NewMTRClusterFanControl() MTRClusterFanControl {
	return getMTRClusterFanControlClass().New()
}




