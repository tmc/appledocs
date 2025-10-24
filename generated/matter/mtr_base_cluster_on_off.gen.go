// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterOnOff] class.
var (
	MTRBaseClusterOnOffClass     _MTRBaseClusterOnOffClass
	MTRBaseClusterOnOffClassOnce sync.Once
)

func getMTRBaseClusterOnOffClass() _MTRBaseClusterOnOffClass {
	MTRBaseClusterOnOffClassOnce.Do(func() {
		MTRBaseClusterOnOffClass = _MTRBaseClusterOnOffClass{objc.GetClass("MTRBaseClusterOnOff")}
	})
	return MTRBaseClusterOnOffClass
}

type _MTRBaseClusterOnOffClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterOnOff] class.
type IMTRBaseClusterOnOff interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOnOff
type MTRBaseClusterOnOff struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterOnOffFrom constructs a [MTRBaseClusterOnOff] from an unsafe.Pointer.
func MTRBaseClusterOnOffFrom(ptr unsafe.Pointer) MTRBaseClusterOnOff {
	return MTRBaseClusterOnOff{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterOnOffClass) Alloc() MTRBaseClusterOnOff {
	rv := objc.Send[MTRBaseClusterOnOff](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterOnOffClass) New() MTRBaseClusterOnOff {
	rv := objc.Send[MTRBaseClusterOnOff](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterOnOff) Init() MTRBaseClusterOnOff {
	rv := objc.Send[MTRBaseClusterOnOff](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterOnOff) Autorelease() MTRBaseClusterOnOff {
	rv := objc.Send[MTRBaseClusterOnOff](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterOnOff creates a new MTRBaseClusterOnOff instance.
func NewMTRBaseClusterOnOff() MTRBaseClusterOnOff {
	return getMTRBaseClusterOnOffClass().New()
}




