// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider_test

import (
	"github.com/tmc/appledocs/generated/fileprovider"
)

// Suppress unused import errors
var _ = fileprovider.NewFileProviderDomain

// ExampleNewFileProviderDomainWithIdentifierDisplayName demonstrates how to create a FileProviderDomain instance using NewFileProviderDomainWithIdentifierDisplayName.
// Creates a new file provider domain with the specified identifier and display name.
func ExampleNewFileProviderDomainWithIdentifierDisplayName() {
	_ = fileprovider.NewFileProviderDomainWithIdentifierDisplayName(
		fileprovider.FileProviderDomainIdentifier{}, // identifier FileProviderDomainIdentifier
		"displayName", // displayName string
	)
	// Output:
}

// ExampleNewFileProviderDomainWithIdentifierDisplayNamePathRelativeToDocumentStorage demonstrates how to create a FileProviderDomain instance using NewFileProviderDomainWithIdentifierDisplayNamePathRelativeToDocumentStorage.
// Returns a newly instantiated domain.
func ExampleNewFileProviderDomainWithIdentifierDisplayNamePathRelativeToDocumentStorage() {
	_ = fileprovider.NewFileProviderDomainWithIdentifierDisplayNamePathRelativeToDocumentStorage(
		fileprovider.FileProviderDomainIdentifier{}, // identifier FileProviderDomainIdentifier
		"displayName", // displayName string
		"/tmp/test",   // pathRelativeToDocumentStorage string
	)
	// Output:
}
