// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementClusterCancelRequestParams] class.
var (
	MTRDeviceEnergyManagementClusterCancelRequestParamsClass     _MTRDeviceEnergyManagementClusterCancelRequestParamsClass
	MTRDeviceEnergyManagementClusterCancelRequestParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterCancelRequestParamsClass() _MTRDeviceEnergyManagementClusterCancelRequestParamsClass {
	MTRDeviceEnergyManagementClusterCancelRequestParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterCancelRequestParamsClass = _MTRDeviceEnergyManagementClusterCancelRequestParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterCancelRequestParams")}
	})
	return MTRDeviceEnergyManagementClusterCancelRequestParamsClass
}

type _MTRDeviceEnergyManagementClusterCancelRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterCancelRequestParams] class.
type IMTRDeviceEnergyManagementClusterCancelRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelRequestParams
type MTRDeviceEnergyManagementClusterCancelRequestParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterCancelRequestParamsFrom constructs a [MTRDeviceEnergyManagementClusterCancelRequestParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterCancelRequestParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterCancelRequestParams {
	return MTRDeviceEnergyManagementClusterCancelRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterCancelRequestParamsClass) Alloc() MTRDeviceEnergyManagementClusterCancelRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterCancelRequestParamsClass) New() MTRDeviceEnergyManagementClusterCancelRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterCancelRequestParams) Init() MTRDeviceEnergyManagementClusterCancelRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterCancelRequestParams) Autorelease() MTRDeviceEnergyManagementClusterCancelRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterCancelRequestParams creates a new MTRDeviceEnergyManagementClusterCancelRequestParams instance.
func NewMTRDeviceEnergyManagementClusterCancelRequestParams() MTRDeviceEnergyManagementClusterCancelRequestParams {
	return getMTRDeviceEnergyManagementClusterCancelRequestParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterCancelRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterCancelRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterCancelRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterCancelRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



