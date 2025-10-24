// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterMoveToColorParams */


/* debug [class_header]: Header for MTRColorControlClusterMoveToColorParams */
// The class instance for the [MTRColorControlClusterMoveToColorParams] class.
var (
	MTRColorControlClusterMoveToColorParamsClass     _MTRColorControlClusterMoveToColorParamsClass
	MTRColorControlClusterMoveToColorParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToColorParamsClass() _MTRColorControlClusterMoveToColorParamsClass {
	MTRColorControlClusterMoveToColorParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToColorParamsClass = _MTRColorControlClusterMoveToColorParamsClass{objc.GetClass("MTRColorControlClusterMoveToColorParams")}
	})
	return MTRColorControlClusterMoveToColorParamsClass
}

type _MTRColorControlClusterMoveToColorParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterMoveToColorParams */
// An interface definition for the [MTRColorControlClusterMoveToColorParams] class.
type IMTRColorControlClusterMoveToColorParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterMoveToColorParams */
	// properties:
	ColorX() objc.IObject /* cross-framework: NSNumber */
	SetColorX(value objc.IObject /* cross-framework: NSNumber */)
	ColorY() objc.IObject /* cross-framework: NSNumber */
	SetColorY(value objc.IObject /* cross-framework: NSNumber */)
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

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterMoveToColorParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterMoveToColorParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToColorParamsClass) Alloc() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterMoveToColorParamsClass) New() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToColorParams) Init() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToColorParams) Autorelease() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToColorParams creates a new MTRColorControlClusterMoveToColorParams instance.
func NewMTRColorControlClusterMoveToColorParams() MTRColorControlClusterMoveToColorParams {
	return getMTRColorControlClusterMoveToColorParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterMoveToColorParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams
type MTRColorControlClusterMoveToColorParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToColorParamsFrom constructs a [MTRColorControlClusterMoveToColorParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToColorParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToColorParams {
	return MTRColorControlClusterMoveToColorParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterMoveToColorParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterMoveToColorParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterMoveToColorParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterMoveToColorParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterMoveToColorParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams/colorX
func (m_ MTRColorControlClusterMoveToColorParams) ColorX() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorX"))
	return rv
}/* debug [instance_properties/getter]: colorX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams/colorX
func (m_ MTRColorControlClusterMoveToColorParams) SetColorX(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorX:"), value)
}/* debug [instance_properties/setter]: colorX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams/colorY
func (m_ MTRColorControlClusterMoveToColorParams) ColorY() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorY"))
	return rv
}/* debug [instance_properties/getter]: colorY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams/colorY
func (m_ MTRColorControlClusterMoveToColorParams) SetColorY(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorY:"), value)
}/* debug [instance_properties/setter]: colorY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams/optionsMask
func (m_ MTRColorControlClusterMoveToColorParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams/optionsMask
func (m_ MTRColorControlClusterMoveToColorParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams/optionsOverride
func (m_ MTRColorControlClusterMoveToColorParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams/optionsOverride
func (m_ MTRColorControlClusterMoveToColorParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveToColorParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveToColorParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveToColorParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveToColorParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams/transitionTime
func (m_ MTRColorControlClusterMoveToColorParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams/transitionTime
func (m_ MTRColorControlClusterMoveToColorParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterMoveToColorParams */



