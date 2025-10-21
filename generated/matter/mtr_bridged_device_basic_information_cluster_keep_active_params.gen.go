// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBridgedDeviceBasicInformationClusterKeepActiveParams] class.
var (
	MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass     _MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass
	MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClassOnce sync.Once
)

func getMTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass() _MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass {
	MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClassOnce.Do(func() {
		MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass = _MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass{objc.GetClass("MTRBridgedDeviceBasicInformationClusterKeepActiveParams")}
	})
	return MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass
}

type _MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBridgedDeviceBasicInformationClusterKeepActiveParams] class.
type IMTRBridgedDeviceBasicInformationClusterKeepActiveParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams
type MTRBridgedDeviceBasicInformationClusterKeepActiveParams struct {
	objectivec.Object
}

// MTRBridgedDeviceBasicInformationClusterKeepActiveParamsFrom constructs a [MTRBridgedDeviceBasicInformationClusterKeepActiveParams] from an unsafe.Pointer.
func MTRBridgedDeviceBasicInformationClusterKeepActiveParamsFrom(ptr unsafe.Pointer) MTRBridgedDeviceBasicInformationClusterKeepActiveParams {
	return MTRBridgedDeviceBasicInformationClusterKeepActiveParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass) Alloc() MTRBridgedDeviceBasicInformationClusterKeepActiveParams {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterKeepActiveParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass) New() MTRBridgedDeviceBasicInformationClusterKeepActiveParams {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterKeepActiveParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) Init() MTRBridgedDeviceBasicInformationClusterKeepActiveParams {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterKeepActiveParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) Autorelease() MTRBridgedDeviceBasicInformationClusterKeepActiveParams {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterKeepActiveParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBridgedDeviceBasicInformationClusterKeepActiveParams creates a new MTRBridgedDeviceBasicInformationClusterKeepActiveParams instance.
func NewMTRBridgedDeviceBasicInformationClusterKeepActiveParams() MTRBridgedDeviceBasicInformationClusterKeepActiveParams {
	return getMTRBridgedDeviceBasicInformationClusterKeepActiveParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/serverSideProcessingTimeout
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/serverSideProcessingTimeout
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/stayActiveDuration
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) StayActiveDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("stayActiveDuration"))
	return rv
}


// SetStayActiveDuration sets the value of the stayActiveDuration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/stayActiveDuration
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) SetStayActiveDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStayActiveDuration:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/timedInvokeTimeoutMs
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/timedInvokeTimeoutMs
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/timeoutMs
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) TimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timeoutMs"))
	return rv
}


// SetTimeoutMs sets the value of the timeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/timeoutMs
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) SetTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeoutMs:"), value)
}


