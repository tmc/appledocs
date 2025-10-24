// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterBarrierControl] class.
var (
	MTRClusterBarrierControlClass     _MTRClusterBarrierControlClass
	MTRClusterBarrierControlClassOnce sync.Once
)

func getMTRClusterBarrierControlClass() _MTRClusterBarrierControlClass {
	MTRClusterBarrierControlClassOnce.Do(func() {
		MTRClusterBarrierControlClass = _MTRClusterBarrierControlClass{objc.GetClass("MTRClusterBarrierControl")}
	})
	return MTRClusterBarrierControlClass
}

type _MTRClusterBarrierControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterBarrierControl] class.
type IMTRClusterBarrierControl interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterBarrierControl
type MTRClusterBarrierControl struct {
	MTRGenericCluster
}

// MTRClusterBarrierControlFrom constructs a [MTRClusterBarrierControl] from an unsafe.Pointer.
func MTRClusterBarrierControlFrom(ptr unsafe.Pointer) MTRClusterBarrierControl {
	return MTRClusterBarrierControl{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterBarrierControlClass) Alloc() MTRClusterBarrierControl {
	rv := objc.Send[MTRClusterBarrierControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterBarrierControlClass) New() MTRClusterBarrierControl {
	rv := objc.Send[MTRClusterBarrierControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterBarrierControl) Init() MTRClusterBarrierControl {
	rv := objc.Send[MTRClusterBarrierControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterBarrierControl) Autorelease() MTRClusterBarrierControl {
	rv := objc.Send[MTRClusterBarrierControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterBarrierControl creates a new MTRClusterBarrierControl instance.
func NewMTRClusterBarrierControl() MTRClusterBarrierControl {
	return getMTRClusterBarrierControlClass().New()
}




