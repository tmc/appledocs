// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRLevelControlClusterStepWithOnOffParams */


/* debug [class_header]: Header for MTRLevelControlClusterStepWithOnOffParams */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRLevelControlClusterStepWithOnOffParams */
// An interface definition for the [MTRLevelControlClusterStepWithOnOffParams] class.
type IMTRLevelControlClusterStepWithOnOffParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRLevelControlClusterStepWithOnOffParams */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRLevelControlClusterStepWithOnOffParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRLevelControlClusterStepWithOnOffParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterStepWithOnOffParamsClass) Alloc() MTRLevelControlClusterStepWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStepWithOnOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRLevelControlClusterStepWithOnOffParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams
type MTRLevelControlClusterStepWithOnOffParams struct {
	objectivec.Object
}

// MTRLevelControlClusterStepWithOnOffParamsFrom constructs a [MTRLevelControlClusterStepWithOnOffParams] from an unsafe.Pointer.
func MTRLevelControlClusterStepWithOnOffParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterStepWithOnOffParams {
	return MTRLevelControlClusterStepWithOnOffParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRLevelControlClusterStepWithOnOffParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRLevelControlClusterStepWithOnOffParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRLevelControlClusterStepWithOnOffParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRLevelControlClusterStepWithOnOffParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRLevelControlClusterStepWithOnOffParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams/optionsMask
func (m_ MTRLevelControlClusterStepWithOnOffParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams/optionsMask
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams/optionsOverride
func (m_ MTRLevelControlClusterStepWithOnOffParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams/optionsOverride
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterStepWithOnOffParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams/stepMode
func (m_ MTRLevelControlClusterStepWithOnOffParams) StepMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepMode"))
	return rv
}/* debug [instance_properties/getter]: stepMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams/stepMode
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetStepMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepMode:"), value)
}/* debug [instance_properties/setter]: stepMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams/stepSize
func (m_ MTRLevelControlClusterStepWithOnOffParams) StepSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepSize"))
	return rv
}/* debug [instance_properties/getter]: stepSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams/stepSize
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetStepSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepSize:"), value)
}/* debug [instance_properties/setter]: stepSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterStepWithOnOffParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams/transitionTime
func (m_ MTRLevelControlClusterStepWithOnOffParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams/transitionTime
func (m_ MTRLevelControlClusterStepWithOnOffParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRLevelControlClusterStepWithOnOffParams */



