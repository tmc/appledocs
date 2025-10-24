// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveHueParams] class.
var (
	MTRColorControlClusterMoveHueParamsClass     _MTRColorControlClusterMoveHueParamsClass
	MTRColorControlClusterMoveHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveHueParamsClass() _MTRColorControlClusterMoveHueParamsClass {
	MTRColorControlClusterMoveHueParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveHueParamsClass = _MTRColorControlClusterMoveHueParamsClass{objc.GetClass("MTRColorControlClusterMoveHueParams")}
	})
	return MTRColorControlClusterMoveHueParamsClass
}

type _MTRColorControlClusterMoveHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveHueParams] class.
type IMTRColorControlClusterMoveHueParams interface {
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams
type MTRColorControlClusterMoveHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveHueParamsFrom constructs a [MTRColorControlClusterMoveHueParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveHueParams {
	return MTRColorControlClusterMoveHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveHueParamsClass) Alloc() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveHueParamsClass) New() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveHueParams) Init() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveHueParams) Autorelease() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveHueParams creates a new MTRColorControlClusterMoveHueParams instance.
func NewMTRColorControlClusterMoveHueParams() MTRColorControlClusterMoveHueParams {
	return getMTRColorControlClusterMoveHueParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/movemode
func (m_ MTRColorControlClusterMoveHueParams) MoveMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("moveMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/movemode
func (m_ MTRColorControlClusterMoveHueParams) SetMoveMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/optionsmask
func (m_ MTRColorControlClusterMoveHueParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/optionsmask
func (m_ MTRColorControlClusterMoveHueParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/optionsoverride
func (m_ MTRColorControlClusterMoveHueParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/optionsoverride
func (m_ MTRColorControlClusterMoveHueParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/rate
func (m_ MTRColorControlClusterMoveHueParams) Rate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/rate
func (m_ MTRColorControlClusterMoveHueParams) SetRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveHueParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/serversideprocessingtimeout
func (m_ MTRColorControlClusterMoveHueParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveHueParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcolorcontrolclustermovehueparams/timedinvoketimeoutms
func (m_ MTRColorControlClusterMoveHueParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



