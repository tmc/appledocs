// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterResumeRequestParams */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterResumeRequestParams */
// The class instance for the [MTRDeviceEnergyManagementClusterResumeRequestParams] class.
var (
	MTRDeviceEnergyManagementClusterResumeRequestParamsClass     _MTRDeviceEnergyManagementClusterResumeRequestParamsClass
	MTRDeviceEnergyManagementClusterResumeRequestParamsClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterResumeRequestParamsClass() _MTRDeviceEnergyManagementClusterResumeRequestParamsClass {
	MTRDeviceEnergyManagementClusterResumeRequestParamsClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterResumeRequestParamsClass = _MTRDeviceEnergyManagementClusterResumeRequestParamsClass{objc.GetClass("MTRDeviceEnergyManagementClusterResumeRequestParams")}
	})
	return MTRDeviceEnergyManagementClusterResumeRequestParamsClass
}

type _MTRDeviceEnergyManagementClusterResumeRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterResumeRequestParams */
// An interface definition for the [MTRDeviceEnergyManagementClusterResumeRequestParams] class.
type IMTRDeviceEnergyManagementClusterResumeRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterResumeRequestParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterResumeRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterResumeRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterResumeRequestParamsClass) Alloc() MTRDeviceEnergyManagementClusterResumeRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumeRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterResumeRequestParamsClass) New() MTRDeviceEnergyManagementClusterResumeRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumeRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterResumeRequestParams) Init() MTRDeviceEnergyManagementClusterResumeRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumeRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterResumeRequestParams) Autorelease() MTRDeviceEnergyManagementClusterResumeRequestParams {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumeRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterResumeRequestParams creates a new MTRDeviceEnergyManagementClusterResumeRequestParams instance.
func NewMTRDeviceEnergyManagementClusterResumeRequestParams() MTRDeviceEnergyManagementClusterResumeRequestParams {
	return getMTRDeviceEnergyManagementClusterResumeRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterResumeRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumeRequestParams
type MTRDeviceEnergyManagementClusterResumeRequestParams struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterResumeRequestParamsFrom constructs a [MTRDeviceEnergyManagementClusterResumeRequestParams] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterResumeRequestParamsFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterResumeRequestParams {
	return MTRDeviceEnergyManagementClusterResumeRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterResumeRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterResumeRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterResumeRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterResumeRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterResumeRequestParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumeRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterResumeRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumeRequestParams/serverSideProcessingTimeout
func (m_ MTRDeviceEnergyManagementClusterResumeRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumeRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterResumeRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumeRequestParams/timedInvokeTimeoutMs
func (m_ MTRDeviceEnergyManagementClusterResumeRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterResumeRequestParams */



