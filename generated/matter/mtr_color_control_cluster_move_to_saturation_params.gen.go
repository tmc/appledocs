// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveToSaturationParams] class.
var (
	MTRColorControlClusterMoveToSaturationParamsClass     _MTRColorControlClusterMoveToSaturationParamsClass
	MTRColorControlClusterMoveToSaturationParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToSaturationParamsClass() _MTRColorControlClusterMoveToSaturationParamsClass {
	MTRColorControlClusterMoveToSaturationParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToSaturationParamsClass = _MTRColorControlClusterMoveToSaturationParamsClass{objc.GetClass("MTRColorControlClusterMoveToSaturationParams")}
	})
	return MTRColorControlClusterMoveToSaturationParamsClass
}

type _MTRColorControlClusterMoveToSaturationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveToSaturationParams] class.
type IMTRColorControlClusterMoveToSaturationParams interface {
	objectivec.IObject
	// properties:
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams
type MTRColorControlClusterMoveToSaturationParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToSaturationParamsFrom constructs a [MTRColorControlClusterMoveToSaturationParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToSaturationParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToSaturationParams {
	return MTRColorControlClusterMoveToSaturationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToSaturationParamsClass) Alloc() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveToSaturationParamsClass) New() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToSaturationParams) Init() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToSaturationParams) Autorelease() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToSaturationParams creates a new MTRColorControlClusterMoveToSaturationParams instance.
func NewMTRColorControlClusterMoveToSaturationParams() MTRColorControlClusterMoveToSaturationParams {
	return getMTRColorControlClusterMoveToSaturationParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/optionsmask
func (m_ MTRColorControlClusterMoveToSaturationParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/optionsmask
func (m_ MTRColorControlClusterMoveToSaturationParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/optionsoverride
func (m_ MTRColorControlClusterMoveToSaturationParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/optionsoverride
func (m_ MTRColorControlClusterMoveToSaturationParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/saturation
func (m_ MTRColorControlClusterMoveToSaturationParams) Saturation() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("saturation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/saturation
func (m_ MTRColorControlClusterMoveToSaturationParams) SetSaturation(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSaturation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToSaturationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToSaturationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToSaturationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToSaturationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/transitiontime
func (m_ MTRColorControlClusterMoveToSaturationParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetosaturationparams/transitiontime
func (m_ MTRColorControlClusterMoveToSaturationParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



