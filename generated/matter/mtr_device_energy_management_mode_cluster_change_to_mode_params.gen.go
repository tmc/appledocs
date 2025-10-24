// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementModeClusterChangeToModeParams */


/* debug [class_header]: Header for MTRDeviceEnergyManagementModeClusterChangeToModeParams */
// The class instance for the [MTRDeviceEnergyManagementModeClusterChangeToModeParams] class.
var (
	MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass     _MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass
	MTRDeviceEnergyManagementModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementModeClusterChangeToModeParamsClass() _MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass {
	MTRDeviceEnergyManagementModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass = _MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass{objc.GetClass("MTRDeviceEnergyManagementModeClusterChangeToModeParams")}
	})
	return MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass
}

type _MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementModeClusterChangeToModeParams */
// An interface definition for the [MTRDeviceEnergyManagementModeClusterChangeToModeParams] class.
type IMTRDeviceEnergyManagementModeClusterChangeToModeParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementModeClusterChangeToModeParams */
	// properties:
	NewMode() objc.IObject /* cross-framework: NSNumber */
	SetNewMode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementModeClusterChangeToModeParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementModeClusterChangeToModeParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass) Alloc() MTRDeviceEnergyManagementModeClusterChangeToModeParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementModeClusterChangeToModeParamsClass) New() MTRDeviceEnergyManagementModeClusterChangeToModeParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) Init() MTRDeviceEnergyManagementModeClusterChangeToModeParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) Autorelease() MTRDeviceEnergyManagementModeClusterChangeToModeParams {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementModeClusterChangeToModeParams creates a new MTRDeviceEnergyManagementModeClusterChangeToModeParams instance.
func NewMTRDeviceEnergyManagementModeClusterChangeToModeParams() MTRDeviceEnergyManagementModeClusterChangeToModeParams {
	return getMTRDeviceEnergyManagementModeClusterChangeToModeParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementModeClusterChangeToModeParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeParams
type MTRDeviceEnergyManagementModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementModeClusterChangeToModeParamsFrom constructs a [MTRDeviceEnergyManagementModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementModeClusterChangeToModeParams {
	return MTRDeviceEnergyManagementModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementModeClusterChangeToModeParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementModeClusterChangeToModeParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementModeClusterChangeToModeParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementModeClusterChangeToModeParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementModeClusterChangeToModeParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeParams/newMode
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) NewMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newMode"))
	return rv
}/* debug [instance_properties/getter]: newMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterChangeToModeParams/newMode
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) SetNewMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}/* debug [instance_properties/setter]: newMode */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementmodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementmodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementmodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementmodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRDeviceEnergyManagementModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementModeClusterChangeToModeParams */



