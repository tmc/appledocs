// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRICDManagementClusterStayActiveRequestParams */


/* debug [class_header]: Header for MTRICDManagementClusterStayActiveRequestParams */
// The class instance for the [MTRICDManagementClusterStayActiveRequestParams] class.
var (
	MTRICDManagementClusterStayActiveRequestParamsClass     _MTRICDManagementClusterStayActiveRequestParamsClass
	MTRICDManagementClusterStayActiveRequestParamsClassOnce sync.Once
)

func getMTRICDManagementClusterStayActiveRequestParamsClass() _MTRICDManagementClusterStayActiveRequestParamsClass {
	MTRICDManagementClusterStayActiveRequestParamsClassOnce.Do(func() {
		MTRICDManagementClusterStayActiveRequestParamsClass = _MTRICDManagementClusterStayActiveRequestParamsClass{objc.GetClass("MTRICDManagementClusterStayActiveRequestParams")}
	})
	return MTRICDManagementClusterStayActiveRequestParamsClass
}

type _MTRICDManagementClusterStayActiveRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRICDManagementClusterStayActiveRequestParams */
// An interface definition for the [MTRICDManagementClusterStayActiveRequestParams] class.
type IMTRICDManagementClusterStayActiveRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRICDManagementClusterStayActiveRequestParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StayActiveDuration() objc.IObject /* cross-framework: NSNumber */
	SetStayActiveDuration(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRICDManagementClusterStayActiveRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRICDManagementClusterStayActiveRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRICDManagementClusterStayActiveRequestParamsClass) Alloc() MTRICDManagementClusterStayActiveRequestParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRICDManagementClusterStayActiveRequestParamsClass) New() MTRICDManagementClusterStayActiveRequestParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRICDManagementClusterStayActiveRequestParams) Init() MTRICDManagementClusterStayActiveRequestParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRICDManagementClusterStayActiveRequestParams) Autorelease() MTRICDManagementClusterStayActiveRequestParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRICDManagementClusterStayActiveRequestParams creates a new MTRICDManagementClusterStayActiveRequestParams instance.
func NewMTRICDManagementClusterStayActiveRequestParams() MTRICDManagementClusterStayActiveRequestParams {
	return getMTRICDManagementClusterStayActiveRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRICDManagementClusterStayActiveRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveRequestParams
type MTRICDManagementClusterStayActiveRequestParams struct {
	objectivec.Object
}

// MTRICDManagementClusterStayActiveRequestParamsFrom constructs a [MTRICDManagementClusterStayActiveRequestParams] from an unsafe.Pointer.
func MTRICDManagementClusterStayActiveRequestParamsFrom(ptr unsafe.Pointer) MTRICDManagementClusterStayActiveRequestParams {
	return MTRICDManagementClusterStayActiveRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRICDManagementClusterStayActiveRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRICDManagementClusterStayActiveRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRICDManagementClusterStayActiveRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRICDManagementClusterStayActiveRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRICDManagementClusterStayActiveRequestParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveRequestParams/serverSideProcessingTimeout
func (m_ MTRICDManagementClusterStayActiveRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveRequestParams/serverSideProcessingTimeout
func (m_ MTRICDManagementClusterStayActiveRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterstayactiverequestparams/stayactiveduration
func (m_ MTRICDManagementClusterStayActiveRequestParams) StayActiveDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stayActiveDuration"))
	return rv
}/* debug [instance_properties/getter]: stayActiveDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterstayactiverequestparams/stayactiveduration
func (m_ MTRICDManagementClusterStayActiveRequestParams) SetStayActiveDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStayActiveDuration:"), value)
}/* debug [instance_properties/setter]: stayActiveDuration */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterstayactiverequestparams/timedinvoketimeoutms
func (m_ MTRICDManagementClusterStayActiveRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterstayactiverequestparams/timedinvoketimeoutms
func (m_ MTRICDManagementClusterStayActiveRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRICDManagementClusterStayActiveRequestParams */



