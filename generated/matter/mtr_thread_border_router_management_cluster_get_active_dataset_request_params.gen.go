// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams */


/* debug [class_header]: Header for MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams */
// The class instance for the [MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams] class.
var (
	MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass     _MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass
	MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClassOnce sync.Once
)

func getMTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass() _MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass {
	MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClassOnce.Do(func() {
		MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass = _MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass{objc.GetClass("MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams")}
	})
	return MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass
}

type _MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams */
// An interface definition for the [MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams] class.
type IMTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams */
	// properties:
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass) Alloc() MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass) New() MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams) Init() MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams) Autorelease() MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams creates a new MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams instance.
func NewMTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams() MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams {
	return getMTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams
type MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams struct {
	objectivec.Object
}

// MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsFrom constructs a [MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams] from an unsafe.Pointer.
func MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsFrom(ptr unsafe.Pointer) MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams {
	return MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams */

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclustergetactivedatasetrequestparams/serversideprocessingtimeout
func (m_ MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadborderroutermanagementclustergetactivedatasetrequestparams/serversideprocessingtimeout
func (m_ MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams */



