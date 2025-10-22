// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterEnableActionParams] class.
var (
	MTRActionsClusterEnableActionParamsClass     _MTRActionsClusterEnableActionParamsClass
	MTRActionsClusterEnableActionParamsClassOnce sync.Once
)

func getMTRActionsClusterEnableActionParamsClass() _MTRActionsClusterEnableActionParamsClass {
	MTRActionsClusterEnableActionParamsClassOnce.Do(func() {
		MTRActionsClusterEnableActionParamsClass = _MTRActionsClusterEnableActionParamsClass{objc.GetClass("MTRActionsClusterEnableActionParams")}
	})
	return MTRActionsClusterEnableActionParamsClass
}

type _MTRActionsClusterEnableActionParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterEnableActionParams] class.
type IMTRActionsClusterEnableActionParams interface {
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionParams
type MTRActionsClusterEnableActionParams struct {
	objectivec.Object
}

// MTRActionsClusterEnableActionParamsFrom constructs a [MTRActionsClusterEnableActionParams] from an unsafe.Pointer.
func MTRActionsClusterEnableActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterEnableActionParams {
	return MTRActionsClusterEnableActionParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterEnableActionParamsClass) Alloc() MTRActionsClusterEnableActionParams {
	rv := objc.Send[MTRActionsClusterEnableActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterEnableActionParamsClass) New() MTRActionsClusterEnableActionParams {
	rv := objc.Send[MTRActionsClusterEnableActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterEnableActionParams) Init() MTRActionsClusterEnableActionParams {
	rv := objc.Send[MTRActionsClusterEnableActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterEnableActionParams) Autorelease() MTRActionsClusterEnableActionParams {
	rv := objc.Send[MTRActionsClusterEnableActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterEnableActionParams creates a new MTRActionsClusterEnableActionParams instance.
func NewMTRActionsClusterEnableActionParams() MTRActionsClusterEnableActionParams {
	return getMTRActionsClusterEnableActionParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/actionid
func (m_ MTRActionsClusterEnableActionParams) ActionID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("actionID"))
	return rv
}


// SetActionID sets the value of the actionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/actionid
func (m_ MTRActionsClusterEnableActionParams) SetActionID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/invokeid
func (m_ MTRActionsClusterEnableActionParams) InvokeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("invokeID"))
	return rv
}


// SetInvokeID sets the value of the invokeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/invokeid
func (m_ MTRActionsClusterEnableActionParams) SetInvokeID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterEnableActionParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterEnableActionParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterEnableActionParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterEnableActionParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



