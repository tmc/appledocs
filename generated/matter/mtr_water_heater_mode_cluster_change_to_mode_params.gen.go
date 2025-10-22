// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWaterHeaterModeClusterChangeToModeParams] class.
var (
	MTRWaterHeaterModeClusterChangeToModeParamsClass     _MTRWaterHeaterModeClusterChangeToModeParamsClass
	MTRWaterHeaterModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTRWaterHeaterModeClusterChangeToModeParamsClass() _MTRWaterHeaterModeClusterChangeToModeParamsClass {
	MTRWaterHeaterModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTRWaterHeaterModeClusterChangeToModeParamsClass = _MTRWaterHeaterModeClusterChangeToModeParamsClass{objc.GetClass("MTRWaterHeaterModeClusterChangeToModeParams")}
	})
	return MTRWaterHeaterModeClusterChangeToModeParamsClass
}

type _MTRWaterHeaterModeClusterChangeToModeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWaterHeaterModeClusterChangeToModeParams] class.
type IMTRWaterHeaterModeClusterChangeToModeParams interface {
	objectivec.IObject
	NewMode() foundation.Number
	SetNewMode(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeParams
type MTRWaterHeaterModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTRWaterHeaterModeClusterChangeToModeParamsFrom constructs a [MTRWaterHeaterModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTRWaterHeaterModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTRWaterHeaterModeClusterChangeToModeParams {
	return MTRWaterHeaterModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterModeClusterChangeToModeParamsClass) Alloc() MTRWaterHeaterModeClusterChangeToModeParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWaterHeaterModeClusterChangeToModeParamsClass) New() MTRWaterHeaterModeClusterChangeToModeParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) Init() MTRWaterHeaterModeClusterChangeToModeParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) Autorelease() MTRWaterHeaterModeClusterChangeToModeParams {
	rv := objc.Send[MTRWaterHeaterModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterModeClusterChangeToModeParams creates a new MTRWaterHeaterModeClusterChangeToModeParams instance.
func NewMTRWaterHeaterModeClusterChangeToModeParams() MTRWaterHeaterModeClusterChangeToModeParams {
	return getMTRWaterHeaterModeClusterChangeToModeParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeParams/newMode
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) NewMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("newMode"))
	return rv
}


// SetNewMode sets the value of the newMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeParams/newMode
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) SetNewMode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTRWaterHeaterModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



