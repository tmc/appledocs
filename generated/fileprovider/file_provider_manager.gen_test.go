// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider_test

import (
	"github.com/tmc/appledocs/generated/fileprovider"
)

// Suppress unused import errors
var _ = fileprovider.NewFileProviderManager

// ExampleNewFileProviderManagerForDomain demonstrates how to create a FileProviderManager instance using NewFileProviderManagerForDomain.
// Returns a newly created file provider manager for the specified domain.
func ExampleNewFileProviderManagerForDomain() {
	_ = fileprovider.NewFileProviderManagerForDomain(
		fileprovider.NSFileProviderDomain{}, // domain NSFileProviderDomain
	)
	// Output:
}
