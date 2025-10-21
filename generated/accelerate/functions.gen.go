// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

package accelerate

import (
	"unsafe"

	"github.com/ebitengine/purego"
	coregraphics "github.com/tmc/appledocs/generated/coregraphics"
)


// Accelerate Functions (255 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_BNNSBandPart func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSClipByGlobalNorm func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSClipByNorm func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSClipByValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSComputeNorm func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSComputeNormBackward func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSCreateRandomGenerator func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSCreateRandomGeneratorWithSeed func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSDataLayoutGetRank func(unsafe.Pointer) unsafe.Pointer
	_BNNSDestroyRandomGenerator func(unsafe.Pointer) unsafe.Pointer
	_BNNSDirectApplyInTopK func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSDirectApplyReduction func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSDirectApplyTopK func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSFilterApplyBatch func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSFilterCreateLayerGram func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSFilterCreateLayerLoss func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSFilterCreateLayerReduction func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSFilterDestroy func(unsafe.Pointer) unsafe.Pointer
	_BNNSGather func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGatherND func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGetPointer func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphCompileOptionsGetGenerateDebugInfo func(unsafe.Pointer) unsafe.Pointer
	_BNNSGraphCompileOptionsGetOptimizationPreference func(unsafe.Pointer) unsafe.Pointer
	_BNNSGraphCompileOptionsSetMessageLogCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphCompileOptionsSetMessageLogMask func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphCompileOptionsSetOptimizationPreference func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphCompileOptionsSetOutputFD func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphCompileOptionsSetOutputPath func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphCompileOptionsSetTargetSingleThread func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphContextEnableNanAndInfChecks func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphContextExecute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphContextGetTensor func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphContextGetWorkspaceSize func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphContextMakeStreaming func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphContextSetStreamingAdvanceCount func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphContextSetWorkspaceAllocationCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphGetArgumentInterleaveFactors func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphGetArgumentNames func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphGetArgumentPosition func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphGetFunctionNames func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphGetOutputCount func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSGraphTensorFillStrides func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSLossFilterApplyBackwardBatch func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSLossFilterApplyBatch func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSNDArrayFullyConnectedSparsifySparseCOO func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSNDArrayFullyConnectedSparsifySparseCSR func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSNDArrayGetDataSize func(unsafe.Pointer) unsafe.Pointer
	_BNNSOptimizerStep func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSRandomGeneratorGetState func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSScatter func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSScatterND func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSShuffle func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSTensorGetAllocationSize func(unsafe.Pointer) unsafe.Pointer
	_BNNSTile func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSTileBackward func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_BNNSTranspose func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_catlas_sset func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
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
	_sparse_vector_add_with_scale_dense_double_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_vector_add_with_scale_dense_float_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_vector_triangular_solve_dense_double_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sparse_vector_triangular_solve_dense_float_complex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_DCT_Execute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_DFT_ExecuteD func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_DFT_Interleaved_CreateSetupD func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_DFT_Interleaved_DestroySetup func(unsafe.Pointer) unsafe.Pointer
	_vDSP_biquad func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_biquad_CreateSetup func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_biquad_CreateSetupD func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_biquad_DestroySetupD func(unsafe.Pointer) unsafe.Pointer
	_vDSP_biquad_SetCoefficientsDouble func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_biquad_SetCoefficientsSingle func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_ctoz func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_ctozD func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_fft2d_zip func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_fft_zrip func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_fftm_zrip func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_normalize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_normalizeD func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_sve func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_sve_svesq func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_sve_svesqD func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_svsD func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_vpoly func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_vsub func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_vsubD func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_vtmerg func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_ztoc func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vDSP_ztocD func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageAffineWarpD_ARGBFFFF func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageAlphaBlend_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageAlphaBlend_PlanarF func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageBuffer_InitWithCGImage func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, coregraphics.CGImageRef, unsafe.Pointer) unsafe.Pointer
	_vImageBuffer_InitWithCVPixelBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageCVImageFormat_CopyConversionMatrix func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageCVImageFormat_CreateWithCVPixelBuffer func(unsafe.Pointer) unsafe.Pointer
	_vImageCVImageFormat_GetConversionMatrix func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageContrastStretch_ARGBFFFF func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConvert_ARGB1555toRGB565 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConvert_ARGB8888To420Yp8_Cb8_Cr8 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConvert_ARGB8888To420Yp8_CbCr8 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConvert_ARGBToYpCbCr_GenerateConversion func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConvert_AnyToAny func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConvert_RGBA8888toRGB888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConvert_YpCbCrToARGB_GenerateConversion func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConverter_CreateForCGToCVImageFormat func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConverter_CreateForCVToCGImageFormat func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConverter_CreateWithCGColorConversionInfo func(coregraphics.CGColorConversionInfoRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConverter_CreateWithCGImageFormat func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConverter_CreateWithColorSyncCodeFragment func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageConverter_GetNumberOfSourceBuffers func(unsafe.Pointer) unsafe.Pointer
	_vImageConvolveWithBias_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageDestroyResamplingFilter func(unsafe.Pointer) unsafe.Pointer
	_vImageDilate_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageDilate_ARGBFFFF func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageEndsInContrastStretch_PlanarF func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageErode_PlanarF func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageHorizontalShear_ARGB16U func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageHorizontalShear_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageMatrixMultiply_ARGB8888ToPlanar8 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageMax_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageMax_ARGBFFFF func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageMin_ARGBFFFF func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageMultiDimensionalInterpolatedLookupTable_Planar16Q12 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageMultiDimensionalInterpolatedLookupTable_PlanarF func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageMultidimensionalTable_Create func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageMultidimensionalTable_Release func(unsafe.Pointer) unsafe.Pointer
	_vImageMultidimensionalTable_Retain func(unsafe.Pointer) unsafe.Pointer
	_vImagePremultipliedAlphaBlendLighten_RGBA8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImagePremultipliedAlphaBlend_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImagePremultipliedAlphaBlend_Planar8 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImagePremultipliedAlphaBlend_PlanarF func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImagePremultipliedConstAlphaBlend_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImagePremultiplyData_ARGB16U func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImagePremultiplyData_RGBA8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageScale_Planar16F func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageScale_Planar8 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageSepConvolve_Planar16U func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageSymmetricPiecewisePolynomial_PlanarF func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageTentConvolve_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageUnpremultiplyData_ARGB8888 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageVerticalShear_ARGB16F func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vImageVerticalShear_ARGBFFFF func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
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
	_vvacos func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvacosf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvacosh func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvacoshf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvasin func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvasinf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvasinh func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvasinhf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvatan func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvatan2 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvatan2f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvatanf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvatanh func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvatanhf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvceil func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvceilf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvcopysign func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvcopysignf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvcos func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvcosf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvcosh func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvcoshf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvcosisin func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvcosisinf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvcospi func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvcospif func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvdiv func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvdivf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvexp func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvexp2 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvexp2f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvexpf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvexpm1 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvexpm1f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvfabs func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvfabsf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvfloor func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvfloorf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvfmod func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvfmodf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvint func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvintf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvlog func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvlog10 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvlog10f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvlog1p func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvlog1pf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvlog2 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvlog2f func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvlogb func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvlogbf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvlogf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvnextafter func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvnextafterf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvnint func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvnintf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvpow func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvpowf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvrec func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvrecf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvremainder func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvremainderf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvrsqrt func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvrsqrtf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvsin func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvsincos func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvsincosf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvsinf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvsinh func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvsinhf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvsinpi func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvsinpif func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvsqrt func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvsqrtf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvtan func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvtanf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvtanh func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvtanhf func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvtanpi func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_vvtanpif func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_BNNSBandPart, lib, "BNNSBandPart")
	tryRegister(&_BNNSClipByGlobalNorm, lib, "BNNSClipByGlobalNorm")
	tryRegister(&_BNNSClipByNorm, lib, "BNNSClipByNorm")
	tryRegister(&_BNNSClipByValue, lib, "BNNSClipByValue")
	tryRegister(&_BNNSComputeNorm, lib, "BNNSComputeNorm")
	tryRegister(&_BNNSComputeNormBackward, lib, "BNNSComputeNormBackward")
	tryRegister(&_BNNSCopy, lib, "BNNSCopy")
	tryRegister(&_BNNSCreateRandomGenerator, lib, "BNNSCreateRandomGenerator")
	tryRegister(&_BNNSCreateRandomGeneratorWithSeed, lib, "BNNSCreateRandomGeneratorWithSeed")
	tryRegister(&_BNNSDataLayoutGetRank, lib, "BNNSDataLayoutGetRank")
	tryRegister(&_BNNSDestroyRandomGenerator, lib, "BNNSDestroyRandomGenerator")
	tryRegister(&_BNNSDirectApplyInTopK, lib, "BNNSDirectApplyInTopK")
	tryRegister(&_BNNSDirectApplyReduction, lib, "BNNSDirectApplyReduction")
	tryRegister(&_BNNSDirectApplyTopK, lib, "BNNSDirectApplyTopK")
	tryRegister(&_BNNSFilterApplyBatch, lib, "BNNSFilterApplyBatch")
	tryRegister(&_BNNSFilterCreateLayerGram, lib, "BNNSFilterCreateLayerGram")
	tryRegister(&_BNNSFilterCreateLayerLoss, lib, "BNNSFilterCreateLayerLoss")
	tryRegister(&_BNNSFilterCreateLayerReduction, lib, "BNNSFilterCreateLayerReduction")
	tryRegister(&_BNNSFilterDestroy, lib, "BNNSFilterDestroy")
	tryRegister(&_BNNSGather, lib, "BNNSGather")
	tryRegister(&_BNNSGatherND, lib, "BNNSGatherND")
	tryRegister(&_BNNSGetPointer, lib, "BNNSGetPointer")
	tryRegister(&_BNNSGraphCompileOptionsGetGenerateDebugInfo, lib, "BNNSGraphCompileOptionsGetGenerateDebugInfo")
	tryRegister(&_BNNSGraphCompileOptionsGetOptimizationPreference, lib, "BNNSGraphCompileOptionsGetOptimizationPreference")
	tryRegister(&_BNNSGraphCompileOptionsSetMessageLogCallback, lib, "BNNSGraphCompileOptionsSetMessageLogCallback")
	tryRegister(&_BNNSGraphCompileOptionsSetMessageLogMask, lib, "BNNSGraphCompileOptionsSetMessageLogMask")
	tryRegister(&_BNNSGraphCompileOptionsSetOptimizationPreference, lib, "BNNSGraphCompileOptionsSetOptimizationPreference")
	tryRegister(&_BNNSGraphCompileOptionsSetOutputFD, lib, "BNNSGraphCompileOptionsSetOutputFD")
	tryRegister(&_BNNSGraphCompileOptionsSetOutputPath, lib, "BNNSGraphCompileOptionsSetOutputPath")
	tryRegister(&_BNNSGraphCompileOptionsSetTargetSingleThread, lib, "BNNSGraphCompileOptionsSetTargetSingleThread")
	tryRegister(&_BNNSGraphContextEnableNanAndInfChecks, lib, "BNNSGraphContextEnableNanAndInfChecks")
	tryRegister(&_BNNSGraphContextExecute, lib, "BNNSGraphContextExecute")
	tryRegister(&_BNNSGraphContextGetTensor, lib, "BNNSGraphContextGetTensor")
	tryRegister(&_BNNSGraphContextGetWorkspaceSize, lib, "BNNSGraphContextGetWorkspaceSize")
	tryRegister(&_BNNSGraphContextMakeStreaming, lib, "BNNSGraphContextMakeStreaming")
	tryRegister(&_BNNSGraphContextSetStreamingAdvanceCount, lib, "BNNSGraphContextSetStreamingAdvanceCount")
	tryRegister(&_BNNSGraphContextSetWorkspaceAllocationCallback, lib, "BNNSGraphContextSetWorkspaceAllocationCallback")
	tryRegister(&_BNNSGraphGetArgumentInterleaveFactors, lib, "BNNSGraphGetArgumentInterleaveFactors")
	tryRegister(&_BNNSGraphGetArgumentNames, lib, "BNNSGraphGetArgumentNames")
	tryRegister(&_BNNSGraphGetArgumentPosition, lib, "BNNSGraphGetArgumentPosition")
	tryRegister(&_BNNSGraphGetFunctionNames, lib, "BNNSGraphGetFunctionNames")
	tryRegister(&_BNNSGraphGetOutputCount, lib, "BNNSGraphGetOutputCount")
	tryRegister(&_BNNSGraphTensorFillStrides, lib, "BNNSGraphTensorFillStrides")
	tryRegister(&_BNNSLossFilterApplyBackwardBatch, lib, "BNNSLossFilterApplyBackwardBatch")
	tryRegister(&_BNNSLossFilterApplyBatch, lib, "BNNSLossFilterApplyBatch")
	tryRegister(&_BNNSNDArrayFullyConnectedSparsifySparseCOO, lib, "BNNSNDArrayFullyConnectedSparsifySparseCOO")
	tryRegister(&_BNNSNDArrayFullyConnectedSparsifySparseCSR, lib, "BNNSNDArrayFullyConnectedSparsifySparseCSR")
	tryRegister(&_BNNSNDArrayGetDataSize, lib, "BNNSNDArrayGetDataSize")
	tryRegister(&_BNNSOptimizerStep, lib, "BNNSOptimizerStep")
	tryRegister(&_BNNSRandomGeneratorGetState, lib, "BNNSRandomGeneratorGetState")
	tryRegister(&_BNNSScatter, lib, "BNNSScatter")
	tryRegister(&_BNNSScatterND, lib, "BNNSScatterND")
	tryRegister(&_BNNSShuffle, lib, "BNNSShuffle")
	tryRegister(&_BNNSTensorGetAllocationSize, lib, "BNNSTensorGetAllocationSize")
	tryRegister(&_BNNSTile, lib, "BNNSTile")
	tryRegister(&_BNNSTileBackward, lib, "BNNSTileBackward")
	tryRegister(&_BNNSTranspose, lib, "BNNSTranspose")
	tryRegister(&_catlas_sset, lib, "catlas_sset")
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
	tryRegister(&_vDSP_DCT_Execute, lib, "vDSP_DCT_Execute")
	tryRegister(&_vDSP_DFT_ExecuteD, lib, "vDSP_DFT_ExecuteD")
	tryRegister(&_vDSP_DFT_Interleaved_CreateSetupD, lib, "vDSP_DFT_Interleaved_CreateSetupD")
	tryRegister(&_vDSP_DFT_Interleaved_DestroySetup, lib, "vDSP_DFT_Interleaved_DestroySetup")
	tryRegister(&_vDSP_biquad, lib, "vDSP_biquad")
	tryRegister(&_vDSP_biquad_CreateSetup, lib, "vDSP_biquad_CreateSetup")
	tryRegister(&_vDSP_biquad_CreateSetupD, lib, "vDSP_biquad_CreateSetupD")
	tryRegister(&_vDSP_biquad_DestroySetupD, lib, "vDSP_biquad_DestroySetupD")
	tryRegister(&_vDSP_biquad_SetCoefficientsDouble, lib, "vDSP_biquad_SetCoefficientsDouble")
	tryRegister(&_vDSP_biquad_SetCoefficientsSingle, lib, "vDSP_biquad_SetCoefficientsSingle")
	tryRegister(&_vDSP_ctoz, lib, "vDSP_ctoz")
	tryRegister(&_vDSP_ctozD, lib, "vDSP_ctozD")
	tryRegister(&_vDSP_fft2d_zip, lib, "vDSP_fft2d_zip")
	tryRegister(&_vDSP_fft_zrip, lib, "vDSP_fft_zrip")
	tryRegister(&_vDSP_fftm_zrip, lib, "vDSP_fftm_zrip")
	tryRegister(&_vDSP_normalize, lib, "vDSP_normalize")
	tryRegister(&_vDSP_normalizeD, lib, "vDSP_normalizeD")
	tryRegister(&_vDSP_sve, lib, "vDSP_sve")
	tryRegister(&_vDSP_sve_svesq, lib, "vDSP_sve_svesq")
	tryRegister(&_vDSP_sve_svesqD, lib, "vDSP_sve_svesqD")
	tryRegister(&_vDSP_svsD, lib, "vDSP_svsD")
	tryRegister(&_vDSP_vpoly, lib, "vDSP_vpoly")
	tryRegister(&_vDSP_vsub, lib, "vDSP_vsub")
	tryRegister(&_vDSP_vsubD, lib, "vDSP_vsubD")
	tryRegister(&_vDSP_vtmerg, lib, "vDSP_vtmerg")
	tryRegister(&_vDSP_ztoc, lib, "vDSP_ztoc")
	tryRegister(&_vDSP_ztocD, lib, "vDSP_ztocD")
	tryRegister(&_vImageAffineWarpD_ARGBFFFF, lib, "vImageAffineWarpD_ARGBFFFF")
	tryRegister(&_vImageAlphaBlend_ARGB8888, lib, "vImageAlphaBlend_ARGB8888")
	tryRegister(&_vImageAlphaBlend_PlanarF, lib, "vImageAlphaBlend_PlanarF")
	tryRegister(&_vImageBuffer_InitWithCGImage, lib, "vImageBuffer_InitWithCGImage")
	tryRegister(&_vImageBuffer_InitWithCVPixelBuffer, lib, "vImageBuffer_InitWithCVPixelBuffer")
	tryRegister(&_vImageCVImageFormat_CopyConversionMatrix, lib, "vImageCVImageFormat_CopyConversionMatrix")
	tryRegister(&_vImageCVImageFormat_CreateWithCVPixelBuffer, lib, "vImageCVImageFormat_CreateWithCVPixelBuffer")
	tryRegister(&_vImageCVImageFormat_GetConversionMatrix, lib, "vImageCVImageFormat_GetConversionMatrix")
	tryRegister(&_vImageContrastStretch_ARGBFFFF, lib, "vImageContrastStretch_ARGBFFFF")
	tryRegister(&_vImageConvert_ARGB1555toRGB565, lib, "vImageConvert_ARGB1555toRGB565")
	tryRegister(&_vImageConvert_ARGB8888To420Yp8_Cb8_Cr8, lib, "vImageConvert_ARGB8888To420Yp8_Cb8_Cr8")
	tryRegister(&_vImageConvert_ARGB8888To420Yp8_CbCr8, lib, "vImageConvert_ARGB8888To420Yp8_CbCr8")
	tryRegister(&_vImageConvert_ARGBToYpCbCr_GenerateConversion, lib, "vImageConvert_ARGBToYpCbCr_GenerateConversion")
	tryRegister(&_vImageConvert_AnyToAny, lib, "vImageConvert_AnyToAny")
	tryRegister(&_vImageConvert_RGBA8888toRGB888, lib, "vImageConvert_RGBA8888toRGB888")
	tryRegister(&_vImageConvert_YpCbCrToARGB_GenerateConversion, lib, "vImageConvert_YpCbCrToARGB_GenerateConversion")
	tryRegister(&_vImageConverter_CreateForCGToCVImageFormat, lib, "vImageConverter_CreateForCGToCVImageFormat")
	tryRegister(&_vImageConverter_CreateForCVToCGImageFormat, lib, "vImageConverter_CreateForCVToCGImageFormat")
	tryRegister(&_vImageConverter_CreateWithCGColorConversionInfo, lib, "vImageConverter_CreateWithCGColorConversionInfo")
	tryRegister(&_vImageConverter_CreateWithCGImageFormat, lib, "vImageConverter_CreateWithCGImageFormat")
	tryRegister(&_vImageConverter_CreateWithColorSyncCodeFragment, lib, "vImageConverter_CreateWithColorSyncCodeFragment")
	tryRegister(&_vImageConverter_GetNumberOfSourceBuffers, lib, "vImageConverter_GetNumberOfSourceBuffers")
	tryRegister(&_vImageConvolveWithBias_ARGB8888, lib, "vImageConvolveWithBias_ARGB8888")
	tryRegister(&_vImageDestroyResamplingFilter, lib, "vImageDestroyResamplingFilter")
	tryRegister(&_vImageDilate_ARGB8888, lib, "vImageDilate_ARGB8888")
	tryRegister(&_vImageDilate_ARGBFFFF, lib, "vImageDilate_ARGBFFFF")
	tryRegister(&_vImageEndsInContrastStretch_PlanarF, lib, "vImageEndsInContrastStretch_PlanarF")
	tryRegister(&_vImageErode_PlanarF, lib, "vImageErode_PlanarF")
	tryRegister(&_vImageHorizontalShear_ARGB16U, lib, "vImageHorizontalShear_ARGB16U")
	tryRegister(&_vImageHorizontalShear_ARGB8888, lib, "vImageHorizontalShear_ARGB8888")
	tryRegister(&_vImageMatrixMultiply_ARGB8888ToPlanar8, lib, "vImageMatrixMultiply_ARGB8888ToPlanar8")
	tryRegister(&_vImageMax_ARGB8888, lib, "vImageMax_ARGB8888")
	tryRegister(&_vImageMax_ARGBFFFF, lib, "vImageMax_ARGBFFFF")
	tryRegister(&_vImageMin_ARGBFFFF, lib, "vImageMin_ARGBFFFF")
	tryRegister(&_vImageMultiDimensionalInterpolatedLookupTable_Planar16Q12, lib, "vImageMultiDimensionalInterpolatedLookupTable_Planar16Q12")
	tryRegister(&_vImageMultiDimensionalInterpolatedLookupTable_PlanarF, lib, "vImageMultiDimensionalInterpolatedLookupTable_PlanarF")
	tryRegister(&_vImageMultidimensionalTable_Create, lib, "vImageMultidimensionalTable_Create")
	tryRegister(&_vImageMultidimensionalTable_Release, lib, "vImageMultidimensionalTable_Release")
	tryRegister(&_vImageMultidimensionalTable_Retain, lib, "vImageMultidimensionalTable_Retain")
	tryRegister(&_vImagePremultipliedAlphaBlendLighten_RGBA8888, lib, "vImagePremultipliedAlphaBlendLighten_RGBA8888")
	tryRegister(&_vImagePremultipliedAlphaBlend_ARGB8888, lib, "vImagePremultipliedAlphaBlend_ARGB8888")
	tryRegister(&_vImagePremultipliedAlphaBlend_Planar8, lib, "vImagePremultipliedAlphaBlend_Planar8")
	tryRegister(&_vImagePremultipliedAlphaBlend_PlanarF, lib, "vImagePremultipliedAlphaBlend_PlanarF")
	tryRegister(&_vImagePremultipliedConstAlphaBlend_ARGB8888, lib, "vImagePremultipliedConstAlphaBlend_ARGB8888")
	tryRegister(&_vImagePremultiplyData_ARGB16U, lib, "vImagePremultiplyData_ARGB16U")
	tryRegister(&_vImagePremultiplyData_RGBA8888, lib, "vImagePremultiplyData_RGBA8888")
	tryRegister(&_vImageScale_Planar16F, lib, "vImageScale_Planar16F")
	tryRegister(&_vImageScale_Planar8, lib, "vImageScale_Planar8")
	tryRegister(&_vImageSepConvolve_Planar16U, lib, "vImageSepConvolve_Planar16U")
	tryRegister(&_vImageSymmetricPiecewisePolynomial_PlanarF, lib, "vImageSymmetricPiecewisePolynomial_PlanarF")
	tryRegister(&_vImageTentConvolve_ARGB8888, lib, "vImageTentConvolve_ARGB8888")
	tryRegister(&_vImageUnpremultiplyData_ARGB8888, lib, "vImageUnpremultiplyData_ARGB8888")
	tryRegister(&_vImageVerticalShear_ARGB16F, lib, "vImageVerticalShear_ARGB16F")
	tryRegister(&_vImageVerticalShear_ARGBFFFF, lib, "vImageVerticalShear_ARGBFFFF")
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



// Copies the specified subdiagonals and superdiagonals of a matrix, and sets other elements to zero. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSBandPart(_:_:_:_:_:)
func BNNSBandPart(num_lower unsafe.Pointer, num_upper unsafe.Pointer, input unsafe.Pointer, output unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSBandPart(num_lower, num_upper, input, output, filter_params)
	}


// Clips a tensor’s values to a maximum global Euclidean norm. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSClipByGlobalNorm(_:_:_:_:_:)
func BNNSClipByGlobalNorm(dest unsafe.Pointer, src unsafe.Pointer, count unsafe.Pointer, max_norm unsafe.Pointer, use_norm unsafe.Pointer) unsafe.Pointer {
	return _BNNSClipByGlobalNorm(dest, src, count, max_norm, use_norm)
	}


// Clips a tensor’s values to a maximum Euclidean norm. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSClipByNorm(_:_:_:_:)
func BNNSClipByNorm(dest unsafe.Pointer, src unsafe.Pointer, max_norm unsafe.Pointer, axis_flags unsafe.Pointer) unsafe.Pointer {
	return _BNNSClipByNorm(dest, src, max_norm, axis_flags)
	}


// Clips a tensor’s values to the specified minimum and maximum values. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSClipByValue(_:_:_:_:)
func BNNSClipByValue(dest unsafe.Pointer, src unsafe.Pointer, min_val unsafe.Pointer, max_val unsafe.Pointer) unsafe.Pointer {
	return _BNNSClipByValue(dest, src, min_val, max_val)
	}


// Computes the specified norm over an entire tensor or the specified axes. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSComputeNorm(_:_:_:_:)
func BNNSComputeNorm(dest unsafe.Pointer, src unsafe.Pointer, norm_type unsafe.Pointer, axis_flags unsafe.Pointer) unsafe.Pointer {
	return _BNNSComputeNorm(dest, src, norm_type, axis_flags)
	}


// Backpropogates gradients for the compute norm function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSComputeNormBackward(_:_:_:_:_:_:)
func BNNSComputeNormBackward(in unsafe.Pointer, in_delta unsafe.Pointer, out unsafe.Pointer, out_delta unsafe.Pointer, norm_type unsafe.Pointer, axis_flags unsafe.Pointer) unsafe.Pointer {
	return _BNNSComputeNormBackward(in, in_delta, out, out_delta, norm_type, axis_flags)
	}


// Copies the contents of an n-dimensional array descriptor to another of the same shape. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSCopy(_:_:_:)
func BNNSCopy(dest unsafe.Pointer, src unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSCopy(dest, src, filter_params)
	}


// Returns a new random number generator using an internally generated random seed. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSCreateRandomGenerator(_:_:)
func BNNSCreateRandomGenerator(method unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSCreateRandomGenerator(method, filter_params)
	}


// Returns a new random number generator using the specified seed. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSCreateRandomGeneratorWithSeed(_:_:_:)
func BNNSCreateRandomGeneratorWithSeed(method unsafe.Pointer, seed unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSCreateRandomGeneratorWithSeed(method, seed, filter_params)
	}


// BNNSDataLayoutGetRank is a Accelerate function. [Full Topic]
//
// Added in macOS 14.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSDataLayoutGetRank(_:)
func BNNSDataLayoutGetRank(layout unsafe.Pointer) unsafe.Pointer {
	return _BNNSDataLayoutGetRank(layout)
	}


// Destroys a random number generator. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSDestroyRandomGenerator(_:)
func BNNSDestroyRandomGenerator(generator unsafe.Pointer) {
	_BNNSDestroyRandomGenerator(generator)
	}


// Applies an in-top-k filter directly to an input. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSDirectApplyInTopK(_:_:_:_:_:_:_:_:_:_:)
func BNNSDirectApplyInTopK(K unsafe.Pointer, axis unsafe.Pointer, batch_size unsafe.Pointer, input unsafe.Pointer, input_batch_stride unsafe.Pointer, test_indices unsafe.Pointer, test_indices_batch_stride unsafe.Pointer, output unsafe.Pointer, output_batch_stride unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSDirectApplyInTopK(K, axis, batch_size, input, input_batch_stride, test_indices, test_indices_batch_stride, output, output_batch_stride, filter_params)
	}


// Applies a reduction operation directly to an input tensor. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSDirectApplyReduction(_:_:)
func BNNSDirectApplyReduction(layer_params unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSDirectApplyReduction(layer_params, filter_params)
	}


// Applies a top-k filter directly to an input. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSDirectApplyTopK(_:_:_:_:_:_:_:_:_:_:)
func BNNSDirectApplyTopK(K unsafe.Pointer, axis unsafe.Pointer, batch_size unsafe.Pointer, input unsafe.Pointer, input_batch_stride unsafe.Pointer, best_values unsafe.Pointer, best_values_batch_stride unsafe.Pointer, best_indices unsafe.Pointer, best_indices_batch_stride unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSDirectApplyTopK(K, axis, batch_size, input, input_batch_stride, best_values, best_values_batch_stride, best_indices, best_indices_batch_stride, filter_params)
	}


// Applies a filter to a set of input objects, writing the result to a set of output objects. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFilterApplyBatch(_:_:_:_:_:_:)
func BNNSFilterApplyBatch(filter unsafe.Pointer, batch_size unsafe.Pointer, in unsafe.Pointer, in_stride unsafe.Pointer, out unsafe.Pointer, out_stride unsafe.Pointer) unsafe.Pointer {
	return _BNNSFilterApplyBatch(filter, batch_size, in, in_stride, out, out_stride)
	}


// Returns a new Gram matrix layer. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerGram(_:_:)
func BNNSFilterCreateLayerGram(layer_params unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSFilterCreateLayerGram(layer_params, filter_params)
	}


// Returns a new loss layer. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerLoss(_:_:)
func BNNSFilterCreateLayerLoss(layer_params unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSFilterCreateLayerLoss(layer_params, filter_params)
	}


// Returns a new reduction layer. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFilterCreateLayerReduction(_:_:)
func BNNSFilterCreateLayerReduction(layer_params unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSFilterCreateLayerReduction(layer_params, filter_params)
	}


// Destroys the specified filter, releasing all resources allocated for it. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSFilterDestroy(_:)
func BNNSFilterDestroy(filter unsafe.Pointer) {
	_BNNSFilterDestroy(filter)
	}


// Gathers the elements of a tensor along a single axis. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGather(_:_:_:_:_:)
func BNNSGather(axis unsafe.Pointer, input unsafe.Pointer, indices unsafe.Pointer, output unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSGather(axis, input, indices, output, filter_params)
	}


// Gathers the slices of a tensor. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGatherND(_:_:_:_:)
func BNNSGatherND(input unsafe.Pointer, indices unsafe.Pointer, output unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSGatherND(input, indices, output, filter_params)
	}


// Returns an n-dimensional array descriptor that contains a reference to a filter-data member. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGetPointer(_:_:)
func BNNSGetPointer(filter unsafe.Pointer, target unsafe.Pointer) unsafe.Pointer {
	return _BNNSGetPointer(filter, target)
	}


// Returns the option for the compiled graph to include debugging information. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsGetGenerateDebugInfo(_:)
func BNNSGraphCompileOptionsGetGenerateDebugInfo(options unsafe.Pointer) unsafe.Pointer {
	return _BNNSGraphCompileOptionsGetGenerateDebugInfo(options)
	}


// Returns the option for the compiled graph to optimize for either size or performance. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsGetOptimizationPreference(_:)
func BNNSGraphCompileOptionsGetOptimizationPreference(options unsafe.Pointer) unsafe.Pointer {
	return _BNNSGraphCompileOptionsGetOptimizationPreference(options)
	}


// Specifies a customized callback function that reports compile-time messages. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsSetMessageLogCallback(_:_:_:)
func BNNSGraphCompileOptionsSetMessageLogCallback(options unsafe.Pointer, log_callback unsafe.Pointer, additional_logging_arguments unsafe.Pointer) {
	_BNNSGraphCompileOptionsSetMessageLogCallback(options, log_callback, additional_logging_arguments)
	}


// Sets the mask for compile-time messages. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsSetMessageLogMask(_:_:)
func BNNSGraphCompileOptionsSetMessageLogMask(options unsafe.Pointer, log_level_mask unsafe.Pointer) {
	_BNNSGraphCompileOptionsSetMessageLogMask(options, log_level_mask)
	}


// Sets the option for the compiled graph to optimize for either size or performance. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsSetOptimizationPreference(_:_:)
func BNNSGraphCompileOptionsSetOptimizationPreference(options unsafe.Pointer, preference unsafe.Pointer) {
	_BNNSGraphCompileOptionsSetOptimizationPreference(options, preference)
	}


// Sets the option for graph compilation to generate the graph object directly to the specified file descriptor. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsSetOutputFD(_:_:)
func BNNSGraphCompileOptionsSetOutputFD(options unsafe.Pointer, fd unsafe.Pointer) {
	_BNNSGraphCompileOptionsSetOutputFD(options, fd)
	}


// Sets the option for graph compilation to generate the graph object directly to the specified file. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsSetOutputPath(_:_:)
func BNNSGraphCompileOptionsSetOutputPath(options unsafe.Pointer, path unsafe.Pointer) {
	_BNNSGraphCompileOptionsSetOutputPath(options, path)
	}


// Sets the option for the compiled graph to execute on a single thread. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphCompileOptionsSetTargetSingleThread(_:_:)
func BNNSGraphCompileOptionsSetTargetSingleThread(options unsafe.Pointer, value unsafe.Pointer) {
	_BNNSGraphCompileOptionsSetTargetSingleThread(options, value)
	}


// Specifies that the context checks intermediate tensors for NaNs and infinities. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextEnableNanAndInfChecks(_:_:)
func BNNSGraphContextEnableNanAndInfChecks(context unsafe.Pointer, enable_check_for_nans_inf unsafe.Pointer) {
	_BNNSGraphContextEnableNanAndInfChecks(context, enable_check_for_nans_inf)
	}


// Executes the specified function with the given context. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextExecute(_:_:_:_:_:_:)
func BNNSGraphContextExecute(context unsafe.Pointer, function unsafe.Pointer, argument_count unsafe.Pointer, arguments unsafe.Pointer, workspace_size unsafe.Pointer, workspace unsafe.Pointer) unsafe.Pointer {
	return _BNNSGraphContextExecute(context, function, argument_count, arguments, workspace_size, workspace)
	}


// Sets the properties of a tensor for the specified function argument. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextGetTensor(_:_:_:_:_:)
func BNNSGraphContextGetTensor(context unsafe.Pointer, function unsafe.Pointer, argument unsafe.Pointer, fill_known_dynamic_shapes unsafe.Pointer, tensor unsafe.Pointer) unsafe.Pointer {
	return _BNNSGraphContextGetTensor(context, function, argument, fill_known_dynamic_shapes, tensor)
	}


// Returns the minimum size, in bytes, of the workspace that graph context execution requires. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextGetWorkspaceSize(_:_:)
func BNNSGraphContextGetWorkspaceSize(context unsafe.Pointer, function unsafe.Pointer) unsafe.Pointer {
	return _BNNSGraphContextGetWorkspaceSize(context, function)
	}


// Returns an allocated and initialized graph context with streaming support from the specified graph. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextMakeStreaming(_:_:_:_:)
func BNNSGraphContextMakeStreaming(graph unsafe.Pointer, function unsafe.Pointer, initial_states_count unsafe.Pointer, initial_states unsafe.Pointer) unsafe.Pointer {
	return _BNNSGraphContextMakeStreaming(graph, function, initial_states_count, initial_states)
	}


// Sets the streaming advancement amount for cases with dynamically shaped inputs. [Full Topic]
//
// Added in macOS 15.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextSetStreamingAdvanceCount(_:_:)
func BNNSGraphContextSetStreamingAdvanceCount(context unsafe.Pointer, advance_count unsafe.Pointer) unsafe.Pointer {
	return _BNNSGraphContextSetStreamingAdvanceCount(context, advance_count)
	}


// Sets the allocation and deallocation callbacks for internal workspace. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphContextSetWorkspaceAllocationCallback(_:_:_:_:_:)
func BNNSGraphContextSetWorkspaceAllocationCallback(context unsafe.Pointer, realloc unsafe.Pointer, free unsafe.Pointer, user_memory_context_size unsafe.Pointer, user_memory_context unsafe.Pointer) unsafe.Pointer {
	return _BNNSGraphContextSetWorkspaceAllocationCallback(context, realloc, free, user_memory_context_size, user_memory_context)
	}


// Returns the interleave factors for arguments, if present [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetArgumentInterleaveFactors(_:_:_:_:_:)
func BNNSGraphGetArgumentInterleaveFactors(graph unsafe.Pointer, function unsafe.Pointer, argument_count unsafe.Pointer, argument_interleave unsafe.Pointer, argument_interleave_counts unsafe.Pointer) unsafe.Pointer {
	return _BNNSGraphGetArgumentInterleaveFactors(graph, function, argument_count, argument_interleave, argument_interleave_counts)
	}


// Extracts the names of arguments for the given function argument. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetArgumentNames(_:_:_:_:)
func BNNSGraphGetArgumentNames(graph unsafe.Pointer, function unsafe.Pointer, argument_names_count unsafe.Pointer, argument_names unsafe.Pointer) unsafe.Pointer {
	return _BNNSGraphGetArgumentNames(graph, function, argument_names_count, argument_names)
	}


// Returns the index into the arguments array for the given function argument. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetArgumentPosition(_:_:_:)
func BNNSGraphGetArgumentPosition(graph unsafe.Pointer, function unsafe.Pointer, argument unsafe.Pointer) unsafe.Pointer {
	return _BNNSGraphGetArgumentPosition(graph, function, argument)
	}


// Extracts the names of callable functions in the graph. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetFunctionNames(_:_:_:)
func BNNSGraphGetFunctionNames(graph unsafe.Pointer, function_name_count unsafe.Pointer, function_names unsafe.Pointer) unsafe.Pointer {
	return _BNNSGraphGetFunctionNames(graph, function_name_count, function_names)
	}


// Returns the number of output arguments for the given function argument. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphGetOutputCount(_:_:)
func BNNSGraphGetOutputCount(graph unsafe.Pointer, function unsafe.Pointer) unsafe.Pointer {
	return _BNNSGraphGetOutputCount(graph, function)
	}


// Sets the stride of the specifed tensor for compatibility with the given model’s input or output argument based on its current shape. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSGraphTensorFillStrides(_:_:_:_:)
func BNNSGraphTensorFillStrides(graph unsafe.Pointer, function unsafe.Pointer, argument unsafe.Pointer, tensor unsafe.Pointer) unsafe.Pointer {
	return _BNNSGraphTensorFillStrides(graph, function, argument, tensor)
	}


// Applies a loss filter backward to generate gradients. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLossFilterApplyBackwardBatch(_:_:_:_:_:_:_:_:_:_:_:_:)
func BNNSLossFilterApplyBackwardBatch(filter unsafe.Pointer, batch_size unsafe.Pointer, in unsafe.Pointer, in_stride unsafe.Pointer, in_delta unsafe.Pointer, in_delta_stride unsafe.Pointer, labels unsafe.Pointer, labels_stride unsafe.Pointer, weights unsafe.Pointer, weights_size unsafe.Pointer, out_delta unsafe.Pointer, out_delta_stride unsafe.Pointer) unsafe.Pointer {
	return _BNNSLossFilterApplyBackwardBatch(filter, batch_size, in, in_stride, in_delta, in_delta_stride, labels, labels_stride, weights, weights_size, out_delta, out_delta_stride)
	}


// Applies a loss filter to a set of input objects, writing the result to a set of output objects. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSLossFilterApplyBatch(_:_:_:_:_:_:_:_:_:_:_:)
func BNNSLossFilterApplyBatch(filter unsafe.Pointer, batch_size unsafe.Pointer, in unsafe.Pointer, in_stride unsafe.Pointer, labels unsafe.Pointer, labels_stride unsafe.Pointer, weights unsafe.Pointer, weights_size unsafe.Pointer, out unsafe.Pointer, in_delta unsafe.Pointer, in_delta_stride unsafe.Pointer) unsafe.Pointer {
	return _BNNSLossFilterApplyBatch(filter, batch_size, in, in_stride, labels, labels_stride, weights, weights_size, out, in_delta, in_delta_stride)
	}


// Converts a sparse tensor from the standardized coordinate list (COO) layout to a device-specific sparse layout that BNNS fully connected layers use. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSNDArrayFullyConnectedSparsifySparseCOO(_:_:_:_:_:_:_:_:_:)
func BNNSNDArrayFullyConnectedSparsifySparseCOO(in_dense_shape unsafe.Pointer, in_indices unsafe.Pointer, in_values unsafe.Pointer, out unsafe.Pointer, sparse_params unsafe.Pointer, batch_size unsafe.Pointer, workspace unsafe.Pointer, workspace_size unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSNDArrayFullyConnectedSparsifySparseCOO(in_dense_shape, in_indices, in_values, out, sparse_params, batch_size, workspace, workspace_size, filter_params)
	}


// Converts a sparse tensor from the standardized compressed sparse row (CSR) layout to a device-specific sparse layout that BNNS fully connected layers use. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSNDArrayFullyConnectedSparsifySparseCSR(_:_:_:_:_:_:_:_:_:_:)
func BNNSNDArrayFullyConnectedSparsifySparseCSR(in_dense_shape unsafe.Pointer, in_column_indices unsafe.Pointer, in_row_starts unsafe.Pointer, in_values unsafe.Pointer, out unsafe.Pointer, sparse_params unsafe.Pointer, batch_size unsafe.Pointer, workspace unsafe.Pointer, workspace_size unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSNDArrayFullyConnectedSparsifySparseCSR(in_dense_shape, in_column_indices, in_row_starts, in_values, out, sparse_params, batch_size, workspace, workspace_size, filter_params)
	}


// Returns the size, in bytes, that an array descriptor requires. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSNDArrayGetDataSize(_:)
func BNNSNDArrayGetDataSize(array unsafe.Pointer) unsafe.Pointer {
	return _BNNSNDArrayGetDataSize(array)
	}


// Applies a single optimization step to one or more parameters. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSOptimizerStep(_:_:_:_:_:_:_:)
func BNNSOptimizerStep(function unsafe.Pointer, OptimizerAlgFields unsafe.Pointer, number_of_parameters unsafe.Pointer, parameters unsafe.Pointer, gradients unsafe.Pointer, accumulators unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSOptimizerStep(function, OptimizerAlgFields, number_of_parameters, parameters, gradients, accumulators, filter_params)
	}


// Returns the state of a random number generator. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSRandomGeneratorGetState(_:_:_:)
func BNNSRandomGeneratorGetState(generator unsafe.Pointer, state_size unsafe.Pointer, state unsafe.Pointer) unsafe.Pointer {
	return _BNNSRandomGeneratorGetState(generator, state_size, state)
	}


// Scatters the elements of a tensor along a single axis. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSScatter(_:_:_:_:_:_:)
func BNNSScatter(axis unsafe.Pointer, op unsafe.Pointer, input unsafe.Pointer, indices unsafe.Pointer, output unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSScatter(axis, op, input, indices, output, filter_params)
	}


// Scatters the slices of a tensor. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSScatterND(_:_:_:_:_:)
func BNNSScatterND(op unsafe.Pointer, input unsafe.Pointer, indices unsafe.Pointer, output unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSScatterND(op, input, indices, output, filter_params)
	}


// Rearranges elements in a tensor according to shuffle type. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSShuffle(_:_:_:_:)
func BNNSShuffle(type_ unsafe.Pointer, input unsafe.Pointer, output unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSShuffle(type_, input, output, filter_params)
	}


// Returns the minimum allocation size, in bytes, of the specified tensor. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSTensorGetAllocationSize(_:)
func BNNSTensorGetAllocationSize(tensor unsafe.Pointer) unsafe.Pointer {
	return _BNNSTensorGetAllocationSize(tensor)
	}


// Generates an output tensor by tiling an input tensor multiple times. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSTile(_:_:_:)
func BNNSTile(input unsafe.Pointer, output unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSTile(input, output, filter_params)
	}


// Applies a tile filter backward to generate an input gradient. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSTileBackward(_:_:_:)
func BNNSTileBackward(in_delta unsafe.Pointer, out_delta unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSTileBackward(in_delta, out_delta, filter_params)
	}


// Transposes a tensor by swapping two of its dimensions. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/BNNSTranspose(_:_:_:_:_:)
func BNNSTranspose(dest unsafe.Pointer, src unsafe.Pointer, axis0 unsafe.Pointer, axis1 unsafe.Pointer, filter_params unsafe.Pointer) unsafe.Pointer {
	return _BNNSTranspose(dest, src, axis0, axis1, filter_params)
	}


// Modifies a vector (single-precision) in place, setting each element to a given value. [Full Topic]
//
// Added in macOS 13.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/catlas_sset(_:_:_:_:)
func catlas_sset(N unsafe.Pointer, ALPHA unsafe.Pointer, X unsafe.Pointer, INCX unsafe.Pointer) {
	_catlas_sset(N, ALPHA, X, INCX)
	}


// sparse_inner_product_dense_double_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_inner_product_dense_double_complex
func sparse_inner_product_dense_double_complex(nz unsafe.Pointer, x unsafe.Pointer, indx unsafe.Pointer, y unsafe.Pointer, incy unsafe.Pointer) unsafe.Pointer {
	return _sparse_inner_product_dense_double_complex(nz, x, indx, y, incy)
	}


// sparse_inner_product_dense_float_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_inner_product_dense_float_complex
func sparse_inner_product_dense_float_complex(nz unsafe.Pointer, x unsafe.Pointer, indx unsafe.Pointer, y unsafe.Pointer, incy unsafe.Pointer) unsafe.Pointer {
	return _sparse_inner_product_dense_float_complex(nz, x, indx, y, incy)
	}


// sparse_inner_product_sparse_double_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_inner_product_sparse_double_complex
func sparse_inner_product_sparse_double_complex(nzx unsafe.Pointer, nzy unsafe.Pointer, x unsafe.Pointer, indx unsafe.Pointer, y unsafe.Pointer, indy unsafe.Pointer) unsafe.Pointer {
	return _sparse_inner_product_sparse_double_complex(nzx, nzy, x, indx, y, indy)
	}


// sparse_inner_product_sparse_float_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_inner_product_sparse_float_complex
func sparse_inner_product_sparse_float_complex(nzx unsafe.Pointer, nzy unsafe.Pointer, x unsafe.Pointer, indx unsafe.Pointer, y unsafe.Pointer, indy unsafe.Pointer) unsafe.Pointer {
	return _sparse_inner_product_sparse_float_complex(nzx, nzy, x, indx, y, indy)
	}


// sparse_insert_entry_double_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_insert_entry_double_complex
func sparse_insert_entry_double_complex(A unsafe.Pointer, val unsafe.Pointer, i unsafe.Pointer, j unsafe.Pointer) unsafe.Pointer {
	return _sparse_insert_entry_double_complex(A, val, i, j)
	}


// sparse_insert_entry_float_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_insert_entry_float_complex
func sparse_insert_entry_float_complex(A unsafe.Pointer, val unsafe.Pointer, i unsafe.Pointer, j unsafe.Pointer) unsafe.Pointer {
	return _sparse_insert_entry_float_complex(A, val, i, j)
	}


// sparse_matrix_product_dense_double_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_product_dense_double_complex
func sparse_matrix_product_dense_double_complex(order unsafe.Pointer, transa unsafe.Pointer, n unsafe.Pointer, alpha unsafe.Pointer, A unsafe.Pointer, B unsafe.Pointer, ldb unsafe.Pointer, C unsafe.Pointer, ldc unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_product_dense_double_complex(order, transa, n, alpha, A, B, ldb, C, ldc)
	}


// sparse_matrix_product_dense_float_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_product_dense_float_complex
func sparse_matrix_product_dense_float_complex(order unsafe.Pointer, transa unsafe.Pointer, n unsafe.Pointer, alpha unsafe.Pointer, A unsafe.Pointer, B unsafe.Pointer, ldb unsafe.Pointer, C unsafe.Pointer, ldc unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_product_dense_float_complex(order, transa, n, alpha, A, B, ldb, C, ldc)
	}


// sparse_matrix_product_sparse_double_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_product_sparse_double_complex
func sparse_matrix_product_sparse_double_complex(order unsafe.Pointer, transa unsafe.Pointer, alpha unsafe.Pointer, A unsafe.Pointer, B unsafe.Pointer, C unsafe.Pointer, ldc unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_product_sparse_double_complex(order, transa, alpha, A, B, C, ldc)
	}


// sparse_matrix_product_sparse_float_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_product_sparse_float_complex
func sparse_matrix_product_sparse_float_complex(order unsafe.Pointer, transa unsafe.Pointer, alpha unsafe.Pointer, A unsafe.Pointer, B unsafe.Pointer, C unsafe.Pointer, ldc unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_product_sparse_float_complex(order, transa, alpha, A, B, C, ldc)
	}


// sparse_matrix_trace_double_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_trace_double_complex
func sparse_matrix_trace_double_complex(A unsafe.Pointer, offset unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_trace_double_complex(A, offset)
	}


// sparse_matrix_trace_float_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_trace_float_complex
func sparse_matrix_trace_float_complex(A unsafe.Pointer, offset unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_trace_float_complex(A, offset)
	}


// sparse_matrix_triangular_solve_dense_double_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_triangular_solve_dense_double_complex
func sparse_matrix_triangular_solve_dense_double_complex(order unsafe.Pointer, transt unsafe.Pointer, nrhs unsafe.Pointer, alpha unsafe.Pointer, T unsafe.Pointer, B unsafe.Pointer, ldb unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_triangular_solve_dense_double_complex(order, transt, nrhs, alpha, T, B, ldb)
	}


// sparse_matrix_triangular_solve_dense_float_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_triangular_solve_dense_float_complex
func sparse_matrix_triangular_solve_dense_float_complex(order unsafe.Pointer, transt unsafe.Pointer, nrhs unsafe.Pointer, alpha unsafe.Pointer, T unsafe.Pointer, B unsafe.Pointer, ldb unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_triangular_solve_dense_float_complex(order, transt, nrhs, alpha, T, B, ldb)
	}


// sparse_matrix_vector_product_dense_double_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_vector_product_dense_double_complex
func sparse_matrix_vector_product_dense_double_complex(transa unsafe.Pointer, alpha unsafe.Pointer, A unsafe.Pointer, x unsafe.Pointer, incx unsafe.Pointer, y unsafe.Pointer, incy unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_vector_product_dense_double_complex(transa, alpha, A, x, incx, y, incy)
	}


// sparse_matrix_vector_product_dense_float_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_matrix_vector_product_dense_float_complex
func sparse_matrix_vector_product_dense_float_complex(transa unsafe.Pointer, alpha unsafe.Pointer, A unsafe.Pointer, x unsafe.Pointer, incx unsafe.Pointer, y unsafe.Pointer, incy unsafe.Pointer) unsafe.Pointer {
	return _sparse_matrix_vector_product_dense_float_complex(transa, alpha, A, x, incx, y, incy)
	}


// sparse_outer_product_dense_double_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_outer_product_dense_double_complex
func sparse_outer_product_dense_double_complex(M unsafe.Pointer, N unsafe.Pointer, nz unsafe.Pointer, alpha unsafe.Pointer, x unsafe.Pointer, incx unsafe.Pointer, y unsafe.Pointer, indy unsafe.Pointer, C unsafe.Pointer) unsafe.Pointer {
	return _sparse_outer_product_dense_double_complex(M, N, nz, alpha, x, incx, y, indy, C)
	}


// sparse_outer_product_dense_float_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_outer_product_dense_float_complex
func sparse_outer_product_dense_float_complex(M unsafe.Pointer, N unsafe.Pointer, nz unsafe.Pointer, alpha unsafe.Pointer, x unsafe.Pointer, incx unsafe.Pointer, y unsafe.Pointer, indy unsafe.Pointer, C unsafe.Pointer) unsafe.Pointer {
	return _sparse_outer_product_dense_float_complex(M, N, nz, alpha, x, incx, y, indy, C)
	}


// sparse_vector_add_with_scale_dense_double_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_vector_add_with_scale_dense_double_complex
func sparse_vector_add_with_scale_dense_double_complex(nz unsafe.Pointer, alpha unsafe.Pointer, x unsafe.Pointer, indx unsafe.Pointer, y unsafe.Pointer, incy unsafe.Pointer) {
	_sparse_vector_add_with_scale_dense_double_complex(nz, alpha, x, indx, y, incy)
	}


// sparse_vector_add_with_scale_dense_float_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_vector_add_with_scale_dense_float_complex
func sparse_vector_add_with_scale_dense_float_complex(nz unsafe.Pointer, alpha unsafe.Pointer, x unsafe.Pointer, indx unsafe.Pointer, y unsafe.Pointer, incy unsafe.Pointer) {
	_sparse_vector_add_with_scale_dense_float_complex(nz, alpha, x, indx, y, incy)
	}


// sparse_vector_triangular_solve_dense_double_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_vector_triangular_solve_dense_double_complex
func sparse_vector_triangular_solve_dense_double_complex(transt unsafe.Pointer, alpha unsafe.Pointer, T unsafe.Pointer, x unsafe.Pointer, incx unsafe.Pointer) unsafe.Pointer {
	return _sparse_vector_triangular_solve_dense_double_complex(transt, alpha, T, x, incx)
	}


// sparse_vector_triangular_solve_dense_float_complex is a Accelerate function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/sparse_vector_triangular_solve_dense_float_complex
func sparse_vector_triangular_solve_dense_float_complex(transt unsafe.Pointer, alpha unsafe.Pointer, T unsafe.Pointer, x unsafe.Pointer, incx unsafe.Pointer) unsafe.Pointer {
	return _sparse_vector_triangular_solve_dense_float_complex(transt, alpha, T, x, incx)
	}


// Calculates the discrete cosine transform for a vector. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DCT_Execute
func vDSP_DCT_Execute(__Setup unsafe.Pointer, __Input unsafe.Pointer, __Output unsafe.Pointer) {
	_vDSP_DCT_Execute(__Setup, __Input, __Output)
	}


// Calculates the discrete double-precision Fourier transform for a vector. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_ExecuteD
func vDSP_DFT_ExecuteD(__Setup unsafe.Pointer, __Ir unsafe.Pointer, __Ii unsafe.Pointer, __Or unsafe.Pointer, __Oi unsafe.Pointer) {
	_vDSP_DFT_ExecuteD(__Setup, __Ir, __Ii, __Or, __Oi)
	}


// Returns a setup structure that contains precalculated data for forward and inverse, double-precision interleaved discrete Fourier transform (DFT) functions. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Interleaved_CreateSetupD(_:_:_:_:)
func vDSP_DFT_Interleaved_CreateSetupD(Previous unsafe.Pointer, Length unsafe.Pointer, Direction unsafe.Pointer, RealtoComplex unsafe.Pointer) unsafe.Pointer {
	return _vDSP_DFT_Interleaved_CreateSetupD(Previous, Length, Direction, RealtoComplex)
	}


// Releases a single-precision discrete Fourier transform (DFT) setup structure. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_DFT_Interleaved_DestroySetup(_:)
func vDSP_DFT_Interleaved_DestroySetup(Setup unsafe.Pointer) {
	_vDSP_DFT_Interleaved_DestroySetup(Setup)
	}


// Applies a single-precision single-channel biquadratic IIR filter. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquad
func vDSP_biquad(__Setup unsafe.Pointer, __Delay unsafe.Pointer, __X unsafe.Pointer, __IX unsafe.Pointer, __Y unsafe.Pointer, __IY unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_biquad(__Setup, __Delay, __X, __IX, __Y, __IY, __N)
	}


// Builds a data structure that contains precalculated data for use by a single-precision cascaded biquadratic filter function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquad_CreateSetup
func vDSP_biquad_CreateSetup(__Coefficients unsafe.Pointer, __M unsafe.Pointer) unsafe.Pointer {
	return _vDSP_biquad_CreateSetup(__Coefficients, __M)
	}


// Builds a data structure that contains precalculated data for use by a double-precision cascaded biquadratic filter function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquad_CreateSetupD
func vDSP_biquad_CreateSetupD(__Coefficients unsafe.Pointer, __M unsafe.Pointer) unsafe.Pointer {
	return _vDSP_biquad_CreateSetupD(__Coefficients, __M)
	}


// Destroys a double-precision biquadratic filter setup object. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquad_DestroySetupD
func vDSP_biquad_DestroySetupD(__setup unsafe.Pointer) {
	_vDSP_biquad_DestroySetupD(__setup)
	}


// Sets double-precision coefficients of the specified single-channel biquadratic filter setup object. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquad_SetCoefficientsDouble
func vDSP_biquad_SetCoefficientsDouble(__setup unsafe.Pointer, __coeffs unsafe.Pointer, __start_sec unsafe.Pointer, __nsec unsafe.Pointer) {
	_vDSP_biquad_SetCoefficientsDouble(__setup, __coeffs, __start_sec, __nsec)
	}


// Sets single-precision coefficients of the specified single-channel biquadratic filter setup object. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_biquad_SetCoefficientsSingle
func vDSP_biquad_SetCoefficientsSingle(__setup unsafe.Pointer, __coeffs unsafe.Pointer, __start_sec unsafe.Pointer, __nsec unsafe.Pointer) {
	_vDSP_biquad_SetCoefficientsSingle(__setup, __coeffs, __start_sec, __nsec)
	}


// Copies the contents of an interleaved single-precision complex vector to a split complex vector. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_ctoz
func vDSP_ctoz(__C unsafe.Pointer, __IC unsafe.Pointer, __Z unsafe.Pointer, __IZ unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_ctoz(__C, __IC, __Z, __IZ, __N)
	}


// Copies the contents of an interleaved double-precision complex vector to a split complex vector. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_ctozD
func vDSP_ctozD(__C unsafe.Pointer, __IC unsafe.Pointer, __Z unsafe.Pointer, __IZ unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_ctozD(__C, __IC, __Z, __IZ, __N)
	}


// Computes a 2D forward or inverse in-place, single-precision complex FFT. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_fft2d_zip
func vDSP_fft2d_zip(__Setup unsafe.Pointer, __C unsafe.Pointer, __IC0 unsafe.Pointer, __IC1 unsafe.Pointer, __Log2N0 unsafe.Pointer, __Log2N1 unsafe.Pointer, __Direction unsafe.Pointer) {
	_vDSP_fft2d_zip(__Setup, __C, __IC0, __IC1, __Log2N0, __Log2N1, __Direction)
	}


// Computes a forward or inverse in-place, single-precision real FFT. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_fft_zrip
func vDSP_fft_zrip(__Setup unsafe.Pointer, __C unsafe.Pointer, __IC unsafe.Pointer, __Log2N unsafe.Pointer, __Direction unsafe.Pointer) {
	_vDSP_fft_zrip(__Setup, __C, __IC, __Log2N, __Direction)
	}


// Computes a forward or inverse in-place, single-precision real FFT on multiple signals. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_fftm_zrip
func vDSP_fftm_zrip(__Setup unsafe.Pointer, __C unsafe.Pointer, __IC unsafe.Pointer, __IM unsafe.Pointer, __Log2N unsafe.Pointer, __M unsafe.Pointer, __Direction unsafe.Pointer) {
	_vDSP_fftm_zrip(__Setup, __C, __IC, __IM, __Log2N, __M, __Direction)
	}


// Computes single-precision mean and standard deviation, and then calculates new elements to have a zero mean and a unit standard deviation. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_normalize
func vDSP_normalize(__A unsafe.Pointer, __IA unsafe.Pointer, __C unsafe.Pointer, __IC unsafe.Pointer, __Mean unsafe.Pointer, __StandardDeviation unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_normalize(__A, __IA, __C, __IC, __Mean, __StandardDeviation, __N)
	}


// Computes double-precision mean and standard deviation, and then calculates new elements to have a zero mean and a unit standard deviation. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_normalizeD
func vDSP_normalizeD(__A unsafe.Pointer, __IA unsafe.Pointer, __C unsafe.Pointer, __IC unsafe.Pointer, __Mean unsafe.Pointer, __StandardDeviation unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_normalizeD(__A, __IA, __C, __IC, __Mean, __StandardDeviation, __N)
	}


// Calculates the sum of values in a single-precision vector. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_sve
func vDSP_sve(__A unsafe.Pointer, __I unsafe.Pointer, __C unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_sve(__A, __I, __C, __N)
	}


// Calculates the sum of values and the sum of squares in a single-precision vector. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_sve_svesq
func vDSP_sve_svesq(__A unsafe.Pointer, __IA unsafe.Pointer, __Sum unsafe.Pointer, __SumOfSquares unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_sve_svesq(__A, __IA, __Sum, __SumOfSquares, __N)
	}


// Calculates the sum of values and the sum of squares in a double-precision vector. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_sve_svesqD
func vDSP_sve_svesqD(__A unsafe.Pointer, __IA unsafe.Pointer, __Sum unsafe.Pointer, __SumOfSquares unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_sve_svesqD(__A, __IA, __Sum, __SumOfSquares, __N)
	}


// Calculates the sum of signed squares in a double-precision vector. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_svsD
func vDSP_svsD(__A unsafe.Pointer, __IA unsafe.Pointer, __C unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_svsD(__A, __IA, __C, __N)
	}


// Evaluates a single-precision polynomial using specified coefficients, variables, and strides. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_vpoly
func vDSP_vpoly(__A unsafe.Pointer, __IA unsafe.Pointer, __B unsafe.Pointer, __IB unsafe.Pointer, __C unsafe.Pointer, __IC unsafe.Pointer, __N unsafe.Pointer, __P unsafe.Pointer) {
	_vDSP_vpoly(__A, __IA, __B, __IB, __C, __IC, __N, __P)
	}


// Calculates the single-precision element-wise subtraction of two vectors, using the specified stride. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_vsub
func vDSP_vsub(__B unsafe.Pointer, __IB unsafe.Pointer, __A unsafe.Pointer, __IA unsafe.Pointer, __C unsafe.Pointer, __IC unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_vsub(__B, __IB, __A, __IA, __C, __IC, __N)
	}


// Calculates the double-precision element-wise subtraction of two vectors, using the specified stride. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_vsubD
func vDSP_vsubD(__B unsafe.Pointer, __IB unsafe.Pointer, __A unsafe.Pointer, __IA unsafe.Pointer, __C unsafe.Pointer, __IC unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_vsubD(__B, __IB, __A, __IA, __C, __IC, __N)
	}


// Performs a tapered merge between two single-precision vectors. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_vtmerg
func vDSP_vtmerg(__A unsafe.Pointer, __IA unsafe.Pointer, __B unsafe.Pointer, __IB unsafe.Pointer, __C unsafe.Pointer, __IC unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_vtmerg(__A, __IA, __B, __IB, __C, __IC, __N)
	}


// Copies the contents of a split single-precision complex vector to an interleaved vector. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_ztoc
func vDSP_ztoc(__Z unsafe.Pointer, __IZ unsafe.Pointer, __C unsafe.Pointer, __IC unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_ztoc(__Z, __IZ, __C, __IC, __N)
	}


// Copies the contents of a split double-precision complex vector to an interleaved vector. [Full Topic]
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vDSP_ztocD
func vDSP_ztocD(__Z unsafe.Pointer, __IZ unsafe.Pointer, __C unsafe.Pointer, __IC unsafe.Pointer, __N unsafe.Pointer) {
	_vDSP_ztocD(__Z, __IZ, __C, __IC, __N)
	}


// Applies a double-precision affine transformation to a 32-bit-per-channel, 4-channel interleaved image. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageAffineWarpD_ARGBFFFF(_:_:_:_:_:_:)
func vImageAffineWarpD_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, transform unsafe.Pointer, backColor unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageAffineWarpD_ARGBFFFF(src, dest, tempBuffer, transform, backColor, flags)
	}


// Performs nonpremultiplied alpha compositing of two 8-bit-per-channel, 4-channel ARGB buffers. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageAlphaBlend_ARGB8888(_:_:_:_:)
func vImageAlphaBlend_ARGB8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageAlphaBlend_ARGB8888(srcTop, srcBottom, dest, flags)
	}


// Performs nonpremultiplied alpha compositing of two 8-bit planar buffers. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageAlphaBlend_PlanarF(_:_:_:_:_:_:_:)
func vImageAlphaBlend_PlanarF(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, srcBottomAlpha unsafe.Pointer, alpha unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageAlphaBlend_PlanarF(srcTop, srcTopAlpha, srcBottom, srcBottomAlpha, alpha, dest, flags)
	}


// Initializes a vImage buffer with the contents of a Core Graphics image. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageBuffer_InitWithCGImage(_:_:_:_:_:)
func vImageBuffer_InitWithCGImage(buf unsafe.Pointer, format unsafe.Pointer, backgroundColor unsafe.Pointer, image coregraphics.CGImageRef, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageBuffer_InitWithCGImage(buf, format, backgroundColor, image, flags)
	}


// Initializes a vImage buffer with a copy of the contents of a Core Video pixel buffer. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageBuffer_InitWithCVPixelBuffer(_:_:_:_:_:_:)
func vImageBuffer_InitWithCVPixelBuffer(buffer unsafe.Pointer, desiredFormat unsafe.Pointer, cvPixelBuffer unsafe.Pointer, cvImageFormat unsafe.Pointer, backgroundColor unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageBuffer_InitWithCVPixelBuffer(buffer, desiredFormat, cvPixelBuffer, cvImageFormat, backgroundColor, flags)
	}


// Copies an RGB-to-YpCbCr conversion matrix to an image format’s internal matrix. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_CopyConversionMatrix(_:_:_:)
func vImageCVImageFormat_CopyConversionMatrix(format unsafe.Pointer, matrix unsafe.Pointer, inType unsafe.Pointer) unsafe.Pointer {
	return _vImageCVImageFormat_CopyConversionMatrix(format, matrix, inType)
	}


// Creates the description of the image encoding in an existing Core Video pixel buffer. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_CreateWithCVPixelBuffer(_:)
func vImageCVImageFormat_CreateWithCVPixelBuffer(buffer unsafe.Pointer) unsafe.Pointer {
	return _vImageCVImageFormat_CreateWithCVPixelBuffer(buffer)
	}


// Returns a pointer to the RGB-to-YpCbCr conversion matrix of a Core Video image format. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageCVImageFormat_GetConversionMatrix(_:_:)
func vImageCVImageFormat_GetConversionMatrix(format unsafe.Pointer, outType unsafe.Pointer) unsafe.Pointer {
	return _vImageCVImageFormat_GetConversionMatrix(format, outType)
	}


// Performs contrast stretching on a 32-bit-per-channel, 4-channel interleaved buffer. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageContrastStretch_ARGBFFFF(_:_:_:_:_:_:_:)
func vImageContrastStretch_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, histogram_entries unsafe.Pointer, minVal unsafe.Pointer, maxVal unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageContrastStretch_ARGBFFFF(src, dest, tempBuffer, histogram_entries, minVal, maxVal, flags)
	}


// Removes the alpha channel from an ARGB1555 buffer to produce an RGB565 result. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB1555toRGB565(_:_:_:)
func vImageConvert_ARGB1555toRGB565(src unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageConvert_ARGB1555toRGB565(src, dest, flags)
	}


// Converts an 8-bit-per-channel, 4-channel ARGB buffer to planar Yp, Cb, and Cr buffers. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888To420Yp8_Cb8_Cr8(_:_:_:_:_:_:_:)
func vImageConvert_ARGB8888To420Yp8_Cb8_Cr8(src unsafe.Pointer, destYp unsafe.Pointer, destCb unsafe.Pointer, destCr unsafe.Pointer, info unsafe.Pointer, permuteMap unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageConvert_ARGB8888To420Yp8_Cb8_Cr8(src, destYp, destCb, destCr, info, permuteMap, flags)
	}


// Converts an 8-bit-per-channel, 4-channel ARGB buffer to a planar Yp buffer and a 2-channel CbCr buffer. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGB8888To420Yp8_CbCr8(_:_:_:_:_:_:)
func vImageConvert_ARGB8888To420Yp8_CbCr8(src unsafe.Pointer, destYp unsafe.Pointer, destCbCr unsafe.Pointer, info unsafe.Pointer, permuteMap unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageConvert_ARGB8888To420Yp8_CbCr8(src, destYp, destCbCr, info, permuteMap, flags)
	}


// Generates the information that describes the conversion from ARGB to YpCbCr. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConvert_ARGBToYpCbCr_GenerateConversion(_:_:_:_:_:_:)
func vImageConvert_ARGBToYpCbCr_GenerateConversion(matrix unsafe.Pointer, pixelRange unsafe.Pointer, outInfo unsafe.Pointer, inARGBType unsafe.Pointer, outYpCbCrType unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageConvert_ARGBToYpCbCr_GenerateConversion(matrix, pixelRange, outInfo, inARGBType, outYpCbCrType, flags)
	}


// Converts the pixels in a vImage buffer to another format, using the specified converter. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConvert_AnyToAny(_:_:_:_:_:)
func vImageConvert_AnyToAny(converter unsafe.Pointer, srcs unsafe.Pointer, dests unsafe.Pointer, tempBuffer unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageConvert_AnyToAny(converter, srcs, dests, tempBuffer, flags)
	}


// Removes the alpha channel from an 8-bit-per-channel RGBA buffer to produce an 8-bit-per-channel RGB result. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConvert_RGBA8888toRGB888(_:_:_:)
func vImageConvert_RGBA8888toRGB888(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _vImageConvert_RGBA8888toRGB888(p0, p1, p2)
	}


// Generates the information that describes the conversion from YpCbCr to ARGB. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConvert_YpCbCrToARGB_GenerateConversion(_:_:_:_:_:_:)
func vImageConvert_YpCbCrToARGB_GenerateConversion(matrix unsafe.Pointer, pixelRange unsafe.Pointer, outInfo unsafe.Pointer, inYpCbCrType unsafe.Pointer, outARGBType unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageConvert_YpCbCrToARGB_GenerateConversion(matrix, pixelRange, outInfo, inYpCbCrType, outARGBType, flags)
	}


// Creates a vImage converter that converts a Core Graphics-formatted image to a Core Video-formatted image. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConverter_CreateForCGToCVImageFormat(_:_:_:_:_:)
func vImageConverter_CreateForCGToCVImageFormat(srcFormat unsafe.Pointer, destFormat unsafe.Pointer, backgroundColor unsafe.Pointer, flags unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _vImageConverter_CreateForCGToCVImageFormat(srcFormat, destFormat, backgroundColor, flags, error_)
	}


// Creates a vImage converter that converts a Core Video-formatted image to a Core Graphics-formatted image. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConverter_CreateForCVToCGImageFormat(_:_:_:_:_:)
func vImageConverter_CreateForCVToCGImageFormat(srcFormat unsafe.Pointer, destFormat unsafe.Pointer, backgroundColor unsafe.Pointer, flags unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _vImageConverter_CreateForCVToCGImageFormat(srcFormat, destFormat, backgroundColor, flags, error_)
	}


// Creates an any-to-any converter that uses a color conversion information object to convert from one image format to another. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConverter_CreateWithCGColorConversionInfo(_:_:_:_:_:_:)
func vImageConverter_CreateWithCGColorConversionInfo(colorConversionInfoRef coregraphics.CGColorConversionInfoRef, sFormat unsafe.Pointer, dFormat unsafe.Pointer, bg unsafe.Pointer, flags unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _vImageConverter_CreateWithCGColorConversionInfo(colorConversionInfoRef, sFormat, dFormat, bg, flags, error_)
	}


// Creates a vImage converter that converts from one vImage Core Graphics image format to another. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConverter_CreateWithCGImageFormat(_:_:_:_:_:)
func vImageConverter_CreateWithCGImageFormat(srcFormat unsafe.Pointer, destFormat unsafe.Pointer, backgroundColor unsafe.Pointer, flags unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _vImageConverter_CreateWithCGImageFormat(srcFormat, destFormat, backgroundColor, flags, error_)
	}


// Creates a vImage converter to convert from one vImage Core Graphics image format to another, using custom ColorSync transform. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConverter_CreateWithColorSyncCodeFragment(_:_:_:_:_:_:)
func vImageConverter_CreateWithColorSyncCodeFragment(codeFragment unsafe.Pointer, srcFormat unsafe.Pointer, destFormat unsafe.Pointer, backgroundColor unsafe.Pointer, flags unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _vImageConverter_CreateWithColorSyncCodeFragment(codeFragment, srcFormat, destFormat, backgroundColor, flags, error_)
	}


// Returns the number of source buffers consumed by the converter. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConverter_GetNumberOfSourceBuffers(_:)
func vImageConverter_GetNumberOfSourceBuffers(converter unsafe.Pointer) unsafe.Pointer {
	return _vImageConverter_GetNumberOfSourceBuffers(converter)
	}


// Convolves an 8-bit-per-channel, 4-channel interleaved image by a 2D kernel and adds a bias. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageConvolveWithBias_ARGB8888(_:_:_:_:_:_:_:_:_:_:_:_:)
func vImageConvolveWithBias_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, kernel unsafe.Pointer, kernel_height unsafe.Pointer, kernel_width unsafe.Pointer, divisor unsafe.Pointer, bias unsafe.Pointer, backgroundColor unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageConvolveWithBias_ARGB8888(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel_height, kernel_width, divisor, bias, backgroundColor, flags)
	}


// Disposes of a resampling filter object. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageDestroyResamplingFilter(_:)
func vImageDestroyResamplingFilter(filter unsafe.Pointer) {
	_vImageDestroyResamplingFilter(filter)
	}


// Dilates an 8-bit-per-channel, 4-channel interleaved buffer. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageDilate_ARGB8888(_:_:_:_:_:_:_:_:)
func vImageDilate_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, kernel unsafe.Pointer, kernel_height unsafe.Pointer, kernel_width unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageDilate_ARGB8888(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel_height, kernel_width, flags)
	}


// Dilates a 32-bit-per-channel, 4-channel interleaved buffer. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageDilate_ARGBFFFF(_:_:_:_:_:_:_:_:)
func vImageDilate_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, kernel unsafe.Pointer, kernel_height unsafe.Pointer, kernel_width unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageDilate_ARGBFFFF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel_height, kernel_width, flags)
	}


// Performs ends-in contrast stretching on a 32-bit planar buffer. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageEndsInContrastStretch_PlanarF(_:_:_:_:_:_:_:_:_:)
func vImageEndsInContrastStretch_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, percent_low unsafe.Pointer, percent_high unsafe.Pointer, histogram_entries unsafe.Pointer, minVal unsafe.Pointer, maxVal unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageEndsInContrastStretch_PlanarF(src, dest, tempBuffer, percent_low, percent_high, histogram_entries, minVal, maxVal, flags)
	}


// Erodes a 32-bit planar buffer. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageErode_PlanarF(_:_:_:_:_:_:_:_:)
func vImageErode_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, kernel unsafe.Pointer, kernel_height unsafe.Pointer, kernel_width unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageErode_PlanarF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, kernel, kernel_height, kernel_width, flags)
	}


// Performs a single-precision horizontal shear on a region of interest within an unsigned 16-bit-per-channel, 4-channel interleaved image. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_ARGB16U(_:_:_:_:_:_:_:_:_:)
func vImageHorizontalShear_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, xTranslate unsafe.Pointer, shearSlope unsafe.Pointer, filter unsafe.Pointer, backColor unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageHorizontalShear_ARGB16U(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
	}


// Performs a single-precision horizontal shear on a region of interest within an 8-bit-per-channel, 4-channel interleaved image. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageHorizontalShear_ARGB8888(_:_:_:_:_:_:_:_:_:)
func vImageHorizontalShear_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, xTranslate unsafe.Pointer, shearSlope unsafe.Pointer, filter unsafe.Pointer, backColor unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageHorizontalShear_ARGB8888(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, xTranslate, shearSlope, filter, backColor, flags)
	}


// Multiplies each pixel in an interleaved four-channel, 8-bit source image by a matrix to produce a planar 8-bit destination image. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMatrixMultiply_ARGB8888ToPlanar8(_:_:_:_:_:_:_:)
func vImageMatrixMultiply_ARGB8888ToPlanar8(src unsafe.Pointer, dest unsafe.Pointer, matrix unsafe.Pointer, divisor unsafe.Pointer, pre_bias unsafe.Pointer, post_bias unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageMatrixMultiply_ARGB8888ToPlanar8(src, dest, matrix, divisor, pre_bias, post_bias, flags)
	}


// Maximizes an 8-bit-per-channel, 4-channel interleaved buffer. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMax_ARGB8888(_:_:_:_:_:_:_:_:)
func vImageMax_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, kernel_height unsafe.Pointer, kernel_width unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageMax_ARGB8888(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, flags)
	}


// Maximizes a 32-bit-per-channel, 4-channel interleaved buffer. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMax_ARGBFFFF(_:_:_:_:_:_:_:_:)
func vImageMax_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, kernel_height unsafe.Pointer, kernel_width unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageMax_ARGBFFFF(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, flags)
	}


// Minimizes an 8-bit-per-channel, 4-channel interleaved buffer. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMin_ARGBFFFF(_:_:_:_:_:_:_:_:)
func vImageMin_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, kernel_height unsafe.Pointer, kernel_width unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageMin_ARGBFFFF(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, flags)
	}


// Uses a multidimensional lookup table to transform a 16Q12 planar image. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMultiDimensionalInterpolatedLookupTable_Planar16Q12(_:_:_:_:_:_:)
func vImageMultiDimensionalInterpolatedLookupTable_Planar16Q12(srcs unsafe.Pointer, dests unsafe.Pointer, tempBuffer unsafe.Pointer, table unsafe.Pointer, method unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageMultiDimensionalInterpolatedLookupTable_Planar16Q12(srcs, dests, tempBuffer, table, method, flags)
	}


// Uses a multidimensional lookup table to transform a 32-bit planar image. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMultiDimensionalInterpolatedLookupTable_PlanarF(_:_:_:_:_:_:)
func vImageMultiDimensionalInterpolatedLookupTable_PlanarF(srcs unsafe.Pointer, dests unsafe.Pointer, tempBuffer unsafe.Pointer, table unsafe.Pointer, method unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageMultiDimensionalInterpolatedLookupTable_PlanarF(srcs, dests, tempBuffer, table, method, flags)
	}


// Creates a multidimensional lookup table. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMultidimensionalTable_Create(_:_:_:_:_:_:_:)
func vImageMultidimensionalTable_Create(tableData unsafe.Pointer, numSrcChannels unsafe.Pointer, numDestChannels unsafe.Pointer, table_entries_per_dimension unsafe.Pointer, hint unsafe.Pointer, flags unsafe.Pointer, err unsafe.Pointer) unsafe.Pointer {
	return _vImageMultidimensionalTable_Create(tableData, numSrcChannels, numDestChannels, table_entries_per_dimension, hint, flags, err)
	}


// Releases a multidimensional table. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMultidimensionalTable_Release(_:)
func vImageMultidimensionalTable_Release(table unsafe.Pointer) unsafe.Pointer {
	return _vImageMultidimensionalTable_Release(table)
	}


// Retains a multidimensional table. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageMultidimensionalTable_Retain(_:)
func vImageMultidimensionalTable_Retain(table unsafe.Pointer) unsafe.Pointer {
	return _vImageMultidimensionalTable_Retain(table)
	}


// Performs alpha compositing of two 8-bit-per-channel, 4-channel BGRA buffers using the lighten blend mode. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlendLighten_RGBA8888(_:_:_:_:)
func vImagePremultipliedAlphaBlendLighten_RGBA8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImagePremultipliedAlphaBlendLighten_RGBA8888(srcTop, srcBottom, dest, flags)
	}


// Performs premultiplied alpha compositing of two 8-bit-per-channel, 4-channel ARGB buffers. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlend_ARGB8888(_:_:_:_:)
func vImagePremultipliedAlphaBlend_ARGB8888(srcTop unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImagePremultipliedAlphaBlend_ARGB8888(srcTop, srcBottom, dest, flags)
	}


// Performs premultiplied alpha compositing of two 8-bit planar buffers. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlend_Planar8(_:_:_:_:_:)
func vImagePremultipliedAlphaBlend_Planar8(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImagePremultipliedAlphaBlend_Planar8(srcTop, srcTopAlpha, srcBottom, dest, flags)
	}


// Performs premultiplied alpha compositing of two 32-bit planar buffers. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedAlphaBlend_PlanarF(_:_:_:_:_:)
func vImagePremultipliedAlphaBlend_PlanarF(srcTop unsafe.Pointer, srcTopAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImagePremultipliedAlphaBlend_PlanarF(srcTop, srcTopAlpha, srcBottom, dest, flags)
	}


// Performs premultiplied alpha compositing of two 8-bit-per-channel, 4-channel interleaved buffers and applies an extra alpha value to the top buffer. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImagePremultipliedConstAlphaBlend_ARGB8888(_:_:_:_:_:)
func vImagePremultipliedConstAlphaBlend_ARGB8888(srcTop unsafe.Pointer, constAlpha unsafe.Pointer, srcBottom unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImagePremultipliedConstAlphaBlend_ARGB8888(srcTop, constAlpha, srcBottom, dest, flags)
	}


// Transforms an unsigned 16-bit-per-channel, 4-channel ARGB buffer from nonpremultiplied alpha format to premultiplied alpha format. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImagePremultiplyData_ARGB16U(_:_:_:)
func vImagePremultiplyData_ARGB16U(src unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImagePremultiplyData_ARGB16U(src, dest, flags)
	}


// Transforms an 8-bit-per-channel, 4-channel RGBA buffer from nonpremultiplied alpha format to premultiplied alpha format. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImagePremultiplyData_RGBA8888(_:_:_:)
func vImagePremultiplyData_RGBA8888(src unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImagePremultiplyData_RGBA8888(src, dest, flags)
	}


// Scales a floating-point 16-bit planar image to fit a destination buffer. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageScale_Planar16F(_:_:_:_:)
func vImageScale_Planar16F(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageScale_Planar16F(src, dest, tempBuffer, flags)
	}


// Scales an 8-bit planar image to fit a destination buffer. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageScale_Planar8(_:_:_:_:)
func vImageScale_Planar8(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageScale_Planar8(src, dest, tempBuffer, flags)
	}


// Convolves an unsigned 16-bit planar image by separate horizontal and vertical separable kernels. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageSepConvolve_Planar16U(_:_:_:_:_:_:_:_:_:_:_:_:)
func vImageSepConvolve_Planar16U(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, kernelX unsafe.Pointer, kernelX_width unsafe.Pointer, kernelY unsafe.Pointer, kernelY_width unsafe.Pointer, bias unsafe.Pointer, backgroundColor unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageSepConvolve_Planar16U(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernelX, kernelX_width, kernelY, kernelY_width, bias, backgroundColor, flags)
	}


// Applies a set of symmetric piecewise polynomials to transform a 32-bit planar image. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageSymmetricPiecewisePolynomial_PlanarF(_:_:_:_:_:_:_:)
func vImageSymmetricPiecewisePolynomial_PlanarF(src unsafe.Pointer, dest unsafe.Pointer, coefficients unsafe.Pointer, boundaries unsafe.Pointer, order unsafe.Pointer, log2segments unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageSymmetricPiecewisePolynomial_PlanarF(src, dest, coefficients, boundaries, order, log2segments, flags)
	}


// Applies a tent filter to an 8-bit-per-channel, 4-channel interleaved source image. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageTentConvolve_ARGB8888(_:_:_:_:_:_:_:_:_:)
func vImageTentConvolve_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, tempBuffer unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, kernel_height unsafe.Pointer, kernel_width unsafe.Pointer, backgroundColor unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageTentConvolve_ARGB8888(src, dest, tempBuffer, srcOffsetToROI_X, srcOffsetToROI_Y, kernel_height, kernel_width, backgroundColor, flags)
	}


// Transforms an 8-bit-per-channel, 4-channel ARGB buffer from premultiplied alpha format to nonpremultiplied alpha format. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageUnpremultiplyData_ARGB8888(_:_:_:)
func vImageUnpremultiplyData_ARGB8888(src unsafe.Pointer, dest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageUnpremultiplyData_ARGB8888(src, dest, flags)
	}


// Performs a single-precision vertical shear on a region of interest within a floating-point 16-bit-per-channel, 4-channel interleaved image. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_ARGB16F(_:_:_:_:_:_:_:_:_:)
func vImageVerticalShear_ARGB16F(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, yTranslate unsafe.Pointer, shearSlope unsafe.Pointer, filter unsafe.Pointer, backColor unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageVerticalShear_ARGB16F(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
	}


// Performs a single-precision vertical shear on a region of interest within a 32-bit-per-channel, 4-channel interleaved image. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vImageVerticalShear_ARGBFFFF(_:_:_:_:_:_:_:_:_:)
func vImageVerticalShear_ARGBFFFF(src unsafe.Pointer, dest unsafe.Pointer, srcOffsetToROI_X unsafe.Pointer, srcOffsetToROI_Y unsafe.Pointer, yTranslate unsafe.Pointer, shearSlope unsafe.Pointer, filter unsafe.Pointer, backColor unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _vImageVerticalShear_ARGBFFFF(src, dest, srcOffsetToROI_X, srcOffsetToROI_Y, yTranslate, shearSlope, filter, backColor, flags)
	}


// 128-bit arithmetic (signed) shift. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vA128Shift(_:_:)
func vA128Shift(vA unsafe.Pointer, vShiftFactor unsafe.Pointer) unsafe.Pointer {
	return _vA128Shift(vA, vShiftFactor)
	}


// 128-bit logical left shift. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vLL128Shift(_:_:)
func vLL128Shift(vA unsafe.Pointer, vShiftFactor unsafe.Pointer) unsafe.Pointer {
	return _vLL128Shift(vA, vShiftFactor)
	}


// 128-bit logical right shift. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vLR128Shift(_:_:)
func vLR128Shift(vA unsafe.Pointer, vShiftFactor unsafe.Pointer) unsafe.Pointer {
	return _vLR128Shift(vA, vShiftFactor)
	}


// Signed 128-bit addition (modular arithmetic). [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vS128Add(_:_:)
func vS128Add(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vS128Add(vA, vB)
	}


// Signed 128-bit addition with saturation (clipping). [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vS128AddS(_:_:)
func vS128AddS(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vS128AddS(vA, vB)
	}


// Signed 128-bit subtraction (modular arithmetic). [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vS128Sub(_:_:)
func vS128Sub(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vS128Sub(vA, vB)
	}


// Signed 128-bit subtraction with saturation (clipping). [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vS128SubS(_:_:)
func vS128SubS(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vS128SubS(vA, vB)
	}


// Signed 64-bit multiplication; results are twice as wide as multiplicands, odd-numbered elements of multiplicand vectors are used. Note the big-endian convention: the leftmost element is element 0. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vS64FullMulOdd(_:_:)
func vS64FullMulOdd(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vS64FullMulOdd(vA, vB)
	}


// Signed 64-bit subtraction with saturation (clipping). [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vS64SubS(_:_:)
func vS64SubS(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vS64SubS(vA, vB)
	}


// Unsigned 128-bit addition (modular arithmetic). [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vU128Add(_:_:)
func vU128Add(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vU128Add(vA, vB)
	}


// Unsigned 128-bit addition with saturation (clipping). [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vU128AddS(_:_:)
func vU128AddS(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vU128AddS(vA, vB)
	}


// Unsigned 128-bit subtraction (modular arithmetic). [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vU128Sub(_:_:)
func vU128Sub(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vU128Sub(vA, vB)
	}


// Unsigned 128-bit subtraction with saturation (clipping). [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vU128SubS(_:_:)
func vU128SubS(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vU128SubS(vA, vB)
	}


// Unsigned 64-bit multiplication; results are twice as wide as multiplicands, odd-numbered elements of multiplicand vectors are used. Note the big-endian convention: the leftmost element is element 0. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vU64FullMulOdd(_:_:)
func vU64FullMulOdd(vA unsafe.Pointer, vB unsafe.Pointer) unsafe.Pointer {
	return _vU64FullMulOdd(vA, vB)
	}


// Calculates the arccosine of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvacos(_:_:_:)
func vvacos(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvacos(p0, p1, p2)
	}


// Calculates the arccosine of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvacosf(_:_:_:)
func vvacosf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvacosf(p0, p1, p2)
	}


// Calculates the inverse hyperbolic cosine of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvacosh(_:_:_:)
func vvacosh(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvacosh(p0, p1, p2)
	}


// Calculates the inverse hyperbolic cosine of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvacoshf(_:_:_:)
func vvacoshf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvacoshf(p0, p1, p2)
	}


// Calculates the arcsine of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvasin(_:_:_:)
func vvasin(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvasin(p0, p1, p2)
	}


// Calculates the arcsine of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvasinf(_:_:_:)
func vvasinf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvasinf(p0, p1, p2)
	}


// Calculates the inverse hyperbolic sine of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvasinh(_:_:_:)
func vvasinh(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvasinh(p0, p1, p2)
	}


// Calculates the inverse hyperbolic sine of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvasinhf(_:_:_:)
func vvasinhf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvasinhf(p0, p1, p2)
	}


// Calculates the arctangent of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvatan(_:_:_:)
func vvatan(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvatan(p0, p1, p2)
	}


// Calculates the arctangent of each pair of elements in two arrays of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvatan2(_:_:_:_:)
func vvatan2(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvatan2(p0, p1, p2, p3)
	}


// Calculates the arctangent of each pair of elements in two arrays of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvatan2f(_:_:_:_:)
func vvatan2f(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvatan2f(p0, p1, p2, p3)
	}


// Calculates the arctangent of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvatanf(_:_:_:)
func vvatanf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvatanf(p0, p1, p2)
	}


// Calculates the inverse hyperbolic tangent of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvatanh(_:_:_:)
func vvatanh(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvatanh(p0, p1, p2)
	}


// Calculates the inverse hyperbolic tangent of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvatanhf(_:_:_:)
func vvatanhf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvatanhf(p0, p1, p2)
	}


// Calculates the ceiling of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvceil(_:_:_:)
func vvceil(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvceil(p0, p1, p2)
	}


// Calculates the ceiling of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvceilf(_:_:_:)
func vvceilf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvceilf(p0, p1, p2)
	}


// Copies an array, setting the sign of each element based on a second array of double-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcopysign(_:_:_:_:)
func vvcopysign(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvcopysign(p0, p1, p2, p3)
	}


// Copies an array, setting the sign of each element based on a second array of single-precision values. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcopysignf(_:_:_:_:)
func vvcopysignf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvcopysignf(p0, p1, p2, p3)
	}


// Calculates the cosine of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcos(_:_:_:)
func vvcos(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvcos(p0, p1, p2)
	}


// Calculates the cosine of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcosf(_:_:_:)
func vvcosf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvcosf(p0, p1, p2)
	}


// Calculates the hyperbolic cosine of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcosh(_:_:_:)
func vvcosh(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvcosh(p0, p1, p2)
	}


// Calculates the hyperbolic cosine of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcoshf(_:_:_:)
func vvcoshf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvcoshf(p0, p1, p2)
	}


// Calculates the cosine and sine of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcosisin(_:_:_:)
func vvcosisin(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvcosisin(p0, p1, p2)
	}


// Calculates the cosine and sine of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcosisinf(_:_:_:)
func vvcosisinf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvcosisinf(p0, p1, p2)
	}


// Calculates the cosine of pi multiplied by each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcospi(_:_:_:)
func vvcospi(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvcospi(p0, p1, p2)
	}


// Calculates the cosine of pi multiplied by each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvcospif(_:_:_:)
func vvcospif(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvcospif(p0, p1, p2)
	}


// Divides each element in an array by the corresponding value in a second array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvdiv(_:_:_:_:)
func vvdiv(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvdiv(p0, p1, p2, p3)
	}


// Divides each element in an array by the corresponding value in a second array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvdivf(_:_:_:_:)
func vvdivf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvdivf(p0, p1, p2, p3)
	}


// Calculates raised to the power of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvexp(_:_:_:)
func vvexp(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvexp(p0, p1, p2)
	}


// Calculates 2 raised to the power of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvexp2(_:_:_:)
func vvexp2(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvexp2(p0, p1, p2)
	}


// Calculates 2 raised to the power of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvexp2f(_:_:_:)
func vvexp2f(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvexp2f(p0, p1, p2)
	}


// Calculates raised to the power of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvexpf(_:_:_:)
func vvexpf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvexpf(p0, p1, p2)
	}


// Calculates for each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvexpm1(_:_:_:)
func vvexpm1(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvexpm1(p0, p1, p2)
	}


// Calculates for each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvexpm1f(_:_:_:)
func vvexpm1f(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvexpm1f(p0, p1, p2)
	}


// Calculates the absolute value for each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvfabs(_:_:_:)
func vvfabs(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvfabs(p0, p1, p2)
	}


// Calculates the absolute value for each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvfabsf(_:_:_:)
func vvfabsf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvfabsf(p0, p1, p2)
	}


// Calculates the floor of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvfloor(_:_:_:)
func vvfloor(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvfloor(p0, p1, p2)
	}


// Calculates the floor of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvfloorf(_:_:_:)
func vvfloorf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvfloorf(p0, p1, p2)
	}


// Calculates the modulus after dividing each element in an array by the corresponding element in a second array of double-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvfmod(_:_:_:_:)
func vvfmod(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvfmod(p0, p1, p2, p3)
	}


// Calculates the modulus after dividing each element in an array by the corresponding element in a second array of single-precision values. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvfmodf(_:_:_:_:)
func vvfmodf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvfmodf(p0, p1, p2, p3)
	}


// Calculates the integer truncation for each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvint(_:_:_:)
func vvint(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvint(p0, p1, p2)
	}


// Calculates the integer truncation for each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvintf(_:_:_:)
func vvintf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvintf(p0, p1, p2)
	}


// Calculates the natural logarithm for each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlog(_:_:_:)
func vvlog(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvlog(p0, p1, p2)
	}


// Calculates the base 10 logarithm of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlog10(_:_:_:)
func vvlog10(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvlog10(p0, p1, p2)
	}


// Calculates the base 10 logarithm of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlog10f(_:_:_:)
func vvlog10f(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvlog10f(p0, p1, p2)
	}


// Calculates for each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlog1p(_:_:_:)
func vvlog1p(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvlog1p(p0, p1, p2)
	}


// Calculates for each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlog1pf(_:_:_:)
func vvlog1pf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvlog1pf(p0, p1, p2)
	}


// Calculates the base 2 logarithm of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlog2(_:_:_:)
func vvlog2(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvlog2(p0, p1, p2)
	}


// Calculates the base 2 logarithm of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlog2f(_:_:_:)
func vvlog2f(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvlog2f(p0, p1, p2)
	}


// Calculates the unbiased exponent of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlogb(_:_:_:)
func vvlogb(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvlogb(p0, p1, p2)
	}


// Calculates the unbiased exponent of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlogbf(_:_:_:)
func vvlogbf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvlogbf(p0, p1, p2)
	}


// Calculates the natural logarithm for each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvlogf(_:_:_:)
func vvlogf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvlogf(p0, p1, p2)
	}


// Calculates the next machine-representable value for each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvnextafter(_:_:_:_:)
func vvnextafter(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvnextafter(p0, p1, p2, p3)
	}


// Calculates the next machine-representable value for each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvnextafterf(_:_:_:_:)
func vvnextafterf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvnextafterf(p0, p1, p2, p3)
	}


// Calculates the nearest integer for each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvnint(_:_:_:)
func vvnint(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvnint(p0, p1, p2)
	}


// Calculates the nearest integer for each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvnintf(_:_:_:)
func vvnintf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvnintf(p0, p1, p2)
	}


// Raises each element in an array to the power of the corresponding element in a second array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvpow(_:_:_:_:)
func vvpow(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvpow(p0, p1, p2, p3)
	}


// Raises each element in an array to the power of the corresponding element in a second array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvpowf(_:_:_:_:)
func vvpowf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvpowf(p0, p1, p2, p3)
	}


// Calculates the reciprocal of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvrec(_:_:_:)
func vvrec(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvrec(p0, p1, p2)
	}


// Calculates the reciprocal of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvrecf(_:_:_:)
func vvrecf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvrecf(p0, p1, p2)
	}


// Calculates the remainder after dividing each element in an array by the corresponding element in a second array of double-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvremainder(_:_:_:_:)
func vvremainder(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvremainder(p0, p1, p2, p3)
	}


// Calculates the remainder after dividing each element in an array by the corresponding element in a second array of single-precision values. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvremainderf(_:_:_:_:)
func vvremainderf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvremainderf(p0, p1, p2, p3)
	}


// Calculates the reciprocal square root of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvrsqrt(_:_:_:)
func vvrsqrt(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvrsqrt(p0, p1, p2)
	}


// Calculates the reciprocal square root of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvrsqrtf(_:_:_:)
func vvrsqrtf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvrsqrtf(p0, p1, p2)
	}


// Calculates the sine of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsin(_:_:_:)
func vvsin(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvsin(p0, p1, p2)
	}


// Calculates the cosine and sine of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsincos(_:_:_:_:)
func vvsincos(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvsincos(p0, p1, p2, p3)
	}


// Calculates the cosine and sine of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsincosf(_:_:_:_:)
func vvsincosf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) {
	_vvsincosf(p0, p1, p2, p3)
	}


// Calculates the sine of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsinf(_:_:_:)
func vvsinf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvsinf(p0, p1, p2)
	}


// Calculates the hyperbolic sine of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsinh(_:_:_:)
func vvsinh(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvsinh(p0, p1, p2)
	}


// Calculates the hyperbolic sine of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsinhf(_:_:_:)
func vvsinhf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvsinhf(p0, p1, p2)
	}


// Calculates the sine of pi multiplied by each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsinpi(_:_:_:)
func vvsinpi(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvsinpi(p0, p1, p2)
	}


// Calculates the sine of pi multiplied by each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsinpif(_:_:_:)
func vvsinpif(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvsinpif(p0, p1, p2)
	}


// Calculates the square root of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsqrt(_:_:_:)
func vvsqrt(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvsqrt(p0, p1, p2)
	}


// Calculates the square root of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvsqrtf(_:_:_:)
func vvsqrtf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvsqrtf(p0, p1, p2)
	}


// Calculates the tangent of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvtan(_:_:_:)
func vvtan(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvtan(p0, p1, p2)
	}


// Calculates the tangent of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvtanf(_:_:_:)
func vvtanf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvtanf(p0, p1, p2)
	}


// Calculates the hyperbolic tangent of each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvtanh(_:_:_:)
func vvtanh(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvtanh(p0, p1, p2)
	}


// Calculates the hyperbolic tangent of each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvtanhf(_:_:_:)
func vvtanhf(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvtanhf(p0, p1, p2)
	}


// Calculates the tangent of pi multiplied by each element in an array of double-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvtanpi(_:_:_:)
func vvtanpi(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvtanpi(p0, p1, p2)
	}


// Calculates the tangent of pi multiplied by each element in an array of single-precision values. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Accelerate/vvtanpif(_:_:_:)
func vvtanpif(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 unsafe.Pointer) {
	_vvtanpif(p0, p1, p2)
	}




