// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PSpeechSynthesizerDelegate is the AVSpeechSynthesizerDelegate protocol interface.
//
// A delegate protocol that contains optional methods you can implement to respond to events that occur during speech synthesis.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//   - watchOS +
//
// See: doc://com.apple.avfaudio/documentation/AVFAudio/AVSpeechSynthesizerDelegate
type PSpeechSynthesizerDelegate interface {
	// Optional methods
	SpeechSynthesizerDidCancelSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)
	HasSpeechSynthesizerDidCancelSpeechUtterance() bool
	SpeechSynthesizerDidContinueSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)
	HasSpeechSynthesizerDidContinueSpeechUtterance() bool
	SpeechSynthesizerDidFinishSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)
	HasSpeechSynthesizerDidFinishSpeechUtterance() bool
	SpeechSynthesizerDidPauseSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)
	HasSpeechSynthesizerDidPauseSpeechUtterance() bool
	SpeechSynthesizerDidStartSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)
	HasSpeechSynthesizerDidStartSpeechUtterance() bool
	SpeechSynthesizerWillSpeakMarkerUtterance(synthesizer IAVSpeechSynthesizer, marker IAVSpeechSynthesisMarker, utterance IAVSpeechUtterance)
	HasSpeechSynthesizerWillSpeakMarkerUtterance() bool
	SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance(synthesizer IAVSpeechSynthesizer, characterRange corefoundation.Range, utterance IAVSpeechUtterance)
	HasSpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance() bool
}

// SpeechSynthesizerDelegate is a delegate implementation builder for the PSpeechSynthesizerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SpeechSynthesizerDelegate struct {
	_SpeechSynthesizerDidCancelSpeechUtterance func(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)
	_SpeechSynthesizerDidContinueSpeechUtterance func(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)
	_SpeechSynthesizerDidFinishSpeechUtterance func(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)
	_SpeechSynthesizerDidPauseSpeechUtterance func(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)
	_SpeechSynthesizerDidStartSpeechUtterance func(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)
	_SpeechSynthesizerWillSpeakMarkerUtterance func(synthesizer IAVSpeechSynthesizer, marker IAVSpeechSynthesisMarker, utterance IAVSpeechUtterance)
	_SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance func(synthesizer IAVSpeechSynthesizer, characterRange corefoundation.Range, utterance IAVSpeechUtterance)
}

// SetSpeechSynthesizerDidCancelSpeechUtterance sets the handler for the SpeechSynthesizerDidCancelSpeechUtterance delegate method.
//
// Tells the delegate when the synthesizer cancels speaking an utterance.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerDidCancelSpeechUtterance(f func(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)) {
	d._SpeechSynthesizerDidCancelSpeechUtterance = f
}

// SetSpeechSynthesizerDidContinueSpeechUtterance sets the handler for the SpeechSynthesizerDidContinueSpeechUtterance delegate method.
//
// Tells the delegate when the synthesizer resumes speaking an utterance after pausing.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerDidContinueSpeechUtterance(f func(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)) {
	d._SpeechSynthesizerDidContinueSpeechUtterance = f
}

// SetSpeechSynthesizerDidFinishSpeechUtterance sets the handler for the SpeechSynthesizerDidFinishSpeechUtterance delegate method.
//
// Tells the delegate when the synthesizer finishes speaking an utterance.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerDidFinishSpeechUtterance(f func(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)) {
	d._SpeechSynthesizerDidFinishSpeechUtterance = f
}

// SetSpeechSynthesizerDidPauseSpeechUtterance sets the handler for the SpeechSynthesizerDidPauseSpeechUtterance delegate method.
//
// Tells the delegate when the synthesizer pauses while speaking an utterance.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerDidPauseSpeechUtterance(f func(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)) {
	d._SpeechSynthesizerDidPauseSpeechUtterance = f
}

// SetSpeechSynthesizerDidStartSpeechUtterance sets the handler for the SpeechSynthesizerDidStartSpeechUtterance delegate method.
//
// Tells the delegate when the synthesizer begins speaking an utterance.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerDidStartSpeechUtterance(f func(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance)) {
	d._SpeechSynthesizerDidStartSpeechUtterance = f
}

// SetSpeechSynthesizerWillSpeakMarkerUtterance sets the handler for the SpeechSynthesizerWillSpeakMarkerUtterance delegate method.
//
// Tells the delegate when the synthesizer is about to speak a marker of an utterance’s text.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerWillSpeakMarkerUtterance(f func(synthesizer IAVSpeechSynthesizer, marker IAVSpeechSynthesisMarker, utterance IAVSpeechUtterance)) {
	d._SpeechSynthesizerWillSpeakMarkerUtterance = f
}

// SetSpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance sets the handler for the SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance delegate method.
//
// Tells the delegate when the synthesizer is about to speak a portion of an utterance’s text.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance(f func(synthesizer IAVSpeechSynthesizer, characterRange corefoundation.Range, utterance IAVSpeechUtterance)) {
	d._SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance = f
}

// SpeechSynthesizerDidCancelSpeechUtterance implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerDidCancelSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance) {
	if d._SpeechSynthesizerDidCancelSpeechUtterance != nil {
		d._SpeechSynthesizerDidCancelSpeechUtterance(synthesizer, utterance)
	}
}

// HasSpeechSynthesizerDidCancelSpeechUtterance returns true if a handler for SpeechSynthesizerDidCancelSpeechUtterance has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerDidCancelSpeechUtterance() bool {
	return d._SpeechSynthesizerDidCancelSpeechUtterance != nil
}

// SpeechSynthesizerDidContinueSpeechUtterance implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerDidContinueSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance) {
	if d._SpeechSynthesizerDidContinueSpeechUtterance != nil {
		d._SpeechSynthesizerDidContinueSpeechUtterance(synthesizer, utterance)
	}
}

// HasSpeechSynthesizerDidContinueSpeechUtterance returns true if a handler for SpeechSynthesizerDidContinueSpeechUtterance has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerDidContinueSpeechUtterance() bool {
	return d._SpeechSynthesizerDidContinueSpeechUtterance != nil
}

// SpeechSynthesizerDidFinishSpeechUtterance implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerDidFinishSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance) {
	if d._SpeechSynthesizerDidFinishSpeechUtterance != nil {
		d._SpeechSynthesizerDidFinishSpeechUtterance(synthesizer, utterance)
	}
}

// HasSpeechSynthesizerDidFinishSpeechUtterance returns true if a handler for SpeechSynthesizerDidFinishSpeechUtterance has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerDidFinishSpeechUtterance() bool {
	return d._SpeechSynthesizerDidFinishSpeechUtterance != nil
}

// SpeechSynthesizerDidPauseSpeechUtterance implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerDidPauseSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance) {
	if d._SpeechSynthesizerDidPauseSpeechUtterance != nil {
		d._SpeechSynthesizerDidPauseSpeechUtterance(synthesizer, utterance)
	}
}

// HasSpeechSynthesizerDidPauseSpeechUtterance returns true if a handler for SpeechSynthesizerDidPauseSpeechUtterance has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerDidPauseSpeechUtterance() bool {
	return d._SpeechSynthesizerDidPauseSpeechUtterance != nil
}

// SpeechSynthesizerDidStartSpeechUtterance implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerDidStartSpeechUtterance(synthesizer IAVSpeechSynthesizer, utterance IAVSpeechUtterance) {
	if d._SpeechSynthesizerDidStartSpeechUtterance != nil {
		d._SpeechSynthesizerDidStartSpeechUtterance(synthesizer, utterance)
	}
}

// HasSpeechSynthesizerDidStartSpeechUtterance returns true if a handler for SpeechSynthesizerDidStartSpeechUtterance has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerDidStartSpeechUtterance() bool {
	return d._SpeechSynthesizerDidStartSpeechUtterance != nil
}

// SpeechSynthesizerWillSpeakMarkerUtterance implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerWillSpeakMarkerUtterance(synthesizer IAVSpeechSynthesizer, marker IAVSpeechSynthesisMarker, utterance IAVSpeechUtterance) {
	if d._SpeechSynthesizerWillSpeakMarkerUtterance != nil {
		d._SpeechSynthesizerWillSpeakMarkerUtterance(synthesizer, marker, utterance)
	}
}

// HasSpeechSynthesizerWillSpeakMarkerUtterance returns true if a handler for SpeechSynthesizerWillSpeakMarkerUtterance has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerWillSpeakMarkerUtterance() bool {
	return d._SpeechSynthesizerWillSpeakMarkerUtterance != nil
}

// SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance(synthesizer IAVSpeechSynthesizer, characterRange corefoundation.Range, utterance IAVSpeechUtterance) {
	if d._SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance != nil {
		d._SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance(synthesizer, characterRange, utterance)
	}
}

// HasSpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance returns true if a handler for SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance() bool {
	return d._SpeechSynthesizerWillSpeakRangeOfSpeechStringUtterance != nil
}
