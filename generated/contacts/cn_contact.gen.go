// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNContact */


/* debug [class_header]: Header for CNContact */
// The class instance for the [CNContact] class.
var (
	CNContactClass     _CNContactClass
	CNContactClassOnce sync.Once
)

func getCNContactClass() _CNContactClass {
	CNContactClassOnce.Do(func() {
		CNContactClass = _CNContactClass{objc.GetClass("CNContact")}
	})
	return CNContactClass
}

type _CNContactClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNContact */
// An interface definition for the [CNContact] class.
type ICNContact interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNContact */
	// properties:
	Birthday() foundation.DateComponents
	ContactRelations() []CNLabeledValue
	ContactType() CNContactType
	Dates() []CNLabeledValue
	DepartmentName() objc.IObject /* cross-framework: NSString */
	EmailAddresses() []CNLabeledValue
	FamilyName() objc.IObject /* cross-framework: NSString */
	GivenName() objc.IObject /* cross-framework: NSString */
	Identifier() objc.IObject /* cross-framework: NSString */
	ImageData() objc.IObject /* cross-framework: NSData */
	ImageDataAvailable() bool
	InstantMessageAddresses() []CNLabeledValue
	JobTitle() objc.IObject /* cross-framework: NSString */
	MiddleName() objc.IObject /* cross-framework: NSString */
	NamePrefix() objc.IObject /* cross-framework: NSString */
	NameSuffix() objc.IObject /* cross-framework: NSString */
	Nickname() objc.IObject /* cross-framework: NSString */
	NonGregorianBirthday() foundation.DateComponents
	Note() objc.IObject /* cross-framework: NSString */
	OrganizationName() objc.IObject /* cross-framework: NSString */
	PhoneNumbers() []CNLabeledValue
	PhoneticFamilyName() objc.IObject /* cross-framework: NSString */
	PhoneticGivenName() objc.IObject /* cross-framework: NSString */
	PhoneticMiddleName() objc.IObject /* cross-framework: NSString */
	PhoneticOrganizationName() objc.IObject /* cross-framework: NSString */
	PostalAddresses() []CNLabeledValue
	PreviousFamilyName() objc.IObject /* cross-framework: NSString */
	SocialProfiles() []CNLabeledValue
	ThumbnailImageData() objc.IObject /* cross-framework: NSData */
	UrlAddresses() []CNLabeledValue
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNContact */
	// methods:
	AreKeysAvailable(keyDescriptors []objc.ID) bool
	IsKeyAvailable(key objc.IObject /* cross-framework: NSString */) bool
	IsUnifiedWithContactWithIdentifier(contactIdentifier objc.IObject /* cross-framework: NSString */) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNContact */
// Alloc allocates a new instance without initialization.
func (cc _CNContactClass) Alloc() CNContact {
	rv := objc.Send[CNContact](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNContactClass) New() CNContact {
	rv := objc.Send[CNContact](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContact) Init() CNContact {
	rv := objc.Send[CNContact](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContact) Autorelease() CNContact {
	rv := objc.Send[CNContact](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContact creates a new CNContact instance.
func NewCNContact() CNContact {
	return getCNContactClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNContact */
// An immutable object that stores information about a single contact, such as the contact’s first name, phone numbers, and addresses.
//
// A object stores an immutable copy of a contact’s information, so you cannot change the information in this object directly. Contact objects are thread-safe, so you may access them from any thread of your app. To modify a contact’s information, call the method to obtain a object with the same information. After modifying the mutable contact, save your changes back to the contacts database using the object. Every contact in the contacts database has a unique ID, which you access using the property. The mutable and immutable versions of the same contact have the same identifier.


// An immutable object that stores information about a single contact, such as the contact’s first name, phone numbers, and addresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact
type CNContact struct {
	objectivec.Object
}

// CNContactFrom constructs a [CNContact] from an unsafe.Pointer.
//
// An immutable object that stores information about a single contact, such as the contact’s first name, phone numbers, and addresses.
func CNContactFrom(ptr unsafe.Pointer) CNContact {
	return CNContact{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNContact *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNContact */

// Returns a comparator to sort contacts with the specified order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/comparator(forNameSortOrder:)
func (cc _CNContactClass) ComparatorForNameSortOrder(sortOrder CNContactSortOrder) Comparator /* not a class type */ {
	rv := objc.Send[Comparator](objc.ID(cc.class), objc.Sel("comparatorForNameSortOrder:"), sortOrder)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ComparatorForNameSortOrder) */


// Fetches all the keys required for the contact sort comparator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/descriptorForAllComparatorKeys()
func (cc _CNContactClass) DescriptorForAllComparatorKeys() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorForAllComparatorKeys"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorForAllComparatorKeys) */


// Returns a string containing the localized contact property name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/localizedString(forKey:)
func (cc _CNContactClass) LocalizedStringForKey(key objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("localizedStringForKey:"), key)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocalizedStringForKey) */


// Returns a predicate to find the contacts matching the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/predicateForContacts(matchingName:)
func (cc _CNContactClass) PredicateForContactsMatchingName(name objc.IObject /* cross-framework: NSString */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContactsMatchingName:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForContactsMatchingName) */


// Returns a predicate to find the contacts whose phone number matches the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/predicateForContacts(matching:)
func (cc _CNContactClass) PredicateForContactsMatchingPhoneNumber(phoneNumber ICNPhoneNumber) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContactsMatchingPhoneNumber:"), phoneNumber)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForContactsMatchingPhoneNumber) */


// Returns a predicate to find the contacts whose email address matches the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/predicateForContacts(matchingEmailAddress:)
func (cc _CNContactClass) PredicateForContactsMatchingEmailAddress(emailAddress objc.IObject /* cross-framework: NSString */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContactsMatchingEmailAddress:"), emailAddress)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForContactsMatchingEmailAddress) */


// Returns a predicate to find the contacts matching the specified identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/predicateForContacts(withIdentifiers:)
func (cc _CNContactClass) PredicateForContactsWithIdentifiers(identifiers []string) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContactsWithIdentifiers:"), identifiers)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForContactsWithIdentifiers) */


// Returns a predicate to find the contacts in the specified container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/predicateForContactsInContainer(withIdentifier:)
func (cc _CNContactClass) PredicateForContactsInContainerWithIdentifier(containerIdentifier objc.IObject /* cross-framework: NSString */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContactsInContainerWithIdentifier:"), containerIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForContactsInContainerWithIdentifier) */


// Returns a predicate to find the contacts that are members in the specified group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/predicateForContactsInGroup(withIdentifier:)
func (cc _CNContactClass) PredicateForContactsInGroupWithIdentifier(groupIdentifier objc.IObject /* cross-framework: NSString */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContactsInGroupWithIdentifier:"), groupIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForContactsInGroupWithIdentifier) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNContact */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNContact */

// Determines whether all contact property values for the specified keys are fetched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/areKeysAvailable(_:)
func (c_ CNContact) AreKeysAvailable(keyDescriptors []objc.ID) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("areKeysAvailable:"), keyDescriptors)
	return rv
}/* debug [instance_methods/method]: AreKeysAvailable */


// Determines whether the contact property value for the specified key is fetched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/isKeyAvailable(_:)
func (c_ CNContact) IsKeyAvailable(key objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isKeyAvailable:"), key)
	return rv
}/* debug [instance_methods/method]: IsKeyAvailable */


// Returns a Boolean indicating whether the current contact is a unified contact and includes a contact with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/isUnifiedWithContact(withIdentifier:)
func (c_ CNContact) IsUnifiedWithContactWithIdentifier(contactIdentifier objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isUnifiedWithContactWithIdentifier:"), contactIdentifier)
	return rv
}/* debug [instance_methods/method]: IsUnifiedWithContactWithIdentifier */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNContact */

// A date component for the Gregorian birthday of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/birthday
func (c_ CNContact) Birthday() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](c_.ID, objc.Sel("birthday"))
	return rv
}/* debug [instance_properties/getter]: birthday */


// An array of labeled relations for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/contactRelations
func (c_ CNContact) ContactRelations() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("contactRelations"))
	return rv
}/* debug [instance_properties/getter]: contactRelations */


// An enum identifying the contact type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/contactType
func (c_ CNContact) ContactType() CNContactType {
	rv := objc.Send[CNContactType](c_.ID, objc.Sel("contactType"))
	return rv
}/* debug [instance_properties/getter]: contactType */


// An array containing labeled Gregorian dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/dates
func (c_ CNContact) Dates() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("dates"))
	return rv
}/* debug [instance_properties/getter]: dates */


// The name of the department associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/departmentName
func (c_ CNContact) DepartmentName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("departmentName"))
	return rv
}/* debug [instance_properties/getter]: departmentName */


// An array of labeled email addresses for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/emailAddresses
func (c_ CNContact) EmailAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("emailAddresses"))
	return rv
}/* debug [instance_properties/getter]: emailAddresses */


// The family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/familyName
func (c_ CNContact) FamilyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("familyName"))
	return rv
}/* debug [instance_properties/getter]: familyName */


// The given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/givenName
func (c_ CNContact) GivenName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("givenName"))
	return rv
}/* debug [instance_properties/getter]: givenName */


// A value that uniquely identifies a contact on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/identifier
func (c_ CNContact) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The profile picture of a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/imageData
func (c_ CNContact) ImageData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("imageData"))
	return rv
}/* debug [instance_properties/getter]: imageData */


// A Boolean indicating whether a contact has a profile picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/imageDataAvailable
func (c_ CNContact) ImageDataAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("imageDataAvailable"))
	return rv
}/* debug [instance_properties/getter]: imageDataAvailable */


// An array of labeled IM addresses for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/instantMessageAddresses
func (c_ CNContact) InstantMessageAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("instantMessageAddresses"))
	return rv
}/* debug [instance_properties/getter]: instantMessageAddresses */


// The contact’s job title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/jobTitle
func (c_ CNContact) JobTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("jobTitle"))
	return rv
}/* debug [instance_properties/getter]: jobTitle */


// The middle name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/middleName
func (c_ CNContact) MiddleName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("middleName"))
	return rv
}/* debug [instance_properties/getter]: middleName */


// The name prefix of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/namePrefix
func (c_ CNContact) NamePrefix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("namePrefix"))
	return rv
}/* debug [instance_properties/getter]: namePrefix */


// The name suffix of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/nameSuffix
func (c_ CNContact) NameSuffix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("nameSuffix"))
	return rv
}/* debug [instance_properties/getter]: nameSuffix */


// The nickname of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/nickname
func (c_ CNContact) Nickname() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("nickname"))
	return rv
}/* debug [instance_properties/getter]: nickname */


// A date component for the non-Gregorian birthday of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/nonGregorianBirthday
func (c_ CNContact) NonGregorianBirthday() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](c_.ID, objc.Sel("nonGregorianBirthday"))
	return rv
}/* debug [instance_properties/getter]: nonGregorianBirthday */


// A string containing notes for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/note
func (c_ CNContact) Note() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("note"))
	return rv
}/* debug [instance_properties/getter]: note */


// The name of the organization associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/organizationName
func (c_ CNContact) OrganizationName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("organizationName"))
	return rv
}/* debug [instance_properties/getter]: organizationName */


// An array of labeled phone numbers for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/phoneNumbers
func (c_ CNContact) PhoneNumbers() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("phoneNumbers"))
	return rv
}/* debug [instance_properties/getter]: phoneNumbers */


// A string for the phonetic family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/phoneticFamilyName
func (c_ CNContact) PhoneticFamilyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("phoneticFamilyName"))
	return rv
}/* debug [instance_properties/getter]: phoneticFamilyName */


// The phonetic given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/phoneticGivenName
func (c_ CNContact) PhoneticGivenName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("phoneticGivenName"))
	return rv
}/* debug [instance_properties/getter]: phoneticGivenName */


// The phonetic middle name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/phoneticMiddleName
func (c_ CNContact) PhoneticMiddleName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("phoneticMiddleName"))
	return rv
}/* debug [instance_properties/getter]: phoneticMiddleName */


// The phonetic name of the organization associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/phoneticOrganizationName
func (c_ CNContact) PhoneticOrganizationName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("phoneticOrganizationName"))
	return rv
}/* debug [instance_properties/getter]: phoneticOrganizationName */


// An array of labeled postal addresses for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/postalAddresses
func (c_ CNContact) PostalAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("postalAddresses"))
	return rv
}/* debug [instance_properties/getter]: postalAddresses */


// A string for the previous family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/previousFamilyName
func (c_ CNContact) PreviousFamilyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("previousFamilyName"))
	return rv
}/* debug [instance_properties/getter]: previousFamilyName */


// An array of labeled social profiles for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/socialProfiles
func (c_ CNContact) SocialProfiles() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("socialProfiles"))
	return rv
}/* debug [instance_properties/getter]: socialProfiles */


// The thumbnail version of the contact’s profile picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/thumbnailImageData
func (c_ CNContact) ThumbnailImageData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("thumbnailImageData"))
	return rv
}/* debug [instance_properties/getter]: thumbnailImageData */


// An array of labeled URL addresses for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/urlAddresses
func (c_ CNContact) UrlAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("urlAddresses"))
	return rv
}/* debug [instance_properties/getter]: urlAddresses */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNContact */



