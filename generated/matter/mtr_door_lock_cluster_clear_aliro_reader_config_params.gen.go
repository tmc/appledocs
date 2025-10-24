// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterClearAliroReaderConfigParams */


/* debug [class_header]: Header for MTRDoorLockClusterClearAliroReaderConfigParams */
// The class instance for the [MTRDoorLockClusterClearAliroReaderConfigParams] class.
var (
	MTRDoorLockClusterClearAliroReaderConfigParamsClass     _MTRDoorLockClusterClearAliroReaderConfigParamsClass
	MTRDoorLockClusterClearAliroReaderConfigParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearAliroReaderConfigParamsClass() _MTRDoorLockClusterClearAliroReaderConfigParamsClass {
	MTRDoorLockClusterClearAliroReaderConfigParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearAliroReaderConfigParamsClass = _MTRDoorLockClusterClearAliroReaderConfigParamsClass{objc.GetClass("MTRDoorLockClusterClearAliroReaderConfigParams")}
	})
	return MTRDoorLockClusterClearAliroReaderConfigParamsClass
}

type _MTRDoorLockClusterClearAliroReaderConfigParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterClearAliroReaderConfigParams */
// An interface definition for the [MTRDoorLockClusterClearAliroReaderConfigParams] class.
type IMTRDoorLockClusterClearAliroReaderConfigParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterClearAliroReaderConfigParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterClearAliroReaderConfigParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterClearAliroReaderConfigParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearAliroReaderConfigParamsClass) Alloc() MTRDoorLockClusterClearAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterClearAliroReaderConfigParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterClearAliroReaderConfigParamsClass) New() MTRDoorLockClusterClearAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterClearAliroReaderConfigParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearAliroReaderConfigParams) Init() MTRDoorLockClusterClearAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterClearAliroReaderConfigParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearAliroReaderConfigParams) Autorelease() MTRDoorLockClusterClearAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterClearAliroReaderConfigParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearAliroReaderConfigParams creates a new MTRDoorLockClusterClearAliroReaderConfigParams instance.
func NewMTRDoorLockClusterClearAliroReaderConfigParams() MTRDoorLockClusterClearAliroReaderConfigParams {
	return getMTRDoorLockClusterClearAliroReaderConfigParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterClearAliroReaderConfigParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearAliroReaderConfigParams
type MTRDoorLockClusterClearAliroReaderConfigParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearAliroReaderConfigParamsFrom constructs a [MTRDoorLockClusterClearAliroReaderConfigParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearAliroReaderConfigParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearAliroReaderConfigParams {
	return MTRDoorLockClusterClearAliroReaderConfigParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterClearAliroReaderConfigParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterClearAliroReaderConfigParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterClearAliroReaderConfigParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterClearAliroReaderConfigParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterClearAliroReaderConfigParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearAliroReaderConfigParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterClearAliroReaderConfigParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearAliroReaderConfigParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterClearAliroReaderConfigParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearaliroreaderconfigparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearAliroReaderConfigParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearaliroreaderconfigparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearAliroReaderConfigParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterClearAliroReaderConfigParams */



