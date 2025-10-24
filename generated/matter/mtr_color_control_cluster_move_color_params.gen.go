// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveColorParams] class.
var (
	MTRColorControlClusterMoveColorParamsClass     _MTRColorControlClusterMoveColorParamsClass
	MTRColorControlClusterMoveColorParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveColorParamsClass() _MTRColorControlClusterMoveColorParamsClass {
	MTRColorControlClusterMoveColorParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveColorParamsClass = _MTRColorControlClusterMoveColorParamsClass{objc.GetClass("MTRColorControlClusterMoveColorParams")}
	})
	return MTRColorControlClusterMoveColorParamsClass
}

type _MTRColorControlClusterMoveColorParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveColorParams] class.
type IMTRColorControlClusterMoveColorParams interface {
	objectivec.IObject
	// properties:
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	RateX() objc.IObject /* cross-framework: NSNumber */
	SetRateX(value objc.IObject /* cross-framework: NSNumber */)
	RateY() objc.IObject /* cross-framework: NSNumber */
	SetRateY(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams
type MTRColorControlClusterMoveColorParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveColorParamsFrom constructs a [MTRColorControlClusterMoveColorParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveColorParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveColorParams {
	return MTRColorControlClusterMoveColorParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveColorParamsClass) Alloc() MTRColorControlClusterMoveColorParams {
	rv := objc.Send[MTRColorControlClusterMoveColorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveColorParamsClass) New() MTRColorControlClusterMoveColorParams {
	rv := objc.Send[MTRColorControlClusterMoveColorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveColorParams) Init() MTRColorControlClusterMoveColorParams {
	rv := objc.Send[MTRColorControlClusterMoveColorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveColorParams) Autorelease() MTRColorControlClusterMoveColorParams {
	rv := objc.Send[MTRColorControlClusterMoveColorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveColorParams creates a new MTRColorControlClusterMoveColorParams instance.
func NewMTRColorControlClusterMoveColorParams() MTRColorControlClusterMoveColorParams {
	return getMTRColorControlClusterMoveColorParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolorparams/optionsmask
func (m_ MTRColorControlClusterMoveColorParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolorparams/optionsmask
func (m_ MTRColorControlClusterMoveColorParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolorparams/optionsoverride
func (m_ MTRColorControlClusterMoveColorParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolorparams/optionsoverride
func (m_ MTRColorControlClusterMoveColorParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolorparams/ratex
func (m_ MTRColorControlClusterMoveColorParams) RateX() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rateX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolorparams/ratex
func (m_ MTRColorControlClusterMoveColorParams) SetRateX(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRateX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolorparams/ratey
func (m_ MTRColorControlClusterMoveColorParams) RateY() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rateY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolorparams/ratey
func (m_ MTRColorControlClusterMoveColorParams) SetRateY(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRateY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolorparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveColorParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolorparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveColorParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolorparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveColorParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolorparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveColorParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



