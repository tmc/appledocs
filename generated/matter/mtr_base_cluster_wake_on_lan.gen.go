// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterWakeOnLan] class.
var (
	MTRBaseClusterWakeOnLanClass     _MTRBaseClusterWakeOnLanClass
	MTRBaseClusterWakeOnLanClassOnce sync.Once
)

func getMTRBaseClusterWakeOnLanClass() _MTRBaseClusterWakeOnLanClass {
	MTRBaseClusterWakeOnLanClassOnce.Do(func() {
		MTRBaseClusterWakeOnLanClass = _MTRBaseClusterWakeOnLanClass{objc.GetClass("MTRBaseClusterWakeOnLan")}
	})
	return MTRBaseClusterWakeOnLanClass
}

type _MTRBaseClusterWakeOnLanClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterWakeOnLan] class.
type IMTRBaseClusterWakeOnLan interface {
	IMTRBaseClusterWakeOnLAN
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWakeOnLan-1dp5d
type MTRBaseClusterWakeOnLan struct {
	MTRBaseClusterWakeOnLAN
}

// MTRBaseClusterWakeOnLanFrom constructs a [MTRBaseClusterWakeOnLan] from an unsafe.Pointer.
func MTRBaseClusterWakeOnLanFrom(ptr unsafe.Pointer) MTRBaseClusterWakeOnLan {
	return MTRBaseClusterWakeOnLan{
		MTRBaseClusterWakeOnLAN: MTRBaseClusterWakeOnLANFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterWakeOnLanClass) Alloc() MTRBaseClusterWakeOnLan {
	rv := objc.Send[MTRBaseClusterWakeOnLan](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterWakeOnLanClass) New() MTRBaseClusterWakeOnLan {
	rv := objc.Send[MTRBaseClusterWakeOnLan](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterWakeOnLan) Init() MTRBaseClusterWakeOnLan {
	rv := objc.Send[MTRBaseClusterWakeOnLan](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterWakeOnLan) Autorelease() MTRBaseClusterWakeOnLan {
	rv := objc.Send[MTRBaseClusterWakeOnLan](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterWakeOnLan creates a new MTRBaseClusterWakeOnLan instance.
func NewMTRBaseClusterWakeOnLan() MTRBaseClusterWakeOnLan {
	return getMTRBaseClusterWakeOnLanClass().New()
}




