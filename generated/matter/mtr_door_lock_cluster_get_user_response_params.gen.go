// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterGetUserResponseParams] class.
var (
	MTRDoorLockClusterGetUserResponseParamsClass     _MTRDoorLockClusterGetUserResponseParamsClass
	MTRDoorLockClusterGetUserResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetUserResponseParamsClass() _MTRDoorLockClusterGetUserResponseParamsClass {
	MTRDoorLockClusterGetUserResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetUserResponseParamsClass = _MTRDoorLockClusterGetUserResponseParamsClass{objc.GetClass("MTRDoorLockClusterGetUserResponseParams")}
	})
	return MTRDoorLockClusterGetUserResponseParamsClass
}

type _MTRDoorLockClusterGetUserResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetUserResponseParams] class.
type IMTRDoorLockClusterGetUserResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams
type MTRDoorLockClusterGetUserResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetUserResponseParamsFrom constructs a [MTRDoorLockClusterGetUserResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetUserResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetUserResponseParams {
	return MTRDoorLockClusterGetUserResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetUserResponseParamsClass) Alloc() MTRDoorLockClusterGetUserResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetUserResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetUserResponseParamsClass) New() MTRDoorLockClusterGetUserResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetUserResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetUserResponseParams) Init() MTRDoorLockClusterGetUserResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetUserResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetUserResponseParams) Autorelease() MTRDoorLockClusterGetUserResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetUserResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetUserResponseParams creates a new MTRDoorLockClusterGetUserResponseParams instance.
func NewMTRDoorLockClusterGetUserResponseParams() MTRDoorLockClusterGetUserResponseParams {
	return getMTRDoorLockClusterGetUserResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetUserResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetUserResponseParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/creatorfabricindex
func (m_ MTRDoorLockClusterGetUserResponseParams) CreatorFabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("creatorFabricIndex"))
	return rv
}


// SetCreatorFabricIndex sets the value of the creatorFabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/creatorfabricindex
func (m_ MTRDoorLockClusterGetUserResponseParams) SetCreatorFabricIndex(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCreatorFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/useruniqueid-761ye
func (m_ MTRDoorLockClusterGetUserResponseParams) UserUniqueID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userUniqueID"))
	return rv
}


// SetUserUniqueID sets the value of the userUniqueID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/useruniqueid-761ye
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserUniqueID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserUniqueID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/credentialrule
func (m_ MTRDoorLockClusterGetUserResponseParams) CredentialRule() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("credentialRule"))
	return rv
}


// SetCredentialRule sets the value of the credentialRule property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/credentialrule
func (m_ MTRDoorLockClusterGetUserResponseParams) SetCredentialRule(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialRule:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/userstatus
func (m_ MTRDoorLockClusterGetUserResponseParams) UserStatus() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userStatus"))
	return rv
}


// SetUserStatus sets the value of the userStatus property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/userstatus
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserStatus(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/username
func (m_ MTRDoorLockClusterGetUserResponseParams) UserName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("userName"))
	return rv
}


// SetUserName sets the value of the userName property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/username
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/nextuserindex
func (m_ MTRDoorLockClusterGetUserResponseParams) NextUserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nextUserIndex"))
	return rv
}


// SetNextUserIndex sets the value of the nextUserIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/nextuserindex
func (m_ MTRDoorLockClusterGetUserResponseParams) SetNextUserIndex(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextUserIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/usertype
func (m_ MTRDoorLockClusterGetUserResponseParams) UserType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userType"))
	return rv
}


// SetUserType sets the value of the userType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/usertype
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserType(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/useruniqueid-761xi
func (m_ MTRDoorLockClusterGetUserResponseParams) UserUniqueId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userUniqueId"))
	return rv
}


// SetUserUniqueId sets the value of the userUniqueId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/useruniqueid-761xi
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserUniqueId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserUniqueId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/lastmodifiedfabricindex
func (m_ MTRDoorLockClusterGetUserResponseParams) LastModifiedFabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lastModifiedFabricIndex"))
	return rv
}


// SetLastModifiedFabricIndex sets the value of the lastModifiedFabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/lastmodifiedfabricindex
func (m_ MTRDoorLockClusterGetUserResponseParams) SetLastModifiedFabricIndex(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLastModifiedFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/credentials
func (m_ MTRDoorLockClusterGetUserResponseParams) Credentials() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("credentials"))
	return rv
}


// SetCredentials sets the value of the credentials property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/credentials
func (m_ MTRDoorLockClusterGetUserResponseParams) SetCredentials(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentials:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/userindex
func (m_ MTRDoorLockClusterGetUserResponseParams) UserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userIndex"))
	return rv
}


// SetUserIndex sets the value of the userIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/userindex
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserIndex(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}



