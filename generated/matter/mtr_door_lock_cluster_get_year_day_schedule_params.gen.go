// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterGetYearDayScheduleParams */


/* debug [class_header]: Header for MTRDoorLockClusterGetYearDayScheduleParams */
// The class instance for the [MTRDoorLockClusterGetYearDayScheduleParams] class.
var (
	MTRDoorLockClusterGetYearDayScheduleParamsClass     _MTRDoorLockClusterGetYearDayScheduleParamsClass
	MTRDoorLockClusterGetYearDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetYearDayScheduleParamsClass() _MTRDoorLockClusterGetYearDayScheduleParamsClass {
	MTRDoorLockClusterGetYearDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetYearDayScheduleParamsClass = _MTRDoorLockClusterGetYearDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterGetYearDayScheduleParams")}
	})
	return MTRDoorLockClusterGetYearDayScheduleParamsClass
}

type _MTRDoorLockClusterGetYearDayScheduleParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterGetYearDayScheduleParams */
// An interface definition for the [MTRDoorLockClusterGetYearDayScheduleParams] class.
type IMTRDoorLockClusterGetYearDayScheduleParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterGetYearDayScheduleParams */
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

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterGetYearDayScheduleParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterGetYearDayScheduleParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetYearDayScheduleParamsClass) Alloc() MTRDoorLockClusterGetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterGetYearDayScheduleParamsClass) New() MTRDoorLockClusterGetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) Init() MTRDoorLockClusterGetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) Autorelease() MTRDoorLockClusterGetYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetYearDayScheduleParams creates a new MTRDoorLockClusterGetYearDayScheduleParams instance.
func NewMTRDoorLockClusterGetYearDayScheduleParams() MTRDoorLockClusterGetYearDayScheduleParams {
	return getMTRDoorLockClusterGetYearDayScheduleParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterGetYearDayScheduleParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleParams
type MTRDoorLockClusterGetYearDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetYearDayScheduleParamsFrom constructs a [MTRDoorLockClusterGetYearDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetYearDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetYearDayScheduleParams {
	return MTRDoorLockClusterGetYearDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterGetYearDayScheduleParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterGetYearDayScheduleParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterGetYearDayScheduleParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterGetYearDayScheduleParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterGetYearDayScheduleParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleParams/userIndex
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleParams/userIndex
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleParams/yearDayIndex
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) YearDayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("yearDayIndex"))
	return rv
}/* debug [instance_properties/getter]: yearDayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleParams/yearDayIndex
func (m_ MTRDoorLockClusterGetYearDayScheduleParams) SetYearDayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setYearDayIndex:"), value)
}/* debug [instance_properties/setter]: yearDayIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterGetYearDayScheduleParams */



