// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterSetHolidayScheduleParams */


/* debug [class_header]: Header for MTRDoorLockClusterSetHolidayScheduleParams */
// The class instance for the [MTRDoorLockClusterSetHolidayScheduleParams] class.
var (
	MTRDoorLockClusterSetHolidayScheduleParamsClass     _MTRDoorLockClusterSetHolidayScheduleParamsClass
	MTRDoorLockClusterSetHolidayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetHolidayScheduleParamsClass() _MTRDoorLockClusterSetHolidayScheduleParamsClass {
	MTRDoorLockClusterSetHolidayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetHolidayScheduleParamsClass = _MTRDoorLockClusterSetHolidayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterSetHolidayScheduleParams")}
	})
	return MTRDoorLockClusterSetHolidayScheduleParamsClass
}

type _MTRDoorLockClusterSetHolidayScheduleParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterSetHolidayScheduleParams */
// An interface definition for the [MTRDoorLockClusterSetHolidayScheduleParams] class.
type IMTRDoorLockClusterSetHolidayScheduleParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterSetHolidayScheduleParams */
	// properties:
	HolidayIndex() objc.IObject /* cross-framework: NSNumber */
	SetHolidayIndex(value objc.IObject /* cross-framework: NSNumber */)
	LocalEndTime() objc.IObject /* cross-framework: NSNumber */
	SetLocalEndTime(value objc.IObject /* cross-framework: NSNumber */)
	LocalStartTime() objc.IObject /* cross-framework: NSNumber */
	SetLocalStartTime(value objc.IObject /* cross-framework: NSNumber */)
	OperatingMode() objc.IObject /* cross-framework: NSNumber */
	SetOperatingMode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterSetHolidayScheduleParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterSetHolidayScheduleParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetHolidayScheduleParamsClass) Alloc() MTRDoorLockClusterSetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetHolidayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterSetHolidayScheduleParamsClass) New() MTRDoorLockClusterSetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetHolidayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) Init() MTRDoorLockClusterSetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetHolidayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) Autorelease() MTRDoorLockClusterSetHolidayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetHolidayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetHolidayScheduleParams creates a new MTRDoorLockClusterSetHolidayScheduleParams instance.
func NewMTRDoorLockClusterSetHolidayScheduleParams() MTRDoorLockClusterSetHolidayScheduleParams {
	return getMTRDoorLockClusterSetHolidayScheduleParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterSetHolidayScheduleParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetHolidayScheduleParams
type MTRDoorLockClusterSetHolidayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetHolidayScheduleParamsFrom constructs a [MTRDoorLockClusterSetHolidayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetHolidayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetHolidayScheduleParams {
	return MTRDoorLockClusterSetHolidayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterSetHolidayScheduleParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterSetHolidayScheduleParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterSetHolidayScheduleParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterSetHolidayScheduleParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterSetHolidayScheduleParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetHolidayScheduleParams/holidayIndex
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) HolidayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("holidayIndex"))
	return rv
}/* debug [instance_properties/getter]: holidayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetHolidayScheduleParams/holidayIndex
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetHolidayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHolidayIndex:"), value)
}/* debug [instance_properties/setter]: holidayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetHolidayScheduleParams/localEndTime
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) LocalEndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("localEndTime"))
	return rv
}/* debug [instance_properties/getter]: localEndTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetHolidayScheduleParams/localEndTime
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetLocalEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalEndTime:"), value)
}/* debug [instance_properties/setter]: localEndTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetHolidayScheduleParams/localStartTime
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) LocalStartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("localStartTime"))
	return rv
}/* debug [instance_properties/getter]: localStartTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetHolidayScheduleParams/localStartTime
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetLocalStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalStartTime:"), value)
}/* debug [instance_properties/setter]: localStartTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetHolidayScheduleParams/operatingMode
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) OperatingMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operatingMode"))
	return rv
}/* debug [instance_properties/getter]: operatingMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetHolidayScheduleParams/operatingMode
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetOperatingMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperatingMode:"), value)
}/* debug [instance_properties/setter]: operatingMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetHolidayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetHolidayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetHolidayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetHolidayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetHolidayScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterSetHolidayScheduleParams */



