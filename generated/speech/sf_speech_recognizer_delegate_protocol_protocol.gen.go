// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PSFSpeechRecognizerDelegate is the SFSpeechRecognizerDelegate protocol interface.
//
// A protocol that you adopt in your objects to track the availability of a speech recognizer.
//
// Availability:
//   - Mac Catalyst 10.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//
// See: doc://com.apple.speech/documentation/Speech/SFSpeechRecognizerDelegate
type PSFSpeechRecognizerDelegate interface {
	// Optional methods
	SpeechRecognizerAvailabilityDidChange(speechRecognizer ISFSpeechRecognizer, available bool)
	HasSpeechRecognizerAvailabilityDidChange() bool
}

// SFSpeechRecognizerDelegate is a delegate implementation builder for the PSFSpeechRecognizerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SFSpeechRecognizerDelegate struct {
	_SpeechRecognizerAvailabilityDidChange func(speechRecognizer ISFSpeechRecognizer, available bool)
}

// SetSpeechRecognizerAvailabilityDidChange sets the handler for the SpeechRecognizerAvailabilityDidChange delegate method.
//
// Tells the delegate that the availability of its associated speech recognizer changed.
func (d *SFSpeechRecognizerDelegate) SetSpeechRecognizerAvailabilityDidChange(f func(speechRecognizer ISFSpeechRecognizer, available bool)) {
	d._SpeechRecognizerAvailabilityDidChange = f
}

// SpeechRecognizerAvailabilityDidChange implements the PSFSpeechRecognizerDelegate interface.
func (d *SFSpeechRecognizerDelegate) SpeechRecognizerAvailabilityDidChange(speechRecognizer ISFSpeechRecognizer, available bool) {
	if d._SpeechRecognizerAvailabilityDidChange != nil {
		d._SpeechRecognizerAvailabilityDidChange(speechRecognizer, available)
	}
}

// HasSpeechRecognizerAvailabilityDidChange returns true if a handler for SpeechRecognizerAvailabilityDidChange has been set.
func (d *SFSpeechRecognizerDelegate) HasSpeechRecognizerAvailabilityDidChange() bool {
	return d._SpeechRecognizerAvailabilityDidChange != nil
}
