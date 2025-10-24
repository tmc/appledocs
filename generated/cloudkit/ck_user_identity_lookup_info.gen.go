// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKUserIdentityLookupInfo */


/* debug [class_header]: Header for CKUserIdentityLookupInfo */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKUserIdentityLookupInfo */
// An interface definition for the [CKUserIdentityLookupInfo] class.
type ICKUserIdentityLookupInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKUserIdentityLookupInfo */
	// properties:
	EmailAddress() objc.IObject /* cross-framework: NSString */
	PhoneNumber() objc.IObject /* cross-framework: NSString */
	UserRecordID() ICKRecordID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKUserIdentityLookupInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKUserIdentityLookupInfo */
// Alloc allocates a new instance without initialization.
func (cc _CKUserIdentityLookupInfoClass) Alloc() CKUserIdentityLookupInfo {
	rv := objc.Send[CKUserIdentityLookupInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKUserIdentityLookupInfo */
// The criteria to use when searching for discoverable iCloud users.
//
// Use this object when you want to discover the identities of your app’s users with , or to create a share’s participants with . You create individual instances by providing an email address, phone number, or user record ID. Alternatively, create an array of objects all at once by using one of the convenience methods, such as .


// The criteria to use when searching for discoverable iCloud users.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKUserIdentityLookupInfo */

// Creates a lookup info for the specified email address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/init(emailAddress:)
func NewCKUserIdentityLookupInfoWithEmailAddress(emailAddress objc.IObject /* cross-framework: NSString */) CKUserIdentityLookupInfo {
	instance := getCKUserIdentityLookupInfoClass().Alloc()
	rv := objc.Send[CKUserIdentityLookupInfo](instance.ID, objc.Sel("initWithEmailAddress:"), emailAddress)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKUserIdentityLookupInfoWithEmailAddress */


// Creates a lookup info for the specified phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/init(phoneNumber:)
func NewCKUserIdentityLookupInfoWithPhoneNumber(phoneNumber objc.IObject /* cross-framework: NSString */) CKUserIdentityLookupInfo {
	instance := getCKUserIdentityLookupInfoClass().Alloc()
	rv := objc.Send[CKUserIdentityLookupInfo](instance.ID, objc.Sel("initWithPhoneNumber:"), phoneNumber)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKUserIdentityLookupInfoWithPhoneNumber */


// Creates a lookup info for the specified user record ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/init(userRecordID:)
func NewCKUserIdentityLookupInfoWithUserRecordID(userRecordID ICKRecordID) CKUserIdentityLookupInfo {
	instance := getCKUserIdentityLookupInfoClass().Alloc()
	rv := objc.Send[CKUserIdentityLookupInfo](instance.ID, objc.Sel("initWithUserRecordID:"), userRecordID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKUserIdentityLookupInfoWithUserRecordID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKUserIdentityLookupInfo */

// Returns an array of lookup infos for the specifed user record IDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/lookupInfos(with:)
func (cc _CKUserIdentityLookupInfoClass) LookupInfosWithRecordIDs(recordIDs []CKRecordID) []CKUserIdentityLookupInfo {
	rv := objc.Send[[]CKUserIdentityLookupInfo](objc.ID(cc.class), objc.Sel("lookupInfosWithRecordIDs:"), recordIDs)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LookupInfosWithRecordIDs) */


// Returns an array of lookup infos for the specifed email addresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/lookupInfos(withEmails:)
func (cc _CKUserIdentityLookupInfoClass) LookupInfosWithEmails(emails []string) []CKUserIdentityLookupInfo {
	rv := objc.Send[[]CKUserIdentityLookupInfo](objc.ID(cc.class), objc.Sel("lookupInfosWithEmails:"), emails)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LookupInfosWithEmails) */


// Returns an array of lookup infos for the specifed phone numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/lookupInfos(withPhoneNumbers:)
func (cc _CKUserIdentityLookupInfoClass) LookupInfosWithPhoneNumbers(phoneNumbers []string) []CKUserIdentityLookupInfo {
	rv := objc.Send[[]CKUserIdentityLookupInfo](objc.ID(cc.class), objc.Sel("lookupInfosWithPhoneNumbers:"), phoneNumbers)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LookupInfosWithPhoneNumbers) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKUserIdentityLookupInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKUserIdentityLookupInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKUserIdentityLookupInfo */

// The user’s email address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/emailAddress
func (c_ CKUserIdentityLookupInfo) EmailAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("emailAddress"))
	return rv
}/* debug [instance_properties/getter]: emailAddress */


// The user’s phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/phoneNumber
func (c_ CKUserIdentityLookupInfo) PhoneNumber() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("phoneNumber"))
	return rv
}/* debug [instance_properties/getter]: phoneNumber */


// The ID of the user record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/userRecordID
func (c_ CKUserIdentityLookupInfo) UserRecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("userRecordID"))
	return rv
}/* debug [instance_properties/getter]: userRecordID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKUserIdentityLookupInfo */


