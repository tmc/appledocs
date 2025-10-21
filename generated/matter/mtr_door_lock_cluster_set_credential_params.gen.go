// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterSetCredentialParams] class.
var (
	MTRDoorLockClusterSetCredentialParamsClass     _MTRDoorLockClusterSetCredentialParamsClass
	MTRDoorLockClusterSetCredentialParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetCredentialParamsClass() _MTRDoorLockClusterSetCredentialParamsClass {
	MTRDoorLockClusterSetCredentialParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetCredentialParamsClass = _MTRDoorLockClusterSetCredentialParamsClass{objc.GetClass("MTRDoorLockClusterSetCredentialParams")}
	})
	return MTRDoorLockClusterSetCredentialParamsClass
}

type _MTRDoorLockClusterSetCredentialParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterSetCredentialParams] class.
type IMTRDoorLockClusterSetCredentialParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams
type MTRDoorLockClusterSetCredentialParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetCredentialParamsFrom constructs a [MTRDoorLockClusterSetCredentialParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetCredentialParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetCredentialParams {
	return MTRDoorLockClusterSetCredentialParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetCredentialParamsClass) Alloc() MTRDoorLockClusterSetCredentialParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterSetCredentialParamsClass) New() MTRDoorLockClusterSetCredentialParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetCredentialParams) Init() MTRDoorLockClusterSetCredentialParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetCredentialParams) Autorelease() MTRDoorLockClusterSetCredentialParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetCredentialParams creates a new MTRDoorLockClusterSetCredentialParams instance.
func NewMTRDoorLockClusterSetCredentialParams() MTRDoorLockClusterSetCredentialParams {
	return getMTRDoorLockClusterSetCredentialParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/credential
func (m_ MTRDoorLockClusterSetCredentialParams) Credential() MTRDoorLockClusterCredentialStruct {
	rv := objc.Send[MTRDoorLockClusterCredentialStruct](m_.ID, objc.Sel("credential"))
	return rv
}


// SetCredential sets the value of the credential property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/credential
func (m_ MTRDoorLockClusterSetCredentialParams) SetCredential(value IMTRDoorLockClusterCredentialStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredential:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/credentialdata
func (m_ MTRDoorLockClusterSetCredentialParams) CredentialData() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("credentialData"))
	return rv
}


// SetCredentialData sets the value of the credentialData property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/credentialdata
func (m_ MTRDoorLockClusterSetCredentialParams) SetCredentialData(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/operationtype
func (m_ MTRDoorLockClusterSetCredentialParams) OperationType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("operationType"))
	return rv
}


// SetOperationType sets the value of the operationType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/operationtype
func (m_ MTRDoorLockClusterSetCredentialParams) SetOperationType(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetCredentialParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetCredentialParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetCredentialParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetCredentialParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/userindex
func (m_ MTRDoorLockClusterSetCredentialParams) UserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userIndex"))
	return rv
}


// SetUserIndex sets the value of the userIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/userindex
func (m_ MTRDoorLockClusterSetCredentialParams) SetUserIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/userstatus
func (m_ MTRDoorLockClusterSetCredentialParams) UserStatus() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userStatus"))
	return rv
}


// SetUserStatus sets the value of the userStatus property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/userstatus
func (m_ MTRDoorLockClusterSetCredentialParams) SetUserStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/usertype
func (m_ MTRDoorLockClusterSetCredentialParams) UserType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userType"))
	return rv
}


// SetUserType sets the value of the userType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/usertype
func (m_ MTRDoorLockClusterSetCredentialParams) SetUserType(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserType:"), value)
}



