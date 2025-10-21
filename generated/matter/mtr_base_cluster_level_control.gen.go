// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterLevelControl] class.
var (
	MTRBaseClusterLevelControlClass     _MTRBaseClusterLevelControlClass
	MTRBaseClusterLevelControlClassOnce sync.Once
)

func getMTRBaseClusterLevelControlClass() _MTRBaseClusterLevelControlClass {
	MTRBaseClusterLevelControlClassOnce.Do(func() {
		MTRBaseClusterLevelControlClass = _MTRBaseClusterLevelControlClass{objc.GetClass("MTRBaseClusterLevelControl")}
	})
	return MTRBaseClusterLevelControlClass
}

type _MTRBaseClusterLevelControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterLevelControl] class.
type IMTRBaseClusterLevelControl interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl
type MTRBaseClusterLevelControl struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterLevelControlFrom constructs a [MTRBaseClusterLevelControl] from an unsafe.Pointer.
func MTRBaseClusterLevelControlFrom(ptr unsafe.Pointer) MTRBaseClusterLevelControl {
	return MTRBaseClusterLevelControl{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterLevelControlClass) Alloc() MTRBaseClusterLevelControl {
	rv := objc.Send[MTRBaseClusterLevelControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterLevelControlClass) New() MTRBaseClusterLevelControl {
	rv := objc.Send[MTRBaseClusterLevelControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterLevelControl) Init() MTRBaseClusterLevelControl {
	rv := objc.Send[MTRBaseClusterLevelControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterLevelControl) Autorelease() MTRBaseClusterLevelControl {
	rv := objc.Send[MTRBaseClusterLevelControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterLevelControl creates a new MTRBaseClusterLevelControl instance.
func NewMTRBaseClusterLevelControl() MTRBaseClusterLevelControl {
	return getMTRBaseClusterLevelControlClass().New()
}




