// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterStartActionWithDurationParams] class.
var (
	MTRActionsClusterStartActionWithDurationParamsClass     _MTRActionsClusterStartActionWithDurationParamsClass
	MTRActionsClusterStartActionWithDurationParamsClassOnce sync.Once
)

func getMTRActionsClusterStartActionWithDurationParamsClass() _MTRActionsClusterStartActionWithDurationParamsClass {
	MTRActionsClusterStartActionWithDurationParamsClassOnce.Do(func() {
		MTRActionsClusterStartActionWithDurationParamsClass = _MTRActionsClusterStartActionWithDurationParamsClass{objc.GetClass("MTRActionsClusterStartActionWithDurationParams")}
	})
	return MTRActionsClusterStartActionWithDurationParamsClass
}

type _MTRActionsClusterStartActionWithDurationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterStartActionWithDurationParams] class.
type IMTRActionsClusterStartActionWithDurationParams interface {
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStartActionWithDurationParams
type MTRActionsClusterStartActionWithDurationParams struct {
	objectivec.Object
}

// MTRActionsClusterStartActionWithDurationParamsFrom constructs a [MTRActionsClusterStartActionWithDurationParams] from an unsafe.Pointer.
func MTRActionsClusterStartActionWithDurationParamsFrom(ptr unsafe.Pointer) MTRActionsClusterStartActionWithDurationParams {
	return MTRActionsClusterStartActionWithDurationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterStartActionWithDurationParamsClass) Alloc() MTRActionsClusterStartActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterStartActionWithDurationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterStartActionWithDurationParamsClass) New() MTRActionsClusterStartActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterStartActionWithDurationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterStartActionWithDurationParams) Init() MTRActionsClusterStartActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterStartActionWithDurationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterStartActionWithDurationParams) Autorelease() MTRActionsClusterStartActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterStartActionWithDurationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterStartActionWithDurationParams creates a new MTRActionsClusterStartActionWithDurationParams instance.
func NewMTRActionsClusterStartActionWithDurationParams() MTRActionsClusterStartActionWithDurationParams {
	return getMTRActionsClusterStartActionWithDurationParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionwithdurationparams/actionid
func (m_ MTRActionsClusterStartActionWithDurationParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionwithdurationparams/actionid
func (m_ MTRActionsClusterStartActionWithDurationParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionwithdurationparams/duration
func (m_ MTRActionsClusterStartActionWithDurationParams) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionwithdurationparams/duration
func (m_ MTRActionsClusterStartActionWithDurationParams) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionwithdurationparams/invokeid
func (m_ MTRActionsClusterStartActionWithDurationParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionwithdurationparams/invokeid
func (m_ MTRActionsClusterStartActionWithDurationParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionwithdurationparams/serversideprocessingtimeout
func (m_ MTRActionsClusterStartActionWithDurationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionwithdurationparams/serversideprocessingtimeout
func (m_ MTRActionsClusterStartActionWithDurationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionwithdurationparams/timedinvoketimeoutms
func (m_ MTRActionsClusterStartActionWithDurationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstartactionwithdurationparams/timedinvoketimeoutms
func (m_ MTRActionsClusterStartActionWithDurationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



