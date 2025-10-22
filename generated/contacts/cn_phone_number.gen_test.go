// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts_test

import (
	"github.com/tmc/appledocs/generated/contacts"
)

// Suppress unused import errors
var _ = contacts.NewCNPhoneNumber

// ExampleNewCNPhoneNumber demonstrates how to create a CNPhoneNumber instance.
func ExampleNewCNPhoneNumber() {
	_ = contacts.NewCNPhoneNumber()
	// Output:
}
// ExampleNewCNPhoneNumberWithStringValue demonstrates how to create a CNPhoneNumber instance using NewCNPhoneNumberWithStringValue.
// Returns a new phone number object initialized with the specified phone number string.
func ExampleNewCNPhoneNumberWithStringValue() {
	_ = contacts.NewCNPhoneNumberWithStringValue(
		"string", // string string
	)
	// Output:
}
