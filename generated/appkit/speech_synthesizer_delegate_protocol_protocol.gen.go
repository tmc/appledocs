// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PSpeechSynthesizerDelegate is the NSSpeechSynthesizerDelegate protocol interface.
//
// A set of optional methods implemented by delegates of   objects.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 14.0)
//
// See: doc://com.apple.appkit/documentation/AppKit/NSSpeechSynthesizerDelegate
type PSpeechSynthesizerDelegate interface {
	// Optional methods
	SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage(sender SpeechSynthesizer /* not a class type */, characterIndex uint, string_ foundation.foundation.INSString, message foundation.foundation.INSString)
	HasSpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage() bool
	SpeechSynthesizerDidEncounterSyncMessage(sender SpeechSynthesizer /* not a class type */, message foundation.foundation.INSString)
	HasSpeechSynthesizerDidEncounterSyncMessage() bool
	SpeechSynthesizerDidFinishSpeaking(sender SpeechSynthesizer /* not a class type */, finishedSpeaking bool)
	HasSpeechSynthesizerDidFinishSpeaking() bool
	SpeechSynthesizerWillSpeakPhoneme(sender SpeechSynthesizer /* not a class type */, phonemeOpcode objectivec.IObject)
	HasSpeechSynthesizerWillSpeakPhoneme() bool
	SpeechSynthesizerWillSpeakWordOfString(sender SpeechSynthesizer /* not a class type */, characterRange foundation.Range, string_ foundation.foundation.INSString)
	HasSpeechSynthesizerWillSpeakWordOfString() bool
}

// SpeechSynthesizerDelegate is a delegate implementation builder for the PSpeechSynthesizerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SpeechSynthesizerDelegate struct {
	_SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage func(sender SpeechSynthesizer /* not a class type */, characterIndex uint, string_ foundation.foundation.INSString, message foundation.foundation.INSString)
	_SpeechSynthesizerDidEncounterSyncMessage func(sender SpeechSynthesizer /* not a class type */, message foundation.foundation.INSString)
	_SpeechSynthesizerDidFinishSpeaking func(sender SpeechSynthesizer /* not a class type */, finishedSpeaking bool)
	_SpeechSynthesizerWillSpeakPhoneme func(sender SpeechSynthesizer /* not a class type */, phonemeOpcode objectivec.IObject)
	_SpeechSynthesizerWillSpeakWordOfString func(sender SpeechSynthesizer /* not a class type */, characterRange foundation.Range, string_ foundation.foundation.INSString)
}

// SetSpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage sets the handler for the SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage delegate method.
//
// Sent to the delegate when a speech synthesizer encounters an error in text being synthesized.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage(f func(sender SpeechSynthesizer /* not a class type */, characterIndex uint, string_ foundation.foundation.INSString, message foundation.foundation.INSString)) {
	d._SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage = f
}

// SetSpeechSynthesizerDidEncounterSyncMessage sets the handler for the SpeechSynthesizerDidEncounterSyncMessage delegate method.
//
// Sent to the delegate when a speech synthesizer encounters a synchronization error.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerDidEncounterSyncMessage(f func(sender SpeechSynthesizer /* not a class type */, message foundation.foundation.INSString)) {
	d._SpeechSynthesizerDidEncounterSyncMessage = f
}

// SetSpeechSynthesizerDidFinishSpeaking sets the handler for the SpeechSynthesizerDidFinishSpeaking delegate method.
//
// Sent when an   object finishes speaking through the sound output device.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerDidFinishSpeaking(f func(sender SpeechSynthesizer /* not a class type */, finishedSpeaking bool)) {
	d._SpeechSynthesizerDidFinishSpeaking = f
}

// SetSpeechSynthesizerWillSpeakPhoneme sets the handler for the SpeechSynthesizerWillSpeakPhoneme delegate method.
//
// Sent just before a synthesized phoneme is spoken through the sound output device.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerWillSpeakPhoneme(f func(sender SpeechSynthesizer /* not a class type */, phonemeOpcode objectivec.IObject)) {
	d._SpeechSynthesizerWillSpeakPhoneme = f
}

// SetSpeechSynthesizerWillSpeakWordOfString sets the handler for the SpeechSynthesizerWillSpeakWordOfString delegate method.
//
// Sent just before a synthesized word is spoken through the sound output device.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerWillSpeakWordOfString(f func(sender SpeechSynthesizer /* not a class type */, characterRange foundation.Range, string_ foundation.foundation.INSString)) {
	d._SpeechSynthesizerWillSpeakWordOfString = f
}

// SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage(sender SpeechSynthesizer /* not a class type */, characterIndex uint, string_ foundation.foundation.INSString, message foundation.foundation.INSString) {
	if d._SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage != nil {
		d._SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage(sender, characterIndex, string_, message)
	}
}

// HasSpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage returns true if a handler for SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage() bool {
	return d._SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage != nil
}

// SpeechSynthesizerDidEncounterSyncMessage implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerDidEncounterSyncMessage(sender SpeechSynthesizer /* not a class type */, message foundation.foundation.INSString) {
	if d._SpeechSynthesizerDidEncounterSyncMessage != nil {
		d._SpeechSynthesizerDidEncounterSyncMessage(sender, message)
	}
}

// HasSpeechSynthesizerDidEncounterSyncMessage returns true if a handler for SpeechSynthesizerDidEncounterSyncMessage has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerDidEncounterSyncMessage() bool {
	return d._SpeechSynthesizerDidEncounterSyncMessage != nil
}

// SpeechSynthesizerDidFinishSpeaking implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerDidFinishSpeaking(sender SpeechSynthesizer /* not a class type */, finishedSpeaking bool) {
	if d._SpeechSynthesizerDidFinishSpeaking != nil {
		d._SpeechSynthesizerDidFinishSpeaking(sender, finishedSpeaking)
	}
}

// HasSpeechSynthesizerDidFinishSpeaking returns true if a handler for SpeechSynthesizerDidFinishSpeaking has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerDidFinishSpeaking() bool {
	return d._SpeechSynthesizerDidFinishSpeaking != nil
}

// SpeechSynthesizerWillSpeakPhoneme implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerWillSpeakPhoneme(sender SpeechSynthesizer /* not a class type */, phonemeOpcode objectivec.IObject) {
	if d._SpeechSynthesizerWillSpeakPhoneme != nil {
		d._SpeechSynthesizerWillSpeakPhoneme(sender, phonemeOpcode)
	}
}

// HasSpeechSynthesizerWillSpeakPhoneme returns true if a handler for SpeechSynthesizerWillSpeakPhoneme has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerWillSpeakPhoneme() bool {
	return d._SpeechSynthesizerWillSpeakPhoneme != nil
}

// SpeechSynthesizerWillSpeakWordOfString implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerWillSpeakWordOfString(sender SpeechSynthesizer /* not a class type */, characterRange foundation.Range, string_ foundation.foundation.INSString) {
	if d._SpeechSynthesizerWillSpeakWordOfString != nil {
		d._SpeechSynthesizerWillSpeakWordOfString(sender, characterRange, string_)
	}
}

// HasSpeechSynthesizerWillSpeakWordOfString returns true if a handler for SpeechSynthesizerWillSpeakWordOfString has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerWillSpeakWordOfString() bool {
	return d._SpeechSynthesizerWillSpeakWordOfString != nil
}

// SpeechSynthesizerDelegateObject wraps an existing Objective-C object that conforms to the PSpeechSynthesizerDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type SpeechSynthesizerDelegateObject struct {
	objectivec.Object
}

// NewSpeechSynthesizerDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSSpeechSynthesizerDelegate protocol.
func NewSpeechSynthesizerDelegateObject(obj objectivec.Object) *SpeechSynthesizerDelegateObject {
	return &SpeechSynthesizerDelegateObject{obj}
}

// Make sure SpeechSynthesizerDelegateObject implements PSpeechSynthesizerDelegate.
var _ PSpeechSynthesizerDelegate = (*SpeechSynthesizerDelegateObject)(nil)

// SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage implements the PSpeechSynthesizerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SpeechSynthesizerDelegateObject) SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage(sender SpeechSynthesizer /* not a class type */, characterIndex uint, string_ foundation.foundation.INSString, message foundation.foundation.INSString) {
	objc.Send[objc.ID](o.ID, objc.Sel("speechSynthesizer:didEncounterErrorAtIndex:ofString:message:"), sender, characterIndex, string_, message)
}

// HasSpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage returns true; this is a placeholder for optional method checks.
func (o *SpeechSynthesizerDelegateObject) HasSpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SpeechSynthesizerDidEncounterSyncMessage implements the PSpeechSynthesizerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SpeechSynthesizerDelegateObject) SpeechSynthesizerDidEncounterSyncMessage(sender SpeechSynthesizer /* not a class type */, message foundation.foundation.INSString) {
	objc.Send[objc.ID](o.ID, objc.Sel("speechSynthesizer:didEncounterSyncMessage:"), sender, message)
}

// HasSpeechSynthesizerDidEncounterSyncMessage returns true; this is a placeholder for optional method checks.
func (o *SpeechSynthesizerDelegateObject) HasSpeechSynthesizerDidEncounterSyncMessage() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SpeechSynthesizerDidFinishSpeaking implements the PSpeechSynthesizerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SpeechSynthesizerDelegateObject) SpeechSynthesizerDidFinishSpeaking(sender SpeechSynthesizer /* not a class type */, finishedSpeaking bool) {
	objc.Send[objc.ID](o.ID, objc.Sel("speechSynthesizer:didFinishSpeaking:"), sender, finishedSpeaking)
}

// HasSpeechSynthesizerDidFinishSpeaking returns true; this is a placeholder for optional method checks.
func (o *SpeechSynthesizerDelegateObject) HasSpeechSynthesizerDidFinishSpeaking() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SpeechSynthesizerWillSpeakPhoneme implements the PSpeechSynthesizerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SpeechSynthesizerDelegateObject) SpeechSynthesizerWillSpeakPhoneme(sender SpeechSynthesizer /* not a class type */, phonemeOpcode objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("speechSynthesizer:willSpeakPhoneme:"), sender, phonemeOpcode)
}

// HasSpeechSynthesizerWillSpeakPhoneme returns true; this is a placeholder for optional method checks.
func (o *SpeechSynthesizerDelegateObject) HasSpeechSynthesizerWillSpeakPhoneme() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SpeechSynthesizerWillSpeakWordOfString implements the PSpeechSynthesizerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SpeechSynthesizerDelegateObject) SpeechSynthesizerWillSpeakWordOfString(sender SpeechSynthesizer /* not a class type */, characterRange foundation.Range, string_ foundation.foundation.INSString) {
	objc.Send[objc.ID](o.ID, objc.Sel("speechSynthesizer:willSpeakWord:ofString:"), sender, characterRange, string_)
}

// HasSpeechSynthesizerWillSpeakWordOfString returns true; this is a placeholder for optional method checks.
func (o *SpeechSynthesizerDelegateObject) HasSpeechSynthesizerWillSpeakWordOfString() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
