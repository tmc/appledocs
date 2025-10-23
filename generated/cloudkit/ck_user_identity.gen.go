// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKUserIdentity] class.
var (
	CKUserIdentityClass     _CKUserIdentityClass
	CKUserIdentityClassOnce sync.Once
)

func getCKUserIdentityClass() _CKUserIdentityClass {
	CKUserIdentityClassOnce.Do(func() {
		CKUserIdentityClass = _CKUserIdentityClass{objc.GetClass("CKUserIdentity")}
	})
	return CKUserIdentityClass
}

type _CKUserIdentityClass struct {
	class objc.Class
}

// An interface definition for the [CKUserIdentity] class.
type ICKUserIdentity interface {
	objectivec.IObject
	UserRecordID() CKRecordID
	ContactIdentifiers() string
	SetContactIdentifiers(value string)
	HasiCloudAccount() bool
	SetHasiCloudAccount(value bool)
	LookupInfo() ICKUserIdentityLookupInfo
	SetLookupInfo(value ICKUserIdentityLookupInfo)
	NameComponents() foundation.PersonNameComponents
	SetNameComponents(value foundation.PersonNameComponents)
}

// The identity of a user.
//
// A user identity provides identifiable data about an iCloud user, including their name, user record ID, and an email address or phone number. CloudKit retrieves this information from the user’s iCloud account. A user must give their consent to be discoverable before CloudKit can provide this data to your app. For more information, see . You don’t create instances of this class. Instead, CloudKit provides them in certain contexts. A share’s owner has a user identity, as does each of its participants. When creating participants, CloudKit tries to find iCloud accounts it can use to populate their identities. If CloudKit doesn’t find an account, it sets the identity’s property to . You can also discover the identities of your app’s users by executing one of the user discovery operations: and . Identities that CloudKit discovers using correspond to entries in the device’s Contacts database. These identities contain the identifiers of their Contact records, which you can use to fetch those records from the Contacts database. For more information, see .


// The identity of a user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity
type CKUserIdentity struct {
	objectivec.Object
}

// CKUserIdentityFrom constructs a [CKUserIdentity] from an unsafe.Pointer.
//
// The identity of a user.
func CKUserIdentityFrom(ptr unsafe.Pointer) CKUserIdentity {
	return CKUserIdentity{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKUserIdentityClass) Alloc() CKUserIdentity {
	rv := objc.Send[CKUserIdentity](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKUserIdentityClass) New() CKUserIdentity {
	rv := objc.Send[CKUserIdentity](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKUserIdentity) Init() CKUserIdentity {
	rv := objc.Send[CKUserIdentity](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKUserIdentity) Autorelease() CKUserIdentity {
	rv := objc.Send[CKUserIdentity](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKUserIdentity creates a new CKUserIdentity instance.
func NewCKUserIdentity() CKUserIdentity {
	return getCKUserIdentityClass().New()
}



// The user record ID for the corresponding user record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/userRecordID
func (c_ CKUserIdentity) UserRecordID() CKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("userRecordID"))
	return rv
}


// Identifiers that match contacts in the local Contacts database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/contactidentifiers
func (c_ CKUserIdentity) ContactIdentifiers() string {
	rv := objc.Send[string](c_.ID, objc.Sel("contactIdentifiers"))
	return rv
}


// Identifiers that match contacts in the local Contacts database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/contactidentifiers
func (c_ CKUserIdentity) SetContactIdentifiers(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactIdentifiers:"), objc.String(value))
}


// A Boolean value that indicates whether the user has an iCloud account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/hasicloudaccount
func (c_ CKUserIdentity) HasiCloudAccount() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasiCloudAccount"))
	return rv
}


// A Boolean value that indicates whether the user has an iCloud account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/hasicloudaccount
func (c_ CKUserIdentity) SetHasiCloudAccount(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasiCloudAccount:"), value)
}


// The lookup info for retrieving the user identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/lookupinfo-swift.property
func (c_ CKUserIdentity) LookupInfo() ICKUserIdentityLookupInfo {
	rv := objc.Send[CKUserIdentityLookupInfo](c_.ID, objc.Sel("lookupInfo"))
	return rv
}


// The lookup info for retrieving the user identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/lookupinfo-swift.property
func (c_ CKUserIdentity) SetLookupInfo(value ICKUserIdentityLookupInfo) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLookupInfo:"), value)
}


// The user’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/namecomponents
func (c_ CKUserIdentity) NameComponents() foundation.PersonNameComponents {
	rv := objc.Send[foundation.PersonNameComponents](c_.ID, objc.Sel("nameComponents"))
	return rv
}


// The user’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckuseridentity/namecomponents
func (c_ CKUserIdentity) SetNameComponents(value foundation.PersonNameComponents) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNameComponents:"), value)
}



