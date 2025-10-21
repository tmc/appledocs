// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKUserIdentityLookupInfo] class.
var (
	CKUserIdentityLookupInfoClass     _CKUserIdentityLookupInfoClass
	CKUserIdentityLookupInfoClassOnce sync.Once
)

func getCKUserIdentityLookupInfoClass() _CKUserIdentityLookupInfoClass {
	CKUserIdentityLookupInfoClassOnce.Do(func() {
		CKUserIdentityLookupInfoClass = _CKUserIdentityLookupInfoClass{objc.GetClass("CKUserIdentityLookupInfo")}
	})
	return CKUserIdentityLookupInfoClass
}

type _CKUserIdentityLookupInfoClass struct {
	class objc.Class
}

// An interface definition for the [CKUserIdentityLookupInfo] class.
type ICKUserIdentityLookupInfo interface {
	objectivec.IObject
}

// The criteria to use when searching for discoverable iCloud users.
//
// Use this object when you want to discover the identities of your app’s users with , or to create a share’s participants with . You create individual instances by providing an email address, phone number, or user record ID. Alternatively, create an array of objects all at once by using one of the convenience methods, such as .
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class
type CKUserIdentityLookupInfo struct {
	objectivec.Object
}

// CKUserIdentityLookupInfoFrom constructs a [CKUserIdentityLookupInfo] from an unsafe.Pointer.
//
// The criteria to use when searching for discoverable iCloud users.
func CKUserIdentityLookupInfoFrom(ptr unsafe.Pointer) CKUserIdentityLookupInfo {
	return CKUserIdentityLookupInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKUserIdentityLookupInfoClass) Alloc() CKUserIdentityLookupInfo {
	rv := objc.Send[CKUserIdentityLookupInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKUserIdentityLookupInfoClass) New() CKUserIdentityLookupInfo {
	rv := objc.Send[CKUserIdentityLookupInfo](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKUserIdentityLookupInfo) Init() CKUserIdentityLookupInfo {
	rv := objc.Send[CKUserIdentityLookupInfo](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKUserIdentityLookupInfo) Autorelease() CKUserIdentityLookupInfo {
	rv := objc.Send[CKUserIdentityLookupInfo](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKUserIdentityLookupInfo creates a new CKUserIdentityLookupInfo instance.
func NewCKUserIdentityLookupInfo() CKUserIdentityLookupInfo {
	return getCKUserIdentityLookupInfoClass().New()
}




// Creates a lookup info for the specified email address.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/init(emailAddress:)
func NewCKUserIdentityLookupInfoWithEmailAddress(emailAddress string) CKUserIdentityLookupInfo {
	instance := getCKUserIdentityLookupInfoClass().Alloc()
	rv := objc.Send[CKUserIdentityLookupInfo](instance.ID, objc.Sel("initWithEmailAddress:"), objc.String(emailAddress))
	rv.Autorelease()
	return rv
}



// Creates a lookup info for the specified phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/init(phoneNumber:)
func NewCKUserIdentityLookupInfoWithPhoneNumber(phoneNumber string) CKUserIdentityLookupInfo {
	instance := getCKUserIdentityLookupInfoClass().Alloc()
	rv := objc.Send[CKUserIdentityLookupInfo](instance.ID, objc.Sel("initWithPhoneNumber:"), objc.String(phoneNumber))
	rv.Autorelease()
	return rv
}



// Creates a lookup info for the specified user record ID.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/init(userRecordID:)
func NewCKUserIdentityLookupInfoWithUserRecordID(userRecordID unsafe.Pointer) CKUserIdentityLookupInfo {
	instance := getCKUserIdentityLookupInfoClass().Alloc()
	rv := objc.Send[CKUserIdentityLookupInfo](instance.ID, objc.Sel("initWithUserRecordID:"), userRecordID)
	rv.Autorelease()
	return rv
}


// Returns an array of lookup infos for the specifed user record IDs.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/lookupInfos(with:)
func (cc _CKUserIdentityLookupInfoClass) LookupInfosWithRecordIDs(recordIDs unsafe.Pointer) []CKUserIdentityLookupInfo {
	rv := objc.Send[[]CKUserIdentityLookupInfo](objc.ID(cc.class), objc.Sel("lookupInfosWithRecordIDs:"), recordIDs)
	return rv
}

// Returns an array of lookup infos for the specifed email addresses.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/lookupInfos(withEmails:)
func (cc _CKUserIdentityLookupInfoClass) LookupInfosWithEmails(emails unsafe.Pointer) []CKUserIdentityLookupInfo {
	rv := objc.Send[[]CKUserIdentityLookupInfo](objc.ID(cc.class), objc.Sel("lookupInfosWithEmails:"), emails)
	return rv
}

// Returns an array of lookup infos for the specifed phone numbers.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/lookupInfos(withPhoneNumbers:)
func (cc _CKUserIdentityLookupInfoClass) LookupInfosWithPhoneNumbers(phoneNumbers unsafe.Pointer) []CKUserIdentityLookupInfo {
	rv := objc.Send[[]CKUserIdentityLookupInfo](objc.ID(cc.class), objc.Sel("lookupInfosWithPhoneNumbers:"), phoneNumbers)
	return rv
}

// The user’s email address.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/emailAddress
func (c_ CKUserIdentityLookupInfo) EmailAddress() string {
	rv := objc.Send[string](c_.ID, objc.Sel("emailAddress"))
	return rv
}

// The user’s phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/phoneNumber
func (c_ CKUserIdentityLookupInfo) PhoneNumber() string {
	rv := objc.Send[string](c_.ID, objc.Sel("phoneNumber"))
	return rv
}

// The ID of the user record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/userRecordID
func (c_ CKUserIdentityLookupInfo) UserRecordID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("userRecordID"))
	return rv
}


