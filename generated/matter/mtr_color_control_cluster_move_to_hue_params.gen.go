// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveToHueParams] class.
var (
	MTRColorControlClusterMoveToHueParamsClass     _MTRColorControlClusterMoveToHueParamsClass
	MTRColorControlClusterMoveToHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToHueParamsClass() _MTRColorControlClusterMoveToHueParamsClass {
	MTRColorControlClusterMoveToHueParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToHueParamsClass = _MTRColorControlClusterMoveToHueParamsClass{objc.GetClass("MTRColorControlClusterMoveToHueParams")}
	})
	return MTRColorControlClusterMoveToHueParamsClass
}

type _MTRColorControlClusterMoveToHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveToHueParams] class.
type IMTRColorControlClusterMoveToHueParams interface {
	objectivec.IObject
	// properties:
	Direction() objc.IObject /* cross-framework: NSNumber */
	SetDirection(value objc.IObject /* cross-framework: NSNumber */)
	Hue() objc.IObject /* cross-framework: NSNumber */
	SetHue(value objc.IObject /* cross-framework: NSNumber */)
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TransitionTime() objc.IObject /* cross-framework: NSNumber */
	SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams
type MTRColorControlClusterMoveToHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToHueParamsFrom constructs a [MTRColorControlClusterMoveToHueParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToHueParams {
	return MTRColorControlClusterMoveToHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToHueParamsClass) Alloc() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveToHueParamsClass) New() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToHueParams) Init() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToHueParams) Autorelease() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToHueParams creates a new MTRColorControlClusterMoveToHueParams instance.
func NewMTRColorControlClusterMoveToHueParams() MTRColorControlClusterMoveToHueParams {
	return getMTRColorControlClusterMoveToHueParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/direction
func (m_ MTRColorControlClusterMoveToHueParams) Direction() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("direction"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/direction
func (m_ MTRColorControlClusterMoveToHueParams) SetDirection(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDirection:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/hue
func (m_ MTRColorControlClusterMoveToHueParams) Hue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("hue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/hue
func (m_ MTRColorControlClusterMoveToHueParams) SetHue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/optionsmask
func (m_ MTRColorControlClusterMoveToHueParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/optionsmask
func (m_ MTRColorControlClusterMoveToHueParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/optionsoverride
func (m_ MTRColorControlClusterMoveToHueParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/optionsoverride
func (m_ MTRColorControlClusterMoveToHueParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToHueParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToHueParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToHueParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToHueParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/transitiontime
func (m_ MTRColorControlClusterMoveToHueParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetohueparams/transitiontime
func (m_ MTRColorControlClusterMoveToHueParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



