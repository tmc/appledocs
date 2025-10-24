// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PAudioRecorderDelegate is the AVAudioRecorderDelegate protocol interface.
//
// A protocol that defines the methods to respond to audio recording events and encoding errors.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.7+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 4.0+
//
// See: doc://com.apple.avfaudio/documentation/AVFAudio/AVAudioRecorderDelegate
type PAudioRecorderDelegate interface {
	// Optional methods
	AudioRecorderBeginInterruption(recorder IAVAudioRecorder)
	HasAudioRecorderBeginInterruption() bool
	AudioRecorderDidFinishRecordingSuccessfully(recorder IAVAudioRecorder, flag bool)
	HasAudioRecorderDidFinishRecordingSuccessfully() bool
	AudioRecorderEncodeErrorDidOccurError(recorder IAVAudioRecorder, error_ objc.IObject /* cross-framework: Error */)
	HasAudioRecorderEncodeErrorDidOccurError() bool
	AudioRecorderEndInterruption(recorder IAVAudioRecorder)
	HasAudioRecorderEndInterruption() bool
	AudioRecorderEndInterruptionWithFlags(recorder IAVAudioRecorder, flags uint)
	HasAudioRecorderEndInterruptionWithFlags() bool
	AudioRecorderEndInterruptionWithOptions(recorder IAVAudioRecorder, flags uint)
	HasAudioRecorderEndInterruptionWithOptions() bool
}

// AudioRecorderDelegate is a delegate implementation builder for the PAudioRecorderDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AudioRecorderDelegate struct {
	_AudioRecorderBeginInterruption func(recorder IAVAudioRecorder)
	_AudioRecorderDidFinishRecordingSuccessfully func(recorder IAVAudioRecorder, flag bool)
	_AudioRecorderEncodeErrorDidOccurError func(recorder IAVAudioRecorder, error_ objc.IObject /* cross-framework: Error */)
	_AudioRecorderEndInterruption func(recorder IAVAudioRecorder)
	_AudioRecorderEndInterruptionWithFlags func(recorder IAVAudioRecorder, flags uint)
	_AudioRecorderEndInterruptionWithOptions func(recorder IAVAudioRecorder, flags uint)
}

// SetAudioRecorderBeginInterruption sets the handler for the AudioRecorderBeginInterruption delegate method.
//
// Tells the delegate that the system interrupted the audio recording.
func (d *AudioRecorderDelegate) SetAudioRecorderBeginInterruption(f func(recorder IAVAudioRecorder)) {
	d._AudioRecorderBeginInterruption = f
}

// SetAudioRecorderDidFinishRecordingSuccessfully sets the handler for the AudioRecorderDidFinishRecordingSuccessfully delegate method.
//
// Tells the delegate when recording stops or finishes due to reaching its time limit.
func (d *AudioRecorderDelegate) SetAudioRecorderDidFinishRecordingSuccessfully(f func(recorder IAVAudioRecorder, flag bool)) {
	d._AudioRecorderDidFinishRecordingSuccessfully = f
}

// SetAudioRecorderEncodeErrorDidOccurError sets the handler for the AudioRecorderEncodeErrorDidOccurError delegate method.
//
// Tells the delegate that the audio recorder encountered an encoding error during recording.
func (d *AudioRecorderDelegate) SetAudioRecorderEncodeErrorDidOccurError(f func(recorder IAVAudioRecorder, error_ objc.IObject /* cross-framework: Error */)) {
	d._AudioRecorderEncodeErrorDidOccurError = f
}

// SetAudioRecorderEndInterruption sets the handler for the AudioRecorderEndInterruption delegate method.
//
// Tells the delegate that the audio session interruption ended.
func (d *AudioRecorderDelegate) SetAudioRecorderEndInterruption(f func(recorder IAVAudioRecorder)) {
	d._AudioRecorderEndInterruption = f
}

// SetAudioRecorderEndInterruptionWithFlags sets the handler for the AudioRecorderEndInterruptionWithFlags delegate method.
//
// Tells the delegate that the audio session interruption ended with flags.
func (d *AudioRecorderDelegate) SetAudioRecorderEndInterruptionWithFlags(f func(recorder IAVAudioRecorder, flags uint)) {
	d._AudioRecorderEndInterruptionWithFlags = f
}

// SetAudioRecorderEndInterruptionWithOptions sets the handler for the AudioRecorderEndInterruptionWithOptions delegate method.
//
// Tells the delegate that the audio session interruption ended with options.
func (d *AudioRecorderDelegate) SetAudioRecorderEndInterruptionWithOptions(f func(recorder IAVAudioRecorder, flags uint)) {
	d._AudioRecorderEndInterruptionWithOptions = f
}

// AudioRecorderBeginInterruption implements the PAudioRecorderDelegate interface.
func (d *AudioRecorderDelegate) AudioRecorderBeginInterruption(recorder IAVAudioRecorder) {
	if d._AudioRecorderBeginInterruption != nil {
		d._AudioRecorderBeginInterruption(recorder)
	}
}

// HasAudioRecorderBeginInterruption returns true if a handler for AudioRecorderBeginInterruption has been set.
func (d *AudioRecorderDelegate) HasAudioRecorderBeginInterruption() bool {
	return d._AudioRecorderBeginInterruption != nil
}

// AudioRecorderDidFinishRecordingSuccessfully implements the PAudioRecorderDelegate interface.
func (d *AudioRecorderDelegate) AudioRecorderDidFinishRecordingSuccessfully(recorder IAVAudioRecorder, flag bool) {
	if d._AudioRecorderDidFinishRecordingSuccessfully != nil {
		d._AudioRecorderDidFinishRecordingSuccessfully(recorder, flag)
	}
}

// HasAudioRecorderDidFinishRecordingSuccessfully returns true if a handler for AudioRecorderDidFinishRecordingSuccessfully has been set.
func (d *AudioRecorderDelegate) HasAudioRecorderDidFinishRecordingSuccessfully() bool {
	return d._AudioRecorderDidFinishRecordingSuccessfully != nil
}

// AudioRecorderEncodeErrorDidOccurError implements the PAudioRecorderDelegate interface.
func (d *AudioRecorderDelegate) AudioRecorderEncodeErrorDidOccurError(recorder IAVAudioRecorder, error_ objc.IObject /* cross-framework: Error */) {
	if d._AudioRecorderEncodeErrorDidOccurError != nil {
		d._AudioRecorderEncodeErrorDidOccurError(recorder, error_)
	}
}

// HasAudioRecorderEncodeErrorDidOccurError returns true if a handler for AudioRecorderEncodeErrorDidOccurError has been set.
func (d *AudioRecorderDelegate) HasAudioRecorderEncodeErrorDidOccurError() bool {
	return d._AudioRecorderEncodeErrorDidOccurError != nil
}

// AudioRecorderEndInterruption implements the PAudioRecorderDelegate interface.
func (d *AudioRecorderDelegate) AudioRecorderEndInterruption(recorder IAVAudioRecorder) {
	if d._AudioRecorderEndInterruption != nil {
		d._AudioRecorderEndInterruption(recorder)
	}
}

// HasAudioRecorderEndInterruption returns true if a handler for AudioRecorderEndInterruption has been set.
func (d *AudioRecorderDelegate) HasAudioRecorderEndInterruption() bool {
	return d._AudioRecorderEndInterruption != nil
}

// AudioRecorderEndInterruptionWithFlags implements the PAudioRecorderDelegate interface.
func (d *AudioRecorderDelegate) AudioRecorderEndInterruptionWithFlags(recorder IAVAudioRecorder, flags uint) {
	if d._AudioRecorderEndInterruptionWithFlags != nil {
		d._AudioRecorderEndInterruptionWithFlags(recorder, flags)
	}
}

// HasAudioRecorderEndInterruptionWithFlags returns true if a handler for AudioRecorderEndInterruptionWithFlags has been set.
func (d *AudioRecorderDelegate) HasAudioRecorderEndInterruptionWithFlags() bool {
	return d._AudioRecorderEndInterruptionWithFlags != nil
}

// AudioRecorderEndInterruptionWithOptions implements the PAudioRecorderDelegate interface.
func (d *AudioRecorderDelegate) AudioRecorderEndInterruptionWithOptions(recorder IAVAudioRecorder, flags uint) {
	if d._AudioRecorderEndInterruptionWithOptions != nil {
		d._AudioRecorderEndInterruptionWithOptions(recorder, flags)
	}
}

// HasAudioRecorderEndInterruptionWithOptions returns true if a handler for AudioRecorderEndInterruptionWithOptions has been set.
func (d *AudioRecorderDelegate) HasAudioRecorderEndInterruptionWithOptions() bool {
	return d._AudioRecorderEndInterruptionWithOptions != nil
}
