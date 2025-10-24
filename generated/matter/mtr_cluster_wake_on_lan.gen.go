// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterWakeOnLan] class.
var (
	MTRClusterWakeOnLanClass     _MTRClusterWakeOnLanClass
	MTRClusterWakeOnLanClassOnce sync.Once
)

func getMTRClusterWakeOnLanClass() _MTRClusterWakeOnLanClass {
	MTRClusterWakeOnLanClassOnce.Do(func() {
		MTRClusterWakeOnLanClass = _MTRClusterWakeOnLanClass{objc.GetClass("MTRClusterWakeOnLan")}
	})
	return MTRClusterWakeOnLanClass
}

type _MTRClusterWakeOnLanClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterWakeOnLan] class.
type IMTRClusterWakeOnLan interface {
	IMTRClusterWakeOnLAN
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWakeOnLan-6e8me
type MTRClusterWakeOnLan struct {
	MTRClusterWakeOnLAN
}

// MTRClusterWakeOnLanFrom constructs a [MTRClusterWakeOnLan] from an unsafe.Pointer.
func MTRClusterWakeOnLanFrom(ptr unsafe.Pointer) MTRClusterWakeOnLan {
	return MTRClusterWakeOnLan{
		MTRClusterWakeOnLAN: MTRClusterWakeOnLANFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterWakeOnLanClass) Alloc() MTRClusterWakeOnLan {
	rv := objc.Send[MTRClusterWakeOnLan](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterWakeOnLanClass) New() MTRClusterWakeOnLan {
	rv := objc.Send[MTRClusterWakeOnLan](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterWakeOnLan) Init() MTRClusterWakeOnLan {
	rv := objc.Send[MTRClusterWakeOnLan](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterWakeOnLan) Autorelease() MTRClusterWakeOnLan {
	rv := objc.Send[MTRClusterWakeOnLan](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterWakeOnLan creates a new MTRClusterWakeOnLan instance.
func NewMTRClusterWakeOnLan() MTRClusterWakeOnLan {
	return getMTRClusterWakeOnLanClass().New()
}
