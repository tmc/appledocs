// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook_test

import (
	"github.com/tmc/appledocs/generated/addressbook"
)

// Suppress unused import errors
var _ = addressbook.NewABRecord

// ExampleNewABRecord demonstrates how to create a ABRecord instance.
// Initializes a record using the shared address book.
func ExampleNewABRecord() {
	_ = addressbook.NewABRecord()
	// Output:
}
// ExampleNewABRecordWithAddressBook demonstrates how to create a ABRecord instance using NewABRecordWithAddressBook.
// Initializes a record using the given address book.
func ExampleNewABRecordWithAddressBook() {
	_ = addressbook.NewABRecordWithAddressBook(
		addressbook.ABAddressBook{}, // addressBook ABAddressBook
	)
	// Output:
}
