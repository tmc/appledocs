// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRGeneralCommissioningClusterArmFailSafeParams] class.
var (
	MTRGeneralCommissioningClusterArmFailSafeParamsClass     _MTRGeneralCommissioningClusterArmFailSafeParamsClass
	MTRGeneralCommissioningClusterArmFailSafeParamsClassOnce sync.Once
)

func getMTRGeneralCommissioningClusterArmFailSafeParamsClass() _MTRGeneralCommissioningClusterArmFailSafeParamsClass {
	MTRGeneralCommissioningClusterArmFailSafeParamsClassOnce.Do(func() {
		MTRGeneralCommissioningClusterArmFailSafeParamsClass = _MTRGeneralCommissioningClusterArmFailSafeParamsClass{objc.GetClass("MTRGeneralCommissioningClusterArmFailSafeParams")}
	})
	return MTRGeneralCommissioningClusterArmFailSafeParamsClass
}

type _MTRGeneralCommissioningClusterArmFailSafeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralCommissioningClusterArmFailSafeParams] class.
type IMTRGeneralCommissioningClusterArmFailSafeParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralCommissioningClusterArmFailSafeParams
type MTRGeneralCommissioningClusterArmFailSafeParams struct {
	objectivec.Object
}

// MTRGeneralCommissioningClusterArmFailSafeParamsFrom constructs a [MTRGeneralCommissioningClusterArmFailSafeParams] from an unsafe.Pointer.
func MTRGeneralCommissioningClusterArmFailSafeParamsFrom(ptr unsafe.Pointer) MTRGeneralCommissioningClusterArmFailSafeParams {
	return MTRGeneralCommissioningClusterArmFailSafeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralCommissioningClusterArmFailSafeParamsClass) Alloc() MTRGeneralCommissioningClusterArmFailSafeParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralCommissioningClusterArmFailSafeParamsClass) New() MTRGeneralCommissioningClusterArmFailSafeParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) Init() MTRGeneralCommissioningClusterArmFailSafeParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralCommissioningClusterArmFailSafeParams) Autorelease() MTRGeneralCommissioningClusterArmFailSafeParams {
	rv := objc.Send[MTRGeneralCommissioningClusterArmFailSafeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralCommissioningClusterArmFailSafeParams creates a new MTRGeneralCommissioningClusterArmFailSafeParams instance.
func NewMTRGeneralCommissioningClusterArmFailSafeParams() MTRGeneralCommissioningClusterArmFailSafeParams {
	return getMTRGeneralCommissioningClusterArmFailSafeParamsClass().New()
}




