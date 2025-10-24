// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRLevelControlClusterMoveToClosestFrequencyParams */


/* debug [class_header]: Header for MTRLevelControlClusterMoveToClosestFrequencyParams */
// The class instance for the [MTRLevelControlClusterMoveToClosestFrequencyParams] class.
var (
	MTRLevelControlClusterMoveToClosestFrequencyParamsClass     _MTRLevelControlClusterMoveToClosestFrequencyParamsClass
	MTRLevelControlClusterMoveToClosestFrequencyParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveToClosestFrequencyParamsClass() _MTRLevelControlClusterMoveToClosestFrequencyParamsClass {
	MTRLevelControlClusterMoveToClosestFrequencyParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveToClosestFrequencyParamsClass = _MTRLevelControlClusterMoveToClosestFrequencyParamsClass{objc.GetClass("MTRLevelControlClusterMoveToClosestFrequencyParams")}
	})
	return MTRLevelControlClusterMoveToClosestFrequencyParamsClass
}

type _MTRLevelControlClusterMoveToClosestFrequencyParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRLevelControlClusterMoveToClosestFrequencyParams */
// An interface definition for the [MTRLevelControlClusterMoveToClosestFrequencyParams] class.
type IMTRLevelControlClusterMoveToClosestFrequencyParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRLevelControlClusterMoveToClosestFrequencyParams */
	// properties:
	Frequency() objc.IObject /* cross-framework: NSNumber */
	SetFrequency(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRLevelControlClusterMoveToClosestFrequencyParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRLevelControlClusterMoveToClosestFrequencyParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveToClosestFrequencyParamsClass) Alloc() MTRLevelControlClusterMoveToClosestFrequencyParams {
	rv := objc.Send[MTRLevelControlClusterMoveToClosestFrequencyParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRLevelControlClusterMoveToClosestFrequencyParamsClass) New() MTRLevelControlClusterMoveToClosestFrequencyParams {
	rv := objc.Send[MTRLevelControlClusterMoveToClosestFrequencyParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) Init() MTRLevelControlClusterMoveToClosestFrequencyParams {
	rv := objc.Send[MTRLevelControlClusterMoveToClosestFrequencyParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) Autorelease() MTRLevelControlClusterMoveToClosestFrequencyParams {
	rv := objc.Send[MTRLevelControlClusterMoveToClosestFrequencyParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveToClosestFrequencyParams creates a new MTRLevelControlClusterMoveToClosestFrequencyParams instance.
func NewMTRLevelControlClusterMoveToClosestFrequencyParams() MTRLevelControlClusterMoveToClosestFrequencyParams {
	return getMTRLevelControlClusterMoveToClosestFrequencyParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRLevelControlClusterMoveToClosestFrequencyParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToClosestFrequencyParams
type MTRLevelControlClusterMoveToClosestFrequencyParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveToClosestFrequencyParamsFrom constructs a [MTRLevelControlClusterMoveToClosestFrequencyParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveToClosestFrequencyParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveToClosestFrequencyParams {
	return MTRLevelControlClusterMoveToClosestFrequencyParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRLevelControlClusterMoveToClosestFrequencyParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRLevelControlClusterMoveToClosestFrequencyParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRLevelControlClusterMoveToClosestFrequencyParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRLevelControlClusterMoveToClosestFrequencyParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRLevelControlClusterMoveToClosestFrequencyParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToClosestFrequencyParams/frequency
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) Frequency() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("frequency"))
	return rv
}/* debug [instance_properties/getter]: frequency */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToClosestFrequencyParams/frequency
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) SetFrequency(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFrequency:"), value)
}/* debug [instance_properties/setter]: frequency */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToClosestFrequencyParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToClosestFrequencyParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToClosestFrequencyParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToClosestFrequencyParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterMoveToClosestFrequencyParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRLevelControlClusterMoveToClosestFrequencyParams */



