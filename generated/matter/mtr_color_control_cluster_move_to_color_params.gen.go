// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveToColorParams] class.
var (
	MTRColorControlClusterMoveToColorParamsClass     _MTRColorControlClusterMoveToColorParamsClass
	MTRColorControlClusterMoveToColorParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToColorParamsClass() _MTRColorControlClusterMoveToColorParamsClass {
	MTRColorControlClusterMoveToColorParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToColorParamsClass = _MTRColorControlClusterMoveToColorParamsClass{objc.GetClass("MTRColorControlClusterMoveToColorParams")}
	})
	return MTRColorControlClusterMoveToColorParamsClass
}

type _MTRColorControlClusterMoveToColorParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveToColorParams] class.
type IMTRColorControlClusterMoveToColorParams interface {
	objectivec.IObject
	// properties:
	ColorX() objc.IObject /* cross-framework: NSNumber */
	SetColorX(value objc.IObject /* cross-framework: NSNumber */)
	ColorY() objc.IObject /* cross-framework: NSNumber */
	SetColorY(value objc.IObject /* cross-framework: NSNumber */)
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams
type MTRColorControlClusterMoveToColorParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToColorParamsFrom constructs a [MTRColorControlClusterMoveToColorParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToColorParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToColorParams {
	return MTRColorControlClusterMoveToColorParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToColorParamsClass) Alloc() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveToColorParamsClass) New() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToColorParams) Init() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToColorParams) Autorelease() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToColorParams creates a new MTRColorControlClusterMoveToColorParams instance.
func NewMTRColorControlClusterMoveToColorParams() MTRColorControlClusterMoveToColorParams {
	return getMTRColorControlClusterMoveToColorParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/colorx
func (m_ MTRColorControlClusterMoveToColorParams) ColorX() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/colorx
func (m_ MTRColorControlClusterMoveToColorParams) SetColorX(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/colory
func (m_ MTRColorControlClusterMoveToColorParams) ColorY() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/colory
func (m_ MTRColorControlClusterMoveToColorParams) SetColorY(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/optionsmask
func (m_ MTRColorControlClusterMoveToColorParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/optionsmask
func (m_ MTRColorControlClusterMoveToColorParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/optionsoverride
func (m_ MTRColorControlClusterMoveToColorParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/optionsoverride
func (m_ MTRColorControlClusterMoveToColorParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToColorParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToColorParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToColorParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToColorParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/transitiontime
func (m_ MTRColorControlClusterMoveToColorParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolorparams/transitiontime
func (m_ MTRColorControlClusterMoveToColorParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



