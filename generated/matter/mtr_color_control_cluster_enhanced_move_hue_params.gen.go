// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterEnhancedMoveHueParams */


/* debug [class_header]: Header for MTRColorControlClusterEnhancedMoveHueParams */
// The class instance for the [MTRColorControlClusterEnhancedMoveHueParams] class.
var (
	MTRColorControlClusterEnhancedMoveHueParamsClass     _MTRColorControlClusterEnhancedMoveHueParamsClass
	MTRColorControlClusterEnhancedMoveHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterEnhancedMoveHueParamsClass() _MTRColorControlClusterEnhancedMoveHueParamsClass {
	MTRColorControlClusterEnhancedMoveHueParamsClassOnce.Do(func() {
		MTRColorControlClusterEnhancedMoveHueParamsClass = _MTRColorControlClusterEnhancedMoveHueParamsClass{objc.GetClass("MTRColorControlClusterEnhancedMoveHueParams")}
	})
	return MTRColorControlClusterEnhancedMoveHueParamsClass
}

type _MTRColorControlClusterEnhancedMoveHueParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterEnhancedMoveHueParams */
// An interface definition for the [MTRColorControlClusterEnhancedMoveHueParams] class.
type IMTRColorControlClusterEnhancedMoveHueParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterEnhancedMoveHueParams */
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

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterEnhancedMoveHueParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterEnhancedMoveHueParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterEnhancedMoveHueParamsClass) Alloc() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterEnhancedMoveHueParamsClass) New() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterEnhancedMoveHueParams) Init() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterEnhancedMoveHueParams) Autorelease() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterEnhancedMoveHueParams creates a new MTRColorControlClusterEnhancedMoveHueParams instance.
func NewMTRColorControlClusterEnhancedMoveHueParams() MTRColorControlClusterEnhancedMoveHueParams {
	return getMTRColorControlClusterEnhancedMoveHueParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterEnhancedMoveHueParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams
type MTRColorControlClusterEnhancedMoveHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterEnhancedMoveHueParamsFrom constructs a [MTRColorControlClusterEnhancedMoveHueParams] from an unsafe.Pointer.
func MTRColorControlClusterEnhancedMoveHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterEnhancedMoveHueParams {
	return MTRColorControlClusterEnhancedMoveHueParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterEnhancedMoveHueParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterEnhancedMoveHueParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterEnhancedMoveHueParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterEnhancedMoveHueParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterEnhancedMoveHueParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams/moveMode
func (m_ MTRColorControlClusterEnhancedMoveHueParams) MoveMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("moveMode"))
	return rv
}/* debug [instance_properties/getter]: moveMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams/moveMode
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetMoveMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}/* debug [instance_properties/setter]: moveMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams/optionsMask
func (m_ MTRColorControlClusterEnhancedMoveHueParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams/optionsMask
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams/optionsOverride
func (m_ MTRColorControlClusterEnhancedMoveHueParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams/optionsOverride
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams/rate
func (m_ MTRColorControlClusterEnhancedMoveHueParams) Rate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams/rate
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterEnhancedMoveHueParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterEnhancedMoveHueParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterEnhancedMoveHueParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterEnhancedMoveHueParams */



