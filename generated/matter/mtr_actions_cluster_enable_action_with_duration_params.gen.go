// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterEnableActionWithDurationParams] class.
var (
	MTRActionsClusterEnableActionWithDurationParamsClass     _MTRActionsClusterEnableActionWithDurationParamsClass
	MTRActionsClusterEnableActionWithDurationParamsClassOnce sync.Once
)

func getMTRActionsClusterEnableActionWithDurationParamsClass() _MTRActionsClusterEnableActionWithDurationParamsClass {
	MTRActionsClusterEnableActionWithDurationParamsClassOnce.Do(func() {
		MTRActionsClusterEnableActionWithDurationParamsClass = _MTRActionsClusterEnableActionWithDurationParamsClass{objc.GetClass("MTRActionsClusterEnableActionWithDurationParams")}
	})
	return MTRActionsClusterEnableActionWithDurationParamsClass
}

type _MTRActionsClusterEnableActionWithDurationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterEnableActionWithDurationParams] class.
type IMTRActionsClusterEnableActionWithDurationParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionWithDurationParams
type MTRActionsClusterEnableActionWithDurationParams struct {
	objectivec.Object
}

// MTRActionsClusterEnableActionWithDurationParamsFrom constructs a [MTRActionsClusterEnableActionWithDurationParams] from an unsafe.Pointer.
func MTRActionsClusterEnableActionWithDurationParamsFrom(ptr unsafe.Pointer) MTRActionsClusterEnableActionWithDurationParams {
	return MTRActionsClusterEnableActionWithDurationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterEnableActionWithDurationParamsClass) Alloc() MTRActionsClusterEnableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterEnableActionWithDurationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterEnableActionWithDurationParamsClass) New() MTRActionsClusterEnableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterEnableActionWithDurationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterEnableActionWithDurationParams) Init() MTRActionsClusterEnableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterEnableActionWithDurationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterEnableActionWithDurationParams) Autorelease() MTRActionsClusterEnableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterEnableActionWithDurationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterEnableActionWithDurationParams creates a new MTRActionsClusterEnableActionWithDurationParams instance.
func NewMTRActionsClusterEnableActionWithDurationParams() MTRActionsClusterEnableActionWithDurationParams {
	return getMTRActionsClusterEnableActionWithDurationParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/actionid
func (m_ MTRActionsClusterEnableActionWithDurationParams) ActionID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("actionID"))
	return rv
}


// SetActionID sets the value of the actionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/actionid
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetActionID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/duration
func (m_ MTRActionsClusterEnableActionWithDurationParams) Duration() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/duration
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetDuration(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/invokeid
func (m_ MTRActionsClusterEnableActionWithDurationParams) InvokeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("invokeID"))
	return rv
}


// SetInvokeID sets the value of the invokeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/invokeid
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetInvokeID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/serversideprocessingtimeout
func (m_ MTRActionsClusterEnableActionWithDurationParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/serversideprocessingtimeout
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/timedinvoketimeoutms
func (m_ MTRActionsClusterEnableActionWithDurationParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/timedinvoketimeoutms
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



