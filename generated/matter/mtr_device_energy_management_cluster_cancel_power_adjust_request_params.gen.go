// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams */
// The class instance for the [MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams] class.
var (
	MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass     _MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass
	MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass() _MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass {
	MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass = _MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams")}
	})
	return MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass
}

type _MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams */
// An interface definition for the [MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams] class.
type IMTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass) Alloc() MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass) New() MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams) Init() MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams) Autorelease() MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams creates a new MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams instance.
func NewMTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams() MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams {
	return getMTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams
type MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsFrom constructs a [MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams {
	return MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustercancelpoweradjustrequestparams/timedinvoketimeoutms
func (m_ MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustercancelpoweradjustrequestparams/timedinvoketimeoutms
func (m_ MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams */



