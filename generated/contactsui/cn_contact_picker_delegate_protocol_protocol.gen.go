// Code generated from Apple documentation for ContactsUI. DO NOT EDIT.

package contactsui

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/contacts"
)

// PCNContactPickerDelegate is the CNContactPickerDelegate protocol interface.
//
// The methods that you implement to respond to contact-picker user events.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - visionOS 1.0+
//
// See: doc://com.apple.Contacts-UI/documentation/ContactsUI/CNContactPickerDelegate
type PCNContactPickerDelegate interface {
	// Optional methods
	ContactPickerDidSelectContactProperty(picker ICNContactPicker, contactProperty contacts.CNContactProperty)
	HasContactPickerDidSelectContactProperty() bool
	ContactPickerDidSelectContacts(picker ICNContactPickerViewController, contacts []contacts.CNContact)
	HasContactPickerDidSelectContacts() bool
	ContactPickerDidSelectContact(picker ICNContactPicker, contact contacts.CNContact)
	HasContactPickerDidSelectContact() bool
	ContactPickerDidSelectContactProperties(picker ICNContactPickerViewController, contactProperties []contacts.CNContactProperty)
	HasContactPickerDidSelectContactProperties() bool
	ContactPickerDidCancel(picker ICNContactPickerViewController)
	HasContactPickerDidCancel() bool
	ContactPickerDidClose(picker ICNContactPicker)
	HasContactPickerDidClose() bool
	ContactPickerWillClose(picker ICNContactPicker)
	HasContactPickerWillClose() bool
}

// CNContactPickerDelegate is a delegate implementation builder for the PCNContactPickerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CNContactPickerDelegate struct {
	_ContactPickerDidSelectContactProperty func(picker ICNContactPicker, contactProperty contacts.CNContactProperty)
	_ContactPickerDidSelectContacts func(picker ICNContactPickerViewController, contacts []contacts.CNContact)
	_ContactPickerDidSelectContact func(picker ICNContactPicker, contact contacts.CNContact)
	_ContactPickerDidSelectContactProperties func(picker ICNContactPickerViewController, contactProperties []contacts.CNContactProperty)
	_ContactPickerDidCancel func(picker ICNContactPickerViewController)
	_ContactPickerDidClose func(picker ICNContactPicker)
	_ContactPickerWillClose func(picker ICNContactPicker)
}

// SetContactPickerDidSelectContactProperty sets the handler for the ContactPickerDidSelectContactProperty delegate method.
//
// Called when a property of the contact has been selected by the user.
func (d *CNContactPickerDelegate) SetContactPickerDidSelectContactProperty(f func(picker ICNContactPicker, contactProperty contacts.CNContactProperty)) {
	d._ContactPickerDidSelectContactProperty = f
}

// SetContactPickerDidSelectContacts sets the handler for the ContactPickerDidSelectContacts delegate method.
//
// Called after contacts have been selected by the user.
func (d *CNContactPickerDelegate) SetContactPickerDidSelectContacts(f func(picker ICNContactPickerViewController, contacts []contacts.CNContact)) {
	d._ContactPickerDidSelectContacts = f
}

// SetContactPickerDidSelectContact sets the handler for the ContactPickerDidSelectContact delegate method.
//
// Called after a contact has been selected by the user.
func (d *CNContactPickerDelegate) SetContactPickerDidSelectContact(f func(picker ICNContactPicker, contact contacts.CNContact)) {
	d._ContactPickerDidSelectContact = f
}

// SetContactPickerDidSelectContactProperties sets the handler for the ContactPickerDidSelectContactProperties delegate method.
//
// Called after contact properties have been selected by the user.
func (d *CNContactPickerDelegate) SetContactPickerDidSelectContactProperties(f func(picker ICNContactPickerViewController, contactProperties []contacts.CNContactProperty)) {
	d._ContactPickerDidSelectContactProperties = f
}

// SetContactPickerDidCancel sets the handler for the ContactPickerDidCancel delegate method.
//
// In iOS, called when the user taps Cancel.
func (d *CNContactPickerDelegate) SetContactPickerDidCancel(f func(picker ICNContactPickerViewController)) {
	d._ContactPickerDidCancel = f
}

// SetContactPickerDidClose sets the handler for the ContactPickerDidClose delegate method.
//
// In macOS, called when the contact picker’s popover has closed.
func (d *CNContactPickerDelegate) SetContactPickerDidClose(f func(picker ICNContactPicker)) {
	d._ContactPickerDidClose = f
}

// SetContactPickerWillClose sets the handler for the ContactPickerWillClose delegate method.
//
// In macOS, called when the contact picker’s popover is about to close.
func (d *CNContactPickerDelegate) SetContactPickerWillClose(f func(picker ICNContactPicker)) {
	d._ContactPickerWillClose = f
}

// ContactPickerDidSelectContactProperty implements the PCNContactPickerDelegate interface.
func (d *CNContactPickerDelegate) ContactPickerDidSelectContactProperty(picker ICNContactPicker, contactProperty contacts.CNContactProperty) {
	if d._ContactPickerDidSelectContactProperty != nil {
		d._ContactPickerDidSelectContactProperty(picker, contactProperty)
	}
}

// HasContactPickerDidSelectContactProperty returns true if a handler for ContactPickerDidSelectContactProperty has been set.
func (d *CNContactPickerDelegate) HasContactPickerDidSelectContactProperty() bool {
	return d._ContactPickerDidSelectContactProperty != nil
}

// ContactPickerDidSelectContacts implements the PCNContactPickerDelegate interface.
func (d *CNContactPickerDelegate) ContactPickerDidSelectContacts(picker ICNContactPickerViewController, contacts []contacts.CNContact) {
	if d._ContactPickerDidSelectContacts != nil {
		d._ContactPickerDidSelectContacts(picker, contacts)
	}
}

// HasContactPickerDidSelectContacts returns true if a handler for ContactPickerDidSelectContacts has been set.
func (d *CNContactPickerDelegate) HasContactPickerDidSelectContacts() bool {
	return d._ContactPickerDidSelectContacts != nil
}

// ContactPickerDidSelectContact implements the PCNContactPickerDelegate interface.
func (d *CNContactPickerDelegate) ContactPickerDidSelectContact(picker ICNContactPicker, contact contacts.CNContact) {
	if d._ContactPickerDidSelectContact != nil {
		d._ContactPickerDidSelectContact(picker, contact)
	}
}

// HasContactPickerDidSelectContact returns true if a handler for ContactPickerDidSelectContact has been set.
func (d *CNContactPickerDelegate) HasContactPickerDidSelectContact() bool {
	return d._ContactPickerDidSelectContact != nil
}

// ContactPickerDidSelectContactProperties implements the PCNContactPickerDelegate interface.
func (d *CNContactPickerDelegate) ContactPickerDidSelectContactProperties(picker ICNContactPickerViewController, contactProperties []contacts.CNContactProperty) {
	if d._ContactPickerDidSelectContactProperties != nil {
		d._ContactPickerDidSelectContactProperties(picker, contactProperties)
	}
}

// HasContactPickerDidSelectContactProperties returns true if a handler for ContactPickerDidSelectContactProperties has been set.
func (d *CNContactPickerDelegate) HasContactPickerDidSelectContactProperties() bool {
	return d._ContactPickerDidSelectContactProperties != nil
}

// ContactPickerDidCancel implements the PCNContactPickerDelegate interface.
func (d *CNContactPickerDelegate) ContactPickerDidCancel(picker ICNContactPickerViewController) {
	if d._ContactPickerDidCancel != nil {
		d._ContactPickerDidCancel(picker)
	}
}

// HasContactPickerDidCancel returns true if a handler for ContactPickerDidCancel has been set.
func (d *CNContactPickerDelegate) HasContactPickerDidCancel() bool {
	return d._ContactPickerDidCancel != nil
}

// ContactPickerDidClose implements the PCNContactPickerDelegate interface.
func (d *CNContactPickerDelegate) ContactPickerDidClose(picker ICNContactPicker) {
	if d._ContactPickerDidClose != nil {
		d._ContactPickerDidClose(picker)
	}
}

// HasContactPickerDidClose returns true if a handler for ContactPickerDidClose has been set.
func (d *CNContactPickerDelegate) HasContactPickerDidClose() bool {
	return d._ContactPickerDidClose != nil
}

// ContactPickerWillClose implements the PCNContactPickerDelegate interface.
func (d *CNContactPickerDelegate) ContactPickerWillClose(picker ICNContactPicker) {
	if d._ContactPickerWillClose != nil {
		d._ContactPickerWillClose(picker)
	}
}

// HasContactPickerWillClose returns true if a handler for ContactPickerWillClose has been set.
func (d *CNContactPickerDelegate) HasContactPickerWillClose() bool {
	return d._ContactPickerWillClose != nil
}
