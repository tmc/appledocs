// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRTemperatureControlClusterSetTemperatureParams */


/* debug [class_header]: Header for MTRTemperatureControlClusterSetTemperatureParams */
// The class instance for the [MTRTemperatureControlClusterSetTemperatureParams] class.
var (
	MTRTemperatureControlClusterSetTemperatureParamsClass     _MTRTemperatureControlClusterSetTemperatureParamsClass
	MTRTemperatureControlClusterSetTemperatureParamsClassOnce sync.Once
)

func getMTRTemperatureControlClusterSetTemperatureParamsClass() _MTRTemperatureControlClusterSetTemperatureParamsClass {
	MTRTemperatureControlClusterSetTemperatureParamsClassOnce.Do(func() {
		MTRTemperatureControlClusterSetTemperatureParamsClass = _MTRTemperatureControlClusterSetTemperatureParamsClass{objc.GetClass("MTRTemperatureControlClusterSetTemperatureParams")}
	})
	return MTRTemperatureControlClusterSetTemperatureParamsClass
}

type _MTRTemperatureControlClusterSetTemperatureParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTemperatureControlClusterSetTemperatureParams */
// An interface definition for the [MTRTemperatureControlClusterSetTemperatureParams] class.
type IMTRTemperatureControlClusterSetTemperatureParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRTemperatureControlClusterSetTemperatureParams */
	// properties:
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TargetTemperature() objc.IObject /* cross-framework: NSNumber */
	SetTargetTemperature(value objc.IObject /* cross-framework: NSNumber */)
	TargetTemperatureLevel() objc.IObject /* cross-framework: NSNumber */
	SetTargetTemperatureLevel(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTemperatureControlClusterSetTemperatureParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTemperatureControlClusterSetTemperatureParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRTemperatureControlClusterSetTemperatureParamsClass) Alloc() MTRTemperatureControlClusterSetTemperatureParams {
	rv := objc.Send[MTRTemperatureControlClusterSetTemperatureParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTemperatureControlClusterSetTemperatureParamsClass) New() MTRTemperatureControlClusterSetTemperatureParams {
	rv := objc.Send[MTRTemperatureControlClusterSetTemperatureParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTemperatureControlClusterSetTemperatureParams) Init() MTRTemperatureControlClusterSetTemperatureParams {
	rv := objc.Send[MTRTemperatureControlClusterSetTemperatureParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTemperatureControlClusterSetTemperatureParams) Autorelease() MTRTemperatureControlClusterSetTemperatureParams {
	rv := objc.Send[MTRTemperatureControlClusterSetTemperatureParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTemperatureControlClusterSetTemperatureParams creates a new MTRTemperatureControlClusterSetTemperatureParams instance.
func NewMTRTemperatureControlClusterSetTemperatureParams() MTRTemperatureControlClusterSetTemperatureParams {
	return getMTRTemperatureControlClusterSetTemperatureParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTemperatureControlClusterSetTemperatureParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams
type MTRTemperatureControlClusterSetTemperatureParams struct {
	objectivec.Object
}

// MTRTemperatureControlClusterSetTemperatureParamsFrom constructs a [MTRTemperatureControlClusterSetTemperatureParams] from an unsafe.Pointer.
func MTRTemperatureControlClusterSetTemperatureParamsFrom(ptr unsafe.Pointer) MTRTemperatureControlClusterSetTemperatureParams {
	return MTRTemperatureControlClusterSetTemperatureParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTemperatureControlClusterSetTemperatureParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTemperatureControlClusterSetTemperatureParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTemperatureControlClusterSetTemperatureParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTemperatureControlClusterSetTemperatureParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTemperatureControlClusterSetTemperatureParams */

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/timedInvokeTimeoutMs
func (m_ MTRTemperatureControlClusterSetTemperatureParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/timedInvokeTimeoutMs
func (m_ MTRTemperatureControlClusterSetTemperatureParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtemperaturecontrolclustersettemperatureparams/serversideprocessingtimeout
func (m_ MTRTemperatureControlClusterSetTemperatureParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtemperaturecontrolclustersettemperatureparams/serversideprocessingtimeout
func (m_ MTRTemperatureControlClusterSetTemperatureParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtemperaturecontrolclustersettemperatureparams/targettemperature
func (m_ MTRTemperatureControlClusterSetTemperatureParams) TargetTemperature() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("targetTemperature"))
	return rv
}/* debug [instance_properties/getter]: targetTemperature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtemperaturecontrolclustersettemperatureparams/targettemperature
func (m_ MTRTemperatureControlClusterSetTemperatureParams) SetTargetTemperature(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetTemperature:"), value)
}/* debug [instance_properties/setter]: targetTemperature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtemperaturecontrolclustersettemperatureparams/targettemperaturelevel
func (m_ MTRTemperatureControlClusterSetTemperatureParams) TargetTemperatureLevel() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("targetTemperatureLevel"))
	return rv
}/* debug [instance_properties/getter]: targetTemperatureLevel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtemperaturecontrolclustersettemperatureparams/targettemperaturelevel
func (m_ MTRTemperatureControlClusterSetTemperatureParams) SetTargetTemperatureLevel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetTemperatureLevel:"), value)
}/* debug [instance_properties/setter]: targetTemperatureLevel */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTemperatureControlClusterSetTemperatureParams */



