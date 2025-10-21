// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

// Enum types and constants
// GLKFogMode - A mode that describes how the fog component is calculated for the fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKFogMode
type GLKFogMode uint

const (
// GLKFogModeExp - The fog component is calculated as   and clamped to the range  .
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKFogMode/exp
GLKFogModeExp GLKFogMode = 0
// GLKFogModeExp2 - The fog component is calculated as   and clamped to the range  .
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKFogMode/exp2
GLKFogModeExp2 GLKFogMode = 0
// GLKFogModeLinear - The fog component is calculated as   and clamped to the range  .
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKFogMode/linear
GLKFogModeLinear GLKFogMode = 0
)

// GLKLightingType - A constant that describes how lighting is calculated by an effect.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKLightingType
type GLKLightingType uint

const (
// GLKLightingTypePerPixel - Indicates that the inputs to the lighting calculation are interpolated across a triangle and the lighting calculations are performed at each fragment.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKLightingType/perPixel
GLKLightingTypePerPixel GLKLightingType = 0
// GLKLightingTypePerVertex - Indicates that the lighting calculations are performed at each vertex in a triangle and then interpolated across the triangle.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKLightingType/perVertex
GLKLightingTypePerVertex GLKLightingType = 0
)

// GLKTextureEnvMode - The mode used to combine the texture with other color components.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureEnvMode
type GLKTextureEnvMode uint

const (
// GLKTextureEnvModeDecal - The output color is calculated by using the texture’s alpha component to blend the texture’s color with the input color.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureEnvMode/decal
GLKTextureEnvModeDecal GLKTextureEnvMode = 0
// GLKTextureEnvModeModulate - The output color is calculated by multiplying the texture’s color by the input color.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureEnvMode/modulate
GLKTextureEnvModeModulate GLKTextureEnvMode = 0
// GLKTextureEnvModeReplace - The output color is set to the color fetched from the texture. The input color is ignored.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureEnvMode/replace
GLKTextureEnvModeReplace GLKTextureEnvMode = 0
)

// GLKTextureInfoAlphaState - Values that describe the alpha information stored in a source image’s pixel data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfoAlphaState
type GLKTextureInfoAlphaState uint

const (
// GLKTextureInfoAlphaStateNonPremultiplied - Indicates that the color values in the texture were not premultiplied by the alpha value.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfoAlphaState/nonPremultiplied
GLKTextureInfoAlphaStateNonPremultiplied GLKTextureInfoAlphaState = 0
// GLKTextureInfoAlphaStateNone - Indicates that the texture has no alpha information.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfoAlphaState/none
GLKTextureInfoAlphaStateNone GLKTextureInfoAlphaState = 0
// GLKTextureInfoAlphaStatePremultiplied - Indicates that the color values in the texture have already been premultiplied by the alpha value.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfoAlphaState/premultiplied
GLKTextureInfoAlphaStatePremultiplied GLKTextureInfoAlphaState = 0
)

// GLKTextureInfoOrigin - The location of the origin in the original source image.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfoOrigin
type GLKTextureInfoOrigin uint

const (
// GLKTextureInfoOriginBottomLeft - The origin of the texture is in the bottom-left corner.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfoOrigin/bottomLeft
GLKTextureInfoOriginBottomLeft GLKTextureInfoOrigin = 0
// GLKTextureInfoOriginTopLeft - The origin of the texture is in the top-left corner.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfoOrigin/topLeft
GLKTextureInfoOriginTopLeft GLKTextureInfoOrigin = 0
// GLKTextureInfoOriginUnknown - The origin of the texture is not supported.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfoOrigin/unknown
GLKTextureInfoOriginUnknown GLKTextureInfoOrigin = 0
)

// GLKTextureLoaderError - Values to be returned when a texture loader encounters an error.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code
type GLKTextureLoaderError uint

const (
// GLKTextureLoaderErrorAlphaPremultiplicationFailure - The texture source data does not allow the alpha to be premultiplied.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/alphaPremultiplicationFailure
GLKTextureLoaderErrorAlphaPremultiplicationFailure GLKTextureLoaderError = 0
// GLKTextureLoaderErrorCompressedTextureUpload - A compressed texture could not be uploaded.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/compressedTextureUpload
GLKTextureLoaderErrorCompressedTextureUpload GLKTextureLoaderError = 0
// GLKTextureLoaderErrorCubeMapInvalidNumFiles - The incorrect number of files were specified for the cube map.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/cubeMapInvalidNumFiles
GLKTextureLoaderErrorCubeMapInvalidNumFiles GLKTextureLoaderError = 0
// GLKTextureLoaderErrorDataPreprocessingFailure - The data could not be preprocessed correctly.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/dataPreprocessingFailure
GLKTextureLoaderErrorDataPreprocessingFailure GLKTextureLoaderError = 0
// GLKTextureLoaderErrorFileOrURLNotFound - A file could not be found at the path provided.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/fileOrURLNotFound
GLKTextureLoaderErrorFileOrURLNotFound GLKTextureLoaderError = 0
// GLKTextureLoaderErrorIncompatibleFormatSRGB - The decoded data was in an incompatible format for an sRGB texture.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/incompatibleFormatSRGB
GLKTextureLoaderErrorIncompatibleFormatSRGB GLKTextureLoaderError = 0
// GLKTextureLoaderErrorInvalidCGImage - The image provided was invalid.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/invalidCGImage
GLKTextureLoaderErrorInvalidCGImage GLKTextureLoaderError = 0
// GLKTextureLoaderErrorInvalidEAGLContext - The EAGL context was not a valid context.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/invalidEAGLContext
GLKTextureLoaderErrorInvalidEAGLContext GLKTextureLoaderError = 0
// GLKTextureLoaderErrorInvalidNSData - The data provided is not in a recognized image format.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/invalidNSData
GLKTextureLoaderErrorInvalidNSData GLKTextureLoaderError = 0
// GLKTextureLoaderErrorMipmapUnsupported - The texture source data does not allow mipmaps to be generated.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/mipmapUnsupported
GLKTextureLoaderErrorMipmapUnsupported GLKTextureLoaderError = 0
// GLKTextureLoaderErrorPVRAtlasUnsupported - Cube maps may not be compressed in PVRTC format.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/pvrAtlasUnsupported
GLKTextureLoaderErrorPVRAtlasUnsupported GLKTextureLoaderError = 0
// GLKTextureLoaderErrorReorientationFailure - The texture source data does not allow the image to be reoriented.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/reorientationFailure
GLKTextureLoaderErrorReorientationFailure GLKTextureLoaderError = 0
// GLKTextureLoaderErrorUncompressedTextureUpload - An uncompressed texture could not be uploaded.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/uncompressedTextureUpload
GLKTextureLoaderErrorUncompressedTextureUpload GLKTextureLoaderError = 0
// GLKTextureLoaderErrorUnknownFileType - The file was in an unrecognized format.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/unknownFileType
GLKTextureLoaderErrorUnknownFileType GLKTextureLoaderError = 0
// GLKTextureLoaderErrorUnknownPathType - The path type was unrecognized.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/unknownPathType
GLKTextureLoaderErrorUnknownPathType GLKTextureLoaderError = 0
// GLKTextureLoaderErrorUnsupportedBitDepth - The data in the source image has an unsupported bit depth.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/unsupportedBitDepth
GLKTextureLoaderErrorUnsupportedBitDepth GLKTextureLoaderError = 0
// GLKTextureLoaderErrorUnsupportedCubeMapDimensions - The cube map’s dimensions are incorrect.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/unsupportedCubeMapDimensions
GLKTextureLoaderErrorUnsupportedCubeMapDimensions GLKTextureLoaderError = 0
// GLKTextureLoaderErrorUnsupportedOrientation - The texture source data is stored with an unsupported origin position.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/unsupportedOrientation
GLKTextureLoaderErrorUnsupportedOrientation GLKTextureLoaderError = 0
// GLKTextureLoaderErrorUnsupportedPVRFormat - The data in the PVRTC compressed format is in an unsupported format.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/unsupportedPVRFormat
GLKTextureLoaderErrorUnsupportedPVRFormat GLKTextureLoaderError = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code/unsupportedTextureTarget
GLKTextureLoaderErrorUnsupportedTextureTarget GLKTextureLoaderError = 0
)

// GLKTextureTarget - The kind of texture pointed to by the property.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureTarget
type GLKTextureTarget uint

const (
// GLKTextureTarget2D - The texture is a 2D texture.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureTarget/target2D
GLKTextureTarget2D GLKTextureTarget = 0
// GLKTextureTargetCt - The number of items in the enumeration.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureTarget/targetCt
GLKTextureTargetCt GLKTextureTarget = 0
// GLKTextureTargetCubeMap - The texture is a set of six textures that make up a cube map.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureTarget/targetCubeMap
GLKTextureTargetCubeMap GLKTextureTarget = 0
)

// GLKVertexAttrib - Values used as indices in OpenGL code to associate vertex data with an attribute in a named shader effect.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKVertexAttrib
type GLKVertexAttrib uint

const (
// GLKVertexAttribColor - This index is used to provide the vertex color to a shader.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKVertexAttrib/color
GLKVertexAttribColor GLKVertexAttrib = 0
// GLKVertexAttribNormal - This index is used to provide the vertex normal to a shader.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKVertexAttrib/normal
GLKVertexAttribNormal GLKVertexAttrib = 0
// GLKVertexAttribPosition - This index is used to provide the vertex position to a shader.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKVertexAttrib/position
GLKVertexAttribPosition GLKVertexAttrib = 0
// GLKVertexAttribTexCoord0 - This index is used to provide a set of texture coordinates to a shader.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKVertexAttrib/texCoord0
GLKVertexAttribTexCoord0 GLKVertexAttrib = 0
// GLKVertexAttribTexCoord1 - This index is used to provide the second set of texture coordinates to a shader.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKVertexAttrib/texCoord1
GLKVertexAttribTexCoord1 GLKVertexAttrib = 0
)

// GLKViewDrawableColorFormat - The format of the color renderbuffer.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewDrawableColorFormat
type GLKViewDrawableColorFormat uint

const (
// GLKViewDrawableColorFormatRGB565 - An RGB565 format.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewDrawableColorFormat/RGB565
GLKViewDrawableColorFormatRGB565 GLKViewDrawableColorFormat = 0
// GLKViewDrawableColorFormatRGBA8888 - An RGBA8888 format.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewDrawableColorFormat/RGBA8888
GLKViewDrawableColorFormatRGBA8888 GLKViewDrawableColorFormat = 0
// GLKViewDrawableColorFormatSRGBA8888 - An sRGBA8888 format.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewDrawableColorFormat/SRGBA8888
GLKViewDrawableColorFormatSRGBA8888 GLKViewDrawableColorFormat = 0
)

// GLKViewDrawableDepthFormat - The format of the depth renderbuffer.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewDrawableDepthFormat
type GLKViewDrawableDepthFormat uint

const (
// GLKViewDrawableDepthFormat16 - A 16-bit depth entry for each pixel.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewDrawableDepthFormat/format16
GLKViewDrawableDepthFormat16 GLKViewDrawableDepthFormat = 0
// GLKViewDrawableDepthFormat24 - A 24-bit depth entry for each pixel.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewDrawableDepthFormat/format24
GLKViewDrawableDepthFormat24 GLKViewDrawableDepthFormat = 0
// GLKViewDrawableDepthFormatNone - The underlying framebuffer object has no depth buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewDrawableDepthFormat/formatNone
GLKViewDrawableDepthFormatNone GLKViewDrawableDepthFormat = 0
)

// GLKViewDrawableMultisample - The format of the multisampling buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewDrawableMultisample
type GLKViewDrawableMultisample uint

const (
// GLKViewDrawableMultisample4X - Multisampling is enabled.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewDrawableMultisample/multisample4X
GLKViewDrawableMultisample4X GLKViewDrawableMultisample = 0
// GLKViewDrawableMultisampleNone - Multisampling is not enabled.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewDrawableMultisample/multisampleNone
GLKViewDrawableMultisampleNone GLKViewDrawableMultisample = 0
)

// GLKViewDrawableStencilFormat - The format of the stencil renderbuffer.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewDrawableStencilFormat
type GLKViewDrawableStencilFormat uint

const (
// GLKViewDrawableStencilFormat8 - An 8-bit stencil entry for each pixel.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewDrawableStencilFormat/format8
GLKViewDrawableStencilFormat8 GLKViewDrawableStencilFormat = 0
// GLKViewDrawableStencilFormatNone - The underlying framebuffer object has no stencil buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewDrawableStencilFormat/formatNone
GLKViewDrawableStencilFormatNone GLKViewDrawableStencilFormat = 0
)


