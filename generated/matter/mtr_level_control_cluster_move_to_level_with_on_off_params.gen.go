// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterMoveToLevelWithOnOffParams] class.
var (
	MTRLevelControlClusterMoveToLevelWithOnOffParamsClass     _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass
	MTRLevelControlClusterMoveToLevelWithOnOffParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveToLevelWithOnOffParamsClass() _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass {
	MTRLevelControlClusterMoveToLevelWithOnOffParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveToLevelWithOnOffParamsClass = _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass{objc.GetClass("MTRLevelControlClusterMoveToLevelWithOnOffParams")}
	})
	return MTRLevelControlClusterMoveToLevelWithOnOffParamsClass
}

type _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterMoveToLevelWithOnOffParams] class.
type IMTRLevelControlClusterMoveToLevelWithOnOffParams interface {
	objectivec.IObject
	// properties:
	Level() objc.IObject /* cross-framework: NSNumber */
	SetLevel(value objc.IObject /* cross-framework: NSNumber */)
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams
type MTRLevelControlClusterMoveToLevelWithOnOffParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveToLevelWithOnOffParamsFrom constructs a [MTRLevelControlClusterMoveToLevelWithOnOffParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveToLevelWithOnOffParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveToLevelWithOnOffParams {
	return MTRLevelControlClusterMoveToLevelWithOnOffParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass) Alloc() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass) New() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) Init() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) Autorelease() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveToLevelWithOnOffParams creates a new MTRLevelControlClusterMoveToLevelWithOnOffParams instance.
func NewMTRLevelControlClusterMoveToLevelWithOnOffParams() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	return getMTRLevelControlClusterMoveToLevelWithOnOffParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/level
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) Level() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("level"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/level
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetLevel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLevel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/optionsmask
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/optionsmask
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/optionsoverride
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/optionsoverride
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/transitiontime
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelwithonoffparams/transitiontime
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



