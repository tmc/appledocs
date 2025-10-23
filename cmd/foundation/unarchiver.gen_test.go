// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewUnarchiver

// ExampleNewUnarchiverForReadingWithData demonstrates how to create a Unarchiver instance using NewUnarchiverForReadingWithData.
// Returns an   object initialized to read an archive from a given data object.
func ExampleNewUnarchiverForReadingWithData() {
	_ = foundation.NewUnarchiverForReadingWithData(
		foundation.NSData{}, // data NSData
	)
	// Output:
}
