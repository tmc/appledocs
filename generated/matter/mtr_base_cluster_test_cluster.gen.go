// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterTestCluster] class.
var (
	MTRBaseClusterTestClusterClass     _MTRBaseClusterTestClusterClass
	MTRBaseClusterTestClusterClassOnce sync.Once
)

func getMTRBaseClusterTestClusterClass() _MTRBaseClusterTestClusterClass {
	MTRBaseClusterTestClusterClassOnce.Do(func() {
		MTRBaseClusterTestClusterClass = _MTRBaseClusterTestClusterClass{objc.GetClass("MTRBaseClusterTestCluster")}
	})
	return MTRBaseClusterTestClusterClass
}

type _MTRBaseClusterTestClusterClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterTestCluster] class.
type IMTRBaseClusterTestCluster interface {
	IMTRBaseClusterUnitTesting
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTestCluster
type MTRBaseClusterTestCluster struct {
	MTRBaseClusterUnitTesting
}

// MTRBaseClusterTestClusterFrom constructs a [MTRBaseClusterTestCluster] from an unsafe.Pointer.
func MTRBaseClusterTestClusterFrom(ptr unsafe.Pointer) MTRBaseClusterTestCluster {
	return MTRBaseClusterTestCluster{
		MTRBaseClusterUnitTesting: MTRBaseClusterUnitTestingFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterTestClusterClass) Alloc() MTRBaseClusterTestCluster {
	rv := objc.Send[MTRBaseClusterTestCluster](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterTestClusterClass) New() MTRBaseClusterTestCluster {
	rv := objc.Send[MTRBaseClusterTestCluster](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterTestCluster) Init() MTRBaseClusterTestCluster {
	rv := objc.Send[MTRBaseClusterTestCluster](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterTestCluster) Autorelease() MTRBaseClusterTestCluster {
	rv := objc.Send[MTRBaseClusterTestCluster](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterTestCluster creates a new MTRBaseClusterTestCluster instance.
func NewMTRBaseClusterTestCluster() MTRBaseClusterTestCluster {
	return getMTRBaseClusterTestClusterClass().New()
}




