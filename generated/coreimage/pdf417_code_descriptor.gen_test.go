// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)


// ExampleNewPDF417CodeDescriptorWithPayloadIsCompactRowCountColumnCount demonstrates how to create a PDF417CodeDescriptor instance using NewPDF417CodeDescriptorWithPayloadIsCompactRowCountColumnCount.
// Initializes an PDF417 code descriptor for the given payload and parameters.
func ExampleNewPDF417CodeDescriptorWithPayloadIsCompactRowCountColumnCount() {
	_ = coreimage.NewPDF417CodeDescriptorWithPayloadIsCompactRowCountColumnCount(
		nil, // errorCorrectedPayload unsafe.Pointer
		false, // isCompact bool
		0, // rowCount int
		0, // columnCount int
	)
	// Output:
}


