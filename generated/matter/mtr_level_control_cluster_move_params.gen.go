// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRLevelControlClusterMoveParams */


/* debug [class_header]: Header for MTRLevelControlClusterMoveParams */
// The class instance for the [MTRLevelControlClusterMoveParams] class.
var (
	MTRLevelControlClusterMoveParamsClass     _MTRLevelControlClusterMoveParamsClass
	MTRLevelControlClusterMoveParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveParamsClass() _MTRLevelControlClusterMoveParamsClass {
	MTRLevelControlClusterMoveParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveParamsClass = _MTRLevelControlClusterMoveParamsClass{objc.GetClass("MTRLevelControlClusterMoveParams")}
	})
	return MTRLevelControlClusterMoveParamsClass
}

type _MTRLevelControlClusterMoveParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRLevelControlClusterMoveParams */
// An interface definition for the [MTRLevelControlClusterMoveParams] class.
type IMTRLevelControlClusterMoveParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRLevelControlClusterMoveParams */
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

	
/* debug [class_interface_methods]: Methods for MTRLevelControlClusterMoveParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRLevelControlClusterMoveParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveParamsClass) Alloc() MTRLevelControlClusterMoveParams {
	rv := objc.Send[MTRLevelControlClusterMoveParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRLevelControlClusterMoveParamsClass) New() MTRLevelControlClusterMoveParams {
	rv := objc.Send[MTRLevelControlClusterMoveParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveParams) Init() MTRLevelControlClusterMoveParams {
	rv := objc.Send[MTRLevelControlClusterMoveParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveParams) Autorelease() MTRLevelControlClusterMoveParams {
	rv := objc.Send[MTRLevelControlClusterMoveParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveParams creates a new MTRLevelControlClusterMoveParams instance.
func NewMTRLevelControlClusterMoveParams() MTRLevelControlClusterMoveParams {
	return getMTRLevelControlClusterMoveParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRLevelControlClusterMoveParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams
type MTRLevelControlClusterMoveParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveParamsFrom constructs a [MTRLevelControlClusterMoveParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveParams {
	return MTRLevelControlClusterMoveParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRLevelControlClusterMoveParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRLevelControlClusterMoveParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRLevelControlClusterMoveParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRLevelControlClusterMoveParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRLevelControlClusterMoveParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams/moveMode
func (m_ MTRLevelControlClusterMoveParams) MoveMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("moveMode"))
	return rv
}/* debug [instance_properties/getter]: moveMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams/moveMode
func (m_ MTRLevelControlClusterMoveParams) SetMoveMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}/* debug [instance_properties/setter]: moveMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams/optionsMask
func (m_ MTRLevelControlClusterMoveParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams/optionsMask
func (m_ MTRLevelControlClusterMoveParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams/optionsOverride
func (m_ MTRLevelControlClusterMoveParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams/optionsOverride
func (m_ MTRLevelControlClusterMoveParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams/rate
func (m_ MTRLevelControlClusterMoveParams) Rate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams/rate
func (m_ MTRLevelControlClusterMoveParams) SetRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterMoveParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterMoveParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterMoveParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterMoveParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRLevelControlClusterMoveParams */



