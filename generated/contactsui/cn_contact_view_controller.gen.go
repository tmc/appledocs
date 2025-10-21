// Code generated from Apple documentation for ContactsUI. DO NOT EDIT.

package contactsui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/contacts"
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
	HighlightPropertyWithKeyIdentifier(key appkit.string, identifier appkit.string)
}

// A view controller that displays a new, unknown, or existing contact.
//
// Present a object when you want to display information about one of the user’s contacts. At creation time, you specify the type of contact you want to display: new, unknown, or existing.
//
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
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(for:)
func NewCNContactViewControllerForContact(contact contacts.ICNContact) CNContactViewController {
	rv := objc.Send[CNContactViewController](objc.ID(getCNContactViewControllerClass().class), objc.Sel("viewControllerForContact:"), contact)
	return rv
}



// Initializes a view controller for a new contact.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(forNewContact:)
func NewCNContactViewControllerForNewContact(contact contacts.ICNContact) CNContactViewController {
	rv := objc.Send[CNContactViewController](objc.ID(getCNContactViewControllerClass().class), objc.Sel("viewControllerForNewContact:"), contact)
	return rv
}



// Initializes a view controller for an unknown contact.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(forUnknownContact:)
func NewCNContactViewControllerForUnknownContact(contact contacts.ICNContact) CNContactViewController {
	rv := objc.Send[CNContactViewController](objc.ID(getCNContactViewControllerClass().class), objc.Sel("viewControllerForUnknownContact:"), contact)
	return rv
}


// Returns the descriptor for all the keys that must be fetched on the contact before setting it on the view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/descriptorForRequiredKeys()
func (cc _CNContactViewControllerClass) DescriptorForRequiredKeys() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("descriptorForRequiredKeys"))
	return rv
}

// Initializes a view controller for an existing contact.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(for:)
func (cc _CNContactViewControllerClass) ViewControllerForContact(contact contacts.ICNContact) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("viewControllerForContact:"), contact)
	return rv
}

// Initializes a view controller for a new contact.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(forNewContact:)
func (cc _CNContactViewControllerClass) ViewControllerForNewContact(contact contacts.ICNContact) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("viewControllerForNewContact:"), contact)
	return rv
}

// Initializes a view controller for an unknown contact.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/init(forUnknownContact:)
func (cc _CNContactViewControllerClass) ViewControllerForUnknownContact(contact contacts.ICNContact) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("viewControllerForUnknownContact:"), contact)
	return rv
}

// Highlights the property of the contact being displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/highlightProperty(withKey:identifier:)
func (c_ CNContactViewController) HighlightPropertyWithKeyIdentifier(key appkit.string, identifier appkit.string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("highlightPropertyWithKey:identifier:"), key, identifier)
}

// Determines whether to display buttons for actions such as sending a text message or initiating a FaceTime call.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/allowsActions
func (c_ CNContactViewController) AllowsActions() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsActions"))
	return rv
}


// SetAllowsActions sets the value of the allowsActions property.
// Determines whether to display buttons for actions such as sending a text message or initiating a FaceTime call.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/allowsActions
func (c_ CNContactViewController) SetAllowsActions(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsActions:"), value)
}

// Determines whether the user can edit the contact’s information.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/allowsEditing
func (c_ CNContactViewController) AllowsEditing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsEditing"))
	return rv
}


// SetAllowsEditing sets the value of the allowsEditing property.
// Determines whether the user can edit the contact’s information.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/allowsEditing
func (c_ CNContactViewController) SetAllowsEditing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsEditing:"), value)
}

// The name to use if the contact has no display name.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/alternateName
func (c_ CNContactViewController) AlternateName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("alternateName"))
	return rv
}


// SetAlternateName sets the value of the alternateName property.
// The name to use if the contact has no display name.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/alternateName
func (c_ CNContactViewController) SetAlternateName(value appkit.string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlternateName:"), value)
}

// The contact being displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/contact
func (c_ CNContactViewController) Contact() contacts.CNContact {
	rv := objc.Send[contacts.CNContact](c_.ID, objc.Sel("contact"))
	return rv
}

// The contact store from which the contact was fetched or to which it will be saved.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/contactStore
func (c_ CNContactViewController) ContactStore() contacts.CNContactStore {
	rv := objc.Send[contacts.CNContactStore](c_.ID, objc.Sel("contactStore"))
	return rv
}


// SetContactStore sets the value of the contactStore property.
// The contact store from which the contact was fetched or to which it will be saved.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/contactStore
func (c_ CNContactViewController) SetContactStore(value contacts.ICNContactStore) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactStore:"), value)
}

// The delegate to be notified.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/delegate
func (c_ CNContactViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate to be notified.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/delegate
func (c_ CNContactViewController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}

// The contact property keys to display.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/displayedPropertyKeys
func (c_ CNContactViewController) DisplayedPropertyKeys() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("displayedPropertyKeys"))
	return rv
}


// SetDisplayedPropertyKeys sets the value of the displayedPropertyKeys property.
// The contact property keys to display.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/displayedPropertyKeys
func (c_ CNContactViewController) SetDisplayedPropertyKeys(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDisplayedPropertyKeys:"), value)
}

// The message displayed under the name of the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/message
func (c_ CNContactViewController) Message() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("message"))
	return rv
}


// SetMessage sets the value of the message property.
// The message displayed under the name of the contact.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/message
func (c_ CNContactViewController) SetMessage(value appkit.string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMessage:"), value)
}

// The container in which to add a new contact.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/parentContainer
func (c_ CNContactViewController) ParentContainer() contacts.CNContainer {
	rv := objc.Send[contacts.CNContainer](c_.ID, objc.Sel("parentContainer"))
	return rv
}


// SetParentContainer sets the value of the parentContainer property.
// The container in which to add a new contact.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/parentContainer
func (c_ CNContactViewController) SetParentContainer(value contacts.ICNContainer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParentContainer:"), value)
}

// The group in which to add a new contact.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/parentGroup
func (c_ CNContactViewController) ParentGroup() contacts.CNGroup {
	rv := objc.Send[contacts.CNGroup](c_.ID, objc.Sel("parentGroup"))
	return rv
}


// SetParentGroup sets the value of the parentGroup property.
// The group in which to add a new contact.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/parentGroup
func (c_ CNContactViewController) SetParentGroup(value contacts.ICNGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParentGroup:"), value)
}

// Determines whether to display data from contacts that are linked to the contact being displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/shouldShowLinkedContacts
func (c_ CNContactViewController) ShouldShowLinkedContacts() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldShowLinkedContacts"))
	return rv
}


// SetShouldShowLinkedContacts sets the value of the shouldShowLinkedContacts property.
// Determines whether to display data from contacts that are linked to the contact being displayed.

//
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/shouldShowLinkedContacts
func (c_ CNContactViewController) SetShouldShowLinkedContacts(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldShowLinkedContacts:"), value)
}


