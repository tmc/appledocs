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
	// properties:
	CreatorFabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetCreatorFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	CredentialRule() objc.IObject /* cross-framework: NSNumber */
	SetCredentialRule(value objc.IObject /* cross-framework: NSNumber */)
	Credentials() unsafe.Pointer
	SetCredentials(value unsafe.Pointer)
	LastModifiedFabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetLastModifiedFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	NextUserIndex() objc.IObject /* cross-framework: NSNumber */
	SetNextUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	UserName() objc.IObject /* cross-framework: NSString */
	SetUserName(value objc.IObject /* cross-framework: NSString */)
	UserStatus() objc.IObject /* cross-framework: NSNumber */
	SetUserStatus(value objc.IObject /* cross-framework: NSNumber */)
	UserType() objc.IObject /* cross-framework: NSNumber */
	SetUserType(value objc.IObject /* cross-framework: NSNumber */)
	UserUniqueID() objc.IObject /* cross-framework: NSNumber */
	SetUserUniqueID(value objc.IObject /* cross-framework: NSNumber */)
	UserUniqueId() objc.IObject /* cross-framework: NSNumber */
	SetUserUniqueId(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/creatorfabricindex
func (m_ MTRDoorLockClusterGetUserResponseParams) CreatorFabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("creatorFabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/creatorfabricindex
func (m_ MTRDoorLockClusterGetUserResponseParams) SetCreatorFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCreatorFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/credentialrule
func (m_ MTRDoorLockClusterGetUserResponseParams) CredentialRule() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("credentialRule"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/credentialrule
func (m_ MTRDoorLockClusterGetUserResponseParams) SetCredentialRule(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialRule:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/credentials
func (m_ MTRDoorLockClusterGetUserResponseParams) Credentials() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("credentials"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/credentials
func (m_ MTRDoorLockClusterGetUserResponseParams) SetCredentials(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentials:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/lastmodifiedfabricindex
func (m_ MTRDoorLockClusterGetUserResponseParams) LastModifiedFabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lastModifiedFabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/lastmodifiedfabricindex
func (m_ MTRDoorLockClusterGetUserResponseParams) SetLastModifiedFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLastModifiedFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/nextuserindex
func (m_ MTRDoorLockClusterGetUserResponseParams) NextUserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nextUserIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/nextuserindex
func (m_ MTRDoorLockClusterGetUserResponseParams) SetNextUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextUserIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetUserResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetUserResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/userindex
func (m_ MTRDoorLockClusterGetUserResponseParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/userindex
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/username
func (m_ MTRDoorLockClusterGetUserResponseParams) UserName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("userName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/username
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/userstatus
func (m_ MTRDoorLockClusterGetUserResponseParams) UserStatus() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userStatus"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/userstatus
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/usertype
func (m_ MTRDoorLockClusterGetUserResponseParams) UserType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/usertype
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/useruniqueid-761ye
func (m_ MTRDoorLockClusterGetUserResponseParams) UserUniqueID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userUniqueID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/useruniqueid-761ye
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserUniqueID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserUniqueID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/useruniqueid-761xi
func (m_ MTRDoorLockClusterGetUserResponseParams) UserUniqueId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userUniqueId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserresponseparams/useruniqueid-761xi
func (m_ MTRDoorLockClusterGetUserResponseParams) SetUserUniqueId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserUniqueId:"), value)
}



