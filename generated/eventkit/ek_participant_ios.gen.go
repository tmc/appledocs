//go:build darwin && ios

// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for EKParticipant


// Returns the address book record that represents the participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipant/abRecord(with:)
func (e_ EKParticipant) ABRecordWithAddressBook(addressBook ABAddressBookRef /* typedef */) ABRecordRef /* typedef */ {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("ABRecordWithAddressBook:"), addressBook)
	return rv
}

// iOS-only properties





