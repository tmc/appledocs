// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREnergyEVSEClusterGetTargetsResponseParams] class.
var (
	MTREnergyEVSEClusterGetTargetsResponseParamsClass     _MTREnergyEVSEClusterGetTargetsResponseParamsClass
	MTREnergyEVSEClusterGetTargetsResponseParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterGetTargetsResponseParamsClass() _MTREnergyEVSEClusterGetTargetsResponseParamsClass {
	MTREnergyEVSEClusterGetTargetsResponseParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterGetTargetsResponseParamsClass = _MTREnergyEVSEClusterGetTargetsResponseParamsClass{objc.GetClass("MTREnergyEVSEClusterGetTargetsResponseParams")}
	})
	return MTREnergyEVSEClusterGetTargetsResponseParamsClass
}

type _MTREnergyEVSEClusterGetTargetsResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterGetTargetsResponseParams] class.
type IMTREnergyEVSEClusterGetTargetsResponseParams interface {
	objectivec.IObject
	ChargingTargetSchedules() objc.ID
	SetChargingTargetSchedules(value objc.ID)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsResponseParams
type MTREnergyEVSEClusterGetTargetsResponseParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterGetTargetsResponseParamsFrom constructs a [MTREnergyEVSEClusterGetTargetsResponseParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterGetTargetsResponseParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterGetTargetsResponseParams {
	return MTREnergyEVSEClusterGetTargetsResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterGetTargetsResponseParamsClass) Alloc() MTREnergyEVSEClusterGetTargetsResponseParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterGetTargetsResponseParamsClass) New() MTREnergyEVSEClusterGetTargetsResponseParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterGetTargetsResponseParams) Init() MTREnergyEVSEClusterGetTargetsResponseParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterGetTargetsResponseParams) Autorelease() MTREnergyEVSEClusterGetTargetsResponseParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterGetTargetsResponseParams creates a new MTREnergyEVSEClusterGetTargetsResponseParams instance.
func NewMTREnergyEVSEClusterGetTargetsResponseParams() MTREnergyEVSEClusterGetTargetsResponseParams {
	return getMTREnergyEVSEClusterGetTargetsResponseParamsClass().New()
}




// Initialize an MTREnergyEVSEClusterGetTargetsResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsResponseParams/init(responseValue:)
func NewMTREnergyEVSEClusterGetTargetsResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTREnergyEVSEClusterGetTargetsResponseParams {
	instance := getMTREnergyEVSEClusterGetTargetsResponseParamsClass().Alloc()
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsResponseParams/chargingTargetSchedules
func (m_ MTREnergyEVSEClusterGetTargetsResponseParams) ChargingTargetSchedules() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("chargingTargetSchedules"))
	return rv
}


// SetChargingTargetSchedules sets the value of the chargingTargetSchedules property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsResponseParams/chargingTargetSchedules
func (m_ MTREnergyEVSEClusterGetTargetsResponseParams) SetChargingTargetSchedules(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChargingTargetSchedules:"), value)
}


