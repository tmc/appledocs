// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)


// ExampleNewArrayWithArrayCopyItems demonstrates how to create a Array instance using NewArrayWithArrayCopyItems.
// Initializes a newly allocated array using   as the source of data objects for the array.
func ExampleNewArrayWithArrayCopyItems() {
	_ = foundation.NewArrayWithArrayCopyItems(
		nil, // array unsafe.Pointer
		false, // flag bool
	)
	// Output:
}

// ExampleNewArrayWithCoder demonstrates how to create a Array instance using NewArrayWithCoder.
func ExampleNewArrayWithCoder() {
	_ = foundation.NewArrayWithCoder(
		nil, // coder unsafe.Pointer
	)
	// Output:
}

// ExampleNewArrayWithContentsOfURL demonstrates how to create a Array instance using NewArrayWithContentsOfURL.
// Initializes a newly allocated array with the contents of the location specified by a given URL.
func ExampleNewArrayWithContentsOfURL() {
	_ = foundation.NewArrayWithContentsOfURL(
		nil, // url unsafe.Pointer
	)
	// Output:
}

// ExampleNewArrayWithObject demonstrates how to create a Array instance using NewArrayWithObject.
// Creates and returns an array containing a given object.
func ExampleNewArrayWithObject() {
	_ = foundation.NewArrayWithObject(
		nil, // anObject unsafe.Pointer
	)
	// Output:
}

// ExampleNewArrayWithObjectsCount demonstrates how to create a Array instance using NewArrayWithObjectsCount.
// Initializes a newly allocated array to include a given number of objects from a given C array.
func ExampleNewArrayWithObjectsCount() {
	_ = foundation.NewArrayWithObjectsCount(
		nil, // objects unsafe.Pointer
		0, // cnt uint
	)
	// Output:
}

// ExampleNewArrayWithObjects demonstrates how to create a Array instance using NewArrayWithObjects.
// Initializes a newly allocated array by placing in it the objects in the argument list.
func ExampleNewArrayWithObjects() {
	_ = foundation.NewArrayWithObjects(
		nil, // firstObj unsafe.Pointer
	)
	// Output:
}

// ExampleNewArray demonstrates how to create a Array instance.
// Initializes a newly allocated array.
func ExampleNewArray() {
	_ = foundation.NewArray()
	// Output:
}

// ExampleNewArrayWithContentsOfFile demonstrates how to create a Array instance using NewArrayWithContentsOfFile.
// Initializes a newly allocated array with the contents of the file specified by a given path.
func ExampleNewArrayWithContentsOfFile() {
	_ = foundation.NewArrayWithContentsOfFile(
		"path", // path string
	)
	// Output:
}

// ExampleNewArrayWithContentsOfURLError demonstrates how to create a Array instance using NewArrayWithContentsOfURLError.
func ExampleNewArrayWithContentsOfURLError() {
	_ = foundation.NewArrayWithContentsOfURLError(
		nil, // url unsafe.Pointer
		nil, // error unsafe.Pointer
	)
	// Output:
}

// ExampleNewArrayWithArray demonstrates how to create a Array instance using NewArrayWithArray.
// Initializes a newly allocated array by placing in it the objects contained in a given array.
func ExampleNewArrayWithArray() {
	_ = foundation.NewArrayWithArray(
		nil, // array unsafe.Pointer
	)
	// Output:
}


