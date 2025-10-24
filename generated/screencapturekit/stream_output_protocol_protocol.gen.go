// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

// PStreamOutput is the SCStreamOutput protocol interface.
//
// A delegate protocol your app implements to receive capture stream output events.
//
// Availability:
//   - Mac Catalyst 18.2+
//   - macOS 12.3+
//
// See: doc://com.apple.screencapturekit/documentation/ScreenCaptureKit/SCStreamOutput
type PStreamOutput interface {
	// Optional methods
	StreamDidOutputSampleBufferOfType(stream ISCStream, sampleBuffer SampleBufferRef /* not a class type */, type_ StreamOutputType)
	HasStreamDidOutputSampleBufferOfType() bool
}
