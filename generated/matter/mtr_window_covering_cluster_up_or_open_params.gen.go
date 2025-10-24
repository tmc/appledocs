// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWindowCoveringClusterUpOrOpenParams] class.
var (
	MTRWindowCoveringClusterUpOrOpenParamsClass     _MTRWindowCoveringClusterUpOrOpenParamsClass
	MTRWindowCoveringClusterUpOrOpenParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterUpOrOpenParamsClass() _MTRWindowCoveringClusterUpOrOpenParamsClass {
	MTRWindowCoveringClusterUpOrOpenParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterUpOrOpenParamsClass = _MTRWindowCoveringClusterUpOrOpenParamsClass{objc.GetClass("MTRWindowCoveringClusterUpOrOpenParams")}
	})
	return MTRWindowCoveringClusterUpOrOpenParamsClass
}

type _MTRWindowCoveringClusterUpOrOpenParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWindowCoveringClusterUpOrOpenParams] class.
type IMTRWindowCoveringClusterUpOrOpenParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterUpOrOpenParams
type MTRWindowCoveringClusterUpOrOpenParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterUpOrOpenParamsFrom constructs a [MTRWindowCoveringClusterUpOrOpenParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterUpOrOpenParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterUpOrOpenParams {
	return MTRWindowCoveringClusterUpOrOpenParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterUpOrOpenParamsClass) Alloc() MTRWindowCoveringClusterUpOrOpenParams {
	rv := objc.Send[MTRWindowCoveringClusterUpOrOpenParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWindowCoveringClusterUpOrOpenParamsClass) New() MTRWindowCoveringClusterUpOrOpenParams {
	rv := objc.Send[MTRWindowCoveringClusterUpOrOpenParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterUpOrOpenParams) Init() MTRWindowCoveringClusterUpOrOpenParams {
	rv := objc.Send[MTRWindowCoveringClusterUpOrOpenParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterUpOrOpenParams) Autorelease() MTRWindowCoveringClusterUpOrOpenParams {
	rv := objc.Send[MTRWindowCoveringClusterUpOrOpenParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterUpOrOpenParams creates a new MTRWindowCoveringClusterUpOrOpenParams instance.
func NewMTRWindowCoveringClusterUpOrOpenParams() MTRWindowCoveringClusterUpOrOpenParams {
	return getMTRWindowCoveringClusterUpOrOpenParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclusteruporopenparams/serversideprocessingtimeout
func (m_ MTRWindowCoveringClusterUpOrOpenParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclusteruporopenparams/serversideprocessingtimeout
func (m_ MTRWindowCoveringClusterUpOrOpenParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclusteruporopenparams/timedinvoketimeoutms
func (m_ MTRWindowCoveringClusterUpOrOpenParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclusteruporopenparams/timedinvoketimeoutms
func (m_ MTRWindowCoveringClusterUpOrOpenParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



