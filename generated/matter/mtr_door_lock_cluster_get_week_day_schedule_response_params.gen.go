// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterGetWeekDayScheduleResponseParams */


/* debug [class_header]: Header for MTRDoorLockClusterGetWeekDayScheduleResponseParams */
// The class instance for the [MTRDoorLockClusterGetWeekDayScheduleResponseParams] class.
var (
	MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass     _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass
	MTRDoorLockClusterGetWeekDayScheduleResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetWeekDayScheduleResponseParamsClass() _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass {
	MTRDoorLockClusterGetWeekDayScheduleResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass = _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass{objc.GetClass("MTRDoorLockClusterGetWeekDayScheduleResponseParams")}
	})
	return MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass
}

type _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterGetWeekDayScheduleResponseParams */
// An interface definition for the [MTRDoorLockClusterGetWeekDayScheduleResponseParams] class.
type IMTRDoorLockClusterGetWeekDayScheduleResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterGetWeekDayScheduleResponseParams */
	// properties:
	DaysMask() objc.IObject /* cross-framework: NSNumber */
	SetDaysMask(value objc.IObject /* cross-framework: NSNumber */)
	EndHour() objc.IObject /* cross-framework: NSNumber */
	SetEndHour(value objc.IObject /* cross-framework: NSNumber */)
	EndMinute() objc.IObject /* cross-framework: NSNumber */
	SetEndMinute(value objc.IObject /* cross-framework: NSNumber */)
	StartHour() objc.IObject /* cross-framework: NSNumber */
	SetStartHour(value objc.IObject /* cross-framework: NSNumber */)
	StartMinute() objc.IObject /* cross-framework: NSNumber */
	SetStartMinute(value objc.IObject /* cross-framework: NSNumber */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	WeekDayIndex() objc.IObject /* cross-framework: NSNumber */
	SetWeekDayIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterGetWeekDayScheduleResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterGetWeekDayScheduleResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass) Alloc() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterGetWeekDayScheduleResponseParamsClass) New() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) Init() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) Autorelease() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetWeekDayScheduleResponseParams creates a new MTRDoorLockClusterGetWeekDayScheduleResponseParams instance.
func NewMTRDoorLockClusterGetWeekDayScheduleResponseParams() MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	return getMTRDoorLockClusterGetWeekDayScheduleResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterGetWeekDayScheduleResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams
type MTRDoorLockClusterGetWeekDayScheduleResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetWeekDayScheduleResponseParamsFrom constructs a [MTRDoorLockClusterGetWeekDayScheduleResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetWeekDayScheduleResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	return MTRDoorLockClusterGetWeekDayScheduleResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterGetWeekDayScheduleResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/init(responseValue:)
func NewMTRDoorLockClusterGetWeekDayScheduleResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRDoorLockClusterGetWeekDayScheduleResponseParams {
	instance := getMTRDoorLockClusterGetWeekDayScheduleResponseParamsClass().Alloc()
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDoorLockClusterGetWeekDayScheduleResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterGetWeekDayScheduleResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterGetWeekDayScheduleResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterGetWeekDayScheduleResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterGetWeekDayScheduleResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/daysMask
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) DaysMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("daysMask"))
	return rv
}/* debug [instance_properties/getter]: daysMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/daysMask
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetDaysMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDaysMask:"), value)
}/* debug [instance_properties/setter]: daysMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/endHour
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) EndHour() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endHour"))
	return rv
}/* debug [instance_properties/getter]: endHour */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/endHour
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetEndHour(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndHour:"), value)
}/* debug [instance_properties/setter]: endHour */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/endMinute
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) EndMinute() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endMinute"))
	return rv
}/* debug [instance_properties/getter]: endMinute */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/endMinute
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetEndMinute(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndMinute:"), value)
}/* debug [instance_properties/setter]: endMinute */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/startHour
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) StartHour() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startHour"))
	return rv
}/* debug [instance_properties/getter]: startHour */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/startHour
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetStartHour(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartHour:"), value)
}/* debug [instance_properties/setter]: startHour */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/startMinute
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) StartMinute() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startMinute"))
	return rv
}/* debug [instance_properties/getter]: startMinute */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/startMinute
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetStartMinute(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartMinute:"), value)
}/* debug [instance_properties/setter]: startMinute */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/status
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/status
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/userIndex
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/userIndex
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/weekDayIndex
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) WeekDayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("weekDayIndex"))
	return rv
}/* debug [instance_properties/getter]: weekDayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleResponseParams/weekDayIndex
func (m_ MTRDoorLockClusterGetWeekDayScheduleResponseParams) SetWeekDayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWeekDayIndex:"), value)
}/* debug [instance_properties/setter]: weekDayIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterGetWeekDayScheduleResponseParams */


