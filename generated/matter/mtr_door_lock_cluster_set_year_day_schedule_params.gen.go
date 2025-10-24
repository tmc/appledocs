// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterSetYearDayScheduleParams */


/* debug [class_header]: Header for MTRDoorLockClusterSetYearDayScheduleParams */
// The class instance for the [MTRDoorLockClusterSetYearDayScheduleParams] class.
var (
	MTRDoorLockClusterSetYearDayScheduleParamsClass     _MTRDoorLockClusterSetYearDayScheduleParamsClass
	MTRDoorLockClusterSetYearDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetYearDayScheduleParamsClass() _MTRDoorLockClusterSetYearDayScheduleParamsClass {
	MTRDoorLockClusterSetYearDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetYearDayScheduleParamsClass = _MTRDoorLockClusterSetYearDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterSetYearDayScheduleParams")}
	})
	return MTRDoorLockClusterSetYearDayScheduleParamsClass
}

type _MTRDoorLockClusterSetYearDayScheduleParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterSetYearDayScheduleParams */
// An interface definition for the [MTRDoorLockClusterSetYearDayScheduleParams] class.
type IMTRDoorLockClusterSetYearDayScheduleParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterSetYearDayScheduleParams */
	// properties:
	LocalEndTime() objc.IObject /* cross-framework: NSNumber */
	SetLocalEndTime(value objc.IObject /* cross-framework: NSNumber */)
	LocalStartTime() objc.IObject /* cross-framework: NSNumber */
	SetLocalStartTime(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	YearDayIndex() objc.IObject /* cross-framework: NSNumber */
	SetYearDayIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterSetYearDayScheduleParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterSetYearDayScheduleParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetYearDayScheduleParamsClass) Alloc() MTRDoorLockClusterSetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetYearDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterSetYearDayScheduleParamsClass) New() MTRDoorLockClusterSetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetYearDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) Init() MTRDoorLockClusterSetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetYearDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) Autorelease() MTRDoorLockClusterSetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterSetYearDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetYearDayScheduleParams creates a new MTRDoorLockClusterSetYearDayScheduleParams instance.
func NewMTRDoorLockClusterSetYearDayScheduleParams() MTRDoorLockClusterSetYearDayScheduleParams {
	return getMTRDoorLockClusterSetYearDayScheduleParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterSetYearDayScheduleParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams
type MTRDoorLockClusterSetYearDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetYearDayScheduleParamsFrom constructs a [MTRDoorLockClusterSetYearDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetYearDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetYearDayScheduleParams {
	return MTRDoorLockClusterSetYearDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterSetYearDayScheduleParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterSetYearDayScheduleParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterSetYearDayScheduleParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterSetYearDayScheduleParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterSetYearDayScheduleParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams/localEndTime
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) LocalEndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("localEndTime"))
	return rv
}/* debug [instance_properties/getter]: localEndTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams/localEndTime
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) SetLocalEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalEndTime:"), value)
}/* debug [instance_properties/setter]: localEndTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams/localStartTime
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) LocalStartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("localStartTime"))
	return rv
}/* debug [instance_properties/getter]: localStartTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams/localStartTime
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) SetLocalStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalStartTime:"), value)
}/* debug [instance_properties/setter]: localStartTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams/userIndex
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams/userIndex
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams/yearDayIndex
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) YearDayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("yearDayIndex"))
	return rv
}/* debug [instance_properties/getter]: yearDayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetYearDayScheduleParams/yearDayIndex
func (m_ MTRDoorLockClusterSetYearDayScheduleParams) SetYearDayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setYearDayIndex:"), value)
}/* debug [instance_properties/setter]: yearDayIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterSetYearDayScheduleParams */



