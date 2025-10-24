// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDishwasherAlarmClusterModifyEnabledAlarmsParams */


/* debug [class_header]: Header for MTRDishwasherAlarmClusterModifyEnabledAlarmsParams */
// The class instance for the [MTRDishwasherAlarmClusterModifyEnabledAlarmsParams] class.
var (
	MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass     _MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass
	MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClassOnce sync.Once
)

func getMTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass() _MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass {
	MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClassOnce.Do(func() {
		MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass = _MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass{objc.GetClass("MTRDishwasherAlarmClusterModifyEnabledAlarmsParams")}
	})
	return MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass
}

type _MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDishwasherAlarmClusterModifyEnabledAlarmsParams */
// An interface definition for the [MTRDishwasherAlarmClusterModifyEnabledAlarmsParams] class.
type IMTRDishwasherAlarmClusterModifyEnabledAlarmsParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDishwasherAlarmClusterModifyEnabledAlarmsParams */
	// properties:
	Mask() objc.IObject /* cross-framework: NSNumber */
	SetMask(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDishwasherAlarmClusterModifyEnabledAlarmsParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDishwasherAlarmClusterModifyEnabledAlarmsParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass) Alloc() MTRDishwasherAlarmClusterModifyEnabledAlarmsParams {
	rv := objc.Send[MTRDishwasherAlarmClusterModifyEnabledAlarmsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass) New() MTRDishwasherAlarmClusterModifyEnabledAlarmsParams {
	rv := objc.Send[MTRDishwasherAlarmClusterModifyEnabledAlarmsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) Init() MTRDishwasherAlarmClusterModifyEnabledAlarmsParams {
	rv := objc.Send[MTRDishwasherAlarmClusterModifyEnabledAlarmsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) Autorelease() MTRDishwasherAlarmClusterModifyEnabledAlarmsParams {
	rv := objc.Send[MTRDishwasherAlarmClusterModifyEnabledAlarmsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDishwasherAlarmClusterModifyEnabledAlarmsParams creates a new MTRDishwasherAlarmClusterModifyEnabledAlarmsParams instance.
func NewMTRDishwasherAlarmClusterModifyEnabledAlarmsParams() MTRDishwasherAlarmClusterModifyEnabledAlarmsParams {
	return getMTRDishwasherAlarmClusterModifyEnabledAlarmsParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDishwasherAlarmClusterModifyEnabledAlarmsParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterModifyEnabledAlarmsParams
type MTRDishwasherAlarmClusterModifyEnabledAlarmsParams struct {
	objectivec.Object
}

// MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsFrom constructs a [MTRDishwasherAlarmClusterModifyEnabledAlarmsParams] from an unsafe.Pointer.
func MTRDishwasherAlarmClusterModifyEnabledAlarmsParamsFrom(ptr unsafe.Pointer) MTRDishwasherAlarmClusterModifyEnabledAlarmsParams {
	return MTRDishwasherAlarmClusterModifyEnabledAlarmsParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDishwasherAlarmClusterModifyEnabledAlarmsParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDishwasherAlarmClusterModifyEnabledAlarmsParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDishwasherAlarmClusterModifyEnabledAlarmsParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDishwasherAlarmClusterModifyEnabledAlarmsParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDishwasherAlarmClusterModifyEnabledAlarmsParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterModifyEnabledAlarmsParams/mask
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) Mask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mask"))
	return rv
}/* debug [instance_properties/getter]: mask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterModifyEnabledAlarmsParams/mask
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) SetMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMask:"), value)
}/* debug [instance_properties/setter]: mask */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwasheralarmclustermodifyenabledalarmsparams/serversideprocessingtimeout
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwasheralarmclustermodifyenabledalarmsparams/serversideprocessingtimeout
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwasheralarmclustermodifyenabledalarmsparams/timedinvoketimeoutms
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwasheralarmclustermodifyenabledalarmsparams/timedinvoketimeoutms
func (m_ MTRDishwasherAlarmClusterModifyEnabledAlarmsParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDishwasherAlarmClusterModifyEnabledAlarmsParams */



