// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterInstantActionParams] class.
var (
	MTRActionsClusterInstantActionParamsClass     _MTRActionsClusterInstantActionParamsClass
	MTRActionsClusterInstantActionParamsClassOnce sync.Once
)

func getMTRActionsClusterInstantActionParamsClass() _MTRActionsClusterInstantActionParamsClass {
	MTRActionsClusterInstantActionParamsClassOnce.Do(func() {
		MTRActionsClusterInstantActionParamsClass = _MTRActionsClusterInstantActionParamsClass{objc.GetClass("MTRActionsClusterInstantActionParams")}
	})
	return MTRActionsClusterInstantActionParamsClass
}

type _MTRActionsClusterInstantActionParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterInstantActionParams] class.
type IMTRActionsClusterInstantActionParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionParams
type MTRActionsClusterInstantActionParams struct {
	objectivec.Object
}

// MTRActionsClusterInstantActionParamsFrom constructs a [MTRActionsClusterInstantActionParams] from an unsafe.Pointer.
func MTRActionsClusterInstantActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterInstantActionParams {
	return MTRActionsClusterInstantActionParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterInstantActionParamsClass) Alloc() MTRActionsClusterInstantActionParams {
	rv := objc.Send[MTRActionsClusterInstantActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterInstantActionParamsClass) New() MTRActionsClusterInstantActionParams {
	rv := objc.Send[MTRActionsClusterInstantActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterInstantActionParams) Init() MTRActionsClusterInstantActionParams {
	rv := objc.Send[MTRActionsClusterInstantActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterInstantActionParams) Autorelease() MTRActionsClusterInstantActionParams {
	rv := objc.Send[MTRActionsClusterInstantActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterInstantActionParams creates a new MTRActionsClusterInstantActionParams instance.
func NewMTRActionsClusterInstantActionParams() MTRActionsClusterInstantActionParams {
	return getMTRActionsClusterInstantActionParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterInstantActionParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterInstantActionParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/invokeid
func (m_ MTRActionsClusterInstantActionParams) InvokeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("invokeID"))
	return rv
}


// SetInvokeID sets the value of the invokeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/invokeid
func (m_ MTRActionsClusterInstantActionParams) SetInvokeID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/actionid
func (m_ MTRActionsClusterInstantActionParams) ActionID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("actionID"))
	return rv
}


// SetActionID sets the value of the actionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/actionid
func (m_ MTRActionsClusterInstantActionParams) SetActionID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterInstantActionParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterInstantActionParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}



