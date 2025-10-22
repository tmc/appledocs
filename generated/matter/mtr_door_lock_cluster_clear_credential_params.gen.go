// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterClearCredentialParams] class.
var (
	MTRDoorLockClusterClearCredentialParamsClass     _MTRDoorLockClusterClearCredentialParamsClass
	MTRDoorLockClusterClearCredentialParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearCredentialParamsClass() _MTRDoorLockClusterClearCredentialParamsClass {
	MTRDoorLockClusterClearCredentialParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearCredentialParamsClass = _MTRDoorLockClusterClearCredentialParamsClass{objc.GetClass("MTRDoorLockClusterClearCredentialParams")}
	})
	return MTRDoorLockClusterClearCredentialParamsClass
}

type _MTRDoorLockClusterClearCredentialParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterClearCredentialParams] class.
type IMTRDoorLockClusterClearCredentialParams interface {
	objectivec.IObject
	Credential() MTRDoorLockClusterCredentialStruct
	SetCredential(value IMTRDoorLockClusterCredentialStruct)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearCredentialParams
type MTRDoorLockClusterClearCredentialParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearCredentialParamsFrom constructs a [MTRDoorLockClusterClearCredentialParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearCredentialParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearCredentialParams {
	return MTRDoorLockClusterClearCredentialParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearCredentialParamsClass) Alloc() MTRDoorLockClusterClearCredentialParams {
	rv := objc.Send[MTRDoorLockClusterClearCredentialParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterClearCredentialParamsClass) New() MTRDoorLockClusterClearCredentialParams {
	rv := objc.Send[MTRDoorLockClusterClearCredentialParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearCredentialParams) Init() MTRDoorLockClusterClearCredentialParams {
	rv := objc.Send[MTRDoorLockClusterClearCredentialParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearCredentialParams) Autorelease() MTRDoorLockClusterClearCredentialParams {
	rv := objc.Send[MTRDoorLockClusterClearCredentialParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearCredentialParams creates a new MTRDoorLockClusterClearCredentialParams instance.
func NewMTRDoorLockClusterClearCredentialParams() MTRDoorLockClusterClearCredentialParams {
	return getMTRDoorLockClusterClearCredentialParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearcredentialparams/credential
func (m_ MTRDoorLockClusterClearCredentialParams) Credential() MTRDoorLockClusterCredentialStruct {
	rv := objc.Send[MTRDoorLockClusterCredentialStruct](m_.ID, objc.Sel("credential"))
	return rv
}


// SetCredential sets the value of the credential property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearcredentialparams/credential
func (m_ MTRDoorLockClusterClearCredentialParams) SetCredential(value IMTRDoorLockClusterCredentialStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredential:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearcredentialparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterClearCredentialParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearcredentialparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterClearCredentialParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearcredentialparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearCredentialParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearcredentialparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearCredentialParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



