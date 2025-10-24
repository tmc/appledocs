// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterStepColorTemperatureParams */


/* debug [class_header]: Header for MTRColorControlClusterStepColorTemperatureParams */
// The class instance for the [MTRColorControlClusterStepColorTemperatureParams] class.
var (
	MTRColorControlClusterStepColorTemperatureParamsClass     _MTRColorControlClusterStepColorTemperatureParamsClass
	MTRColorControlClusterStepColorTemperatureParamsClassOnce sync.Once
)

func getMTRColorControlClusterStepColorTemperatureParamsClass() _MTRColorControlClusterStepColorTemperatureParamsClass {
	MTRColorControlClusterStepColorTemperatureParamsClassOnce.Do(func() {
		MTRColorControlClusterStepColorTemperatureParamsClass = _MTRColorControlClusterStepColorTemperatureParamsClass{objc.GetClass("MTRColorControlClusterStepColorTemperatureParams")}
	})
	return MTRColorControlClusterStepColorTemperatureParamsClass
}

type _MTRColorControlClusterStepColorTemperatureParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterStepColorTemperatureParams */
// An interface definition for the [MTRColorControlClusterStepColorTemperatureParams] class.
type IMTRColorControlClusterStepColorTemperatureParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterStepColorTemperatureParams */
	// properties:
	ColorTemperatureMaximumMireds() objc.IObject /* cross-framework: NSNumber */
	SetColorTemperatureMaximumMireds(value objc.IObject /* cross-framework: NSNumber */)
	ColorTemperatureMinimumMireds() objc.IObject /* cross-framework: NSNumber */
	SetColorTemperatureMinimumMireds(value objc.IObject /* cross-framework: NSNumber */)
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

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterStepColorTemperatureParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterStepColorTemperatureParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterStepColorTemperatureParamsClass) Alloc() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterStepColorTemperatureParamsClass) New() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterStepColorTemperatureParams) Init() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterStepColorTemperatureParams) Autorelease() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterStepColorTemperatureParams creates a new MTRColorControlClusterStepColorTemperatureParams instance.
func NewMTRColorControlClusterStepColorTemperatureParams() MTRColorControlClusterStepColorTemperatureParams {
	return getMTRColorControlClusterStepColorTemperatureParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterStepColorTemperatureParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams
type MTRColorControlClusterStepColorTemperatureParams struct {
	objectivec.Object
}

// MTRColorControlClusterStepColorTemperatureParamsFrom constructs a [MTRColorControlClusterStepColorTemperatureParams] from an unsafe.Pointer.
func MTRColorControlClusterStepColorTemperatureParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterStepColorTemperatureParams {
	return MTRColorControlClusterStepColorTemperatureParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterStepColorTemperatureParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterStepColorTemperatureParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterStepColorTemperatureParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterStepColorTemperatureParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterStepColorTemperatureParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/colorTemperatureMaximumMireds
func (m_ MTRColorControlClusterStepColorTemperatureParams) ColorTemperatureMaximumMireds() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorTemperatureMaximumMireds"))
	return rv
}/* debug [instance_properties/getter]: colorTemperatureMaximumMireds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/colorTemperatureMaximumMireds
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetColorTemperatureMaximumMireds(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMaximumMireds:"), value)
}/* debug [instance_properties/setter]: colorTemperatureMaximumMireds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/colorTemperatureMinimumMireds
func (m_ MTRColorControlClusterStepColorTemperatureParams) ColorTemperatureMinimumMireds() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorTemperatureMinimumMireds"))
	return rv
}/* debug [instance_properties/getter]: colorTemperatureMinimumMireds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/colorTemperatureMinimumMireds
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetColorTemperatureMinimumMireds(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMinimumMireds:"), value)
}/* debug [instance_properties/setter]: colorTemperatureMinimumMireds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/optionsMask
func (m_ MTRColorControlClusterStepColorTemperatureParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/optionsMask
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/optionsOverride
func (m_ MTRColorControlClusterStepColorTemperatureParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/optionsOverride
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterStepColorTemperatureParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/stepMode
func (m_ MTRColorControlClusterStepColorTemperatureParams) StepMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepMode"))
	return rv
}/* debug [instance_properties/getter]: stepMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/stepMode
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetStepMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepMode:"), value)
}/* debug [instance_properties/setter]: stepMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/stepSize
func (m_ MTRColorControlClusterStepColorTemperatureParams) StepSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepSize"))
	return rv
}/* debug [instance_properties/getter]: stepSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/stepSize
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetStepSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepSize:"), value)
}/* debug [instance_properties/setter]: stepSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterStepColorTemperatureParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/transitionTime
func (m_ MTRColorControlClusterStepColorTemperatureParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams/transitionTime
func (m_ MTRColorControlClusterStepColorTemperatureParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterStepColorTemperatureParams */



