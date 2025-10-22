// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThermostatClusterSetActivePresetRequestParams] class.
var (
	MTRThermostatClusterSetActivePresetRequestParamsClass     _MTRThermostatClusterSetActivePresetRequestParamsClass
	MTRThermostatClusterSetActivePresetRequestParamsClassOnce sync.Once
)

func getMTRThermostatClusterSetActivePresetRequestParamsClass() _MTRThermostatClusterSetActivePresetRequestParamsClass {
	MTRThermostatClusterSetActivePresetRequestParamsClassOnce.Do(func() {
		MTRThermostatClusterSetActivePresetRequestParamsClass = _MTRThermostatClusterSetActivePresetRequestParamsClass{objc.GetClass("MTRThermostatClusterSetActivePresetRequestParams")}
	})
	return MTRThermostatClusterSetActivePresetRequestParamsClass
}

type _MTRThermostatClusterSetActivePresetRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterSetActivePresetRequestParams] class.
type IMTRThermostatClusterSetActivePresetRequestParams interface {
	objectivec.IObject
	PresetHandle() foundation.NSData
	SetPresetHandle(value foundation.IData)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams
type MTRThermostatClusterSetActivePresetRequestParams struct {
	objectivec.Object
}

// MTRThermostatClusterSetActivePresetRequestParamsFrom constructs a [MTRThermostatClusterSetActivePresetRequestParams] from an unsafe.Pointer.
func MTRThermostatClusterSetActivePresetRequestParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterSetActivePresetRequestParams {
	return MTRThermostatClusterSetActivePresetRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterSetActivePresetRequestParamsClass) Alloc() MTRThermostatClusterSetActivePresetRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActivePresetRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterSetActivePresetRequestParamsClass) New() MTRThermostatClusterSetActivePresetRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActivePresetRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterSetActivePresetRequestParams) Init() MTRThermostatClusterSetActivePresetRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActivePresetRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterSetActivePresetRequestParams) Autorelease() MTRThermostatClusterSetActivePresetRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActivePresetRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterSetActivePresetRequestParams creates a new MTRThermostatClusterSetActivePresetRequestParams instance.
func NewMTRThermostatClusterSetActivePresetRequestParams() MTRThermostatClusterSetActivePresetRequestParams {
	return getMTRThermostatClusterSetActivePresetRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/presetHandle
func (m_ MTRThermostatClusterSetActivePresetRequestParams) PresetHandle() foundation.NSData {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("presetHandle"))
	return rv
}


// SetPresetHandle sets the value of the presetHandle property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/presetHandle
func (m_ MTRThermostatClusterSetActivePresetRequestParams) SetPresetHandle(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetHandle:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterSetActivePresetRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterSetActivePresetRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterSetActivePresetRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterSetActivePresetRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



