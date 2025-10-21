// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementClusterResumeRequestParams] class.
var (
	MTRDeviceEnergyManagementClusterResumeRequestParamsClass     _MTRDeviceEnergyManagementClusterResumeRequestParamsClass
	MTRDeviceEnergyManagementClusterResumeRequestParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterResumeRequestParamsClass() _MTRDeviceEnergyManagementClusterResumeRequestParamsClass {
	MTRDeviceEnergyManagementClusterResumeRequestParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterResumeRequestParamsClass = _MTRDeviceEnergyManagementClusterResumeRequestParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterResumeRequestParams")}
	})
	return MTRDeviceEnergyManagementClusterResumeRequestParamsClass
}

type _MTRDeviceEnergyManagementClusterResumeRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterResumeRequestParams] class.
type IMTRDeviceEnergyManagementClusterResumeRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumeRequestParams
type MTRDeviceEnergyManagementClusterResumeRequestParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterResumeRequestParamsFrom constructs a [MTRDeviceEnergyManagementClusterResumeRequestParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterResumeRequestParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterResumeRequestParams {
	return MTRDeviceEnergyManagementClusterResumeRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterResumeRequestParamsClass) Alloc() MTRDeviceEnergyManagementClusterResumeRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumeRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterResumeRequestParamsClass) New() MTRDeviceEnergyManagementClusterResumeRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumeRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterResumeRequestParams) Init() MTRDeviceEnergyManagementClusterResumeRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumeRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterResumeRequestParams) Autorelease() MTRDeviceEnergyManagementClusterResumeRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumeRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterResumeRequestParams creates a new MTRDeviceEnergyManagementClusterResumeRequestParams instance.
func NewMTRDeviceEnergyManagementClusterResumeRequestParams() MTRDeviceEnergyManagementClusterResumeRequestParams {
	return getMTRDeviceEnergyManagementClusterResumeRequestParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumeRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterResumeRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumeRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterResumeRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumeRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterResumeRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumeRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterResumeRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



