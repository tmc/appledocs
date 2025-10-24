// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage


// Enum types and constants

// DataMatrixCodeECCVersion - Constants indicating the Data Matrix code ECC version.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/ECCVersion-swift.enum
type DataMatrixCodeECCVersion uint

const (
	// DataMatrixCodeECCVersion000 - Indicates error correction using convolutional code error correction with no data protection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/ECCVersion-swift.enum/v000
	DataMatrixCodeECCVersion000 DataMatrixCodeECCVersion = 0
	// DataMatrixCodeECCVersion050 - Indicates 1/4 of the symbol is dedicated to convolutional code error correction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/ECCVersion-swift.enum/v050
	DataMatrixCodeECCVersion050 DataMatrixCodeECCVersion = 0
	// DataMatrixCodeECCVersion080 - Indicates 1/3 of the symbol is dedicated to convolutional code error correction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/ECCVersion-swift.enum/v080
	DataMatrixCodeECCVersion080 DataMatrixCodeECCVersion = 0
	// DataMatrixCodeECCVersion100 - Indicates 1/2 of the symbol is dedicated to convolutional code error correction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/ECCVersion-swift.enum/v100
	DataMatrixCodeECCVersion100 DataMatrixCodeECCVersion = 0
	// DataMatrixCodeECCVersion140 - Indicates 3/4 of the symbol is dedicated to convolutional code error correction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/ECCVersion-swift.enum/v140
	DataMatrixCodeECCVersion140 DataMatrixCodeECCVersion = 0
	// DataMatrixCodeECCVersion200 - Indicates error correction using Reed-Solomon error correction. Data protection overhead varies based on symbol size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/ECCVersion-swift.enum/v200
	DataMatrixCodeECCVersion200 DataMatrixCodeECCVersion = 0
)


// QRCodeErrorCorrectionLevel - Constants indicating the percentage of the symbol that is dedicated to error correction.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/ErrorCorrectionLevel-swift.enum
type QRCodeErrorCorrectionLevel uint

const (
	// QRCodeErrorCorrectionLevelH - Indicates that approximately 65% of the symbol data is dedicated to error correction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/ErrorCorrectionLevel-swift.enum/levelH
	QRCodeErrorCorrectionLevelH QRCodeErrorCorrectionLevel = 0
	// QRCodeErrorCorrectionLevelL - Indicates that approximately 20% of the symbol data is dedicated to error correction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/ErrorCorrectionLevel-swift.enum/levelL
	QRCodeErrorCorrectionLevelL QRCodeErrorCorrectionLevel = 0
	// QRCodeErrorCorrectionLevelM - Indicates that approximately 37% of the symbol data is dedicated to error correction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/ErrorCorrectionLevel-swift.enum/levelM
	QRCodeErrorCorrectionLevelM QRCodeErrorCorrectionLevel = 0
	// QRCodeErrorCorrectionLevelQ - Indicates that approximately 55% of the symbol data is dedicated to error correction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/ErrorCorrectionLevel-swift.enum/levelQ
	QRCodeErrorCorrectionLevelQ QRCodeErrorCorrectionLevel = 0
)


// RenderDestinationAlphaMode - Different ways of representing alpha.
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


