// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams] class.
var (
	MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass     _MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass
	MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass() _MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass {
	MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass = _MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams")}
	})
	return MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass
}

type _MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams] class.
type IMTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams
type MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsFrom constructs a [MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams {
	return MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass) Alloc() MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass) New() MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams) Init() MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams) Autorelease() MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams creates a new MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams instance.
func NewMTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams() MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams {
	return getMTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



