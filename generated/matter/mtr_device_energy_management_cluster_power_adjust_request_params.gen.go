// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementClusterPowerAdjustRequestParams] class.
var (
	MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass     _MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass
	MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass() _MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass {
	MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass = _MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterPowerAdjustRequestParams")}
	})
	return MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass
}

type _MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterPowerAdjustRequestParams] class.
type IMTRDeviceEnergyManagementClusterPowerAdjustRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustRequestParams
type MTRDeviceEnergyManagementClusterPowerAdjustRequestParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsFrom constructs a [MTRDeviceEnergyManagementClusterPowerAdjustRequestParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterPowerAdjustRequestParams {
	return MTRDeviceEnergyManagementClusterPowerAdjustRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass) Alloc() MTRDeviceEnergyManagementClusterPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass) New() MTRDeviceEnergyManagementClusterPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) Init() MTRDeviceEnergyManagementClusterPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) Autorelease() MTRDeviceEnergyManagementClusterPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterPowerAdjustRequestParams creates a new MTRDeviceEnergyManagementClusterPowerAdjustRequestParams instance.
func NewMTRDeviceEnergyManagementClusterPowerAdjustRequestParams() MTRDeviceEnergyManagementClusterPowerAdjustRequestParams {
	return getMTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustRequestParams/cause
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) Cause() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("cause"))
	return rv
}


// SetCause sets the value of the cause property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustRequestParams/cause
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) SetCause(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustRequestParams/duration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) Duration() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustRequestParams/duration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) SetDuration(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustRequestParams/power
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) Power() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("power"))
	return rv
}


// SetPower sets the value of the power property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustRequestParams/power
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) SetPower(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPower:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



