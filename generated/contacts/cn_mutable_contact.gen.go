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
	Birthday() foundation.DateComponents
	SetBirthday(value foundation.IDateComponents)
	ContactRelations() []CNLabeledValue
	SetContactRelations(value []CNLabeledValue)
	ContactType() CNContactType
	SetContactType(value CNContactType)
	Dates() []CNLabeledValue
	SetDates(value []CNLabeledValue)
	DepartmentName() string
	SetDepartmentName(value string)
	EmailAddresses() []CNLabeledValue
	SetEmailAddresses(value []CNLabeledValue)
	FamilyName() string
	SetFamilyName(value string)
	GivenName() string
	SetGivenName(value string)
	ImageData() foundation.NSData
	SetImageData(value foundation.IData)
	InstantMessageAddresses() []CNLabeledValue
	SetInstantMessageAddresses(value []CNLabeledValue)
	JobTitle() string
	SetJobTitle(value string)
	MiddleName() string
	SetMiddleName(value string)
	NamePrefix() string
	SetNamePrefix(value string)
	NameSuffix() string
	SetNameSuffix(value string)
	Nickname() string
	SetNickname(value string)
	NonGregorianBirthday() foundation.DateComponents
	SetNonGregorianBirthday(value foundation.IDateComponents)
	Note() string
	SetNote(value string)
	OrganizationName() string
	SetOrganizationName(value string)
	PhoneNumbers() []CNLabeledValue
	SetPhoneNumbers(value []CNLabeledValue)
	PhoneticFamilyName() string
	SetPhoneticFamilyName(value string)
	PhoneticGivenName() string
	SetPhoneticGivenName(value string)
	PhoneticMiddleName() string
	SetPhoneticMiddleName(value string)
	PhoneticOrganizationName() string
	SetPhoneticOrganizationName(value string)
	PostalAddresses() []CNLabeledValue
	SetPostalAddresses(value []CNLabeledValue)
	PreviousFamilyName() string
	SetPreviousFamilyName(value string)
	SocialProfiles() []CNLabeledValue
	SetSocialProfiles(value []CNLabeledValue)
	UrlAddresses() []CNLabeledValue
	SetUrlAddresses(value []CNLabeledValue)
	CNContactPropertyNotFetchedExceptionName() string
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



// A date component for the Gregorian birthday of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/birthday

func (c_ CNMutableContact) Birthday() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](c_.ID, objc.Sel("birthday"))
	return rv
}


// A date component for the Gregorian birthday of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/birthday

func (c_ CNMutableContact) SetBirthday(value foundation.IDateComponents) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBirthday:"), value)
}


// An array of labeled contact relations for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/contactRelations

func (c_ CNMutableContact) ContactRelations() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("contactRelations"))
	return rv
}


// An array of labeled contact relations for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/contactRelations

func (c_ CNMutableContact) SetContactRelations(value []CNLabeledValue) {
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


// An enum identifying the contact type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/contactType

func (c_ CNMutableContact) ContactType() CNContactType {
	rv := objc.Send[CNContactType](c_.ID, objc.Sel("contactType"))
	return rv
}


// An enum identifying the contact type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/contactType

func (c_ CNMutableContact) SetContactType(value CNContactType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactType:"), value)
}


// An array containing labeled Gregorian dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/dates

func (c_ CNMutableContact) Dates() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("dates"))
	return rv
}


// An array containing labeled Gregorian dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/dates

func (c_ CNMutableContact) SetDates(value []CNLabeledValue) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setDates:"), nsArray)
}


// The name of the department associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/departmentName

func (c_ CNMutableContact) DepartmentName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("departmentName"))
	return rv
}


// The name of the department associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/departmentName

func (c_ CNMutableContact) SetDepartmentName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDepartmentName:"), objc.String(value))
}


// An array of labeled email addresses for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/emailAddresses

func (c_ CNMutableContact) EmailAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("emailAddresses"))
	return rv
}


// An array of labeled email addresses for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/emailAddresses

func (c_ CNMutableContact) SetEmailAddresses(value []CNLabeledValue) {
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


// The family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/familyName

func (c_ CNMutableContact) FamilyName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("familyName"))
	return rv
}


// The family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/familyName

func (c_ CNMutableContact) SetFamilyName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFamilyName:"), objc.String(value))
}


// The given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/givenName

func (c_ CNMutableContact) GivenName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("givenName"))
	return rv
}


// The given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/givenName

func (c_ CNMutableContact) SetGivenName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGivenName:"), objc.String(value))
}


// The profile picture of a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/imageData

func (c_ CNMutableContact) ImageData() foundation.NSData {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("imageData"))
	return rv
}


// The profile picture of a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/imageData

func (c_ CNMutableContact) SetImageData(value foundation.IData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageData:"), value)
}


// An array of labeled IM addresses for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/instantMessageAddresses

func (c_ CNMutableContact) InstantMessageAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("instantMessageAddresses"))
	return rv
}


// An array of labeled IM addresses for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/instantMessageAddresses

func (c_ CNMutableContact) SetInstantMessageAddresses(value []CNLabeledValue) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstantMessageAddresses:"), nsArray)
}


// The contact’s job title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/jobTitle

func (c_ CNMutableContact) JobTitle() string {
	rv := objc.Send[string](c_.ID, objc.Sel("jobTitle"))
	return rv
}


// The contact’s job title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/jobTitle

func (c_ CNMutableContact) SetJobTitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setJobTitle:"), objc.String(value))
}


// The middle name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/middleName

func (c_ CNMutableContact) MiddleName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("middleName"))
	return rv
}


// The middle name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/middleName

func (c_ CNMutableContact) SetMiddleName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMiddleName:"), objc.String(value))
}


// The name prefix of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/namePrefix

func (c_ CNMutableContact) NamePrefix() string {
	rv := objc.Send[string](c_.ID, objc.Sel("namePrefix"))
	return rv
}


// The name prefix of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/namePrefix

func (c_ CNMutableContact) SetNamePrefix(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNamePrefix:"), objc.String(value))
}


// The name suffix of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/nameSuffix

func (c_ CNMutableContact) NameSuffix() string {
	rv := objc.Send[string](c_.ID, objc.Sel("nameSuffix"))
	return rv
}


// The name suffix of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/nameSuffix

func (c_ CNMutableContact) SetNameSuffix(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNameSuffix:"), objc.String(value))
}


// The nickname of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/nickname

func (c_ CNMutableContact) Nickname() string {
	rv := objc.Send[string](c_.ID, objc.Sel("nickname"))
	return rv
}


// The nickname of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/nickname

func (c_ CNMutableContact) SetNickname(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNickname:"), objc.String(value))
}


// A date component for the non-Gregorian birthday of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/nonGregorianBirthday

func (c_ CNMutableContact) NonGregorianBirthday() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](c_.ID, objc.Sel("nonGregorianBirthday"))
	return rv
}


// A date component for the non-Gregorian birthday of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/nonGregorianBirthday

func (c_ CNMutableContact) SetNonGregorianBirthday(value foundation.IDateComponents) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNonGregorianBirthday:"), value)
}


// A string containing notes for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/note

func (c_ CNMutableContact) Note() string {
	rv := objc.Send[string](c_.ID, objc.Sel("note"))
	return rv
}


// A string containing notes for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/note

func (c_ CNMutableContact) SetNote(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNote:"), objc.String(value))
}


// The name of the organization associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/organizationName

func (c_ CNMutableContact) OrganizationName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("organizationName"))
	return rv
}


// The name of the organization associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/organizationName

func (c_ CNMutableContact) SetOrganizationName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrganizationName:"), objc.String(value))
}


// An array of labeled phone numbers for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneNumbers

func (c_ CNMutableContact) PhoneNumbers() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("phoneNumbers"))
	return rv
}


// An array of labeled phone numbers for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneNumbers

func (c_ CNMutableContact) SetPhoneNumbers(value []CNLabeledValue) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneNumbers:"), nsArray)
}


// The phonetic family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticFamilyName

func (c_ CNMutableContact) PhoneticFamilyName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("phoneticFamilyName"))
	return rv
}


// The phonetic family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticFamilyName

func (c_ CNMutableContact) SetPhoneticFamilyName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneticFamilyName:"), objc.String(value))
}


// The phonetic given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticGivenName

func (c_ CNMutableContact) PhoneticGivenName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("phoneticGivenName"))
	return rv
}


// The phonetic given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticGivenName

func (c_ CNMutableContact) SetPhoneticGivenName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneticGivenName:"), objc.String(value))
}


// The phonetic middle name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticMiddleName

func (c_ CNMutableContact) PhoneticMiddleName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("phoneticMiddleName"))
	return rv
}


// The phonetic middle name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticMiddleName

func (c_ CNMutableContact) SetPhoneticMiddleName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneticMiddleName:"), objc.String(value))
}


// The phonetic name of the organization associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticOrganizationName

func (c_ CNMutableContact) PhoneticOrganizationName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("phoneticOrganizationName"))
	return rv
}


// The phonetic name of the organization associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticOrganizationName

func (c_ CNMutableContact) SetPhoneticOrganizationName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneticOrganizationName:"), objc.String(value))
}


// An array of labeled postal addresses for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/postalAddresses

func (c_ CNMutableContact) PostalAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("postalAddresses"))
	return rv
}


// An array of labeled postal addresses for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/postalAddresses

func (c_ CNMutableContact) SetPostalAddresses(value []CNLabeledValue) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setPostalAddresses:"), nsArray)
}


// The previous family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/previousFamilyName

func (c_ CNMutableContact) PreviousFamilyName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("previousFamilyName"))
	return rv
}


// The previous family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/previousFamilyName

func (c_ CNMutableContact) SetPreviousFamilyName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviousFamilyName:"), objc.String(value))
}


// An array of labeled social profiles for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/socialProfiles

func (c_ CNMutableContact) SocialProfiles() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("socialProfiles"))
	return rv
}


// An array of labeled social profiles for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/socialProfiles

func (c_ CNMutableContact) SetSocialProfiles(value []CNLabeledValue) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setSocialProfiles:"), nsArray)
}


// An array of labeled URL addresses for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/urlAddresses

func (c_ CNMutableContact) UrlAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("urlAddresses"))
	return rv
}


// An array of labeled URL addresses for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/urlAddresses

func (c_ CNMutableContact) SetUrlAddresses(value []CNLabeledValue) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setUrlAddresses:"), nsArray)
}


// Exception thrown when an accessed property was not fetched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactpropertynotfetchedexceptionname

func (c_ CNMutableContact) CNContactPropertyNotFetchedExceptionName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNContactPropertyNotFetchedExceptionName"))
	return rv
}



