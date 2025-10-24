// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PAudioPlayerDelegate is the AVAudioPlayerDelegate protocol interface.
//
// A protocol that defines the methods to respond to audio playback events and decoding errors.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//   - watchOS 3.0+
//
// See: doc://com.apple.avfaudio/documentation/AVFAudio/AVAudioPlayerDelegate
type PAudioPlayerDelegate interface {
	// Optional methods
	AudioPlayerBeginInterruption(player IAVAudioPlayer)
	HasAudioPlayerBeginInterruption() bool
	AudioPlayerDecodeErrorDidOccurError(player IAVAudioPlayer, error_ objc.IObject /* cross-framework: Error */)
	HasAudioPlayerDecodeErrorDidOccurError() bool
	AudioPlayerDidFinishPlayingSuccessfully(player IAVAudioPlayer, flag bool)
	HasAudioPlayerDidFinishPlayingSuccessfully() bool
	AudioPlayerEndInterruption(player IAVAudioPlayer)
	HasAudioPlayerEndInterruption() bool
	AudioPlayerEndInterruptionWithFlags(player IAVAudioPlayer, flags uint)
	HasAudioPlayerEndInterruptionWithFlags() bool
	AudioPlayerEndInterruptionWithOptions(player IAVAudioPlayer, flags uint)
	HasAudioPlayerEndInterruptionWithOptions() bool
}

// AudioPlayerDelegate is a delegate implementation builder for the PAudioPlayerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AudioPlayerDelegate struct {
	_AudioPlayerBeginInterruption func(player IAVAudioPlayer)
	_AudioPlayerDecodeErrorDidOccurError func(player IAVAudioPlayer, error_ objc.IObject /* cross-framework: Error */)
	_AudioPlayerDidFinishPlayingSuccessfully func(player IAVAudioPlayer, flag bool)
	_AudioPlayerEndInterruption func(player IAVAudioPlayer)
	_AudioPlayerEndInterruptionWithFlags func(player IAVAudioPlayer, flags uint)
	_AudioPlayerEndInterruptionWithOptions func(player IAVAudioPlayer, flags uint)
}

// SetAudioPlayerBeginInterruption sets the handler for the AudioPlayerBeginInterruption delegate method.
//
// Tells the delegate when the system interrupts the audio player’s playback.
func (d *AudioPlayerDelegate) SetAudioPlayerBeginInterruption(f func(player IAVAudioPlayer)) {
	d._AudioPlayerBeginInterruption = f
}

// SetAudioPlayerDecodeErrorDidOccurError sets the handler for the AudioPlayerDecodeErrorDidOccurError delegate method.
//
// Tells the delegate when an audio player encounters a decoding error during playback.
func (d *AudioPlayerDelegate) SetAudioPlayerDecodeErrorDidOccurError(f func(player IAVAudioPlayer, error_ objc.IObject /* cross-framework: Error */)) {
	d._AudioPlayerDecodeErrorDidOccurError = f
}

// SetAudioPlayerDidFinishPlayingSuccessfully sets the handler for the AudioPlayerDidFinishPlayingSuccessfully delegate method.
//
// Tells the delegate when the audio finishes playing.
func (d *AudioPlayerDelegate) SetAudioPlayerDidFinishPlayingSuccessfully(f func(player IAVAudioPlayer, flag bool)) {
	d._AudioPlayerDidFinishPlayingSuccessfully = f
}

// SetAudioPlayerEndInterruption sets the handler for the AudioPlayerEndInterruption delegate method.
//
// Tells the delegate when the audio session interruption ends.
func (d *AudioPlayerDelegate) SetAudioPlayerEndInterruption(f func(player IAVAudioPlayer)) {
	d._AudioPlayerEndInterruption = f
}

// SetAudioPlayerEndInterruptionWithFlags sets the handler for the AudioPlayerEndInterruptionWithFlags delegate method.
//
// Tells the delegate when the audio session interruption ends with flags.
func (d *AudioPlayerDelegate) SetAudioPlayerEndInterruptionWithFlags(f func(player IAVAudioPlayer, flags uint)) {
	d._AudioPlayerEndInterruptionWithFlags = f
}

// SetAudioPlayerEndInterruptionWithOptions sets the handler for the AudioPlayerEndInterruptionWithOptions delegate method.
//
// Tells the delegate when the audio session interruption ends with options.
func (d *AudioPlayerDelegate) SetAudioPlayerEndInterruptionWithOptions(f func(player IAVAudioPlayer, flags uint)) {
	d._AudioPlayerEndInterruptionWithOptions = f
}

// AudioPlayerBeginInterruption implements the PAudioPlayerDelegate interface.
func (d *AudioPlayerDelegate) AudioPlayerBeginInterruption(player IAVAudioPlayer) {
	if d._AudioPlayerBeginInterruption != nil {
		d._AudioPlayerBeginInterruption(player)
	}
}

// HasAudioPlayerBeginInterruption returns true if a handler for AudioPlayerBeginInterruption has been set.
func (d *AudioPlayerDelegate) HasAudioPlayerBeginInterruption() bool {
	return d._AudioPlayerBeginInterruption != nil
}

// AudioPlayerDecodeErrorDidOccurError implements the PAudioPlayerDelegate interface.
func (d *AudioPlayerDelegate) AudioPlayerDecodeErrorDidOccurError(player IAVAudioPlayer, error_ objc.IObject /* cross-framework: Error */) {
	if d._AudioPlayerDecodeErrorDidOccurError != nil {
		d._AudioPlayerDecodeErrorDidOccurError(player, error_)
	}
}

// HasAudioPlayerDecodeErrorDidOccurError returns true if a handler for AudioPlayerDecodeErrorDidOccurError has been set.
func (d *AudioPlayerDelegate) HasAudioPlayerDecodeErrorDidOccurError() bool {
	return d._AudioPlayerDecodeErrorDidOccurError != nil
}

// AudioPlayerDidFinishPlayingSuccessfully implements the PAudioPlayerDelegate interface.
func (d *AudioPlayerDelegate) AudioPlayerDidFinishPlayingSuccessfully(player IAVAudioPlayer, flag bool) {
	if d._AudioPlayerDidFinishPlayingSuccessfully != nil {
		d._AudioPlayerDidFinishPlayingSuccessfully(player, flag)
	}
}

// HasAudioPlayerDidFinishPlayingSuccessfully returns true if a handler for AudioPlayerDidFinishPlayingSuccessfully has been set.
func (d *AudioPlayerDelegate) HasAudioPlayerDidFinishPlayingSuccessfully() bool {
	return d._AudioPlayerDidFinishPlayingSuccessfully != nil
}

// AudioPlayerEndInterruption implements the PAudioPlayerDelegate interface.
func (d *AudioPlayerDelegate) AudioPlayerEndInterruption(player IAVAudioPlayer) {
	if d._AudioPlayerEndInterruption != nil {
		d._AudioPlayerEndInterruption(player)
	}
}

// HasAudioPlayerEndInterruption returns true if a handler for AudioPlayerEndInterruption has been set.
func (d *AudioPlayerDelegate) HasAudioPlayerEndInterruption() bool {
	return d._AudioPlayerEndInterruption != nil
}

// AudioPlayerEndInterruptionWithFlags implements the PAudioPlayerDelegate interface.
func (d *AudioPlayerDelegate) AudioPlayerEndInterruptionWithFlags(player IAVAudioPlayer, flags uint) {
	if d._AudioPlayerEndInterruptionWithFlags != nil {
		d._AudioPlayerEndInterruptionWithFlags(player, flags)
	}
}

// HasAudioPlayerEndInterruptionWithFlags returns true if a handler for AudioPlayerEndInterruptionWithFlags has been set.
func (d *AudioPlayerDelegate) HasAudioPlayerEndInterruptionWithFlags() bool {
	return d._AudioPlayerEndInterruptionWithFlags != nil
}

// AudioPlayerEndInterruptionWithOptions implements the PAudioPlayerDelegate interface.
func (d *AudioPlayerDelegate) AudioPlayerEndInterruptionWithOptions(player IAVAudioPlayer, flags uint) {
	if d._AudioPlayerEndInterruptionWithOptions != nil {
		d._AudioPlayerEndInterruptionWithOptions(player, flags)
	}
}

// HasAudioPlayerEndInterruptionWithOptions returns true if a handler for AudioPlayerEndInterruptionWithOptions has been set.
func (d *AudioPlayerDelegate) HasAudioPlayerEndInterruptionWithOptions() bool {
	return d._AudioPlayerEndInterruptionWithOptions != nil
}
