// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/label
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/label
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/productID
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) ProductID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("productID"))
	return rv
}


// SetProductID sets the value of the productID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/productID
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetProductID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/requestID
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) RequestID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("requestID"))
	return rv
}


// SetRequestID sets the value of the requestID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/requestID
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetRequestID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestID:"), value)
}
// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/serverSideProcessingTimeout
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/serverSideProcessingTimeout
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/timedInvokeTimeoutMs
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/timedInvokeTimeoutMs
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/vendorID
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) VendorID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("vendorID"))
	return rv
}


// SetVendorID sets the value of the vendorID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterRequestCommissioningApprovalParams/vendorID
func (m_ MTRCommissionerControlClusterRequestCommissioningApprovalParams) SetVendorID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}


