// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PABPersonPickerDelegate is the ABPersonPickerDelegate protocol interface.
//
// Methods you use to respond to user selections in a person picker.
//
// Availability:
//   - macOS 10.9+
//
// See: doc://com.apple.addressbook/documentation/AddressBook/ABPersonPickerDelegate
type PABPersonPickerDelegate interface {
	// Required methods
	PersonPickerDidChoosePersonPropertyIdentifier(picker IABPersonPicker, person IABPerson, property objc.IObject /* cross-framework: NSString */, identifier objc.IObject /* cross-framework: NSString */)/* debug [protocol_interface/required_method]: PersonPickerDidChoosePersonPropertyIdentifier */
	PersonPickerDidClose(picker IABPersonPicker)/* debug [protocol_interface/required_method]: PersonPickerDidClose */
}

// ABPersonPickerDelegate is a delegate implementation builder for the PABPersonPickerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ABPersonPickerDelegate struct {
	_PersonPickerDidChoosePersonPropertyIdentifier func(picker IABPersonPicker, person IABPerson, property objc.IObject /* cross-framework: NSString */, identifier objc.IObject /* cross-framework: NSString */)
	_PersonPickerDidClose func(picker IABPersonPicker)
}

// SetPersonPickerDidChoosePersonPropertyIdentifier sets the handler for the PersonPickerDidChoosePersonPropertyIdentifier delegate method.
//
// Informs the delegate when the user selects a person, or a specific property of a person.
func (d *ABPersonPickerDelegate) SetPersonPickerDidChoosePersonPropertyIdentifier(f func(picker IABPersonPicker, person IABPerson, property objc.IObject /* cross-framework: NSString */, identifier objc.IObject /* cross-framework: NSString */)) {
	d._PersonPickerDidChoosePersonPropertyIdentifier = f
}

// SetPersonPickerDidClose sets the handler for the PersonPickerDidClose delegate method.
//
// Notifies the delegate when the user closes the  picker.
func (d *ABPersonPickerDelegate) SetPersonPickerDidClose(f func(picker IABPersonPicker)) {
	d._PersonPickerDidClose = f
}

// PersonPickerDidChoosePersonPropertyIdentifier implements the PABPersonPickerDelegate interface.
func (d *ABPersonPickerDelegate) PersonPickerDidChoosePersonPropertyIdentifier(picker IABPersonPicker, person IABPerson, property objc.IObject /* cross-framework: NSString */, identifier objc.IObject /* cross-framework: NSString */) {
	if d._PersonPickerDidChoosePersonPropertyIdentifier != nil {
		d._PersonPickerDidChoosePersonPropertyIdentifier(picker, person, property, identifier)
	}
}

// HasPersonPickerDidChoosePersonPropertyIdentifier returns true if a handler for PersonPickerDidChoosePersonPropertyIdentifier has been set.
func (d *ABPersonPickerDelegate) HasPersonPickerDidChoosePersonPropertyIdentifier() bool {
	return d._PersonPickerDidChoosePersonPropertyIdentifier != nil
}

// PersonPickerDidClose implements the PABPersonPickerDelegate interface.
func (d *ABPersonPickerDelegate) PersonPickerDidClose(picker IABPersonPicker) {
	if d._PersonPickerDidClose != nil {
		d._PersonPickerDidClose(picker)
	}
}

// HasPersonPickerDidClose returns true if a handler for PersonPickerDidClose has been set.
func (d *ABPersonPickerDelegate) HasPersonPickerDidClose() bool {
	return d._PersonPickerDidClose != nil
}
