// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRGenericBaseCluster] class.
var (
	MTRGenericBaseClusterClass     _MTRGenericBaseClusterClass
	MTRGenericBaseClusterClassOnce sync.Once
)

func getMTRGenericBaseClusterClass() _MTRGenericBaseClusterClass {
	MTRGenericBaseClusterClassOnce.Do(func() {
		MTRGenericBaseClusterClass = _MTRGenericBaseClusterClass{objc.GetClass("MTRGenericBaseCluster")}
	})
	return MTRGenericBaseClusterClass
}

type _MTRGenericBaseClusterClass struct {
	class objc.Class
}

// An interface definition for the [MTRGenericBaseCluster] class.
type IMTRGenericBaseCluster interface {
	IMTRCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGenericBaseCluster
type MTRGenericBaseCluster struct {
	MTRCluster
}

// MTRGenericBaseClusterFrom constructs a [MTRGenericBaseCluster] from an unsafe.Pointer.
func MTRGenericBaseClusterFrom(ptr unsafe.Pointer) MTRGenericBaseCluster {
	return MTRGenericBaseCluster{
		MTRCluster: MTRClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGenericBaseClusterClass) Alloc() MTRGenericBaseCluster {
	rv := objc.Send[MTRGenericBaseCluster](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGenericBaseClusterClass) New() MTRGenericBaseCluster {
	rv := objc.Send[MTRGenericBaseCluster](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGenericBaseCluster) Init() MTRGenericBaseCluster {
	rv := objc.Send[MTRGenericBaseCluster](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGenericBaseCluster) Autorelease() MTRGenericBaseCluster {
	rv := objc.Send[MTRGenericBaseCluster](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGenericBaseCluster creates a new MTRGenericBaseCluster instance.
func NewMTRGenericBaseCluster() MTRGenericBaseCluster {
	return getMTRGenericBaseClusterClass().New()
}




