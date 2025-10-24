// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterColorLoopSetParams] class.
var (
	MTRColorControlClusterColorLoopSetParamsClass     _MTRColorControlClusterColorLoopSetParamsClass
	MTRColorControlClusterColorLoopSetParamsClassOnce sync.Once
)

func getMTRColorControlClusterColorLoopSetParamsClass() _MTRColorControlClusterColorLoopSetParamsClass {
	MTRColorControlClusterColorLoopSetParamsClassOnce.Do(func() {
		MTRColorControlClusterColorLoopSetParamsClass = _MTRColorControlClusterColorLoopSetParamsClass{objc.GetClass("MTRColorControlClusterColorLoopSetParams")}
	})
	return MTRColorControlClusterColorLoopSetParamsClass
}

type _MTRColorControlClusterColorLoopSetParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterColorLoopSetParams] class.
type IMTRColorControlClusterColorLoopSetParams interface {
	objectivec.IObject
	// properties:
	Action() objc.IObject /* cross-framework: NSNumber */
	SetAction(value objc.IObject /* cross-framework: NSNumber */)
	Direction() objc.IObject /* cross-framework: NSNumber */
	SetDirection(value objc.IObject /* cross-framework: NSNumber */)
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StartHue() objc.IObject /* cross-framework: NSNumber */
	SetStartHue(value objc.IObject /* cross-framework: NSNumber */)
	Time() objc.IObject /* cross-framework: NSNumber */
	SetTime(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UpdateFlags() objc.IObject /* cross-framework: NSNumber */
	SetUpdateFlags(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams
type MTRColorControlClusterColorLoopSetParams struct {
	objectivec.Object
}

// MTRColorControlClusterColorLoopSetParamsFrom constructs a [MTRColorControlClusterColorLoopSetParams] from an unsafe.Pointer.
func MTRColorControlClusterColorLoopSetParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterColorLoopSetParams {
	return MTRColorControlClusterColorLoopSetParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterColorLoopSetParamsClass) Alloc() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterColorLoopSetParamsClass) New() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterColorLoopSetParams) Init() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterColorLoopSetParams) Autorelease() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterColorLoopSetParams creates a new MTRColorControlClusterColorLoopSetParams instance.
func NewMTRColorControlClusterColorLoopSetParams() MTRColorControlClusterColorLoopSetParams {
	return getMTRColorControlClusterColorLoopSetParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/action
func (m_ MTRColorControlClusterColorLoopSetParams) Action() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("action"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/action
func (m_ MTRColorControlClusterColorLoopSetParams) SetAction(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/direction
func (m_ MTRColorControlClusterColorLoopSetParams) Direction() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("direction"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/direction
func (m_ MTRColorControlClusterColorLoopSetParams) SetDirection(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDirection:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/optionsmask
func (m_ MTRColorControlClusterColorLoopSetParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/optionsmask
func (m_ MTRColorControlClusterColorLoopSetParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/optionsoverride
func (m_ MTRColorControlClusterColorLoopSetParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/optionsoverride
func (m_ MTRColorControlClusterColorLoopSetParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterColorLoopSetParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterColorLoopSetParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/starthue
func (m_ MTRColorControlClusterColorLoopSetParams) StartHue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startHue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/starthue
func (m_ MTRColorControlClusterColorLoopSetParams) SetStartHue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartHue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/time
func (m_ MTRColorControlClusterColorLoopSetParams) Time() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("time"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/time
func (m_ MTRColorControlClusterColorLoopSetParams) SetTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterColorLoopSetParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterColorLoopSetParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/updateflags
func (m_ MTRColorControlClusterColorLoopSetParams) UpdateFlags() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("updateFlags"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustercolorloopsetparams/updateflags
func (m_ MTRColorControlClusterColorLoopSetParams) SetUpdateFlags(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateFlags:"), value)
}



