// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

package accelerate

// Type aliases and typedefs
// BNNSAlloc - A type-alias for a user-provided memory allocation function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSAlloc
// BNNSAlloc has base type: int (*)(void **, unsigned long, unsigned long)
type BNNSAlloc uintptr
// BNNSFilter - An opaque type that represents a filter.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFilter
// BNNSFilter has base type: void *
type BNNSFilter uintptr
// BNNSFree - A type-alias for a user-provided memory deallocation function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFree
// BNNSFree has base type: void (*)(void *)
type BNNSFree uintptr
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
// Pixel_16F type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_16F
// Pixel_16F has base type: uint16_t
type Pixel_16F uintptr
// Pixel_8 - A type for a planar, 8-bits-per-channel, unsigned pixel.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_8
// Pixel_8 has base type: uint8_t
type Pixel_8 uintptr
// Pixel_ARGB_16F type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_ARGB_16F
// Pixel_ARGB_16F has base type: uint16_t[4]
type Pixel_ARGB_16F uintptr
// Pixel_F - A type for a planar, 32-bits-per-channel, floating-point pixel.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/Pixel_F
// Pixel_F has base type: float
type Pixel_F uintptr
// bnns_graph_compile_message_fn_t - The graph compile-message logging callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_compile_message_fn_t
// bnns_graph_compile_message_fn_t has base type: void (*)(BNNSGraphMessageLevel, const char *, const char *, bnns_user_message_data_t *)
type bnns_graph_compile_message_fn_t uintptr
// bnns_graph_execute_message_fn_t - The graph execute-message logging callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_execute_message_fn_t
// bnns_graph_execute_message_fn_t has base type: void (*)(BNNSGraphMessageLevel, const char *, const char *, bnns_user_message_data_t *)
type bnns_graph_execute_message_fn_t uintptr
// bnns_graph_free_all_fn_t - The workspace and output deallocation function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/bnns_graph_free_all_fn_t
// bnns_graph_free_all_fn_t has base type: void (*)(void *, unsigned long)
type bnns_graph_free_all_fn_t uintptr
// la_count_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/la_count_t
// la_count_t has base type: unsigned long
type la_count_t uintptr
// sparse_dimension - The dimension type.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_dimension
// sparse_dimension has base type: uint64_t
type sparse_dimension uintptr
// sparse_index - The index type.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_index
// sparse_index has base type: int64_t
type sparse_index uintptr
// sparse_matrix_double_complex type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_double_complex
// sparse_matrix_double_complex has base type: struct sparse_m_double_complex *
type sparse_matrix_double_complex uintptr
// sparse_matrix_float - Sparse matrix opaque type for float.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_float
// sparse_matrix_float has base type: struct sparse_m_float *
type sparse_matrix_float uintptr
// sparse_stride - The stride type.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_stride
// sparse_stride has base type: int64_t
type sparse_stride uintptr
// vDSP_DFT_Setup - An opaque type that contains setup information for a single-precision discrete Fourier transform (DFT).
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Setup
// vDSP_DFT_Setup has base type: struct vDSP_DFT_SetupStruct *
type vDSP_DFT_Setup uintptr
// vDSP_Length - An unsigned-integer value that represents the size of vectors and the indices of elements in vectors.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_Length
// vDSP_Length has base type: unsigned long
type vDSP_Length uintptr
// vDSP_Stride - An integer value that represents the differences between indices of elements, including the lengths of strides.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_Stride
// vDSP_Stride has base type: long
type vDSP_Stride uintptr
// vDSP_biquad_Setup - A data structure that contains precalculated data for use by the single-precision cascaded biquadratic IIR filter function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquad_Setup
// vDSP_biquad_Setup has base type: struct vDSP_biquad_SetupStruct *
type vDSP_biquad_Setup uintptr
// vDSP_biquad_SetupD - A data structure that contains precalculated data for use by the double-precision cascaded biquadratic IIR filter function.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquad_SetupD
// vDSP_biquad_SetupD has base type: struct vDSP_biquad_SetupStructD *
type vDSP_biquad_SetupD uintptr
// vImageCVImageFormatRef - A mutable description of image encoding in a Core Video pixel buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat
// vImageCVImageFormatRef has base type: struct vImageCVImageFormat *
type vImageCVImageFormatRef uintptr
// vImageConverterRef - A description of a conversion from one image format to another.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConverter
// vImageConverterRef has base type: struct vImageConverter *
type vImageConverterRef uintptr
// vImageMatrixType - An enumeration of RGB -> Y’CbCr conversion matrix types.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMatrixType
// vImageMatrixType has base type: uint32_t
type vImageMatrixType uintptr
// vSInt32 - A 128-bit vector packed with   values.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vSInt32
// vSInt32 has base type: __attribute__((__vector_size__(4 * sizeof(int)))) int
type vSInt32 uintptr

