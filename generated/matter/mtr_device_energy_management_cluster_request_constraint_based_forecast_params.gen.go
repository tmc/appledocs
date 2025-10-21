// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams] class.
var (
	MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass     _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass
	MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass() _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass {
	MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass = _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams")}
	})
	return MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass
}

type _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams] class.
type IMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams
type MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsFrom constructs a [MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	return MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass) Alloc() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass) New() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) Init() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) Autorelease() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams creates a new MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams instance.
func NewMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	return getMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/cause
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) Cause() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("cause"))
	return rv
}


// SetCause sets the value of the cause property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/cause
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) SetCause(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/constraints
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) Constraints() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("constraints"))
	return rv
}


// SetConstraints sets the value of the constraints property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/constraints
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) SetConstraints(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConstraints:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



