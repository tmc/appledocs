// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRLevelControlClusterStepParams */


/* debug [class_header]: Header for MTRLevelControlClusterStepParams */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRLevelControlClusterStepParams */
// An interface definition for the [MTRLevelControlClusterStepParams] class.
type IMTRLevelControlClusterStepParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRLevelControlClusterStepParams */
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

	
/* debug [class_interface_methods]: Methods for MTRLevelControlClusterStepParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRLevelControlClusterStepParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterStepParamsClass) Alloc() MTRLevelControlClusterStepParams {
	rv := objc.Send[MTRLevelControlClusterStepParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRLevelControlClusterStepParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams
type MTRLevelControlClusterStepParams struct {
	objectivec.Object
}

// MTRLevelControlClusterStepParamsFrom constructs a [MTRLevelControlClusterStepParams] from an unsafe.Pointer.
func MTRLevelControlClusterStepParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterStepParams {
	return MTRLevelControlClusterStepParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRLevelControlClusterStepParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRLevelControlClusterStepParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRLevelControlClusterStepParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRLevelControlClusterStepParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRLevelControlClusterStepParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams/optionsMask
func (m_ MTRLevelControlClusterStepParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams/optionsMask
func (m_ MTRLevelControlClusterStepParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams/optionsOverride
func (m_ MTRLevelControlClusterStepParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams/optionsOverride
func (m_ MTRLevelControlClusterStepParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterStepParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterStepParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams/stepMode
func (m_ MTRLevelControlClusterStepParams) StepMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepMode"))
	return rv
}/* debug [instance_properties/getter]: stepMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams/stepMode
func (m_ MTRLevelControlClusterStepParams) SetStepMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepMode:"), value)
}/* debug [instance_properties/setter]: stepMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams/stepSize
func (m_ MTRLevelControlClusterStepParams) StepSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepSize"))
	return rv
}/* debug [instance_properties/getter]: stepSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams/stepSize
func (m_ MTRLevelControlClusterStepParams) SetStepSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepSize:"), value)
}/* debug [instance_properties/setter]: stepSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterStepParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterStepParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams/transitionTime
func (m_ MTRLevelControlClusterStepParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepParams/transitionTime
func (m_ MTRLevelControlClusterStepParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRLevelControlClusterStepParams */



