// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)


// ExampleNewDataMatrixCodeDescriptorWithPayloadRowCountColumnCountEccVersion demonstrates how to create a DataMatrixCodeDescriptor instance using NewDataMatrixCodeDescriptorWithPayloadRowCountColumnCountEccVersion.
// Initializes a Data Matrix code descriptor for the given payload and parameters.
func ExampleNewDataMatrixCodeDescriptorWithPayloadRowCountColumnCountEccVersion() {
	_ = coreimage.NewDataMatrixCodeDescriptorWithPayloadRowCountColumnCountEccVersion(
		nil, // errorCorrectedPayload unsafe.Pointer
		0, // rowCount int
		0, // columnCount int
		nil, // eccVersion unsafe.Pointer
	)
	// Output:
}


