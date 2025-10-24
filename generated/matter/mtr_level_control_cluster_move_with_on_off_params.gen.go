// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRLevelControlClusterMoveWithOnOffParams */


/* debug [class_header]: Header for MTRLevelControlClusterMoveWithOnOffParams */
// The class instance for the [MTRLevelControlClusterMoveWithOnOffParams] class.
var (
	MTRLevelControlClusterMoveWithOnOffParamsClass     _MTRLevelControlClusterMoveWithOnOffParamsClass
	MTRLevelControlClusterMoveWithOnOffParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveWithOnOffParamsClass() _MTRLevelControlClusterMoveWithOnOffParamsClass {
	MTRLevelControlClusterMoveWithOnOffParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveWithOnOffParamsClass = _MTRLevelControlClusterMoveWithOnOffParamsClass{objc.GetClass("MTRLevelControlClusterMoveWithOnOffParams")}
	})
	return MTRLevelControlClusterMoveWithOnOffParamsClass
}

type _MTRLevelControlClusterMoveWithOnOffParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRLevelControlClusterMoveWithOnOffParams */
// An interface definition for the [MTRLevelControlClusterMoveWithOnOffParams] class.
type IMTRLevelControlClusterMoveWithOnOffParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRLevelControlClusterMoveWithOnOffParams */
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

	
/* debug [class_interface_methods]: Methods for MTRLevelControlClusterMoveWithOnOffParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRLevelControlClusterMoveWithOnOffParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveWithOnOffParamsClass) Alloc() MTRLevelControlClusterMoveWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveWithOnOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRLevelControlClusterMoveWithOnOffParamsClass) New() MTRLevelControlClusterMoveWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveWithOnOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveWithOnOffParams) Init() MTRLevelControlClusterMoveWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveWithOnOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveWithOnOffParams) Autorelease() MTRLevelControlClusterMoveWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveWithOnOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveWithOnOffParams creates a new MTRLevelControlClusterMoveWithOnOffParams instance.
func NewMTRLevelControlClusterMoveWithOnOffParams() MTRLevelControlClusterMoveWithOnOffParams {
	return getMTRLevelControlClusterMoveWithOnOffParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRLevelControlClusterMoveWithOnOffParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams
type MTRLevelControlClusterMoveWithOnOffParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveWithOnOffParamsFrom constructs a [MTRLevelControlClusterMoveWithOnOffParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveWithOnOffParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveWithOnOffParams {
	return MTRLevelControlClusterMoveWithOnOffParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRLevelControlClusterMoveWithOnOffParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRLevelControlClusterMoveWithOnOffParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRLevelControlClusterMoveWithOnOffParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRLevelControlClusterMoveWithOnOffParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRLevelControlClusterMoveWithOnOffParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams/moveMode
func (m_ MTRLevelControlClusterMoveWithOnOffParams) MoveMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("moveMode"))
	return rv
}/* debug [instance_properties/getter]: moveMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams/moveMode
func (m_ MTRLevelControlClusterMoveWithOnOffParams) SetMoveMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}/* debug [instance_properties/setter]: moveMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams/optionsMask
func (m_ MTRLevelControlClusterMoveWithOnOffParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams/optionsMask
func (m_ MTRLevelControlClusterMoveWithOnOffParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams/optionsOverride
func (m_ MTRLevelControlClusterMoveWithOnOffParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams/optionsOverride
func (m_ MTRLevelControlClusterMoveWithOnOffParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams/rate
func (m_ MTRLevelControlClusterMoveWithOnOffParams) Rate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams/rate
func (m_ MTRLevelControlClusterMoveWithOnOffParams) SetRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterMoveWithOnOffParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterMoveWithOnOffParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterMoveWithOnOffParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterMoveWithOnOffParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRLevelControlClusterMoveWithOnOffParams */



