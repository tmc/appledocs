// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRCommissionerControlClusterRequestCommissioningApprovalParams */


/* debug [class_header]: Header for MTRCommissionerControlClusterRequestCommissioningApprovalParams */
// The class instance for the [MTRCommissionerControlClusterRequestCommissioningApprovalParams] class.
var (
	MTRCommissionerControlClusterRequestCommissioningApprovalParamsClass     _MTRCommissionerControlClusterRequestCommissioningApprovalParamsClass
	MTRCommissionerControlClusterRequestCommissioningApprovalParamsClassOnce sync.Once
)

func getMTRCommissionerControlClusterRequestCommissioningApprovalParamsClass() _MTRCommissionerControlClusterRequestCommissioningApprovalParamsClass {
	MTRCommissionerControlClusterRequestCommissioningApprovalParamsClassOnce.Do(func() {
		MTRCommissionerControlClusterRequestCommissioningApprovalParamsClass = _MTRCommissionerControlClusterRequestCommissioningApprovalParamsClass{objc.GetClass("MTRCommissionerControlClusterRequestCommissioningApprovalParams")}
	})
	return MTRCommissionerControlClusterRequestCommissioningApprovalParamsClass
}

type _MTRCommissionerControlClusterRequestCommissioningApprovalParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRCommissionerControlClusterRequestCommissioningApprovalParams */
// An interface definition for the [MTRCommissionerControlClusterRequestCommissioningApprovalParams] class.
type IMTRCommissionerControlClusterRequestCommissioningApprovalParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRCommissionerControlClusterRequestCommissioningApprovalParams */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	ProductID() objc.IObject /* cross-framework: NSNumber */
	SetProductID(value objc.IObject /* cross-framework: NSNumber */)
	RequestID() objc.IObject /* cross-framework: NSNumber */
	SetRequestID(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	VendorID() objc.IObject /* cross-framework: NSNumber */
	SetVendorID(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRCommissionerControlClusterRequestCommissioningApprovalParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRCommissionerControlClusterRequestCommissioningApprovalParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRCommissionerControlClusterRequestCommissioningApprovalParamsClass) Alloc() MTRCommissionerControlClusterRequestCommissioningApprovalParams {
	rv := objc.Send[MTRCommissionerControlClusterRequestCommissioningApprovalParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRCommissionerControlClusterRequestCommissioningApprovalParamsClass) New() MTRCommissionerControlClusterRequestCommissioningApprovalParams {
	rv := objc.Send[MTRCommissionerControlClusterRequestCommissioningApprovalParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) Init() MTRCommissionerControlClusterRequestCommissioningApprovalParams {
	rv := objc.Send[MTRCommissionerControlClusterRequestCommissioningApprovalParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) Autorelease() MTRCommissionerControlClusterRequestCommissioningApprovalParams {
	rv := objc.Send[MTRCommissionerControlClusterRequestCommissioningApprovalParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommissionerControlClusterRequestCommissioningApprovalParams creates a new MTRCommissionerControlClusterRequestCommissioningApprovalParams instance.
func NewMTRCommissionerControlClusterRequestCommissioningApprovalParams() MTRCommissionerControlClusterRequestCommissioningApprovalParams {
	return getMTRCommissionerControlClusterRequestCommissioningApprovalParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRCommissionerControlClusterRequestCommissioningApprovalParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams
type MTRCommissionerControlClusterRequestCommissioningApprovalParams struct {
	objectivec.Object
}

// MTRCommissionerControlClusterRequestCommissioningApprovalParamsFrom constructs a [MTRCommissionerControlClusterRequestCommissioningApprovalParams] from an unsafe.Pointer.
func MTRCommissionerControlClusterRequestCommissioningApprovalParamsFrom(ptr unsafe.Pointer) MTRCommissionerControlClusterRequestCommissioningApprovalParams {
	return MTRCommissionerControlClusterRequestCommissioningApprovalParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRCommissionerControlClusterRequestCommissioningApprovalParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRCommissionerControlClusterRequestCommissioningApprovalParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRCommissionerControlClusterRequestCommissioningApprovalParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRCommissionerControlClusterRequestCommissioningApprovalParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRCommissionerControlClusterRequestCommissioningApprovalParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/label
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/label
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterrequestcommissioningapprovalparams/productid
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) ProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productID"))
	return rv
}/* debug [instance_properties/getter]: productID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterrequestcommissioningapprovalparams/productid
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetProductID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}/* debug [instance_properties/setter]: productID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterrequestcommissioningapprovalparams/requestid
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) RequestID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("requestID"))
	return rv
}/* debug [instance_properties/getter]: requestID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterrequestcommissioningapprovalparams/requestid
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetRequestID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestID:"), value)
}/* debug [instance_properties/setter]: requestID */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterrequestcommissioningapprovalparams/serversideprocessingtimeout
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterrequestcommissioningapprovalparams/serversideprocessingtimeout
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterrequestcommissioningapprovalparams/timedinvoketimeoutms
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterrequestcommissioningapprovalparams/timedinvoketimeoutms
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterrequestcommissioningapprovalparams/vendorid
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}/* debug [instance_properties/getter]: vendorID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterrequestcommissioningapprovalparams/vendorid
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}/* debug [instance_properties/setter]: vendorID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRCommissionerControlClusterRequestCommissioningApprovalParams */



