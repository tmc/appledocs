// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThermostatClusterSetActiveScheduleRequestParams] class.
var (
	MTRThermostatClusterSetActiveScheduleRequestParamsClass     _MTRThermostatClusterSetActiveScheduleRequestParamsClass
	MTRThermostatClusterSetActiveScheduleRequestParamsClassOnce sync.Once
)

func getMTRThermostatClusterSetActiveScheduleRequestParamsClass() _MTRThermostatClusterSetActiveScheduleRequestParamsClass {
	MTRThermostatClusterSetActiveScheduleRequestParamsClassOnce.Do(func() {
		MTRThermostatClusterSetActiveScheduleRequestParamsClass = _MTRThermostatClusterSetActiveScheduleRequestParamsClass{objc.GetClass("MTRThermostatClusterSetActiveScheduleRequestParams")}
	})
	return MTRThermostatClusterSetActiveScheduleRequestParamsClass
}

type _MTRThermostatClusterSetActiveScheduleRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterSetActiveScheduleRequestParams] class.
type IMTRThermostatClusterSetActiveScheduleRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActiveScheduleRequestParams
type MTRThermostatClusterSetActiveScheduleRequestParams struct {
	objectivec.Object
}

// MTRThermostatClusterSetActiveScheduleRequestParamsFrom constructs a [MTRThermostatClusterSetActiveScheduleRequestParams] from an unsafe.Pointer.
func MTRThermostatClusterSetActiveScheduleRequestParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterSetActiveScheduleRequestParams {
	return MTRThermostatClusterSetActiveScheduleRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterSetActiveScheduleRequestParamsClass) Alloc() MTRThermostatClusterSetActiveScheduleRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActiveScheduleRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterSetActiveScheduleRequestParamsClass) New() MTRThermostatClusterSetActiveScheduleRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActiveScheduleRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) Init() MTRThermostatClusterSetActiveScheduleRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActiveScheduleRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) Autorelease() MTRThermostatClusterSetActiveScheduleRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActiveScheduleRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterSetActiveScheduleRequestParams creates a new MTRThermostatClusterSetActiveScheduleRequestParams instance.
func NewMTRThermostatClusterSetActiveScheduleRequestParams() MTRThermostatClusterSetActiveScheduleRequestParams {
	return getMTRThermostatClusterSetActiveScheduleRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActiveScheduleRequestParams/scheduleHandle
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) ScheduleHandle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("scheduleHandle"))
	return rv
}


// SetScheduleHandle sets the value of the scheduleHandle property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActiveScheduleRequestParams/scheduleHandle
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) SetScheduleHandle(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScheduleHandle:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActiveScheduleRequestParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActiveScheduleRequestParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActiveScheduleRequestParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActiveScheduleRequestParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



