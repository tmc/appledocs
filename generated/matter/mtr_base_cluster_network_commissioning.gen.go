// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterNetworkCommissioning] class.
var (
	MTRBaseClusterNetworkCommissioningClass     _MTRBaseClusterNetworkCommissioningClass
	MTRBaseClusterNetworkCommissioningClassOnce sync.Once
)

func getMTRBaseClusterNetworkCommissioningClass() _MTRBaseClusterNetworkCommissioningClass {
	MTRBaseClusterNetworkCommissioningClassOnce.Do(func() {
		MTRBaseClusterNetworkCommissioningClass = _MTRBaseClusterNetworkCommissioningClass{objc.GetClass("MTRBaseClusterNetworkCommissioning")}
	})
	return MTRBaseClusterNetworkCommissioningClass
}

type _MTRBaseClusterNetworkCommissioningClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterNetworkCommissioning] class.
type IMTRBaseClusterNetworkCommissioning interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterNetworkCommissioning
type MTRBaseClusterNetworkCommissioning struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterNetworkCommissioningFrom constructs a [MTRBaseClusterNetworkCommissioning] from an unsafe.Pointer.
func MTRBaseClusterNetworkCommissioningFrom(ptr unsafe.Pointer) MTRBaseClusterNetworkCommissioning {
	return MTRBaseClusterNetworkCommissioning{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterNetworkCommissioningClass) Alloc() MTRBaseClusterNetworkCommissioning {
	rv := objc.Send[MTRBaseClusterNetworkCommissioning](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterNetworkCommissioningClass) New() MTRBaseClusterNetworkCommissioning {
	rv := objc.Send[MTRBaseClusterNetworkCommissioning](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterNetworkCommissioning) Init() MTRBaseClusterNetworkCommissioning {
	rv := objc.Send[MTRBaseClusterNetworkCommissioning](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterNetworkCommissioning) Autorelease() MTRBaseClusterNetworkCommissioning {
	rv := objc.Send[MTRBaseClusterNetworkCommissioning](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterNetworkCommissioning creates a new MTRBaseClusterNetworkCommissioning instance.
func NewMTRBaseClusterNetworkCommissioning() MTRBaseClusterNetworkCommissioning {
	return getMTRBaseClusterNetworkCommissioningClass().New()
}




