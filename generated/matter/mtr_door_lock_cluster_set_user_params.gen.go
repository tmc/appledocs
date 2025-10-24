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
	// properties:
	CredentialRule() objc.IObject /* cross-framework: NSNumber */
	SetCredentialRule(value objc.IObject /* cross-framework: NSNumber */)
	OperationType() objc.IObject /* cross-framework: NSNumber */
	SetOperationType(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/credentialrule
func (m_ MTRDoorLockClusterSetUserParams) CredentialRule() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("credentialRule"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/credentialrule
func (m_ MTRDoorLockClusterSetUserParams) SetCredentialRule(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialRule:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/operationtype
func (m_ MTRDoorLockClusterSetUserParams) OperationType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operationType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/operationtype
func (m_ MTRDoorLockClusterSetUserParams) SetOperationType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetUserParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetUserParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetUserParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetUserParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/userindex
func (m_ MTRDoorLockClusterSetUserParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/userindex
func (m_ MTRDoorLockClusterSetUserParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/username
func (m_ MTRDoorLockClusterSetUserParams) UserName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("userName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/username
func (m_ MTRDoorLockClusterSetUserParams) SetUserName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/userstatus
func (m_ MTRDoorLockClusterSetUserParams) UserStatus() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userStatus"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/userstatus
func (m_ MTRDoorLockClusterSetUserParams) SetUserStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/usertype
func (m_ MTRDoorLockClusterSetUserParams) UserType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/usertype
func (m_ MTRDoorLockClusterSetUserParams) SetUserType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/useruniqueid-22tka
func (m_ MTRDoorLockClusterSetUserParams) UserUniqueID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userUniqueID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/useruniqueid-22tka
func (m_ MTRDoorLockClusterSetUserParams) SetUserUniqueID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserUniqueID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/useruniqueid-22tje
func (m_ MTRDoorLockClusterSetUserParams) UserUniqueId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userUniqueId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetuserparams/useruniqueid-22tje
func (m_ MTRDoorLockClusterSetUserParams) SetUserUniqueId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserUniqueId:"), value)
}



