// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterMoveToHueAndSaturationParams */


/* debug [class_header]: Header for MTRColorControlClusterMoveToHueAndSaturationParams */
// The class instance for the [MTRColorControlClusterMoveToHueAndSaturationParams] class.
var (
	MTRColorControlClusterMoveToHueAndSaturationParamsClass     _MTRColorControlClusterMoveToHueAndSaturationParamsClass
	MTRColorControlClusterMoveToHueAndSaturationParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToHueAndSaturationParamsClass() _MTRColorControlClusterMoveToHueAndSaturationParamsClass {
	MTRColorControlClusterMoveToHueAndSaturationParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToHueAndSaturationParamsClass = _MTRColorControlClusterMoveToHueAndSaturationParamsClass{objc.GetClass("MTRColorControlClusterMoveToHueAndSaturationParams")}
	})
	return MTRColorControlClusterMoveToHueAndSaturationParamsClass
}

type _MTRColorControlClusterMoveToHueAndSaturationParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterMoveToHueAndSaturationParams */
// An interface definition for the [MTRColorControlClusterMoveToHueAndSaturationParams] class.
type IMTRColorControlClusterMoveToHueAndSaturationParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterMoveToHueAndSaturationParams */
	// properties:
	Hue() objc.IObject /* cross-framework: NSNumber */
	SetHue(value objc.IObject /* cross-framework: NSNumber */)
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

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterMoveToHueAndSaturationParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterMoveToHueAndSaturationParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToHueAndSaturationParamsClass) Alloc() MTRColorControlClusterMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueAndSaturationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterMoveToHueAndSaturationParamsClass) New() MTRColorControlClusterMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueAndSaturationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) Init() MTRColorControlClusterMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueAndSaturationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) Autorelease() MTRColorControlClusterMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueAndSaturationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToHueAndSaturationParams creates a new MTRColorControlClusterMoveToHueAndSaturationParams instance.
func NewMTRColorControlClusterMoveToHueAndSaturationParams() MTRColorControlClusterMoveToHueAndSaturationParams {
	return getMTRColorControlClusterMoveToHueAndSaturationParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterMoveToHueAndSaturationParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams
type MTRColorControlClusterMoveToHueAndSaturationParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToHueAndSaturationParamsFrom constructs a [MTRColorControlClusterMoveToHueAndSaturationParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToHueAndSaturationParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToHueAndSaturationParams {
	return MTRColorControlClusterMoveToHueAndSaturationParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterMoveToHueAndSaturationParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterMoveToHueAndSaturationParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterMoveToHueAndSaturationParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterMoveToHueAndSaturationParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterMoveToHueAndSaturationParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams/hue
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) Hue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("hue"))
	return rv
}/* debug [instance_properties/getter]: hue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams/hue
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) SetHue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHue:"), value)
}/* debug [instance_properties/setter]: hue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams/optionsMask
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams/optionsMask
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams/optionsOverride
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams/optionsOverride
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams/saturation
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) Saturation() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("saturation"))
	return rv
}/* debug [instance_properties/getter]: saturation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams/saturation
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) SetSaturation(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSaturation:"), value)
}/* debug [instance_properties/setter]: saturation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams/transitionTime
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueAndSaturationParams/transitionTime
func (m_ MTRColorControlClusterMoveToHueAndSaturationParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterMoveToHueAndSaturationParams */



