// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

package accelerate

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Accelerate Functions (145 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_BNNSDirectApplyBroadcastMatMul func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_BNNSFilterCreateLayerBroadcastMatMul func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphContextEnableNanAndInfChecks func(unsafe.Pointer, bool)
	_BNNSMatMul func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_BNNSMatMulWorkspaceSize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_cblas_sgemm func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, []float32, unsafe.Pointer, []float32, unsafe.Pointer, unsafe.Pointer, []float32, unsafe.Pointer)
	_sparse_inner_product_dense_double_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_inner_product_dense_float_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_inner_product_sparse_double_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_inner_product_sparse_float_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_insert_entry_double_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_insert_entry_float_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_matrix_product_dense_double_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_matrix_product_dense_float_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_matrix_product_sparse_double_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_matrix_product_sparse_float_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_matrix_trace_double_complex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_matrix_trace_float_complex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_matrix_triangular_solve_dense_double_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_matrix_triangular_solve_dense_float_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_matrix_vector_product_dense_double_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_matrix_vector_product_dense_float_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_outer_product_dense_double_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_outer_product_dense_float_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_vector_add_with_scale_dense_double_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_sparse_vector_add_with_scale_dense_float_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_sparse_vector_triangular_solve_dense_double_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_vector_triangular_solve_dense_float_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_DFT_Interleaved_DestroySetup func(unsafe.Pointer)
	_vDSP_biquad func(unsafe.Pointer, []float32, []float32, unsafe.Pointer, []float32, unsafe.Pointer, unsafe.Pointer)
	_vDSP_biquad_CreateSetup func([]float64, unsafe.Pointer) unsafe.Pointer
	_vDSP_biquad_SetCoefficientsSingle func(unsafe.Pointer, []float32, unsafe.Pointer, unsafe.Pointer)
	_vDSP_fft2d_zip func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_vDSP_fftm_zrip func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_vDSP_vtmerg func([]float32, unsafe.Pointer, []float32, unsafe.Pointer, []float32, unsafe.Pointer, unsafe.Pointer)
	_vDSP_ztoc func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_vImageAlphaBlend_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageBuffer_InitWithCVPixelBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, []float64, unsafe.Pointer) unsafe.Pointer
	_vImageContrastStretch_ARGBFFFF func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConvert_ARGBToYpCbCr_GenerateConversion func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConvert_YpCbCrToARGB_GenerateConversion func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConvolveWithBias_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint32, uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageDilate_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageMatrixMultiply_ARGB8888ToPlanar8 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageMultidimensionalTable_Create func(unsafe.Pointer, uint32, uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageMultidimensionalTable_Release func(unsafe.Pointer) unsafe.Pointer
	_vImageMultidimensionalTable_Retain func(unsafe.Pointer) unsafe.Pointer
	_vImagePremultipliedAlphaBlend_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImagePremultipliedConstAlphaBlend_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImagePremultiplyData_RGBA8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageUnpremultiplyData_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vA128Shift func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vLL128Shift func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vLR128Shift func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vS128Add func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vS128AddS func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vS128Sub func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vS128SubS func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vS64FullMulOdd func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vS64SubS func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vU128Add func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vU128AddS func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vU128Sub func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vU128SubS func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vU64FullMulOdd func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvacos func([]float64, []float64, []int)
	_vvacosf func([]float32, []float32, []int)
	_vvacosh func([]float64, []float64, []int)
	_vvacoshf func([]float32, []float32, []int)
	_vvasin func([]float64, []float64, []int)
	_vvasinf func([]float32, []float32, []int)
	_vvasinh func([]float64, []float64, []int)
	_vvasinhf func([]float32, []float32, []int)
	_vvatan func([]float64, []float64, []int)
	_vvatan2 func([]float64, []float64, []float64, []int)
	_vvatan2f func([]float32, []float32, []float32, []int)
	_vvatanf func([]float32, []float32, []int)
	_vvatanh func([]float64, []float64, []int)
	_vvatanhf func([]float32, []float32, []int)
	_vvceil func([]float64, []float64, []int)
	_vvceilf func([]float32, []float32, []int)
	_vvcopysign func([]float64, []float64, []float64, []int)
	_vvcopysignf func([]float32, []float32, []float32, []int)
	_vvcos func([]float64, []float64, []int)
	_vvcosf func([]float32, []float32, []int)
	_vvcosh func([]float64, []float64, []int)
	_vvcoshf func([]float32, []float32, []int)
	_vvcosisin func(unsafe.Pointer, []float64, []int)
	_vvcosisinf func(unsafe.Pointer, []float32, []int)
	_vvcospi func([]float64, []float64, []int)
	_vvcospif func([]float32, []float32, []int)
	_vvdiv func([]float64, []float64, []float64, []int)
	_vvdivf func([]float32, []float32, []float32, []int)
	_vvexp func([]float64, []float64, []int)
	_vvexp2 func([]float64, []float64, []int)
	_vvexp2f func([]float32, []float32, []int)
	_vvexpf func([]float32, []float32, []int)
	_vvexpm1 func([]float64, []float64, []int)
	_vvexpm1f func([]float32, []float32, []int)
	_vvfabs func([]float64, []float64, []int)
	_vvfabsf func([]float32, []float32, []int)
	_vvfloor func([]float64, []float64, []int)
	_vvfloorf func([]float32, []float32, []int)
	_vvfmod func([]float64, []float64, []float64, []int)
	_vvfmodf func([]float32, []float32, []float32, []int)
	_vvint func([]float64, []float64, []int)
	_vvintf func([]float32, []float32, []int)
	_vvlog func([]float64, []float64, []int)
	_vvlog10 func([]float64, []float64, []int)
	_vvlog10f func([]float32, []float32, []int)
	_vvlog1p func([]float64, []float64, []int)
	_vvlog1pf func([]float32, []float32, []int)
	_vvlog2 func([]float64, []float64, []int)
	_vvlog2f func([]float32, []float32, []int)
	_vvlogb func([]float64, []float64, []int)
	_vvlogbf func([]float32, []float32, []int)
	_vvlogf func([]float32, []float32, []int)
	_vvnextafter func([]float64, []float64, []float64, []int)
	_vvnextafterf func([]float32, []float32, []float32, []int)
	_vvnint func([]float64, []float64, []int)
	_vvnintf func([]float32, []float32, []int)
	_vvpow func([]float64, []float64, []float64, []int)
	_vvpowf func([]float32, []float32, []float32, []int)
	_vvrec func([]float64, []float64, []int)
	_vvrecf func([]float32, []float32, []int)
	_vvremainder func([]float64, []float64, []float64, []int)
	_vvremainderf func([]float32, []float32, []float32, []int)
	_vvrsqrt func([]float64, []float64, []int)
	_vvrsqrtf func([]float32, []float32, []int)
	_vvsin func([]float64, []float64, []int)
	_vvsincos func([]float64, []float64, []float64, []int)
	_vvsincosf func([]float32, []float32, []float32, []int)
	_vvsinf func([]float32, []float32, []int)
	_vvsinh func([]float64, []float64, []int)
	_vvsinhf func([]float32, []float32, []int)
	_vvsinpi func([]float64, []float64, []int)
	_vvsinpif func([]float32, []float32, []int)
	_vvsqrt func([]float64, []float64, []int)
	_vvsqrtf func([]float32, []float32, []int)
	_vvtan func([]float64, []float64, []int)
	_vvtanf func([]float32, []float32, []int)
	_vvtanh func([]float64, []float64, []int)
	_vvtanhf func([]float32, []float32, []int)
	_vvtanpi func([]float64, []float64, []int)
	_vvtanpif func([]float32, []float32, []int)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_BNNSDirectApplyBroadcastMatMul, lib, "BNNSDirectApplyBroadcastMatMul")
	tryRegister(&_BNNSFilterCreateLayerBroadcastMatMul, lib, "BNNSFilterCreateLayerBroadcastMatMul")
	tryRegister(&_BNNSGraphContextEnableNanAndInfChecks, lib, "BNNSGraphContextEnableNanAndInfChecks")
	tryRegister(&_BNNSMatMul, lib, "BNNSMatMul")
	tryRegister(&_BNNSMatMulWorkspaceSize, lib, "BNNSMatMulWorkspaceSize")
	tryRegister(&_cblas_sgemm, lib, "cblas_sgemm")
	tryRegister(&_sparse_inner_product_dense_double_complex, lib, "sparse_inner_product_dense_double_complex")
	tryRegister(&_sparse_inner_product_dense_float_complex, lib, "sparse_inner_product_dense_float_complex")
	tryRegister(&_sparse_inner_product_sparse_double_complex, lib, "sparse_inner_product_sparse_double_complex")
	tryRegister(&_sparse_inner_product_sparse_float_complex, lib, "sparse_inner_product_sparse_float_complex")
	tryRegister(&_sparse_insert_entry_double_complex, lib, "sparse_insert_entry_double_complex")
	tryRegister(&_sparse_insert_entry_float_complex, lib, "sparse_insert_entry_float_complex")
	tryRegister(&_sparse_matrix_product_dense_double_complex, lib, "sparse_matrix_product_dense_double_complex")
	tryRegister(&_sparse_matrix_product_dense_float_complex, lib, "sparse_matrix_product_dense_float_complex")
	tryRegister(&_sparse_matrix_product_sparse_double_complex, lib, "sparse_matrix_product_sparse_double_complex")
	tryRegister(&_sparse_matrix_product_sparse_float_complex, lib, "sparse_matrix_product_sparse_float_complex")
	tryRegister(&_sparse_matrix_trace_double_complex, lib, "sparse_matrix_trace_double_complex")
	tryRegister(&_sparse_matrix_trace_float_complex, lib, "sparse_matrix_trace_float_complex")
	tryRegister(&_sparse_matrix_triangular_solve_dense_double_complex, lib, "sparse_matrix_triangular_solve_dense_double_complex")
	tryRegister(&_sparse_matrix_triangular_solve_dense_float_complex, lib, "sparse_matrix_triangular_solve_dense_float_complex")
	tryRegister(&_sparse_matrix_vector_product_dense_double_complex, lib, "sparse_matrix_vector_product_dense_double_complex")
	tryRegister(&_sparse_matrix_vector_product_dense_float_complex, lib, "sparse_matrix_vector_product_dense_float_complex")
	tryRegister(&_sparse_outer_product_dense_double_complex, lib, "sparse_outer_product_dense_double_complex")
	tryRegister(&_sparse_outer_product_dense_float_complex, lib, "sparse_outer_product_dense_float_complex")
	tryRegister(&_sparse_vector_add_with_scale_dense_double_complex, lib, "sparse_vector_add_with_scale_dense_double_complex")
	tryRegister(&_sparse_vector_add_with_scale_dense_float_complex, lib, "sparse_vector_add_with_scale_dense_float_complex")
	tryRegister(&_sparse_vector_triangular_solve_dense_double_complex, lib, "sparse_vector_triangular_solve_dense_double_complex")
	tryRegister(&_sparse_vector_triangular_solve_dense_float_complex, lib, "sparse_vector_triangular_solve_dense_float_complex")
	tryRegister(&_vDSP_DFT_Interleaved_DestroySetup, lib, "vDSP_DFT_Interleaved_DestroySetup")
	tryRegister(&_vDSP_biquad, lib, "vDSP_biquad")
	tryRegister(&_vDSP_biquad_CreateSetup, lib, "vDSP_biquad_CreateSetup")
	tryRegister(&_vDSP_biquad_SetCoefficientsSingle, lib, "vDSP_biquad_SetCoefficientsSingle")
	tryRegister(&_vDSP_fft2d_zip, lib, "vDSP_fft2d_zip")
	tryRegister(&_vDSP_fftm_zrip, lib, "vDSP_fftm_zrip")
	tryRegister(&_vDSP_vtmerg, lib, "vDSP_vtmerg")
	tryRegister(&_vDSP_ztoc, lib, "vDSP_ztoc")
	tryRegister(&_vImageAlphaBlend_ARGB8888, lib, "vImageAlphaBlend_ARGB8888")
	tryRegister(&_vImageBuffer_InitWithCVPixelBuffer, lib, "vImageBuffer_InitWithCVPixelBuffer")
	tryRegister(&_vImageContrastStretch_ARGBFFFF, lib, "vImageContrastStretch_ARGBFFFF")
	tryRegister(&_vImageConvert_ARGBToYpCbCr_GenerateConversion, lib, "vImageConvert_ARGBToYpCbCr_GenerateConversion")
	tryRegister(&_vImageConvert_YpCbCrToARGB_GenerateConversion, lib, "vImageConvert_YpCbCrToARGB_GenerateConversion")
	tryRegister(&_vImageConvolveWithBias_ARGB8888, lib, "vImageConvolveWithBias_ARGB8888")
	tryRegister(&_vImageDilate_ARGB8888, lib, "vImageDilate_ARGB8888")
	tryRegister(&_vImageMatrixMultiply_ARGB8888ToPlanar8, lib, "vImageMatrixMultiply_ARGB8888ToPlanar8")
	tryRegister(&_vImageMultidimensionalTable_Create, lib, "vImageMultidimensionalTable_Create")
	tryRegister(&_vImageMultidimensionalTable_Release, lib, "vImageMultidimensionalTable_Release")
	tryRegister(&_vImageMultidimensionalTable_Retain, lib, "vImageMultidimensionalTable_Retain")
	tryRegister(&_vImagePremultipliedAlphaBlend_ARGB8888, lib, "vImagePremultipliedAlphaBlend_ARGB8888")
	tryRegister(&_vImagePremultipliedConstAlphaBlend_ARGB8888, lib, "vImagePremultipliedConstAlphaBlend_ARGB8888")
	tryRegister(&_vImagePremultiplyData_RGBA8888, lib, "vImagePremultiplyData_RGBA8888")
	tryRegister(&_vImageUnpremultiplyData_ARGB8888, lib, "vImageUnpremultiplyData_ARGB8888")
	tryRegister(&_vA128Shift, lib, "vA128Shift")
	tryRegister(&_vLL128Shift, lib, "vLL128Shift")
	tryRegister(&_vLR128Shift, lib, "vLR128Shift")
	tryRegister(&_vS128Add, lib, "vS128Add")
	tryRegister(&_vS128AddS, lib, "vS128AddS")
	tryRegister(&_vS128Sub, lib, "vS128Sub")
	tryRegister(&_vS128SubS, lib, "vS128SubS")
	tryRegister(&_vS64FullMulOdd, lib, "vS64FullMulOdd")
	tryRegister(&_vS64SubS, lib, "vS64SubS")
	tryRegister(&_vU128Add, lib, "vU128Add")
	tryRegister(&_vU128AddS, lib, "vU128AddS")
	tryRegister(&_vU128Sub, lib, "vU128Sub")
	tryRegister(&_vU128SubS, lib, "vU128SubS")
	tryRegister(&_vU64FullMulOdd, lib, "vU64FullMulOdd")
	tryRegister(&_vvacos, lib, "vvacos")
	tryRegister(&_vvacosf, lib, "vvacosf")
	tryRegister(&_vvacosh, lib, "vvacosh")
	tryRegister(&_vvacoshf, lib, "vvacoshf")
	tryRegister(&_vvasin, lib, "vvasin")
	tryRegister(&_vvasinf, lib, "vvasinf")
	tryRegister(&_vvasinh, lib, "vvasinh")
	tryRegister(&_vvasinhf, lib, "vvasinhf")
	tryRegister(&_vvatan, lib, "vvatan")
	tryRegister(&_vvatan2, lib, "vvatan2")
	tryRegister(&_vvatan2f, lib, "vvatan2f")
	tryRegister(&_vvatanf, lib, "vvatanf")
	tryRegister(&_vvatanh, lib, "vvatanh")
	tryRegister(&_vvatanhf, lib, "vvatanhf")
	tryRegister(&_vvceil, lib, "vvceil")
	tryRegister(&_vvceilf, lib, "vvceilf")
	tryRegister(&_vvcopysign, lib, "vvcopysign")
	tryRegister(&_vvcopysignf, lib, "vvcopysignf")
	tryRegister(&_vvcos, lib, "vvcos")
	tryRegister(&_vvcosf, lib, "vvcosf")
	tryRegister(&_vvcosh, lib, "vvcosh")
	tryRegister(&_vvcoshf, lib, "vvcoshf")
	tryRegister(&_vvcosisin, lib, "vvcosisin")
	tryRegister(&_vvcosisinf, lib, "vvcosisinf")
	tryRegister(&_vvcospi, lib, "vvcospi")
	tryRegister(&_vvcospif, lib, "vvcospif")
	tryRegister(&_vvdiv, lib, "vvdiv")
	tryRegister(&_vvdivf, lib, "vvdivf")
	tryRegister(&_vvexp, lib, "vvexp")
	tryRegister(&_vvexp2, lib, "vvexp2")
	tryRegister(&_vvexp2f, lib, "vvexp2f")
	tryRegister(&_vvexpf, lib, "vvexpf")
	tryRegister(&_vvexpm1, lib, "vvexpm1")
	tryRegister(&_vvexpm1f, lib, "vvexpm1f")
	tryRegister(&_vvfabs, lib, "vvfabs")
	tryRegister(&_vvfabsf, lib, "vvfabsf")
	tryRegister(&_vvfloor, lib, "vvfloor")
	tryRegister(&_vvfloorf, lib, "vvfloorf")
	tryRegister(&_vvfmod, lib, "vvfmod")
	tryRegister(&_vvfmodf, lib, "vvfmodf")
	tryRegister(&_vvint, lib, "vvint")
	tryRegister(&_vvintf, lib, "vvintf")
	tryRegister(&_vvlog, lib, "vvlog")
	tryRegister(&_vvlog10, lib, "vvlog10")
	tryRegister(&_vvlog10f, lib, "vvlog10f")
	tryRegister(&_vvlog1p, lib, "vvlog1p")
	tryRegister(&_vvlog1pf, lib, "vvlog1pf")
	tryRegister(&_vvlog2, lib, "vvlog2")
	tryRegister(&_vvlog2f, lib, "vvlog2f")
	tryRegister(&_vvlogb, lib, "vvlogb")
	tryRegister(&_vvlogbf, lib, "vvlogbf")
	tryRegister(&_vvlogf, lib, "vvlogf")
	tryRegister(&_vvnextafter, lib, "vvnextafter")
	tryRegister(&_vvnextafterf, lib, "vvnextafterf")
	tryRegister(&_vvnint, lib, "vvnint")
	tryRegister(&_vvnintf, lib, "vvnintf")
	tryRegister(&_vvpow, lib, "vvpow")
	tryRegister(&_vvpowf, lib, "vvpowf")
	tryRegister(&_vvrec, lib, "vvrec")
	tryRegister(&_vvrecf, lib, "vvrecf")
	tryRegister(&_vvremainder, lib, "vvremainder")
	tryRegister(&_vvremainderf, lib, "vvremainderf")
	tryRegister(&_vvrsqrt, lib, "vvrsqrt")
	tryRegister(&_vvrsqrtf, lib, "vvrsqrtf")
	tryRegister(&_vvsin, lib, "vvsin")
	tryRegister(&_vvsincos, lib, "vvsincos")
	tryRegister(&_vvsincosf, lib, "vvsincosf")
	tryRegister(&_vvsinf, lib, "vvsinf")
	tryRegister(&_vvsinh, lib, "vvsinh")
	tryRegister(&_vvsinhf, lib, "vvsinhf")
	tryRegister(&_vvsinpi, lib, "vvsinpi")
	tryRegister(&_vvsinpif, lib, "vvsinpif")
	tryRegister(&_vvsqrt, lib, "vvsqrt")
	tryRegister(&_vvsqrtf, lib, "vvsqrtf")
	tryRegister(&_vvtan, lib, "vvtan")
	tryRegister(&_vvtanf, lib, "vvtanf")
	tryRegister(&_vvtanh, lib, "vvtanh")
	tryRegister(&_vvtanhf, lib, "vvtanhf")
	tryRegister(&_vvtanpi, lib, "vvtanpi")
	tryRegister(&_vvtanpif, lib, "vvtanpif")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Applies a broadcast matrix multiplication operation directly to two input matrices.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 11.0.
// Applies a broadcast matrix multiplication operation directly to two input matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSDirectApplyBroadcastMatMul(_:_:_:_:_:_:_:)
func BNNSDirectApplyBroadcastMatMul(transA unsafe.Pointer, transB unsafe.Pointer, alpha unsafe.Pointer, inputA unsafe.Pointer, inputB unsafe.Pointer, output unsafe.Pointer, filter_params unsafe.Pointer) {
	_BNNSDirectApplyBroadcastMatMul(transA, transB, alpha, inputA, inputB, output, filter_params)
}

// Returns a new broadcast matrix multiply layer.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 11.0.
// Returns a new broadcast matrix multiply layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerBroadcastMatMul(_:_:)
func BNNSFilterCreateLayerBroadcastMatMul(layer_params unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSFilterCreateLayerBroadcastMatMul(layer_params, filter_params)
}

// Specifies that the context checks intermediate tensors for NaNs and infinities.
//
// Added in macOS 15.0.
// Specifies that the context checks intermediate tensors for NaNs and infinities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextEnableNanAndInfChecks(_:_:)
func BNNSGraphContextEnableNanAndInfChecks(context unsafe.Pointer, enable_check_for_nans_inf bool) {
	_BNNSGraphContextEnableNanAndInfChecks(context, enable_check_for_nans_inf)
}

// Applies a matrix multiplication operation directly to two input matrices.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 13.0.
// Applies a matrix multiplication operation directly to two input matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSMatMul(_:_:_:_:_:_:_:_:)
func BNNSMatMul(transA unsafe.Pointer, transB unsafe.Pointer, alpha unsafe.Pointer, inputA unsafe.Pointer, inputB unsafe.Pointer, output unsafe.Pointer, workspace unsafe.Pointer, filter_params unsafe.Pointer) int {
	return _BNNSMatMul(transA, transB, alpha, inputA, inputB, output, workspace, filter_params)
}

// Returns the workspace size that a matrix multiply operation requires.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 13.0.
// Returns the workspace size that a matrix multiply operation requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSMatMulWorkspaceSize(_:_:_:_:_:_:_:)
func BNNSMatMulWorkspaceSize(transA unsafe.Pointer, transB unsafe.Pointer, alpha unsafe.Pointer, inputA unsafe.Pointer, inputB unsafe.Pointer, output unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSMatMulWorkspaceSize(transA, transB, alpha, inputA, inputB, output, filter_params)
}

// Multiplies two matrices (single-precision).
//
// Added in macOS 13.3.
// Multiplies two matrices (single-precision).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/cblas_sgemm(_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func cblas_sgemm(ORDER unsafe.Pointer, TRANSA unsafe.Pointer, TRANSB unsafe.Pointer, M unsafe.Pointer, N unsafe.Pointer, K unsafe.Pointer, ALPHA unsafe.Pointer, A []float32, LDA unsafe.Pointer, B []float32, LDB unsafe.Pointer, BETA unsafe.Pointer, C []float32, LDC unsafe.Pointer) {
	_cblas_sgemm(ORDER, TRANSA, TRANSB, M, N, K, ALPHA, A, LDA, B, LDB, BETA, C, LDC)
}

// sparse_inner_product_dense_double_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_inner_product_dense_double_complex
func sparse_inner_product_dense_double_complex(nz unsafe.Pointer, x unsafe.Pointer, indx unsafe.Pointer, y unsafe.Pointer, incy unsafe.Pointer) unsafe.Pointer {
	return _sparse_inner_product_dense_double_complex(nz, x, indx, y, incy)
}

// sparse_inner_product_dense_float_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_inner_product_dense_float_complex
func sparse_inner_product_dense_float_complex(nz unsafe.Pointer, x unsafe.Pointer, indx unsafe.Pointer, y unsafe.Pointer, incy unsafe.Pointer) unsafe.Pointer {
	return _sparse_inner_product_dense_float_complex(nz, x, indx, y, incy)
}

// sparse_inner_product_sparse_double_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_inner_product_sparse_double_complex
func sparse_inner_product_sparse_double_complex(nzx unsafe.Pointer, nzy unsafe.Pointer, x unsafe.Pointer, indx unsafe.Pointer, y unsafe.Pointer, indy unsafe.Pointer) unsafe.Pointer {
	return _sparse_inner_product_sparse_double_complex(nzx, nzy, x, indx, y, indy)
}

// sparse_inner_product_sparse_float_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_inner_product_sparse_float_complex
func sparse_inner_product_sparse_float_complex(nzx unsafe.Pointer, nzy unsafe.Pointer, x unsafe.Pointer, indx unsafe.Pointer, y unsafe.Pointer, indy unsafe.Pointer) unsafe.Pointer {
	return _sparse_inner_product_sparse_float_complex(nzx, nzy, x, indx, y, indy)
}

// sparse_insert_entry_double_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_insert_entry_double_complex
func sparse_insert_entry_double_complex(A unsafe.Pointer, val unsafe.Pointer, i unsafe.Pointer, j unsafe.Pointer) unsafe.Pointer {
	return _sparse_insert_entry_double_complex(A, val, i, j)
}

// sparse_insert_entry_float_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_insert_entry_float_complex
func sparse_insert_entry_float_complex(A unsafe.Pointer, val unsafe.Pointer, i unsafe.Pointer, j unsafe.Pointer) unsafe.Pointer {
	return _sparse_insert_entry_float_complex(A, val, i, j)
}

// sparse_matrix_product_dense_double_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_product_dense_double_complex
func sparse_matrix_product_dense_double_complex(order unsafe.Pointer, transa unsafe.Pointer, n unsafe.Pointer, alpha unsafe.Pointer, A unsafe.Pointer, B unsafe.Pointer, ldb unsafe.Pointer, C unsafe.Pointer, ldc unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_product_dense_double_complex(order, transa, n, alpha, A, B, ldb, C, ldc)
}

// sparse_matrix_product_dense_float_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_product_dense_float_complex
func sparse_matrix_product_dense_float_complex(order unsafe.Pointer, transa unsafe.Pointer, n unsafe.Pointer, alpha unsafe.Pointer, A unsafe.Pointer, B unsafe.Pointer, ldb unsafe.Pointer, C unsafe.Pointer, ldc unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_product_dense_float_complex(order, transa, n, alpha, A, B, ldb, C, ldc)
}

// sparse_matrix_product_sparse_double_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_product_sparse_double_complex
func sparse_matrix_product_sparse_double_complex(order unsafe.Pointer, transa unsafe.Pointer, alpha unsafe.Pointer, A unsafe.Pointer, B unsafe.Pointer, C unsafe.Pointer, ldc unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_product_sparse_double_complex(order, transa, alpha, A, B, C, ldc)
}

// sparse_matrix_product_sparse_float_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_product_sparse_float_complex
func sparse_matrix_product_sparse_float_complex(order unsafe.Pointer, transa unsafe.Pointer, alpha unsafe.Pointer, A unsafe.Pointer, B unsafe.Pointer, C unsafe.Pointer, ldc unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_product_sparse_float_complex(order, transa, alpha, A, B, C, ldc)
}

// sparse_matrix_trace_double_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_trace_double_complex
func sparse_matrix_trace_double_complex(A unsafe.Pointer, offset unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_trace_double_complex(A, offset)
}

// sparse_matrix_trace_float_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_trace_float_complex
func sparse_matrix_trace_float_complex(A unsafe.Pointer, offset unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_trace_float_complex(A, offset)
}

// sparse_matrix_triangular_solve_dense_double_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_triangular_solve_dense_double_complex
func sparse_matrix_triangular_solve_dense_double_complex(order unsafe.Pointer, transt unsafe.Pointer, nrhs unsafe.Pointer, alpha unsafe.Pointer, T unsafe.Pointer, B unsafe.Pointer, ldb unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_triangular_solve_dense_double_complex(order, transt, nrhs, alpha, T, B, ldb)
}

// sparse_matrix_triangular_solve_dense_float_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_triangular_solve_dense_float_complex
func sparse_matrix_triangular_solve_dense_float_complex(order unsafe.Pointer, transt unsafe.Pointer, nrhs unsafe.Pointer, alpha unsafe.Pointer, T unsafe.Pointer, B unsafe.Pointer, ldb unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_triangular_solve_dense_float_complex(order, transt, nrhs, alpha, T, B, ldb)
}

// sparse_matrix_vector_product_dense_double_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_vector_product_dense_double_complex
func sparse_matrix_vector_product_dense_double_complex(transa unsafe.Pointer, alpha unsafe.Pointer, A unsafe.Pointer, x unsafe.Pointer, incx unsafe.Pointer, y unsafe.Pointer, incy unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_vector_product_dense_double_complex(transa, alpha, A, x, incx, y, incy)
}

// sparse_matrix_vector_product_dense_float_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_vector_product_dense_float_complex
func sparse_matrix_vector_product_dense_float_complex(transa unsafe.Pointer, alpha unsafe.Pointer, A unsafe.Pointer, x unsafe.Pointer, incx unsafe.Pointer, y unsafe.Pointer, incy unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_vector_product_dense_float_complex(transa, alpha, A, x, incx, y, incy)
}

// sparse_outer_product_dense_double_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_outer_product_dense_double_complex
func sparse_outer_product_dense_double_complex(M unsafe.Pointer, N unsafe.Pointer, nz unsafe.Pointer, alpha unsafe.Pointer, x unsafe.Pointer, incx unsafe.Pointer, y unsafe.Pointer, indy unsafe.Pointer, C unsafe.Pointer) unsafe.Pointer {
	return _sparse_outer_product_dense_double_complex(M, N, nz, alpha, x, incx, y, indy, C)
}

// sparse_outer_product_dense_float_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_outer_product_dense_float_complex
func sparse_outer_product_dense_float_complex(M unsafe.Pointer, N unsafe.Pointer, nz unsafe.Pointer, alpha unsafe.Pointer, x unsafe.Pointer, incx unsafe.Pointer, y unsafe.Pointer, indy unsafe.Pointer, C unsafe.Pointer) unsafe.Pointer {
	return _sparse_outer_product_dense_float_complex(M, N, nz, alpha, x, incx, y, indy, C)
}

// sparse_vector_add_with_scale_dense_double_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_vector_add_with_scale_dense_double_complex
func sparse_vector_add_with_scale_dense_double_complex(nz unsafe.Pointer, alpha unsafe.Pointer, x unsafe.Pointer, indx unsafe.Pointer, y unsafe.Pointer, incy unsafe.Pointer) {
	_sparse_vector_add_with_scale_dense_double_complex(nz, alpha, x, indx, y, incy)
}

// sparse_vector_add_with_scale_dense_float_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_vector_add_with_scale_dense_float_complex
func sparse_vector_add_with_scale_dense_float_complex(nz unsafe.Pointer, alpha unsafe.Pointer, x unsafe.Pointer, indx unsafe.Pointer, y unsafe.Pointer, incy unsafe.Pointer) {
	_sparse_vector_add_with_scale_dense_float_complex(nz, alpha, x, indx, y, incy)
}

// sparse_vector_triangular_solve_dense_double_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_vector_triangular_solve_dense_double_complex
func sparse_vector_triangular_solve_dense_double_complex(transt unsafe.Pointer, alpha unsafe.Pointer, T unsafe.Pointer, x unsafe.Pointer, incx unsafe.Pointer) unsafe.Pointer {
	return _sparse_vector_triangular_solve_dense_double_complex(transt, alpha, T, x, incx)
}

// sparse_vector_triangular_solve_dense_float_complex is a Accelerate function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_vector_triangular_solve_dense_float_complex
func sparse_vector_triangular_solve_dense_float_complex(transt unsafe.Pointer, alpha unsafe.Pointer, T unsafe.Pointer, x unsafe.Pointer, incx unsafe.Pointer) unsafe.Pointer {
	return _sparse_vector_triangular_solve_dense_float_complex(transt, alpha, T, x, incx)
}

// Releases a single-precision discrete Fourier transform (DFT) setup structure.
//
// Added in macOS 12.0.
// Releases a single-precision discrete Fourier transform (DFT) setup structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Interleaved_DestroySetup(_:)
func vDSP_DFT_Interleaved_DestroySetup(Setup unsafe.Pointer) {
	_vDSP_DFT_Interleaved_DestroySetup(Setup)
}

// Applies a single-precision single-channel biquadratic IIR filter.
//
// Added in macOS 10.9.
// Applies a single-precision single-channel biquadratic IIR filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquad
func vDSP_biquad(__Setup unsafe.Pointer, __Delay []float32, __X []float32, __IX unsafe.Pointer, __Y []float32, __IY unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_biquad(__Setup, __Delay, __X, __IX, __Y, __IY, __N)
}

// Builds a data structure that contains precalculated data for use by a single-precision cascaded biquadratic filter function.
//
// Added in macOS 10.9.
// Builds a data structure that contains precalculated data for use by a single-precision cascaded biquadratic filter function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquad_CreateSetup
func vDSP_biquad_CreateSetup(__Coefficients []float64, __M unsafe.Pointer) unsafe.Pointer {
	return _vDSP_biquad_CreateSetup(__Coefficients, __M)
}

// Sets single-precision coefficients of the specified single-channel biquadratic filter setup object.
//
// Added in macOS 12.0.
// Sets single-precision coefficients of the specified single-channel biquadratic filter setup object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquad_SetCoefficientsSingle
func vDSP_biquad_SetCoefficientsSingle(__setup unsafe.Pointer, __coeffs []float32, __start_sec unsafe.Pointer, __nsec unsafe.Pointer) {
	_vDSP_biquad_SetCoefficientsSingle(__setup, __coeffs, __start_sec, __nsec)
}

// Computes a 2D forward or inverse in-place, single-precision complex FFT.
//
// Added in macOS 10.0.
// Computes a 2D forward or inverse in-place, single-precision complex FFT.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_fft2d_zip
func vDSP_fft2d_zip(__Setup unsafe.Pointer, __C unsafe.Pointer, __IC0 unsafe.Pointer, __IC1 unsafe.Pointer, __Log2N0 unsafe.Pointer, __Log2N1 unsafe.Pointer, __Direction unsafe.Pointer) {
	_vDSP_fft2d_zip(__Setup, __C, __IC0, __IC1, __Log2N0, __Log2N1, __Direction)
}

// Computes a forward or inverse in-place, single-precision real FFT on multiple signals.
//
// Added in macOS 10.2.
// Computes a forward or inverse in-place, single-precision real FFT on multiple signals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_fftm_zrip
func vDSP_fftm_zrip(__Setup unsafe.Pointer, __C unsafe.Pointer, __IC unsafe.Pointer, __IM unsafe.Pointer, __Log2N unsafe.Pointer, __M unsafe.Pointer, __Direction unsafe.Pointer) {
	_vDSP_fftm_zrip(__Setup, __C, __IC, __IM, __Log2N, __M, __Direction)
}

// Performs a tapered merge between two single-precision vectors.
//
// Added in macOS 10.4.
// Performs a tapered merge between two single-precision vectors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_vtmerg
func vDSP_vtmerg(__A []float32, __IA unsafe.Pointer, __B []float32, __IB unsafe.Pointer, __C []float32, __IC unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_vtmerg(__A, __IA, __B, __IB, __C, __IC, __N)
}

// Copies the contents of a split single-precision complex vector to an interleaved vector.
//
// Added in macOS 10.0.
// Copies the contents of a split single-precision complex vector to an interleaved vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_ztoc
func vDSP_ztoc(__Z unsafe.Pointer, __IZ unsafe.Pointer, __C unsafe.Pointer, __IC unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_ztoc(__Z, __IZ, __C, __IC, __N)
}

// Performs nonpremultiplied alpha compositing of two 8-bit-per-channel, 4-channel ARGB buffers.
//
// Added in macOS 10.3.
// Performs nonpremultiplied alpha compositing of two 8-bit-per-channel, 4-channel ARGB buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageAlphaBlend_ARGB8888(_:_:_:_:)
func vImageAlphaBlend_ARGB8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageAlphaBlend_ARGB8888(srcTop, srcBottom, dest, flags)
}

// Initializes a vImage buffer with a copy of the contents of a Core Video pixel buffer.
//
// Added in macOS 10.10.
// Initializes a vImage buffer with a copy of the contents of a Core Video pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageBuffer_InitWithCVPixelBuffer(_:_:_:_:_:_:)
func vImageBuffer_InitWithCVPixelBuffer(buffer unsafe.Pointer, desiredFormat unsafe.Pointer, cvPixelBuffer unsafe.Pointer, cvImageFormat unsafe.Pointer, backgroundColor []float64, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageBuffer_InitWithCVPixelBuffer(buffer, desiredFormat, cvPixelBuffer, cvImageFormat, backgroundColor, flags)
}

// Performs contrast stretching on a 32-bit-per-channel, 4-channel interleaved buffer.
//
// Added in macOS 10.3.
// Performs contrast stretching on a 32-bit-per-channel, 4-channel interleaved buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageContrastStretch_ARGBFFFF(_:_:_:_:_:_:_:)
func vImageContrastStretch_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, histogram_entries unsafe.Pointer, minVal unsafe.Pointer, maxVal unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageContrastStretch_ARGBFFFF(src, dest, tempBuffer, histogram_entries, minVal, maxVal, flags)
}

// Generates the information that describes the conversion from ARGB to YpCbCr.
//
// Added in macOS 10.10.
// Generates the information that describes the conversion from ARGB to YpCbCr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGBToYpCbCr_GenerateConversion(_:_:_:_:_:_:)
func vImageConvert_ARGBToYpCbCr_GenerateConversion(matrix unsafe.Pointer, pixelRange unsafe.Pointer, outInfo unsafe.Pointer, inARGBType unsafe.Pointer, outYpCbCrType unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageConvert_ARGBToYpCbCr_GenerateConversion(matrix, pixelRange, outInfo, inARGBType, outYpCbCrType, flags)
}

// Generates the information that describes the conversion from YpCbCr to ARGB.
//
// Added in macOS 10.10.
// Generates the information that describes the conversion from YpCbCr to ARGB.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConvert_YpCbCrToARGB_GenerateConversion(_:_:_:_:_:_:)
func vImageConvert_YpCbCrToARGB_GenerateConversion(matrix unsafe.Pointer, pixelRange unsafe.Pointer, outInfo unsafe.Pointer, inYpCbCrType unsafe.Pointer, outARGBType unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageConvert_YpCbCrToARGB_GenerateConversion(matrix, pixelRange, outInfo, inYpCbCrType, outARGBType, flags)
}

// Convolves an 8-bit-per-channel, 4-channel interleaved image by a 2D kernel and adds a bias.
//
// Added in macOS 10.4.
// Convolves an 8-bit-per-channel, 4-channel interleaved image by a 2D kernel and adds a bias.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConvolveWithBias_ARGB8888(_:_:_:_:_:_:_:_:_:_:_:_:)
func vImageConvolveWithBias_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, kernel unsafe.Pointer, kernel_height uint32, kernel_width uint32, divisor unsafe.Pointer, bias unsafe.Pointer, backgroundColor unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageConvolveWithBias_ARGB8888(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel_height, kernel_width, divisor, bias, backgroundColor, flags)
}

// Dilates an 8-bit-per-channel, 4-channel interleaved buffer.
//
// Added in macOS 10.3.
// Dilates an 8-bit-per-channel, 4-channel interleaved buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageDilate_ARGB8888(_:_:_:_:_:_:_:_:)
func vImageDilate_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, kernel unsafe.Pointer, kernel_height unsafe.Pointer, kernel_width unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageDilate_ARGB8888(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel_height, kernel_width, flags)
}

// Multiplies each pixel in an interleaved four-channel, 8-bit source image by a matrix to produce a planar 8-bit destination image.
//
// Added in macOS 10.11.
// Multiplies each pixel in an interleaved four-channel, 8-bit source image by a matrix to produce a planar 8-bit destination image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMatrixMultiply_ARGB8888ToPlanar8(_:_:_:_:_:_:_:)
func vImageMatrixMultiply_ARGB8888ToPlanar8(src unsafe.Pointer, dest unsafe.Pointer, matrix unsafe.Pointer, divisor unsafe.Pointer, pre_bias unsafe.Pointer, post_bias unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageMatrixMultiply_ARGB8888ToPlanar8(src, dest, matrix, divisor, pre_bias, post_bias, flags)
}

// Creates a multidimensional lookup table.
//
// Added in macOS 10.9.
// Creates a multidimensional lookup table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMultidimensionalTable_Create(_:_:_:_:_:_:_:)
func vImageMultidimensionalTable_Create(tableData unsafe.Pointer, numSrcChannels uint32, numDestChannels uint32, table_entries_per_dimension unsafe.Pointer, hint unsafe.Pointer, flags unsafe.Pointer, err unsafe.Pointer) unsafe.Pointer {
	return _vImageMultidimensionalTable_Create(tableData, numSrcChannels, numDestChannels, table_entries_per_dimension, hint, flags, err)
}

// Releases a multidimensional table.
//
// Added in macOS 10.9.
// Releases a multidimensional table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMultidimensionalTable_Release(_:)
func vImageMultidimensionalTable_Release(table unsafe.Pointer) unsafe.Pointer {
	return _vImageMultidimensionalTable_Release(table)
}

// Retains a multidimensional table.
//
// Added in macOS 10.9.
// Retains a multidimensional table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMultidimensionalTable_Retain(_:)
func vImageMultidimensionalTable_Retain(table unsafe.Pointer) unsafe.Pointer {
	return _vImageMultidimensionalTable_Retain(table)
}

// Performs premultiplied alpha compositing of two 8-bit-per-channel, 4-channel ARGB buffers.
//
// Added in macOS 10.3.
// Performs premultiplied alpha compositing of two 8-bit-per-channel, 4-channel ARGB buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlend_ARGB8888(_:_:_:_:)
func vImagePremultipliedAlphaBlend_ARGB8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImagePremultipliedAlphaBlend_ARGB8888(srcTop, srcBottom, dest, flags)
}

// Performs premultiplied alpha compositing of two 8-bit-per-channel, 4-channel interleaved buffers and applies an extra alpha value to the top buffer.
//
// Added in macOS 10.4.
// Performs premultiplied alpha compositing of two 8-bit-per-channel, 4-channel interleaved buffers and applies an extra alpha value to the top buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedConstAlphaBlend_ARGB8888(_:_:_:_:_:)
func vImagePremultipliedConstAlphaBlend_ARGB8888(srcTop unsafe.Pointer, constAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImagePremultipliedConstAlphaBlend_ARGB8888(srcTop, constAlpha, srcBottom, dest, flags)
}

// Transforms an 8-bit-per-channel, 4-channel RGBA buffer from nonpremultiplied alpha format to premultiplied alpha format.
//
// Added in macOS 10.4.
// Transforms an 8-bit-per-channel, 4-channel RGBA buffer from nonpremultiplied alpha format to premultiplied alpha format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImagePremultiplyData_RGBA8888(_:_:_:)
func vImagePremultiplyData_RGBA8888(src unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImagePremultiplyData_RGBA8888(src, dest, flags)
}

// Transforms an 8-bit-per-channel, 4-channel ARGB buffer from premultiplied alpha format to nonpremultiplied alpha format.
//
// Added in macOS 10.3.
// Transforms an 8-bit-per-channel, 4-channel ARGB buffer from premultiplied alpha format to nonpremultiplied alpha format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageUnpremultiplyData_ARGB8888(_:_:_:)
func vImageUnpremultiplyData_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageUnpremultiplyData_ARGB8888(src, dest, flags)
}

// 128-bit arithmetic (signed) shift.
//
// Added in macOS 10.0.
// 128-bit arithmetic (signed) shift.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vA128Shift(_:_:)
func vA128Shift(vA unsafe.Pointer, vShiftFactor unsafe.Pointer) unsafe.Pointer {
	return _vA128Shift(vA, vShiftFactor)
}

// 128-bit logical left shift.
//
// Added in macOS 10.5.
// 128-bit logical left shift.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vLL128Shift(_:_:)
func vLL128Shift(vA unsafe.Pointer, vShiftFactor unsafe.Pointer) unsafe.Pointer {
	return _vLL128Shift(vA, vShiftFactor)
}

// 128-bit logical right shift.
//
// Added in macOS 10.5.
// 128-bit logical right shift.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vLR128Shift(_:_:)
func vLR128Shift(vA unsafe.Pointer, vShiftFactor unsafe.Pointer) unsafe.Pointer {
	return _vLR128Shift(vA, vShiftFactor)
}

// Signed 128-bit addition (modular arithmetic).
//
// Added in macOS 10.0.
// Signed 128-bit addition (modular arithmetic).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vS128Add(_:_:)
func vS128Add(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vS128Add(vA, vB)
}

// Signed 128-bit addition with saturation (clipping).
//
// Added in macOS 10.0.
// Signed 128-bit addition with saturation (clipping).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vS128AddS(_:_:)
func vS128AddS(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vS128AddS(vA, vB)
}

// Signed 128-bit subtraction (modular arithmetic).
//
// Added in macOS 10.0.
// Signed 128-bit subtraction (modular arithmetic).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vS128Sub(_:_:)
func vS128Sub(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vS128Sub(vA, vB)
}

// Signed 128-bit subtraction with saturation (clipping).
//
// Added in macOS 10.0.
// Signed 128-bit subtraction with saturation (clipping).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vS128SubS(_:_:)
func vS128SubS(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vS128SubS(vA, vB)
}

// Signed 64-bit multiplication; results are twice as wide as multiplicands, odd-numbered elements of multiplicand vectors are used. Note the big-endian convention: the leftmost element is element 0.
//
// Added in macOS 10.0.
// Signed 64-bit multiplication; results are twice as wide as multiplicands, odd-numbered elements of multiplicand vectors are used. Note the big-endian convention: the leftmost element is element 0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vS64FullMulOdd(_:_:)
func vS64FullMulOdd(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vS64FullMulOdd(vA, vB)
}

// Signed 64-bit subtraction with saturation (clipping).
//
// Added in macOS 10.0.
// Signed 64-bit subtraction with saturation (clipping).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vS64SubS(_:_:)
func vS64SubS(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vS64SubS(vA, vB)
}

// Unsigned 128-bit addition (modular arithmetic).
//
// Added in macOS 10.0.
// Unsigned 128-bit addition (modular arithmetic).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vU128Add(_:_:)
func vU128Add(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vU128Add(vA, vB)
}

// Unsigned 128-bit addition with saturation (clipping).
//
// Added in macOS 10.0.
// Unsigned 128-bit addition with saturation (clipping).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vU128AddS(_:_:)
func vU128AddS(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vU128AddS(vA, vB)
}

// Unsigned 128-bit subtraction (modular arithmetic).
//
// Added in macOS 10.0.
// Unsigned 128-bit subtraction (modular arithmetic).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vU128Sub(_:_:)
func vU128Sub(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vU128Sub(vA, vB)
}

// Unsigned 128-bit subtraction with saturation (clipping).
//
// Added in macOS 10.0.
// Unsigned 128-bit subtraction with saturation (clipping).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vU128SubS(_:_:)
func vU128SubS(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vU128SubS(vA, vB)
}

// Unsigned 64-bit multiplication; results are twice as wide as multiplicands, odd-numbered elements of multiplicand vectors are used. Note the big-endian convention: the leftmost element is element 0.
//
// Added in macOS 10.0.
// Unsigned 64-bit multiplication; results are twice as wide as multiplicands, odd-numbered elements of multiplicand vectors are used. Note the big-endian convention: the leftmost element is element 0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vU64FullMulOdd(_:_:)
func vU64FullMulOdd(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vU64FullMulOdd(vA, vB)
}

// Calculates the arccosine of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the arccosine of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvacos(_:_:_:)
func vvacos(p0 []float64, p1 []float64, p2 []int) {
	_vvacos(p0, p1, p2)
}

// Calculates the arccosine of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the arccosine of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvacosf(_:_:_:)
func vvacosf(p0 []float32, p1 []float32, p2 []int) {
	_vvacosf(p0, p1, p2)
}

// Calculates the inverse hyperbolic cosine of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the inverse hyperbolic cosine of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvacosh(_:_:_:)
func vvacosh(p0 []float64, p1 []float64, p2 []int) {
	_vvacosh(p0, p1, p2)
}

// Calculates the inverse hyperbolic cosine of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the inverse hyperbolic cosine of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvacoshf(_:_:_:)
func vvacoshf(p0 []float32, p1 []float32, p2 []int) {
	_vvacoshf(p0, p1, p2)
}

// Calculates the arcsine of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the arcsine of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvasin(_:_:_:)
func vvasin(p0 []float64, p1 []float64, p2 []int) {
	_vvasin(p0, p1, p2)
}

// Calculates the arcsine of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the arcsine of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvasinf(_:_:_:)
func vvasinf(p0 []float32, p1 []float32, p2 []int) {
	_vvasinf(p0, p1, p2)
}

// Calculates the inverse hyperbolic sine of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the inverse hyperbolic sine of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvasinh(_:_:_:)
func vvasinh(p0 []float64, p1 []float64, p2 []int) {
	_vvasinh(p0, p1, p2)
}

// Calculates the inverse hyperbolic sine of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the inverse hyperbolic sine of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvasinhf(_:_:_:)
func vvasinhf(p0 []float32, p1 []float32, p2 []int) {
	_vvasinhf(p0, p1, p2)
}

// Calculates the arctangent of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the arctangent of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvatan(_:_:_:)
func vvatan(p0 []float64, p1 []float64, p2 []int) {
	_vvatan(p0, p1, p2)
}

// Calculates the arctangent of each pair of elements in two arrays of double-precision values.
//
// Added in macOS 10.4.
// Calculates the arctangent of each pair of elements in two arrays of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvatan2(_:_:_:_:)
func vvatan2(p0 []float64, p1 []float64, p2 []float64, p3 []int) {
	_vvatan2(p0, p1, p2, p3)
}

// Calculates the arctangent of each pair of elements in two arrays of single-precision values.
//
// Added in macOS 10.4.
// Calculates the arctangent of each pair of elements in two arrays of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvatan2f(_:_:_:_:)
func vvatan2f(p0 []float32, p1 []float32, p2 []float32, p3 []int) {
	_vvatan2f(p0, p1, p2, p3)
}

// Calculates the arctangent of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the arctangent of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvatanf(_:_:_:)
func vvatanf(p0 []float32, p1 []float32, p2 []int) {
	_vvatanf(p0, p1, p2)
}

// Calculates the inverse hyperbolic tangent of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the inverse hyperbolic tangent of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvatanh(_:_:_:)
func vvatanh(p0 []float64, p1 []float64, p2 []int) {
	_vvatanh(p0, p1, p2)
}

// Calculates the inverse hyperbolic tangent of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the inverse hyperbolic tangent of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvatanhf(_:_:_:)
func vvatanhf(p0 []float32, p1 []float32, p2 []int) {
	_vvatanhf(p0, p1, p2)
}

// Calculates the ceiling of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the ceiling of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvceil(_:_:_:)
func vvceil(p0 []float64, p1 []float64, p2 []int) {
	_vvceil(p0, p1, p2)
}

// Calculates the ceiling of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the ceiling of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvceilf(_:_:_:)
func vvceilf(p0 []float32, p1 []float32, p2 []int) {
	_vvceilf(p0, p1, p2)
}

// Copies an array, setting the sign of each element based on a second array of double-precision values.
//
// Added in macOS 10.7.
// Copies an array, setting the sign of each element based on a second array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcopysign(_:_:_:_:)
func vvcopysign(p0 []float64, p1 []float64, p2 []float64, p3 []int) {
	_vvcopysign(p0, p1, p2, p3)
}

// Copies an array, setting the sign of each element based on a second array of single-precision values.
//
// Added in macOS 10.5.
// Copies an array, setting the sign of each element based on a second array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcopysignf(_:_:_:_:)
func vvcopysignf(p0 []float32, p1 []float32, p2 []float32, p3 []int) {
	_vvcopysignf(p0, p1, p2, p3)
}

// Calculates the cosine of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the cosine of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcos(_:_:_:)
func vvcos(p0 []float64, p1 []float64, p2 []int) {
	_vvcos(p0, p1, p2)
}

// Calculates the cosine of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the cosine of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcosf(_:_:_:)
func vvcosf(p0 []float32, p1 []float32, p2 []int) {
	_vvcosf(p0, p1, p2)
}

// Calculates the hyperbolic cosine of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the hyperbolic cosine of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcosh(_:_:_:)
func vvcosh(p0 []float64, p1 []float64, p2 []int) {
	_vvcosh(p0, p1, p2)
}

// Calculates the hyperbolic cosine of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the hyperbolic cosine of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcoshf(_:_:_:)
func vvcoshf(p0 []float32, p1 []float32, p2 []int) {
	_vvcoshf(p0, p1, p2)
}

// Calculates the cosine and sine of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the cosine and sine of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcosisin(_:_:_:)
func vvcosisin(p0 unsafe.Pointer, p1 []float64, p2 []int) {
	_vvcosisin(p0, p1, p2)
}

// Calculates the cosine and sine of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the cosine and sine of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcosisinf(_:_:_:)
func vvcosisinf(p0 unsafe.Pointer, p1 []float32, p2 []int) {
	_vvcosisinf(p0, p1, p2)
}

// Calculates the cosine of pi multiplied by each element in an array of double-precision values.
//
// Added in macOS 10.7.
// Calculates the cosine of pi multiplied by each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcospi(_:_:_:)
func vvcospi(p0 []float64, p1 []float64, p2 []int) {
	_vvcospi(p0, p1, p2)
}

// Calculates the cosine of pi multiplied by each element in an array of single-precision values.
//
// Added in macOS 10.7.
// Calculates the cosine of pi multiplied by each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcospif(_:_:_:)
func vvcospif(p0 []float32, p1 []float32, p2 []int) {
	_vvcospif(p0, p1, p2)
}

// Divides each element in an array by the corresponding value in a second array of double-precision values.
//
// Added in macOS 10.4.
// Divides each element in an array by the corresponding value in a second array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvdiv(_:_:_:_:)
func vvdiv(p0 []float64, p1 []float64, p2 []float64, p3 []int) {
	_vvdiv(p0, p1, p2, p3)
}

// Divides each element in an array by the corresponding value in a second array of single-precision values.
//
// Added in macOS 10.4.
// Divides each element in an array by the corresponding value in a second array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvdivf(_:_:_:_:)
func vvdivf(p0 []float32, p1 []float32, p2 []float32, p3 []int) {
	_vvdivf(p0, p1, p2, p3)
}

// Calculates raised to the power of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates raised to the power of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvexp(_:_:_:)
func vvexp(p0 []float64, p1 []float64, p2 []int) {
	_vvexp(p0, p1, p2)
}

// Calculates 2 raised to the power of each element in an array of double-precision values.
//
// Added in macOS 10.7.
// Calculates 2 raised to the power of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvexp2(_:_:_:)
func vvexp2(p0 []float64, p1 []float64, p2 []int) {
	_vvexp2(p0, p1, p2)
}

// Calculates 2 raised to the power of each element in an array of single-precision values.
//
// Added in macOS 10.7.
// Calculates 2 raised to the power of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvexp2f(_:_:_:)
func vvexp2f(p0 []float32, p1 []float32, p2 []int) {
	_vvexp2f(p0, p1, p2)
}

// Calculates raised to the power of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates raised to the power of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvexpf(_:_:_:)
func vvexpf(p0 []float32, p1 []float32, p2 []int) {
	_vvexpf(p0, p1, p2)
}

// Calculates for each element in an array of double-precision values.
//
// Added in macOS 10.7.
// Calculates for each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvexpm1(_:_:_:)
func vvexpm1(p0 []float64, p1 []float64, p2 []int) {
	_vvexpm1(p0, p1, p2)
}

// Calculates for each element in an array of single-precision values.
//
// Added in macOS 10.5.
// Calculates for each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvexpm1f(_:_:_:)
func vvexpm1f(p0 []float32, p1 []float32, p2 []int) {
	_vvexpm1f(p0, p1, p2)
}

// Calculates the absolute value for each element in an array of double-precision values.
//
// Added in macOS 10.7.
// Calculates the absolute value for each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvfabs(_:_:_:)
func vvfabs(p0 []float64, p1 []float64, p2 []int) {
	_vvfabs(p0, p1, p2)
}

// Calculates the absolute value for each element in an array of single-precision values.
//
// Added in macOS 10.7.
// Calculates the absolute value for each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvfabsf(_:_:_:)
func vvfabsf(p0 []float32, p1 []float32, p2 []int) {
	_vvfabsf(p0, p1, p2)
}

// Calculates the floor of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the floor of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvfloor(_:_:_:)
func vvfloor(p0 []float64, p1 []float64, p2 []int) {
	_vvfloor(p0, p1, p2)
}

// Calculates the floor of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the floor of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvfloorf(_:_:_:)
func vvfloorf(p0 []float32, p1 []float32, p2 []int) {
	_vvfloorf(p0, p1, p2)
}

// Calculates the modulus after dividing each element in an array by the corresponding element in a second array of double-precision values.
//
// Added in macOS 10.7.
// Calculates the modulus after dividing each element in an array by the corresponding element in a second array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvfmod(_:_:_:_:)
func vvfmod(p0 []float64, p1 []float64, p2 []float64, p3 []int) {
	_vvfmod(p0, p1, p2, p3)
}

// Calculates the modulus after dividing each element in an array by the corresponding element in a second array of single-precision values.
//
// Added in macOS 10.5.
// Calculates the modulus after dividing each element in an array by the corresponding element in a second array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvfmodf(_:_:_:_:)
func vvfmodf(p0 []float32, p1 []float32, p2 []float32, p3 []int) {
	_vvfmodf(p0, p1, p2, p3)
}

// Calculates the integer truncation for each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the integer truncation for each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvint(_:_:_:)
func vvint(p0 []float64, p1 []float64, p2 []int) {
	_vvint(p0, p1, p2)
}

// Calculates the integer truncation for each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the integer truncation for each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvintf(_:_:_:)
func vvintf(p0 []float32, p1 []float32, p2 []int) {
	_vvintf(p0, p1, p2)
}

// Calculates the natural logarithm for each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the natural logarithm for each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlog(_:_:_:)
func vvlog(p0 []float64, p1 []float64, p2 []int) {
	_vvlog(p0, p1, p2)
}

// Calculates the base 10 logarithm of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the base 10 logarithm of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlog10(_:_:_:)
func vvlog10(p0 []float64, p1 []float64, p2 []int) {
	_vvlog10(p0, p1, p2)
}

// Calculates the base 10 logarithm of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the base 10 logarithm of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlog10f(_:_:_:)
func vvlog10f(p0 []float32, p1 []float32, p2 []int) {
	_vvlog10f(p0, p1, p2)
}

// Calculates for each element in an array of double-precision values.
//
// Added in macOS 10.7.
// Calculates for each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlog1p(_:_:_:)
func vvlog1p(p0 []float64, p1 []float64, p2 []int) {
	_vvlog1p(p0, p1, p2)
}

// Calculates for each element in an array of single-precision values.
//
// Added in macOS 10.5.
// Calculates for each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlog1pf(_:_:_:)
func vvlog1pf(p0 []float32, p1 []float32, p2 []int) {
	_vvlog1pf(p0, p1, p2)
}

// Calculates the base 2 logarithm of each element in an array of double-precision values.
//
// Added in macOS 10.7.
// Calculates the base 2 logarithm of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlog2(_:_:_:)
func vvlog2(p0 []float64, p1 []float64, p2 []int) {
	_vvlog2(p0, p1, p2)
}

// Calculates the base 2 logarithm of each element in an array of single-precision values.
//
// Added in macOS 10.7.
// Calculates the base 2 logarithm of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlog2f(_:_:_:)
func vvlog2f(p0 []float32, p1 []float32, p2 []int) {
	_vvlog2f(p0, p1, p2)
}

// Calculates the unbiased exponent of each element in an array of double-precision values.
//
// Added in macOS 10.7.
// Calculates the unbiased exponent of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlogb(_:_:_:)
func vvlogb(p0 []float64, p1 []float64, p2 []int) {
	_vvlogb(p0, p1, p2)
}

// Calculates the unbiased exponent of each element in an array of single-precision values.
//
// Added in macOS 10.5.
// Calculates the unbiased exponent of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlogbf(_:_:_:)
func vvlogbf(p0 []float32, p1 []float32, p2 []int) {
	_vvlogbf(p0, p1, p2)
}

// Calculates the natural logarithm for each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the natural logarithm for each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlogf(_:_:_:)
func vvlogf(p0 []float32, p1 []float32, p2 []int) {
	_vvlogf(p0, p1, p2)
}

// Calculates the next machine-representable value for each element in an array of double-precision values.
//
// Added in macOS 10.7.
// Calculates the next machine-representable value for each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvnextafter(_:_:_:_:)
func vvnextafter(p0 []float64, p1 []float64, p2 []float64, p3 []int) {
	_vvnextafter(p0, p1, p2, p3)
}

// Calculates the next machine-representable value for each element in an array of single-precision values.
//
// Added in macOS 10.5.
// Calculates the next machine-representable value for each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvnextafterf(_:_:_:_:)
func vvnextafterf(p0 []float32, p1 []float32, p2 []float32, p3 []int) {
	_vvnextafterf(p0, p1, p2, p3)
}

// Calculates the nearest integer for each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the nearest integer for each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvnint(_:_:_:)
func vvnint(p0 []float64, p1 []float64, p2 []int) {
	_vvnint(p0, p1, p2)
}

// Calculates the nearest integer for each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the nearest integer for each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvnintf(_:_:_:)
func vvnintf(p0 []float32, p1 []float32, p2 []int) {
	_vvnintf(p0, p1, p2)
}

// Raises each element in an array to the power of the corresponding element in a second array of double-precision values.
//
// Added in macOS 10.4.
// Raises each element in an array to the power of the corresponding element in a second array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvpow(_:_:_:_:)
func vvpow(p0 []float64, p1 []float64, p2 []float64, p3 []int) {
	_vvpow(p0, p1, p2, p3)
}

// Raises each element in an array to the power of the corresponding element in a second array of single-precision values.
//
// Added in macOS 10.4.
// Raises each element in an array to the power of the corresponding element in a second array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvpowf(_:_:_:_:)
func vvpowf(p0 []float32, p1 []float32, p2 []float32, p3 []int) {
	_vvpowf(p0, p1, p2, p3)
}

// Calculates the reciprocal of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the reciprocal of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvrec(_:_:_:)
func vvrec(p0 []float64, p1 []float64, p2 []int) {
	_vvrec(p0, p1, p2)
}

// Calculates the reciprocal of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the reciprocal of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvrecf(_:_:_:)
func vvrecf(p0 []float32, p1 []float32, p2 []int) {
	_vvrecf(p0, p1, p2)
}

// Calculates the remainder after dividing each element in an array by the corresponding element in a second array of double-precision values.
//
// Added in macOS 10.7.
// Calculates the remainder after dividing each element in an array by the corresponding element in a second array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvremainder(_:_:_:_:)
func vvremainder(p0 []float64, p1 []float64, p2 []float64, p3 []int) {
	_vvremainder(p0, p1, p2, p3)
}

// Calculates the remainder after dividing each element in an array by the corresponding element in a second array of single-precision values.
//
// Added in macOS 10.5.
// Calculates the remainder after dividing each element in an array by the corresponding element in a second array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvremainderf(_:_:_:_:)
func vvremainderf(p0 []float32, p1 []float32, p2 []float32, p3 []int) {
	_vvremainderf(p0, p1, p2, p3)
}

// Calculates the reciprocal square root of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the reciprocal square root of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvrsqrt(_:_:_:)
func vvrsqrt(p0 []float64, p1 []float64, p2 []int) {
	_vvrsqrt(p0, p1, p2)
}

// Calculates the reciprocal square root of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the reciprocal square root of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvrsqrtf(_:_:_:)
func vvrsqrtf(p0 []float32, p1 []float32, p2 []int) {
	_vvrsqrtf(p0, p1, p2)
}

// Calculates the sine of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the sine of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsin(_:_:_:)
func vvsin(p0 []float64, p1 []float64, p2 []int) {
	_vvsin(p0, p1, p2)
}

// Calculates the cosine and sine of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the cosine and sine of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsincos(_:_:_:_:)
func vvsincos(p0 []float64, p1 []float64, p2 []float64, p3 []int) {
	_vvsincos(p0, p1, p2, p3)
}

// Calculates the cosine and sine of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the cosine and sine of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsincosf(_:_:_:_:)
func vvsincosf(p0 []float32, p1 []float32, p2 []float32, p3 []int) {
	_vvsincosf(p0, p1, p2, p3)
}

// Calculates the sine of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the sine of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsinf(_:_:_:)
func vvsinf(p0 []float32, p1 []float32, p2 []int) {
	_vvsinf(p0, p1, p2)
}

// Calculates the hyperbolic sine of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the hyperbolic sine of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsinh(_:_:_:)
func vvsinh(p0 []float64, p1 []float64, p2 []int) {
	_vvsinh(p0, p1, p2)
}

// Calculates the hyperbolic sine of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the hyperbolic sine of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsinhf(_:_:_:)
func vvsinhf(p0 []float32, p1 []float32, p2 []int) {
	_vvsinhf(p0, p1, p2)
}

// Calculates the sine of pi multiplied by each element in an array of double-precision values.
//
// Added in macOS 10.7.
// Calculates the sine of pi multiplied by each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsinpi(_:_:_:)
func vvsinpi(p0 []float64, p1 []float64, p2 []int) {
	_vvsinpi(p0, p1, p2)
}

// Calculates the sine of pi multiplied by each element in an array of single-precision values.
//
// Added in macOS 10.7.
// Calculates the sine of pi multiplied by each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsinpif(_:_:_:)
func vvsinpif(p0 []float32, p1 []float32, p2 []int) {
	_vvsinpif(p0, p1, p2)
}

// Calculates the square root of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the square root of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsqrt(_:_:_:)
func vvsqrt(p0 []float64, p1 []float64, p2 []int) {
	_vvsqrt(p0, p1, p2)
}

// Calculates the square root of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the square root of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsqrtf(_:_:_:)
func vvsqrtf(p0 []float32, p1 []float32, p2 []int) {
	_vvsqrtf(p0, p1, p2)
}

// Calculates the tangent of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the tangent of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvtan(_:_:_:)
func vvtan(p0 []float64, p1 []float64, p2 []int) {
	_vvtan(p0, p1, p2)
}

// Calculates the tangent of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the tangent of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvtanf(_:_:_:)
func vvtanf(p0 []float32, p1 []float32, p2 []int) {
	_vvtanf(p0, p1, p2)
}

// Calculates the hyperbolic tangent of each element in an array of double-precision values.
//
// Added in macOS 10.4.
// Calculates the hyperbolic tangent of each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvtanh(_:_:_:)
func vvtanh(p0 []float64, p1 []float64, p2 []int) {
	_vvtanh(p0, p1, p2)
}

// Calculates the hyperbolic tangent of each element in an array of single-precision values.
//
// Added in macOS 10.4.
// Calculates the hyperbolic tangent of each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvtanhf(_:_:_:)
func vvtanhf(p0 []float32, p1 []float32, p2 []int) {
	_vvtanhf(p0, p1, p2)
}

// Calculates the tangent of pi multiplied by each element in an array of double-precision values.
//
// Added in macOS 10.7.
// Calculates the tangent of pi multiplied by each element in an array of double-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvtanpi(_:_:_:)
func vvtanpi(p0 []float64, p1 []float64, p2 []int) {
	_vvtanpi(p0, p1, p2)
}

// Calculates the tangent of pi multiplied by each element in an array of single-precision values.
//
// Added in macOS 10.7.
// Calculates the tangent of pi multiplied by each element in an array of single-precision values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvtanpif(_:_:_:)
func vvtanpif(p0 []float32, p1 []float32, p2 []int) {
	_vvtanpif(p0, p1, p2)
}



