// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterMoveHueParams */


/* debug [class_header]: Header for MTRColorControlClusterMoveHueParams */
// The class instance for the [MTRColorControlClusterMoveHueParams] class.
var (
	MTRColorControlClusterMoveHueParamsClass     _MTRColorControlClusterMoveHueParamsClass
	MTRColorControlClusterMoveHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveHueParamsClass() _MTRColorControlClusterMoveHueParamsClass {
	MTRColorControlClusterMoveHueParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveHueParamsClass = _MTRColorControlClusterMoveHueParamsClass{objc.GetClass("MTRColorControlClusterMoveHueParams")}
	})
	return MTRColorControlClusterMoveHueParamsClass
}

type _MTRColorControlClusterMoveHueParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterMoveHueParams */
// An interface definition for the [MTRColorControlClusterMoveHueParams] class.
type IMTRColorControlClusterMoveHueParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterMoveHueParams */
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

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterMoveHueParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterMoveHueParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveHueParamsClass) Alloc() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterMoveHueParamsClass) New() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveHueParams) Init() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveHueParams) Autorelease() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveHueParams creates a new MTRColorControlClusterMoveHueParams instance.
func NewMTRColorControlClusterMoveHueParams() MTRColorControlClusterMoveHueParams {
	return getMTRColorControlClusterMoveHueParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterMoveHueParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams
type MTRColorControlClusterMoveHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveHueParamsFrom constructs a [MTRColorControlClusterMoveHueParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveHueParams {
	return MTRColorControlClusterMoveHueParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterMoveHueParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterMoveHueParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterMoveHueParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterMoveHueParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterMoveHueParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams/moveMode
func (m_ MTRColorControlClusterMoveHueParams) MoveMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("moveMode"))
	return rv
}/* debug [instance_properties/getter]: moveMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams/moveMode
func (m_ MTRColorControlClusterMoveHueParams) SetMoveMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoveMode:"), value)
}/* debug [instance_properties/setter]: moveMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams/optionsMask
func (m_ MTRColorControlClusterMoveHueParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams/optionsMask
func (m_ MTRColorControlClusterMoveHueParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams/optionsOverride
func (m_ MTRColorControlClusterMoveHueParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams/optionsOverride
func (m_ MTRColorControlClusterMoveHueParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams/rate
func (m_ MTRColorControlClusterMoveHueParams) Rate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams/rate
func (m_ MTRColorControlClusterMoveHueParams) SetRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveHueParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveHueParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveHueParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveHueParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterMoveHueParams */



