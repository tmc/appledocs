// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterPauseRequestParams */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterPauseRequestParams */
// The class instance for the [MTRDeviceEnergyManagementClusterPauseRequestParams] class.
var (
	MTRDeviceEnergyManagementClusterPauseRequestParamsClass     _MTRDeviceEnergyManagementClusterPauseRequestParamsClass
	MTRDeviceEnergyManagementClusterPauseRequestParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterPauseRequestParamsClass() _MTRDeviceEnergyManagementClusterPauseRequestParamsClass {
	MTRDeviceEnergyManagementClusterPauseRequestParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterPauseRequestParamsClass = _MTRDeviceEnergyManagementClusterPauseRequestParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterPauseRequestParams")}
	})
	return MTRDeviceEnergyManagementClusterPauseRequestParamsClass
}

type _MTRDeviceEnergyManagementClusterPauseRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterPauseRequestParams */
// An interface definition for the [MTRDeviceEnergyManagementClusterPauseRequestParams] class.
type IMTRDeviceEnergyManagementClusterPauseRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterPauseRequestParams */
	// properties:
	Cause() objc.IObject /* cross-framework: NSNumber */
	SetCause(value objc.IObject /* cross-framework: NSNumber */)
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterPauseRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterPauseRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterPauseRequestParamsClass) Alloc() MTRDeviceEnergyManagementClusterPauseRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPauseRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterPauseRequestParamsClass) New() MTRDeviceEnergyManagementClusterPauseRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPauseRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) Init() MTRDeviceEnergyManagementClusterPauseRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPauseRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) Autorelease() MTRDeviceEnergyManagementClusterPauseRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPauseRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterPauseRequestParams creates a new MTRDeviceEnergyManagementClusterPauseRequestParams instance.
func NewMTRDeviceEnergyManagementClusterPauseRequestParams() MTRDeviceEnergyManagementClusterPauseRequestParams {
	return getMTRDeviceEnergyManagementClusterPauseRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterPauseRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPauseRequestParams
type MTRDeviceEnergyManagementClusterPauseRequestParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterPauseRequestParamsFrom constructs a [MTRDeviceEnergyManagementClusterPauseRequestParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterPauseRequestParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterPauseRequestParams {
	return MTRDeviceEnergyManagementClusterPauseRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterPauseRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterPauseRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterPauseRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterPauseRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterPauseRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPauseRequestParams/cause
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) Cause() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cause"))
	return rv
}/* debug [instance_properties/getter]: cause */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPauseRequestParams/cause
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) SetCause(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}/* debug [instance_properties/setter]: cause */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpauserequestparams/duration
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpauserequestparams/duration
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpauserequestparams/serversideprocessingtimeout
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpauserequestparams/serversideprocessingtimeout
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpauserequestparams/timedinvoketimeoutms
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpauserequestparams/timedinvoketimeoutms
func (m_ MTRDeviceEnergyManagementClusterPauseRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterPauseRequestParams */



