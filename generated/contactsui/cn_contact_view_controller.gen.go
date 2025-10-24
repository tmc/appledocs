// Code generated from Apple documentation for ContactsUI. DO NOT EDIT.

package contactsui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/contacts"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNContactViewController */


/* debug [class_header]: Header for CNContactViewController */
// The class instance for the [CNContactViewController] class.
var (
	CNContactViewControllerClass     _CNContactViewControllerClass
	CNContactViewControllerClassOnce sync.Once
)

func getCNContactViewControllerClass() _CNContactViewControllerClass {
	CNContactViewControllerClassOnce.Do(func() {
		CNContactViewControllerClass = _CNContactViewControllerClass{objc.GetClass("CNContactViewController")}
	})
	return CNContactViewControllerClass
}

type _CNContactViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNContactViewController */
// An interface definition for the [CNContactViewController] class.
type ICNContactViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for CNContactViewController */
	// properties:
	Contact() contacts.CNContact
	SetContact(value contacts.CNContact)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNContactViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNContactViewController */
// Alloc allocates a new instance without initialization.
func (cc _CNContactViewControllerClass) Alloc() CNContactViewController {
	rv := objc.Send[CNContactViewController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNContactViewControllerClass) New() CNContactViewController {
	rv := objc.Send[CNContactViewController](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContactViewController) Init() CNContactViewController {
	rv := objc.Send[CNContactViewController](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContactViewController) Autorelease() CNContactViewController {
	rv := objc.Send[CNContactViewController](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContactViewController creates a new CNContactViewController instance.
func NewCNContactViewController() CNContactViewController {
	return getCNContactViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNContactViewController */
// A view controller that displays a new, unknown, or existing contact.
//
// Present a object when you want to display information about one of the user’s contacts. At creation time, you specify the type of contact you want to display: new, unknown, or existing.


// A view controller that displays a new, unknown, or existing contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController
type CNContactViewController struct {
	ViewController
}

// CNContactViewControllerFrom constructs a [CNContactViewController] from an unsafe.Pointer.
//
// A view controller that displays a new, unknown, or existing contact.
func CNContactViewControllerFrom(ptr unsafe.Pointer) CNContactViewController {
	return CNContactViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNContactViewController */

// Initializes a view controller for an existing contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(for:)
func NewCNContactViewControllerForContact(contact contacts.CNContact) CNContactViewController {
	rv := objc.Send[CNContactViewController](objc.ID(getCNContactViewControllerClass().class), objc.Sel("viewControllerForContact:"), contact)
	return rv
}/* debug [class_init_methods/constructor]: NewCNContactViewControllerForContact */


// Initializes a view controller for a new contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(forNewContact:)
func NewCNContactViewControllerForNewContact(contact contacts.CNContact) CNContactViewController {
	rv := objc.Send[CNContactViewController](objc.ID(getCNContactViewControllerClass().class), objc.Sel("viewControllerForNewContact:"), contact)
	return rv
}/* debug [class_init_methods/constructor]: NewCNContactViewControllerForNewContact */


// Initializes a view controller for an unknown contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(forUnknownContact:)
func NewCNContactViewControllerForUnknownContact(contact contacts.CNContact) CNContactViewController {
	rv := objc.Send[CNContactViewController](objc.ID(getCNContactViewControllerClass().class), objc.Sel("viewControllerForUnknownContact:"), contact)
	return rv
}/* debug [class_init_methods/constructor]: NewCNContactViewControllerForUnknownContact */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNContactViewController */

// Returns the descriptor for all the keys that must be fetched on the contact before setting it on the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/descriptorForRequiredKeys()
func (cc _CNContactViewControllerClass) DescriptorForRequiredKeys() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorForRequiredKeys"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorForRequiredKeys) */


// Initializes a view controller for an existing contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(for:)
func (cc _CNContactViewControllerClass) ViewControllerForContact(contact contacts.CNContact) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("viewControllerForContact:"), contact)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ViewControllerForContact) */


// Initializes a view controller for a new contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(forNewContact:)
func (cc _CNContactViewControllerClass) ViewControllerForNewContact(contact contacts.CNContact) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("viewControllerForNewContact:"), contact)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ViewControllerForNewContact) */


// Initializes a view controller for an unknown contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(forUnknownContact:)
func (cc _CNContactViewControllerClass) ViewControllerForUnknownContact(contact contacts.CNContact) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("viewControllerForUnknownContact:"), contact)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ViewControllerForUnknownContact) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNContactViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNContactViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNContactViewController */

// The contact being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/contact
func (c_ CNContactViewController) Contact() contacts.CNContact {
	rv := objc.Send[contacts.CNContact](c_.ID, objc.Sel("contact"))
	return rv
}/* debug [instance_properties/getter]: contact */


// The contact being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/contact
func (c_ CNContactViewController) SetContact(value contacts.CNContact) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContact:"), value)
}/* debug [instance_properties/setter]: contact */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNContactViewController */


