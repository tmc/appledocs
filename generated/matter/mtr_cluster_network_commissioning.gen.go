// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterNetworkCommissioning] class.
var (
	MTRClusterNetworkCommissioningClass     _MTRClusterNetworkCommissioningClass
	MTRClusterNetworkCommissioningClassOnce sync.Once
)

func getMTRClusterNetworkCommissioningClass() _MTRClusterNetworkCommissioningClass {
	MTRClusterNetworkCommissioningClassOnce.Do(func() {
		MTRClusterNetworkCommissioningClass = _MTRClusterNetworkCommissioningClass{objc.GetClass("MTRClusterNetworkCommissioning")}
	})
	return MTRClusterNetworkCommissioningClass
}

type _MTRClusterNetworkCommissioningClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterNetworkCommissioning] class.
type IMTRClusterNetworkCommissioning interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterNetworkCommissioning
type MTRClusterNetworkCommissioning struct {
	MTRGenericCluster
}

// MTRClusterNetworkCommissioningFrom constructs a [MTRClusterNetworkCommissioning] from an unsafe.Pointer.
func MTRClusterNetworkCommissioningFrom(ptr unsafe.Pointer) MTRClusterNetworkCommissioning {
	return MTRClusterNetworkCommissioning{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterNetworkCommissioningClass) Alloc() MTRClusterNetworkCommissioning {
	rv := objc.Send[MTRClusterNetworkCommissioning](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterNetworkCommissioningClass) New() MTRClusterNetworkCommissioning {
	rv := objc.Send[MTRClusterNetworkCommissioning](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterNetworkCommissioning) Init() MTRClusterNetworkCommissioning {
	rv := objc.Send[MTRClusterNetworkCommissioning](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterNetworkCommissioning) Autorelease() MTRClusterNetworkCommissioning {
	rv := objc.Send[MTRClusterNetworkCommissioning](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterNetworkCommissioning creates a new MTRClusterNetworkCommissioning instance.
func NewMTRClusterNetworkCommissioning() MTRClusterNetworkCommissioning {
	return getMTRClusterNetworkCommissioningClass().New()
}




