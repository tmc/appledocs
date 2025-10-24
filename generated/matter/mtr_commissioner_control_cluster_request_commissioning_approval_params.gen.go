// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRCommissionerControlClusterRequestCommissioningApprovalParams] class.
type IMTRCommissionerControlClusterRequestCommissioningApprovalParams interface {
	objectivec.IObject
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
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams
type MTRCommissionerControlClusterRequestCommissioningApprovalParams struct {
	objectivec.Object
}

// MTRCommissionerControlClusterRequestCommissioningApprovalParamsFrom constructs a [MTRCommissionerControlClusterRequestCommissioningApprovalParams] from an unsafe.Pointer.
func MTRCommissionerControlClusterRequestCommissioningApprovalParamsFrom(ptr unsafe.Pointer) MTRCommissionerControlClusterRequestCommissioningApprovalParams {
	return MTRCommissionerControlClusterRequestCommissioningApprovalParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRCommissionerControlClusterRequestCommissioningApprovalParamsClass) Alloc() MTRCommissionerControlClusterRequestCommissioningApprovalParams {
	rv := objc.Send[MTRCommissionerControlClusterRequestCommissioningApprovalParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/label
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/label
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/productID
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) ProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/productID
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetProductID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/requestID
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) RequestID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("requestID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/requestID
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetRequestID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestID:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/serverSideProcessingTimeout
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/serverSideProcessingTimeout
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/timedInvokeTimeoutMs
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/timedInvokeTimeoutMs
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/vendorID
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/vendorID
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}



