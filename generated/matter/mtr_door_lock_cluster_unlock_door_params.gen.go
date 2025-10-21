// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterUnlockDoorParams] class.
var (
	MTRDoorLockClusterUnlockDoorParamsClass     _MTRDoorLockClusterUnlockDoorParamsClass
	MTRDoorLockClusterUnlockDoorParamsClassOnce sync.Once
)

func getMTRDoorLockClusterUnlockDoorParamsClass() _MTRDoorLockClusterUnlockDoorParamsClass {
	MTRDoorLockClusterUnlockDoorParamsClassOnce.Do(func() {
		MTRDoorLockClusterUnlockDoorParamsClass = _MTRDoorLockClusterUnlockDoorParamsClass{objc.GetClass("MTRDoorLockClusterUnlockDoorParams")}
	})
	return MTRDoorLockClusterUnlockDoorParamsClass
}

type _MTRDoorLockClusterUnlockDoorParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterUnlockDoorParams] class.
type IMTRDoorLockClusterUnlockDoorParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockDoorParams
type MTRDoorLockClusterUnlockDoorParams struct {
	objectivec.Object
}

// MTRDoorLockClusterUnlockDoorParamsFrom constructs a [MTRDoorLockClusterUnlockDoorParams] from an unsafe.Pointer.
func MTRDoorLockClusterUnlockDoorParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterUnlockDoorParams {
	return MTRDoorLockClusterUnlockDoorParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterUnlockDoorParamsClass) Alloc() MTRDoorLockClusterUnlockDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnlockDoorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterUnlockDoorParamsClass) New() MTRDoorLockClusterUnlockDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnlockDoorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterUnlockDoorParams) Init() MTRDoorLockClusterUnlockDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnlockDoorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterUnlockDoorParams) Autorelease() MTRDoorLockClusterUnlockDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnlockDoorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterUnlockDoorParams creates a new MTRDoorLockClusterUnlockDoorParams instance.
func NewMTRDoorLockClusterUnlockDoorParams() MTRDoorLockClusterUnlockDoorParams {
	return getMTRDoorLockClusterUnlockDoorParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunlockdoorparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterUnlockDoorParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunlockdoorparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterUnlockDoorParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunlockdoorparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterUnlockDoorParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunlockdoorparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterUnlockDoorParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunlockdoorparams/pincode
func (m_ MTRDoorLockClusterUnlockDoorParams) PinCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pinCode"))
	return rv
}


// SetPinCode sets the value of the pinCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterunlockdoorparams/pincode
func (m_ MTRDoorLockClusterUnlockDoorParams) SetPinCode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPinCode:"), value)
}



