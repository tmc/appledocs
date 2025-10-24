// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROvenCavityOperationalStateClusterStartParams */


/* debug [class_header]: Header for MTROvenCavityOperationalStateClusterStartParams */
// The class instance for the [MTROvenCavityOperationalStateClusterStartParams] class.
var (
	MTROvenCavityOperationalStateClusterStartParamsClass     _MTROvenCavityOperationalStateClusterStartParamsClass
	MTROvenCavityOperationalStateClusterStartParamsClassOnce sync.Once
)

func getMTROvenCavityOperationalStateClusterStartParamsClass() _MTROvenCavityOperationalStateClusterStartParamsClass {
	MTROvenCavityOperationalStateClusterStartParamsClassOnce.Do(func() {
		MTROvenCavityOperationalStateClusterStartParamsClass = _MTROvenCavityOperationalStateClusterStartParamsClass{objc.GetClass("MTROvenCavityOperationalStateClusterStartParams")}
	})
	return MTROvenCavityOperationalStateClusterStartParamsClass
}

type _MTROvenCavityOperationalStateClusterStartParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROvenCavityOperationalStateClusterStartParams */
// An interface definition for the [MTROvenCavityOperationalStateClusterStartParams] class.
type IMTROvenCavityOperationalStateClusterStartParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROvenCavityOperationalStateClusterStartParams */
	// properties:
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROvenCavityOperationalStateClusterStartParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROvenCavityOperationalStateClusterStartParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROvenCavityOperationalStateClusterStartParamsClass) Alloc() MTROvenCavityOperationalStateClusterStartParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStartParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROvenCavityOperationalStateClusterStartParamsClass) New() MTROvenCavityOperationalStateClusterStartParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStartParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenCavityOperationalStateClusterStartParams) Init() MTROvenCavityOperationalStateClusterStartParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStartParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenCavityOperationalStateClusterStartParams) Autorelease() MTROvenCavityOperationalStateClusterStartParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStartParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenCavityOperationalStateClusterStartParams creates a new MTROvenCavityOperationalStateClusterStartParams instance.
func NewMTROvenCavityOperationalStateClusterStartParams() MTROvenCavityOperationalStateClusterStartParams {
	return getMTROvenCavityOperationalStateClusterStartParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROvenCavityOperationalStateClusterStartParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStartParams
type MTROvenCavityOperationalStateClusterStartParams struct {
	objectivec.Object
}

// MTROvenCavityOperationalStateClusterStartParamsFrom constructs a [MTROvenCavityOperationalStateClusterStartParams] from an unsafe.Pointer.
func MTROvenCavityOperationalStateClusterStartParamsFrom(ptr unsafe.Pointer) MTROvenCavityOperationalStateClusterStartParams {
	return MTROvenCavityOperationalStateClusterStartParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROvenCavityOperationalStateClusterStartParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROvenCavityOperationalStateClusterStartParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROvenCavityOperationalStateClusterStartParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROvenCavityOperationalStateClusterStartParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROvenCavityOperationalStateClusterStartParams */

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStartParams/timedInvokeTimeoutMs
func (m_ MTROvenCavityOperationalStateClusterStartParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStartParams/timedInvokeTimeoutMs
func (m_ MTROvenCavityOperationalStateClusterStartParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovencavityoperationalstateclusterstartparams/serversideprocessingtimeout
func (m_ MTROvenCavityOperationalStateClusterStartParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovencavityoperationalstateclusterstartparams/serversideprocessingtimeout
func (m_ MTROvenCavityOperationalStateClusterStartParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROvenCavityOperationalStateClusterStartParams */



