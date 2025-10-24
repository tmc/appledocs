// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterMoveToLevelParams] class.
var (
	MTRLevelControlClusterMoveToLevelParamsClass     _MTRLevelControlClusterMoveToLevelParamsClass
	MTRLevelControlClusterMoveToLevelParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveToLevelParamsClass() _MTRLevelControlClusterMoveToLevelParamsClass {
	MTRLevelControlClusterMoveToLevelParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveToLevelParamsClass = _MTRLevelControlClusterMoveToLevelParamsClass{objc.GetClass("MTRLevelControlClusterMoveToLevelParams")}
	})
	return MTRLevelControlClusterMoveToLevelParamsClass
}

type _MTRLevelControlClusterMoveToLevelParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterMoveToLevelParams] class.
type IMTRLevelControlClusterMoveToLevelParams interface {
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams
type MTRLevelControlClusterMoveToLevelParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveToLevelParamsFrom constructs a [MTRLevelControlClusterMoveToLevelParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveToLevelParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveToLevelParams {
	return MTRLevelControlClusterMoveToLevelParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveToLevelParamsClass) Alloc() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterMoveToLevelParamsClass) New() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveToLevelParams) Init() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveToLevelParams) Autorelease() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveToLevelParams creates a new MTRLevelControlClusterMoveToLevelParams instance.
func NewMTRLevelControlClusterMoveToLevelParams() MTRLevelControlClusterMoveToLevelParams {
	return getMTRLevelControlClusterMoveToLevelParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/level
func (m_ MTRLevelControlClusterMoveToLevelParams) Level() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("level"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/level
func (m_ MTRLevelControlClusterMoveToLevelParams) SetLevel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLevel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/optionsmask
func (m_ MTRLevelControlClusterMoveToLevelParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/optionsmask
func (m_ MTRLevelControlClusterMoveToLevelParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/optionsoverride
func (m_ MTRLevelControlClusterMoveToLevelParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/optionsoverride
func (m_ MTRLevelControlClusterMoveToLevelParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveToLevelParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterMoveToLevelParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveToLevelParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterMoveToLevelParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/transitiontime
func (m_ MTRLevelControlClusterMoveToLevelParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclustermovetolevelparams/transitiontime
func (m_ MTRLevelControlClusterMoveToLevelParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



