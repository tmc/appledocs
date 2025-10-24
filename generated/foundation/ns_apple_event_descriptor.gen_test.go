// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewAppleEventDescriptor

// ExampleNewAppleEventDescriptorListDescriptor demonstrates how to create a AppleEventDescriptor instance using NewAppleEventDescriptorListDescriptor.
// Initializes a newly allocated instance as an empty list descriptor.
func ExampleNewAppleEventDescriptorListDescriptor() {
	_ = foundation.NewAppleEventDescriptorListDescriptor()
	// Output:
}
// ExampleNewAppleEventDescriptorRecordDescriptor demonstrates how to create a AppleEventDescriptor instance using NewAppleEventDescriptorRecordDescriptor.
// Initializes a newly allocated instance as a descriptor that is an Apple event record.
func ExampleNewAppleEventDescriptorRecordDescriptor() {
	_ = foundation.NewAppleEventDescriptorRecordDescriptor()
	// Output:
}
// ExampleNewAppleEventDescriptorWithDouble demonstrates how to create a AppleEventDescriptor instance using NewAppleEventDescriptorWithDouble.
func ExampleNewAppleEventDescriptorWithDouble() {
	_ = foundation.NewAppleEventDescriptorWithDouble(
		0.0, // doubleValue float64
	)
	// Output:
}
// ExampleNewAppleEventDescriptorWithEnumCode demonstrates how to create a AppleEventDescriptor instance using NewAppleEventDescriptorWithEnumCode.
// Creates a descriptor initialized with type   that stores the specified enumerator data type value.
func ExampleNewAppleEventDescriptorWithEnumCode() {
	_ = foundation.NewAppleEventDescriptorWithEnumCode(
		foundation.uint32 /* not a class type */{}, // enumerator uint32 /* not a class type */
	)
	// Output:
}
// ExampleNewAppleEventDescriptorWithTypeCode demonstrates how to create a AppleEventDescriptor instance using NewAppleEventDescriptorWithTypeCode.
// Creates a descriptor initialized with type   that stores the specified type value.
func ExampleNewAppleEventDescriptorWithTypeCode() {
	_ = foundation.NewAppleEventDescriptorWithTypeCode(
		foundation.uint32 /* not a class type */{}, // typeCode uint32 /* not a class type */
	)
	// Output:
}
