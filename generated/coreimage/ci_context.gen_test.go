// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)

// Suppress unused import errors
var _ = coreimage.NewContext

// ExampleNewContext demonstrates how to create a Context instance.
// Initializes a context without a specific rendering destination, using default options.
func ExampleNewContext() {
	_ = coreimage.NewContext()
	// Output:
}
// ExampleContext_ClearCaches demonstrates using ClearCaches on a Context instance.
// Frees any cached data, such as temporary images, associated with the context and runs the garbage collector.
func ExampleContext_ClearCaches() {
	obj := coreimage.NewContext()
	obj.ClearCaches()
	// Output:
	}

// ExampleContext_ReclaimResources demonstrates using ReclaimResources on a Context instance.
// Runs the garbage collector to reclaim any resources that the context no longer requires.
func ExampleContext_ReclaimResources() {
	obj := coreimage.NewContext()
	obj.ReclaimResources()
	// Output:
	}

