// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterTestCluster] class.
var (
	MTRClusterTestClusterClass     _MTRClusterTestClusterClass
	MTRClusterTestClusterClassOnce sync.Once
)

func getMTRClusterTestClusterClass() _MTRClusterTestClusterClass {
	MTRClusterTestClusterClassOnce.Do(func() {
		MTRClusterTestClusterClass = _MTRClusterTestClusterClass{objc.GetClass("MTRClusterTestCluster")}
	})
	return MTRClusterTestClusterClass
}

type _MTRClusterTestClusterClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterTestCluster] class.
type IMTRClusterTestCluster interface {
	IMTRClusterUnitTesting
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTestCluster
type MTRClusterTestCluster struct {
	MTRClusterUnitTesting
}

// MTRClusterTestClusterFrom constructs a [MTRClusterTestCluster] from an unsafe.Pointer.
func MTRClusterTestClusterFrom(ptr unsafe.Pointer) MTRClusterTestCluster {
	return MTRClusterTestCluster{
		MTRClusterUnitTesting: MTRClusterUnitTestingFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterTestClusterClass) Alloc() MTRClusterTestCluster {
	rv := objc.Send[MTRClusterTestCluster](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterTestClusterClass) New() MTRClusterTestCluster {
	rv := objc.Send[MTRClusterTestCluster](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterTestCluster) Init() MTRClusterTestCluster {
	rv := objc.Send[MTRClusterTestCluster](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterTestCluster) Autorelease() MTRClusterTestCluster {
	rv := objc.Send[MTRClusterTestCluster](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterTestCluster creates a new MTRClusterTestCluster instance.
func NewMTRClusterTestCluster() MTRClusterTestCluster {
	return getMTRClusterTestClusterClass().New()
}




