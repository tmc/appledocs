// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROvenModeClusterChangeToModeParams] class.
var (
	MTROvenModeClusterChangeToModeParamsClass     _MTROvenModeClusterChangeToModeParamsClass
	MTROvenModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTROvenModeClusterChangeToModeParamsClass() _MTROvenModeClusterChangeToModeParamsClass {
	MTROvenModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTROvenModeClusterChangeToModeParamsClass = _MTROvenModeClusterChangeToModeParamsClass{objc.GetClass("MTROvenModeClusterChangeToModeParams")}
	})
	return MTROvenModeClusterChangeToModeParamsClass
}

type _MTROvenModeClusterChangeToModeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROvenModeClusterChangeToModeParams] class.
type IMTROvenModeClusterChangeToModeParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams
type MTROvenModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTROvenModeClusterChangeToModeParamsFrom constructs a [MTROvenModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTROvenModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTROvenModeClusterChangeToModeParams {
	return MTROvenModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROvenModeClusterChangeToModeParamsClass) Alloc() MTROvenModeClusterChangeToModeParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROvenModeClusterChangeToModeParamsClass) New() MTROvenModeClusterChangeToModeParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenModeClusterChangeToModeParams) Init() MTROvenModeClusterChangeToModeParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenModeClusterChangeToModeParams) Autorelease() MTROvenModeClusterChangeToModeParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenModeClusterChangeToModeParams creates a new MTROvenModeClusterChangeToModeParams instance.
func NewMTROvenModeClusterChangeToModeParams() MTROvenModeClusterChangeToModeParams {
	return getMTROvenModeClusterChangeToModeParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams/newMode
func (m_ MTROvenModeClusterChangeToModeParams) NewMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("newMode"))
	return rv
}


// SetNewMode sets the value of the newMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams/newMode
func (m_ MTROvenModeClusterChangeToModeParams) SetNewMode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTROvenModeClusterChangeToModeParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTROvenModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTROvenModeClusterChangeToModeParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTROvenModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



