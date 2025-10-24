// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

// PLossCallback is the MPSNNLossCallback protocol interface.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSNNLossCallback
type PLossCallback interface {
	// Required methods
	ScalarWeight()/* debug [protocol_interface/required_method]: ScalarWeight */
	ScalarWeightForSourceImageDestinationImage(sourceImage IImage, destinationImage IImage) float32/* debug [protocol_interface/required_method]: ScalarWeightForSourceImageDestinationImage */
}
