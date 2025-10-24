// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterLevelControl] class.
var (
	MTRClusterLevelControlClass     _MTRClusterLevelControlClass
	MTRClusterLevelControlClassOnce sync.Once
)

func getMTRClusterLevelControlClass() _MTRClusterLevelControlClass {
	MTRClusterLevelControlClassOnce.Do(func() {
		MTRClusterLevelControlClass = _MTRClusterLevelControlClass{objc.GetClass("MTRClusterLevelControl")}
	})
	return MTRClusterLevelControlClass
}

type _MTRClusterLevelControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterLevelControl] class.
type IMTRClusterLevelControl interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLevelControl
type MTRClusterLevelControl struct {
	MTRGenericCluster
}

// MTRClusterLevelControlFrom constructs a [MTRClusterLevelControl] from an unsafe.Pointer.
func MTRClusterLevelControlFrom(ptr unsafe.Pointer) MTRClusterLevelControl {
	return MTRClusterLevelControl{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterLevelControlClass) Alloc() MTRClusterLevelControl {
	rv := objc.Send[MTRClusterLevelControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterLevelControlClass) New() MTRClusterLevelControl {
	rv := objc.Send[MTRClusterLevelControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterLevelControl) Init() MTRClusterLevelControl {
	rv := objc.Send[MTRClusterLevelControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterLevelControl) Autorelease() MTRClusterLevelControl {
	rv := objc.Send[MTRClusterLevelControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterLevelControl creates a new MTRClusterLevelControl instance.
func NewMTRClusterLevelControl() MTRClusterLevelControl {
	return getMTRClusterLevelControlClass().New()
}




