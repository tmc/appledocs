// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterEnhancedMoveToHueParams */


/* debug [class_header]: Header for MTRColorControlClusterEnhancedMoveToHueParams */
// The class instance for the [MTRColorControlClusterEnhancedMoveToHueParams] class.
var (
	MTRColorControlClusterEnhancedMoveToHueParamsClass     _MTRColorControlClusterEnhancedMoveToHueParamsClass
	MTRColorControlClusterEnhancedMoveToHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterEnhancedMoveToHueParamsClass() _MTRColorControlClusterEnhancedMoveToHueParamsClass {
	MTRColorControlClusterEnhancedMoveToHueParamsClassOnce.Do(func() {
		MTRColorControlClusterEnhancedMoveToHueParamsClass = _MTRColorControlClusterEnhancedMoveToHueParamsClass{objc.GetClass("MTRColorControlClusterEnhancedMoveToHueParams")}
	})
	return MTRColorControlClusterEnhancedMoveToHueParamsClass
}

type _MTRColorControlClusterEnhancedMoveToHueParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterEnhancedMoveToHueParams */
// An interface definition for the [MTRColorControlClusterEnhancedMoveToHueParams] class.
type IMTRColorControlClusterEnhancedMoveToHueParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterEnhancedMoveToHueParams */
	// properties:
	Direction() objc.IObject /* cross-framework: NSNumber */
	SetDirection(value objc.IObject /* cross-framework: NSNumber */)
	EnhancedHue() objc.IObject /* cross-framework: NSNumber */
	SetEnhancedHue(value objc.IObject /* cross-framework: NSNumber */)
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TransitionTime() objc.IObject /* cross-framework: NSNumber */
	SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterEnhancedMoveToHueParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterEnhancedMoveToHueParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterEnhancedMoveToHueParamsClass) Alloc() MTRColorControlClusterEnhancedMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterEnhancedMoveToHueParamsClass) New() MTRColorControlClusterEnhancedMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) Init() MTRColorControlClusterEnhancedMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) Autorelease() MTRColorControlClusterEnhancedMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterEnhancedMoveToHueParams creates a new MTRColorControlClusterEnhancedMoveToHueParams instance.
func NewMTRColorControlClusterEnhancedMoveToHueParams() MTRColorControlClusterEnhancedMoveToHueParams {
	return getMTRColorControlClusterEnhancedMoveToHueParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterEnhancedMoveToHueParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams
type MTRColorControlClusterEnhancedMoveToHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterEnhancedMoveToHueParamsFrom constructs a [MTRColorControlClusterEnhancedMoveToHueParams] from an unsafe.Pointer.
func MTRColorControlClusterEnhancedMoveToHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterEnhancedMoveToHueParams {
	return MTRColorControlClusterEnhancedMoveToHueParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterEnhancedMoveToHueParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterEnhancedMoveToHueParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterEnhancedMoveToHueParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterEnhancedMoveToHueParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterEnhancedMoveToHueParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams/direction
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) Direction() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("direction"))
	return rv
}/* debug [instance_properties/getter]: direction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams/direction
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) SetDirection(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDirection:"), value)
}/* debug [instance_properties/setter]: direction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams/enhancedHue
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) EnhancedHue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("enhancedHue"))
	return rv
}/* debug [instance_properties/getter]: enhancedHue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams/enhancedHue
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) SetEnhancedHue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnhancedHue:"), value)
}/* debug [instance_properties/setter]: enhancedHue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams/optionsMask
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams/optionsMask
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams/optionsOverride
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams/optionsOverride
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams/transitionTime
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams/transitionTime
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterEnhancedMoveToHueParams */



