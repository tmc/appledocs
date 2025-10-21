// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterPauseActionParams] class.
var (
	MTRActionsClusterPauseActionParamsClass     _MTRActionsClusterPauseActionParamsClass
	MTRActionsClusterPauseActionParamsClassOnce sync.Once
)

func getMTRActionsClusterPauseActionParamsClass() _MTRActionsClusterPauseActionParamsClass {
	MTRActionsClusterPauseActionParamsClassOnce.Do(func() {
		MTRActionsClusterPauseActionParamsClass = _MTRActionsClusterPauseActionParamsClass{objc.GetClass("MTRActionsClusterPauseActionParams")}
	})
	return MTRActionsClusterPauseActionParamsClass
}

type _MTRActionsClusterPauseActionParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterPauseActionParams] class.
type IMTRActionsClusterPauseActionParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionParams
type MTRActionsClusterPauseActionParams struct {
	objectivec.Object
}

// MTRActionsClusterPauseActionParamsFrom constructs a [MTRActionsClusterPauseActionParams] from an unsafe.Pointer.
func MTRActionsClusterPauseActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterPauseActionParams {
	return MTRActionsClusterPauseActionParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterPauseActionParamsClass) Alloc() MTRActionsClusterPauseActionParams {
	rv := objc.Send[MTRActionsClusterPauseActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterPauseActionParamsClass) New() MTRActionsClusterPauseActionParams {
	rv := objc.Send[MTRActionsClusterPauseActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterPauseActionParams) Init() MTRActionsClusterPauseActionParams {
	rv := objc.Send[MTRActionsClusterPauseActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterPauseActionParams) Autorelease() MTRActionsClusterPauseActionParams {
	rv := objc.Send[MTRActionsClusterPauseActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterPauseActionParams creates a new MTRActionsClusterPauseActionParams instance.
func NewMTRActionsClusterPauseActionParams() MTRActionsClusterPauseActionParams {
	return getMTRActionsClusterPauseActionParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterPauseActionParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterPauseActionParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterPauseActionParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterPauseActionParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionparams/actionid
func (m_ MTRActionsClusterPauseActionParams) ActionID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("actionID"))
	return rv
}


// SetActionID sets the value of the actionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionparams/actionid
func (m_ MTRActionsClusterPauseActionParams) SetActionID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionparams/invokeid
func (m_ MTRActionsClusterPauseActionParams) InvokeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("invokeID"))
	return rv
}


// SetInvokeID sets the value of the invokeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionparams/invokeid
func (m_ MTRActionsClusterPauseActionParams) SetInvokeID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}



