// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterClearHolidayScheduleParams */


/* debug [class_header]: Header for MTRDoorLockClusterClearHolidayScheduleParams */
// The class instance for the [MTRDoorLockClusterClearHolidayScheduleParams] class.
var (
	MTRDoorLockClusterClearHolidayScheduleParamsClass     _MTRDoorLockClusterClearHolidayScheduleParamsClass
	MTRDoorLockClusterClearHolidayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearHolidayScheduleParamsClass() _MTRDoorLockClusterClearHolidayScheduleParamsClass {
	MTRDoorLockClusterClearHolidayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearHolidayScheduleParamsClass = _MTRDoorLockClusterClearHolidayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterClearHolidayScheduleParams")}
	})
	return MTRDoorLockClusterClearHolidayScheduleParamsClass
}

type _MTRDoorLockClusterClearHolidayScheduleParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterClearHolidayScheduleParams */
// An interface definition for the [MTRDoorLockClusterClearHolidayScheduleParams] class.
type IMTRDoorLockClusterClearHolidayScheduleParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterClearHolidayScheduleParams */
	// properties:
	HolidayIndex() objc.IObject /* cross-framework: NSNumber */
	SetHolidayIndex(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterClearHolidayScheduleParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterClearHolidayScheduleParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearHolidayScheduleParamsClass) Alloc() MTRDoorLockClusterClearHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearHolidayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterClearHolidayScheduleParamsClass) New() MTRDoorLockClusterClearHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearHolidayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) Init() MTRDoorLockClusterClearHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearHolidayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) Autorelease() MTRDoorLockClusterClearHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearHolidayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearHolidayScheduleParams creates a new MTRDoorLockClusterClearHolidayScheduleParams instance.
func NewMTRDoorLockClusterClearHolidayScheduleParams() MTRDoorLockClusterClearHolidayScheduleParams {
	return getMTRDoorLockClusterClearHolidayScheduleParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterClearHolidayScheduleParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearHolidayScheduleParams
type MTRDoorLockClusterClearHolidayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearHolidayScheduleParamsFrom constructs a [MTRDoorLockClusterClearHolidayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearHolidayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearHolidayScheduleParams {
	return MTRDoorLockClusterClearHolidayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterClearHolidayScheduleParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterClearHolidayScheduleParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterClearHolidayScheduleParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterClearHolidayScheduleParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterClearHolidayScheduleParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearHolidayScheduleParams/holidayIndex
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) HolidayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("holidayIndex"))
	return rv
}/* debug [instance_properties/getter]: holidayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearHolidayScheduleParams/holidayIndex
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) SetHolidayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHolidayIndex:"), value)
}/* debug [instance_properties/setter]: holidayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearHolidayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearHolidayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearHolidayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearHolidayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterClearHolidayScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterClearHolidayScheduleParams */



