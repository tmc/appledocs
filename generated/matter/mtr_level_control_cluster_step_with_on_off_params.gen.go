// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterStepWithOnOffParams] class.
var (
	MTRLevelControlClusterStepWithOnOffParamsClass     _MTRLevelControlClusterStepWithOnOffParamsClass
	MTRLevelControlClusterStepWithOnOffParamsClassOnce sync.Once
)

func getMTRLevelControlClusterStepWithOnOffParamsClass() _MTRLevelControlClusterStepWithOnOffParamsClass {
	MTRLevelControlClusterStepWithOnOffParamsClassOnce.Do(func() {
		MTRLevelControlClusterStepWithOnOffParamsClass = _MTRLevelControlClusterStepWithOnOffParamsClass{objc.GetClass("MTRLevelControlClusterStepWithOnOffParams")}
	})
	return MTRLevelControlClusterStepWithOnOffParamsClass
}

type _MTRLevelControlClusterStepWithOnOffParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterStepWithOnOffParams] class.
type IMTRLevelControlClusterStepWithOnOffParams interface {
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams
type MTRLevelControlClusterStepWithOnOffParams struct {
	objectivec.Object
}

// MTRLevelControlClusterStepWithOnOffParamsFrom constructs a [MTRLevelControlClusterStepWithOnOffParams] from an unsafe.Pointer.
func MTRLevelControlClusterStepWithOnOffParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterStepWithOnOffParams {
	return MTRLevelControlClusterStepWithOnOffParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterStepWithOnOffParamsClass) Alloc() MTRLevelControlClusterStepWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStepWithOnOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterStepWithOnOffParamsClass) New() MTRLevelControlClusterStepWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStepWithOnOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterStepWithOnOffParams) Init() MTRLevelControlClusterStepWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStepWithOnOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterStepWithOnOffParams) Autorelease() MTRLevelControlClusterStepWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStepWithOnOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterStepWithOnOffParams creates a new MTRLevelControlClusterStepWithOnOffParams instance.
func NewMTRLevelControlClusterStepWithOnOffParams() MTRLevelControlClusterStepWithOnOffParams {
	return getMTRLevelControlClusterStepWithOnOffParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/optionsmask
func (m_ MTRLevelControlClusterStepWithOnOffParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/optionsmask
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/optionsoverride
func (m_ MTRLevelControlClusterStepWithOnOffParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/optionsoverride
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterStepWithOnOffParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/serversideprocessingtimeout
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/stepmode
func (m_ MTRLevelControlClusterStepWithOnOffParams) StepMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/stepmode
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetStepMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/stepsize
func (m_ MTRLevelControlClusterStepWithOnOffParams) StepSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/stepsize
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetStepSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepSize:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterStepWithOnOffParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/timedinvoketimeoutms
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/transitiontime
func (m_ MTRLevelControlClusterStepWithOnOffParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlevelcontrolclusterstepwithonoffparams/transitiontime
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



