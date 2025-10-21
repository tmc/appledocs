// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterEnhancedMoveToHueParams] class.
var (
	MTRColorControlClusterEnhancedMoveToHueParamsClass     _MTRColorControlClusterEnhancedMoveToHueParamsClass
	MTRColorControlClusterEnhancedMoveToHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterEnhancedMoveToHueParamsClass() _MTRColorControlClusterEnhancedMoveToHueParamsClass {
	MTRColorControlClusterEnhancedMoveToHueParamsClassOnce.Do(func() {
		MTRColorControlClusterEnhancedMoveToHueParamsClass = _MTRColorControlClusterEnhancedMoveToHueParamsClass{objc.GetClass("MTRColorControlClusterEnhancedMoveToHueParams")}
	})
	return MTRColorControlClusterEnhancedMoveToHueParamsClass
}

type _MTRColorControlClusterEnhancedMoveToHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterEnhancedMoveToHueParams] class.
type IMTRColorControlClusterEnhancedMoveToHueParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams
type MTRColorControlClusterEnhancedMoveToHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterEnhancedMoveToHueParamsFrom constructs a [MTRColorControlClusterEnhancedMoveToHueParams] from an unsafe.Pointer.
func MTRColorControlClusterEnhancedMoveToHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterEnhancedMoveToHueParams {
	return MTRColorControlClusterEnhancedMoveToHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterEnhancedMoveToHueParamsClass) Alloc() MTRColorControlClusterEnhancedMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterEnhancedMoveToHueParamsClass) New() MTRColorControlClusterEnhancedMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) Init() MTRColorControlClusterEnhancedMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) Autorelease() MTRColorControlClusterEnhancedMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterEnhancedMoveToHueParams creates a new MTRColorControlClusterEnhancedMoveToHueParams instance.
func NewMTRColorControlClusterEnhancedMoveToHueParams() MTRColorControlClusterEnhancedMoveToHueParams {
	return getMTRColorControlClusterEnhancedMoveToHueParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueparams/direction
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) Direction() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("direction"))
	return rv
}


// SetDirection sets the value of the direction property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueparams/direction
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) SetDirection(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDirection:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueparams/enhancedhue
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) EnhancedHue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("enhancedHue"))
	return rv
}


// SetEnhancedHue sets the value of the enhancedHue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueparams/enhancedhue
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) SetEnhancedHue(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnhancedHue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueparams/optionsmask
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) OptionsMask() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// SetOptionsMask sets the value of the optionsMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueparams/optionsmask
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) SetOptionsMask(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueparams/optionsoverride
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) OptionsOverride() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// SetOptionsOverride sets the value of the optionsOverride property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueparams/optionsoverride
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) SetOptionsOverride(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueparams/transitiontime
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) TransitionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// SetTransitionTime sets the value of the transitionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovetohueparams/transitiontime
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) SetTransitionTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



