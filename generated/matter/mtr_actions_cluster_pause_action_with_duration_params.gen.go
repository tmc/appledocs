// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterPauseActionWithDurationParams] class.
var (
	MTRActionsClusterPauseActionWithDurationParamsClass     _MTRActionsClusterPauseActionWithDurationParamsClass
	MTRActionsClusterPauseActionWithDurationParamsClassOnce sync.Once
)

func getMTRActionsClusterPauseActionWithDurationParamsClass() _MTRActionsClusterPauseActionWithDurationParamsClass {
	MTRActionsClusterPauseActionWithDurationParamsClassOnce.Do(func() {
		MTRActionsClusterPauseActionWithDurationParamsClass = _MTRActionsClusterPauseActionWithDurationParamsClass{objc.GetClass("MTRActionsClusterPauseActionWithDurationParams")}
	})
	return MTRActionsClusterPauseActionWithDurationParamsClass
}

type _MTRActionsClusterPauseActionWithDurationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterPauseActionWithDurationParams] class.
type IMTRActionsClusterPauseActionWithDurationParams interface {
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionWithDurationParams
type MTRActionsClusterPauseActionWithDurationParams struct {
	objectivec.Object
}

// MTRActionsClusterPauseActionWithDurationParamsFrom constructs a [MTRActionsClusterPauseActionWithDurationParams] from an unsafe.Pointer.
func MTRActionsClusterPauseActionWithDurationParamsFrom(ptr unsafe.Pointer) MTRActionsClusterPauseActionWithDurationParams {
	return MTRActionsClusterPauseActionWithDurationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterPauseActionWithDurationParamsClass) Alloc() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterPauseActionWithDurationParamsClass) New() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterPauseActionWithDurationParams) Init() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterPauseActionWithDurationParams) Autorelease() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterPauseActionWithDurationParams creates a new MTRActionsClusterPauseActionWithDurationParams instance.
func NewMTRActionsClusterPauseActionWithDurationParams() MTRActionsClusterPauseActionWithDurationParams {
	return getMTRActionsClusterPauseActionWithDurationParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/actionid
func (m_ MTRActionsClusterPauseActionWithDurationParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/actionid
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/duration
func (m_ MTRActionsClusterPauseActionWithDurationParams) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/duration
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/invokeid
func (m_ MTRActionsClusterPauseActionWithDurationParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/invokeid
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/serversideprocessingtimeout
func (m_ MTRActionsClusterPauseActionWithDurationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/serversideprocessingtimeout
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/timedinvoketimeoutms
func (m_ MTRActionsClusterPauseActionWithDurationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterpauseactionwithdurationparams/timedinvoketimeoutms
func (m_ MTRActionsClusterPauseActionWithDurationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



