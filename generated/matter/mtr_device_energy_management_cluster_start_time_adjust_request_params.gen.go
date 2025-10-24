// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams] class.
var (
	MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsClass     _MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsClass
	MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsClass() _MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsClass {
	MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsClass = _MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams")}
	})
	return MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsClass
}

type _MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams] class.
type IMTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams interface {
	objectivec.IObject
	// properties:
	Cause() objc.IObject /* cross-framework: NSNumber */
	SetCause(value objc.IObject /* cross-framework: NSNumber */)
	RequestedStartTime() objc.IObject /* cross-framework: NSNumber */
	SetRequestedStartTime(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams
type MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsFrom constructs a [MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams {
	return MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsClass) Alloc() MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsClass) New() MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams) Init() MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams) Autorelease() MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams creates a new MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams instance.
func NewMTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams() MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams {
	return getMTRDeviceEnergyManagementClusterStartTimeAdjustRequestParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams/cause
func (m_ MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams) Cause() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cause"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams/cause
func (m_ MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams) SetCause(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams/requestedStartTime
func (m_ MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams) RequestedStartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("requestedStartTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams/requestedStartTime
func (m_ MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams) SetRequestedStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestedStartTime:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



