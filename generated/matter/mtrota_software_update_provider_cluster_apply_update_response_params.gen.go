// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams] class.
var (
	MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass     _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass
	MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClassOnce sync.Once
)

func getMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass() _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass {
	MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClassOnce.Do(func() {
		MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass = _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass{objc.GetClass("MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams")}
	})
	return MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass
}

type _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams] class.
type IMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams interface {
	objectivec.IObject
	// properties:
	Action() objc.IObject /* cross-framework: NSNumber */
	SetAction(value objc.IObject /* cross-framework: NSNumber */)
	DelayedActionTime() objc.IObject /* cross-framework: NSNumber */
	SetDelayedActionTime(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams-36zc9
type MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom constructs a [MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	return MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass) Alloc() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass) New() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) Init() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) Autorelease() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams creates a new MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams instance.
func NewMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	return getMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdateresponseparams-36zc9/action
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) Action() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("action"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdateresponseparams-36zc9/action
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) SetAction(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdateresponseparams-36zc9/delayedactiontime
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) DelayedActionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("delayedActionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdateresponseparams-36zc9/delayedactiontime
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) SetDelayedActionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelayedActionTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdateresponseparams-36zc9/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdateresponseparams-36zc9/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



