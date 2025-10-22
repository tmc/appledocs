// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts_test

import (
	"github.com/tmc/appledocs/generated/contacts"
)

// Suppress unused import errors
var _ = contacts.NewCNInstantMessageAddress

// ExampleNewCNInstantMessageAddressWithUsernameService demonstrates how to create a CNInstantMessageAddress instance using NewCNInstantMessageAddressWithUsernameService.
// Returns a   object initialized with the specified user name and service.
func ExampleNewCNInstantMessageAddressWithUsernameService() {
	_ = contacts.NewCNInstantMessageAddressWithUsernameService(
		"username", // username string
		"service", // service string
	)
	// Output:
}
