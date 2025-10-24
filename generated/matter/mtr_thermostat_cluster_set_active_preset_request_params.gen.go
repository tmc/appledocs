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
	// properties:
	PresetHandle() objc.IObject /* cross-framework: NSData */
	SetPresetHandle(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/presetHandle
func (m_ MTRThermostatClusterSetActivePresetRequestParams) PresetHandle() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("presetHandle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/presetHandle
func (m_ MTRThermostatClusterSetActivePresetRequestParams) SetPresetHandle(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetHandle:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterSetActivePresetRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterSetActivePresetRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterSetActivePresetRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActivePresetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterSetActivePresetRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



