// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWindowCoveringClusterGoToLiftValueParams] class.
var (
	MTRWindowCoveringClusterGoToLiftValueParamsClass     _MTRWindowCoveringClusterGoToLiftValueParamsClass
	MTRWindowCoveringClusterGoToLiftValueParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterGoToLiftValueParamsClass() _MTRWindowCoveringClusterGoToLiftValueParamsClass {
	MTRWindowCoveringClusterGoToLiftValueParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterGoToLiftValueParamsClass = _MTRWindowCoveringClusterGoToLiftValueParamsClass{objc.GetClass("MTRWindowCoveringClusterGoToLiftValueParams")}
	})
	return MTRWindowCoveringClusterGoToLiftValueParamsClass
}

type _MTRWindowCoveringClusterGoToLiftValueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWindowCoveringClusterGoToLiftValueParams] class.
type IMTRWindowCoveringClusterGoToLiftValueParams interface {
	objectivec.IObject
	// properties:
	LiftValue() objc.IObject /* cross-framework: NSNumber */
	SetLiftValue(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftValueParams
type MTRWindowCoveringClusterGoToLiftValueParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterGoToLiftValueParamsFrom constructs a [MTRWindowCoveringClusterGoToLiftValueParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterGoToLiftValueParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterGoToLiftValueParams {
	return MTRWindowCoveringClusterGoToLiftValueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterGoToLiftValueParamsClass) Alloc() MTRWindowCoveringClusterGoToLiftValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftValueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWindowCoveringClusterGoToLiftValueParamsClass) New() MTRWindowCoveringClusterGoToLiftValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftValueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) Init() MTRWindowCoveringClusterGoToLiftValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftValueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) Autorelease() MTRWindowCoveringClusterGoToLiftValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftValueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterGoToLiftValueParams creates a new MTRWindowCoveringClusterGoToLiftValueParams instance.
func NewMTRWindowCoveringClusterGoToLiftValueParams() MTRWindowCoveringClusterGoToLiftValueParams {
	return getMTRWindowCoveringClusterGoToLiftValueParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergotoliftvalueparams/liftvalue
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) LiftValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("liftValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergotoliftvalueparams/liftvalue
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) SetLiftValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLiftValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergotoliftvalueparams/serversideprocessingtimeout
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergotoliftvalueparams/serversideprocessingtimeout
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergotoliftvalueparams/timedinvoketimeoutms
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwindowcoveringclustergotoliftvalueparams/timedinvoketimeoutms
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



