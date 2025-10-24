// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CNMutableContact] class.
var (
	CNMutableContactClass     _CNMutableContactClass
	CNMutableContactClassOnce sync.Once
)

func getCNMutableContactClass() _CNMutableContactClass {
	CNMutableContactClassOnce.Do(func() {
		CNMutableContactClass = _CNMutableContactClass{objc.GetClass("CNMutableContact")}
	})
	return CNMutableContactClass
}

type _CNMutableContactClass struct {
	class objc.Class
}

// An interface definition for the [CNMutableContact] class.
type ICNMutableContact interface {
	ICNContact
	// properties:
	ContactRelations() []ICNLabeledValue
	SetContactRelations(value []ICNLabeledValue)
	EmailAddresses() []ICNLabeledValue
	SetEmailAddresses(value []ICNLabeledValue)
	ImageData() objc.IObject /* cross-framework: NSData */
	SetImageData(value objc.IObject /* cross-framework: NSData */)
	NamePrefix() objc.IObject /* cross-framework: NSString */
	SetNamePrefix(value objc.IObject /* cross-framework: NSString */)
	Nickname() objc.IObject /* cross-framework: NSString */
	SetNickname(value objc.IObject /* cross-framework: NSString */)
	PreviousFamilyName() objc.IObject /* cross-framework: NSString */
	SetPreviousFamilyName(value objc.IObject /* cross-framework: NSString */)
	CNContactPropertyNotFetchedExceptionName() objc.IObject /* cross-framework: NSString */
	Birthday() objc.IObject /* cross-framework: DateComponents */
	SetBirthday(value objc.IObject /* cross-framework: DateComponents */)
	ContactType() CNContactType
	SetContactType(value CNContactType)
	Dates() objc.IObject /* cross-framework: DateComponents */
	SetDates(value objc.IObject /* cross-framework: DateComponents */)
	DepartmentName() objc.IObject /* cross-framework: NSString */
	SetDepartmentName(value objc.IObject /* cross-framework: NSString */)
	FamilyName() objc.IObject /* cross-framework: NSString */
	SetFamilyName(value objc.IObject /* cross-framework: NSString */)
	GivenName() objc.IObject /* cross-framework: NSString */
	SetGivenName(value objc.IObject /* cross-framework: NSString */)
	Id() objc.IObject /* cross-framework: UUID */
	SetId(value objc.IObject /* cross-framework: UUID */)
	InstantMessageAddresses() ICNInstantMessageAddress
	SetInstantMessageAddresses(value ICNInstantMessageAddress)
	JobTitle() objc.IObject /* cross-framework: NSString */
	SetJobTitle(value objc.IObject /* cross-framework: NSString */)
	MiddleName() objc.IObject /* cross-framework: NSString */
	SetMiddleName(value objc.IObject /* cross-framework: NSString */)
	NameSuffix() objc.IObject /* cross-framework: NSString */
	SetNameSuffix(value objc.IObject /* cross-framework: NSString */)
	NonGregorianBirthday() objc.IObject /* cross-framework: DateComponents */
	SetNonGregorianBirthday(value objc.IObject /* cross-framework: DateComponents */)
	Note() objc.IObject /* cross-framework: NSString */
	SetNote(value objc.IObject /* cross-framework: NSString */)
	OrganizationName() objc.IObject /* cross-framework: NSString */
	SetOrganizationName(value objc.IObject /* cross-framework: NSString */)
	PhoneNumbers() ICNPhoneNumber
	SetPhoneNumbers(value ICNPhoneNumber)
	PhoneticFamilyName() objc.IObject /* cross-framework: NSString */
	SetPhoneticFamilyName(value objc.IObject /* cross-framework: NSString */)
	PhoneticGivenName() objc.IObject /* cross-framework: NSString */
	SetPhoneticGivenName(value objc.IObject /* cross-framework: NSString */)
	PhoneticMiddleName() objc.IObject /* cross-framework: NSString */
	SetPhoneticMiddleName(value objc.IObject /* cross-framework: NSString */)
	PhoneticOrganizationName() objc.IObject /* cross-framework: NSString */
	SetPhoneticOrganizationName(value objc.IObject /* cross-framework: NSString */)
	PostalAddresses() ICNPostalAddress
	SetPostalAddresses(value ICNPostalAddress)
	SocialProfiles() ICNSocialProfile
	SetSocialProfiles(value ICNSocialProfile)
	UrlAddresses() objc.IObject /* cross-framework: NSString */
	SetUrlAddresses(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A mutable object that stores information about a single contact, such as the contact’s first name, phone numbers, and addresses.
//
// objects are not a thread-safe class. To access the contact information in a thread-safe manner, use a object instead. You may modify only those properties whose values you fetched from the contacts database. When fetching a contact, you specify which properties you want to retrieve from the database. The contact store then populates the properties of a object with those values. After creating a mutable copy of that object, you can modify only those properties for which a value exists. If you attempt to access a property that is not available, the object throws a exception. To remove the value for a property, set string and array properties to empty, and set all other properties to .


// A mutable object that stores information about a single contact, such as the contact’s first name, phone numbers, and addresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact
type CNMutableContact struct {
	CNContact
}

// CNMutableContactFrom constructs a [CNMutableContact] from an unsafe.Pointer.
//
// A mutable object that stores information about a single contact, such as the contact’s first name, phone numbers, and addresses.
func CNMutableContactFrom(ptr unsafe.Pointer) CNMutableContact {
	return CNMutableContact{
		CNContact: CNContactFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNMutableContactClass) Alloc() CNMutableContact {
	rv := objc.Send[CNMutableContact](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNMutableContactClass) New() CNMutableContact {
	rv := objc.Send[CNMutableContact](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNMutableContact) Init() CNMutableContact {
	rv := objc.Send[CNMutableContact](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNMutableContact) Autorelease() CNMutableContact {
	rv := objc.Send[CNMutableContact](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNMutableContact creates a new CNMutableContact instance.
func NewCNMutableContact() CNMutableContact {
	return getCNMutableContactClass().New()
}



// An array of labeled contact relations for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/contactRelations
func (c_ CNMutableContact) ContactRelations() []ICNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("contactRelations"))
	return rv
}


// An array of labeled contact relations for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/contactRelations
func (c_ CNMutableContact) SetContactRelations(value []ICNLabeledValue) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactRelations:"), nsArray)
}


// An array of labeled email addresses for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/emailAddresses
func (c_ CNMutableContact) EmailAddresses() []ICNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("emailAddresses"))
	return rv
}


// An array of labeled email addresses for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/emailAddresses
func (c_ CNMutableContact) SetEmailAddresses(value []ICNLabeledValue) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmailAddresses:"), nsArray)
}


// The profile picture of a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/imageData
func (c_ CNMutableContact) ImageData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("imageData"))
	return rv
}


// The profile picture of a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/imageData
func (c_ CNMutableContact) SetImageData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageData:"), value)
}


// The name prefix of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/namePrefix
func (c_ CNMutableContact) NamePrefix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("namePrefix"))
	return rv
}


// The name prefix of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/namePrefix
func (c_ CNMutableContact) SetNamePrefix(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNamePrefix:"), value)
}


// The nickname of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/nickname
func (c_ CNMutableContact) Nickname() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("nickname"))
	return rv
}


// The nickname of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/nickname
func (c_ CNMutableContact) SetNickname(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNickname:"), value)
}


// The previous family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/previousFamilyName
func (c_ CNMutableContact) PreviousFamilyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("previousFamilyName"))
	return rv
}


// The previous family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/previousFamilyName
func (c_ CNMutableContact) SetPreviousFamilyName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviousFamilyName:"), value)
}


// Exception thrown when an accessed property was not fetched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactpropertynotfetchedexceptionname
func (c_ CNMutableContact) CNContactPropertyNotFetchedExceptionName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNContactPropertyNotFetchedExceptionName"))
	return rv
}


// A date component for the Gregorian birthday of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/birthday
func (c_ CNMutableContact) Birthday() objc.IObject /* cross-framework: DateComponents */ {
	rv := objc.Send[foundation.DateComponents](c_.ID, objc.Sel("birthday"))
	return rv
}


// A date component for the Gregorian birthday of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/birthday
func (c_ CNMutableContact) SetBirthday(value objc.IObject /* cross-framework: DateComponents */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBirthday:"), value)
}


// An enum identifying the contact type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/contacttype
func (c_ CNMutableContact) ContactType() CNContactType {
	rv := objc.Send[CNContactType](c_.ID, objc.Sel("contactType"))
	return rv
}


// An enum identifying the contact type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/contacttype
func (c_ CNMutableContact) SetContactType(value CNContactType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactType:"), value)
}


// An array containing labeled Gregorian dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/dates
func (c_ CNMutableContact) Dates() objc.IObject /* cross-framework: DateComponents */ {
	rv := objc.Send[foundation.DateComponents](c_.ID, objc.Sel("dates"))
	return rv
}


// An array containing labeled Gregorian dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/dates
func (c_ CNMutableContact) SetDates(value objc.IObject /* cross-framework: DateComponents */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDates:"), value)
}


// The name of the department associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/departmentname
func (c_ CNMutableContact) DepartmentName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("departmentName"))
	return rv
}


// The name of the department associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/departmentname
func (c_ CNMutableContact) SetDepartmentName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDepartmentName:"), value)
}


// The family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/familyname
func (c_ CNMutableContact) FamilyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("familyName"))
	return rv
}


// The family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/familyname
func (c_ CNMutableContact) SetFamilyName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFamilyName:"), value)
}


// The given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/givenname
func (c_ CNMutableContact) GivenName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("givenName"))
	return rv
}


// The given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/givenname
func (c_ CNMutableContact) SetGivenName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGivenName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/id
func (c_ CNMutableContact) Id() objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("id"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/id
func (c_ CNMutableContact) SetId(value objc.IObject /* cross-framework: UUID */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setId:"), value)
}


// An array of labeled IM addresses for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/instantmessageaddresses
func (c_ CNMutableContact) InstantMessageAddresses() ICNInstantMessageAddress {
	rv := objc.Send[CNInstantMessageAddress](c_.ID, objc.Sel("instantMessageAddresses"))
	return rv
}


// An array of labeled IM addresses for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/instantmessageaddresses
func (c_ CNMutableContact) SetInstantMessageAddresses(value ICNInstantMessageAddress) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstantMessageAddresses:"), value)
}


// The contact’s job title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/jobtitle
func (c_ CNMutableContact) JobTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("jobTitle"))
	return rv
}


// The contact’s job title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/jobtitle
func (c_ CNMutableContact) SetJobTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setJobTitle:"), value)
}


// The middle name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/middlename
func (c_ CNMutableContact) MiddleName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("middleName"))
	return rv
}


// The middle name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/middlename
func (c_ CNMutableContact) SetMiddleName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMiddleName:"), value)
}


// The name suffix of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/namesuffix
func (c_ CNMutableContact) NameSuffix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("nameSuffix"))
	return rv
}


// The name suffix of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/namesuffix
func (c_ CNMutableContact) SetNameSuffix(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNameSuffix:"), value)
}


// A date component for the non-Gregorian birthday of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/nongregorianbirthday
func (c_ CNMutableContact) NonGregorianBirthday() objc.IObject /* cross-framework: DateComponents */ {
	rv := objc.Send[foundation.DateComponents](c_.ID, objc.Sel("nonGregorianBirthday"))
	return rv
}


// A date component for the non-Gregorian birthday of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/nongregorianbirthday
func (c_ CNMutableContact) SetNonGregorianBirthday(value objc.IObject /* cross-framework: DateComponents */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNonGregorianBirthday:"), value)
}


// A string containing notes for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/note
func (c_ CNMutableContact) Note() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("note"))
	return rv
}


// A string containing notes for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/note
func (c_ CNMutableContact) SetNote(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNote:"), value)
}


// The name of the organization associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/organizationname
func (c_ CNMutableContact) OrganizationName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("organizationName"))
	return rv
}


// The name of the organization associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/organizationname
func (c_ CNMutableContact) SetOrganizationName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrganizationName:"), value)
}


// An array of labeled phone numbers for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/phonenumbers
func (c_ CNMutableContact) PhoneNumbers() ICNPhoneNumber {
	rv := objc.Send[CNPhoneNumber](c_.ID, objc.Sel("phoneNumbers"))
	return rv
}


// An array of labeled phone numbers for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/phonenumbers
func (c_ CNMutableContact) SetPhoneNumbers(value ICNPhoneNumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneNumbers:"), value)
}


// The phonetic family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/phoneticfamilyname
func (c_ CNMutableContact) PhoneticFamilyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("phoneticFamilyName"))
	return rv
}


// The phonetic family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/phoneticfamilyname
func (c_ CNMutableContact) SetPhoneticFamilyName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneticFamilyName:"), value)
}


// The phonetic given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/phoneticgivenname
func (c_ CNMutableContact) PhoneticGivenName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("phoneticGivenName"))
	return rv
}


// The phonetic given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/phoneticgivenname
func (c_ CNMutableContact) SetPhoneticGivenName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneticGivenName:"), value)
}


// The phonetic middle name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/phoneticmiddlename
func (c_ CNMutableContact) PhoneticMiddleName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("phoneticMiddleName"))
	return rv
}


// The phonetic middle name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/phoneticmiddlename
func (c_ CNMutableContact) SetPhoneticMiddleName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneticMiddleName:"), value)
}


// The phonetic name of the organization associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/phoneticorganizationname
func (c_ CNMutableContact) PhoneticOrganizationName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("phoneticOrganizationName"))
	return rv
}


// The phonetic name of the organization associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/phoneticorganizationname
func (c_ CNMutableContact) SetPhoneticOrganizationName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneticOrganizationName:"), value)
}


// An array of labeled postal addresses for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/postaladdresses
func (c_ CNMutableContact) PostalAddresses() ICNPostalAddress {
	rv := objc.Send[CNPostalAddress](c_.ID, objc.Sel("postalAddresses"))
	return rv
}


// An array of labeled postal addresses for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/postaladdresses
func (c_ CNMutableContact) SetPostalAddresses(value ICNPostalAddress) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPostalAddresses:"), value)
}


// An array of labeled social profiles for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/socialprofiles
func (c_ CNMutableContact) SocialProfiles() ICNSocialProfile {
	rv := objc.Send[CNSocialProfile](c_.ID, objc.Sel("socialProfiles"))
	return rv
}


// An array of labeled social profiles for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/socialprofiles
func (c_ CNMutableContact) SetSocialProfiles(value ICNSocialProfile) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSocialProfiles:"), value)
}


// An array of labeled URL addresses for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/urladdresses
func (c_ CNMutableContact) UrlAddresses() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("urlAddresses"))
	return rv
}


// An array of labeled URL addresses for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablecontact/urladdresses
func (c_ CNMutableContact) SetUrlAddresses(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUrlAddresses:"), value)
}



