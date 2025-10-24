// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWindowCoveringClusterStopMotionParams] class.
var (
	MTRWindowCoveringClusterStopMotionParamsClass     _MTRWindowCoveringClusterStopMotionParamsClass
	MTRWindowCoveringClusterStopMotionParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterStopMotionParamsClass() _MTRWindowCoveringClusterStopMotionParamsClass {
	MTRWindowCoveringClusterStopMotionParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterStopMotionParamsClass = _MTRWindowCoveringClusterStopMotionParamsClass{objc.GetClass("MTRWindowCoveringClusterStopMotionParams")}
	})
	return MTRWindowCoveringClusterStopMotionParamsClass
}

type _MTRWindowCoveringClusterStopMotionParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWindowCoveringClusterStopMotionParams] class.
type IMTRWindowCoveringClusterStopMotionParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterStopMotionParams
type MTRWindowCoveringClusterStopMotionParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterStopMotionParamsFrom constructs a [MTRWindowCoveringClusterStopMotionParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterStopMotionParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterStopMotionParams {
	return MTRWindowCoveringClusterStopMotionParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterStopMotionParamsClass) Alloc() MTRWindowCoveringClusterStopMotionParams {
	rv := objc.Send[MTRWindowCoveringClusterStopMotionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWindowCoveringClusterStopMotionParamsClass) New() MTRWindowCoveringClusterStopMotionParams {
	rv := objc.Send[MTRWindowCoveringClusterStopMotionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterStopMotionParams) Init() MTRWindowCoveringClusterStopMotionParams {
	rv := objc.Send[MTRWindowCoveringClusterStopMotionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterStopMotionParams) Autorelease() MTRWindowCoveringClusterStopMotionParams {
	rv := objc.Send[MTRWindowCoveringClusterStopMotionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterStopMotionParams creates a new MTRWindowCoveringClusterStopMotionParams instance.
func NewMTRWindowCoveringClusterStopMotionParams() MTRWindowCoveringClusterStopMotionParams {
	return getMTRWindowCoveringClusterStopMotionParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclusterstopmotionparams/serversideprocessingtimeout
func (m_ MTRWindowCoveringClusterStopMotionParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclusterstopmotionparams/serversideprocessingtimeout
func (m_ MTRWindowCoveringClusterStopMotionParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclusterstopmotionparams/timedinvoketimeoutms
func (m_ MTRWindowCoveringClusterStopMotionParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclusterstopmotionparams/timedinvoketimeoutms
func (m_ MTRWindowCoveringClusterStopMotionParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



