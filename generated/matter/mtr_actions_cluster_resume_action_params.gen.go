// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterResumeActionParams] class.
var (
	MTRActionsClusterResumeActionParamsClass     _MTRActionsClusterResumeActionParamsClass
	MTRActionsClusterResumeActionParamsClassOnce sync.Once
)

func getMTRActionsClusterResumeActionParamsClass() _MTRActionsClusterResumeActionParamsClass {
	MTRActionsClusterResumeActionParamsClassOnce.Do(func() {
		MTRActionsClusterResumeActionParamsClass = _MTRActionsClusterResumeActionParamsClass{objc.GetClass("MTRActionsClusterResumeActionParams")}
	})
	return MTRActionsClusterResumeActionParamsClass
}

type _MTRActionsClusterResumeActionParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterResumeActionParams] class.
type IMTRActionsClusterResumeActionParams interface {
	objectivec.IObject
	ActionID() foundation.Number
	SetActionID(value foundation.INumber)
	InvokeID() foundation.Number
	SetInvokeID(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterResumeActionParams
type MTRActionsClusterResumeActionParams struct {
	objectivec.Object
}

// MTRActionsClusterResumeActionParamsFrom constructs a [MTRActionsClusterResumeActionParams] from an unsafe.Pointer.
func MTRActionsClusterResumeActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterResumeActionParams {
	return MTRActionsClusterResumeActionParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterResumeActionParamsClass) Alloc() MTRActionsClusterResumeActionParams {
	rv := objc.Send[MTRActionsClusterResumeActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterResumeActionParamsClass) New() MTRActionsClusterResumeActionParams {
	rv := objc.Send[MTRActionsClusterResumeActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterResumeActionParams) Init() MTRActionsClusterResumeActionParams {
	rv := objc.Send[MTRActionsClusterResumeActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterResumeActionParams) Autorelease() MTRActionsClusterResumeActionParams {
	rv := objc.Send[MTRActionsClusterResumeActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterResumeActionParams creates a new MTRActionsClusterResumeActionParams instance.
func NewMTRActionsClusterResumeActionParams() MTRActionsClusterResumeActionParams {
	return getMTRActionsClusterResumeActionParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterresumeactionparams/actionid
func (m_ MTRActionsClusterResumeActionParams) ActionID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("actionID"))
	return rv
}


// SetActionID sets the value of the actionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterresumeactionparams/actionid
func (m_ MTRActionsClusterResumeActionParams) SetActionID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterresumeactionparams/invokeid
func (m_ MTRActionsClusterResumeActionParams) InvokeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("invokeID"))
	return rv
}


// SetInvokeID sets the value of the invokeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterresumeactionparams/invokeid
func (m_ MTRActionsClusterResumeActionParams) SetInvokeID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterresumeactionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterResumeActionParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterresumeactionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterResumeActionParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterresumeactionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterResumeActionParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterresumeactionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterResumeActionParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



