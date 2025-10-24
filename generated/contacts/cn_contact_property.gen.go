// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	GivenName() objc.IObject /* cross-framework: NSString */
	SetGivenName(value objc.IObject /* cross-framework: NSString */)
	JobTitle() objc.IObject /* cross-framework: NSString */
	SetJobTitle(value objc.IObject /* cross-framework: NSString */)
	PhoneNumbers() ICNPhoneNumber
	SetPhoneNumbers(value ICNPhoneNumber)
	Contact() ICNContact
	SetContact(value ICNContact)
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	Key() objc.IObject /* cross-framework: NSString */
	SetKey(value objc.IObject /* cross-framework: NSString */)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Value() unsafe.Pointer
	SetValue(value unsafe.Pointer)
	CNContactPropertyNotFetchedExceptionName() objc.IObject /* cross-framework: NSString */
	// methods:
}

// An object that represents a property of a contact.
//
// A contact (that is, an instance of ) has properties, such as , , and . Each property is represented by an instance of , which provides a tuple that can contain three or five values, depending on whether the property is a member of an array of labeled values. For example, the property is a member of an array of labeled values, so the tuple contains the contact, key, value, identifier, and label. For the property, which is not contained in a labeled array, returns a tuple that contains the contact, key, and value. The class is used by to return the user’s selected property.


// An object that represents a property of a contact.
//
// [Full Topic]
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



// The given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/givenname
func (c_ CNContactProperty) GivenName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("givenName"))
	return rv
}


// The given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/givenname
func (c_ CNContactProperty) SetGivenName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGivenName:"), value)
}


// The contact’s job title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/jobtitle
func (c_ CNContactProperty) JobTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("jobTitle"))
	return rv
}


// The contact’s job title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/jobtitle
func (c_ CNContactProperty) SetJobTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setJobTitle:"), value)
}


// An array of labeled phone numbers for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/phonenumbers
func (c_ CNContactProperty) PhoneNumbers() ICNPhoneNumber {
	rv := objc.Send[CNPhoneNumber](c_.ID, objc.Sel("phoneNumbers"))
	return rv
}


// An array of labeled phone numbers for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/phonenumbers
func (c_ CNContactProperty) SetPhoneNumbers(value ICNPhoneNumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneNumbers:"), value)
}


// The associated contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/contact
func (c_ CNContactProperty) Contact() ICNContact {
	rv := objc.Send[CNContact](c_.ID, objc.Sel("contact"))
	return rv
}


// The associated contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/contact
func (c_ CNContactProperty) SetContact(value ICNContact) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContact:"), value)
}


// The identifier of the labeled value in the array of labeled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/identifier
func (c_ CNContactProperty) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}


// The identifier of the labeled value in the array of labeled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/identifier
func (c_ CNContactProperty) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIdentifier:"), value)
}


// The key of the contact property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/key
func (c_ CNContactProperty) Key() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("key"))
	return rv
}


// The key of the contact property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/key
func (c_ CNContactProperty) SetKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKey:"), value)
}


// The label of the labeled value of the property array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/label
func (c_ CNContactProperty) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("label"))
	return rv
}


// The label of the labeled value of the property array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/label
func (c_ CNContactProperty) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabel:"), value)
}


// The value of the property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/value
func (c_ CNContactProperty) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("value"))
	return rv
}


// The value of the property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactproperty/value
func (c_ CNContactProperty) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setValue:"), value)
}


// Exception thrown when an accessed property was not fetched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactpropertynotfetchedexceptionname
func (c_ CNContactProperty) CNContactPropertyNotFetchedExceptionName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNContactPropertyNotFetchedExceptionName"))
	return rv
}



