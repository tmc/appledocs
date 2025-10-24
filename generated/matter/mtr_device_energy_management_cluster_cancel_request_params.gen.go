// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterCancelRequestParams */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterCancelRequestParams */
// The class instance for the [MTRDeviceEnergyManagementClusterCancelRequestParams] class.
var (
	MTRDeviceEnergyManagementClusterCancelRequestParamsClass     _MTRDeviceEnergyManagementClusterCancelRequestParamsClass
	MTRDeviceEnergyManagementClusterCancelRequestParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterCancelRequestParamsClass() _MTRDeviceEnergyManagementClusterCancelRequestParamsClass {
	MTRDeviceEnergyManagementClusterCancelRequestParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterCancelRequestParamsClass = _MTRDeviceEnergyManagementClusterCancelRequestParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterCancelRequestParams")}
	})
	return MTRDeviceEnergyManagementClusterCancelRequestParamsClass
}

type _MTRDeviceEnergyManagementClusterCancelRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterCancelRequestParams */
// An interface definition for the [MTRDeviceEnergyManagementClusterCancelRequestParams] class.
type IMTRDeviceEnergyManagementClusterCancelRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterCancelRequestParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterCancelRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterCancelRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterCancelRequestParamsClass) Alloc() MTRDeviceEnergyManagementClusterCancelRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterCancelRequestParamsClass) New() MTRDeviceEnergyManagementClusterCancelRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterCancelRequestParams) Init() MTRDeviceEnergyManagementClusterCancelRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterCancelRequestParams) Autorelease() MTRDeviceEnergyManagementClusterCancelRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterCancelRequestParams creates a new MTRDeviceEnergyManagementClusterCancelRequestParams instance.
func NewMTRDeviceEnergyManagementClusterCancelRequestParams() MTRDeviceEnergyManagementClusterCancelRequestParams {
	return getMTRDeviceEnergyManagementClusterCancelRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterCancelRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelRequestParams
type MTRDeviceEnergyManagementClusterCancelRequestParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterCancelRequestParamsFrom constructs a [MTRDeviceEnergyManagementClusterCancelRequestParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterCancelRequestParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterCancelRequestParams {
	return MTRDeviceEnergyManagementClusterCancelRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterCancelRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterCancelRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterCancelRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterCancelRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterCancelRequestParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterCancelRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterCancelRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustercancelrequestparams/timedinvoketimeoutms
func (m_ MTRDeviceEnergyManagementClusterCancelRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustercancelrequestparams/timedinvoketimeoutms
func (m_ MTRDeviceEnergyManagementClusterCancelRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterCancelRequestParams */



