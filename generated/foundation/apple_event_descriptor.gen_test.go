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
// ExampleNewAppleEventDescriptorWithApplicationURL demonstrates how to create a AppleEventDescriptor instance using NewAppleEventDescriptorWithApplicationURL.
func ExampleNewAppleEventDescriptorWithApplicationURL() {
	_ = foundation.NewAppleEventDescriptorWithApplicationURL(
		foundation.URL{}, // applicationURL URL
	)
	// Output:
}
// ExampleNewAppleEventDescriptorWithBundleIdentifier demonstrates how to create a AppleEventDescriptor instance using NewAppleEventDescriptorWithBundleIdentifier.
func ExampleNewAppleEventDescriptorWithBundleIdentifier() {
	_ = foundation.NewAppleEventDescriptorWithBundleIdentifier(
		"bundleIdentifier", // bundleIdentifier string
	)
	// Output:
}
// ExampleNewAppleEventDescriptorWithDate demonstrates how to create a AppleEventDescriptor instance using NewAppleEventDescriptorWithDate.
func ExampleNewAppleEventDescriptorWithDate() {
	_ = foundation.NewAppleEventDescriptorWithDate(
		foundation.NSDate{}, // date NSDate
	)
	// Output:
}
// ExampleNewAppleEventDescriptorWithDouble demonstrates how to create a AppleEventDescriptor instance using NewAppleEventDescriptorWithDouble.
func ExampleNewAppleEventDescriptorWithDouble() {
	_ = foundation.NewAppleEventDescriptorWithDouble(
		0.0, // doubleValue float64
	)
	// Output:
}
// ExampleNewAppleEventDescriptorWithFileURL demonstrates how to create a AppleEventDescriptor instance using NewAppleEventDescriptorWithFileURL.
func ExampleNewAppleEventDescriptorWithFileURL() {
	_ = foundation.NewAppleEventDescriptorWithFileURL(
		foundation.URL{}, // fileURL URL
	)
	// Output:
}
// ExampleNewAppleEventDescriptorWithString demonstrates how to create a AppleEventDescriptor instance using NewAppleEventDescriptorWithString.
// Creates a descriptor initialized with type   that stores the text from the specified string.
func ExampleNewAppleEventDescriptorWithString() {
	_ = foundation.NewAppleEventDescriptorWithString(
		"string", // string string
	)
	// Output:
}
