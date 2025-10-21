// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREnergyEVSEClusterDisableParams] class.
var (
	MTREnergyEVSEClusterDisableParamsClass     _MTREnergyEVSEClusterDisableParamsClass
	MTREnergyEVSEClusterDisableParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterDisableParamsClass() _MTREnergyEVSEClusterDisableParamsClass {
	MTREnergyEVSEClusterDisableParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterDisableParamsClass = _MTREnergyEVSEClusterDisableParamsClass{objc.GetClass("MTREnergyEVSEClusterDisableParams")}
	})
	return MTREnergyEVSEClusterDisableParamsClass
}

type _MTREnergyEVSEClusterDisableParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterDisableParams] class.
type IMTREnergyEVSEClusterDisableParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterDisableParams
type MTREnergyEVSEClusterDisableParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterDisableParamsFrom constructs a [MTREnergyEVSEClusterDisableParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterDisableParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterDisableParams {
	return MTREnergyEVSEClusterDisableParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterDisableParamsClass) Alloc() MTREnergyEVSEClusterDisableParams {
	rv := objc.Send[MTREnergyEVSEClusterDisableParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterDisableParamsClass) New() MTREnergyEVSEClusterDisableParams {
	rv := objc.Send[MTREnergyEVSEClusterDisableParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterDisableParams) Init() MTREnergyEVSEClusterDisableParams {
	rv := objc.Send[MTREnergyEVSEClusterDisableParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterDisableParams) Autorelease() MTREnergyEVSEClusterDisableParams {
	rv := objc.Send[MTREnergyEVSEClusterDisableParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterDisableParams creates a new MTREnergyEVSEClusterDisableParams instance.
func NewMTREnergyEVSEClusterDisableParams() MTREnergyEVSEClusterDisableParams {
	return getMTREnergyEVSEClusterDisableParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterDisableParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterDisableParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterDisableParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterDisableParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterDisableParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterDisableParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterDisableParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterDisableParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



