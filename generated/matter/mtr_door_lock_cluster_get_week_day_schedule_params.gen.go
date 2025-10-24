// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterGetWeekDayScheduleParams */


/* debug [class_header]: Header for MTRDoorLockClusterGetWeekDayScheduleParams */
// The class instance for the [MTRDoorLockClusterGetWeekDayScheduleParams] class.
var (
	MTRDoorLockClusterGetWeekDayScheduleParamsClass     _MTRDoorLockClusterGetWeekDayScheduleParamsClass
	MTRDoorLockClusterGetWeekDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetWeekDayScheduleParamsClass() _MTRDoorLockClusterGetWeekDayScheduleParamsClass {
	MTRDoorLockClusterGetWeekDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetWeekDayScheduleParamsClass = _MTRDoorLockClusterGetWeekDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterGetWeekDayScheduleParams")}
	})
	return MTRDoorLockClusterGetWeekDayScheduleParamsClass
}

type _MTRDoorLockClusterGetWeekDayScheduleParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterGetWeekDayScheduleParams */
// An interface definition for the [MTRDoorLockClusterGetWeekDayScheduleParams] class.
type IMTRDoorLockClusterGetWeekDayScheduleParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterGetWeekDayScheduleParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	WeekDayIndex() objc.IObject /* cross-framework: NSNumber */
	SetWeekDayIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterGetWeekDayScheduleParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterGetWeekDayScheduleParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetWeekDayScheduleParamsClass) Alloc() MTRDoorLockClusterGetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterGetWeekDayScheduleParamsClass) New() MTRDoorLockClusterGetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) Init() MTRDoorLockClusterGetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) Autorelease() MTRDoorLockClusterGetWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetWeekDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetWeekDayScheduleParams creates a new MTRDoorLockClusterGetWeekDayScheduleParams instance.
func NewMTRDoorLockClusterGetWeekDayScheduleParams() MTRDoorLockClusterGetWeekDayScheduleParams {
	return getMTRDoorLockClusterGetWeekDayScheduleParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterGetWeekDayScheduleParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleParams
type MTRDoorLockClusterGetWeekDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetWeekDayScheduleParamsFrom constructs a [MTRDoorLockClusterGetWeekDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetWeekDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetWeekDayScheduleParams {
	return MTRDoorLockClusterGetWeekDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterGetWeekDayScheduleParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterGetWeekDayScheduleParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterGetWeekDayScheduleParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterGetWeekDayScheduleParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterGetWeekDayScheduleParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleParams/userIndex
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleParams/userIndex
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleParams/weekDayIndex
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) WeekDayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("weekDayIndex"))
	return rv
}/* debug [instance_properties/getter]: weekDayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetWeekDayScheduleParams/weekDayIndex
func (m_ MTRDoorLockClusterGetWeekDayScheduleParams) SetWeekDayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWeekDayIndex:"), value)
}/* debug [instance_properties/setter]: weekDayIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterGetWeekDayScheduleParams */



