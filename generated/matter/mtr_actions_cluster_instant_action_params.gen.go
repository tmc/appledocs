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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/actionid
func (m_ MTRActionsClusterInstantActionParams) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/actionid
func (m_ MTRActionsClusterInstantActionParams) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/invokeid
func (m_ MTRActionsClusterInstantActionParams) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/invokeid
func (m_ MTRActionsClusterInstantActionParams) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterInstantActionParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/serversideprocessingtimeout
func (m_ MTRActionsClusterInstantActionParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterInstantActionParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterinstantactionparams/timedinvoketimeoutms
func (m_ MTRActionsClusterInstantActionParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



