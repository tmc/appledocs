// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterSetUserParams] class.
var (
	MTRDoorLockClusterSetUserParamsClass     _MTRDoorLockClusterSetUserParamsClass
	MTRDoorLockClusterSetUserParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetUserParamsClass() _MTRDoorLockClusterSetUserParamsClass {
	MTRDoorLockClusterSetUserParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetUserParamsClass = _MTRDoorLockClusterSetUserParamsClass{objc.GetClass("MTRDoorLockClusterSetUserParams")}
	})
	return MTRDoorLockClusterSetUserParamsClass
}

type _MTRDoorLockClusterSetUserParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterSetUserParams] class.
type IMTRDoorLockClusterSetUserParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams
type MTRDoorLockClusterSetUserParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetUserParamsFrom constructs a [MTRDoorLockClusterSetUserParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetUserParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetUserParams {
	return MTRDoorLockClusterSetUserParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetUserParamsClass) Alloc() MTRDoorLockClusterSetUserParams {
	rv := objc.Send[MTRDoorLockClusterSetUserParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterSetUserParamsClass) New() MTRDoorLockClusterSetUserParams {
	rv := objc.Send[MTRDoorLockClusterSetUserParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetUserParams) Init() MTRDoorLockClusterSetUserParams {
	rv := objc.Send[MTRDoorLockClusterSetUserParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetUserParams) Autorelease() MTRDoorLockClusterSetUserParams {
	rv := objc.Send[MTRDoorLockClusterSetUserParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetUserParams creates a new MTRDoorLockClusterSetUserParams instance.
func NewMTRDoorLockClusterSetUserParams() MTRDoorLockClusterSetUserParams {
	return getMTRDoorLockClusterSetUserParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/credentialrule
func (m_ MTRDoorLockClusterSetUserParams) CredentialRule() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("credentialRule"))
	return rv
}


// SetCredentialRule sets the value of the credentialRule property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/credentialrule
func (m_ MTRDoorLockClusterSetUserParams) SetCredentialRule(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialRule:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/operationtype
func (m_ MTRDoorLockClusterSetUserParams) OperationType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("operationType"))
	return rv
}


// SetOperationType sets the value of the operationType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/operationtype
func (m_ MTRDoorLockClusterSetUserParams) SetOperationType(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetUserParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetUserParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetUserParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetUserParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/userindex
func (m_ MTRDoorLockClusterSetUserParams) UserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userIndex"))
	return rv
}


// SetUserIndex sets the value of the userIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/userindex
func (m_ MTRDoorLockClusterSetUserParams) SetUserIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/username
func (m_ MTRDoorLockClusterSetUserParams) UserName() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("userName"))
	return rv
}


// SetUserName sets the value of the userName property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/username
func (m_ MTRDoorLockClusterSetUserParams) SetUserName(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserName:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/userstatus
func (m_ MTRDoorLockClusterSetUserParams) UserStatus() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userStatus"))
	return rv
}


// SetUserStatus sets the value of the userStatus property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/userstatus
func (m_ MTRDoorLockClusterSetUserParams) SetUserStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/usertype
func (m_ MTRDoorLockClusterSetUserParams) UserType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userType"))
	return rv
}


// SetUserType sets the value of the userType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/usertype
func (m_ MTRDoorLockClusterSetUserParams) SetUserType(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/useruniqueid-22tka
func (m_ MTRDoorLockClusterSetUserParams) UserUniqueID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userUniqueID"))
	return rv
}


// SetUserUniqueID sets the value of the userUniqueID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/useruniqueid-22tka
func (m_ MTRDoorLockClusterSetUserParams) SetUserUniqueID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserUniqueID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/useruniqueid-22tje
func (m_ MTRDoorLockClusterSetUserParams) UserUniqueId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userUniqueId"))
	return rv
}


// SetUserUniqueId sets the value of the userUniqueId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/useruniqueid-22tje
func (m_ MTRDoorLockClusterSetUserParams) SetUserUniqueId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserUniqueId:"), value)
}



