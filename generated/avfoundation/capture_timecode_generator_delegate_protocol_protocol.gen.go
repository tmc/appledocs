// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// PCaptureTimecodeGeneratorDelegate is the AVCaptureTimecodeGeneratorDelegate protocol interface.
//
// A protocol for receiving real-time timecode updates and error notifications from a timecode generator.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureTimecodeGeneratorDelegate
type PCaptureTimecodeGeneratorDelegate interface {
	// Required methods
	TimecodeGeneratorDidReceiveUpdateFromSource(generator IAVCaptureTimecodeGenerator, timecode CaptureTimecode, source IAVCaptureTimecodeSource)
	TimecodeGeneratorDidUpdateAvailableSources(generator IAVCaptureTimecodeGenerator, availableSources []CaptureTimecodeSource)
	TimecodeGeneratorTransitionedToSynchronizationStatusForSource(generator IAVCaptureTimecodeGenerator, synchronizationStatus CaptureTimecodeGeneratorSynchronizationStatus, source IAVCaptureTimecodeSource)
}
