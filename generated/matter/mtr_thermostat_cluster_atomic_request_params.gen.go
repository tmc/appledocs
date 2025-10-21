// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThermostatClusterAtomicRequestParams] class.
var (
	MTRThermostatClusterAtomicRequestParamsClass     _MTRThermostatClusterAtomicRequestParamsClass
	MTRThermostatClusterAtomicRequestParamsClassOnce sync.Once
)

func getMTRThermostatClusterAtomicRequestParamsClass() _MTRThermostatClusterAtomicRequestParamsClass {
	MTRThermostatClusterAtomicRequestParamsClassOnce.Do(func() {
		MTRThermostatClusterAtomicRequestParamsClass = _MTRThermostatClusterAtomicRequestParamsClass{objc.GetClass("MTRThermostatClusterAtomicRequestParams")}
	})
	return MTRThermostatClusterAtomicRequestParamsClass
}

type _MTRThermostatClusterAtomicRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterAtomicRequestParams] class.
type IMTRThermostatClusterAtomicRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicRequestParams
type MTRThermostatClusterAtomicRequestParams struct {
	objectivec.Object
}

// MTRThermostatClusterAtomicRequestParamsFrom constructs a [MTRThermostatClusterAtomicRequestParams] from an unsafe.Pointer.
func MTRThermostatClusterAtomicRequestParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterAtomicRequestParams {
	return MTRThermostatClusterAtomicRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterAtomicRequestParamsClass) Alloc() MTRThermostatClusterAtomicRequestParams {
	rv := objc.Send[MTRThermostatClusterAtomicRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterAtomicRequestParamsClass) New() MTRThermostatClusterAtomicRequestParams {
	rv := objc.Send[MTRThermostatClusterAtomicRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterAtomicRequestParams) Init() MTRThermostatClusterAtomicRequestParams {
	rv := objc.Send[MTRThermostatClusterAtomicRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterAtomicRequestParams) Autorelease() MTRThermostatClusterAtomicRequestParams {
	rv := objc.Send[MTRThermostatClusterAtomicRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterAtomicRequestParams creates a new MTRThermostatClusterAtomicRequestParams instance.
func NewMTRThermostatClusterAtomicRequestParams() MTRThermostatClusterAtomicRequestParams {
	return getMTRThermostatClusterAtomicRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicRequestParams/attributeRequests
func (m_ MTRThermostatClusterAtomicRequestParams) AttributeRequests() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("attributeRequests"))
	return rv
}


// SetAttributeRequests sets the value of the attributeRequests property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicRequestParams/attributeRequests
func (m_ MTRThermostatClusterAtomicRequestParams) SetAttributeRequests(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributeRequests:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicRequestParams/requestType
func (m_ MTRThermostatClusterAtomicRequestParams) RequestType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("requestType"))
	return rv
}


// SetRequestType sets the value of the requestType property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicRequestParams/requestType
func (m_ MTRThermostatClusterAtomicRequestParams) SetRequestType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestType:"), value)
}
// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicRequestParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterAtomicRequestParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicRequestParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterAtomicRequestParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicRequestParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterAtomicRequestParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicRequestParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterAtomicRequestParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicRequestParams/timeout
func (m_ MTRThermostatClusterAtomicRequestParams) Timeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timeout"))
	return rv
}


// SetTimeout sets the value of the timeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterAtomicRequestParams/timeout
func (m_ MTRThermostatClusterAtomicRequestParams) SetTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeout:"), value)
}


