// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterEnhancedMoveToHueAndSaturationParams] class.
var (
	MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass     _MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass
	MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClassOnce sync.Once
)

func getMTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass() _MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass {
	MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClassOnce.Do(func() {
		MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass = _MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass{objc.GetClass("MTRColorControlClusterEnhancedMoveToHueAndSaturationParams")}
	})
	return MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass
}

type _MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterEnhancedMoveToHueAndSaturationParams] class.
type IMTRColorControlClusterEnhancedMoveToHueAndSaturationParams interface {
	objectivec.IObject
	// properties:
	EnhancedHue() objc.IObject /* cross-framework: NSNumber */
	SetEnhancedHue(value objc.IObject /* cross-framework: NSNumber */)
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	Saturation() objc.IObject /* cross-framework: NSNumber */
	SetSaturation(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TransitionTime() objc.IObject /* cross-framework: NSNumber */
	SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams
type MTRColorControlClusterEnhancedMoveToHueAndSaturationParams struct {
	objectivec.Object
}

// MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsFrom constructs a [MTRColorControlClusterEnhancedMoveToHueAndSaturationParams] from an unsafe.Pointer.
func MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterEnhancedMoveToHueAndSaturationParams {
	return MTRColorControlClusterEnhancedMoveToHueAndSaturationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass) Alloc() MTRColorControlClusterEnhancedMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueAndSaturationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass) New() MTRColorControlClusterEnhancedMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueAndSaturationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) Init() MTRColorControlClusterEnhancedMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueAndSaturationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) Autorelease() MTRColorControlClusterEnhancedMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueAndSaturationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterEnhancedMoveToHueAndSaturationParams creates a new MTRColorControlClusterEnhancedMoveToHueAndSaturationParams instance.
func NewMTRColorControlClusterEnhancedMoveToHueAndSaturationParams() MTRColorControlClusterEnhancedMoveToHueAndSaturationParams {
	return getMTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueandsaturationparams/enhancedhue
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) EnhancedHue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("enhancedHue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueandsaturationparams/enhancedhue
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) SetEnhancedHue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnhancedHue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueandsaturationparams/optionsmask
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueandsaturationparams/optionsmask
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueandsaturationparams/optionsoverride
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueandsaturationparams/optionsoverride
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueandsaturationparams/saturation
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) Saturation() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("saturation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueandsaturationparams/saturation
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) SetSaturation(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSaturation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueandsaturationparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueandsaturationparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueandsaturationparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueandsaturationparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueandsaturationparams/transitiontime
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueandsaturationparams/transitiontime
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



