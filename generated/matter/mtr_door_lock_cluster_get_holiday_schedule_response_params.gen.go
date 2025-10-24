// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterGetHolidayScheduleResponseParams */


/* debug [class_header]: Header for MTRDoorLockClusterGetHolidayScheduleResponseParams */
// The class instance for the [MTRDoorLockClusterGetHolidayScheduleResponseParams] class.
var (
	MTRDoorLockClusterGetHolidayScheduleResponseParamsClass     _MTRDoorLockClusterGetHolidayScheduleResponseParamsClass
	MTRDoorLockClusterGetHolidayScheduleResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetHolidayScheduleResponseParamsClass() _MTRDoorLockClusterGetHolidayScheduleResponseParamsClass {
	MTRDoorLockClusterGetHolidayScheduleResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetHolidayScheduleResponseParamsClass = _MTRDoorLockClusterGetHolidayScheduleResponseParamsClass{objc.GetClass("MTRDoorLockClusterGetHolidayScheduleResponseParams")}
	})
	return MTRDoorLockClusterGetHolidayScheduleResponseParamsClass
}

type _MTRDoorLockClusterGetHolidayScheduleResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterGetHolidayScheduleResponseParams */
// An interface definition for the [MTRDoorLockClusterGetHolidayScheduleResponseParams] class.
type IMTRDoorLockClusterGetHolidayScheduleResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterGetHolidayScheduleResponseParams */
	// properties:
	HolidayIndex() objc.IObject /* cross-framework: NSNumber */
	SetHolidayIndex(value objc.IObject /* cross-framework: NSNumber */)
	LocalEndTime() objc.IObject /* cross-framework: NSNumber */
	SetLocalEndTime(value objc.IObject /* cross-framework: NSNumber */)
	LocalStartTime() objc.IObject /* cross-framework: NSNumber */
	SetLocalStartTime(value objc.IObject /* cross-framework: NSNumber */)
	OperatingMode() objc.IObject /* cross-framework: NSNumber */
	SetOperatingMode(value objc.IObject /* cross-framework: NSNumber */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterGetHolidayScheduleResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterGetHolidayScheduleResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetHolidayScheduleResponseParamsClass) Alloc() MTRDoorLockClusterGetHolidayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterGetHolidayScheduleResponseParamsClass) New() MTRDoorLockClusterGetHolidayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) Init() MTRDoorLockClusterGetHolidayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) Autorelease() MTRDoorLockClusterGetHolidayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetHolidayScheduleResponseParams creates a new MTRDoorLockClusterGetHolidayScheduleResponseParams instance.
func NewMTRDoorLockClusterGetHolidayScheduleResponseParams() MTRDoorLockClusterGetHolidayScheduleResponseParams {
	return getMTRDoorLockClusterGetHolidayScheduleResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterGetHolidayScheduleResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams
type MTRDoorLockClusterGetHolidayScheduleResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetHolidayScheduleResponseParamsFrom constructs a [MTRDoorLockClusterGetHolidayScheduleResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetHolidayScheduleResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetHolidayScheduleResponseParams {
	return MTRDoorLockClusterGetHolidayScheduleResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterGetHolidayScheduleResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams/init(responseValue:)
func NewMTRDoorLockClusterGetHolidayScheduleResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRDoorLockClusterGetHolidayScheduleResponseParams {
	instance := getMTRDoorLockClusterGetHolidayScheduleResponseParamsClass().Alloc()
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDoorLockClusterGetHolidayScheduleResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterGetHolidayScheduleResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterGetHolidayScheduleResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterGetHolidayScheduleResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterGetHolidayScheduleResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams/holidayIndex
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) HolidayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("holidayIndex"))
	return rv
}/* debug [instance_properties/getter]: holidayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams/holidayIndex
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetHolidayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHolidayIndex:"), value)
}/* debug [instance_properties/setter]: holidayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams/localEndTime
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) LocalEndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("localEndTime"))
	return rv
}/* debug [instance_properties/getter]: localEndTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams/localEndTime
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetLocalEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalEndTime:"), value)
}/* debug [instance_properties/setter]: localEndTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams/localStartTime
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) LocalStartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("localStartTime"))
	return rv
}/* debug [instance_properties/getter]: localStartTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams/localStartTime
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetLocalStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalStartTime:"), value)
}/* debug [instance_properties/setter]: localStartTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams/operatingMode
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) OperatingMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operatingMode"))
	return rv
}/* debug [instance_properties/getter]: operatingMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams/operatingMode
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetOperatingMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperatingMode:"), value)
}/* debug [instance_properties/setter]: operatingMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams/status
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams/status
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleResponseParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetHolidayScheduleResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterGetHolidayScheduleResponseParams */


