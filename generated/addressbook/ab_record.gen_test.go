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
// ExampleABRecord_IsReadOnly demonstrates using IsReadOnly on a ABRecord instance.
// Returns whether a record is read-only.
func ExampleABRecord_IsReadOnly() {
	obj := addressbook.NewABRecord()
	_ = obj.IsReadOnly()
	// Output:
	}

