// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization_test

import (
	"github.com/tmc/appledocs/generated/virtualization"
)

// Suppress unused import errors
var _ = virtualization.NewVZSingleDirectoryShare

// ExampleNewVZSingleDirectoryShareWithDirectory demonstrates how to create a VZSingleDirectoryShare instance using NewVZSingleDirectoryShareWithDirectory.
// Creates a directory share with a directory that you specify on the host.
func ExampleNewVZSingleDirectoryShareWithDirectory() {
	_ = virtualization.NewVZSingleDirectoryShareWithDirectory(
		virtualization.VZSharedDirectory{}, // directory VZSharedDirectory
	)
	// Output:
}
