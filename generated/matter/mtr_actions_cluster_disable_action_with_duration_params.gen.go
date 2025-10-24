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
	// properties:
	ActionID() objc.IObject /* cross-framework: NSNumber */
	SetActionID(value objc.IObject /* cross-framework: NSNumber */)
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	InvokeID() objc.IObject /* cross-framework: NSNumber */
	SetInvokeID(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/actionid
func (m_ MTRActionsClusterDisableActionWithDurationParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/actionid
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/duration
func (m_ MTRActionsClusterDisableActionWithDurationParams) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/duration
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/invokeid
func (m_ MTRActionsClusterDisableActionWithDurationParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/invokeid
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/serversideprocessingtimeout
func (m_ MTRActionsClusterDisableActionWithDurationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/serversideprocessingtimeout
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/timedinvoketimeoutms
func (m_ MTRActionsClusterDisableActionWithDurationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterdisableactionwithdurationparams/timedinvoketimeoutms
func (m_ MTRActionsClusterDisableActionWithDurationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



