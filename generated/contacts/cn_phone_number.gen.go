// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNPhoneNumber */


/* debug [class_header]: Header for CNPhoneNumber */
// The class instance for the [CNPhoneNumber] class.
var (
	CNPhoneNumberClass     _CNPhoneNumberClass
	CNPhoneNumberClassOnce sync.Once
)

func getCNPhoneNumberClass() _CNPhoneNumberClass {
	CNPhoneNumberClassOnce.Do(func() {
		CNPhoneNumberClass = _CNPhoneNumberClass{objc.GetClass("CNPhoneNumber")}
	})
	return CNPhoneNumberClass
}

type _CNPhoneNumberClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNPhoneNumber */
// An interface definition for the [CNPhoneNumber] class.
type ICNPhoneNumber interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNPhoneNumber */
	// properties:
	StringValue() objc.IObject /* cross-framework: NSString */
	CNContactPhoneNumbersKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNPhoneNumber */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNPhoneNumber */
// Alloc allocates a new instance without initialization.
func (cc _CNPhoneNumberClass) Alloc() CNPhoneNumber {
	rv := objc.Send[CNPhoneNumber](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNPhoneNumberClass) New() CNPhoneNumber {
	rv := objc.Send[CNPhoneNumber](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNPhoneNumber) Init() CNPhoneNumber {
	rv := objc.Send[CNPhoneNumber](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNPhoneNumber) Autorelease() CNPhoneNumber {
	rv := objc.Send[CNPhoneNumber](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNPhoneNumber creates a new CNPhoneNumber instance.
func NewCNPhoneNumber() CNPhoneNumber {
	return getCNPhoneNumberClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNPhoneNumber */
// An immutable object representing a phone number for a contact.
//
// objects are thread-safe, and you may access their properties from any thread of your app.


// An immutable object representing a phone number for a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPhoneNumber
type CNPhoneNumber struct {
	objectivec.Object
}

// CNPhoneNumberFrom constructs a [CNPhoneNumber] from an unsafe.Pointer.
//
// An immutable object representing a phone number for a contact.
func CNPhoneNumberFrom(ptr unsafe.Pointer) CNPhoneNumber {
	return CNPhoneNumber{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNPhoneNumber */

// Returns a new phone number object initialized with the specified phone number string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPhoneNumber/init(stringValue:)
func NewCNPhoneNumberWithStringValue(string_ objc.IObject /* cross-framework: NSString */) CNPhoneNumber {
	instance := getCNPhoneNumberClass().Alloc()
	rv := objc.Send[CNPhoneNumber](instance.ID, objc.Sel("initWithStringValue:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNPhoneNumberWithStringValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNPhoneNumber */

// Returns a new phone number object initialized with the specified phone number string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPhoneNumber/phoneNumberWithStringValue:
func (cc _CNPhoneNumberClass) PhoneNumberWithStringValue(stringValue objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("phoneNumberWithStringValue:"), stringValue)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PhoneNumberWithStringValue) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNPhoneNumber */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNPhoneNumber */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNPhoneNumber */

// The string value of the phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNPhoneNumber/stringValue
func (c_ CNPhoneNumber) StringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("stringValue"))
	return rv
}/* debug [instance_properties/getter]: stringValue */


// A phone numbers of a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactphonenumberskey
func (c_ CNPhoneNumber) CNContactPhoneNumbersKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNContactPhoneNumbersKey"))
	return rv
}/* debug [instance_properties/getter]: CNContactPhoneNumbersKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNPhoneNumber */


