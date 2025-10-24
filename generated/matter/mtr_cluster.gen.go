// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRCluster] class.
var (
	MTRClusterClass     _MTRClusterClass
	MTRClusterClassOnce sync.Once
)

func getMTRClusterClass() _MTRClusterClass {
	MTRClusterClassOnce.Do(func() {
		MTRClusterClass = _MTRClusterClass{objc.GetClass("MTRCluster")}
	})
	return MTRClusterClass
}

type _MTRClusterClass struct {
	class objc.Class
}

// An interface definition for the [MTRCluster] class.
type IMTRCluster interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCluster
type MTRCluster struct {
	objectivec.Object
}

// MTRClusterFrom constructs a [MTRCluster] from an unsafe.Pointer.
func MTRClusterFrom(ptr unsafe.Pointer) MTRCluster {
	return MTRCluster{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterClass) Alloc() MTRCluster {
	rv := objc.Send[MTRCluster](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterClass) New() MTRCluster {
	rv := objc.Send[MTRCluster](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCluster) Init() MTRCluster {
	rv := objc.Send[MTRCluster](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCluster) Autorelease() MTRCluster {
	rv := objc.Send[MTRCluster](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCluster creates a new MTRCluster instance.
func NewMTRCluster() MTRCluster {
	return getMTRClusterClass().New()
}




