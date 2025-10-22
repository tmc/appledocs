// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterClearUserParams] class.
var (
	MTRDoorLockClusterClearUserParamsClass     _MTRDoorLockClusterClearUserParamsClass
	MTRDoorLockClusterClearUserParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearUserParamsClass() _MTRDoorLockClusterClearUserParamsClass {
	MTRDoorLockClusterClearUserParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearUserParamsClass = _MTRDoorLockClusterClearUserParamsClass{objc.GetClass("MTRDoorLockClusterClearUserParams")}
	})
	return MTRDoorLockClusterClearUserParamsClass
}

type _MTRDoorLockClusterClearUserParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterClearUserParams] class.
type IMTRDoorLockClusterClearUserParams interface {
	objectivec.IObject
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
	UserIndex() foundation.Number
	SetUserIndex(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearUserParams
type MTRDoorLockClusterClearUserParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearUserParamsFrom constructs a [MTRDoorLockClusterClearUserParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearUserParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearUserParams {
	return MTRDoorLockClusterClearUserParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearUserParamsClass) Alloc() MTRDoorLockClusterClearUserParams {
	rv := objc.Send[MTRDoorLockClusterClearUserParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterClearUserParamsClass) New() MTRDoorLockClusterClearUserParams {
	rv := objc.Send[MTRDoorLockClusterClearUserParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearUserParams) Init() MTRDoorLockClusterClearUserParams {
	rv := objc.Send[MTRDoorLockClusterClearUserParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearUserParams) Autorelease() MTRDoorLockClusterClearUserParams {
	rv := objc.Send[MTRDoorLockClusterClearUserParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearUserParams creates a new MTRDoorLockClusterClearUserParams instance.
func NewMTRDoorLockClusterClearUserParams() MTRDoorLockClusterClearUserParams {
	return getMTRDoorLockClusterClearUserParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearuserparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterClearUserParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearuserparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterClearUserParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearuserparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearUserParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearuserparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearUserParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearuserparams/userindex
func (m_ MTRDoorLockClusterClearUserParams) UserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userIndex"))
	return rv
}


// SetUserIndex sets the value of the userIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearuserparams/userindex
func (m_ MTRDoorLockClusterClearUserParams) SetUserIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}



