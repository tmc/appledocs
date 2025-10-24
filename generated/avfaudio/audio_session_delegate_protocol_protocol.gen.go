// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PAudioSessionDelegate is the AVAudioSessionDelegate protocol interface.
//
// A protocol that defines responses to changes in state for the audio session.
//
// Availability:
//   - Mac Catalyst 14.0+ (Deprecated in 14.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// See: doc://com.apple.avfaudio/documentation/AVFAudio/AVAudioSessionDelegate
type PAudioSessionDelegate interface {
	// Optional methods
	BeginInterruption()
	HasBeginInterruption() bool
	EndInterruption()
	HasEndInterruption() bool
	EndInterruptionWithFlags(flags uint)
	HasEndInterruptionWithFlags() bool
	InputIsAvailableChanged(isInputAvailable bool)
	HasInputIsAvailableChanged() bool
}

// AudioSessionDelegate is a delegate implementation builder for the PAudioSessionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AudioSessionDelegate struct {
	_BeginInterruption func()
	_EndInterruption func()
	_EndInterruptionWithFlags func(flags uint)
	_InputIsAvailableChanged func(isInputAvailable bool)
}

// SetBeginInterruption sets the handler for the BeginInterruption delegate method.
//
// Called after your audio session is interrupted.
func (d *AudioSessionDelegate) SetBeginInterruption(f func()) {
	d._BeginInterruption = f
}

// SetEndInterruption sets the handler for the EndInterruption delegate method.
//
// Called after your audio session interruption ends.
func (d *AudioSessionDelegate) SetEndInterruption(f func()) {
	d._EndInterruption = f
}

// SetEndInterruptionWithFlags sets the handler for the EndInterruptionWithFlags delegate method.
//
// Called after your audio session interruption ends, with flags indicating the state of the audio session.
func (d *AudioSessionDelegate) SetEndInterruptionWithFlags(f func(flags uint)) {
	d._EndInterruptionWithFlags = f
}

// SetInputIsAvailableChanged sets the handler for the InputIsAvailableChanged delegate method.
//
// Called after the availability of audio input changes on a device.
func (d *AudioSessionDelegate) SetInputIsAvailableChanged(f func(isInputAvailable bool)) {
	d._InputIsAvailableChanged = f
}

// BeginInterruption implements the PAudioSessionDelegate interface.
func (d *AudioSessionDelegate) BeginInterruption() {
	if d._BeginInterruption != nil {
		d._BeginInterruption()
	}
}

// HasBeginInterruption returns true if a handler for BeginInterruption has been set.
func (d *AudioSessionDelegate) HasBeginInterruption() bool {
	return d._BeginInterruption != nil
}

// EndInterruption implements the PAudioSessionDelegate interface.
func (d *AudioSessionDelegate) EndInterruption() {
	if d._EndInterruption != nil {
		d._EndInterruption()
	}
}

// HasEndInterruption returns true if a handler for EndInterruption has been set.
func (d *AudioSessionDelegate) HasEndInterruption() bool {
	return d._EndInterruption != nil
}

// EndInterruptionWithFlags implements the PAudioSessionDelegate interface.
func (d *AudioSessionDelegate) EndInterruptionWithFlags(flags uint) {
	if d._EndInterruptionWithFlags != nil {
		d._EndInterruptionWithFlags(flags)
	}
}

// HasEndInterruptionWithFlags returns true if a handler for EndInterruptionWithFlags has been set.
func (d *AudioSessionDelegate) HasEndInterruptionWithFlags() bool {
	return d._EndInterruptionWithFlags != nil
}

// InputIsAvailableChanged implements the PAudioSessionDelegate interface.
func (d *AudioSessionDelegate) InputIsAvailableChanged(isInputAvailable bool) {
	if d._InputIsAvailableChanged != nil {
		d._InputIsAvailableChanged(isInputAvailable)
	}
}

// HasInputIsAvailableChanged returns true if a handler for InputIsAvailableChanged has been set.
func (d *AudioSessionDelegate) HasInputIsAvailableChanged() bool {
	return d._InputIsAvailableChanged != nil
}
