// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterMoveToSaturationParams */


/* debug [class_header]: Header for MTRColorControlClusterMoveToSaturationParams */
// The class instance for the [MTRColorControlClusterMoveToSaturationParams] class.
var (
	MTRColorControlClusterMoveToSaturationParamsClass     _MTRColorControlClusterMoveToSaturationParamsClass
	MTRColorControlClusterMoveToSaturationParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToSaturationParamsClass() _MTRColorControlClusterMoveToSaturationParamsClass {
	MTRColorControlClusterMoveToSaturationParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToSaturationParamsClass = _MTRColorControlClusterMoveToSaturationParamsClass{objc.GetClass("MTRColorControlClusterMoveToSaturationParams")}
	})
	return MTRColorControlClusterMoveToSaturationParamsClass
}

type _MTRColorControlClusterMoveToSaturationParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterMoveToSaturationParams */
// An interface definition for the [MTRColorControlClusterMoveToSaturationParams] class.
type IMTRColorControlClusterMoveToSaturationParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterMoveToSaturationParams */
	// properties:
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	Saturation() objc.IObject /* cross-framework: NSNumber */
	SetSaturation(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TransitionTime() objc.IObject /* cross-framework: NSNumber */
	SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterMoveToSaturationParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterMoveToSaturationParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToSaturationParamsClass) Alloc() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterMoveToSaturationParamsClass) New() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToSaturationParams) Init() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToSaturationParams) Autorelease() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToSaturationParams creates a new MTRColorControlClusterMoveToSaturationParams instance.
func NewMTRColorControlClusterMoveToSaturationParams() MTRColorControlClusterMoveToSaturationParams {
	return getMTRColorControlClusterMoveToSaturationParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterMoveToSaturationParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams
type MTRColorControlClusterMoveToSaturationParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToSaturationParamsFrom constructs a [MTRColorControlClusterMoveToSaturationParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToSaturationParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToSaturationParams {
	return MTRColorControlClusterMoveToSaturationParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterMoveToSaturationParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterMoveToSaturationParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterMoveToSaturationParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterMoveToSaturationParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterMoveToSaturationParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams/optionsMask
func (m_ MTRColorControlClusterMoveToSaturationParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams/optionsMask
func (m_ MTRColorControlClusterMoveToSaturationParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams/optionsOverride
func (m_ MTRColorControlClusterMoveToSaturationParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams/optionsOverride
func (m_ MTRColorControlClusterMoveToSaturationParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams/saturation
func (m_ MTRColorControlClusterMoveToSaturationParams) Saturation() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("saturation"))
	return rv
}/* debug [instance_properties/getter]: saturation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams/saturation
func (m_ MTRColorControlClusterMoveToSaturationParams) SetSaturation(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSaturation:"), value)
}/* debug [instance_properties/setter]: saturation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveToSaturationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveToSaturationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveToSaturationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveToSaturationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams/transitionTime
func (m_ MTRColorControlClusterMoveToSaturationParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams/transitionTime
func (m_ MTRColorControlClusterMoveToSaturationParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterMoveToSaturationParams */



