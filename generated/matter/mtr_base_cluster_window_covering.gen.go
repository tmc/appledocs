// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterWindowCovering] class.
var (
	MTRBaseClusterWindowCoveringClass     _MTRBaseClusterWindowCoveringClass
	MTRBaseClusterWindowCoveringClassOnce sync.Once
)

func getMTRBaseClusterWindowCoveringClass() _MTRBaseClusterWindowCoveringClass {
	MTRBaseClusterWindowCoveringClassOnce.Do(func() {
		MTRBaseClusterWindowCoveringClass = _MTRBaseClusterWindowCoveringClass{objc.GetClass("MTRBaseClusterWindowCovering")}
	})
	return MTRBaseClusterWindowCoveringClass
}

type _MTRBaseClusterWindowCoveringClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterWindowCovering] class.
type IMTRBaseClusterWindowCovering interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering
type MTRBaseClusterWindowCovering struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterWindowCoveringFrom constructs a [MTRBaseClusterWindowCovering] from an unsafe.Pointer.
func MTRBaseClusterWindowCoveringFrom(ptr unsafe.Pointer) MTRBaseClusterWindowCovering {
	return MTRBaseClusterWindowCovering{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterWindowCoveringClass) Alloc() MTRBaseClusterWindowCovering {
	rv := objc.Send[MTRBaseClusterWindowCovering](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterWindowCoveringClass) New() MTRBaseClusterWindowCovering {
	rv := objc.Send[MTRBaseClusterWindowCovering](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterWindowCovering) Init() MTRBaseClusterWindowCovering {
	rv := objc.Send[MTRBaseClusterWindowCovering](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterWindowCovering) Autorelease() MTRBaseClusterWindowCovering {
	rv := objc.Send[MTRBaseClusterWindowCovering](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterWindowCovering creates a new MTRBaseClusterWindowCovering instance.
func NewMTRBaseClusterWindowCovering() MTRBaseClusterWindowCovering {
	return getMTRBaseClusterWindowCoveringClass().New()
}




