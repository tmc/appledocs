// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PStreamDelegate is the SCStreamDelegate protocol interface.
//
// A delegate protocol your app implements to respond to stream events.
//
// Availability:
//   - Mac Catalyst 18.2+
//   - macOS 12.3+
//
// See: doc://com.apple.screencapturekit/documentation/ScreenCaptureKit/SCStreamDelegate
type PStreamDelegate interface {
	// Optional methods
	OutputVideoEffectDidStartForStream(stream ISCStream)
	HasOutputVideoEffectDidStartForStream() bool
	OutputVideoEffectDidStopForStream(stream ISCStream)
	HasOutputVideoEffectDidStopForStream() bool
	StreamDidStopWithError(stream ISCStream, error_ objc.IObject /* cross-framework: Error */)
	HasStreamDidStopWithError() bool
	StreamDidBecomeActive(stream ISCStream)
	HasStreamDidBecomeActive() bool
	StreamDidBecomeInactive(stream ISCStream)
	HasStreamDidBecomeInactive() bool
}

// StreamDelegate is a delegate implementation builder for the PStreamDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type StreamDelegate struct {
	_OutputVideoEffectDidStartForStream func(stream ISCStream)
	_OutputVideoEffectDidStopForStream func(stream ISCStream)
	_StreamDidStopWithError func(stream ISCStream, error_ objc.IObject /* cross-framework: Error */)
	_StreamDidBecomeActive func(stream ISCStream)
	_StreamDidBecomeInactive func(stream ISCStream)
}

// SetOutputVideoEffectDidStartForStream sets the handler for the OutputVideoEffectDidStartForStream delegate method.
//
// Tells the delegate that Presenter Overlay started.
func (d *StreamDelegate) SetOutputVideoEffectDidStartForStream(f func(stream ISCStream)) {
	d._OutputVideoEffectDidStartForStream = f
}

// SetOutputVideoEffectDidStopForStream sets the handler for the OutputVideoEffectDidStopForStream delegate method.
//
// Tells the delegate that Presenter Overlay stopped.
func (d *StreamDelegate) SetOutputVideoEffectDidStopForStream(f func(stream ISCStream)) {
	d._OutputVideoEffectDidStopForStream = f
}

// SetStreamDidStopWithError sets the handler for the StreamDidStopWithError delegate method.
//
// Tells the delegate that the stream stopped with an error.
func (d *StreamDelegate) SetStreamDidStopWithError(f func(stream ISCStream, error_ objc.IObject /* cross-framework: Error */)) {
	d._StreamDidStopWithError = f
}

// SetStreamDidBecomeActive sets the handler for the StreamDidBecomeActive delegate method.
func (d *StreamDelegate) SetStreamDidBecomeActive(f func(stream ISCStream)) {
	d._StreamDidBecomeActive = f
}

// SetStreamDidBecomeInactive sets the handler for the StreamDidBecomeInactive delegate method.
func (d *StreamDelegate) SetStreamDidBecomeInactive(f func(stream ISCStream)) {
	d._StreamDidBecomeInactive = f
}

// OutputVideoEffectDidStartForStream implements the PStreamDelegate interface.
func (d *StreamDelegate) OutputVideoEffectDidStartForStream(stream ISCStream) {
	if d._OutputVideoEffectDidStartForStream != nil {
		d._OutputVideoEffectDidStartForStream(stream)
	}
}

// HasOutputVideoEffectDidStartForStream returns true if a handler for OutputVideoEffectDidStartForStream has been set.
func (d *StreamDelegate) HasOutputVideoEffectDidStartForStream() bool {
	return d._OutputVideoEffectDidStartForStream != nil
}

// OutputVideoEffectDidStopForStream implements the PStreamDelegate interface.
func (d *StreamDelegate) OutputVideoEffectDidStopForStream(stream ISCStream) {
	if d._OutputVideoEffectDidStopForStream != nil {
		d._OutputVideoEffectDidStopForStream(stream)
	}
}

// HasOutputVideoEffectDidStopForStream returns true if a handler for OutputVideoEffectDidStopForStream has been set.
func (d *StreamDelegate) HasOutputVideoEffectDidStopForStream() bool {
	return d._OutputVideoEffectDidStopForStream != nil
}

// StreamDidStopWithError implements the PStreamDelegate interface.
func (d *StreamDelegate) StreamDidStopWithError(stream ISCStream, error_ objc.IObject /* cross-framework: Error */) {
	if d._StreamDidStopWithError != nil {
		d._StreamDidStopWithError(stream, error_)
	}
}

// HasStreamDidStopWithError returns true if a handler for StreamDidStopWithError has been set.
func (d *StreamDelegate) HasStreamDidStopWithError() bool {
	return d._StreamDidStopWithError != nil
}

// StreamDidBecomeActive implements the PStreamDelegate interface.
func (d *StreamDelegate) StreamDidBecomeActive(stream ISCStream) {
	if d._StreamDidBecomeActive != nil {
		d._StreamDidBecomeActive(stream)
	}
}

// HasStreamDidBecomeActive returns true if a handler for StreamDidBecomeActive has been set.
func (d *StreamDelegate) HasStreamDidBecomeActive() bool {
	return d._StreamDidBecomeActive != nil
}

// StreamDidBecomeInactive implements the PStreamDelegate interface.
func (d *StreamDelegate) StreamDidBecomeInactive(stream ISCStream) {
	if d._StreamDidBecomeInactive != nil {
		d._StreamDidBecomeInactive(stream)
	}
}

// HasStreamDidBecomeInactive returns true if a handler for StreamDidBecomeInactive has been set.
func (d *StreamDelegate) HasStreamDidBecomeInactive() bool {
	return d._StreamDidBecomeInactive != nil
}
