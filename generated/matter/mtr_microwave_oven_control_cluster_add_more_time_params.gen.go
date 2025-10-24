// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMicrowaveOvenControlClusterAddMoreTimeParams */


/* debug [class_header]: Header for MTRMicrowaveOvenControlClusterAddMoreTimeParams */
// The class instance for the [MTRMicrowaveOvenControlClusterAddMoreTimeParams] class.
var (
	MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass     _MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass
	MTRMicrowaveOvenControlClusterAddMoreTimeParamsClassOnce sync.Once
)

func getMTRMicrowaveOvenControlClusterAddMoreTimeParamsClass() _MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass {
	MTRMicrowaveOvenControlClusterAddMoreTimeParamsClassOnce.Do(func() {
		MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass = _MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass{objc.GetClass("MTRMicrowaveOvenControlClusterAddMoreTimeParams")}
	})
	return MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass
}

type _MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMicrowaveOvenControlClusterAddMoreTimeParams */
// An interface definition for the [MTRMicrowaveOvenControlClusterAddMoreTimeParams] class.
type IMTRMicrowaveOvenControlClusterAddMoreTimeParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMicrowaveOvenControlClusterAddMoreTimeParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimeToAdd() objc.IObject /* cross-framework: NSNumber */
	SetTimeToAdd(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMicrowaveOvenControlClusterAddMoreTimeParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMicrowaveOvenControlClusterAddMoreTimeParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass) Alloc() MTRMicrowaveOvenControlClusterAddMoreTimeParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterAddMoreTimeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass) New() MTRMicrowaveOvenControlClusterAddMoreTimeParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterAddMoreTimeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) Init() MTRMicrowaveOvenControlClusterAddMoreTimeParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterAddMoreTimeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) Autorelease() MTRMicrowaveOvenControlClusterAddMoreTimeParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterAddMoreTimeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMicrowaveOvenControlClusterAddMoreTimeParams creates a new MTRMicrowaveOvenControlClusterAddMoreTimeParams instance.
func NewMTRMicrowaveOvenControlClusterAddMoreTimeParams() MTRMicrowaveOvenControlClusterAddMoreTimeParams {
	return getMTRMicrowaveOvenControlClusterAddMoreTimeParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMicrowaveOvenControlClusterAddMoreTimeParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterAddMoreTimeParams
type MTRMicrowaveOvenControlClusterAddMoreTimeParams struct {
	objectivec.Object
}

// MTRMicrowaveOvenControlClusterAddMoreTimeParamsFrom constructs a [MTRMicrowaveOvenControlClusterAddMoreTimeParams] from an unsafe.Pointer.
func MTRMicrowaveOvenControlClusterAddMoreTimeParamsFrom(ptr unsafe.Pointer) MTRMicrowaveOvenControlClusterAddMoreTimeParams {
	return MTRMicrowaveOvenControlClusterAddMoreTimeParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMicrowaveOvenControlClusterAddMoreTimeParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMicrowaveOvenControlClusterAddMoreTimeParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMicrowaveOvenControlClusterAddMoreTimeParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMicrowaveOvenControlClusterAddMoreTimeParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMicrowaveOvenControlClusterAddMoreTimeParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterAddMoreTimeParams/serverSideProcessingTimeout
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterAddMoreTimeParams/serverSideProcessingTimeout
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovencontrolclusteraddmoretimeparams/timetoadd
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) TimeToAdd() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timeToAdd"))
	return rv
}/* debug [instance_properties/getter]: timeToAdd */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovencontrolclusteraddmoretimeparams/timetoadd
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) SetTimeToAdd(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeToAdd:"), value)
}/* debug [instance_properties/setter]: timeToAdd */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovencontrolclusteraddmoretimeparams/timedinvoketimeoutms
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovencontrolclusteraddmoretimeparams/timedinvoketimeoutms
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMicrowaveOvenControlClusterAddMoreTimeParams */



