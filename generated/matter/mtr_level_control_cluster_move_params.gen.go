// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterMoveParams] class.
var (
	MTRLevelControlClusterMoveParamsClass     _MTRLevelControlClusterMoveParamsClass
	MTRLevelControlClusterMoveParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveParamsClass() _MTRLevelControlClusterMoveParamsClass {
	MTRLevelControlClusterMoveParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveParamsClass = _MTRLevelControlClusterMoveParamsClass{objc.GetClass("MTRLevelControlClusterMoveParams")}
	})
	return MTRLevelControlClusterMoveParamsClass
}

type _MTRLevelControlClusterMoveParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterMoveParams] class.
type IMTRLevelControlClusterMoveParams interface {
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams
type MTRLevelControlClusterMoveParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveParamsFrom constructs a [MTRLevelControlClusterMoveParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveParams {
	return MTRLevelControlClusterMoveParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveParamsClass) Alloc() MTRLevelControlClusterMoveParams {
	rv := objc.Send[MTRLevelControlClusterMoveParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterMoveParamsClass) New() MTRLevelControlClusterMoveParams {
	rv := objc.Send[MTRLevelControlClusterMoveParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveParams) Init() MTRLevelControlClusterMoveParams {
	rv := objc.Send[MTRLevelControlClusterMoveParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveParams) Autorelease() MTRLevelControlClusterMoveParams {
	rv := objc.Send[MTRLevelControlClusterMoveParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveParams creates a new MTRLevelControlClusterMoveParams instance.
func NewMTRLevelControlClusterMoveParams() MTRLevelControlClusterMoveParams {
	return getMTRLevelControlClusterMoveParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/movemode
func (m_ MTRLevelControlClusterMoveParams) MoveMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("moveMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/movemode
func (m_ MTRLevelControlClusterMoveParams) SetMoveMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/optionsmask
func (m_ MTRLevelControlClusterMoveParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/optionsmask
func (m_ MTRLevelControlClusterMoveParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/optionsoverride
func (m_ MTRLevelControlClusterMoveParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/optionsoverride
func (m_ MTRLevelControlClusterMoveParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/rate
func (m_ MTRLevelControlClusterMoveParams) Rate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/rate
func (m_ MTRLevelControlClusterMoveParams) SetRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermoveparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



