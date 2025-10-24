//go:build darwin && ios

// Code generated from Apple documentation for ContactsUI. DO NOT EDIT.

package contactsui

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/contacts"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CNContactViewController


// Highlights the property of the contact being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/highlightProperty(withKey:identifier:)
func (c_ CNContactViewController) HighlightPropertyWithKeyIdentifier(key objc.IObject /* cross-framework: NSString */, identifier objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("highlightPropertyWithKey:identifier:"), key, identifier)
}

// iOS-only properties

// Determines whether to display buttons for actions such as sending a text message or initiating a FaceTime call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/allowsActions
func (c_ CNContactViewController) AllowsActions() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsActions"))
	return rv
}
func (c_ CNContactViewController) SetAllowsActions(value bool) {
	c_.ID.Send(objc.RegisterName("setAllowsActions:"), value)
}

// Determines whether the user can edit the contact’s information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/allowsEditing
func (c_ CNContactViewController) AllowsEditing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsEditing"))
	return rv
}
func (c_ CNContactViewController) SetAllowsEditing(value bool) {
	c_.ID.Send(objc.RegisterName("setAllowsEditing:"), value)
}

// The name to use if the contact has no display name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/alternateName
func (c_ CNContactViewController) AlternateName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("alternateName"))
	return rv
}
func (c_ CNContactViewController) SetAlternateName(value objc.IObject /* cross-framework: NSString */) {
	c_.ID.Send(objc.RegisterName("setAlternateName:"), value)
}

// The contact store from which the contact was fetched or to which it will be saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/contactStore
func (c_ CNContactViewController) ContactStore() objc.IObject /* cross-framework: CNContactStore */ {
	rv := objc.Send[contacts.CNContactStore](c_.ID, objc.Sel("contactStore"))
	return rv
}
func (c_ CNContactViewController) SetContactStore(value objc.IObject /* cross-framework: CNContactStore */) {
	c_.ID.Send(objc.RegisterName("setContactStore:"), value)
}

// The delegate to be notified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/delegate
func (c_ CNContactViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}
func (c_ CNContactViewController) SetDelegate(value objc.ID) {
	c_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// The contact property keys to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/displayedPropertyKeys
func (c_ CNContactViewController) DisplayedPropertyKeys() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](c_.ID, objc.Sel("displayedPropertyKeys"))
	return rv
}
func (c_ CNContactViewController) SetDisplayedPropertyKeys(value objc.IObject /* cross-framework: NSArray */) {
	c_.ID.Send(objc.RegisterName("setDisplayedPropertyKeys:"), value)
}

// The message displayed under the name of the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/message
func (c_ CNContactViewController) Message() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("message"))
	return rv
}
func (c_ CNContactViewController) SetMessage(value objc.IObject /* cross-framework: NSString */) {
	c_.ID.Send(objc.RegisterName("setMessage:"), value)
}

// The container in which to add a new contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/parentContainer
func (c_ CNContactViewController) ParentContainer() objc.IObject /* cross-framework: CNContainer */ {
	rv := objc.Send[contacts.CNContainer](c_.ID, objc.Sel("parentContainer"))
	return rv
}
func (c_ CNContactViewController) SetParentContainer(value objc.IObject /* cross-framework: CNContainer */) {
	c_.ID.Send(objc.RegisterName("setParentContainer:"), value)
}

// The group in which to add a new contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/parentGroup
func (c_ CNContactViewController) ParentGroup() objc.IObject /* cross-framework: CNGroup */ {
	rv := objc.Send[contacts.CNGroup](c_.ID, objc.Sel("parentGroup"))
	return rv
}
func (c_ CNContactViewController) SetParentGroup(value objc.IObject /* cross-framework: CNGroup */) {
	c_.ID.Send(objc.RegisterName("setParentGroup:"), value)
}

// Determines whether to display data from contacts that are linked to the contact being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactViewController/shouldShowLinkedContacts
func (c_ CNContactViewController) ShouldShowLinkedContacts() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldShowLinkedContacts"))
	return rv
}
func (c_ CNContactViewController) SetShouldShowLinkedContacts(value bool) {
	c_.ID.Send(objc.RegisterName("setShouldShowLinkedContacts:"), value)
}




