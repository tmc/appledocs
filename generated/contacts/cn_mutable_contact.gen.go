// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CNMutableContact */


/* debug [class_header]: Header for CNMutableContact */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNMutableContact */
// An interface definition for the [CNMutableContact] class.
type ICNMutableContact interface {
	ICNContact
	
/* debug [class_interface_properties]: Properties for CNMutableContact */
	// properties:
	Birthday() foundation.DateComponents
	SetBirthday(value foundation.DateComponents)
	ContactRelations() []CNLabeledValue
	SetContactRelations(value []CNLabeledValue)
	ContactType() CNContactType
	SetContactType(value CNContactType)
	Dates() []CNLabeledValue
	SetDates(value []CNLabeledValue)
	DepartmentName() objc.IObject /* cross-framework: NSString */
	SetDepartmentName(value objc.IObject /* cross-framework: NSString */)
	EmailAddresses() []CNLabeledValue
	SetEmailAddresses(value []CNLabeledValue)
	FamilyName() objc.IObject /* cross-framework: NSString */
	SetFamilyName(value objc.IObject /* cross-framework: NSString */)
	GivenName() objc.IObject /* cross-framework: NSString */
	SetGivenName(value objc.IObject /* cross-framework: NSString */)
	ImageData() objc.IObject /* cross-framework: NSData */
	SetImageData(value objc.IObject /* cross-framework: NSData */)
	InstantMessageAddresses() []CNLabeledValue
	SetInstantMessageAddresses(value []CNLabeledValue)
	JobTitle() objc.IObject /* cross-framework: NSString */
	SetJobTitle(value objc.IObject /* cross-framework: NSString */)
	MiddleName() objc.IObject /* cross-framework: NSString */
	SetMiddleName(value objc.IObject /* cross-framework: NSString */)
	NamePrefix() objc.IObject /* cross-framework: NSString */
	SetNamePrefix(value objc.IObject /* cross-framework: NSString */)
	NameSuffix() objc.IObject /* cross-framework: NSString */
	SetNameSuffix(value objc.IObject /* cross-framework: NSString */)
	Nickname() objc.IObject /* cross-framework: NSString */
	SetNickname(value objc.IObject /* cross-framework: NSString */)
	NonGregorianBirthday() foundation.DateComponents
	SetNonGregorianBirthday(value foundation.DateComponents)
	Note() objc.IObject /* cross-framework: NSString */
	SetNote(value objc.IObject /* cross-framework: NSString */)
	OrganizationName() objc.IObject /* cross-framework: NSString */
	SetOrganizationName(value objc.IObject /* cross-framework: NSString */)
	PhoneNumbers() []CNLabeledValue
	SetPhoneNumbers(value []CNLabeledValue)
	PhoneticFamilyName() objc.IObject /* cross-framework: NSString */
	SetPhoneticFamilyName(value objc.IObject /* cross-framework: NSString */)
	PhoneticGivenName() objc.IObject /* cross-framework: NSString */
	SetPhoneticGivenName(value objc.IObject /* cross-framework: NSString */)
	PhoneticMiddleName() objc.IObject /* cross-framework: NSString */
	SetPhoneticMiddleName(value objc.IObject /* cross-framework: NSString */)
	PhoneticOrganizationName() objc.IObject /* cross-framework: NSString */
	SetPhoneticOrganizationName(value objc.IObject /* cross-framework: NSString */)
	PostalAddresses() []CNLabeledValue
	SetPostalAddresses(value []CNLabeledValue)
	PreviousFamilyName() objc.IObject /* cross-framework: NSString */
	SetPreviousFamilyName(value objc.IObject /* cross-framework: NSString */)
	SocialProfiles() []CNLabeledValue
	SetSocialProfiles(value []CNLabeledValue)
	UrlAddresses() []CNLabeledValue
	SetUrlAddresses(value []CNLabeledValue)
	CNContactPropertyNotFetchedExceptionName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNMutableContact */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNMutableContact */
// Alloc allocates a new instance without initialization.
func (cc _CNMutableContactClass) Alloc() CNMutableContact {
	rv := objc.Send[CNMutableContact](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNMutableContact */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNMutableContact *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNMutableContact */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNMutableContact */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNMutableContact */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNMutableContact */

// A date component for the Gregorian birthday of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/birthday
func (c_ CNMutableContact) Birthday() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](c_.ID, objc.Sel("birthday"))
	return rv
}/* debug [instance_properties/getter]: birthday */


// A date component for the Gregorian birthday of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/birthday
func (c_ CNMutableContact) SetBirthday(value foundation.DateComponents) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBirthday:"), value)
}/* debug [instance_properties/setter]: birthday */


// An array of labeled contact relations for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/contactRelations
func (c_ CNMutableContact) ContactRelations() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("contactRelations"))
	return rv
}/* debug [instance_properties/getter]: contactRelations */


// An array of labeled contact relations for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/contactRelations
func (c_ CNMutableContact) SetContactRelations(value []CNLabeledValue) {
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
}/* debug [instance_properties/setter]: contactRelations */


// An enum identifying the contact type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/contactType
func (c_ CNMutableContact) ContactType() CNContactType {
	rv := objc.Send[CNContactType](c_.ID, objc.Sel("contactType"))
	return rv
}/* debug [instance_properties/getter]: contactType */


// An enum identifying the contact type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/contactType
func (c_ CNMutableContact) SetContactType(value CNContactType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactType:"), value)
}/* debug [instance_properties/setter]: contactType */


// An array containing labeled Gregorian dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/dates
func (c_ CNMutableContact) Dates() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("dates"))
	return rv
}/* debug [instance_properties/getter]: dates */


// An array containing labeled Gregorian dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/dates
func (c_ CNMutableContact) SetDates(value []CNLabeledValue) {
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
}/* debug [instance_properties/setter]: dates */


// The name of the department associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/departmentName
func (c_ CNMutableContact) DepartmentName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("departmentName"))
	return rv
}/* debug [instance_properties/getter]: departmentName */


// The name of the department associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/departmentName
func (c_ CNMutableContact) SetDepartmentName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDepartmentName:"), value)
}/* debug [instance_properties/setter]: departmentName */


// An array of labeled email addresses for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/emailAddresses
func (c_ CNMutableContact) EmailAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("emailAddresses"))
	return rv
}/* debug [instance_properties/getter]: emailAddresses */


// An array of labeled email addresses for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/emailAddresses
func (c_ CNMutableContact) SetEmailAddresses(value []CNLabeledValue) {
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
}/* debug [instance_properties/setter]: emailAddresses */


// The family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/familyName
func (c_ CNMutableContact) FamilyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("familyName"))
	return rv
}/* debug [instance_properties/getter]: familyName */


// The family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/familyName
func (c_ CNMutableContact) SetFamilyName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFamilyName:"), value)
}/* debug [instance_properties/setter]: familyName */


// The given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/givenName
func (c_ CNMutableContact) GivenName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("givenName"))
	return rv
}/* debug [instance_properties/getter]: givenName */


// The given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/givenName
func (c_ CNMutableContact) SetGivenName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGivenName:"), value)
}/* debug [instance_properties/setter]: givenName */


// The profile picture of a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/imageData
func (c_ CNMutableContact) ImageData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("imageData"))
	return rv
}/* debug [instance_properties/getter]: imageData */


// The profile picture of a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/imageData
func (c_ CNMutableContact) SetImageData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageData:"), value)
}/* debug [instance_properties/setter]: imageData */


// An array of labeled IM addresses for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/instantMessageAddresses
func (c_ CNMutableContact) InstantMessageAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("instantMessageAddresses"))
	return rv
}/* debug [instance_properties/getter]: instantMessageAddresses */


// An array of labeled IM addresses for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/instantMessageAddresses
func (c_ CNMutableContact) SetInstantMessageAddresses(value []CNLabeledValue) {
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
}/* debug [instance_properties/setter]: instantMessageAddresses */


// The contact’s job title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/jobTitle
func (c_ CNMutableContact) JobTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("jobTitle"))
	return rv
}/* debug [instance_properties/getter]: jobTitle */


// The contact’s job title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/jobTitle
func (c_ CNMutableContact) SetJobTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setJobTitle:"), value)
}/* debug [instance_properties/setter]: jobTitle */


// The middle name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/middleName
func (c_ CNMutableContact) MiddleName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("middleName"))
	return rv
}/* debug [instance_properties/getter]: middleName */


// The middle name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/middleName
func (c_ CNMutableContact) SetMiddleName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMiddleName:"), value)
}/* debug [instance_properties/setter]: middleName */


// The name prefix of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/namePrefix
func (c_ CNMutableContact) NamePrefix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("namePrefix"))
	return rv
}/* debug [instance_properties/getter]: namePrefix */


// The name prefix of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/namePrefix
func (c_ CNMutableContact) SetNamePrefix(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNamePrefix:"), value)
}/* debug [instance_properties/setter]: namePrefix */


// The name suffix of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/nameSuffix
func (c_ CNMutableContact) NameSuffix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("nameSuffix"))
	return rv
}/* debug [instance_properties/getter]: nameSuffix */


// The name suffix of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/nameSuffix
func (c_ CNMutableContact) SetNameSuffix(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNameSuffix:"), value)
}/* debug [instance_properties/setter]: nameSuffix */


// The nickname of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/nickname
func (c_ CNMutableContact) Nickname() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("nickname"))
	return rv
}/* debug [instance_properties/getter]: nickname */


// The nickname of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/nickname
func (c_ CNMutableContact) SetNickname(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNickname:"), value)
}/* debug [instance_properties/setter]: nickname */


// A date component for the non-Gregorian birthday of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/nonGregorianBirthday
func (c_ CNMutableContact) NonGregorianBirthday() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](c_.ID, objc.Sel("nonGregorianBirthday"))
	return rv
}/* debug [instance_properties/getter]: nonGregorianBirthday */


// A date component for the non-Gregorian birthday of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/nonGregorianBirthday
func (c_ CNMutableContact) SetNonGregorianBirthday(value foundation.DateComponents) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNonGregorianBirthday:"), value)
}/* debug [instance_properties/setter]: nonGregorianBirthday */


// A string containing notes for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/note
func (c_ CNMutableContact) Note() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("note"))
	return rv
}/* debug [instance_properties/getter]: note */


// A string containing notes for the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/note
func (c_ CNMutableContact) SetNote(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNote:"), value)
}/* debug [instance_properties/setter]: note */


// The name of the organization associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/organizationName
func (c_ CNMutableContact) OrganizationName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("organizationName"))
	return rv
}/* debug [instance_properties/getter]: organizationName */


// The name of the organization associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/organizationName
func (c_ CNMutableContact) SetOrganizationName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrganizationName:"), value)
}/* debug [instance_properties/setter]: organizationName */


// An array of labeled phone numbers for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneNumbers
func (c_ CNMutableContact) PhoneNumbers() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("phoneNumbers"))
	return rv
}/* debug [instance_properties/getter]: phoneNumbers */


// An array of labeled phone numbers for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneNumbers
func (c_ CNMutableContact) SetPhoneNumbers(value []CNLabeledValue) {
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
}/* debug [instance_properties/setter]: phoneNumbers */


// The phonetic family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticFamilyName
func (c_ CNMutableContact) PhoneticFamilyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("phoneticFamilyName"))
	return rv
}/* debug [instance_properties/getter]: phoneticFamilyName */


// The phonetic family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticFamilyName
func (c_ CNMutableContact) SetPhoneticFamilyName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneticFamilyName:"), value)
}/* debug [instance_properties/setter]: phoneticFamilyName */


// The phonetic given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticGivenName
func (c_ CNMutableContact) PhoneticGivenName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("phoneticGivenName"))
	return rv
}/* debug [instance_properties/getter]: phoneticGivenName */


// The phonetic given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticGivenName
func (c_ CNMutableContact) SetPhoneticGivenName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneticGivenName:"), value)
}/* debug [instance_properties/setter]: phoneticGivenName */


// The phonetic middle name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticMiddleName
func (c_ CNMutableContact) PhoneticMiddleName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("phoneticMiddleName"))
	return rv
}/* debug [instance_properties/getter]: phoneticMiddleName */


// The phonetic middle name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticMiddleName
func (c_ CNMutableContact) SetPhoneticMiddleName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneticMiddleName:"), value)
}/* debug [instance_properties/setter]: phoneticMiddleName */


// The phonetic name of the organization associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticOrganizationName
func (c_ CNMutableContact) PhoneticOrganizationName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("phoneticOrganizationName"))
	return rv
}/* debug [instance_properties/getter]: phoneticOrganizationName */


// The phonetic name of the organization associated with the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/phoneticOrganizationName
func (c_ CNMutableContact) SetPhoneticOrganizationName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneticOrganizationName:"), value)
}/* debug [instance_properties/setter]: phoneticOrganizationName */


// An array of labeled postal addresses for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/postalAddresses
func (c_ CNMutableContact) PostalAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("postalAddresses"))
	return rv
}/* debug [instance_properties/getter]: postalAddresses */


// An array of labeled postal addresses for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/postalAddresses
func (c_ CNMutableContact) SetPostalAddresses(value []CNLabeledValue) {
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
}/* debug [instance_properties/setter]: postalAddresses */


// The previous family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/previousFamilyName
func (c_ CNMutableContact) PreviousFamilyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("previousFamilyName"))
	return rv
}/* debug [instance_properties/getter]: previousFamilyName */


// The previous family name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/previousFamilyName
func (c_ CNMutableContact) SetPreviousFamilyName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviousFamilyName:"), value)
}/* debug [instance_properties/setter]: previousFamilyName */


// An array of labeled social profiles for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/socialProfiles
func (c_ CNMutableContact) SocialProfiles() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("socialProfiles"))
	return rv
}/* debug [instance_properties/getter]: socialProfiles */


// An array of labeled social profiles for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/socialProfiles
func (c_ CNMutableContact) SetSocialProfiles(value []CNLabeledValue) {
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
}/* debug [instance_properties/setter]: socialProfiles */


// An array of labeled URL addresses for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/urlAddresses
func (c_ CNMutableContact) UrlAddresses() []CNLabeledValue {
	rv := objc.Send[[]CNLabeledValue](c_.ID, objc.Sel("urlAddresses"))
	return rv
}/* debug [instance_properties/getter]: urlAddresses */


// An array of labeled URL addresses for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableContact/urlAddresses
func (c_ CNMutableContact) SetUrlAddresses(value []CNLabeledValue) {
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
}/* debug [instance_properties/setter]: urlAddresses */


// Exception thrown when an accessed property was not fetched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactpropertynotfetchedexceptionname
func (c_ CNMutableContact) CNContactPropertyNotFetchedExceptionName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNContactPropertyNotFetchedExceptionName"))
	return rv
}/* debug [instance_properties/getter]: CNContactPropertyNotFetchedExceptionName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNMutableContact */



