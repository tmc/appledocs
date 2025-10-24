// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNContactProperty */


/* debug [class_header]: Header for CNContactProperty */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNContactProperty */
// An interface definition for the [CNContactProperty] class.
type ICNContactProperty interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNContactProperty */
	// properties:
	Contact() ICNContact
	Identifier() objc.IObject /* cross-framework: NSString */
	Key() objc.IObject /* cross-framework: NSString */
	Label() objc.IObject /* cross-framework: NSString */
	Value() objc.ID
	GivenName() objc.IObject /* cross-framework: NSString */
	SetGivenName(value objc.IObject /* cross-framework: NSString */)
	JobTitle() objc.IObject /* cross-framework: NSString */
	SetJobTitle(value objc.IObject /* cross-framework: NSString */)
	PhoneNumbers() ICNPhoneNumber
	SetPhoneNumbers(value ICNPhoneNumber)
	CNContactPropertyNotFetchedExceptionName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNContactProperty */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNContactProperty */
// Alloc allocates a new instance without initialization.
func (cc _CNContactPropertyClass) Alloc() CNContactProperty {
	rv := objc.Send[CNContactProperty](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNContactProperty */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNContactProperty *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNContactProperty */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNContactProperty */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNContactProperty */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNContactProperty */

// The associated contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactProperty/contact
func (c_ CNContactProperty) Contact() ICNContact {
	rv := objc.Send[CNContact](c_.ID, objc.Sel("contact"))
	return rv
}/* debug [instance_properties/getter]: contact */


// The identifier of the labeled value in the array of labeled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactProperty/identifier
func (c_ CNContactProperty) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The key of the contact property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactProperty/key
func (c_ CNContactProperty) Key() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("key"))
	return rv
}/* debug [instance_properties/getter]: key */


// The label of the labeled value of the property array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactProperty/label
func (c_ CNContactProperty) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// The value of the property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactProperty/value
func (c_ CNContactProperty) Value() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/givenname
func (c_ CNContactProperty) GivenName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("givenName"))
	return rv
}/* debug [instance_properties/getter]: givenName */


// The given name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/givenname
func (c_ CNContactProperty) SetGivenName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGivenName:"), value)
}/* debug [instance_properties/setter]: givenName */


// The contact’s job title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/jobtitle
func (c_ CNContactProperty) JobTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("jobTitle"))
	return rv
}/* debug [instance_properties/getter]: jobTitle */


// The contact’s job title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/jobtitle
func (c_ CNContactProperty) SetJobTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setJobTitle:"), value)
}/* debug [instance_properties/setter]: jobTitle */


// An array of labeled phone numbers for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/phonenumbers
func (c_ CNContactProperty) PhoneNumbers() ICNPhoneNumber {
	rv := objc.Send[CNPhoneNumber](c_.ID, objc.Sel("phoneNumbers"))
	return rv
}/* debug [instance_properties/getter]: phoneNumbers */


// An array of labeled phone numbers for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/phonenumbers
func (c_ CNContactProperty) SetPhoneNumbers(value ICNPhoneNumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneNumbers:"), value)
}/* debug [instance_properties/setter]: phoneNumbers */


// Exception thrown when an accessed property was not fetched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactpropertynotfetchedexceptionname
func (c_ CNContactProperty) CNContactPropertyNotFetchedExceptionName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNContactPropertyNotFetchedExceptionName"))
	return rv
}/* debug [instance_properties/getter]: CNContactPropertyNotFetchedExceptionName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNContactProperty */



