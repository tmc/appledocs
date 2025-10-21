// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterPauseActionWithDurationParams] class.
var (
	MTRActionsClusterPauseActionWithDurationParamsClass     _MTRActionsClusterPauseActionWithDurationParamsClass
	MTRActionsClusterPauseActionWithDurationParamsClassOnce sync.Once
)

func getMTRActionsClusterPauseActionWithDurationParamsClass() _MTRActionsClusterPauseActionWithDurationParamsClass {
	MTRActionsClusterPauseActionWithDurationParamsClassOnce.Do(func() {
		MTRActionsClusterPauseActionWithDurationParamsClass = _MTRActionsClusterPauseActionWithDurationParamsClass{objc.GetClass("MTRActionsClusterPauseActionWithDurationParams")}
	})
	return MTRActionsClusterPauseActionWithDurationParamsClass
}

type _MTRActionsClusterPauseActionWithDurationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterPauseActionWithDurationParams] class.
type IMTRActionsClusterPauseActionWithDurationParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionWithDurationParams
type MTRActionsClusterPauseActionWithDurationParams struct {
	objectivec.Object
}

// MTRActionsClusterPauseActionWithDurationParamsFrom constructs a [MTRActionsClusterPauseActionWithDurationParams] from an unsafe.Pointer.
func MTRActionsClusterPauseActionWithDurationParamsFrom(ptr unsafe.Pointer) MTRActionsClusterPauseActionWithDurationParams {
	return MTRActionsClusterPauseActionWithDurationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterPauseActionWithDurationParamsClass) Alloc() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterPauseActionWithDurationParamsClass) New() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterPauseActionWithDurationParams) Init() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterPauseActionWithDurationParams) Autorelease() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterPauseActionWithDurationParams creates a new MTRActionsClusterPauseActionWithDurationParams instance.
func NewMTRActionsClusterPauseActionWithDurationParams() MTRActionsClusterPauseActionWithDurationParams {
	return getMTRActionsClusterPauseActionWithDurationParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/actionid
func (m_ MTRActionsClusterPauseActionWithDurationParams) ActionID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("actionID"))
	return rv
}


// SetActionID sets the value of the actionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/actionid
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetActionID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/duration
func (m_ MTRActionsClusterPauseActionWithDurationParams) Duration() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/duration
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetDuration(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/invokeid
func (m_ MTRActionsClusterPauseActionWithDurationParams) InvokeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("invokeID"))
	return rv
}


// SetInvokeID sets the value of the invokeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/invokeid
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetInvokeID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/serversideprocessingtimeout
func (m_ MTRActionsClusterPauseActionWithDurationParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/serversideprocessingtimeout
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/timedinvoketimeoutms
func (m_ MTRActionsClusterPauseActionWithDurationParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/timedinvoketimeoutms
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



