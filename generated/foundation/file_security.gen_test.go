// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewFileSecurity

// ExampleNewFileSecurityWithCoder demonstrates how to create a FileSecurity instance using NewFileSecurityWithCoder.
func ExampleNewFileSecurityWithCoder() {
	_ = foundation.NewFileSecurityWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
