// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

// Enum types and constants
// MPSAlphaType - Premultiplication description for the color channels of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAlphaType
type MPSAlphaType uint

const (
	// MPSAlphaTypeAlphaIsOne - Alpha is guaranteed to be 1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAlphaType/alphaIsOne
	MPSAlphaTypeAlphaIsOne MPSAlphaType = 0
	// MPSAlphaTypeNonPremultiplied - The image is not premultiplied by alpha.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAlphaType/nonPremultiplied
	MPSAlphaTypeNonPremultiplied MPSAlphaType = 0
	// MPSAlphaTypePremultiplied - The image is premultiplied by alpha.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAlphaType/premultiplied
	MPSAlphaTypePremultiplied MPSAlphaType = 0
)

// MPSImageEdgeMode - The options used to control the edge behavior of an image filter when it reads outside the bounds of a source texture.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEdgeMode
type MPSImageEdgeMode uint

const (
	// MPSImageEdgeModeClamp - Out-of-bound pixels are clamped to the nearest edge pixel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEdgeMode/clamp
	MPSImageEdgeModeClamp MPSImageEdgeMode = 0
	// MPSImageEdgeModeZero - Out-of-bound pixels are set to   for images without an alpha channel or   for images with an alpha channel, as defined by their pixel format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEdgeMode/zero
	MPSImageEdgeModeZero MPSImageEdgeMode = 0
)

// MPSNNComparisonType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonType
type MPSNNComparisonType uint

// MPSRNNSequenceDirection - Directions that a sequence of inputs can be processed by a recurrent neural network layer.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNSequenceDirection
type MPSRNNSequenceDirection uint


