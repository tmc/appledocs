// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider_test

import (
	"github.com/tmc/appledocs/generated/fileprovider"
)

// Suppress unused import errors
var _ = fileprovider.NewFileProviderKnownFolderLocation

// ExampleNewFileProviderKnownFolderLocationWithExistingItemIdentifier demonstrates how to create a FileProviderKnownFolderLocation instance using NewFileProviderKnownFolderLocationWithExistingItemIdentifier.
// Initialize a location with the item identifier of a folder that already exists on the server.
func ExampleNewFileProviderKnownFolderLocationWithExistingItemIdentifier() {
	_ = fileprovider.NewFileProviderKnownFolderLocationWithExistingItemIdentifier(
		fileprovider.FileProviderItemIdentifier /* typedef */{}, // existingItemIdentifier FileProviderItemIdentifier /* typedef */
	)
	// Output:
}
