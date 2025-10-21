// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNContactProperty] class.
var (
	CNContactPropertyClass     _CNContactPropertyClass
	CNContactPropertyClassOnce sync.Once
)

func getCNContactPropertyClass() _CNContactPropertyClass {
	CNContactPropertyClassOnce.Do(func() {
		CNContactPropertyClass = _CNContactPropertyClass{objc.GetClass("CNContactProperty")}
	})
	return CNContactPropertyClass
}

type _CNContactPropertyClass struct {
	class objc.Class
}

// An interface definition for the [CNContactProperty] class.
type ICNContactProperty interface {
	objectivec.IObject
}

// An object that represents a property of a contact.
//
// A contact (that is, an instance of ) has properties, such as , , and . Each property is represented by an instance of , which provides a tuple that can contain three or five values, depending on whether the property is a member of an array of labeled values. For example, the property is a member of an array of labeled values, so the tuple contains the contact, key, value, identifier, and label. For the property, which is not contained in a labeled array, returns a tuple that contains the contact, key, and value. The class is used by to return the user’s selected property.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactProperty
type CNContactProperty struct {
	objectivec.Object
}

// CNContactPropertyFrom constructs a [CNContactProperty] from an unsafe.Pointer.
//
// An object that represents a property of a contact.
func CNContactPropertyFrom(ptr unsafe.Pointer) CNContactProperty {
	return CNContactProperty{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNContactPropertyClass) Alloc() CNContactProperty {
	rv := objc.Send[CNContactProperty](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNContactPropertyClass) New() CNContactProperty {
	rv := objc.Send[CNContactProperty](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContactProperty) Init() CNContactProperty {
	rv := objc.Send[CNContactProperty](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContactProperty) Autorelease() CNContactProperty {
	rv := objc.Send[CNContactProperty](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContactProperty creates a new CNContactProperty instance.
func NewCNContactProperty() CNContactProperty {
	return getCNContactPropertyClass().New()
}


// The associated contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactProperty/contact
func (c_ CNContactProperty) Contact() CNContact {
	rv := objc.Send[CNContact](c_.ID, objc.Sel("contact"))
	return rv
}

// The identifier of the labeled value in the array of labeled.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactProperty/identifier
func (c_ CNContactProperty) Identifier() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("identifier"))
	return rv
}

// The key of the contact property.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactProperty/key
func (c_ CNContactProperty) Key() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("key"))
	return rv
}

// The label of the labeled value of the property array.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactProperty/label
func (c_ CNContactProperty) Label() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("label"))
	return rv
}

// The value of the property.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactProperty/value
func (c_ CNContactProperty) Value() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("value"))
	return rv
}

// The given name of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/givenname
func (c_ CNContactProperty) GivenName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("givenName"))
	return rv
}


// SetGivenName sets the value of the givenName property.
// The given name of the contact.

//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/givenname
func (c_ CNContactProperty) SetGivenName(value appkit.string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGivenName:"), value)
}

// The contact’s job title.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/jobtitle
func (c_ CNContactProperty) JobTitle() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("jobTitle"))
	return rv
}


// SetJobTitle sets the value of the jobTitle property.
// The contact’s job title.

//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/jobtitle
func (c_ CNContactProperty) SetJobTitle(value appkit.string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setJobTitle:"), value)
}

// An array of labeled phone numbers for a contact.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/phonenumbers
func (c_ CNContactProperty) PhoneNumbers() CNPhoneNumber {
	rv := objc.Send[CNPhoneNumber](c_.ID, objc.Sel("phoneNumbers"))
	return rv
}


// SetPhoneNumbers sets the value of the phoneNumbers property.
// An array of labeled phone numbers for a contact.

//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/phonenumbers
func (c_ CNContactProperty) SetPhoneNumbers(value ICNPhoneNumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneNumbers:"), value)
}

// Exception thrown when an accessed property was not fetched.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactpropertynotfetchedexceptionname
func (c_ CNContactProperty) CNContactPropertyNotFetchedExceptionName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNContactPropertyNotFetchedExceptionName"))
	return rv
}



