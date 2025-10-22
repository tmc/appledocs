// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory_test

import (
	"github.com/tmc/appledocs/generated/opendirectory"
)

// Suppress unused import errors
var _ = opendirectory.NewODModuleEntry

// ExampleNewODModuleEntryWithNameXpcServiceName demonstrates how to create a ODModuleEntry instance using NewODModuleEntryWithNameXpcServiceName.
func ExampleNewODModuleEntryWithNameXpcServiceName() {
	_ = opendirectory.NewODModuleEntryWithNameXpcServiceName(
		"name", // name string
		"xpcServiceName", // xpcServiceName string
	)
	// Output:
}
