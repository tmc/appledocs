// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterLockDoorParams] class.
var (
	MTRDoorLockClusterLockDoorParamsClass     _MTRDoorLockClusterLockDoorParamsClass
	MTRDoorLockClusterLockDoorParamsClassOnce sync.Once
)

func getMTRDoorLockClusterLockDoorParamsClass() _MTRDoorLockClusterLockDoorParamsClass {
	MTRDoorLockClusterLockDoorParamsClassOnce.Do(func() {
		MTRDoorLockClusterLockDoorParamsClass = _MTRDoorLockClusterLockDoorParamsClass{objc.GetClass("MTRDoorLockClusterLockDoorParams")}
	})
	return MTRDoorLockClusterLockDoorParamsClass
}

type _MTRDoorLockClusterLockDoorParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterLockDoorParams] class.
type IMTRDoorLockClusterLockDoorParams interface {
	objectivec.IObject
	// properties:
	PinCode() objc.IObject /* cross-framework: Data */
	SetPinCode(value objc.IObject /* cross-framework: Data */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockDoorParams
type MTRDoorLockClusterLockDoorParams struct {
	objectivec.Object
}

// MTRDoorLockClusterLockDoorParamsFrom constructs a [MTRDoorLockClusterLockDoorParams] from an unsafe.Pointer.
func MTRDoorLockClusterLockDoorParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterLockDoorParams {
	return MTRDoorLockClusterLockDoorParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterLockDoorParamsClass) Alloc() MTRDoorLockClusterLockDoorParams {
	rv := objc.Send[MTRDoorLockClusterLockDoorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterLockDoorParamsClass) New() MTRDoorLockClusterLockDoorParams {
	rv := objc.Send[MTRDoorLockClusterLockDoorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterLockDoorParams) Init() MTRDoorLockClusterLockDoorParams {
	rv := objc.Send[MTRDoorLockClusterLockDoorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterLockDoorParams) Autorelease() MTRDoorLockClusterLockDoorParams {
	rv := objc.Send[MTRDoorLockClusterLockDoorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterLockDoorParams creates a new MTRDoorLockClusterLockDoorParams instance.
func NewMTRDoorLockClusterLockDoorParams() MTRDoorLockClusterLockDoorParams {
	return getMTRDoorLockClusterLockDoorParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockdoorparams/pincode
func (m_ MTRDoorLockClusterLockDoorParams) PinCode() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("pinCode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockdoorparams/pincode
func (m_ MTRDoorLockClusterLockDoorParams) SetPinCode(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPinCode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockdoorparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterLockDoorParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockdoorparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterLockDoorParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockdoorparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterLockDoorParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockdoorparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterLockDoorParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



