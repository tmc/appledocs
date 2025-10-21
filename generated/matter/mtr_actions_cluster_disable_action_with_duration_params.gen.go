// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterDisableActionWithDurationParams] class.
var (
	MTRActionsClusterDisableActionWithDurationParamsClass     _MTRActionsClusterDisableActionWithDurationParamsClass
	MTRActionsClusterDisableActionWithDurationParamsClassOnce sync.Once
)

func getMTRActionsClusterDisableActionWithDurationParamsClass() _MTRActionsClusterDisableActionWithDurationParamsClass {
	MTRActionsClusterDisableActionWithDurationParamsClassOnce.Do(func() {
		MTRActionsClusterDisableActionWithDurationParamsClass = _MTRActionsClusterDisableActionWithDurationParamsClass{objc.GetClass("MTRActionsClusterDisableActionWithDurationParams")}
	})
	return MTRActionsClusterDisableActionWithDurationParamsClass
}

type _MTRActionsClusterDisableActionWithDurationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterDisableActionWithDurationParams] class.
type IMTRActionsClusterDisableActionWithDurationParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionWithDurationParams
type MTRActionsClusterDisableActionWithDurationParams struct {
	objectivec.Object
}

// MTRActionsClusterDisableActionWithDurationParamsFrom constructs a [MTRActionsClusterDisableActionWithDurationParams] from an unsafe.Pointer.
func MTRActionsClusterDisableActionWithDurationParamsFrom(ptr unsafe.Pointer) MTRActionsClusterDisableActionWithDurationParams {
	return MTRActionsClusterDisableActionWithDurationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterDisableActionWithDurationParamsClass) Alloc() MTRActionsClusterDisableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterDisableActionWithDurationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterDisableActionWithDurationParamsClass) New() MTRActionsClusterDisableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterDisableActionWithDurationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterDisableActionWithDurationParams) Init() MTRActionsClusterDisableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterDisableActionWithDurationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterDisableActionWithDurationParams) Autorelease() MTRActionsClusterDisableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterDisableActionWithDurationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterDisableActionWithDurationParams creates a new MTRActionsClusterDisableActionWithDurationParams instance.
func NewMTRActionsClusterDisableActionWithDurationParams() MTRActionsClusterDisableActionWithDurationParams {
	return getMTRActionsClusterDisableActionWithDurationParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/actionid
func (m_ MTRActionsClusterDisableActionWithDurationParams) ActionID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("actionID"))
	return rv
}


// SetActionID sets the value of the actionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/actionid
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetActionID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/duration
func (m_ MTRActionsClusterDisableActionWithDurationParams) Duration() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/duration
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetDuration(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/invokeid
func (m_ MTRActionsClusterDisableActionWithDurationParams) InvokeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("invokeID"))
	return rv
}


// SetInvokeID sets the value of the invokeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/invokeid
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetInvokeID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/serversideprocessingtimeout
func (m_ MTRActionsClusterDisableActionWithDurationParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/serversideprocessingtimeout
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/timedinvoketimeoutms
func (m_ MTRActionsClusterDisableActionWithDurationParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/timedinvoketimeoutms
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



