// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterSetWeekDayScheduleParams */


/* debug [class_header]: Header for MTRDoorLockClusterSetWeekDayScheduleParams */
// The class instance for the [MTRDoorLockClusterSetWeekDayScheduleParams] class.
var (
	MTRDoorLockClusterSetWeekDayScheduleParamsClass     _MTRDoorLockClusterSetWeekDayScheduleParamsClass
	MTRDoorLockClusterSetWeekDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetWeekDayScheduleParamsClass() _MTRDoorLockClusterSetWeekDayScheduleParamsClass {
	MTRDoorLockClusterSetWeekDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetWeekDayScheduleParamsClass = _MTRDoorLockClusterSetWeekDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterSetWeekDayScheduleParams")}
	})
	return MTRDoorLockClusterSetWeekDayScheduleParamsClass
}

type _MTRDoorLockClusterSetWeekDayScheduleParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterSetWeekDayScheduleParams */
// An interface definition for the [MTRDoorLockClusterSetWeekDayScheduleParams] class.
type IMTRDoorLockClusterSetWeekDayScheduleParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterSetWeekDayScheduleParams */
	// properties:
	DaysMask() objc.IObject /* cross-framework: NSNumber */
	SetDaysMask(value objc.IObject /* cross-framework: NSNumber */)
	EndHour() objc.IObject /* cross-framework: NSNumber */
	SetEndHour(value objc.IObject /* cross-framework: NSNumber */)
	EndMinute() objc.IObject /* cross-framework: NSNumber */
	SetEndMinute(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StartHour() objc.IObject /* cross-framework: NSNumber */
	SetStartHour(value objc.IObject /* cross-framework: NSNumber */)
	StartMinute() objc.IObject /* cross-framework: NSNumber */
	SetStartMinute(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	WeekDayIndex() objc.IObject /* cross-framework: NSNumber */
	SetWeekDayIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterSetWeekDayScheduleParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterSetWeekDayScheduleParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetWeekDayScheduleParamsClass) Alloc() MTRDoorLockClusterSetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetWeekDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterSetWeekDayScheduleParamsClass) New() MTRDoorLockClusterSetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetWeekDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) Init() MTRDoorLockClusterSetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetWeekDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) Autorelease() MTRDoorLockClusterSetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetWeekDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetWeekDayScheduleParams creates a new MTRDoorLockClusterSetWeekDayScheduleParams instance.
func NewMTRDoorLockClusterSetWeekDayScheduleParams() MTRDoorLockClusterSetWeekDayScheduleParams {
	return getMTRDoorLockClusterSetWeekDayScheduleParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterSetWeekDayScheduleParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams
type MTRDoorLockClusterSetWeekDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetWeekDayScheduleParamsFrom constructs a [MTRDoorLockClusterSetWeekDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetWeekDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetWeekDayScheduleParams {
	return MTRDoorLockClusterSetWeekDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterSetWeekDayScheduleParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterSetWeekDayScheduleParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterSetWeekDayScheduleParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterSetWeekDayScheduleParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterSetWeekDayScheduleParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/daysMask
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) DaysMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("daysMask"))
	return rv
}/* debug [instance_properties/getter]: daysMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/daysMask
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetDaysMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDaysMask:"), value)
}/* debug [instance_properties/setter]: daysMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/endHour
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) EndHour() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endHour"))
	return rv
}/* debug [instance_properties/getter]: endHour */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/endHour
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetEndHour(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndHour:"), value)
}/* debug [instance_properties/setter]: endHour */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/endMinute
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) EndMinute() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endMinute"))
	return rv
}/* debug [instance_properties/getter]: endMinute */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/endMinute
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetEndMinute(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndMinute:"), value)
}/* debug [instance_properties/setter]: endMinute */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/startHour
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) StartHour() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startHour"))
	return rv
}/* debug [instance_properties/getter]: startHour */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/startHour
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetStartHour(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartHour:"), value)
}/* debug [instance_properties/setter]: startHour */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/startMinute
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) StartMinute() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startMinute"))
	return rv
}/* debug [instance_properties/getter]: startMinute */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/startMinute
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetStartMinute(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartMinute:"), value)
}/* debug [instance_properties/setter]: startMinute */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/userIndex
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/userIndex
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/weekDayIndex
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) WeekDayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("weekDayIndex"))
	return rv
}/* debug [instance_properties/getter]: weekDayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetWeekDayScheduleParams/weekDayIndex
func (m_ MTRDoorLockClusterSetWeekDayScheduleParams) SetWeekDayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWeekDayIndex:"), value)
}/* debug [instance_properties/setter]: weekDayIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterSetWeekDayScheduleParams */



