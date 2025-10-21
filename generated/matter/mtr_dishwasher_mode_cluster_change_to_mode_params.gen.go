// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDishwasherModeClusterChangeToModeParams] class.
var (
	MTRDishwasherModeClusterChangeToModeParamsClass     _MTRDishwasherModeClusterChangeToModeParamsClass
	MTRDishwasherModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTRDishwasherModeClusterChangeToModeParamsClass() _MTRDishwasherModeClusterChangeToModeParamsClass {
	MTRDishwasherModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTRDishwasherModeClusterChangeToModeParamsClass = _MTRDishwasherModeClusterChangeToModeParamsClass{objc.GetClass("MTRDishwasherModeClusterChangeToModeParams")}
	})
	return MTRDishwasherModeClusterChangeToModeParamsClass
}

type _MTRDishwasherModeClusterChangeToModeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDishwasherModeClusterChangeToModeParams] class.
type IMTRDishwasherModeClusterChangeToModeParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeParams
type MTRDishwasherModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTRDishwasherModeClusterChangeToModeParamsFrom constructs a [MTRDishwasherModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTRDishwasherModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTRDishwasherModeClusterChangeToModeParams {
	return MTRDishwasherModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDishwasherModeClusterChangeToModeParamsClass) Alloc() MTRDishwasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDishwasherModeClusterChangeToModeParamsClass) New() MTRDishwasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDishwasherModeClusterChangeToModeParams) Init() MTRDishwasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDishwasherModeClusterChangeToModeParams) Autorelease() MTRDishwasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRDishwasherModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDishwasherModeClusterChangeToModeParams creates a new MTRDishwasherModeClusterChangeToModeParams instance.
func NewMTRDishwasherModeClusterChangeToModeParams() MTRDishwasherModeClusterChangeToModeParams {
	return getMTRDishwasherModeClusterChangeToModeParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeParams/newMode
func (m_ MTRDishwasherModeClusterChangeToModeParams) NewMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("newMode"))
	return rv
}


// SetNewMode sets the value of the newMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeParams/newMode
func (m_ MTRDishwasherModeClusterChangeToModeParams) SetNewMode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTRDishwasherModeClusterChangeToModeParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTRDishwasherModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTRDishwasherModeClusterChangeToModeParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTRDishwasherModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



