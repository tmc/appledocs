// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterModifyForecastRequestParams */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterModifyForecastRequestParams */
// The class instance for the [MTRDeviceEnergyManagementClusterModifyForecastRequestParams] class.
var (
	MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass     _MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass
	MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass() _MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass {
	MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass = _MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterModifyForecastRequestParams")}
	})
	return MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass
}

type _MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterModifyForecastRequestParams */
// An interface definition for the [MTRDeviceEnergyManagementClusterModifyForecastRequestParams] class.
type IMTRDeviceEnergyManagementClusterModifyForecastRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterModifyForecastRequestParams */
	// properties:
	ForecastID() objc.IObject /* cross-framework: NSNumber */
	SetForecastID(value objc.IObject /* cross-framework: NSNumber */)
	Cause() objc.IObject /* cross-framework: NSNumber */
	SetCause(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterModifyForecastRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterModifyForecastRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass) Alloc() MTRDeviceEnergyManagementClusterModifyForecastRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterModifyForecastRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass) New() MTRDeviceEnergyManagementClusterModifyForecastRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterModifyForecastRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) Init() MTRDeviceEnergyManagementClusterModifyForecastRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterModifyForecastRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) Autorelease() MTRDeviceEnergyManagementClusterModifyForecastRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterModifyForecastRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterModifyForecastRequestParams creates a new MTRDeviceEnergyManagementClusterModifyForecastRequestParams instance.
func NewMTRDeviceEnergyManagementClusterModifyForecastRequestParams() MTRDeviceEnergyManagementClusterModifyForecastRequestParams {
	return getMTRDeviceEnergyManagementClusterModifyForecastRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterModifyForecastRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterModifyForecastRequestParams
type MTRDeviceEnergyManagementClusterModifyForecastRequestParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterModifyForecastRequestParamsFrom constructs a [MTRDeviceEnergyManagementClusterModifyForecastRequestParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterModifyForecastRequestParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterModifyForecastRequestParams {
	return MTRDeviceEnergyManagementClusterModifyForecastRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterModifyForecastRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterModifyForecastRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterModifyForecastRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterModifyForecastRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterModifyForecastRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterModifyForecastRequestParams/forecastID
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) ForecastID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("forecastID"))
	return rv
}/* debug [instance_properties/getter]: forecastID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterModifyForecastRequestParams/forecastID
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) SetForecastID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setForecastID:"), value)
}/* debug [instance_properties/setter]: forecastID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustermodifyforecastrequestparams/cause
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) Cause() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cause"))
	return rv
}/* debug [instance_properties/getter]: cause */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustermodifyforecastrequestparams/cause
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) SetCause(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}/* debug [instance_properties/setter]: cause */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustermodifyforecastrequestparams/serversideprocessingtimeout
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustermodifyforecastrequestparams/serversideprocessingtimeout
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustermodifyforecastrequestparams/timedinvoketimeoutms
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclustermodifyforecastrequestparams/timedinvoketimeoutms
func (m_ MTRDeviceEnergyManagementClusterModifyForecastRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterModifyForecastRequestParams */



