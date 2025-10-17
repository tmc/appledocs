// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)


// ExampleNewGridViewWithNumberOfColumnsRows demonstrates how to create a GridView instance using NewGridViewWithNumberOfColumnsRows.
// Creates a newly allocated grid view object with the specified number of columns and rows.
func ExampleNewGridViewWithNumberOfColumnsRows() {
	_ = appkit.NewGridViewWithNumberOfColumnsRows(
		0, // columnCount int
		0, // rowCount int
	)
	// Output:
}

// ExampleNewGridViewWithViews demonstrates how to create a GridView instance using NewGridViewWithViews.
// Creates a newly allocated grid view object with the specified array of arrays of views.
func ExampleNewGridViewWithViews() {
	_ = appkit.NewGridViewWithViews(
		nil, // rows unsafe.Pointer
	)
	// Output:
}

// ExampleNewGridViewWithCoder demonstrates how to create a GridView instance using NewGridViewWithCoder.
// Creates a newly allocated grid view object from the coder.
func ExampleNewGridViewWithCoder() {
	_ = appkit.NewGridViewWithCoder(
		nil, // coder unsafe.Pointer
	)
	// Output:
}

// ExampleNewGridViewWithFrame demonstrates how to create a GridView instance using NewGridViewWithFrame.
// Creates a newly allocated grid view object with the specified frame rectangle.
func ExampleNewGridViewWithFrame() {
	_ = appkit.NewGridViewWithFrame(
		nil, // frameRect unsafe.Pointer
	)
	// Output:
}


