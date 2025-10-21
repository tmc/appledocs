// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

// Enum types and constants
// CIDataMatrixCodeECCVersion - Constants indicating the Data Matrix code ECC version.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/ECCVersion-swift.enum
type DataMatrixCodeECCVersion uint

// CIQRCodeErrorCorrectionLevel - Constants indicating the percentage of the symbol that is dedicated to error correction.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/ErrorCorrectionLevel-swift.enum
type QRCodeErrorCorrectionLevel uint

const (
// QRCodeErrorCorrectionLevelL - Indicates that approximately 20% of the symbol data is dedicated to error correction.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/ErrorCorrectionLevel-swift.enum/levelL
QRCodeErrorCorrectionLevelL QRCodeErrorCorrectionLevel = 0
)

// CIRenderDestinationAlphaMode - Different ways of representing alpha.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestinationAlphaMode
type RenderDestinationAlphaMode uint

const (
// RenderDestinationAlphaNone - Designates a destination with no alpha compositing.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestinationAlphaMode/none
RenderDestinationAlphaNone RenderDestinationAlphaMode = 0
// RenderDestinationAlphaPremultiplied - Designates a destination that expects premultiplied alpha values.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestinationAlphaMode/premultiplied
RenderDestinationAlphaPremultiplied RenderDestinationAlphaMode = 0
// RenderDestinationAlphaUnpremultiplied - Designates a destination that expects non-premultiplied alpha values.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestinationAlphaMode/unpremultiplied
RenderDestinationAlphaUnpremultiplied RenderDestinationAlphaMode = 0
)


