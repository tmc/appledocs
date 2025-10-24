// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterPowerAdjustRequestParams */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterPowerAdjustRequestParams */
// The class instance for the [MTRDeviceEnergyManagementClusterPowerAdjustRequestParams] class.
var (
	MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass     _MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass
	MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass() _MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass {
	MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass = _MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterPowerAdjustRequestParams")}
	})
	return MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass
}

type _MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterPowerAdjustRequestParams */
// An interface definition for the [MTRDeviceEnergyManagementClusterPowerAdjustRequestParams] class.
type IMTRDeviceEnergyManagementClusterPowerAdjustRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterPowerAdjustRequestParams */
	// properties:
	Cause() objc.IObject /* cross-framework: NSNumber */
	SetCause(value objc.IObject /* cross-framework: NSNumber */)
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	Power() objc.IObject /* cross-framework: NSNumber */
	SetPower(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterPowerAdjustRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterPowerAdjustRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass) Alloc() MTRDeviceEnergyManagementClusterPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass) New() MTRDeviceEnergyManagementClusterPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) Init() MTRDeviceEnergyManagementClusterPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) Autorelease() MTRDeviceEnergyManagementClusterPowerAdjustRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterPowerAdjustRequestParams creates a new MTRDeviceEnergyManagementClusterPowerAdjustRequestParams instance.
func NewMTRDeviceEnergyManagementClusterPowerAdjustRequestParams() MTRDeviceEnergyManagementClusterPowerAdjustRequestParams {
	return getMTRDeviceEnergyManagementClusterPowerAdjustRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterPowerAdjustRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustRequestParams
type MTRDeviceEnergyManagementClusterPowerAdjustRequestParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsFrom constructs a [MTRDeviceEnergyManagementClusterPowerAdjustRequestParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterPowerAdjustRequestParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterPowerAdjustRequestParams {
	return MTRDeviceEnergyManagementClusterPowerAdjustRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterPowerAdjustRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterPowerAdjustRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterPowerAdjustRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterPowerAdjustRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterPowerAdjustRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustRequestParams/cause
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) Cause() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cause"))
	return rv
}/* debug [instance_properties/getter]: cause */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustRequestParams/cause
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) SetCause(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}/* debug [instance_properties/setter]: cause */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjustrequestparams/duration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjustrequestparams/duration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjustrequestparams/power
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) Power() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("power"))
	return rv
}/* debug [instance_properties/getter]: power */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjustrequestparams/power
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) SetPower(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPower:"), value)
}/* debug [instance_properties/setter]: power */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjustrequestparams/serversideprocessingtimeout
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjustrequestparams/serversideprocessingtimeout
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjustrequestparams/timedinvoketimeoutms
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjustrequestparams/timedinvoketimeoutms
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterPowerAdjustRequestParams */



