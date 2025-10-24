// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterEnhancedStepHueParams] class.
var (
	MTRColorControlClusterEnhancedStepHueParamsClass     _MTRColorControlClusterEnhancedStepHueParamsClass
	MTRColorControlClusterEnhancedStepHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterEnhancedStepHueParamsClass() _MTRColorControlClusterEnhancedStepHueParamsClass {
	MTRColorControlClusterEnhancedStepHueParamsClassOnce.Do(func() {
		MTRColorControlClusterEnhancedStepHueParamsClass = _MTRColorControlClusterEnhancedStepHueParamsClass{objc.GetClass("MTRColorControlClusterEnhancedStepHueParams")}
	})
	return MTRColorControlClusterEnhancedStepHueParamsClass
}

type _MTRColorControlClusterEnhancedStepHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterEnhancedStepHueParams] class.
type IMTRColorControlClusterEnhancedStepHueParams interface {
	objectivec.IObject
	// properties:
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StepMode() objc.IObject /* cross-framework: NSNumber */
	SetStepMode(value objc.IObject /* cross-framework: NSNumber */)
	StepSize() objc.IObject /* cross-framework: NSNumber */
	SetStepSize(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TransitionTime() objc.IObject /* cross-framework: NSNumber */
	SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams
type MTRColorControlClusterEnhancedStepHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterEnhancedStepHueParamsFrom constructs a [MTRColorControlClusterEnhancedStepHueParams] from an unsafe.Pointer.
func MTRColorControlClusterEnhancedStepHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterEnhancedStepHueParams {
	return MTRColorControlClusterEnhancedStepHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterEnhancedStepHueParamsClass) Alloc() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterEnhancedStepHueParamsClass) New() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterEnhancedStepHueParams) Init() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterEnhancedStepHueParams) Autorelease() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterEnhancedStepHueParams creates a new MTRColorControlClusterEnhancedStepHueParams instance.
func NewMTRColorControlClusterEnhancedStepHueParams() MTRColorControlClusterEnhancedStepHueParams {
	return getMTRColorControlClusterEnhancedStepHueParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/optionsmask
func (m_ MTRColorControlClusterEnhancedStepHueParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/optionsmask
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/optionsoverride
func (m_ MTRColorControlClusterEnhancedStepHueParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/optionsoverride
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterEnhancedStepHueParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/stepmode
func (m_ MTRColorControlClusterEnhancedStepHueParams) StepMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/stepmode
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetStepMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/stepsize
func (m_ MTRColorControlClusterEnhancedStepHueParams) StepSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/stepsize
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetStepSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepSize:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterEnhancedStepHueParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/transitiontime
func (m_ MTRColorControlClusterEnhancedStepHueParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedstephueparams/transitiontime
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



