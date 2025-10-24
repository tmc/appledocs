// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterStepColorTemperatureParams] class.
var (
	MTRColorControlClusterStepColorTemperatureParamsClass     _MTRColorControlClusterStepColorTemperatureParamsClass
	MTRColorControlClusterStepColorTemperatureParamsClassOnce sync.Once
)

func getMTRColorControlClusterStepColorTemperatureParamsClass() _MTRColorControlClusterStepColorTemperatureParamsClass {
	MTRColorControlClusterStepColorTemperatureParamsClassOnce.Do(func() {
		MTRColorControlClusterStepColorTemperatureParamsClass = _MTRColorControlClusterStepColorTemperatureParamsClass{objc.GetClass("MTRColorControlClusterStepColorTemperatureParams")}
	})
	return MTRColorControlClusterStepColorTemperatureParamsClass
}

type _MTRColorControlClusterStepColorTemperatureParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterStepColorTemperatureParams] class.
type IMTRColorControlClusterStepColorTemperatureParams interface {
	objectivec.IObject
	// properties:
	ColorTemperatureMaximumMireds() objc.IObject /* cross-framework: NSNumber */
	SetColorTemperatureMaximumMireds(value objc.IObject /* cross-framework: NSNumber */)
	ColorTemperatureMinimumMireds() objc.IObject /* cross-framework: NSNumber */
	SetColorTemperatureMinimumMireds(value objc.IObject /* cross-framework: NSNumber */)
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams
type MTRColorControlClusterStepColorTemperatureParams struct {
	objectivec.Object
}

// MTRColorControlClusterStepColorTemperatureParamsFrom constructs a [MTRColorControlClusterStepColorTemperatureParams] from an unsafe.Pointer.
func MTRColorControlClusterStepColorTemperatureParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterStepColorTemperatureParams {
	return MTRColorControlClusterStepColorTemperatureParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterStepColorTemperatureParamsClass) Alloc() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterStepColorTemperatureParamsClass) New() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterStepColorTemperatureParams) Init() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterStepColorTemperatureParams) Autorelease() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterStepColorTemperatureParams creates a new MTRColorControlClusterStepColorTemperatureParams instance.
func NewMTRColorControlClusterStepColorTemperatureParams() MTRColorControlClusterStepColorTemperatureParams {
	return getMTRColorControlClusterStepColorTemperatureParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/colortemperaturemaximummireds
func (m_ MTRColorControlClusterStepColorTemperatureParams) ColorTemperatureMaximumMireds() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorTemperatureMaximumMireds"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/colortemperaturemaximummireds
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetColorTemperatureMaximumMireds(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMaximumMireds:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/colortemperatureminimummireds
func (m_ MTRColorControlClusterStepColorTemperatureParams) ColorTemperatureMinimumMireds() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorTemperatureMinimumMireds"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/colortemperatureminimummireds
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetColorTemperatureMinimumMireds(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMinimumMireds:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/optionsmask
func (m_ MTRColorControlClusterStepColorTemperatureParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/optionsmask
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/optionsoverride
func (m_ MTRColorControlClusterStepColorTemperatureParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/optionsoverride
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterStepColorTemperatureParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/stepmode
func (m_ MTRColorControlClusterStepColorTemperatureParams) StepMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/stepmode
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetStepMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/stepsize
func (m_ MTRColorControlClusterStepColorTemperatureParams) StepSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/stepsize
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetStepSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepSize:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterStepColorTemperatureParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/transitiontime
func (m_ MTRColorControlClusterStepColorTemperatureParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterstepcolortemperatureparams/transitiontime
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



