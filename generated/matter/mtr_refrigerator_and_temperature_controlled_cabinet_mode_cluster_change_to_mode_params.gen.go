// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams] class.
var (
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass     _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass() _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass {
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass = _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass{objc.GetClass("MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams")}
	})
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass
}

type _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams] class.
type IMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams interface {
	objectivec.IObject
	// properties:
	NewMode() objc.IObject /* cross-framework: NSNumber */
	SetNewMode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams
type MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsFrom constructs a [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams {
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass) Alloc() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass) New() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) Init() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) Autorelease() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams creates a new MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams instance.
func NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams {
	return getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams/newMode
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) NewMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams/newMode
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) SetNewMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



