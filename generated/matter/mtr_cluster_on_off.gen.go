// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterOnOff] class.
var (
	MTRClusterOnOffClass     _MTRClusterOnOffClass
	MTRClusterOnOffClassOnce sync.Once
)

func getMTRClusterOnOffClass() _MTRClusterOnOffClass {
	MTRClusterOnOffClassOnce.Do(func() {
		MTRClusterOnOffClass = _MTRClusterOnOffClass{objc.GetClass("MTRClusterOnOff")}
	})
	return MTRClusterOnOffClass
}

type _MTRClusterOnOffClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterOnOff] class.
type IMTRClusterOnOff interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOnOff
type MTRClusterOnOff struct {
	MTRGenericCluster
}

// MTRClusterOnOffFrom constructs a [MTRClusterOnOff] from an unsafe.Pointer.
func MTRClusterOnOffFrom(ptr unsafe.Pointer) MTRClusterOnOff {
	return MTRClusterOnOff{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterOnOffClass) Alloc() MTRClusterOnOff {
	rv := objc.Send[MTRClusterOnOff](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterOnOffClass) New() MTRClusterOnOff {
	rv := objc.Send[MTRClusterOnOff](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterOnOff) Init() MTRClusterOnOff {
	rv := objc.Send[MTRClusterOnOff](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterOnOff) Autorelease() MTRClusterOnOff {
	rv := objc.Send[MTRClusterOnOff](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterOnOff creates a new MTRClusterOnOff instance.
func NewMTRClusterOnOff() MTRClusterOnOff {
	return getMTRClusterOnOffClass().New()
}
