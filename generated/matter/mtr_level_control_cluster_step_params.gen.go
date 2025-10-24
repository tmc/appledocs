// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterStepParams] class.
var (
	MTRLevelControlClusterStepParamsClass     _MTRLevelControlClusterStepParamsClass
	MTRLevelControlClusterStepParamsClassOnce sync.Once
)

func getMTRLevelControlClusterStepParamsClass() _MTRLevelControlClusterStepParamsClass {
	MTRLevelControlClusterStepParamsClassOnce.Do(func() {
		MTRLevelControlClusterStepParamsClass = _MTRLevelControlClusterStepParamsClass{objc.GetClass("MTRLevelControlClusterStepParams")}
	})
	return MTRLevelControlClusterStepParamsClass
}

type _MTRLevelControlClusterStepParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterStepParams] class.
type IMTRLevelControlClusterStepParams interface {
	objectivec.IObject
	// properties:
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StepMode() objc.IObject /* cross-framework: NSNumber */
	SetStepMode(value objc.IObject /* cross-framework: NSNumber */)
	StepSize() objc.IObject /* cross-framework: NSNumber */
	SetStepSize(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TransitionTime() objc.IObject /* cross-framework: NSNumber */
	SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams
type MTRLevelControlClusterStepParams struct {
	objectivec.Object
}

// MTRLevelControlClusterStepParamsFrom constructs a [MTRLevelControlClusterStepParams] from an unsafe.Pointer.
func MTRLevelControlClusterStepParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterStepParams {
	return MTRLevelControlClusterStepParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterStepParamsClass) Alloc() MTRLevelControlClusterStepParams {
	rv := objc.Send[MTRLevelControlClusterStepParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterStepParamsClass) New() MTRLevelControlClusterStepParams {
	rv := objc.Send[MTRLevelControlClusterStepParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterStepParams) Init() MTRLevelControlClusterStepParams {
	rv := objc.Send[MTRLevelControlClusterStepParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterStepParams) Autorelease() MTRLevelControlClusterStepParams {
	rv := objc.Send[MTRLevelControlClusterStepParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterStepParams creates a new MTRLevelControlClusterStepParams instance.
func NewMTRLevelControlClusterStepParams() MTRLevelControlClusterStepParams {
	return getMTRLevelControlClusterStepParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/optionsmask
func (m_ MTRLevelControlClusterStepParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/optionsmask
func (m_ MTRLevelControlClusterStepParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/optionsoverride
func (m_ MTRLevelControlClusterStepParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/optionsoverride
func (m_ MTRLevelControlClusterStepParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterStepParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterStepParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/stepmode
func (m_ MTRLevelControlClusterStepParams) StepMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/stepmode
func (m_ MTRLevelControlClusterStepParams) SetStepMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/stepsize
func (m_ MTRLevelControlClusterStepParams) StepSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/stepsize
func (m_ MTRLevelControlClusterStepParams) SetStepSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepSize:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterStepParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterStepParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/transitiontime
func (m_ MTRLevelControlClusterStepParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepparams/transitiontime
func (m_ MTRLevelControlClusterStepParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



