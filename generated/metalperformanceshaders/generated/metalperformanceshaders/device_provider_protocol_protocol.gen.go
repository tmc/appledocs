// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

// PDeviceProvider is the MPSDeviceProvider protocol interface.
//
// An interface that enables the setting of a Metal device for unarchived objects.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSDeviceProvider
type PDeviceProvider interface {
	// Required methods
	MpsMTLDevice()/* debug [protocol_interface/required_method]: MpsMTLDevice */
}
