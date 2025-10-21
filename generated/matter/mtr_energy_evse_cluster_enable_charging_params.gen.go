// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTREnergyEVSEClusterEnableChargingParams] class.
var (
	MTREnergyEVSEClusterEnableChargingParamsClass     _MTREnergyEVSEClusterEnableChargingParamsClass
	MTREnergyEVSEClusterEnableChargingParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterEnableChargingParamsClass() _MTREnergyEVSEClusterEnableChargingParamsClass {
	MTREnergyEVSEClusterEnableChargingParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterEnableChargingParamsClass = _MTREnergyEVSEClusterEnableChargingParamsClass{objc.GetClass("MTREnergyEVSEClusterEnableChargingParams")}
	})
	return MTREnergyEVSEClusterEnableChargingParamsClass
}

type _MTREnergyEVSEClusterEnableChargingParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterEnableChargingParams] class.
type IMTREnergyEVSEClusterEnableChargingParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnableChargingParams
type MTREnergyEVSEClusterEnableChargingParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterEnableChargingParamsFrom constructs a [MTREnergyEVSEClusterEnableChargingParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterEnableChargingParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterEnableChargingParams {
	return MTREnergyEVSEClusterEnableChargingParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterEnableChargingParamsClass) Alloc() MTREnergyEVSEClusterEnableChargingParams {
	rv := objc.Send[MTREnergyEVSEClusterEnableChargingParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterEnableChargingParamsClass) New() MTREnergyEVSEClusterEnableChargingParams {
	rv := objc.Send[MTREnergyEVSEClusterEnableChargingParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterEnableChargingParams) Init() MTREnergyEVSEClusterEnableChargingParams {
	rv := objc.Send[MTREnergyEVSEClusterEnableChargingParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterEnableChargingParams) Autorelease() MTREnergyEVSEClusterEnableChargingParams {
	rv := objc.Send[MTREnergyEVSEClusterEnableChargingParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterEnableChargingParams creates a new MTREnergyEVSEClusterEnableChargingParams instance.
func NewMTREnergyEVSEClusterEnableChargingParams() MTREnergyEVSEClusterEnableChargingParams {
	return getMTREnergyEVSEClusterEnableChargingParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnableChargingParams/chargingEnabledUntil
func (m_ MTREnergyEVSEClusterEnableChargingParams) ChargingEnabledUntil() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("chargingEnabledUntil"))
	return rv
}


// SetChargingEnabledUntil sets the value of the chargingEnabledUntil property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnableChargingParams/chargingEnabledUntil
func (m_ MTREnergyEVSEClusterEnableChargingParams) SetChargingEnabledUntil(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChargingEnabledUntil:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnableChargingParams/maximumChargeCurrent
func (m_ MTREnergyEVSEClusterEnableChargingParams) MaximumChargeCurrent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("maximumChargeCurrent"))
	return rv
}


// SetMaximumChargeCurrent sets the value of the maximumChargeCurrent property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnableChargingParams/maximumChargeCurrent
func (m_ MTREnergyEVSEClusterEnableChargingParams) SetMaximumChargeCurrent(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumChargeCurrent:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnableChargingParams/minimumChargeCurrent
func (m_ MTREnergyEVSEClusterEnableChargingParams) MinimumChargeCurrent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("minimumChargeCurrent"))
	return rv
}


// SetMinimumChargeCurrent sets the value of the minimumChargeCurrent property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnableChargingParams/minimumChargeCurrent
func (m_ MTREnergyEVSEClusterEnableChargingParams) SetMinimumChargeCurrent(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimumChargeCurrent:"), value)
}
// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnableChargingParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterEnableChargingParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnableChargingParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterEnableChargingParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnableChargingParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterEnableChargingParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnableChargingParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterEnableChargingParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


