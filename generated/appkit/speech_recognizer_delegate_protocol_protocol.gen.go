// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	SpeechRecognizerDidRecognizeCommand(sender ISpeechRecognizer, command foundation.foundation.INSString)
	HasSpeechRecognizerDidRecognizeCommand() bool
}

// SpeechRecognizerDelegate is a delegate implementation builder for the PSpeechRecognizerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SpeechRecognizerDelegate struct {
	_SpeechRecognizerDidRecognizeCommand func(sender ISpeechRecognizer, command foundation.foundation.INSString)
}

// SetSpeechRecognizerDidRecognizeCommand sets the handler for the SpeechRecognizerDidRecognizeCommand delegate method.
//
// Invoked when the recognition engine has recognized the application command  .
func (d *SpeechRecognizerDelegate) SetSpeechRecognizerDidRecognizeCommand(f func(sender ISpeechRecognizer, command foundation.foundation.INSString)) {
	d._SpeechRecognizerDidRecognizeCommand = f
}

// SpeechRecognizerDidRecognizeCommand implements the PSpeechRecognizerDelegate interface.
func (d *SpeechRecognizerDelegate) SpeechRecognizerDidRecognizeCommand(sender ISpeechRecognizer, command foundation.foundation.INSString) {
	if d._SpeechRecognizerDidRecognizeCommand != nil {
		d._SpeechRecognizerDidRecognizeCommand(sender, command)
	}
}

// HasSpeechRecognizerDidRecognizeCommand returns true if a handler for SpeechRecognizerDidRecognizeCommand has been set.
func (d *SpeechRecognizerDelegate) HasSpeechRecognizerDidRecognizeCommand() bool {
	return d._SpeechRecognizerDidRecognizeCommand != nil
}

// SpeechRecognizerDelegateObject wraps an existing Objective-C object that conforms to the PSpeechRecognizerDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type SpeechRecognizerDelegateObject struct {
	objectivec.Object
}

// NewSpeechRecognizerDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSSpeechRecognizerDelegate protocol.
func NewSpeechRecognizerDelegateObject(obj objectivec.Object) *SpeechRecognizerDelegateObject {
	return &SpeechRecognizerDelegateObject{obj}
}

// Make sure SpeechRecognizerDelegateObject implements PSpeechRecognizerDelegate.
var _ PSpeechRecognizerDelegate = (*SpeechRecognizerDelegateObject)(nil)

// SpeechRecognizerDidRecognizeCommand implements the PSpeechRecognizerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SpeechRecognizerDelegateObject) SpeechRecognizerDidRecognizeCommand(sender ISpeechRecognizer, command foundation.foundation.INSString) {
	objc.Send[objc.ID](o.ID, objc.Sel("speechRecognizer:didRecognizeCommand:"), sender, command)
}

// HasSpeechRecognizerDidRecognizeCommand returns true; this is a placeholder for optional method checks.
func (o *SpeechRecognizerDelegateObject) HasSpeechRecognizerDidRecognizeCommand() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
