// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterBarrierControl] class.
var (
	MTRBaseClusterBarrierControlClass     _MTRBaseClusterBarrierControlClass
	MTRBaseClusterBarrierControlClassOnce sync.Once
)

func getMTRBaseClusterBarrierControlClass() _MTRBaseClusterBarrierControlClass {
	MTRBaseClusterBarrierControlClassOnce.Do(func() {
		MTRBaseClusterBarrierControlClass = _MTRBaseClusterBarrierControlClass{objc.GetClass("MTRBaseClusterBarrierControl")}
	})
	return MTRBaseClusterBarrierControlClass
}

type _MTRBaseClusterBarrierControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterBarrierControl] class.
type IMTRBaseClusterBarrierControl interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterBarrierControl
type MTRBaseClusterBarrierControl struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterBarrierControlFrom constructs a [MTRBaseClusterBarrierControl] from an unsafe.Pointer.
func MTRBaseClusterBarrierControlFrom(ptr unsafe.Pointer) MTRBaseClusterBarrierControl {
	return MTRBaseClusterBarrierControl{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterBarrierControlClass) Alloc() MTRBaseClusterBarrierControl {
	rv := objc.Send[MTRBaseClusterBarrierControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterBarrierControlClass) New() MTRBaseClusterBarrierControl {
	rv := objc.Send[MTRBaseClusterBarrierControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterBarrierControl) Init() MTRBaseClusterBarrierControl {
	rv := objc.Send[MTRBaseClusterBarrierControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterBarrierControl) Autorelease() MTRBaseClusterBarrierControl {
	rv := objc.Send[MTRBaseClusterBarrierControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterBarrierControl creates a new MTRBaseClusterBarrierControl instance.
func NewMTRBaseClusterBarrierControl() MTRBaseClusterBarrierControl {
	return getMTRBaseClusterBarrierControlClass().New()
}
