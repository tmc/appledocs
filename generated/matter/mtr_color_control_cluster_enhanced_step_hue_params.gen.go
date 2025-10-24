// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterEnhancedStepHueParams */


/* debug [class_header]: Header for MTRColorControlClusterEnhancedStepHueParams */
// The class instance for the [MTRColorControlClusterEnhancedStepHueParams] class.
var (
	MTRColorControlClusterEnhancedStepHueParamsClass     _MTRColorControlClusterEnhancedStepHueParamsClass
	MTRColorControlClusterEnhancedStepHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterEnhancedStepHueParamsClass() _MTRColorControlClusterEnhancedStepHueParamsClass {
	MTRColorControlClusterEnhancedStepHueParamsClassOnce.Do(func() {
		MTRColorControlClusterEnhancedStepHueParamsClass = _MTRColorControlClusterEnhancedStepHueParamsClass{objc.GetClass("MTRColorControlClusterEnhancedStepHueParams")}
	})
	return MTRColorControlClusterEnhancedStepHueParamsClass
}

type _MTRColorControlClusterEnhancedStepHueParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterEnhancedStepHueParams */
// An interface definition for the [MTRColorControlClusterEnhancedStepHueParams] class.
type IMTRColorControlClusterEnhancedStepHueParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterEnhancedStepHueParams */
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

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterEnhancedStepHueParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterEnhancedStepHueParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterEnhancedStepHueParamsClass) Alloc() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterEnhancedStepHueParamsClass) New() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterEnhancedStepHueParams) Init() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterEnhancedStepHueParams) Autorelease() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterEnhancedStepHueParams creates a new MTRColorControlClusterEnhancedStepHueParams instance.
func NewMTRColorControlClusterEnhancedStepHueParams() MTRColorControlClusterEnhancedStepHueParams {
	return getMTRColorControlClusterEnhancedStepHueParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterEnhancedStepHueParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams
type MTRColorControlClusterEnhancedStepHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterEnhancedStepHueParamsFrom constructs a [MTRColorControlClusterEnhancedStepHueParams] from an unsafe.Pointer.
func MTRColorControlClusterEnhancedStepHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterEnhancedStepHueParams {
	return MTRColorControlClusterEnhancedStepHueParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterEnhancedStepHueParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterEnhancedStepHueParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterEnhancedStepHueParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterEnhancedStepHueParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterEnhancedStepHueParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams/optionsMask
func (m_ MTRColorControlClusterEnhancedStepHueParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams/optionsMask
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams/optionsOverride
func (m_ MTRColorControlClusterEnhancedStepHueParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams/optionsOverride
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterEnhancedStepHueParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams/stepMode
func (m_ MTRColorControlClusterEnhancedStepHueParams) StepMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepMode"))
	return rv
}/* debug [instance_properties/getter]: stepMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams/stepMode
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetStepMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepMode:"), value)
}/* debug [instance_properties/setter]: stepMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams/stepSize
func (m_ MTRColorControlClusterEnhancedStepHueParams) StepSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepSize"))
	return rv
}/* debug [instance_properties/getter]: stepSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams/stepSize
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetStepSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepSize:"), value)
}/* debug [instance_properties/setter]: stepSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterEnhancedStepHueParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams/transitionTime
func (m_ MTRColorControlClusterEnhancedStepHueParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams/transitionTime
func (m_ MTRColorControlClusterEnhancedStepHueParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterEnhancedStepHueParams */



