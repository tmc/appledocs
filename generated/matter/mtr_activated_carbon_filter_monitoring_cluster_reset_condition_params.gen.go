// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActivatedCarbonFilterMonitoringClusterResetConditionParams] class.
var (
	MTRActivatedCarbonFilterMonitoringClusterResetConditionParamsClass     _MTRActivatedCarbonFilterMonitoringClusterResetConditionParamsClass
	MTRActivatedCarbonFilterMonitoringClusterResetConditionParamsClassOnce sync.Once
)

func getMTRActivatedCarbonFilterMonitoringClusterResetConditionParamsClass() _MTRActivatedCarbonFilterMonitoringClusterResetConditionParamsClass {
	MTRActivatedCarbonFilterMonitoringClusterResetConditionParamsClassOnce.Do(func() {
		MTRActivatedCarbonFilterMonitoringClusterResetConditionParamsClass = _MTRActivatedCarbonFilterMonitoringClusterResetConditionParamsClass{objc.GetClass("MTRActivatedCarbonFilterMonitoringClusterResetConditionParams")}
	})
	return MTRActivatedCarbonFilterMonitoringClusterResetConditionParamsClass
}

type _MTRActivatedCarbonFilterMonitoringClusterResetConditionParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActivatedCarbonFilterMonitoringClusterResetConditionParams] class.
type IMTRActivatedCarbonFilterMonitoringClusterResetConditionParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActivatedCarbonFilterMonitoringClusterResetConditionParams
type MTRActivatedCarbonFilterMonitoringClusterResetConditionParams struct {
	objectivec.Object
}

// MTRActivatedCarbonFilterMonitoringClusterResetConditionParamsFrom constructs a [MTRActivatedCarbonFilterMonitoringClusterResetConditionParams] from an unsafe.Pointer.
func MTRActivatedCarbonFilterMonitoringClusterResetConditionParamsFrom(ptr unsafe.Pointer) MTRActivatedCarbonFilterMonitoringClusterResetConditionParams {
	return MTRActivatedCarbonFilterMonitoringClusterResetConditionParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActivatedCarbonFilterMonitoringClusterResetConditionParamsClass) Alloc() MTRActivatedCarbonFilterMonitoringClusterResetConditionParams {
	rv := objc.Send[MTRActivatedCarbonFilterMonitoringClusterResetConditionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActivatedCarbonFilterMonitoringClusterResetConditionParamsClass) New() MTRActivatedCarbonFilterMonitoringClusterResetConditionParams {
	rv := objc.Send[MTRActivatedCarbonFilterMonitoringClusterResetConditionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActivatedCarbonFilterMonitoringClusterResetConditionParams) Init() MTRActivatedCarbonFilterMonitoringClusterResetConditionParams {
	rv := objc.Send[MTRActivatedCarbonFilterMonitoringClusterResetConditionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActivatedCarbonFilterMonitoringClusterResetConditionParams) Autorelease() MTRActivatedCarbonFilterMonitoringClusterResetConditionParams {
	rv := objc.Send[MTRActivatedCarbonFilterMonitoringClusterResetConditionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActivatedCarbonFilterMonitoringClusterResetConditionParams creates a new MTRActivatedCarbonFilterMonitoringClusterResetConditionParams instance.
func NewMTRActivatedCarbonFilterMonitoringClusterResetConditionParams() MTRActivatedCarbonFilterMonitoringClusterResetConditionParams {
	return getMTRActivatedCarbonFilterMonitoringClusterResetConditionParamsClass().New()
}




