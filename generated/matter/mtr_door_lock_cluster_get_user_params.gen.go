// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterGetUserParams] class.
var (
	MTRDoorLockClusterGetUserParamsClass     _MTRDoorLockClusterGetUserParamsClass
	MTRDoorLockClusterGetUserParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetUserParamsClass() _MTRDoorLockClusterGetUserParamsClass {
	MTRDoorLockClusterGetUserParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetUserParamsClass = _MTRDoorLockClusterGetUserParamsClass{objc.GetClass("MTRDoorLockClusterGetUserParams")}
	})
	return MTRDoorLockClusterGetUserParamsClass
}

type _MTRDoorLockClusterGetUserParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetUserParams] class.
type IMTRDoorLockClusterGetUserParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserParams
type MTRDoorLockClusterGetUserParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetUserParamsFrom constructs a [MTRDoorLockClusterGetUserParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetUserParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetUserParams {
	return MTRDoorLockClusterGetUserParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetUserParamsClass) Alloc() MTRDoorLockClusterGetUserParams {
	rv := objc.Send[MTRDoorLockClusterGetUserParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetUserParamsClass) New() MTRDoorLockClusterGetUserParams {
	rv := objc.Send[MTRDoorLockClusterGetUserParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetUserParams) Init() MTRDoorLockClusterGetUserParams {
	rv := objc.Send[MTRDoorLockClusterGetUserParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetUserParams) Autorelease() MTRDoorLockClusterGetUserParams {
	rv := objc.Send[MTRDoorLockClusterGetUserParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetUserParams creates a new MTRDoorLockClusterGetUserParams instance.
func NewMTRDoorLockClusterGetUserParams() MTRDoorLockClusterGetUserParams {
	return getMTRDoorLockClusterGetUserParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterGetUserParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterGetUserParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetUserParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetUserParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserparams/userindex
func (m_ MTRDoorLockClusterGetUserParams) UserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userIndex"))
	return rv
}


// SetUserIndex sets the value of the userIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserparams/userindex
func (m_ MTRDoorLockClusterGetUserParams) SetUserIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}



