// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveToHueAndSaturationParams] class.
var (
	MTRColorControlClusterMoveToHueAndSaturationParamsClass     _MTRColorControlClusterMoveToHueAndSaturationParamsClass
	MTRColorControlClusterMoveToHueAndSaturationParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToHueAndSaturationParamsClass() _MTRColorControlClusterMoveToHueAndSaturationParamsClass {
	MTRColorControlClusterMoveToHueAndSaturationParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToHueAndSaturationParamsClass = _MTRColorControlClusterMoveToHueAndSaturationParamsClass{objc.GetClass("MTRColorControlClusterMoveToHueAndSaturationParams")}
	})
	return MTRColorControlClusterMoveToHueAndSaturationParamsClass
}

type _MTRColorControlClusterMoveToHueAndSaturationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveToHueAndSaturationParams] class.
type IMTRColorControlClusterMoveToHueAndSaturationParams interface {
	objectivec.IObject
	// properties:
	Hue() objc.IObject /* cross-framework: NSNumber */
	SetHue(value objc.IObject /* cross-framework: NSNumber */)
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams
type MTRColorControlClusterMoveToHueAndSaturationParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToHueAndSaturationParamsFrom constructs a [MTRColorControlClusterMoveToHueAndSaturationParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToHueAndSaturationParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToHueAndSaturationParams {
	return MTRColorControlClusterMoveToHueAndSaturationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToHueAndSaturationParamsClass) Alloc() MTRColorControlClusterMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueAndSaturationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveToHueAndSaturationParamsClass) New() MTRColorControlClusterMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueAndSaturationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) Init() MTRColorControlClusterMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueAndSaturationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) Autorelease() MTRColorControlClusterMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueAndSaturationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToHueAndSaturationParams creates a new MTRColorControlClusterMoveToHueAndSaturationParams instance.
func NewMTRColorControlClusterMoveToHueAndSaturationParams() MTRColorControlClusterMoveToHueAndSaturationParams {
	return getMTRColorControlClusterMoveToHueAndSaturationParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueandsaturationparams/hue
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) Hue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("hue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueandsaturationparams/hue
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) SetHue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueandsaturationparams/optionsmask
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueandsaturationparams/optionsmask
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueandsaturationparams/optionsoverride
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueandsaturationparams/optionsoverride
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueandsaturationparams/saturation
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) Saturation() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("saturation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueandsaturationparams/saturation
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) SetSaturation(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSaturation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueandsaturationparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueandsaturationparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueandsaturationparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueandsaturationparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueandsaturationparams/transitiontime
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueandsaturationparams/transitiontime
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



