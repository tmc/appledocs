// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams */


/* debug [class_header]: Header for MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams */
// The class instance for the [MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams] class.
var (
	MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass     _MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass
	MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClassOnce sync.Once
)

func getMTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass() _MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass {
	MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClassOnce.Do(func() {
		MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass = _MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass{objc.GetClass("MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams")}
	})
	return MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass
}

type _MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams */
// An interface definition for the [MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams] class.
type IMTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass) Alloc() MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass) New() MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams) Init() MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams) Autorelease() MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams creates a new MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams instance.
func NewMTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams() MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams {
	return getMTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams
type MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams struct {
	objectivec.Object
}

// MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsFrom constructs a [MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams] from an unsafe.Pointer.
func MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsFrom(ptr unsafe.Pointer) MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams {
	return MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams/serverSideProcessingTimeout
func (m_ MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams/serverSideProcessingTimeout
func (m_ MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclustergetpendingdatasetrequestparams/timedinvoketimeoutms
func (m_ MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclustergetpendingdatasetrequestparams/timedinvoketimeoutms
func (m_ MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams */



