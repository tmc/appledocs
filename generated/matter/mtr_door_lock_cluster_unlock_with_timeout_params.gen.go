// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterUnlockWithTimeoutParams] class.
var (
	MTRDoorLockClusterUnlockWithTimeoutParamsClass     _MTRDoorLockClusterUnlockWithTimeoutParamsClass
	MTRDoorLockClusterUnlockWithTimeoutParamsClassOnce sync.Once
)

func getMTRDoorLockClusterUnlockWithTimeoutParamsClass() _MTRDoorLockClusterUnlockWithTimeoutParamsClass {
	MTRDoorLockClusterUnlockWithTimeoutParamsClassOnce.Do(func() {
		MTRDoorLockClusterUnlockWithTimeoutParamsClass = _MTRDoorLockClusterUnlockWithTimeoutParamsClass{objc.GetClass("MTRDoorLockClusterUnlockWithTimeoutParams")}
	})
	return MTRDoorLockClusterUnlockWithTimeoutParamsClass
}

type _MTRDoorLockClusterUnlockWithTimeoutParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterUnlockWithTimeoutParams] class.
type IMTRDoorLockClusterUnlockWithTimeoutParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockWithTimeoutParams
type MTRDoorLockClusterUnlockWithTimeoutParams struct {
	objectivec.Object
}

// MTRDoorLockClusterUnlockWithTimeoutParamsFrom constructs a [MTRDoorLockClusterUnlockWithTimeoutParams] from an unsafe.Pointer.
func MTRDoorLockClusterUnlockWithTimeoutParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterUnlockWithTimeoutParams {
	return MTRDoorLockClusterUnlockWithTimeoutParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterUnlockWithTimeoutParamsClass) Alloc() MTRDoorLockClusterUnlockWithTimeoutParams {
	rv := objc.Send[MTRDoorLockClusterUnlockWithTimeoutParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterUnlockWithTimeoutParamsClass) New() MTRDoorLockClusterUnlockWithTimeoutParams {
	rv := objc.Send[MTRDoorLockClusterUnlockWithTimeoutParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) Init() MTRDoorLockClusterUnlockWithTimeoutParams {
	rv := objc.Send[MTRDoorLockClusterUnlockWithTimeoutParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) Autorelease() MTRDoorLockClusterUnlockWithTimeoutParams {
	rv := objc.Send[MTRDoorLockClusterUnlockWithTimeoutParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterUnlockWithTimeoutParams creates a new MTRDoorLockClusterUnlockWithTimeoutParams instance.
func NewMTRDoorLockClusterUnlockWithTimeoutParams() MTRDoorLockClusterUnlockWithTimeoutParams {
	return getMTRDoorLockClusterUnlockWithTimeoutParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunlockwithtimeoutparams/pincode
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) PinCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pinCode"))
	return rv
}


// SetPinCode sets the value of the pinCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunlockwithtimeoutparams/pincode
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) SetPinCode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPinCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunlockwithtimeoutparams/timeout
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) Timeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timeout"))
	return rv
}


// SetTimeout sets the value of the timeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunlockwithtimeoutparams/timeout
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) SetTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunlockwithtimeoutparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunlockwithtimeoutparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunlockwithtimeoutparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunlockwithtimeoutparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterUnlockWithTimeoutParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}



