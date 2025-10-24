// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWindowCoveringClusterGoToTiltValueParams] class.
var (
	MTRWindowCoveringClusterGoToTiltValueParamsClass     _MTRWindowCoveringClusterGoToTiltValueParamsClass
	MTRWindowCoveringClusterGoToTiltValueParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterGoToTiltValueParamsClass() _MTRWindowCoveringClusterGoToTiltValueParamsClass {
	MTRWindowCoveringClusterGoToTiltValueParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterGoToTiltValueParamsClass = _MTRWindowCoveringClusterGoToTiltValueParamsClass{objc.GetClass("MTRWindowCoveringClusterGoToTiltValueParams")}
	})
	return MTRWindowCoveringClusterGoToTiltValueParamsClass
}

type _MTRWindowCoveringClusterGoToTiltValueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWindowCoveringClusterGoToTiltValueParams] class.
type IMTRWindowCoveringClusterGoToTiltValueParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TiltValue() objc.IObject /* cross-framework: NSNumber */
	SetTiltValue(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToTiltValueParams
type MTRWindowCoveringClusterGoToTiltValueParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterGoToTiltValueParamsFrom constructs a [MTRWindowCoveringClusterGoToTiltValueParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterGoToTiltValueParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterGoToTiltValueParams {
	return MTRWindowCoveringClusterGoToTiltValueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterGoToTiltValueParamsClass) Alloc() MTRWindowCoveringClusterGoToTiltValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToTiltValueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWindowCoveringClusterGoToTiltValueParamsClass) New() MTRWindowCoveringClusterGoToTiltValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToTiltValueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterGoToTiltValueParams) Init() MTRWindowCoveringClusterGoToTiltValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToTiltValueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterGoToTiltValueParams) Autorelease() MTRWindowCoveringClusterGoToTiltValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToTiltValueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterGoToTiltValueParams creates a new MTRWindowCoveringClusterGoToTiltValueParams instance.
func NewMTRWindowCoveringClusterGoToTiltValueParams() MTRWindowCoveringClusterGoToTiltValueParams {
	return getMTRWindowCoveringClusterGoToTiltValueParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergototiltvalueparams/serversideprocessingtimeout
func (m_ MTRWindowCoveringClusterGoToTiltValueParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergototiltvalueparams/serversideprocessingtimeout
func (m_ MTRWindowCoveringClusterGoToTiltValueParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergototiltvalueparams/tiltvalue
func (m_ MTRWindowCoveringClusterGoToTiltValueParams) TiltValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("tiltValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergototiltvalueparams/tiltvalue
func (m_ MTRWindowCoveringClusterGoToTiltValueParams) SetTiltValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTiltValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergototiltvalueparams/timedinvoketimeoutms
func (m_ MTRWindowCoveringClusterGoToTiltValueParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergototiltvalueparams/timedinvoketimeoutms
func (m_ MTRWindowCoveringClusterGoToTiltValueParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



