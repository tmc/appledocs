// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterAdministratorCommissioning] class.
var (
	MTRClusterAdministratorCommissioningClass     _MTRClusterAdministratorCommissioningClass
	MTRClusterAdministratorCommissioningClassOnce sync.Once
)

func getMTRClusterAdministratorCommissioningClass() _MTRClusterAdministratorCommissioningClass {
	MTRClusterAdministratorCommissioningClassOnce.Do(func() {
		MTRClusterAdministratorCommissioningClass = _MTRClusterAdministratorCommissioningClass{objc.GetClass("MTRClusterAdministratorCommissioning")}
	})
	return MTRClusterAdministratorCommissioningClass
}

type _MTRClusterAdministratorCommissioningClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterAdministratorCommissioning] class.
type IMTRClusterAdministratorCommissioning interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterAdministratorCommissioning
type MTRClusterAdministratorCommissioning struct {
	MTRGenericCluster
}

// MTRClusterAdministratorCommissioningFrom constructs a [MTRClusterAdministratorCommissioning] from an unsafe.Pointer.
func MTRClusterAdministratorCommissioningFrom(ptr unsafe.Pointer) MTRClusterAdministratorCommissioning {
	return MTRClusterAdministratorCommissioning{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterAdministratorCommissioningClass) Alloc() MTRClusterAdministratorCommissioning {
	rv := objc.Send[MTRClusterAdministratorCommissioning](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterAdministratorCommissioningClass) New() MTRClusterAdministratorCommissioning {
	rv := objc.Send[MTRClusterAdministratorCommissioning](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterAdministratorCommissioning) Init() MTRClusterAdministratorCommissioning {
	rv := objc.Send[MTRClusterAdministratorCommissioning](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterAdministratorCommissioning) Autorelease() MTRClusterAdministratorCommissioning {
	rv := objc.Send[MTRClusterAdministratorCommissioning](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterAdministratorCommissioning creates a new MTRClusterAdministratorCommissioning instance.
func NewMTRClusterAdministratorCommissioning() MTRClusterAdministratorCommissioning {
	return getMTRClusterAdministratorCommissioningClass().New()
}




