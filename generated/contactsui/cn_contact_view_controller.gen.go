// Code generated from Apple documentation for ContactsUI. DO NOT EDIT.

package contactsui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/contacts"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [CNContactViewController] class.
type ICNContactViewController interface {
	appkit.IViewController
	// properties:
	Contact() objc.IObject /* cross-framework: CNContact */
	// methods:
}

// A view controller that displays a new, unknown, or existing contact.
//
// Present a object when you want to display information about one of the user’s contacts. At creation time, you specify the type of contact you want to display: new, unknown, or existing.


// A view controller that displays a new, unknown, or existing contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController
type CNContactViewController struct {
	appkit.ViewController
}

// CNContactViewControllerFrom constructs a [CNContactViewController] from an unsafe.Pointer.
//
// A view controller that displays a new, unknown, or existing contact.
func CNContactViewControllerFrom(ptr unsafe.Pointer) CNContactViewController {
	return CNContactViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNContactViewControllerClass) Alloc() CNContactViewController {
	rv := objc.Send[CNContactViewController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Initializes a view controller for an existing contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(for:)
func NewCNContactViewControllerForContact(contact objc.IObject /* cross-framework: CNContact */) CNContactViewController {
	rv := objc.Send[CNContactViewController](objc.ID(getCNContactViewControllerClass().class), objc.Sel("viewControllerForContact:"), contact)
	return rv
}


// Initializes a view controller for a new contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(forNewContact:)
func NewCNContactViewControllerForNewContact(contact objc.IObject /* cross-framework: CNContact */) CNContactViewController {
	rv := objc.Send[CNContactViewController](objc.ID(getCNContactViewControllerClass().class), objc.Sel("viewControllerForNewContact:"), contact)
	return rv
}


// Initializes a view controller for an unknown contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(forUnknownContact:)
func NewCNContactViewControllerForUnknownContact(contact objc.IObject /* cross-framework: CNContact */) CNContactViewController {
	rv := objc.Send[CNContactViewController](objc.ID(getCNContactViewControllerClass().class), objc.Sel("viewControllerForUnknownContact:"), contact)
	return rv
}



// Returns the descriptor for all the keys that must be fetched on the contact before setting it on the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/descriptorForRequiredKeys()
func (cc _CNContactViewControllerClass) DescriptorForRequiredKeys() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("descriptorForRequiredKeys"))
	return rv
}


// Initializes a view controller for an existing contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(for:)
func (cc _CNContactViewControllerClass) ViewControllerForContact(contact objc.IObject /* cross-framework: CNContact */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("viewControllerForContact:"), contact)
	return rv
}


// Initializes a view controller for a new contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(forNewContact:)
func (cc _CNContactViewControllerClass) ViewControllerForNewContact(contact objc.IObject /* cross-framework: CNContact */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("viewControllerForNewContact:"), contact)
	return rv
}


// Initializes a view controller for an unknown contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(forUnknownContact:)
func (cc _CNContactViewControllerClass) ViewControllerForUnknownContact(contact objc.IObject /* cross-framework: CNContact */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("viewControllerForUnknownContact:"), contact)
	return rv
}


// The contact being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/contact
func (c_ CNContactViewController) Contact() objc.IObject /* cross-framework: CNContact */ {
	rv := objc.Send[contacts.CNContact](c_.ID, objc.Sel("contact"))
	return rv
}


