// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CNContact] class.
type ICNContact interface {
	objectivec.IObject
	AreKeysAvailable(keyDescriptors []objc.ID) bool
	IsKeyAvailable(key appkit.string) bool
	IsUnifiedWithContactWithIdentifier(contactIdentifier appkit.string) bool
}

// An immutable object that stores information about a single contact, such as the contact’s first name, phone numbers, and addresses.
//
// A object stores an immutable copy of a contact’s information, so you cannot change the information in this object directly. Contact objects are thread-safe, so you may access them from any thread of your app. To modify a contact’s information, call the method to obtain a object with the same information. After modifying the mutable contact, save your changes back to the contacts database using the object. Every contact in the contacts database has a unique ID, which you access using the property. The mutable and immutable versions of the same contact have the same identifier.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CNContactClass) Alloc() CNContact {
	rv := objc.Send[CNContact](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns a comparator to sort contacts with the specified order.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/comparator(forNameSortOrder:)
func (cc _CNContactClass) ComparatorForNameSortOrder(sortOrder ICNContactSortOrder) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("comparatorForNameSortOrder:"), sortOrder)
	return rv
}

// Fetches all the keys required for the contact sort comparator.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/descriptorForAllComparatorKeys()
func (cc _CNContactClass) DescriptorForAllComparatorKeys() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("descriptorForAllComparatorKeys"))
	return rv
}

// Returns a string containing the localized contact property name.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/localizedString(forKey:)
func (cc _CNContactClass) LocalizedStringForKey(key appkit.string) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("localizedStringForKey:"), key)
	return rv
}

// Returns a predicate to find the contacts whose phone number matches the specified value.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/predicateForContacts(matching:)
func (cc _CNContactClass) PredicateForContactsMatchingPhoneNumber(phoneNumber ICNPhoneNumber) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContactsMatchingPhoneNumber:"), phoneNumber)
	return rv
}

// Returns a predicate to find the contacts whose email address matches the specified value.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/predicateForContacts(matchingEmailAddress:)
func (cc _CNContactClass) PredicateForContactsMatchingEmailAddress(emailAddress appkit.string) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContactsMatchingEmailAddress:"), emailAddress)
	return rv
}

// Returns a predicate to find the contacts matching the specified name.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/predicateForContacts(matchingName:)
func (cc _CNContactClass) PredicateForContactsMatchingName(name appkit.string) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContactsMatchingName:"), name)
	return rv
}

// Returns a predicate to find the contacts matching the specified identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/predicateForContacts(withIdentifiers:)
func (cc _CNContactClass) PredicateForContactsWithIdentifiers(identifiers []string) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContactsWithIdentifiers:"), identifiers)
	return rv
}

// Returns a predicate to find the contacts in the specified container.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/predicateForContactsInContainer(withIdentifier:)
func (cc _CNContactClass) PredicateForContactsInContainerWithIdentifier(containerIdentifier appkit.string) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContactsInContainerWithIdentifier:"), containerIdentifier)
	return rv
}

// Returns a predicate to find the contacts that are members in the specified group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/predicateForContactsInGroup(withIdentifier:)
func (cc _CNContactClass) PredicateForContactsInGroupWithIdentifier(groupIdentifier appkit.string) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContactsInGroupWithIdentifier:"), groupIdentifier)
	return rv
}

// Determines whether all contact property values for the specified keys are fetched.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/areKeysAvailable(_:)
func (c_ CNContact) AreKeysAvailable(keyDescriptors []objc.ID) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("areKeysAvailable:"), keyDescriptors)
	return rv
}

// Determines whether the contact property value for the specified key is fetched.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/isKeyAvailable(_:)
func (c_ CNContact) IsKeyAvailable(key appkit.string) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isKeyAvailable:"), key)
	return rv
}

// Returns a Boolean indicating whether the current contact is a unified contact and includes a contact with the specified identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/isUnifiedWithContact(withIdentifier:)
func (c_ CNContact) IsUnifiedWithContactWithIdentifier(contactIdentifier appkit.string) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isUnifiedWithContactWithIdentifier:"), contactIdentifier)
	return rv
}

// A date component for the Gregorian birthday of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/birthday
func (c_ CNContact) Birthday() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](c_.ID, objc.Sel("birthday"))
	return rv
}

// An array of labeled relations for the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/contactRelations
func (c_ CNContact) ContactRelations() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("contactRelations"))
	return rv
}

// An enum identifying the contact type.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/contactType
func (c_ CNContact) ContactType() CNContactType {
	rv := objc.Send[CNContactType](c_.ID, objc.Sel("contactType"))
	return rv
}

// An array containing labeled Gregorian dates.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/dates
func (c_ CNContact) Dates() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("dates"))
	return rv
}

// The name of the department associated with the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/departmentName
func (c_ CNContact) DepartmentName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("departmentName"))
	return rv
}

// An array of labeled email addresses for the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/emailAddresses
func (c_ CNContact) EmailAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("emailAddresses"))
	return rv
}

// The family name of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/familyName
func (c_ CNContact) FamilyName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("familyName"))
	return rv
}

// The given name of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/givenName
func (c_ CNContact) GivenName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("givenName"))
	return rv
}

// A value that uniquely identifies a contact on the device.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/identifier
func (c_ CNContact) Identifier() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("identifier"))
	return rv
}

// The profile picture of a contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/imageData
func (c_ CNContact) ImageData() foundation.NSData {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("imageData"))
	return rv
}

// A Boolean indicating whether a contact has a profile picture.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/imageDataAvailable
func (c_ CNContact) ImageDataAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("imageDataAvailable"))
	return rv
}

// An array of labeled IM addresses for the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/instantMessageAddresses
func (c_ CNContact) InstantMessageAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("instantMessageAddresses"))
	return rv
}

// The contact’s job title.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/jobTitle
func (c_ CNContact) JobTitle() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("jobTitle"))
	return rv
}

// The middle name of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/middleName
func (c_ CNContact) MiddleName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("middleName"))
	return rv
}

// The name prefix of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/namePrefix
func (c_ CNContact) NamePrefix() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("namePrefix"))
	return rv
}

// The name suffix of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/nameSuffix
func (c_ CNContact) NameSuffix() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("nameSuffix"))
	return rv
}

// The nickname of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/nickname
func (c_ CNContact) Nickname() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("nickname"))
	return rv
}

// A date component for the non-Gregorian birthday of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/nonGregorianBirthday
func (c_ CNContact) NonGregorianBirthday() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](c_.ID, objc.Sel("nonGregorianBirthday"))
	return rv
}

// A string containing notes for the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/note
func (c_ CNContact) Note() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("note"))
	return rv
}

// The name of the organization associated with the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/organizationName
func (c_ CNContact) OrganizationName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("organizationName"))
	return rv
}

// An array of labeled phone numbers for a contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/phoneNumbers
func (c_ CNContact) PhoneNumbers() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("phoneNumbers"))
	return rv
}

// A string for the phonetic family name of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/phoneticFamilyName
func (c_ CNContact) PhoneticFamilyName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("phoneticFamilyName"))
	return rv
}

// The phonetic given name of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/phoneticGivenName
func (c_ CNContact) PhoneticGivenName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("phoneticGivenName"))
	return rv
}

// The phonetic middle name of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/phoneticMiddleName
func (c_ CNContact) PhoneticMiddleName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("phoneticMiddleName"))
	return rv
}

// The phonetic name of the organization associated with the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/phoneticOrganizationName
func (c_ CNContact) PhoneticOrganizationName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("phoneticOrganizationName"))
	return rv
}

// An array of labeled postal addresses for a contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/postalAddresses
func (c_ CNContact) PostalAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("postalAddresses"))
	return rv
}

// A string for the previous family name of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/previousFamilyName
func (c_ CNContact) PreviousFamilyName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("previousFamilyName"))
	return rv
}

// An array of labeled social profiles for a contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/socialProfiles
func (c_ CNContact) SocialProfiles() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("socialProfiles"))
	return rv
}

// The thumbnail version of the contact’s profile picture.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/thumbnailImageData
func (c_ CNContact) ThumbnailImageData() foundation.NSData {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("thumbnailImageData"))
	return rv
}

// An array of labeled URL addresses for a contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContact/urlAddresses
func (c_ CNContact) UrlAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("urlAddresses"))
	return rv
}



