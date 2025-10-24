// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterSetpointRaiseLowerParams */


/* debug [class_header]: Header for MTRThermostatClusterSetpointRaiseLowerParams */
// The class instance for the [MTRThermostatClusterSetpointRaiseLowerParams] class.
var (
	MTRThermostatClusterSetpointRaiseLowerParamsClass     _MTRThermostatClusterSetpointRaiseLowerParamsClass
	MTRThermostatClusterSetpointRaiseLowerParamsClassOnce sync.Once
)

func getMTRThermostatClusterSetpointRaiseLowerParamsClass() _MTRThermostatClusterSetpointRaiseLowerParamsClass {
	MTRThermostatClusterSetpointRaiseLowerParamsClassOnce.Do(func() {
		MTRThermostatClusterSetpointRaiseLowerParamsClass = _MTRThermostatClusterSetpointRaiseLowerParamsClass{objc.GetClass("MTRThermostatClusterSetpointRaiseLowerParams")}
	})
	return MTRThermostatClusterSetpointRaiseLowerParamsClass
}

type _MTRThermostatClusterSetpointRaiseLowerParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterSetpointRaiseLowerParams */
// An interface definition for the [MTRThermostatClusterSetpointRaiseLowerParams] class.
type IMTRThermostatClusterSetpointRaiseLowerParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterSetpointRaiseLowerParams */
	// properties:
	Amount() objc.IObject /* cross-framework: NSNumber */
	SetAmount(value objc.IObject /* cross-framework: NSNumber */)
	Mode() objc.IObject /* cross-framework: NSNumber */
	SetMode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterSetpointRaiseLowerParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterSetpointRaiseLowerParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterSetpointRaiseLowerParamsClass) Alloc() MTRThermostatClusterSetpointRaiseLowerParams {
	rv := objc.Send[MTRThermostatClusterSetpointRaiseLowerParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterSetpointRaiseLowerParamsClass) New() MTRThermostatClusterSetpointRaiseLowerParams {
	rv := objc.Send[MTRThermostatClusterSetpointRaiseLowerParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) Init() MTRThermostatClusterSetpointRaiseLowerParams {
	rv := objc.Send[MTRThermostatClusterSetpointRaiseLowerParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) Autorelease() MTRThermostatClusterSetpointRaiseLowerParams {
	rv := objc.Send[MTRThermostatClusterSetpointRaiseLowerParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterSetpointRaiseLowerParams creates a new MTRThermostatClusterSetpointRaiseLowerParams instance.
func NewMTRThermostatClusterSetpointRaiseLowerParams() MTRThermostatClusterSetpointRaiseLowerParams {
	return getMTRThermostatClusterSetpointRaiseLowerParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterSetpointRaiseLowerParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetpointRaiseLowerParams
type MTRThermostatClusterSetpointRaiseLowerParams struct {
	objectivec.Object
}

// MTRThermostatClusterSetpointRaiseLowerParamsFrom constructs a [MTRThermostatClusterSetpointRaiseLowerParams] from an unsafe.Pointer.
func MTRThermostatClusterSetpointRaiseLowerParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterSetpointRaiseLowerParams {
	return MTRThermostatClusterSetpointRaiseLowerParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterSetpointRaiseLowerParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterSetpointRaiseLowerParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterSetpointRaiseLowerParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterSetpointRaiseLowerParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterSetpointRaiseLowerParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetpointRaiseLowerParams/amount
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) Amount() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("amount"))
	return rv
}/* debug [instance_properties/getter]: amount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetpointRaiseLowerParams/amount
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) SetAmount(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAmount:"), value)
}/* debug [instance_properties/setter]: amount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetpointRaiseLowerParams/mode
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetpointRaiseLowerParams/mode
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetpointRaiseLowerParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetpointRaiseLowerParams/serverSideProcessingTimeout
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetpointRaiseLowerParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetpointRaiseLowerParams/timedInvokeTimeoutMs
func (m_ MTRThermostatClusterSetpointRaiseLowerParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterSetpointRaiseLowerParams */



