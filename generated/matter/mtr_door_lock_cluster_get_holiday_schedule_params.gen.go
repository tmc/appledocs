// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterGetHolidayScheduleParams */


/* debug [class_header]: Header for MTRDoorLockClusterGetHolidayScheduleParams */
// The class instance for the [MTRDoorLockClusterGetHolidayScheduleParams] class.
var (
	MTRDoorLockClusterGetHolidayScheduleParamsClass     _MTRDoorLockClusterGetHolidayScheduleParamsClass
	MTRDoorLockClusterGetHolidayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetHolidayScheduleParamsClass() _MTRDoorLockClusterGetHolidayScheduleParamsClass {
	MTRDoorLockClusterGetHolidayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetHolidayScheduleParamsClass = _MTRDoorLockClusterGetHolidayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterGetHolidayScheduleParams")}
	})
	return MTRDoorLockClusterGetHolidayScheduleParamsClass
}

type _MTRDoorLockClusterGetHolidayScheduleParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterGetHolidayScheduleParams */
// An interface definition for the [MTRDoorLockClusterGetHolidayScheduleParams] class.
type IMTRDoorLockClusterGetHolidayScheduleParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterGetHolidayScheduleParams */
	// properties:
	HolidayIndex() objc.IObject /* cross-framework: NSNumber */
	SetHolidayIndex(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterGetHolidayScheduleParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterGetHolidayScheduleParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetHolidayScheduleParamsClass) Alloc() MTRDoorLockClusterGetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterGetHolidayScheduleParamsClass) New() MTRDoorLockClusterGetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) Init() MTRDoorLockClusterGetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) Autorelease() MTRDoorLockClusterGetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetHolidayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetHolidayScheduleParams creates a new MTRDoorLockClusterGetHolidayScheduleParams instance.
func NewMTRDoorLockClusterGetHolidayScheduleParams() MTRDoorLockClusterGetHolidayScheduleParams {
	return getMTRDoorLockClusterGetHolidayScheduleParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterGetHolidayScheduleParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleParams
type MTRDoorLockClusterGetHolidayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetHolidayScheduleParamsFrom constructs a [MTRDoorLockClusterGetHolidayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetHolidayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetHolidayScheduleParams {
	return MTRDoorLockClusterGetHolidayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterGetHolidayScheduleParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterGetHolidayScheduleParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterGetHolidayScheduleParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterGetHolidayScheduleParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterGetHolidayScheduleParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleParams/holidayIndex
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) HolidayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("holidayIndex"))
	return rv
}/* debug [instance_properties/getter]: holidayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleParams/holidayIndex
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) SetHolidayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHolidayIndex:"), value)
}/* debug [instance_properties/setter]: holidayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetHolidayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetHolidayScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterGetHolidayScheduleParams */



