// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXPCListener

// ExampleNewXPCListenerWithMachServiceName demonstrates how to create a XPCListener instance using NewXPCListenerWithMachServiceName.
// Initializes a listener in a LaunchAgent or LaunchDaemon which has a name advertised in a   file.
func ExampleNewXPCListenerWithMachServiceName() {
	_ = foundation.NewXPCListenerWithMachServiceName(
		"name", // name string
	)
	// Output:
}
