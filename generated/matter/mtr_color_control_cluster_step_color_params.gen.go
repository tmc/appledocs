// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterStepColorParams */


/* debug [class_header]: Header for MTRColorControlClusterStepColorParams */
// The class instance for the [MTRColorControlClusterStepColorParams] class.
var (
	MTRColorControlClusterStepColorParamsClass     _MTRColorControlClusterStepColorParamsClass
	MTRColorControlClusterStepColorParamsClassOnce sync.Once
)

func getMTRColorControlClusterStepColorParamsClass() _MTRColorControlClusterStepColorParamsClass {
	MTRColorControlClusterStepColorParamsClassOnce.Do(func() {
		MTRColorControlClusterStepColorParamsClass = _MTRColorControlClusterStepColorParamsClass{objc.GetClass("MTRColorControlClusterStepColorParams")}
	})
	return MTRColorControlClusterStepColorParamsClass
}

type _MTRColorControlClusterStepColorParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterStepColorParams */
// An interface definition for the [MTRColorControlClusterStepColorParams] class.
type IMTRColorControlClusterStepColorParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterStepColorParams */
	// properties:
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StepX() objc.IObject /* cross-framework: NSNumber */
	SetStepX(value objc.IObject /* cross-framework: NSNumber */)
	StepY() objc.IObject /* cross-framework: NSNumber */
	SetStepY(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TransitionTime() objc.IObject /* cross-framework: NSNumber */
	SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterStepColorParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterStepColorParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterStepColorParamsClass) Alloc() MTRColorControlClusterStepColorParams {
	rv := objc.Send[MTRColorControlClusterStepColorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterStepColorParamsClass) New() MTRColorControlClusterStepColorParams {
	rv := objc.Send[MTRColorControlClusterStepColorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterStepColorParams) Init() MTRColorControlClusterStepColorParams {
	rv := objc.Send[MTRColorControlClusterStepColorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterStepColorParams) Autorelease() MTRColorControlClusterStepColorParams {
	rv := objc.Send[MTRColorControlClusterStepColorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterStepColorParams creates a new MTRColorControlClusterStepColorParams instance.
func NewMTRColorControlClusterStepColorParams() MTRColorControlClusterStepColorParams {
	return getMTRColorControlClusterStepColorParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterStepColorParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams
type MTRColorControlClusterStepColorParams struct {
	objectivec.Object
}

// MTRColorControlClusterStepColorParamsFrom constructs a [MTRColorControlClusterStepColorParams] from an unsafe.Pointer.
func MTRColorControlClusterStepColorParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterStepColorParams {
	return MTRColorControlClusterStepColorParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterStepColorParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterStepColorParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterStepColorParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterStepColorParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterStepColorParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams/optionsMask
func (m_ MTRColorControlClusterStepColorParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams/optionsMask
func (m_ MTRColorControlClusterStepColorParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams/optionsOverride
func (m_ MTRColorControlClusterStepColorParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams/optionsOverride
func (m_ MTRColorControlClusterStepColorParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterStepColorParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterStepColorParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams/stepX
func (m_ MTRColorControlClusterStepColorParams) StepX() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepX"))
	return rv
}/* debug [instance_properties/getter]: stepX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams/stepX
func (m_ MTRColorControlClusterStepColorParams) SetStepX(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepX:"), value)
}/* debug [instance_properties/setter]: stepX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams/stepY
func (m_ MTRColorControlClusterStepColorParams) StepY() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stepY"))
	return rv
}/* debug [instance_properties/getter]: stepY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams/stepY
func (m_ MTRColorControlClusterStepColorParams) SetStepY(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepY:"), value)
}/* debug [instance_properties/setter]: stepY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterStepColorParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterStepColorParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams/transitionTime
func (m_ MTRColorControlClusterStepColorParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorParams/transitionTime
func (m_ MTRColorControlClusterStepColorParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterStepColorParams */



