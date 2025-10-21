// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit_test

import (
	"github.com/tmc/appledocs/generated/fskit"
)

// Suppress unused import errors
var _ = fskit.NewFSFileName

// ExampleNewFSFileNameWithString demonstrates how to create a FSFileName instance using NewFSFileNameWithString.
// Creates a filename by copying a character sequence from a string instance.
func ExampleNewFSFileNameWithString() {
	_ = fskit.NewFSFileNameWithString(
		"name", // name string
	)
	// Output:
}
