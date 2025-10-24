// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterMoveToColorTemperatureParams */


/* debug [class_header]: Header for MTRColorControlClusterMoveToColorTemperatureParams */
// The class instance for the [MTRColorControlClusterMoveToColorTemperatureParams] class.
var (
	MTRColorControlClusterMoveToColorTemperatureParamsClass     _MTRColorControlClusterMoveToColorTemperatureParamsClass
	MTRColorControlClusterMoveToColorTemperatureParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToColorTemperatureParamsClass() _MTRColorControlClusterMoveToColorTemperatureParamsClass {
	MTRColorControlClusterMoveToColorTemperatureParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToColorTemperatureParamsClass = _MTRColorControlClusterMoveToColorTemperatureParamsClass{objc.GetClass("MTRColorControlClusterMoveToColorTemperatureParams")}
	})
	return MTRColorControlClusterMoveToColorTemperatureParamsClass
}

type _MTRColorControlClusterMoveToColorTemperatureParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterMoveToColorTemperatureParams */
// An interface definition for the [MTRColorControlClusterMoveToColorTemperatureParams] class.
type IMTRColorControlClusterMoveToColorTemperatureParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterMoveToColorTemperatureParams */
	// properties:
	ColorTemperature() objc.IObject /* cross-framework: NSNumber */
	SetColorTemperature(value objc.IObject /* cross-framework: NSNumber */)
	ColorTemperatureMireds() objc.IObject /* cross-framework: NSNumber */
	SetColorTemperatureMireds(value objc.IObject /* cross-framework: NSNumber */)
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

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterMoveToColorTemperatureParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterMoveToColorTemperatureParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToColorTemperatureParamsClass) Alloc() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterMoveToColorTemperatureParamsClass) New() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) Init() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) Autorelease() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToColorTemperatureParams creates a new MTRColorControlClusterMoveToColorTemperatureParams instance.
func NewMTRColorControlClusterMoveToColorTemperatureParams() MTRColorControlClusterMoveToColorTemperatureParams {
	return getMTRColorControlClusterMoveToColorTemperatureParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterMoveToColorTemperatureParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams
type MTRColorControlClusterMoveToColorTemperatureParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToColorTemperatureParamsFrom constructs a [MTRColorControlClusterMoveToColorTemperatureParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToColorTemperatureParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToColorTemperatureParams {
	return MTRColorControlClusterMoveToColorTemperatureParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterMoveToColorTemperatureParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterMoveToColorTemperatureParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterMoveToColorTemperatureParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterMoveToColorTemperatureParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterMoveToColorTemperatureParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams/colorTemperature
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) ColorTemperature() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorTemperature"))
	return rv
}/* debug [instance_properties/getter]: colorTemperature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams/colorTemperature
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetColorTemperature(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperature:"), value)
}/* debug [instance_properties/setter]: colorTemperature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams/colorTemperatureMireds
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) ColorTemperatureMireds() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorTemperatureMireds"))
	return rv
}/* debug [instance_properties/getter]: colorTemperatureMireds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams/colorTemperatureMireds
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetColorTemperatureMireds(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMireds:"), value)
}/* debug [instance_properties/setter]: colorTemperatureMireds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams/optionsMask
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams/optionsMask
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams/optionsOverride
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams/optionsOverride
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams/transitionTime
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams/transitionTime
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterMoveToColorTemperatureParams */



