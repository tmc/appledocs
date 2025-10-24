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
	// properties:
	Credential() IMTRDoorLockClusterCredentialStruct
	SetCredential(value IMTRDoorLockClusterCredentialStruct)
	CredentialData() objc.IObject /* cross-framework: Data */
	SetCredentialData(value objc.IObject /* cross-framework: Data */)
	OperationType() objc.IObject /* cross-framework: NSNumber */
	SetOperationType(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	UserStatus() objc.IObject /* cross-framework: NSNumber */
	SetUserStatus(value objc.IObject /* cross-framework: NSNumber */)
	UserType() objc.IObject /* cross-framework: NSNumber */
	SetUserType(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/credential
func (m_ MTRDoorLockClusterSetCredentialParams) Credential() IMTRDoorLockClusterCredentialStruct {
	rv := objc.Send[MTRDoorLockClusterCredentialStruct](m_.ID, objc.Sel("credential"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/credential
func (m_ MTRDoorLockClusterSetCredentialParams) SetCredential(value IMTRDoorLockClusterCredentialStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredential:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/credentialdata
func (m_ MTRDoorLockClusterSetCredentialParams) CredentialData() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("credentialData"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/credentialdata
func (m_ MTRDoorLockClusterSetCredentialParams) SetCredentialData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/operationtype
func (m_ MTRDoorLockClusterSetCredentialParams) OperationType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operationType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/operationtype
func (m_ MTRDoorLockClusterSetCredentialParams) SetOperationType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetCredentialParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterSetCredentialParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetCredentialParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetCredentialParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/userindex
func (m_ MTRDoorLockClusterSetCredentialParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/userindex
func (m_ MTRDoorLockClusterSetCredentialParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/userstatus
func (m_ MTRDoorLockClusterSetCredentialParams) UserStatus() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userStatus"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/userstatus
func (m_ MTRDoorLockClusterSetCredentialParams) SetUserStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/usertype
func (m_ MTRDoorLockClusterSetCredentialParams) UserType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialparams/usertype
func (m_ MTRDoorLockClusterSetCredentialParams) SetUserType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserType:"), value)
}



