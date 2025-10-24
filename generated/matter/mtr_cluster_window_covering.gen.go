// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterWindowCovering] class.
var (
	MTRClusterWindowCoveringClass     _MTRClusterWindowCoveringClass
	MTRClusterWindowCoveringClassOnce sync.Once
)

func getMTRClusterWindowCoveringClass() _MTRClusterWindowCoveringClass {
	MTRClusterWindowCoveringClassOnce.Do(func() {
		MTRClusterWindowCoveringClass = _MTRClusterWindowCoveringClass{objc.GetClass("MTRClusterWindowCovering")}
	})
	return MTRClusterWindowCoveringClass
}

type _MTRClusterWindowCoveringClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterWindowCovering] class.
type IMTRClusterWindowCovering interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWindowCovering
type MTRClusterWindowCovering struct {
	MTRGenericCluster
}

// MTRClusterWindowCoveringFrom constructs a [MTRClusterWindowCovering] from an unsafe.Pointer.
func MTRClusterWindowCoveringFrom(ptr unsafe.Pointer) MTRClusterWindowCovering {
	return MTRClusterWindowCovering{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterWindowCoveringClass) Alloc() MTRClusterWindowCovering {
	rv := objc.Send[MTRClusterWindowCovering](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterWindowCoveringClass) New() MTRClusterWindowCovering {
	rv := objc.Send[MTRClusterWindowCovering](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterWindowCovering) Init() MTRClusterWindowCovering {
	rv := objc.Send[MTRClusterWindowCovering](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterWindowCovering) Autorelease() MTRClusterWindowCovering {
	rv := objc.Send[MTRClusterWindowCovering](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterWindowCovering creates a new MTRClusterWindowCovering instance.
func NewMTRClusterWindowCovering() MTRClusterWindowCovering {
	return getMTRClusterWindowCoveringClass().New()
}




