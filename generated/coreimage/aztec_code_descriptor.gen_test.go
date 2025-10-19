// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)


// ExampleNewAztecCodeDescriptorWithPayloadIsCompactLayerCountDataCodewordCount demonstrates how to create a AztecCodeDescriptor instance using NewAztecCodeDescriptorWithPayloadIsCompactLayerCountDataCodewordCount.
// Initializes an Aztec code descriptor for the given payload and parameters.
func ExampleNewAztecCodeDescriptorWithPayloadIsCompactLayerCountDataCodewordCount() {
	_ = coreimage.NewAztecCodeDescriptorWithPayloadIsCompactLayerCountDataCodewordCount(
		nil, // errorCorrectedPayload unsafe.Pointer
		false, // isCompact bool
		0, // layerCount int
		0, // dataCodewordCount int
	)
	// Output:
}


