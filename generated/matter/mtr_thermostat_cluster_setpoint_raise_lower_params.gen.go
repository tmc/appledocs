// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThermostatClusterSetpointRaiseLowerParams] class.
var (
	MTRThermostatClusterSetpointRaiseLowerParamsClass     _MTRThermostatClusterSetpointRaiseLowerParamsClass
	MTRThermostatClusterSetpointRaiseLowerParamsClassOnce sync.Once
)

func getMTRThermostatClusterSetpointRaiseLowerParamsClass() _MTRThermostatClusterSetpointRaiseLowerParamsClass {
	MTRThermostatClusterSetpointRaiseLowerParamsClassOnce.Do(func() {
		MTRThermostatClusterSetpointRaiseLowerParamsClass = _MTRThermostatClusterSetpointRaiseLowerParamsClass{objc.GetClass("MTRThermostatClusterSetpointRaiseLowerParams")}
	})
	return MTRThermostatClusterSetpointRaiseLowerParamsClass
}

type _MTRThermostatClusterSetpointRaiseLowerParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterSetpointRaiseLowerParams] class.
type IMTRThermostatClusterSetpointRaiseLowerParams interface {
	objectivec.IObject
	Amount() foundation.Number
	SetAmount(value foundation.INumber)
	Mode() foundation.Number
	SetMode(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetpointRaiseLowerParams
type MTRThermostatClusterSetpointRaiseLowerParams struct {
	objectivec.Object
}

// MTRThermostatClusterSetpointRaiseLowerParamsFrom constructs a [MTRThermostatClusterSetpointRaiseLowerParams] from an unsafe.Pointer.
func MTRThermostatClusterSetpointRaiseLowerParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterSetpointRaiseLowerParams {
	return MTRThermostatClusterSetpointRaiseLowerParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterSetpointRaiseLowerParamsClass) Alloc() MTRThermostatClusterSetpointRaiseLowerParams {
	rv := objc.Send[MTRThermostatClusterSetpointRaiseLowerParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterSetpointRaiseLowerParamsClass) New() MTRThermostatClusterSetpointRaiseLowerParams {
	rv := objc.Send[MTRThermostatClusterSetpointRaiseLowerParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) Init() MTRThermostatClusterSetpointRaiseLowerParams {
	rv := objc.Send[MTRThermostatClusterSetpointRaiseLowerParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) Autorelease() MTRThermostatClusterSetpointRaiseLowerParams {
	rv := objc.Send[MTRThermostatClusterSetpointRaiseLowerParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterSetpointRaiseLowerParams creates a new MTRThermostatClusterSetpointRaiseLowerParams instance.
func NewMTRThermostatClusterSetpointRaiseLowerParams() MTRThermostatClusterSetpointRaiseLowerParams {
	return getMTRThermostatClusterSetpointRaiseLowerParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/amount
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) Amount() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("amount"))
	return rv
}


// SetAmount sets the value of the amount property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/amount
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) SetAmount(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAmount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/mode
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) Mode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/mode
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) SetMode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/serversideprocessingtimeout
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/serversideprocessingtimeout
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetpointraiselowerparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



