// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKUserIdentityLookupInfo


// ExampleNewCKUserIdentityLookupInfoWithEmailAddress demonstrates how to create a CKUserIdentityLookupInfo instance using NewCKUserIdentityLookupInfoWithEmailAddress.
// Creates a lookup info for the specified email address.
func ExampleNewCKUserIdentityLookupInfoWithEmailAddress() {
	_ = cloudkit.NewCKUserIdentityLookupInfoWithEmailAddress(
		"emailAddress", // emailAddress string
	)
	// Output:
}

// ExampleNewCKUserIdentityLookupInfoWithPhoneNumber demonstrates how to create a CKUserIdentityLookupInfo instance using NewCKUserIdentityLookupInfoWithPhoneNumber.
// Creates a lookup info for the specified phone number.
func ExampleNewCKUserIdentityLookupInfoWithPhoneNumber() {
	_ = cloudkit.NewCKUserIdentityLookupInfoWithPhoneNumber(
		"phoneNumber", // phoneNumber string
	)
	// Output:
}



