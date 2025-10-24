// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterGetYearDayScheduleResponseParams */


/* debug [class_header]: Header for MTRDoorLockClusterGetYearDayScheduleResponseParams */
// The class instance for the [MTRDoorLockClusterGetYearDayScheduleResponseParams] class.
var (
	MTRDoorLockClusterGetYearDayScheduleResponseParamsClass     _MTRDoorLockClusterGetYearDayScheduleResponseParamsClass
	MTRDoorLockClusterGetYearDayScheduleResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetYearDayScheduleResponseParamsClass() _MTRDoorLockClusterGetYearDayScheduleResponseParamsClass {
	MTRDoorLockClusterGetYearDayScheduleResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetYearDayScheduleResponseParamsClass = _MTRDoorLockClusterGetYearDayScheduleResponseParamsClass{objc.GetClass("MTRDoorLockClusterGetYearDayScheduleResponseParams")}
	})
	return MTRDoorLockClusterGetYearDayScheduleResponseParamsClass
}

type _MTRDoorLockClusterGetYearDayScheduleResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterGetYearDayScheduleResponseParams */
// An interface definition for the [MTRDoorLockClusterGetYearDayScheduleResponseParams] class.
type IMTRDoorLockClusterGetYearDayScheduleResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterGetYearDayScheduleResponseParams */
	// properties:
	LocalEndTime() objc.IObject /* cross-framework: NSNumber */
	SetLocalEndTime(value objc.IObject /* cross-framework: NSNumber */)
	LocalStartTime() objc.IObject /* cross-framework: NSNumber */
	SetLocalStartTime(value objc.IObject /* cross-framework: NSNumber */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	YearDayIndex() objc.IObject /* cross-framework: NSNumber */
	SetYearDayIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterGetYearDayScheduleResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterGetYearDayScheduleResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetYearDayScheduleResponseParamsClass) Alloc() MTRDoorLockClusterGetYearDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterGetYearDayScheduleResponseParamsClass) New() MTRDoorLockClusterGetYearDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) Init() MTRDoorLockClusterGetYearDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) Autorelease() MTRDoorLockClusterGetYearDayScheduleResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetYearDayScheduleResponseParams creates a new MTRDoorLockClusterGetYearDayScheduleResponseParams instance.
func NewMTRDoorLockClusterGetYearDayScheduleResponseParams() MTRDoorLockClusterGetYearDayScheduleResponseParams {
	return getMTRDoorLockClusterGetYearDayScheduleResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterGetYearDayScheduleResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams
type MTRDoorLockClusterGetYearDayScheduleResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetYearDayScheduleResponseParamsFrom constructs a [MTRDoorLockClusterGetYearDayScheduleResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetYearDayScheduleResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetYearDayScheduleResponseParams {
	return MTRDoorLockClusterGetYearDayScheduleResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterGetYearDayScheduleResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams/init(responseValue:)
func NewMTRDoorLockClusterGetYearDayScheduleResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRDoorLockClusterGetYearDayScheduleResponseParams {
	instance := getMTRDoorLockClusterGetYearDayScheduleResponseParamsClass().Alloc()
	rv := objc.Send[MTRDoorLockClusterGetYearDayScheduleResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDoorLockClusterGetYearDayScheduleResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterGetYearDayScheduleResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterGetYearDayScheduleResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterGetYearDayScheduleResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterGetYearDayScheduleResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams/localEndTime
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) LocalEndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("localEndTime"))
	return rv
}/* debug [instance_properties/getter]: localEndTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams/localEndTime
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetLocalEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalEndTime:"), value)
}/* debug [instance_properties/setter]: localEndTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams/localStartTime
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) LocalStartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("localStartTime"))
	return rv
}/* debug [instance_properties/getter]: localStartTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams/localStartTime
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetLocalStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocalStartTime:"), value)
}/* debug [instance_properties/setter]: localStartTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams/status
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams/status
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams/userIndex
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams/userIndex
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams/yearDayIndex
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) YearDayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("yearDayIndex"))
	return rv
}/* debug [instance_properties/getter]: yearDayIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetYearDayScheduleResponseParams/yearDayIndex
func (m_ MTRDoorLockClusterGetYearDayScheduleResponseParams) SetYearDayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setYearDayIndex:"), value)
}/* debug [instance_properties/setter]: yearDayIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterGetYearDayScheduleResponseParams */


