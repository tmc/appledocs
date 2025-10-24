// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveToColorTemperatureParams] class.
var (
	MTRColorControlClusterMoveToColorTemperatureParamsClass     _MTRColorControlClusterMoveToColorTemperatureParamsClass
	MTRColorControlClusterMoveToColorTemperatureParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToColorTemperatureParamsClass() _MTRColorControlClusterMoveToColorTemperatureParamsClass {
	MTRColorControlClusterMoveToColorTemperatureParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToColorTemperatureParamsClass = _MTRColorControlClusterMoveToColorTemperatureParamsClass{objc.GetClass("MTRColorControlClusterMoveToColorTemperatureParams")}
	})
	return MTRColorControlClusterMoveToColorTemperatureParamsClass
}

type _MTRColorControlClusterMoveToColorTemperatureParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveToColorTemperatureParams] class.
type IMTRColorControlClusterMoveToColorTemperatureParams interface {
	objectivec.IObject
	// properties:
	ColorTemperature() objc.IObject /* cross-framework: NSNumber */
	SetColorTemperature(value objc.IObject /* cross-framework: NSNumber */)
	ColorTemperatureMireds() objc.IObject /* cross-framework: NSNumber */
	SetColorTemperatureMireds(value objc.IObject /* cross-framework: NSNumber */)
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams
type MTRColorControlClusterMoveToColorTemperatureParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToColorTemperatureParamsFrom constructs a [MTRColorControlClusterMoveToColorTemperatureParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToColorTemperatureParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToColorTemperatureParams {
	return MTRColorControlClusterMoveToColorTemperatureParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToColorTemperatureParamsClass) Alloc() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveToColorTemperatureParamsClass) New() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) Init() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) Autorelease() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToColorTemperatureParams creates a new MTRColorControlClusterMoveToColorTemperatureParams instance.
func NewMTRColorControlClusterMoveToColorTemperatureParams() MTRColorControlClusterMoveToColorTemperatureParams {
	return getMTRColorControlClusterMoveToColorTemperatureParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/colortemperature
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) ColorTemperature() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorTemperature"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/colortemperature
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetColorTemperature(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperature:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/colortemperaturemireds
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) ColorTemperatureMireds() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorTemperatureMireds"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/colortemperaturemireds
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetColorTemperatureMireds(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMireds:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/optionsmask
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/optionsmask
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/optionsoverride
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/optionsoverride
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/transitiontime
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovetocolortemperatureparams/transitiontime
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



