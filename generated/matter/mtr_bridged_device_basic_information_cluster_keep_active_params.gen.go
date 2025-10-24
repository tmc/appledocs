// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StayActiveDuration() objc.IObject /* cross-framework: NSNumber */
	SetStayActiveDuration(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/serverSideProcessingTimeout
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/serverSideProcessingTimeout
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/stayActiveDuration
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) StayActiveDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stayActiveDuration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/stayActiveDuration
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) SetStayActiveDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStayActiveDuration:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/timedInvokeTimeoutMs
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/timedInvokeTimeoutMs
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/timeoutMs
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) TimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterKeepActiveParams/timeoutMs
func (m_ MTRBridgedDeviceBasicInformationClusterKeepActiveParams) SetTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeoutMs:"), value)
}



