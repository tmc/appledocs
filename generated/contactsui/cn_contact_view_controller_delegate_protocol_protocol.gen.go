// Code generated from Apple documentation for ContactsUI. DO NOT EDIT.

package contactsui

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/contacts"
)

// PCNContactViewControllerDelegate is the CNContactViewControllerDelegate protocol interface.
//
// Methods you use to respond to user interactions with a contact view controller.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.Contacts-UI/documentation/ContactsUI/CNContactViewControllerDelegate
type PCNContactViewControllerDelegate interface {
	// Optional methods
	ContactViewControllerDidCompleteWithContact(viewController ICNContactViewController, contact contacts.CNContact)
	HasContactViewControllerDidCompleteWithContact() bool
	ContactViewControllerShouldPerformDefaultActionForContactProperty(viewController ICNContactViewController, property contacts.CNContactProperty) bool
	HasContactViewControllerShouldPerformDefaultActionForContactProperty() bool
}

// CNContactViewControllerDelegate is a delegate implementation builder for the PCNContactViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CNContactViewControllerDelegate struct {
	_ContactViewControllerDidCompleteWithContact func(viewController ICNContactViewController, contact contacts.CNContact)
	_ContactViewControllerShouldPerformDefaultActionForContactProperty func(viewController ICNContactViewController, property contacts.CNContactProperty) bool
}

// SetContactViewControllerDidCompleteWithContact sets the handler for the ContactViewControllerDidCompleteWithContact delegate method.
//
// Called when the view has been presented.
func (d *CNContactViewControllerDelegate) SetContactViewControllerDidCompleteWithContact(f func(viewController ICNContactViewController, contact contacts.CNContact)) {
	d._ContactViewControllerDidCompleteWithContact = f
}

// SetContactViewControllerShouldPerformDefaultActionForContactProperty sets the handler for the ContactViewControllerShouldPerformDefaultActionForContactProperty delegate method.
//
// Called when the user selects a property.
func (d *CNContactViewControllerDelegate) SetContactViewControllerShouldPerformDefaultActionForContactProperty(f func(viewController ICNContactViewController, property contacts.CNContactProperty) bool) {
	d._ContactViewControllerShouldPerformDefaultActionForContactProperty = f
}

// ContactViewControllerDidCompleteWithContact implements the PCNContactViewControllerDelegate interface.
func (d *CNContactViewControllerDelegate) ContactViewControllerDidCompleteWithContact(viewController ICNContactViewController, contact contacts.CNContact) {
	if d._ContactViewControllerDidCompleteWithContact != nil {
		d._ContactViewControllerDidCompleteWithContact(viewController, contact)
	}
}

// HasContactViewControllerDidCompleteWithContact returns true if a handler for ContactViewControllerDidCompleteWithContact has been set.
func (d *CNContactViewControllerDelegate) HasContactViewControllerDidCompleteWithContact() bool {
	return d._ContactViewControllerDidCompleteWithContact != nil
}

// ContactViewControllerShouldPerformDefaultActionForContactProperty implements the PCNContactViewControllerDelegate interface.
func (d *CNContactViewControllerDelegate) ContactViewControllerShouldPerformDefaultActionForContactProperty(viewController ICNContactViewController, property contacts.CNContactProperty) bool {
	if d._ContactViewControllerShouldPerformDefaultActionForContactProperty != nil {
		return d._ContactViewControllerShouldPerformDefaultActionForContactProperty(viewController, property)
	}
	var zero bool
	return zero
}

// HasContactViewControllerShouldPerformDefaultActionForContactProperty returns true if a handler for ContactViewControllerShouldPerformDefaultActionForContactProperty has been set.
func (d *CNContactViewControllerDelegate) HasContactViewControllerShouldPerformDefaultActionForContactProperty() bool {
	return d._ContactViewControllerShouldPerformDefaultActionForContactProperty != nil
}
