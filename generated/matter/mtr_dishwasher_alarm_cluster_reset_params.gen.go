// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDishwasherAlarmClusterResetParams */


/* debug [class_header]: Header for MTRDishwasherAlarmClusterResetParams */
// The class instance for the [MTRDishwasherAlarmClusterResetParams] class.
var (
	MTRDishwasherAlarmClusterResetParamsClass     _MTRDishwasherAlarmClusterResetParamsClass
	MTRDishwasherAlarmClusterResetParamsClassOnce sync.Once
)

func getMTRDishwasherAlarmClusterResetParamsClass() _MTRDishwasherAlarmClusterResetParamsClass {
	MTRDishwasherAlarmClusterResetParamsClassOnce.Do(func() {
		MTRDishwasherAlarmClusterResetParamsClass = _MTRDishwasherAlarmClusterResetParamsClass{objc.GetClass("MTRDishwasherAlarmClusterResetParams")}
	})
	return MTRDishwasherAlarmClusterResetParamsClass
}

type _MTRDishwasherAlarmClusterResetParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDishwasherAlarmClusterResetParams */
// An interface definition for the [MTRDishwasherAlarmClusterResetParams] class.
type IMTRDishwasherAlarmClusterResetParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDishwasherAlarmClusterResetParams */
	// properties:
	Alarms() objc.IObject /* cross-framework: NSNumber */
	SetAlarms(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDishwasherAlarmClusterResetParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDishwasherAlarmClusterResetParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDishwasherAlarmClusterResetParamsClass) Alloc() MTRDishwasherAlarmClusterResetParams {
	rv := objc.Send[MTRDishwasherAlarmClusterResetParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDishwasherAlarmClusterResetParamsClass) New() MTRDishwasherAlarmClusterResetParams {
	rv := objc.Send[MTRDishwasherAlarmClusterResetParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDishwasherAlarmClusterResetParams) Init() MTRDishwasherAlarmClusterResetParams {
	rv := objc.Send[MTRDishwasherAlarmClusterResetParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDishwasherAlarmClusterResetParams) Autorelease() MTRDishwasherAlarmClusterResetParams {
	rv := objc.Send[MTRDishwasherAlarmClusterResetParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDishwasherAlarmClusterResetParams creates a new MTRDishwasherAlarmClusterResetParams instance.
func NewMTRDishwasherAlarmClusterResetParams() MTRDishwasherAlarmClusterResetParams {
	return getMTRDishwasherAlarmClusterResetParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDishwasherAlarmClusterResetParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterResetParams
type MTRDishwasherAlarmClusterResetParams struct {
	objectivec.Object
}

// MTRDishwasherAlarmClusterResetParamsFrom constructs a [MTRDishwasherAlarmClusterResetParams] from an unsafe.Pointer.
func MTRDishwasherAlarmClusterResetParamsFrom(ptr unsafe.Pointer) MTRDishwasherAlarmClusterResetParams {
	return MTRDishwasherAlarmClusterResetParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDishwasherAlarmClusterResetParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDishwasherAlarmClusterResetParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDishwasherAlarmClusterResetParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDishwasherAlarmClusterResetParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDishwasherAlarmClusterResetParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterResetParams/alarms
func (m_ MTRDishwasherAlarmClusterResetParams) Alarms() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("alarms"))
	return rv
}/* debug [instance_properties/getter]: alarms */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterResetParams/alarms
func (m_ MTRDishwasherAlarmClusterResetParams) SetAlarms(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlarms:"), value)
}/* debug [instance_properties/setter]: alarms */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwasheralarmclusterresetparams/serversideprocessingtimeout
func (m_ MTRDishwasherAlarmClusterResetParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwasheralarmclusterresetparams/serversideprocessingtimeout
func (m_ MTRDishwasherAlarmClusterResetParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwasheralarmclusterresetparams/timedinvoketimeoutms
func (m_ MTRDishwasherAlarmClusterResetParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwasheralarmclusterresetparams/timedinvoketimeoutms
func (m_ MTRDishwasherAlarmClusterResetParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDishwasherAlarmClusterResetParams */



