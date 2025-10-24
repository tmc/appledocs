// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterMoveColorTemperatureParams */


/* debug [class_header]: Header for MTRColorControlClusterMoveColorTemperatureParams */
// The class instance for the [MTRColorControlClusterMoveColorTemperatureParams] class.
var (
	MTRColorControlClusterMoveColorTemperatureParamsClass     _MTRColorControlClusterMoveColorTemperatureParamsClass
	MTRColorControlClusterMoveColorTemperatureParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveColorTemperatureParamsClass() _MTRColorControlClusterMoveColorTemperatureParamsClass {
	MTRColorControlClusterMoveColorTemperatureParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveColorTemperatureParamsClass = _MTRColorControlClusterMoveColorTemperatureParamsClass{objc.GetClass("MTRColorControlClusterMoveColorTemperatureParams")}
	})
	return MTRColorControlClusterMoveColorTemperatureParamsClass
}

type _MTRColorControlClusterMoveColorTemperatureParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterMoveColorTemperatureParams */
// An interface definition for the [MTRColorControlClusterMoveColorTemperatureParams] class.
type IMTRColorControlClusterMoveColorTemperatureParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterMoveColorTemperatureParams */
	// properties:
	ColorTemperatureMaximumMireds() objc.IObject /* cross-framework: NSNumber */
	SetColorTemperatureMaximumMireds(value objc.IObject /* cross-framework: NSNumber */)
	ColorTemperatureMinimumMireds() objc.IObject /* cross-framework: NSNumber */
	SetColorTemperatureMinimumMireds(value objc.IObject /* cross-framework: NSNumber */)
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

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterMoveColorTemperatureParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterMoveColorTemperatureParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveColorTemperatureParamsClass) Alloc() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterMoveColorTemperatureParamsClass) New() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveColorTemperatureParams) Init() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveColorTemperatureParams) Autorelease() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveColorTemperatureParams creates a new MTRColorControlClusterMoveColorTemperatureParams instance.
func NewMTRColorControlClusterMoveColorTemperatureParams() MTRColorControlClusterMoveColorTemperatureParams {
	return getMTRColorControlClusterMoveColorTemperatureParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterMoveColorTemperatureParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams
type MTRColorControlClusterMoveColorTemperatureParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveColorTemperatureParamsFrom constructs a [MTRColorControlClusterMoveColorTemperatureParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveColorTemperatureParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveColorTemperatureParams {
	return MTRColorControlClusterMoveColorTemperatureParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterMoveColorTemperatureParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterMoveColorTemperatureParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterMoveColorTemperatureParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterMoveColorTemperatureParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterMoveColorTemperatureParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/colorTemperatureMaximumMireds
func (m_ MTRColorControlClusterMoveColorTemperatureParams) ColorTemperatureMaximumMireds() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorTemperatureMaximumMireds"))
	return rv
}/* debug [instance_properties/getter]: colorTemperatureMaximumMireds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/colorTemperatureMaximumMireds
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetColorTemperatureMaximumMireds(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMaximumMireds:"), value)
}/* debug [instance_properties/setter]: colorTemperatureMaximumMireds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/colorTemperatureMinimumMireds
func (m_ MTRColorControlClusterMoveColorTemperatureParams) ColorTemperatureMinimumMireds() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("colorTemperatureMinimumMireds"))
	return rv
}/* debug [instance_properties/getter]: colorTemperatureMinimumMireds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/colorTemperatureMinimumMireds
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetColorTemperatureMinimumMireds(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTemperatureMinimumMireds:"), value)
}/* debug [instance_properties/setter]: colorTemperatureMinimumMireds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/moveMode
func (m_ MTRColorControlClusterMoveColorTemperatureParams) MoveMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("moveMode"))
	return rv
}/* debug [instance_properties/getter]: moveMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/moveMode
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetMoveMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}/* debug [instance_properties/setter]: moveMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/optionsMask
func (m_ MTRColorControlClusterMoveColorTemperatureParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/optionsMask
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/optionsOverride
func (m_ MTRColorControlClusterMoveColorTemperatureParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/optionsOverride
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/rate
func (m_ MTRColorControlClusterMoveColorTemperatureParams) Rate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/rate
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveColorTemperatureParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveColorTemperatureParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveColorTemperatureParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterMoveColorTemperatureParams */



