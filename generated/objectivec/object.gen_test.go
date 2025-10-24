// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec_test

import (
	"github.com/tmc/appledocs/generated/objectivec"
)

// Suppress unused import errors
var _ = objectivec.NewObject

// ExampleNewObject demonstrates how to create a Object instance.
// Implemented by subclasses to initialize a new object (the receiver) immediately after memory for it has been allocated.
func ExampleNewObject() {
	_ = objectivec.NewObject()
	// Output:
}
// ExampleObject_Dealloc demonstrates using Dealloc on a Object instance.
// Deallocates the memory occupied by the receiver.
func ExampleObject_Dealloc() {
	obj := objectivec.NewObject()
	obj.Dealloc()
	// Output:
}

// ExampleObject_Finalize demonstrates using Finalize on a Object instance.
// The garbage collector invokes this method on the receiver before disposing of the memory it uses.
func ExampleObject_Finalize() {
	obj := objectivec.NewObject()
	obj.Finalize()
	// Output:
}

