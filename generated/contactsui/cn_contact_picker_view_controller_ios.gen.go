//go:build darwin && ios

// Code generated from Apple documentation for ContactsUI. DO NOT EDIT.

package contactsui

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CNContactPickerViewController


// iOS-only properties

// The delegate to be notified when the user selects a contact or a property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/delegate
func (c_ CNContactPickerViewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}
func (c_ CNContactPickerViewController) SetDelegate(value unsafe.Pointer) {
	c_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// The property keys to display in the contact detail card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/displayedPropertyKeys
func (c_ CNContactPickerViewController) DisplayedPropertyKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("displayedPropertyKeys"))
	return rv
}
func (c_ CNContactPickerViewController) SetDisplayedPropertyKeys(value []string) {
	c_.ID.Send(objc.RegisterName("setDisplayedPropertyKeys:"), value)
}

// A predicate to determine the contact selectability in the list of contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/predicateForEnablingContact
func (c_ CNContactPickerViewController) PredicateForEnablingContact() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](c_.ID, objc.Sel("predicateForEnablingContact"))
	return rv
}
func (c_ CNContactPickerViewController) SetPredicateForEnablingContact(value foundation.Predicate) {
	c_.ID.Send(objc.RegisterName("setPredicateForEnablingContact:"), value)
}

// A predicate to control the return of the selected contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/predicateForSelectionOfContact
func (c_ CNContactPickerViewController) PredicateForSelectionOfContact() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](c_.ID, objc.Sel("predicateForSelectionOfContact"))
	return rv
}
func (c_ CNContactPickerViewController) SetPredicateForSelectionOfContact(value foundation.Predicate) {
	c_.ID.Send(objc.RegisterName("setPredicateForSelectionOfContact:"), value)
}

// A predicate to control the properties of the selected contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ContactsUI/CNContactPickerViewController/predicateForSelectionOfProperty
func (c_ CNContactPickerViewController) PredicateForSelectionOfProperty() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](c_.ID, objc.Sel("predicateForSelectionOfProperty"))
	return rv
}
func (c_ CNContactPickerViewController) SetPredicateForSelectionOfProperty(value foundation.Predicate) {
	c_.ID.Send(objc.RegisterName("setPredicateForSelectionOfProperty:"), value)
}





