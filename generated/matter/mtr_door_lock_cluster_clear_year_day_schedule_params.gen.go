// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterClearYearDayScheduleParams */


/* debug [class_header]: Header for MTRDoorLockClusterClearYearDayScheduleParams */
// The class instance for the [MTRDoorLockClusterClearYearDayScheduleParams] class.
var (
	MTRDoorLockClusterClearYearDayScheduleParamsClass     _MTRDoorLockClusterClearYearDayScheduleParamsClass
	MTRDoorLockClusterClearYearDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearYearDayScheduleParamsClass() _MTRDoorLockClusterClearYearDayScheduleParamsClass {
	MTRDoorLockClusterClearYearDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearYearDayScheduleParamsClass = _MTRDoorLockClusterClearYearDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterClearYearDayScheduleParams")}
	})
	return MTRDoorLockClusterClearYearDayScheduleParamsClass
}

type _MTRDoorLockClusterClearYearDayScheduleParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterClearYearDayScheduleParams */
// An interface definition for the [MTRDoorLockClusterClearYearDayScheduleParams] class.
type IMTRDoorLockClusterClearYearDayScheduleParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterClearYearDayScheduleParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	YearDayIndex() objc.IObject /* cross-framework: NSNumber */
	SetYearDayIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterClearYearDayScheduleParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterClearYearDayScheduleParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearYearDayScheduleParamsClass) Alloc() MTRDoorLockClusterClearYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearYearDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterClearYearDayScheduleParamsClass) New() MTRDoorLockClusterClearYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearYearDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) Init() MTRDoorLockClusterClearYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearYearDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) Autorelease() MTRDoorLockClusterClearYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearYearDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearYearDayScheduleParams creates a new MTRDoorLockClusterClearYearDayScheduleParams instance.
func NewMTRDoorLockClusterClearYearDayScheduleParams() MTRDoorLockClusterClearYearDayScheduleParams {
	return getMTRDoorLockClusterClearYearDayScheduleParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterClearYearDayScheduleParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearYearDayScheduleParams
type MTRDoorLockClusterClearYearDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearYearDayScheduleParamsFrom constructs a [MTRDoorLockClusterClearYearDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearYearDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearYearDayScheduleParams {
	return MTRDoorLockClusterClearYearDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterClearYearDayScheduleParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterClearYearDayScheduleParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterClearYearDayScheduleParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterClearYearDayScheduleParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterClearYearDayScheduleParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearYearDayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearYearDayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearYearDayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearYearDayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearYearDayScheduleParams/userIndex
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearYearDayScheduleParams/userIndex
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearYearDayScheduleParams/yearDayIndex
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) YearDayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("yearDayIndex"))
	return rv
}/* debug [instance_properties/getter]: yearDayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearYearDayScheduleParams/yearDayIndex
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) SetYearDayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setYearDayIndex:"), value)
}/* debug [instance_properties/setter]: yearDayIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterClearYearDayScheduleParams */



