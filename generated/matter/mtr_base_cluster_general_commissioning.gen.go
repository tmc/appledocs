// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterGeneralCommissioning] class.
var (
	MTRBaseClusterGeneralCommissioningClass     _MTRBaseClusterGeneralCommissioningClass
	MTRBaseClusterGeneralCommissioningClassOnce sync.Once
)

func getMTRBaseClusterGeneralCommissioningClass() _MTRBaseClusterGeneralCommissioningClass {
	MTRBaseClusterGeneralCommissioningClassOnce.Do(func() {
		MTRBaseClusterGeneralCommissioningClass = _MTRBaseClusterGeneralCommissioningClass{objc.GetClass("MTRBaseClusterGeneralCommissioning")}
	})
	return MTRBaseClusterGeneralCommissioningClass
}

type _MTRBaseClusterGeneralCommissioningClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterGeneralCommissioning] class.
type IMTRBaseClusterGeneralCommissioning interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterGeneralCommissioning
type MTRBaseClusterGeneralCommissioning struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterGeneralCommissioningFrom constructs a [MTRBaseClusterGeneralCommissioning] from an unsafe.Pointer.
func MTRBaseClusterGeneralCommissioningFrom(ptr unsafe.Pointer) MTRBaseClusterGeneralCommissioning {
	return MTRBaseClusterGeneralCommissioning{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterGeneralCommissioningClass) Alloc() MTRBaseClusterGeneralCommissioning {
	rv := objc.Send[MTRBaseClusterGeneralCommissioning](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterGeneralCommissioningClass) New() MTRBaseClusterGeneralCommissioning {
	rv := objc.Send[MTRBaseClusterGeneralCommissioning](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterGeneralCommissioning) Init() MTRBaseClusterGeneralCommissioning {
	rv := objc.Send[MTRBaseClusterGeneralCommissioning](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterGeneralCommissioning) Autorelease() MTRBaseClusterGeneralCommissioning {
	rv := objc.Send[MTRBaseClusterGeneralCommissioning](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterGeneralCommissioning creates a new MTRBaseClusterGeneralCommissioning instance.
func NewMTRBaseClusterGeneralCommissioning() MTRBaseClusterGeneralCommissioning {
	return getMTRBaseClusterGeneralCommissioningClass().New()
}




