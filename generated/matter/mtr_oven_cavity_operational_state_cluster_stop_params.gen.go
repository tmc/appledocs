// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROvenCavityOperationalStateClusterStopParams */


/* debug [class_header]: Header for MTROvenCavityOperationalStateClusterStopParams */
// The class instance for the [MTROvenCavityOperationalStateClusterStopParams] class.
var (
	MTROvenCavityOperationalStateClusterStopParamsClass     _MTROvenCavityOperationalStateClusterStopParamsClass
	MTROvenCavityOperationalStateClusterStopParamsClassOnce sync.Once
)

func getMTROvenCavityOperationalStateClusterStopParamsClass() _MTROvenCavityOperationalStateClusterStopParamsClass {
	MTROvenCavityOperationalStateClusterStopParamsClassOnce.Do(func() {
		MTROvenCavityOperationalStateClusterStopParamsClass = _MTROvenCavityOperationalStateClusterStopParamsClass{objc.GetClass("MTROvenCavityOperationalStateClusterStopParams")}
	})
	return MTROvenCavityOperationalStateClusterStopParamsClass
}

type _MTROvenCavityOperationalStateClusterStopParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROvenCavityOperationalStateClusterStopParams */
// An interface definition for the [MTROvenCavityOperationalStateClusterStopParams] class.
type IMTROvenCavityOperationalStateClusterStopParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROvenCavityOperationalStateClusterStopParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROvenCavityOperationalStateClusterStopParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROvenCavityOperationalStateClusterStopParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROvenCavityOperationalStateClusterStopParamsClass) Alloc() MTROvenCavityOperationalStateClusterStopParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStopParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROvenCavityOperationalStateClusterStopParamsClass) New() MTROvenCavityOperationalStateClusterStopParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStopParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenCavityOperationalStateClusterStopParams) Init() MTROvenCavityOperationalStateClusterStopParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStopParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenCavityOperationalStateClusterStopParams) Autorelease() MTROvenCavityOperationalStateClusterStopParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStopParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenCavityOperationalStateClusterStopParams creates a new MTROvenCavityOperationalStateClusterStopParams instance.
func NewMTROvenCavityOperationalStateClusterStopParams() MTROvenCavityOperationalStateClusterStopParams {
	return getMTROvenCavityOperationalStateClusterStopParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROvenCavityOperationalStateClusterStopParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStopParams
type MTROvenCavityOperationalStateClusterStopParams struct {
	objectivec.Object
}

// MTROvenCavityOperationalStateClusterStopParamsFrom constructs a [MTROvenCavityOperationalStateClusterStopParams] from an unsafe.Pointer.
func MTROvenCavityOperationalStateClusterStopParamsFrom(ptr unsafe.Pointer) MTROvenCavityOperationalStateClusterStopParams {
	return MTROvenCavityOperationalStateClusterStopParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROvenCavityOperationalStateClusterStopParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROvenCavityOperationalStateClusterStopParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROvenCavityOperationalStateClusterStopParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROvenCavityOperationalStateClusterStopParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROvenCavityOperationalStateClusterStopParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStopParams/serverSideProcessingTimeout
func (m_ MTROvenCavityOperationalStateClusterStopParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStopParams/serverSideProcessingTimeout
func (m_ MTROvenCavityOperationalStateClusterStopParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovencavityoperationalstateclusterstopparams/timedinvoketimeoutms
func (m_ MTROvenCavityOperationalStateClusterStopParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovencavityoperationalstateclusterstopparams/timedinvoketimeoutms
func (m_ MTROvenCavityOperationalStateClusterStopParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROvenCavityOperationalStateClusterStopParams */



