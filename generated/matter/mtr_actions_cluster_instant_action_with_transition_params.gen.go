// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterInstantActionWithTransitionParams] class.
var (
	MTRActionsClusterInstantActionWithTransitionParamsClass     _MTRActionsClusterInstantActionWithTransitionParamsClass
	MTRActionsClusterInstantActionWithTransitionParamsClassOnce sync.Once
)

func getMTRActionsClusterInstantActionWithTransitionParamsClass() _MTRActionsClusterInstantActionWithTransitionParamsClass {
	MTRActionsClusterInstantActionWithTransitionParamsClassOnce.Do(func() {
		MTRActionsClusterInstantActionWithTransitionParamsClass = _MTRActionsClusterInstantActionWithTransitionParamsClass{objc.GetClass("MTRActionsClusterInstantActionWithTransitionParams")}
	})
	return MTRActionsClusterInstantActionWithTransitionParamsClass
}

type _MTRActionsClusterInstantActionWithTransitionParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterInstantActionWithTransitionParams] class.
type IMTRActionsClusterInstantActionWithTransitionParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionWithTransitionParams
type MTRActionsClusterInstantActionWithTransitionParams struct {
	objectivec.Object
}

// MTRActionsClusterInstantActionWithTransitionParamsFrom constructs a [MTRActionsClusterInstantActionWithTransitionParams] from an unsafe.Pointer.
func MTRActionsClusterInstantActionWithTransitionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterInstantActionWithTransitionParams {
	return MTRActionsClusterInstantActionWithTransitionParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterInstantActionWithTransitionParamsClass) Alloc() MTRActionsClusterInstantActionWithTransitionParams {
	rv := objc.Send[MTRActionsClusterInstantActionWithTransitionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterInstantActionWithTransitionParamsClass) New() MTRActionsClusterInstantActionWithTransitionParams {
	rv := objc.Send[MTRActionsClusterInstantActionWithTransitionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterInstantActionWithTransitionParams) Init() MTRActionsClusterInstantActionWithTransitionParams {
	rv := objc.Send[MTRActionsClusterInstantActionWithTransitionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterInstantActionWithTransitionParams) Autorelease() MTRActionsClusterInstantActionWithTransitionParams {
	rv := objc.Send[MTRActionsClusterInstantActionWithTransitionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterInstantActionWithTransitionParams creates a new MTRActionsClusterInstantActionWithTransitionParams instance.
func NewMTRActionsClusterInstantActionWithTransitionParams() MTRActionsClusterInstantActionWithTransitionParams {
	return getMTRActionsClusterInstantActionWithTransitionParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionwithtransitionparams/actionid
func (m_ MTRActionsClusterInstantActionWithTransitionParams) ActionID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("actionID"))
	return rv
}


// SetActionID sets the value of the actionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionwithtransitionparams/actionid
func (m_ MTRActionsClusterInstantActionWithTransitionParams) SetActionID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionwithtransitionparams/invokeid
func (m_ MTRActionsClusterInstantActionWithTransitionParams) InvokeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("invokeID"))
	return rv
}


// SetInvokeID sets the value of the invokeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionwithtransitionparams/invokeid
func (m_ MTRActionsClusterInstantActionWithTransitionParams) SetInvokeID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionwithtransitionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterInstantActionWithTransitionParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionwithtransitionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterInstantActionWithTransitionParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionwithtransitionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterInstantActionWithTransitionParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionwithtransitionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterInstantActionWithTransitionParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionwithtransitionparams/transitiontime
func (m_ MTRActionsClusterInstantActionWithTransitionParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionwithtransitionparams/transitiontime
func (m_ MTRActionsClusterInstantActionWithTransitionParams) SetTransitionTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



