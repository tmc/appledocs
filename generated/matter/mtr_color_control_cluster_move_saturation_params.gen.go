// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterMoveSaturationParams */


/* debug [class_header]: Header for MTRColorControlClusterMoveSaturationParams */
// The class instance for the [MTRColorControlClusterMoveSaturationParams] class.
var (
	MTRColorControlClusterMoveSaturationParamsClass     _MTRColorControlClusterMoveSaturationParamsClass
	MTRColorControlClusterMoveSaturationParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveSaturationParamsClass() _MTRColorControlClusterMoveSaturationParamsClass {
	MTRColorControlClusterMoveSaturationParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveSaturationParamsClass = _MTRColorControlClusterMoveSaturationParamsClass{objc.GetClass("MTRColorControlClusterMoveSaturationParams")}
	})
	return MTRColorControlClusterMoveSaturationParamsClass
}

type _MTRColorControlClusterMoveSaturationParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterMoveSaturationParams */
// An interface definition for the [MTRColorControlClusterMoveSaturationParams] class.
type IMTRColorControlClusterMoveSaturationParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterMoveSaturationParams */
	// properties:
	MoveMode() objc.IObject /* cross-framework: NSNumber */
	SetMoveMode(value objc.IObject /* cross-framework: NSNumber */)
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	Rate() objc.IObject /* cross-framework: NSNumber */
	SetRate(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterMoveSaturationParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterMoveSaturationParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveSaturationParamsClass) Alloc() MTRColorControlClusterMoveSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveSaturationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterMoveSaturationParamsClass) New() MTRColorControlClusterMoveSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveSaturationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveSaturationParams) Init() MTRColorControlClusterMoveSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveSaturationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveSaturationParams) Autorelease() MTRColorControlClusterMoveSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveSaturationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveSaturationParams creates a new MTRColorControlClusterMoveSaturationParams instance.
func NewMTRColorControlClusterMoveSaturationParams() MTRColorControlClusterMoveSaturationParams {
	return getMTRColorControlClusterMoveSaturationParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterMoveSaturationParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams
type MTRColorControlClusterMoveSaturationParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveSaturationParamsFrom constructs a [MTRColorControlClusterMoveSaturationParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveSaturationParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveSaturationParams {
	return MTRColorControlClusterMoveSaturationParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterMoveSaturationParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterMoveSaturationParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterMoveSaturationParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterMoveSaturationParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterMoveSaturationParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams/moveMode
func (m_ MTRColorControlClusterMoveSaturationParams) MoveMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("moveMode"))
	return rv
}/* debug [instance_properties/getter]: moveMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams/moveMode
func (m_ MTRColorControlClusterMoveSaturationParams) SetMoveMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}/* debug [instance_properties/setter]: moveMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams/optionsMask
func (m_ MTRColorControlClusterMoveSaturationParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams/optionsMask
func (m_ MTRColorControlClusterMoveSaturationParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams/optionsOverride
func (m_ MTRColorControlClusterMoveSaturationParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams/optionsOverride
func (m_ MTRColorControlClusterMoveSaturationParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams/rate
func (m_ MTRColorControlClusterMoveSaturationParams) Rate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams/rate
func (m_ MTRColorControlClusterMoveSaturationParams) SetRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveSaturationParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveSaturationParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveSaturationParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveSaturationParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterMoveSaturationParams */



