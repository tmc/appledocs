// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams */
// The class instance for the [MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams] class.
var (
	MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass     _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass
	MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass() _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass {
	MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass = _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams")}
	})
	return MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass
}

type _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams */
// An interface definition for the [MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams] class.
type IMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams */
	// properties:
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	Cause() objc.IObject /* cross-framework: NSNumber */
	SetCause(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass) Alloc() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass) New() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) Init() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) Autorelease() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams creates a new MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams instance.
func NewMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams() MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	return getMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams
type MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsFrom constructs a [MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams {
	return MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams */

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterrequestconstraintbasedforecastparams/cause
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) Cause() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cause"))
	return rv
}/* debug [instance_properties/getter]: cause */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterrequestconstraintbasedforecastparams/cause
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) SetCause(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}/* debug [instance_properties/setter]: cause */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterrequestconstraintbasedforecastparams/serversideprocessingtimeout
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterrequestconstraintbasedforecastparams/serversideprocessingtimeout
func (m_ MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams */



