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
	// properties:
	ActionID() objc.IObject /* cross-framework: NSNumber */
	SetActionID(value objc.IObject /* cross-framework: NSNumber */)
	InvokeID() objc.IObject /* cross-framework: NSNumber */
	SetInvokeID(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/actionid
func (m_ MTRActionsClusterEnableActionParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/actionid
func (m_ MTRActionsClusterEnableActionParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/invokeid
func (m_ MTRActionsClusterEnableActionParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/invokeid
func (m_ MTRActionsClusterEnableActionParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterEnableActionParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterEnableActionParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterEnableActionParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterenableactionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterEnableActionParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



