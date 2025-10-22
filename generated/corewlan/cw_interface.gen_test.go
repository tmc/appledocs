// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan_test

import (
	"github.com/tmc/appledocs/generated/corewlan"
)

// Suppress unused import errors
var _ = corewlan.NewCWInterface

// ExampleNewCWInterfaceWithInterfaceName demonstrates how to create a CWInterface instance using NewCWInterfaceWithInterfaceName.
// Convenience method for getting an CWInterface object with the specified name.
func ExampleNewCWInterfaceWithInterfaceName() {
	_ = corewlan.NewCWInterfaceWithInterfaceName(
		"name", // name string
	)
	// Output:
}
// ExampleNewCWInterfaceWithName demonstrates how to create a CWInterface instance using NewCWInterfaceWithName.
// An instance method for obtaining an CWInterface object.
func ExampleNewCWInterfaceWithName() {
	_ = corewlan.NewCWInterfaceWithName(
		"name", // name string
	)
	// Output:
}
