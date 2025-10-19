// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)


// ExampleNewQRCodeDescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel demonstrates how to create a QRCodeDescriptor instance using NewQRCodeDescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel.
// Initializes a QR code descriptor for the given payload and parameters.
func ExampleNewQRCodeDescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel() {
	_ = coreimage.NewQRCodeDescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(
		nil, // errorCorrectedPayload unsafe.Pointer
		0, // symbolVersion int
		nil, // maskPattern unsafe.Pointer
		nil, // errorCorrectionLevel unsafe.Pointer
	)
	// Output:
}


