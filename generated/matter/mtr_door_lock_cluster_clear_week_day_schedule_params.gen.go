// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterClearWeekDayScheduleParams */


/* debug [class_header]: Header for MTRDoorLockClusterClearWeekDayScheduleParams */
// The class instance for the [MTRDoorLockClusterClearWeekDayScheduleParams] class.
var (
	MTRDoorLockClusterClearWeekDayScheduleParamsClass     _MTRDoorLockClusterClearWeekDayScheduleParamsClass
	MTRDoorLockClusterClearWeekDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearWeekDayScheduleParamsClass() _MTRDoorLockClusterClearWeekDayScheduleParamsClass {
	MTRDoorLockClusterClearWeekDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearWeekDayScheduleParamsClass = _MTRDoorLockClusterClearWeekDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterClearWeekDayScheduleParams")}
	})
	return MTRDoorLockClusterClearWeekDayScheduleParamsClass
}

type _MTRDoorLockClusterClearWeekDayScheduleParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterClearWeekDayScheduleParams */
// An interface definition for the [MTRDoorLockClusterClearWeekDayScheduleParams] class.
type IMTRDoorLockClusterClearWeekDayScheduleParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterClearWeekDayScheduleParams */
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

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterClearWeekDayScheduleParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterClearWeekDayScheduleParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearWeekDayScheduleParamsClass) Alloc() MTRDoorLockClusterClearWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearWeekDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterClearWeekDayScheduleParamsClass) New() MTRDoorLockClusterClearWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearWeekDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) Init() MTRDoorLockClusterClearWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearWeekDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) Autorelease() MTRDoorLockClusterClearWeekDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearWeekDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearWeekDayScheduleParams creates a new MTRDoorLockClusterClearWeekDayScheduleParams instance.
func NewMTRDoorLockClusterClearWeekDayScheduleParams() MTRDoorLockClusterClearWeekDayScheduleParams {
	return getMTRDoorLockClusterClearWeekDayScheduleParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterClearWeekDayScheduleParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearWeekDayScheduleParams
type MTRDoorLockClusterClearWeekDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearWeekDayScheduleParamsFrom constructs a [MTRDoorLockClusterClearWeekDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearWeekDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearWeekDayScheduleParams {
	return MTRDoorLockClusterClearWeekDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterClearWeekDayScheduleParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterClearWeekDayScheduleParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterClearWeekDayScheduleParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterClearWeekDayScheduleParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterClearWeekDayScheduleParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearWeekDayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearWeekDayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearWeekDayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearWeekDayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearWeekDayScheduleParams/userIndex
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearWeekDayScheduleParams/userIndex
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearWeekDayScheduleParams/weekDayIndex
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) WeekDayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("weekDayIndex"))
	return rv
}/* debug [instance_properties/getter]: weekDayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearWeekDayScheduleParams/weekDayIndex
func (m_ MTRDoorLockClusterClearWeekDayScheduleParams) SetWeekDayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWeekDayIndex:"), value)
}/* debug [instance_properties/setter]: weekDayIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterClearWeekDayScheduleParams */



