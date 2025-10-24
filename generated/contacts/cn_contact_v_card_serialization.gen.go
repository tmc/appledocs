// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNContactVCardSerialization */


/* debug [class_header]: Header for CNContactVCardSerialization */
// The class instance for the [CNContactVCardSerialization] class.
var (
	CNContactVCardSerializationClass     _CNContactVCardSerializationClass
	CNContactVCardSerializationClassOnce sync.Once
)

func getCNContactVCardSerializationClass() _CNContactVCardSerializationClass {
	CNContactVCardSerializationClassOnce.Do(func() {
		CNContactVCardSerializationClass = _CNContactVCardSerializationClass{objc.GetClass("CNContactVCardSerialization")}
	})
	return CNContactVCardSerializationClass
}

type _CNContactVCardSerializationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNContactVCardSerialization */
// An interface definition for the [CNContactVCardSerialization] class.
type ICNContactVCardSerialization interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNContactVCardSerialization */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNContactVCardSerialization */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNContactVCardSerialization */
// Alloc allocates a new instance without initialization.
func (cc _CNContactVCardSerializationClass) Alloc() CNContactVCardSerialization {
	rv := objc.Send[CNContactVCardSerialization](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNContactVCardSerializationClass) New() CNContactVCardSerialization {
	rv := objc.Send[CNContactVCardSerialization](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContactVCardSerialization) Init() CNContactVCardSerialization {
	rv := objc.Send[CNContactVCardSerialization](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContactVCardSerialization) Autorelease() CNContactVCardSerialization {
	rv := objc.Send[CNContactVCardSerialization](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContactVCardSerialization creates a new CNContactVCardSerialization instance.
func NewCNContactVCardSerialization() CNContactVCardSerialization {
	return getCNContactVCardSerializationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNContactVCardSerialization */
// An object you use to convert to and from a vCard representation of the user’s contacts.


// An object you use to convert to and from a vCard representation of the user’s contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactVCardSerialization
type CNContactVCardSerialization struct {
	objectivec.Object
}

// CNContactVCardSerializationFrom constructs a [CNContactVCardSerialization] from an unsafe.Pointer.
//
// An object you use to convert to and from a vCard representation of the user’s contacts.
func CNContactVCardSerializationFrom(ptr unsafe.Pointer) CNContactVCardSerialization {
	return CNContactVCardSerialization{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNContactVCardSerialization *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNContactVCardSerialization */

// Returns the contacts from the vCard data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactVCardSerialization/contacts(with:)
func (cc _CNContactVCardSerializationClass) ContactsWithDataError(data objc.IObject /* cross-framework: NSData */, error_ objectivec.IObject) []CNContact {
	rv := objc.Send[[]CNContact](objc.ID(cc.class), objc.Sel("contactsWithData:error:"), data, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContactsWithDataError) */


// Returns the vCard representation of the specified contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactVCardSerialization/data(with:)
func (cc _CNContactVCardSerializationClass) DataWithContactsError(contacts []CNContact, error_ objectivec.IObject) foundation.Data {
	rv := objc.Send[foundation.Data](objc.ID(cc.class), objc.Sel("dataWithContacts:error:"), contacts, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataWithContactsError) */


// Use to fetch all contact keys required to create vCard data from a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactVCardSerialization/descriptorForRequiredKeys()
func (cc _CNContactVCardSerializationClass) DescriptorForRequiredKeys() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorForRequiredKeys"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorForRequiredKeys) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNContactVCardSerialization */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNContactVCardSerialization */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNContactVCardSerialization */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNContactVCardSerialization */



