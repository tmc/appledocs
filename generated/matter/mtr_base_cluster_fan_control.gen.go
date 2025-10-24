// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterFanControl] class.
var (
	MTRBaseClusterFanControlClass     _MTRBaseClusterFanControlClass
	MTRBaseClusterFanControlClassOnce sync.Once
)

func getMTRBaseClusterFanControlClass() _MTRBaseClusterFanControlClass {
	MTRBaseClusterFanControlClassOnce.Do(func() {
		MTRBaseClusterFanControlClass = _MTRBaseClusterFanControlClass{objc.GetClass("MTRBaseClusterFanControl")}
	})
	return MTRBaseClusterFanControlClass
}

type _MTRBaseClusterFanControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterFanControl] class.
type IMTRBaseClusterFanControl interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterFanControl
type MTRBaseClusterFanControl struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterFanControlFrom constructs a [MTRBaseClusterFanControl] from an unsafe.Pointer.
func MTRBaseClusterFanControlFrom(ptr unsafe.Pointer) MTRBaseClusterFanControl {
	return MTRBaseClusterFanControl{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterFanControlClass) Alloc() MTRBaseClusterFanControl {
	rv := objc.Send[MTRBaseClusterFanControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterFanControlClass) New() MTRBaseClusterFanControl {
	rv := objc.Send[MTRBaseClusterFanControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterFanControl) Init() MTRBaseClusterFanControl {
	rv := objc.Send[MTRBaseClusterFanControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterFanControl) Autorelease() MTRBaseClusterFanControl {
	rv := objc.Send[MTRBaseClusterFanControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterFanControl creates a new MTRBaseClusterFanControl instance.
func NewMTRBaseClusterFanControl() MTRBaseClusterFanControl {
	return getMTRBaseClusterFanControlClass().New()
}
