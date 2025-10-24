// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams */


/* debug [class_header]: Header for MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams */
// The class instance for the [MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams] class.
var (
	MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass     _MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass
	MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClassOnce sync.Once
)

func getMTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass() _MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass {
	MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClassOnce.Do(func() {
		MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass = _MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass{objc.GetClass("MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams")}
	})
	return MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass
}

type _MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams */
// An interface definition for the [MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams] class.
type IMTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams */
	// properties:
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	PendingDataset() foundation.Data
	SetPendingDataset(value foundation.Data)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass) Alloc() MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass) New() MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) Init() MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) Autorelease() MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams creates a new MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams instance.
func NewMTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams() MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams {
	return getMTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams
type MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams struct {
	objectivec.Object
}

// MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsFrom constructs a [MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams] from an unsafe.Pointer.
func MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsFrom(ptr unsafe.Pointer) MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams {
	return MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams */

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclustersetpendingdatasetrequestparams/pendingdataset
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) PendingDataset() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("pendingDataset"))
	return rv
}/* debug [instance_properties/getter]: pendingDataset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclustersetpendingdatasetrequestparams/pendingdataset
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) SetPendingDataset(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPendingDataset:"), value)
}/* debug [instance_properties/setter]: pendingDataset */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclustersetpendingdatasetrequestparams/serversideprocessingtimeout
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclustersetpendingdatasetrequestparams/serversideprocessingtimeout
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams */



