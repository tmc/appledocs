// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterColorLoopSetParams */


/* debug [class_header]: Header for MTRColorControlClusterColorLoopSetParams */
// The class instance for the [MTRColorControlClusterColorLoopSetParams] class.
var (
	MTRColorControlClusterColorLoopSetParamsClass     _MTRColorControlClusterColorLoopSetParamsClass
	MTRColorControlClusterColorLoopSetParamsClassOnce sync.Once
)

func getMTRColorControlClusterColorLoopSetParamsClass() _MTRColorControlClusterColorLoopSetParamsClass {
	MTRColorControlClusterColorLoopSetParamsClassOnce.Do(func() {
		MTRColorControlClusterColorLoopSetParamsClass = _MTRColorControlClusterColorLoopSetParamsClass{objc.GetClass("MTRColorControlClusterColorLoopSetParams")}
	})
	return MTRColorControlClusterColorLoopSetParamsClass
}

type _MTRColorControlClusterColorLoopSetParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterColorLoopSetParams */
// An interface definition for the [MTRColorControlClusterColorLoopSetParams] class.
type IMTRColorControlClusterColorLoopSetParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterColorLoopSetParams */
	// properties:
	Action() objc.IObject /* cross-framework: NSNumber */
	SetAction(value objc.IObject /* cross-framework: NSNumber */)
	Direction() objc.IObject /* cross-framework: NSNumber */
	SetDirection(value objc.IObject /* cross-framework: NSNumber */)
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StartHue() objc.IObject /* cross-framework: NSNumber */
	SetStartHue(value objc.IObject /* cross-framework: NSNumber */)
	Time() objc.IObject /* cross-framework: NSNumber */
	SetTime(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UpdateFlags() objc.IObject /* cross-framework: NSNumber */
	SetUpdateFlags(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterColorLoopSetParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterColorLoopSetParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterColorLoopSetParamsClass) Alloc() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterColorLoopSetParamsClass) New() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterColorLoopSetParams) Init() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterColorLoopSetParams) Autorelease() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterColorLoopSetParams creates a new MTRColorControlClusterColorLoopSetParams instance.
func NewMTRColorControlClusterColorLoopSetParams() MTRColorControlClusterColorLoopSetParams {
	return getMTRColorControlClusterColorLoopSetParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterColorLoopSetParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams
type MTRColorControlClusterColorLoopSetParams struct {
	objectivec.Object
}

// MTRColorControlClusterColorLoopSetParamsFrom constructs a [MTRColorControlClusterColorLoopSetParams] from an unsafe.Pointer.
func MTRColorControlClusterColorLoopSetParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterColorLoopSetParams {
	return MTRColorControlClusterColorLoopSetParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterColorLoopSetParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterColorLoopSetParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterColorLoopSetParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterColorLoopSetParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterColorLoopSetParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/action
func (m_ MTRColorControlClusterColorLoopSetParams) Action() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/action
func (m_ MTRColorControlClusterColorLoopSetParams) SetAction(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/direction
func (m_ MTRColorControlClusterColorLoopSetParams) Direction() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("direction"))
	return rv
}/* debug [instance_properties/getter]: direction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/direction
func (m_ MTRColorControlClusterColorLoopSetParams) SetDirection(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDirection:"), value)
}/* debug [instance_properties/setter]: direction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/optionsMask
func (m_ MTRColorControlClusterColorLoopSetParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/optionsMask
func (m_ MTRColorControlClusterColorLoopSetParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/optionsOverride
func (m_ MTRColorControlClusterColorLoopSetParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/optionsOverride
func (m_ MTRColorControlClusterColorLoopSetParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterColorLoopSetParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterColorLoopSetParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/startHue
func (m_ MTRColorControlClusterColorLoopSetParams) StartHue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startHue"))
	return rv
}/* debug [instance_properties/getter]: startHue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/startHue
func (m_ MTRColorControlClusterColorLoopSetParams) SetStartHue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartHue:"), value)
}/* debug [instance_properties/setter]: startHue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/time
func (m_ MTRColorControlClusterColorLoopSetParams) Time() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("time"))
	return rv
}/* debug [instance_properties/getter]: time */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/time
func (m_ MTRColorControlClusterColorLoopSetParams) SetTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTime:"), value)
}/* debug [instance_properties/setter]: time */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterColorLoopSetParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterColorLoopSetParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/updateFlags
func (m_ MTRColorControlClusterColorLoopSetParams) UpdateFlags() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("updateFlags"))
	return rv
}/* debug [instance_properties/getter]: updateFlags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams/updateFlags
func (m_ MTRColorControlClusterColorLoopSetParams) SetUpdateFlags(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateFlags:"), value)
}/* debug [instance_properties/setter]: updateFlags */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterColorLoopSetParams */



