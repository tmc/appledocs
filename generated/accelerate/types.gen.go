// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

package accelerate
import (
	"unsafe"
)


// C struct types
// BNNSFilterParameters - A structure that contains common filter parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFilterParameters
type BNNSFilterParameters struct {
	Flags uint32 // A logical OR of zero or more values from BNNS flags.
}// BNNSLayerParametersBroadcastMatMul - A set of parameters that define a broadcast matrix multiply layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLayerParametersBroadcastMatMul
type BNNSLayerParametersBroadcastMatMul struct {
}// DSPComplex - A structure that represents a single-precision complex value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DSPComplex
type DSPComplex struct {
	Real float32 // The real part of the value.
}// DSPSplitComplex - A structure that represents a single-precision complex vector with the real and imaginary parts stored in separate arrays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DSPSplitComplex
type DSPSplitComplex struct {
}// DenseVector_Complex_Float - Contains a dense vector of float complex values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DenseVector_Complex_Float
type DenseVector_Complex_Float struct {
}// DenseVector_Double - A structure that contains a dense vector of double-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DenseVector_Double
type DenseVector_Double struct {
	Count int // The number of items in the vector.
}// SparseMatrixStructure - A description of the sparsity structure of a sparse matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseMatrixStructure
type SparseMatrixStructure struct {
	RowCount int // The number of rows in the matrix.
}// SparseMatrix_Float - A structure that contains a sparse matrix of single-precision, floating-point values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseMatrix_Float
type SparseMatrix_Float struct {
}// SparseSymbolicFactorOptions - A structure that contains options that affect the symbolic stage of a sparse factorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/SparseSymbolicFactorOptions
type SparseSymbolicFactorOptions struct {
}// bnns_user_message_data_t - Additional user-defined logging argument for message-logging callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_user_message_data_t
type bnns_user_message_data_t struct {
}// vImage_Buffer - An image buffer that stores an image’s pixel data, dimensions, and row stride.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_Buffer
type vImage_Buffer struct {
	Width unsafe.Pointer // The width of the image, in pixels.
}// vImage_CGImageFormat - The description of a Core Graphics image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_CGImageFormat
type vImage_CGImageFormat struct {
	BitmapInfo unsafe.Pointer // The component information that describes the color channels.
	BitsPerPixel uint32 // The number of bits that represents one pixel.
	ColorSpace ColorSpaceRef // A description of the position of the pixel data in the image, relative to a reference XYZ color space.
}// vImage_YpCbCrPixelRange - The description of range and clamping information for YpCbCr pixel formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_YpCbCrPixelRange
type vImage_YpCbCrPixelRange struct {
	CbCr_bias int32 // The encoding for   for this video format.
}// vImage_YpCbCrToARGB - The information that describes the conversion from YpCbCr to ARGB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_YpCbCrToARGB
type vImage_YpCbCrToARGB struct {
	Opaque uint8 // The bytes of the opaque representation.
}// vImage_YpCbCrToARGBMatrix - The 3 x 3 matrix that the vImage library uses to convert from YpCbCr to RGB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_YpCbCrToARGBMatrix
type vImage_YpCbCrToARGBMatrix struct {
}



