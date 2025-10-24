// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterSetUtcTimeParams */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterSetUtcTimeParams */
// The class instance for the [MTRTimeSynchronizationClusterSetUtcTimeParams] class.
var (
	MTRTimeSynchronizationClusterSetUtcTimeParamsClass     _MTRTimeSynchronizationClusterSetUtcTimeParamsClass
	MTRTimeSynchronizationClusterSetUtcTimeParamsClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterSetUtcTimeParamsClass() _MTRTimeSynchronizationClusterSetUtcTimeParamsClass {
	MTRTimeSynchronizationClusterSetUtcTimeParamsClassOnce.Do(func() {
		MTRTimeSynchronizationClusterSetUtcTimeParamsClass = _MTRTimeSynchronizationClusterSetUtcTimeParamsClass{objc.GetClass("MTRTimeSynchronizationClusterSetUtcTimeParams")}
	})
	return MTRTimeSynchronizationClusterSetUtcTimeParamsClass
}

type _MTRTimeSynchronizationClusterSetUtcTimeParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterSetUtcTimeParams */
// An interface definition for the [MTRTimeSynchronizationClusterSetUtcTimeParams] class.
type IMTRTimeSynchronizationClusterSetUtcTimeParams interface {
	IMTRTimeSynchronizationClusterSetUTCTimeParams
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterSetUtcTimeParams */
	// properties:
	Granularity() objc.IObject /* cross-framework: NSNumber */
	SetGranularity(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TimeSource() objc.IObject /* cross-framework: NSNumber */
	SetTimeSource(value objc.IObject /* cross-framework: NSNumber */)
	UtcTime() objc.IObject /* cross-framework: NSNumber */
	SetUtcTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterSetUtcTimeParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterSetUtcTimeParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterSetUtcTimeParamsClass) Alloc() MTRTimeSynchronizationClusterSetUtcTimeParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetUtcTimeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterSetUtcTimeParamsClass) New() MTRTimeSynchronizationClusterSetUtcTimeParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetUtcTimeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) Init() MTRTimeSynchronizationClusterSetUtcTimeParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetUtcTimeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) Autorelease() MTRTimeSynchronizationClusterSetUtcTimeParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetUtcTimeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterSetUtcTimeParams creates a new MTRTimeSynchronizationClusterSetUtcTimeParams instance.
func NewMTRTimeSynchronizationClusterSetUtcTimeParams() MTRTimeSynchronizationClusterSetUtcTimeParams {
	return getMTRTimeSynchronizationClusterSetUtcTimeParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterSetUtcTimeParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetUtcTimeParams-2ms2i
type MTRTimeSynchronizationClusterSetUtcTimeParams struct {
	MTRTimeSynchronizationClusterSetUTCTimeParams
}

// MTRTimeSynchronizationClusterSetUtcTimeParamsFrom constructs a [MTRTimeSynchronizationClusterSetUtcTimeParams] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterSetUtcTimeParamsFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterSetUtcTimeParams {
	return MTRTimeSynchronizationClusterSetUtcTimeParams{
		MTRTimeSynchronizationClusterSetUTCTimeParams: MTRTimeSynchronizationClusterSetUTCTimeParamsFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterSetUtcTimeParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterSetUtcTimeParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterSetUtcTimeParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterSetUtcTimeParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterSetUtcTimeParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetUtcTimeParams-2ms2i/granularity
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) Granularity() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("granularity"))
	return rv
}/* debug [instance_properties/getter]: granularity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetUtcTimeParams-2ms2i/granularity
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetGranularity(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGranularity:"), value)
}/* debug [instance_properties/setter]: granularity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetUtcTimeParams-2ms2i/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetUtcTimeParams-2ms2i/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetUtcTimeParams-2ms2i/timedInvokeTimeoutMs
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetUtcTimeParams-2ms2i/timedInvokeTimeoutMs
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetUtcTimeParams-2ms2i/timeSource
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) TimeSource() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timeSource"))
	return rv
}/* debug [instance_properties/getter]: timeSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetUtcTimeParams-2ms2i/timeSource
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetTimeSource(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeSource:"), value)
}/* debug [instance_properties/setter]: timeSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetUtcTimeParams-2ms2i/utcTime
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) UtcTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("utcTime"))
	return rv
}/* debug [instance_properties/getter]: utcTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetUtcTimeParams-2ms2i/utcTime
func (m_ MTRTimeSynchronizationClusterSetUtcTimeParams) SetUtcTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUtcTime:"), value)
}/* debug [instance_properties/setter]: utcTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterSetUtcTimeParams */



