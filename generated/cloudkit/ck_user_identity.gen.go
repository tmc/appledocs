// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKUserIdentity */


/* debug [class_header]: Header for CKUserIdentity */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKUserIdentity */
// An interface definition for the [CKUserIdentity] class.
type ICKUserIdentity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKUserIdentity */
	// properties:
	ContactIdentifiers() []string
	HasiCloudAccount() bool
	LookupInfo() ICKUserIdentityLookupInfo
	NameComponents() foundation.PersonNameComponents
	UserRecordID() ICKRecordID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKUserIdentity */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKUserIdentity */
// Alloc allocates a new instance without initialization.
func (cc _CKUserIdentityClass) Alloc() CKUserIdentity {
	rv := objc.Send[CKUserIdentity](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKUserIdentity */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKUserIdentity *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKUserIdentity */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKUserIdentity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKUserIdentity */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKUserIdentity */

// Identifiers that match contacts in the local Contacts database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/contactIdentifiers
func (c_ CKUserIdentity) ContactIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("contactIdentifiers"))
	return rv
}/* debug [instance_properties/getter]: contactIdentifiers */


// A Boolean value that indicates whether the user has an iCloud account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/hasiCloudAccount
func (c_ CKUserIdentity) HasiCloudAccount() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasiCloudAccount"))
	return rv
}/* debug [instance_properties/getter]: hasiCloudAccount */


// The lookup info for retrieving the user identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/lookupInfo-swift.property
func (c_ CKUserIdentity) LookupInfo() ICKUserIdentityLookupInfo {
	rv := objc.Send[CKUserIdentityLookupInfo](c_.ID, objc.Sel("lookupInfo"))
	return rv
}/* debug [instance_properties/getter]: lookupInfo */


// The user’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/nameComponents
func (c_ CKUserIdentity) NameComponents() foundation.PersonNameComponents {
	rv := objc.Send[foundation.PersonNameComponents](c_.ID, objc.Sel("nameComponents"))
	return rv
}/* debug [instance_properties/getter]: nameComponents */


// The user record ID for the corresponding user record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/userRecordID
func (c_ CKUserIdentity) UserRecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("userRecordID"))
	return rv
}/* debug [instance_properties/getter]: userRecordID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKUserIdentity */



