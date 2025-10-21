// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterAdministratorCommissioning] class.
var (
	MTRBaseClusterAdministratorCommissioningClass     _MTRBaseClusterAdministratorCommissioningClass
	MTRBaseClusterAdministratorCommissioningClassOnce sync.Once
)

func getMTRBaseClusterAdministratorCommissioningClass() _MTRBaseClusterAdministratorCommissioningClass {
	MTRBaseClusterAdministratorCommissioningClassOnce.Do(func() {
		MTRBaseClusterAdministratorCommissioningClass = _MTRBaseClusterAdministratorCommissioningClass{objc.GetClass("MTRBaseClusterAdministratorCommissioning")}
	})
	return MTRBaseClusterAdministratorCommissioningClass
}

type _MTRBaseClusterAdministratorCommissioningClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterAdministratorCommissioning] class.
type IMTRBaseClusterAdministratorCommissioning interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterAdministratorCommissioning
type MTRBaseClusterAdministratorCommissioning struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterAdministratorCommissioningFrom constructs a [MTRBaseClusterAdministratorCommissioning] from an unsafe.Pointer.
func MTRBaseClusterAdministratorCommissioningFrom(ptr unsafe.Pointer) MTRBaseClusterAdministratorCommissioning {
	return MTRBaseClusterAdministratorCommissioning{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterAdministratorCommissioningClass) Alloc() MTRBaseClusterAdministratorCommissioning {
	rv := objc.Send[MTRBaseClusterAdministratorCommissioning](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterAdministratorCommissioningClass) New() MTRBaseClusterAdministratorCommissioning {
	rv := objc.Send[MTRBaseClusterAdministratorCommissioning](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterAdministratorCommissioning) Init() MTRBaseClusterAdministratorCommissioning {
	rv := objc.Send[MTRBaseClusterAdministratorCommissioning](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterAdministratorCommissioning) Autorelease() MTRBaseClusterAdministratorCommissioning {
	rv := objc.Send[MTRBaseClusterAdministratorCommissioning](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterAdministratorCommissioning creates a new MTRBaseClusterAdministratorCommissioning instance.
func NewMTRBaseClusterAdministratorCommissioning() MTRBaseClusterAdministratorCommissioning {
	return getMTRBaseClusterAdministratorCommissioningClass().New()
}




