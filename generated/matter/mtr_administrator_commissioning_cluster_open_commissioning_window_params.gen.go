// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAdministratorCommissioningClusterOpenCommissioningWindowParams] class.
var (
	MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass     _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass
	MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClassOnce sync.Once
)

func getMTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass() _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass {
	MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClassOnce.Do(func() {
		MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass = _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass{objc.GetClass("MTRAdministratorCommissioningClusterOpenCommissioningWindowParams")}
	})
	return MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass
}

type _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAdministratorCommissioningClusterOpenCommissioningWindowParams] class.
type IMTRAdministratorCommissioningClusterOpenCommissioningWindowParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAdministratorCommissioningClusterOpenCommissioningWindowParams
type MTRAdministratorCommissioningClusterOpenCommissioningWindowParams struct {
	objectivec.Object
}

// MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsFrom constructs a [MTRAdministratorCommissioningClusterOpenCommissioningWindowParams] from an unsafe.Pointer.
func MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsFrom(ptr unsafe.Pointer) MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	return MTRAdministratorCommissioningClusterOpenCommissioningWindowParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass) Alloc() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenCommissioningWindowParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass) New() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenCommissioningWindowParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) Init() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenCommissioningWindowParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAdministratorCommissioningClusterOpenCommissioningWindowParams) Autorelease() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	rv := objc.Send[MTRAdministratorCommissioningClusterOpenCommissioningWindowParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAdministratorCommissioningClusterOpenCommissioningWindowParams creates a new MTRAdministratorCommissioningClusterOpenCommissioningWindowParams instance.
func NewMTRAdministratorCommissioningClusterOpenCommissioningWindowParams() MTRAdministratorCommissioningClusterOpenCommissioningWindowParams {
	return getMTRAdministratorCommissioningClusterOpenCommissioningWindowParamsClass().New()
}




