// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterSetActiveScheduleRequestParams */


/* debug [class_header]: Header for MTRThermostatClusterSetActiveScheduleRequestParams */
// The class instance for the [MTRThermostatClusterSetActiveScheduleRequestParams] class.
var (
	MTRThermostatClusterSetActiveScheduleRequestParamsClass     _MTRThermostatClusterSetActiveScheduleRequestParamsClass
	MTRThermostatClusterSetActiveScheduleRequestParamsClassOnce sync.Once
)

func getMTRThermostatClusterSetActiveScheduleRequestParamsClass() _MTRThermostatClusterSetActiveScheduleRequestParamsClass {
	MTRThermostatClusterSetActiveScheduleRequestParamsClassOnce.Do(func() {
		MTRThermostatClusterSetActiveScheduleRequestParamsClass = _MTRThermostatClusterSetActiveScheduleRequestParamsClass{objc.GetClass("MTRThermostatClusterSetActiveScheduleRequestParams")}
	})
	return MTRThermostatClusterSetActiveScheduleRequestParamsClass
}

type _MTRThermostatClusterSetActiveScheduleRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterSetActiveScheduleRequestParams */
// An interface definition for the [MTRThermostatClusterSetActiveScheduleRequestParams] class.
type IMTRThermostatClusterSetActiveScheduleRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterSetActiveScheduleRequestParams */
	// properties:
	ScheduleHandle() objc.IObject /* cross-framework: NSData */
	SetScheduleHandle(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterSetActiveScheduleRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterSetActiveScheduleRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterSetActiveScheduleRequestParamsClass) Alloc() MTRThermostatClusterSetActiveScheduleRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActiveScheduleRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterSetActiveScheduleRequestParamsClass) New() MTRThermostatClusterSetActiveScheduleRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActiveScheduleRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) Init() MTRThermostatClusterSetActiveScheduleRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActiveScheduleRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) Autorelease() MTRThermostatClusterSetActiveScheduleRequestParams {
	rv := objc.Send[MTRThermostatClusterSetActiveScheduleRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterSetActiveScheduleRequestParams creates a new MTRThermostatClusterSetActiveScheduleRequestParams instance.
func NewMTRThermostatClusterSetActiveScheduleRequestParams() MTRThermostatClusterSetActiveScheduleRequestParams {
	return getMTRThermostatClusterSetActiveScheduleRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterSetActiveScheduleRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActiveScheduleRequestParams
type MTRThermostatClusterSetActiveScheduleRequestParams struct {
	objectivec.Object
}

// MTRThermostatClusterSetActiveScheduleRequestParamsFrom constructs a [MTRThermostatClusterSetActiveScheduleRequestParams] from an unsafe.Pointer.
func MTRThermostatClusterSetActiveScheduleRequestParamsFrom(ptr unsafe.Pointer) MTRThermostatClusterSetActiveScheduleRequestParams {
	return MTRThermostatClusterSetActiveScheduleRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterSetActiveScheduleRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterSetActiveScheduleRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterSetActiveScheduleRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterSetActiveScheduleRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterSetActiveScheduleRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActiveScheduleRequestParams/scheduleHandle
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) ScheduleHandle() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("scheduleHandle"))
	return rv
}/* debug [instance_properties/getter]: scheduleHandle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterSetActiveScheduleRequestParams/scheduleHandle
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) SetScheduleHandle(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScheduleHandle:"), value)
}/* debug [instance_properties/setter]: scheduleHandle */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetactiveschedulerequestparams/serversideprocessingtimeout
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetactiveschedulerequestparams/serversideprocessingtimeout
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetactiveschedulerequestparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclustersetactiveschedulerequestparams/timedinvoketimeoutms
func (m_ MTRThermostatClusterSetActiveScheduleRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterSetActiveScheduleRequestParams */



