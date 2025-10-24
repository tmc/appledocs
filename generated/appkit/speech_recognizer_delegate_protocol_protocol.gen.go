// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PSpeechRecognizerDelegate is the NSSpeechRecognizerDelegate protocol interface.
//
// A set of optional methods implemented by delegates of   objects.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSSpeechRecognizerDelegate
type PSpeechRecognizerDelegate interface {
	// Optional methods
	SpeechRecognizerDidRecognizeCommand(sender ISpeechRecognizer, command objc.IObject /* cross-framework: NSString */)
	HasSpeechRecognizerDidRecognizeCommand() bool
}

// SpeechRecognizerDelegate is a delegate implementation builder for the PSpeechRecognizerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SpeechRecognizerDelegate struct {
	_SpeechRecognizerDidRecognizeCommand func(sender ISpeechRecognizer, command objc.IObject /* cross-framework: NSString */)
}

// SetSpeechRecognizerDidRecognizeCommand sets the handler for the SpeechRecognizerDidRecognizeCommand delegate method.
//
// Invoked when the recognition engine has recognized the application command  .
func (d *SpeechRecognizerDelegate) SetSpeechRecognizerDidRecognizeCommand(f func(sender ISpeechRecognizer, command objc.IObject /* cross-framework: NSString */)) {
	d._SpeechRecognizerDidRecognizeCommand = f
}

// SpeechRecognizerDidRecognizeCommand implements the PSpeechRecognizerDelegate interface.
func (d *SpeechRecognizerDelegate) SpeechRecognizerDidRecognizeCommand(sender ISpeechRecognizer, command objc.IObject /* cross-framework: NSString */) {
	if d._SpeechRecognizerDidRecognizeCommand != nil {
		d._SpeechRecognizerDidRecognizeCommand(sender, command)
	}
}

// HasSpeechRecognizerDidRecognizeCommand returns true if a handler for SpeechRecognizerDidRecognizeCommand has been set.
func (d *SpeechRecognizerDelegate) HasSpeechRecognizerDidRecognizeCommand() bool {
	return d._SpeechRecognizerDidRecognizeCommand != nil
}
