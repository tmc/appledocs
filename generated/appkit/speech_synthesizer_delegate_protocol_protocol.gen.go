// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

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
	SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage(sender objc.IObject /* cross-framework: SpeechSynthesizer */, characterIndex uint, string_ objc.IObject /* cross-framework: NSString */, message objc.IObject /* cross-framework: NSString */)
	HasSpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage() bool
	SpeechSynthesizerDidEncounterSyncMessage(sender objc.IObject /* cross-framework: SpeechSynthesizer */, message objc.IObject /* cross-framework: NSString */)
	HasSpeechSynthesizerDidEncounterSyncMessage() bool
	SpeechSynthesizerDidFinishSpeaking(sender objc.IObject /* cross-framework: SpeechSynthesizer */, finishedSpeaking bool)
	HasSpeechSynthesizerDidFinishSpeaking() bool
	SpeechSynthesizerWillSpeakPhoneme(sender objc.IObject /* cross-framework: SpeechSynthesizer */, phonemeOpcode objectivec.IObject)
	HasSpeechSynthesizerWillSpeakPhoneme() bool
	SpeechSynthesizerWillSpeakWordOfString(sender objc.IObject /* cross-framework: SpeechSynthesizer */, characterRange corefoundation.Range, string_ objc.IObject /* cross-framework: NSString */)
	HasSpeechSynthesizerWillSpeakWordOfString() bool
}

// SpeechSynthesizerDelegate is a delegate implementation builder for the PSpeechSynthesizerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SpeechSynthesizerDelegate struct {
	_SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage func(sender objc.IObject /* cross-framework: SpeechSynthesizer */, characterIndex uint, string_ objc.IObject /* cross-framework: NSString */, message objc.IObject /* cross-framework: NSString */)
	_SpeechSynthesizerDidEncounterSyncMessage func(sender objc.IObject /* cross-framework: SpeechSynthesizer */, message objc.IObject /* cross-framework: NSString */)
	_SpeechSynthesizerDidFinishSpeaking func(sender objc.IObject /* cross-framework: SpeechSynthesizer */, finishedSpeaking bool)
	_SpeechSynthesizerWillSpeakPhoneme func(sender objc.IObject /* cross-framework: SpeechSynthesizer */, phonemeOpcode objectivec.IObject)
	_SpeechSynthesizerWillSpeakWordOfString func(sender objc.IObject /* cross-framework: SpeechSynthesizer */, characterRange corefoundation.Range, string_ objc.IObject /* cross-framework: NSString */)
}

// SetSpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage sets the handler for the SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage delegate method.
//
// Sent to the delegate when a speech synthesizer encounters an error in text being synthesized.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage(f func(sender objc.IObject /* cross-framework: SpeechSynthesizer */, characterIndex uint, string_ objc.IObject /* cross-framework: NSString */, message objc.IObject /* cross-framework: NSString */)) {
	d._SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage = f
}

// SetSpeechSynthesizerDidEncounterSyncMessage sets the handler for the SpeechSynthesizerDidEncounterSyncMessage delegate method.
//
// Sent to the delegate when a speech synthesizer encounters a synchronization error.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerDidEncounterSyncMessage(f func(sender objc.IObject /* cross-framework: SpeechSynthesizer */, message objc.IObject /* cross-framework: NSString */)) {
	d._SpeechSynthesizerDidEncounterSyncMessage = f
}

// SetSpeechSynthesizerDidFinishSpeaking sets the handler for the SpeechSynthesizerDidFinishSpeaking delegate method.
//
// Sent when an   object finishes speaking through the sound output device.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerDidFinishSpeaking(f func(sender objc.IObject /* cross-framework: SpeechSynthesizer */, finishedSpeaking bool)) {
	d._SpeechSynthesizerDidFinishSpeaking = f
}

// SetSpeechSynthesizerWillSpeakPhoneme sets the handler for the SpeechSynthesizerWillSpeakPhoneme delegate method.
//
// Sent just before a synthesized phoneme is spoken through the sound output device.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerWillSpeakPhoneme(f func(sender objc.IObject /* cross-framework: SpeechSynthesizer */, phonemeOpcode objectivec.IObject)) {
	d._SpeechSynthesizerWillSpeakPhoneme = f
}

// SetSpeechSynthesizerWillSpeakWordOfString sets the handler for the SpeechSynthesizerWillSpeakWordOfString delegate method.
//
// Sent just before a synthesized word is spoken through the sound output device.
func (d *SpeechSynthesizerDelegate) SetSpeechSynthesizerWillSpeakWordOfString(f func(sender objc.IObject /* cross-framework: SpeechSynthesizer */, characterRange corefoundation.Range, string_ objc.IObject /* cross-framework: NSString */)) {
	d._SpeechSynthesizerWillSpeakWordOfString = f
}

// SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage(sender objc.IObject /* cross-framework: SpeechSynthesizer */, characterIndex uint, string_ objc.IObject /* cross-framework: NSString */, message objc.IObject /* cross-framework: NSString */) {
	if d._SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage != nil {
		d._SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage(sender, characterIndex, string_, message)
	}
}

// HasSpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage returns true if a handler for SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage() bool {
	return d._SpeechSynthesizerDidEncounterErrorAtIndexOfStringMessage != nil
}

// SpeechSynthesizerDidEncounterSyncMessage implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerDidEncounterSyncMessage(sender objc.IObject /* cross-framework: SpeechSynthesizer */, message objc.IObject /* cross-framework: NSString */) {
	if d._SpeechSynthesizerDidEncounterSyncMessage != nil {
		d._SpeechSynthesizerDidEncounterSyncMessage(sender, message)
	}
}

// HasSpeechSynthesizerDidEncounterSyncMessage returns true if a handler for SpeechSynthesizerDidEncounterSyncMessage has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerDidEncounterSyncMessage() bool {
	return d._SpeechSynthesizerDidEncounterSyncMessage != nil
}

// SpeechSynthesizerDidFinishSpeaking implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerDidFinishSpeaking(sender objc.IObject /* cross-framework: SpeechSynthesizer */, finishedSpeaking bool) {
	if d._SpeechSynthesizerDidFinishSpeaking != nil {
		d._SpeechSynthesizerDidFinishSpeaking(sender, finishedSpeaking)
	}
}

// HasSpeechSynthesizerDidFinishSpeaking returns true if a handler for SpeechSynthesizerDidFinishSpeaking has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerDidFinishSpeaking() bool {
	return d._SpeechSynthesizerDidFinishSpeaking != nil
}

// SpeechSynthesizerWillSpeakPhoneme implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerWillSpeakPhoneme(sender objc.IObject /* cross-framework: SpeechSynthesizer */, phonemeOpcode objectivec.IObject) {
	if d._SpeechSynthesizerWillSpeakPhoneme != nil {
		d._SpeechSynthesizerWillSpeakPhoneme(sender, phonemeOpcode)
	}
}

// HasSpeechSynthesizerWillSpeakPhoneme returns true if a handler for SpeechSynthesizerWillSpeakPhoneme has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerWillSpeakPhoneme() bool {
	return d._SpeechSynthesizerWillSpeakPhoneme != nil
}

// SpeechSynthesizerWillSpeakWordOfString implements the PSpeechSynthesizerDelegate interface.
func (d *SpeechSynthesizerDelegate) SpeechSynthesizerWillSpeakWordOfString(sender objc.IObject /* cross-framework: SpeechSynthesizer */, characterRange corefoundation.Range, string_ objc.IObject /* cross-framework: NSString */) {
	if d._SpeechSynthesizerWillSpeakWordOfString != nil {
		d._SpeechSynthesizerWillSpeakWordOfString(sender, characterRange, string_)
	}
}

// HasSpeechSynthesizerWillSpeakWordOfString returns true if a handler for SpeechSynthesizerWillSpeakWordOfString has been set.
func (d *SpeechSynthesizerDelegate) HasSpeechSynthesizerWillSpeakWordOfString() bool {
	return d._SpeechSynthesizerWillSpeakWordOfString != nil
}
