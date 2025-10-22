// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRModeSelectClusterChangeToModeParams] class.
var (
	MTRModeSelectClusterChangeToModeParamsClass     _MTRModeSelectClusterChangeToModeParamsClass
	MTRModeSelectClusterChangeToModeParamsClassOnce sync.Once
)

func getMTRModeSelectClusterChangeToModeParamsClass() _MTRModeSelectClusterChangeToModeParamsClass {
	MTRModeSelectClusterChangeToModeParamsClassOnce.Do(func() {
		MTRModeSelectClusterChangeToModeParamsClass = _MTRModeSelectClusterChangeToModeParamsClass{objc.GetClass("MTRModeSelectClusterChangeToModeParams")}
	})
	return MTRModeSelectClusterChangeToModeParamsClass
}

type _MTRModeSelectClusterChangeToModeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRModeSelectClusterChangeToModeParams] class.
type IMTRModeSelectClusterChangeToModeParams interface {
	objectivec.IObject
	NewMode() foundation.Number
	SetNewMode(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRModeSelectClusterChangeToModeParams
type MTRModeSelectClusterChangeToModeParams struct {
	objectivec.Object
}

// MTRModeSelectClusterChangeToModeParamsFrom constructs a [MTRModeSelectClusterChangeToModeParams] from an unsafe.Pointer.
func MTRModeSelectClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTRModeSelectClusterChangeToModeParams {
	return MTRModeSelectClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRModeSelectClusterChangeToModeParamsClass) Alloc() MTRModeSelectClusterChangeToModeParams {
	rv := objc.Send[MTRModeSelectClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRModeSelectClusterChangeToModeParamsClass) New() MTRModeSelectClusterChangeToModeParams {
	rv := objc.Send[MTRModeSelectClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRModeSelectClusterChangeToModeParams) Init() MTRModeSelectClusterChangeToModeParams {
	rv := objc.Send[MTRModeSelectClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRModeSelectClusterChangeToModeParams) Autorelease() MTRModeSelectClusterChangeToModeParams {
	rv := objc.Send[MTRModeSelectClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRModeSelectClusterChangeToModeParams creates a new MTRModeSelectClusterChangeToModeParams instance.
func NewMTRModeSelectClusterChangeToModeParams() MTRModeSelectClusterChangeToModeParams {
	return getMTRModeSelectClusterChangeToModeParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclusterchangetomodeparams/newmode
func (m_ MTRModeSelectClusterChangeToModeParams) NewMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("newMode"))
	return rv
}


// SetNewMode sets the value of the newMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclusterchangetomodeparams/newmode
func (m_ MTRModeSelectClusterChangeToModeParams) SetNewMode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRModeSelectClusterChangeToModeParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRModeSelectClusterChangeToModeParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRModeSelectClusterChangeToModeParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRModeSelectClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



