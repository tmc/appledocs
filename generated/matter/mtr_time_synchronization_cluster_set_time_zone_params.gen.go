// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterSetTimeZoneParams */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterSetTimeZoneParams */
// The class instance for the [MTRTimeSynchronizationClusterSetTimeZoneParams] class.
var (
	MTRTimeSynchronizationClusterSetTimeZoneParamsClass     _MTRTimeSynchronizationClusterSetTimeZoneParamsClass
	MTRTimeSynchronizationClusterSetTimeZoneParamsClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterSetTimeZoneParamsClass() _MTRTimeSynchronizationClusterSetTimeZoneParamsClass {
	MTRTimeSynchronizationClusterSetTimeZoneParamsClassOnce.Do(func() {
		MTRTimeSynchronizationClusterSetTimeZoneParamsClass = _MTRTimeSynchronizationClusterSetTimeZoneParamsClass{objc.GetClass("MTRTimeSynchronizationClusterSetTimeZoneParams")}
	})
	return MTRTimeSynchronizationClusterSetTimeZoneParamsClass
}

type _MTRTimeSynchronizationClusterSetTimeZoneParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterSetTimeZoneParams */
// An interface definition for the [MTRTimeSynchronizationClusterSetTimeZoneParams] class.
type IMTRTimeSynchronizationClusterSetTimeZoneParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterSetTimeZoneParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterSetTimeZoneParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterSetTimeZoneParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterSetTimeZoneParamsClass) Alloc() MTRTimeSynchronizationClusterSetTimeZoneParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterSetTimeZoneParamsClass) New() MTRTimeSynchronizationClusterSetTimeZoneParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterSetTimeZoneParams) Init() MTRTimeSynchronizationClusterSetTimeZoneParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterSetTimeZoneParams) Autorelease() MTRTimeSynchronizationClusterSetTimeZoneParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterSetTimeZoneParams creates a new MTRTimeSynchronizationClusterSetTimeZoneParams instance.
func NewMTRTimeSynchronizationClusterSetTimeZoneParams() MTRTimeSynchronizationClusterSetTimeZoneParams {
	return getMTRTimeSynchronizationClusterSetTimeZoneParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterSetTimeZoneParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneParams
type MTRTimeSynchronizationClusterSetTimeZoneParams struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterSetTimeZoneParamsFrom constructs a [MTRTimeSynchronizationClusterSetTimeZoneParams] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterSetTimeZoneParamsFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterSetTimeZoneParams {
	return MTRTimeSynchronizationClusterSetTimeZoneParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterSetTimeZoneParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterSetTimeZoneParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterSetTimeZoneParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterSetTimeZoneParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterSetTimeZoneParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetTimeZoneParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetTimeZoneParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersettimezoneparams/timedinvoketimeoutms
func (m_ MTRTimeSynchronizationClusterSetTimeZoneParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersettimezoneparams/timedinvoketimeoutms
func (m_ MTRTimeSynchronizationClusterSetTimeZoneParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterSetTimeZoneParams */



