// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams */


/* debug [class_header]: Header for MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams */
// The class instance for the [MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams] class.
var (
	MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass     _MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass
	MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClassOnce sync.Once
)

func getMTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass() _MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass {
	MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClassOnce.Do(func() {
		MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass = _MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass{objc.GetClass("MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams")}
	})
	return MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass
}

type _MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams */
// An interface definition for the [MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams] class.
type IMTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams */
	// properties:
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	ActiveDataset() foundation.Data
	SetActiveDataset(value foundation.Data)
	Breadcrumb() objc.IObject /* cross-framework: NSNumber */
	SetBreadcrumb(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass) Alloc() MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass) New() MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) Init() MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) Autorelease() MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams creates a new MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams instance.
func NewMTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams() MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams {
	return getMTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams
type MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams struct {
	objectivec.Object
}

// MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsFrom constructs a [MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams] from an unsafe.Pointer.
func MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsFrom(ptr unsafe.Pointer) MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams {
	return MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams */

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclustersetactivedatasetrequestparams/activedataset
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) ActiveDataset() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("activeDataset"))
	return rv
}/* debug [instance_properties/getter]: activeDataset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclustersetactivedatasetrequestparams/activedataset
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) SetActiveDataset(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActiveDataset:"), value)
}/* debug [instance_properties/setter]: activeDataset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclustersetactivedatasetrequestparams/breadcrumb
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) Breadcrumb() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("breadcrumb"))
	return rv
}/* debug [instance_properties/getter]: breadcrumb */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclustersetactivedatasetrequestparams/breadcrumb
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) SetBreadcrumb(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBreadcrumb:"), value)
}/* debug [instance_properties/setter]: breadcrumb */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclustersetactivedatasetrequestparams/serversideprocessingtimeout
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclustersetactivedatasetrequestparams/serversideprocessingtimeout
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams */



