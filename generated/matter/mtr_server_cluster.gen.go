// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRServerCluster] class.
var (
	MTRServerClusterClass     _MTRServerClusterClass
	MTRServerClusterClassOnce sync.Once
)

func getMTRServerClusterClass() _MTRServerClusterClass {
	MTRServerClusterClassOnce.Do(func() {
		MTRServerClusterClass = _MTRServerClusterClass{objc.GetClass("MTRServerCluster")}
	})
	return MTRServerClusterClass
}

type _MTRServerClusterClass struct {
	class objc.Class
}

// An interface definition for the [MTRServerCluster] class.
type IMTRServerCluster interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServerCluster
type MTRServerCluster struct {
	objectivec.Object
}

// MTRServerClusterFrom constructs a [MTRServerCluster] from an unsafe.Pointer.
func MTRServerClusterFrom(ptr unsafe.Pointer) MTRServerCluster {
	return MTRServerCluster{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServerClusterClass) Alloc() MTRServerCluster {
	rv := objc.Send[MTRServerCluster](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServerClusterClass) New() MTRServerCluster {
	rv := objc.Send[MTRServerCluster](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServerCluster) Init() MTRServerCluster {
	rv := objc.Send[MTRServerCluster](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServerCluster) Autorelease() MTRServerCluster {
	rv := objc.Send[MTRServerCluster](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServerCluster creates a new MTRServerCluster instance.
func NewMTRServerCluster() MTRServerCluster {
	return getMTRServerClusterClass().New()
}




