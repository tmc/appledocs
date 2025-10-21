// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralDiagnosticsClusterPayloadTestRequestParams] class.
var (
	MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass     _MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass
	MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass() _MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass {
	MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass = _MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass{objc.GetClass("MTRGeneralDiagnosticsClusterPayloadTestRequestParams")}
	})
	return MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass
}

type _MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralDiagnosticsClusterPayloadTestRequestParams] class.
type IMTRGeneralDiagnosticsClusterPayloadTestRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams
type MTRGeneralDiagnosticsClusterPayloadTestRequestParams struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterPayloadTestRequestParamsFrom constructs a [MTRGeneralDiagnosticsClusterPayloadTestRequestParams] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterPayloadTestRequestParamsFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterPayloadTestRequestParams {
	return MTRGeneralDiagnosticsClusterPayloadTestRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass) Alloc() MTRGeneralDiagnosticsClusterPayloadTestRequestParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass) New() MTRGeneralDiagnosticsClusterPayloadTestRequestParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) Init() MTRGeneralDiagnosticsClusterPayloadTestRequestParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) Autorelease() MTRGeneralDiagnosticsClusterPayloadTestRequestParams {
	rv := objc.Send[MTRGeneralDiagnosticsClusterPayloadTestRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterPayloadTestRequestParams creates a new MTRGeneralDiagnosticsClusterPayloadTestRequestParams instance.
func NewMTRGeneralDiagnosticsClusterPayloadTestRequestParams() MTRGeneralDiagnosticsClusterPayloadTestRequestParams {
	return getMTRGeneralDiagnosticsClusterPayloadTestRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams/count
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) Count() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("count"))
	return rv
}


// SetCount sets the value of the count property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams/count
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) SetCount(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams/enableKey
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) EnableKey() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("enableKey"))
	return rv
}


// SetEnableKey sets the value of the enableKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams/enableKey
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) SetEnableKey(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnableKey:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams/serverSideProcessingTimeout
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams/serverSideProcessingTimeout
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams/timedInvokeTimeoutMs
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams/timedInvokeTimeoutMs
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams/value
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) Value() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterPayloadTestRequestParams/value
func (m_ MTRGeneralDiagnosticsClusterPayloadTestRequestParams) SetValue(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



