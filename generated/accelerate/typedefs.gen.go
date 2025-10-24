// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

package accelerate
import (
"unsafe"
)

// Type aliases and typedefs
// Bnns_graph_compile_message_fn_t - The graph compile-message logging callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_compile_message_fn_t
// bnns_graph_compile_message_fn_t is a callback function
// C type: void (*)(BNNSGraphMessageLevel, const char *, const char *, bnns_user_message_data_t *)
type Bnns_graph_compile_message_fn_t = func(BNNSGraphMessageLevel, string, string, unsafe.Pointer)
// Bnns_graph_execute_message_fn_t - The graph execute-message logging callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_execute_message_fn_t
// bnns_graph_execute_message_fn_t is a callback function
// C type: void (*)(BNNSGraphMessageLevel, const char *, const char *, bnns_user_message_data_t *)
type Bnns_graph_execute_message_fn_t = func(BNNSGraphMessageLevel, string, string, unsafe.Pointer)
// Bnns_graph_free_all_fn_t - The workspace and output deallocation function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_free_all_fn_t
// bnns_graph_free_all_fn_t is a callback function
// C type: void (*)(void *, unsigned long)
type Bnns_graph_free_all_fn_t = func(unsafe.Pointer, uint)
// Bnns_graph_realloc_fn_t - The workspace and output allocation function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_realloc_fn_t
// bnns_graph_realloc_fn_t is a callback function
// C type: int (*)(void *, unsigned long, void **, unsigned long, unsigned long)
type Bnns_graph_realloc_fn_t = func(unsafe.Pointer, uint, unsafe.Pointer, uint, uint) int32
// BNNSAlloc - A type-alias for a user-provided memory allocation function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSAlloc
// BNNSAlloc is a callback function
// C type: int (*)(void **, unsigned long, unsigned long)
type BNNSAlloc = func(unsafe.Pointer, uint, uint) int32
// BNNSFilter - An opaque type that represents a filter.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFilter
// BNNSFilter has base type: void *
type BNNSFilter uintptr
// BNNSFree - A type-alias for a user-provided memory deallocation function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFree
// BNNSFree is a callback function
// C type: void (*)(void *)
type BNNSFree = func(unsafe.Pointer)
// BNNSNearestNeighbors - A k-nearest neighbors object.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSNearestNeighbors
// BNNSNearestNeighbors has base type: void *
type BNNSNearestNeighbors uintptr
// BNNSRandomGenerator - A pointer to a random number generator object.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSRandomGenerator
// BNNSRandomGenerator has base type: void *
type BNNSRandomGenerator uintptr
// COMPLEX type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/COMPLEX
// COMPLEX has base type: DSPComplex
type COMPLEX uintptr
// COMPLEX_SPLIT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/COMPLEX_SPLIT
// COMPLEX_SPLIT has base type: DSPSplitComplex
type COMPLEX_SPLIT uintptr
// DOUBLE_COMPLEX type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DOUBLE_COMPLEX
// DOUBLE_COMPLEX has base type: DSPDoubleComplex
type DOUBLE_COMPLEX uintptr
// DOUBLE_COMPLEX_SPLIT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/DOUBLE_COMPLEX_SPLIT
// DOUBLE_COMPLEX_SPLIT has base type: DSPDoubleSplitComplex
type DOUBLE_COMPLEX_SPLIT uintptr
// FFTDirection - Constants that specify whether to perform a forward or inverse FFT.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/FFTDirection
type FFTDirection int32
// FFTRadix - The radix of the FFT decomposition.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/FFTRadix
type FFTRadix int32
// FFTSetup - An opaque type that contains setup information for a single-precision FFT transform.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/FFTSetup
// FFTSetup has base type: struct OpaqueFFTSetup *
type FFTSetup uintptr
// FFTSetupD - An opaque type that contains setup information for a double-precision FFT transform.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/FFTSetupD
// FFTSetupD has base type: struct OpaqueFFTSetupD *
type FFTSetupD uintptr
// GammaFunction - A type for a gamma function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/GammaFunction
// GammaFunction has base type: void *
type GammaFunction uintptr
// La_attribute_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/la_attribute_t
// la_attribute_t has base type: unsigned long
type La_attribute_t uintptr
// La_count_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/la_count_t
// la_count_t has base type: unsigned long
type La_count_t uintptr
// La_deallocator_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/la_deallocator_t
// la_deallocator_t is a callback function
// C type: void (*)(void *)
type La_deallocator_t = func(unsafe.Pointer)
// La_hint_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/la_hint_t
// la_hint_t has base type: unsigned long
type La_hint_t uintptr
// La_index_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/la_index_t
// la_index_t has base type: long
type La_index_t uintptr
// La_norm_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/la_norm_t
// la_norm_t has base type: unsigned long
type La_norm_t uintptr
// La_object_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/la_object_t
// la_object_t has base type: NSObject<OS_la_object> *
type La_object_t uintptr
// La_scalar_type_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/la_scalar_type_t
type La_scalar_type_t uint32
// La_status_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/la_status_t
// la_status_t has base type: long
type La_status_t uintptr
// Pixel_16F type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_16F
// Pixel_16F has base type: uint16_t
type Pixel_16F uintptr
// Pixel_16F16F type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_16F16F
// Pixel_16F16F has base type: uint16_t[2]
type Pixel_16F16F uintptr
// Pixel_16Q12 - A type for a signed 16-bit, fixed-point number with 12 bits of fractional precision.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_16Q12
// Pixel_16Q12 has base type: int16_t
type Pixel_16Q12 uintptr
// Pixel_16S - A type for a planar, 16-bits-per-channel, signed pixel.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_16S
// Pixel_16S has base type: int16_t
type Pixel_16S uintptr
// Pixel_16S16S type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_16S16S
// Pixel_16S16S has base type: int16_t[2]
type Pixel_16S16S uintptr
// Pixel_16U - A type for a planar, 16-bits-per-channel, unsigned pixel.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_16U
// Pixel_16U has base type: uint16_t
type Pixel_16U uintptr
// Pixel_16U16U - A type for a two-channel, 16-bits-per-channel, unsigned pixel.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_16U16U
// Pixel_16U16U has base type: uint16_t[2]
type Pixel_16U16U uintptr
// Pixel_32U - A type you use for the XRGB2101010 format.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_32U
// Pixel_32U has base type: uint32_t
type Pixel_32U uintptr
// Pixel_8 - A type for a planar, 8-bits-per-channel, unsigned pixel.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_8
// Pixel_8 has base type: uint8_t
type Pixel_8 uintptr
// Pixel_88 - A type for a two-channel, 8-bits-per-channel, unsigned pixel.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_88
// Pixel_88 has base type: uint8_t[2]
type Pixel_88 uintptr
// Pixel_8888 - A type for a four-channel, 8-bits-per-channel, unsigned pixel.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_8888
// Pixel_8888 has base type: uint8_t[4]
type Pixel_8888 uintptr
// Pixel_ARGB_16F type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_ARGB_16F
// Pixel_ARGB_16F has base type: uint16_t[4]
type Pixel_ARGB_16F uintptr
// Pixel_ARGB_16S - A type for a four-channel, 16-bits-per-channel, signed pixel.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_ARGB_16S
// Pixel_ARGB_16S has base type: int16_t[4]
type Pixel_ARGB_16S uintptr
// Pixel_ARGB_16U - A type for a four-channel, 16-bits-per-channel, unsigned pixel.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_ARGB_16U
// Pixel_ARGB_16U has base type: uint16_t[4]
type Pixel_ARGB_16U uintptr
// Pixel_F - A type for a planar, 32-bits-per-channel, floating-point pixel.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_F
// Pixel_F has base type: float
type Pixel_F uintptr
// Pixel_FF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_FF
// Pixel_FF has base type: float[2]
type Pixel_FF uintptr
// Pixel_FFFF - A type for a four-channel, 32-bits-per-channel, floating-point pixel.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_FFFF
// Pixel_FFFF has base type: float[4]
type Pixel_FFFF uintptr
// Quadrature_function_array type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/quadrature_function_array
// quadrature_function_array is a callback function
// C type: void (*)(void *, unsigned long, const double *, double *)
type Quadrature_function_array = func(unsafe.Pointer, uint, float64, float64)
// ResamplingFilter - A pointer to a resampling filter callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/ResamplingFilter
// ResamplingFilter has base type: void *
type ResamplingFilter uintptr
// Sparse_index - The index type.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_index
// sparse_index has base type: int64_t
type Sparse_index uintptr
// Sparse_matrix_double - Sparse matrix opaque type for double.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_double
// sparse_matrix_double has base type: struct sparse_m_double *
type Sparse_matrix_double uintptr
// Sparse_matrix_double_complex type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_double_complex
// sparse_matrix_double_complex has base type: struct sparse_m_double_complex *
type Sparse_matrix_double_complex uintptr
// Sparse_matrix_float - Sparse matrix opaque type for float.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_float
// sparse_matrix_float has base type: struct sparse_m_float *
type Sparse_matrix_float uintptr
// Sparse_stride - The stride type.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_stride
// sparse_stride has base type: int64_t
type Sparse_stride uintptr
// VBool32 - A 128-bit vector packed with   values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vBool32
// vBool32 has base type: __attribute__((__vector_size__(4 * sizeof(unsigned int)))) unsigned int
type VBool32 uintptr
// VDouble - A 128-bit vector packed with   values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDouble
// vDouble has base type: __attribute__((__vector_size__(2 * sizeof(double)))) double
type VDouble uintptr
// VDSP_biquad_Setup - A data structure that contains precalculated data for use by the single-precision cascaded biquadratic IIR filter function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquad_Setup
// vDSP_biquad_Setup has base type: struct vDSP_biquad_SetupStruct *
type VDSP_biquad_Setup uintptr
// VDSP_biquad_SetupD - A data structure that contains precalculated data for use by the double-precision cascaded biquadratic IIR filter function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquad_SetupD
// vDSP_biquad_SetupD has base type: struct vDSP_biquad_SetupStructD *
type VDSP_biquad_SetupD uintptr
// VDSP_biquadm_Setup - A data structure that contains precalculated data for use by a single-precision, multichannel cascaded biquadratic filter function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquadm_Setup
// vDSP_biquadm_Setup has base type: struct vDSP_biquadm_SetupStruct *
type VDSP_biquadm_Setup uintptr
// VDSP_biquadm_SetupD - A data structure that contains precalculated data for use by a double-precision, multichannel cascaded biquadratic filter function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquadm_SetupD
// vDSP_biquadm_SetupD has base type: struct vDSP_biquadm_SetupStructD *
type VDSP_biquadm_SetupD uintptr
// VDSP_DFT_Interleaved_Setup - An opaque type that contains setup information for an interleaved single-precision discrete Fourier transform (DFT).
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Interleaved_Setup
// vDSP_DFT_Interleaved_Setup has base type: struct vDSP_DFT_Interleaved_SetupStruct *
type VDSP_DFT_Interleaved_Setup uintptr
// VDSP_DFT_Interleaved_SetupD - An opaque type that contains setup information for an interleaved double-precision discrete Fourier transform (DFT).
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Interleaved_SetupD
// vDSP_DFT_Interleaved_SetupD has base type: struct vDSP_DFT_Interleaved_SetupStructD *
type VDSP_DFT_Interleaved_SetupD uintptr
// VDSP_DFT_Setup - An opaque type that contains setup information for a single-precision discrete Fourier transform (DFT).
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Setup
// vDSP_DFT_Setup has base type: struct vDSP_DFT_SetupStruct *
type VDSP_DFT_Setup uintptr
// VDSP_DFT_SetupD - An opaque type that contains setup information for a double-precision discrete Fourier transform (DFT).
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_SetupD
// vDSP_DFT_SetupD has base type: struct vDSP_DFT_SetupStructD *
type VDSP_DFT_SetupD uintptr
// VDSP_Length - An unsigned-integer value that represents the size of vectors and the indices of elements in vectors.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_Length
// vDSP_Length has base type: unsigned long
type VDSP_Length uintptr
// VDSP_Stride - An integer value that represents the differences between indices of elements, including the lengths of strides.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_Stride
// vDSP_Stride has base type: long
type VDSP_Stride uintptr
// VFloat - A 128-bit vector packed with   values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vFloat
// vFloat has base type: __attribute__((__vector_size__(4 * sizeof(float)))) float
type VFloat uintptr
// VImage_CGAffineTransform - A structure for values that represent a Core Graphics–compatible affine transformation.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_CGAffineTransform
// vImage_CGAffineTransform has base type: vImage_AffineTransform_Double
type VImage_CGAffineTransform uintptr
// VImage_Error - A type for image errors.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_Error
// vImage_Error has base type: ssize_t
type VImage_Error uintptr
// VImage_Flags - A type for processing options.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_Flags
// vImage_Flags has base type: uint32_t
type VImage_Flags uintptr
// VImage_MultidimensionalTable - An opaque pointer that represents a multidimensional lookup table.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_MultidimensionalTable
// vImage_MultidimensionalTable has base type: struct vImage_MultidimensionalTableData *
type VImage_MultidimensionalTable uintptr
// VImage_WarpInterpolation - Constants for selecting the interpolation mode
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImage_WarpInterpolation
// vImage_WarpInterpolation has base type: int32_t
type VImage_WarpInterpolation uintptr
// VImageBufferTypeCode - Type codes, such as chrominance or luminance, for the contents of a vImage buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageBufferTypeCode
// vImageBufferTypeCode has base type: uint32_t
type VImageBufferTypeCode uintptr
// VImageConstCVImageFormatRef - An immutable description of image encoding in a Core Video pixel buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConstCVImageFormat
// vImageConstCVImageFormatRef has base type: const struct vImageCVImageFormat *
type VImageConstCVImageFormatRef uintptr
// VImageConverterRef - A description of a conversion from one image format to another.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConverter
// vImageConverterRef has base type: struct vImageConverter *
type VImageConverterRef uintptr
// VImageCVImageFormatRef - A mutable description of image encoding in a Core Video pixel buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat
// vImageCVImageFormatRef has base type: struct vImageCVImageFormat *
type VImageCVImageFormatRef uintptr
// VImageCVImageFormatError - Additional error codes for functions that use the vImageCVImageFormatRef
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormatError
// vImageCVImageFormatError has base type: ssize_t
type VImageCVImageFormatError uintptr
// VImageMatrixType - An enumeration of RGB -> Y’CbCr conversion matrix types.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMatrixType
// vImageMatrixType has base type: uint32_t
type VImageMatrixType uintptr
// VImagePixelCount - A type for the number of pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImagePixelCount
// vImagePixelCount has base type: unsigned long
type VImagePixelCount uintptr
// VSInt16 - A 128-bit vector packed with   values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vSInt16
// vSInt16 has base type: __attribute__((__vector_size__(8 * sizeof(short)))) short
type VSInt16 uintptr
// VSInt32 - A 128-bit vector packed with   values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vSInt32
// vSInt32 has base type: __attribute__((__vector_size__(4 * sizeof(int)))) int
type VSInt32 uintptr
// VSInt64 - A 128-bit vector packed with   values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vSInt64
// vSInt64 has base type: __attribute__((__vector_size__(2 * sizeof(long long)))) long long
type VSInt64 uintptr
// VSInt8 - A 128-bit vector packed with   values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vSInt8
// vSInt8 has base type: __attribute__((__vector_size__(16 * sizeof(signed char)))) signed char
type VSInt8 uintptr
// VUInt16 - A 128-bit vector packed with   values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vUInt16
// vUInt16 has base type: __attribute__((__vector_size__(8 * sizeof(unsigned short)))) unsigned short
type VUInt16 uintptr
// VUInt32 - A 128-bit vector packed with   values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vUInt32
// vUInt32 has base type: __attribute__((__vector_size__(4 * sizeof(unsigned int)))) unsigned int
type VUInt32 uintptr
// VUInt64 - A 128-bit vector packed with   values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vUInt64
// vUInt64 has base type: __attribute__((__vector_size__(2 * sizeof(unsigned long long)))) unsigned long long
type VUInt64 uintptr
// VUInt8 - A 128-bit vector packed with   values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vUInt8
// vUInt8 has base type: __attribute__((__vector_size__(16 * sizeof(unsigned char)))) unsigned char
type VUInt8 uintptr

