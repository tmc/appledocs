// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension_test

import (
	"github.com/tmc/appledocs/generated/networkextension"
)

// Suppress unused import errors
var _ = networkextension.NewNEFilterDataVerdict

// ExampleNewNEFilterDataVerdictWithPassBytesPeekBytes demonstrates how to create a NEFilterDataVerdict instance using NewNEFilterDataVerdictWithPassBytesPeekBytes.
// Creates a verdict that tells the system to pass a chunk of network data to its final destination, and specifies the next chunk of data to provide.
func ExampleNewNEFilterDataVerdictWithPassBytesPeekBytes() {
	_ = networkextension.NewNEFilterDataVerdictWithPassBytesPeekBytes(
		0, // passBytes uint
		0, // peekBytes uint
	)
	// Output:
}
