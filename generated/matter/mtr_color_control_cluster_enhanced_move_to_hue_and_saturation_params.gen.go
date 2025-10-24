// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterEnhancedMoveToHueAndSaturationParams */


/* debug [class_header]: Header for MTRColorControlClusterEnhancedMoveToHueAndSaturationParams */
// The class instance for the [MTRColorControlClusterEnhancedMoveToHueAndSaturationParams] class.
var (
	MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass     _MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass
	MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClassOnce sync.Once
)

func getMTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass() _MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass {
	MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClassOnce.Do(func() {
		MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass = _MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass{objc.GetClass("MTRColorControlClusterEnhancedMoveToHueAndSaturationParams")}
	})
	return MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass
}

type _MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterEnhancedMoveToHueAndSaturationParams */
// An interface definition for the [MTRColorControlClusterEnhancedMoveToHueAndSaturationParams] class.
type IMTRColorControlClusterEnhancedMoveToHueAndSaturationParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterEnhancedMoveToHueAndSaturationParams */
	// properties:
	EnhancedHue() objc.IObject /* cross-framework: NSNumber */
	SetEnhancedHue(value objc.IObject /* cross-framework: NSNumber */)
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

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterEnhancedMoveToHueAndSaturationParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterEnhancedMoveToHueAndSaturationParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass) Alloc() MTRColorControlClusterEnhancedMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueAndSaturationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass) New() MTRColorControlClusterEnhancedMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueAndSaturationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) Init() MTRColorControlClusterEnhancedMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueAndSaturationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) Autorelease() MTRColorControlClusterEnhancedMoveToHueAndSaturationParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueAndSaturationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterEnhancedMoveToHueAndSaturationParams creates a new MTRColorControlClusterEnhancedMoveToHueAndSaturationParams instance.
func NewMTRColorControlClusterEnhancedMoveToHueAndSaturationParams() MTRColorControlClusterEnhancedMoveToHueAndSaturationParams {
	return getMTRColorControlClusterEnhancedMoveToHueAndSaturationParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterEnhancedMoveToHueAndSaturationParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams
type MTRColorControlClusterEnhancedMoveToHueAndSaturationParams struct {
	objectivec.Object
}

// MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsFrom constructs a [MTRColorControlClusterEnhancedMoveToHueAndSaturationParams] from an unsafe.Pointer.
func MTRColorControlClusterEnhancedMoveToHueAndSaturationParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterEnhancedMoveToHueAndSaturationParams {
	return MTRColorControlClusterEnhancedMoveToHueAndSaturationParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterEnhancedMoveToHueAndSaturationParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterEnhancedMoveToHueAndSaturationParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterEnhancedMoveToHueAndSaturationParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterEnhancedMoveToHueAndSaturationParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterEnhancedMoveToHueAndSaturationParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams/enhancedHue
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) EnhancedHue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("enhancedHue"))
	return rv
}/* debug [instance_properties/getter]: enhancedHue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams/enhancedHue
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) SetEnhancedHue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnhancedHue:"), value)
}/* debug [instance_properties/setter]: enhancedHue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams/optionsMask
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams/optionsMask
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams/optionsOverride
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams/optionsOverride
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams/saturation
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) Saturation() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("saturation"))
	return rv
}/* debug [instance_properties/getter]: saturation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams/saturation
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) SetSaturation(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSaturation:"), value)
}/* debug [instance_properties/setter]: saturation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams/transitionTime
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueAndSaturationParams/transitionTime
func (m_ MTRColorControlClusterEnhancedMoveToHueAndSaturationParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterEnhancedMoveToHueAndSaturationParams */



