// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterMoveToHueParams */


/* debug [class_header]: Header for MTRColorControlClusterMoveToHueParams */
// The class instance for the [MTRColorControlClusterMoveToHueParams] class.
var (
	MTRColorControlClusterMoveToHueParamsClass     _MTRColorControlClusterMoveToHueParamsClass
	MTRColorControlClusterMoveToHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToHueParamsClass() _MTRColorControlClusterMoveToHueParamsClass {
	MTRColorControlClusterMoveToHueParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToHueParamsClass = _MTRColorControlClusterMoveToHueParamsClass{objc.GetClass("MTRColorControlClusterMoveToHueParams")}
	})
	return MTRColorControlClusterMoveToHueParamsClass
}

type _MTRColorControlClusterMoveToHueParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterMoveToHueParams */
// An interface definition for the [MTRColorControlClusterMoveToHueParams] class.
type IMTRColorControlClusterMoveToHueParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterMoveToHueParams */
	// properties:
	Direction() objc.IObject /* cross-framework: NSNumber */
	SetDirection(value objc.IObject /* cross-framework: NSNumber */)
	Hue() objc.IObject /* cross-framework: NSNumber */
	SetHue(value objc.IObject /* cross-framework: NSNumber */)
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

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterMoveToHueParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterMoveToHueParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToHueParamsClass) Alloc() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterMoveToHueParamsClass) New() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToHueParams) Init() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToHueParams) Autorelease() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToHueParams creates a new MTRColorControlClusterMoveToHueParams instance.
func NewMTRColorControlClusterMoveToHueParams() MTRColorControlClusterMoveToHueParams {
	return getMTRColorControlClusterMoveToHueParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterMoveToHueParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams
type MTRColorControlClusterMoveToHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToHueParamsFrom constructs a [MTRColorControlClusterMoveToHueParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToHueParams {
	return MTRColorControlClusterMoveToHueParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterMoveToHueParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterMoveToHueParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterMoveToHueParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterMoveToHueParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterMoveToHueParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams/direction
func (m_ MTRColorControlClusterMoveToHueParams) Direction() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("direction"))
	return rv
}/* debug [instance_properties/getter]: direction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams/direction
func (m_ MTRColorControlClusterMoveToHueParams) SetDirection(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDirection:"), value)
}/* debug [instance_properties/setter]: direction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams/hue
func (m_ MTRColorControlClusterMoveToHueParams) Hue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("hue"))
	return rv
}/* debug [instance_properties/getter]: hue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams/hue
func (m_ MTRColorControlClusterMoveToHueParams) SetHue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHue:"), value)
}/* debug [instance_properties/setter]: hue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams/optionsMask
func (m_ MTRColorControlClusterMoveToHueParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams/optionsMask
func (m_ MTRColorControlClusterMoveToHueParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams/optionsOverride
func (m_ MTRColorControlClusterMoveToHueParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams/optionsOverride
func (m_ MTRColorControlClusterMoveToHueParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveToHueParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveToHueParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveToHueParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveToHueParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams/transitionTime
func (m_ MTRColorControlClusterMoveToHueParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams/transitionTime
func (m_ MTRColorControlClusterMoveToHueParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterMoveToHueParams */



