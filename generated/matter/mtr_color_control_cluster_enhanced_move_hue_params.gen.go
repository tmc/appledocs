// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterEnhancedMoveHueParams] class.
var (
	MTRColorControlClusterEnhancedMoveHueParamsClass     _MTRColorControlClusterEnhancedMoveHueParamsClass
	MTRColorControlClusterEnhancedMoveHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterEnhancedMoveHueParamsClass() _MTRColorControlClusterEnhancedMoveHueParamsClass {
	MTRColorControlClusterEnhancedMoveHueParamsClassOnce.Do(func() {
		MTRColorControlClusterEnhancedMoveHueParamsClass = _MTRColorControlClusterEnhancedMoveHueParamsClass{objc.GetClass("MTRColorControlClusterEnhancedMoveHueParams")}
	})
	return MTRColorControlClusterEnhancedMoveHueParamsClass
}

type _MTRColorControlClusterEnhancedMoveHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterEnhancedMoveHueParams] class.
type IMTRColorControlClusterEnhancedMoveHueParams interface {
	objectivec.IObject
	// properties:
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams
type MTRColorControlClusterEnhancedMoveHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterEnhancedMoveHueParamsFrom constructs a [MTRColorControlClusterEnhancedMoveHueParams] from an unsafe.Pointer.
func MTRColorControlClusterEnhancedMoveHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterEnhancedMoveHueParams {
	return MTRColorControlClusterEnhancedMoveHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterEnhancedMoveHueParamsClass) Alloc() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterEnhancedMoveHueParamsClass) New() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterEnhancedMoveHueParams) Init() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterEnhancedMoveHueParams) Autorelease() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterEnhancedMoveHueParams creates a new MTRColorControlClusterEnhancedMoveHueParams instance.
func NewMTRColorControlClusterEnhancedMoveHueParams() MTRColorControlClusterEnhancedMoveHueParams {
	return getMTRColorControlClusterEnhancedMoveHueParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/movemode
func (m_ MTRColorControlClusterEnhancedMoveHueParams) MoveMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("moveMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/movemode
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetMoveMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/optionsmask
func (m_ MTRColorControlClusterEnhancedMoveHueParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/optionsmask
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/optionsoverride
func (m_ MTRColorControlClusterEnhancedMoveHueParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/optionsoverride
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/rate
func (m_ MTRColorControlClusterEnhancedMoveHueParams) Rate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/rate
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterEnhancedMoveHueParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterEnhancedMoveHueParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclusterenhancedmovehueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



