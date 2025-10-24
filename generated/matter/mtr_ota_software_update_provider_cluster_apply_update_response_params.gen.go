// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams] class.
var (
	MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass     _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass
	MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClassOnce sync.Once
)

func getMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass() _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass {
	MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClassOnce.Do(func() {
		MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass = _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass{objc.GetClass("MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams")}
	})
	return MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass
}

type _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams] class.
type IMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams interface {
	IMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams-92als
type MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams struct {
	MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams
}

// MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom constructs a [MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams] from an unsafe.Pointer.
func MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	return MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams{
		MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams: MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass) Alloc() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass) New() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) Init() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) Autorelease() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams creates a new MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams instance.
func NewMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	return getMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdateresponseparams-92als/action
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) Action() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("action"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdateresponseparams-92als/action
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) SetAction(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdateresponseparams-92als/delayedactiontime
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) DelayedActionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("delayedActionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdateresponseparams-92als/delayedactiontime
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) SetDelayedActionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelayedActionTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdateresponseparams-92als/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdateresponseparams-92als/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



