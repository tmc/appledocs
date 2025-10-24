// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveColorTemperatureParams] class.
var (
	MTRColorControlClusterMoveColorTemperatureParamsClass     _MTRColorControlClusterMoveColorTemperatureParamsClass
	MTRColorControlClusterMoveColorTemperatureParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveColorTemperatureParamsClass() _MTRColorControlClusterMoveColorTemperatureParamsClass {
	MTRColorControlClusterMoveColorTemperatureParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveColorTemperatureParamsClass = _MTRColorControlClusterMoveColorTemperatureParamsClass{objc.GetClass("MTRColorControlClusterMoveColorTemperatureParams")}
	})
	return MTRColorControlClusterMoveColorTemperatureParamsClass
}

type _MTRColorControlClusterMoveColorTemperatureParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveColorTemperatureParams] class.
type IMTRColorControlClusterMoveColorTemperatureParams interface {
	objectivec.IObject
	// properties:
	ColorTemperatureMaximumMireds() objc.IObject /* cross-framework: NSNumber */
	SetColorTemperatureMaximumMireds(value objc.IObject /* cross-framework: NSNumber */)
	ColorTemperatureMinimumMireds() objc.IObject /* cross-framework: NSNumber */
	SetColorTemperatureMinimumMireds(value objc.IObject /* cross-framework: NSNumber */)
	MoveMode() objc.IObject /* cross-framework: NSNumber */
	SetMoveMode(value objc.IObject /* cross-framework: NSNumber */)
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	Rate() objc.IObject /* cross-framework: NSNumber */
	SetRate(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams
type MTRColorControlClusterMoveColorTemperatureParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveColorTemperatureParamsFrom constructs a [MTRColorControlClusterMoveColorTemperatureParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveColorTemperatureParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveColorTemperatureParams {
	return MTRColorControlClusterMoveColorTemperatureParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveColorTemperatureParamsClass) Alloc() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveColorTemperatureParamsClass) New() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveColorTemperatureParams) Init() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveColorTemperatureParams) Autorelease() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveColorTemperatureParams creates a new MTRColorControlClusterMoveColorTemperatureParams instance.
func NewMTRColorControlClusterMoveColorTemperatureParams() MTRColorControlClusterMoveColorTemperatureParams {
	return getMTRColorControlClusterMoveColorTemperatureParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/colortemperaturemaximummireds
func (m_ MTRColorControlClusterMoveColorTemperatureParams) ColorTemperatureMaximumMireds() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorTemperatureMaximumMireds"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/colortemperaturemaximummireds
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetColorTemperatureMaximumMireds(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMaximumMireds:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/colortemperatureminimummireds
func (m_ MTRColorControlClusterMoveColorTemperatureParams) ColorTemperatureMinimumMireds() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorTemperatureMinimumMireds"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/colortemperatureminimummireds
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetColorTemperatureMinimumMireds(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMinimumMireds:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/movemode
func (m_ MTRColorControlClusterMoveColorTemperatureParams) MoveMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("moveMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/movemode
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetMoveMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/optionsmask
func (m_ MTRColorControlClusterMoveColorTemperatureParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/optionsmask
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/optionsoverride
func (m_ MTRColorControlClusterMoveColorTemperatureParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/optionsoverride
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/rate
func (m_ MTRColorControlClusterMoveColorTemperatureParams) Rate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/rate
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveColorTemperatureParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveColorTemperatureParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovecolortemperatureparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



