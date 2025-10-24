// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

// Enum types and constants
// MTLArgumentBuffersTier - The values that determine the limits and capabilities of argument buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentBuffersTier
type MTLArgumentBuffersTier uint

// MTLIndirectCommandType - The types of commands that you can encode into the indirect command buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType
type MTLIndirectCommandType uint

// MTLMathFloatingPointFunctions - Indicates which FP32 math functions Metal uses.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMathFloatingPointFunctions
type MTLMathFloatingPointFunctions uint

// MTLTextureType - The dimension of each image, including whether multiple images are arranged into an array or a cube.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType
type MTLTextureType uint

// MTLTextureUsage - An enumeration for the various options that determine how you can use a texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage
type MTLTextureUsage uint

const (
	// MTLTextureUsagePixelFormatView - An option to create texture views with a different component layout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/pixelFormatView
	MTLTextureUsagePixelFormatView MTLTextureUsage = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/shaderAtomic
	MTLTextureUsageShaderAtomic MTLTextureUsage = 0
	// MTLTextureUsageShaderWrite - An option for writing to the texture in a shader.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/shaderWrite
	MTLTextureUsageShaderWrite MTLTextureUsage = 0
	// MTLTextureUsageUnknown - An option for a texture whose usage is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/unknown
	MTLTextureUsageUnknown MTLTextureUsage = 0
)

// MTLTransformType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTransformType
type MTLTransformType uint


