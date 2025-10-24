// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

// PGramMatrixCallback is the MPSNNGramMatrixCallback protocol interface.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSNNGramMatrixCallback
type PGramMatrixCallback interface {
	// Required methods
	Alpha()/* debug [protocol_interface/required_method]: Alpha */
	AlphaForSourceImageDestinationImage(sourceImage IImage, destinationImage IImage) float32/* debug [protocol_interface/required_method]: AlphaForSourceImageDestinationImage */
}
