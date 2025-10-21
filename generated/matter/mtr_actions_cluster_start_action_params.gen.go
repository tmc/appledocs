// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterStartActionParams] class.
var (
	MTRActionsClusterStartActionParamsClass     _MTRActionsClusterStartActionParamsClass
	MTRActionsClusterStartActionParamsClassOnce sync.Once
)

func getMTRActionsClusterStartActionParamsClass() _MTRActionsClusterStartActionParamsClass {
	MTRActionsClusterStartActionParamsClassOnce.Do(func() {
		MTRActionsClusterStartActionParamsClass = _MTRActionsClusterStartActionParamsClass{objc.GetClass("MTRActionsClusterStartActionParams")}
	})
	return MTRActionsClusterStartActionParamsClass
}

type _MTRActionsClusterStartActionParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterStartActionParams] class.
type IMTRActionsClusterStartActionParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStartActionParams
type MTRActionsClusterStartActionParams struct {
	objectivec.Object
}

// MTRActionsClusterStartActionParamsFrom constructs a [MTRActionsClusterStartActionParams] from an unsafe.Pointer.
func MTRActionsClusterStartActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterStartActionParams {
	return MTRActionsClusterStartActionParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterStartActionParamsClass) Alloc() MTRActionsClusterStartActionParams {
	rv := objc.Send[MTRActionsClusterStartActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterStartActionParamsClass) New() MTRActionsClusterStartActionParams {
	rv := objc.Send[MTRActionsClusterStartActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterStartActionParams) Init() MTRActionsClusterStartActionParams {
	rv := objc.Send[MTRActionsClusterStartActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterStartActionParams) Autorelease() MTRActionsClusterStartActionParams {
	rv := objc.Send[MTRActionsClusterStartActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterStartActionParams creates a new MTRActionsClusterStartActionParams instance.
func NewMTRActionsClusterStartActionParams() MTRActionsClusterStartActionParams {
	return getMTRActionsClusterStartActionParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionparams/actionid
func (m_ MTRActionsClusterStartActionParams) ActionID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("actionID"))
	return rv
}


// SetActionID sets the value of the actionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionparams/actionid
func (m_ MTRActionsClusterStartActionParams) SetActionID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionparams/invokeid
func (m_ MTRActionsClusterStartActionParams) InvokeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("invokeID"))
	return rv
}


// SetInvokeID sets the value of the invokeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionparams/invokeid
func (m_ MTRActionsClusterStartActionParams) SetInvokeID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterStartActionParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterStartActionParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterStartActionParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterStartActionParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



