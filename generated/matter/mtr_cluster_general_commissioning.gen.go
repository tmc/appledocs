// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterGeneralCommissioning] class.
var (
	MTRClusterGeneralCommissioningClass     _MTRClusterGeneralCommissioningClass
	MTRClusterGeneralCommissioningClassOnce sync.Once
)

func getMTRClusterGeneralCommissioningClass() _MTRClusterGeneralCommissioningClass {
	MTRClusterGeneralCommissioningClassOnce.Do(func() {
		MTRClusterGeneralCommissioningClass = _MTRClusterGeneralCommissioningClass{objc.GetClass("MTRClusterGeneralCommissioning")}
	})
	return MTRClusterGeneralCommissioningClass
}

type _MTRClusterGeneralCommissioningClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterGeneralCommissioning] class.
type IMTRClusterGeneralCommissioning interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterGeneralCommissioning
type MTRClusterGeneralCommissioning struct {
	MTRGenericCluster
}

// MTRClusterGeneralCommissioningFrom constructs a [MTRClusterGeneralCommissioning] from an unsafe.Pointer.
func MTRClusterGeneralCommissioningFrom(ptr unsafe.Pointer) MTRClusterGeneralCommissioning {
	return MTRClusterGeneralCommissioning{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterGeneralCommissioningClass) Alloc() MTRClusterGeneralCommissioning {
	rv := objc.Send[MTRClusterGeneralCommissioning](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterGeneralCommissioningClass) New() MTRClusterGeneralCommissioning {
	rv := objc.Send[MTRClusterGeneralCommissioning](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterGeneralCommissioning) Init() MTRClusterGeneralCommissioning {
	rv := objc.Send[MTRClusterGeneralCommissioning](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterGeneralCommissioning) Autorelease() MTRClusterGeneralCommissioning {
	rv := objc.Send[MTRClusterGeneralCommissioning](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterGeneralCommissioning creates a new MTRClusterGeneralCommissioning instance.
func NewMTRClusterGeneralCommissioning() MTRClusterGeneralCommissioning {
	return getMTRClusterGeneralCommissioningClass().New()
}
