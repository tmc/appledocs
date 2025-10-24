// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementClusterPauseRequestParams] class.
var (
	MTRDeviceEnergyManagementClusterPauseRequestParamsClass     _MTRDeviceEnergyManagementClusterPauseRequestParamsClass
	MTRDeviceEnergyManagementClusterPauseRequestParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterPauseRequestParamsClass() _MTRDeviceEnergyManagementClusterPauseRequestParamsClass {
	MTRDeviceEnergyManagementClusterPauseRequestParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterPauseRequestParamsClass = _MTRDeviceEnergyManagementClusterPauseRequestParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterPauseRequestParams")}
	})
	return MTRDeviceEnergyManagementClusterPauseRequestParamsClass
}

type _MTRDeviceEnergyManagementClusterPauseRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterPauseRequestParams] class.
type IMTRDeviceEnergyManagementClusterPauseRequestParams interface {
	objectivec.IObject
	// properties:
	Cause() objc.IObject /* cross-framework: NSNumber */
	SetCause(value objc.IObject /* cross-framework: NSNumber */)
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPauseRequestParams
type MTRDeviceEnergyManagementClusterPauseRequestParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterPauseRequestParamsFrom constructs a [MTRDeviceEnergyManagementClusterPauseRequestParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterPauseRequestParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterPauseRequestParams {
	return MTRDeviceEnergyManagementClusterPauseRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterPauseRequestParamsClass) Alloc() MTRDeviceEnergyManagementClusterPauseRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPauseRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterPauseRequestParamsClass) New() MTRDeviceEnergyManagementClusterPauseRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPauseRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) Init() MTRDeviceEnergyManagementClusterPauseRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPauseRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) Autorelease() MTRDeviceEnergyManagementClusterPauseRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPauseRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterPauseRequestParams creates a new MTRDeviceEnergyManagementClusterPauseRequestParams instance.
func NewMTRDeviceEnergyManagementClusterPauseRequestParams() MTRDeviceEnergyManagementClusterPauseRequestParams {
	return getMTRDeviceEnergyManagementClusterPauseRequestParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPauseRequestParams/cause
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) Cause() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cause"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPauseRequestParams/cause
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) SetCause(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPauseRequestParams/duration
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPauseRequestParams/duration
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPauseRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPauseRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPauseRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPauseRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



