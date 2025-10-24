// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterStepColorParams] class.
var (
	MTRColorControlClusterStepColorParamsClass     _MTRColorControlClusterStepColorParamsClass
	MTRColorControlClusterStepColorParamsClassOnce sync.Once
)

func getMTRColorControlClusterStepColorParamsClass() _MTRColorControlClusterStepColorParamsClass {
	MTRColorControlClusterStepColorParamsClassOnce.Do(func() {
		MTRColorControlClusterStepColorParamsClass = _MTRColorControlClusterStepColorParamsClass{objc.GetClass("MTRColorControlClusterStepColorParams")}
	})
	return MTRColorControlClusterStepColorParamsClass
}

type _MTRColorControlClusterStepColorParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterStepColorParams] class.
type IMTRColorControlClusterStepColorParams interface {
	objectivec.IObject
	// properties:
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StepX() objc.IObject /* cross-framework: NSNumber */
	SetStepX(value objc.IObject /* cross-framework: NSNumber */)
	StepY() objc.IObject /* cross-framework: NSNumber */
	SetStepY(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TransitionTime() objc.IObject /* cross-framework: NSNumber */
	SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams
type MTRColorControlClusterStepColorParams struct {
	objectivec.Object
}

// MTRColorControlClusterStepColorParamsFrom constructs a [MTRColorControlClusterStepColorParams] from an unsafe.Pointer.
func MTRColorControlClusterStepColorParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterStepColorParams {
	return MTRColorControlClusterStepColorParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterStepColorParamsClass) Alloc() MTRColorControlClusterStepColorParams {
	rv := objc.Send[MTRColorControlClusterStepColorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterStepColorParamsClass) New() MTRColorControlClusterStepColorParams {
	rv := objc.Send[MTRColorControlClusterStepColorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterStepColorParams) Init() MTRColorControlClusterStepColorParams {
	rv := objc.Send[MTRColorControlClusterStepColorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterStepColorParams) Autorelease() MTRColorControlClusterStepColorParams {
	rv := objc.Send[MTRColorControlClusterStepColorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterStepColorParams creates a new MTRColorControlClusterStepColorParams instance.
func NewMTRColorControlClusterStepColorParams() MTRColorControlClusterStepColorParams {
	return getMTRColorControlClusterStepColorParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/optionsmask
func (m_ MTRColorControlClusterStepColorParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/optionsmask
func (m_ MTRColorControlClusterStepColorParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/optionsoverride
func (m_ MTRColorControlClusterStepColorParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/optionsoverride
func (m_ MTRColorControlClusterStepColorParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterStepColorParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterStepColorParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/stepx
func (m_ MTRColorControlClusterStepColorParams) StepX() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/stepx
func (m_ MTRColorControlClusterStepColorParams) SetStepX(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/stepy
func (m_ MTRColorControlClusterStepColorParams) StepY() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/stepy
func (m_ MTRColorControlClusterStepColorParams) SetStepY(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterStepColorParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterStepColorParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/transitiontime
func (m_ MTRColorControlClusterStepColorParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolorparams/transitiontime
func (m_ MTRColorControlClusterStepColorParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



