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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/actionid
func (m_ MTRActionsClusterEnableActionWithDurationParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/actionid
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/duration
func (m_ MTRActionsClusterEnableActionWithDurationParams) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/duration
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/invokeid
func (m_ MTRActionsClusterEnableActionWithDurationParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/invokeid
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/serversideprocessingtimeout
func (m_ MTRActionsClusterEnableActionWithDurationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/serversideprocessingtimeout
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/timedinvoketimeoutms
func (m_ MTRActionsClusterEnableActionWithDurationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionwithdurationparams/timedinvoketimeoutms
func (m_ MTRActionsClusterEnableActionWithDurationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



