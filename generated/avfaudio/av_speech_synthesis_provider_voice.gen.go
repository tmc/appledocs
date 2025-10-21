// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SpeechSynthesisProviderVoice] class.
var (
	SpeechSynthesisProviderVoiceClass     _SpeechSynthesisProviderVoiceClass
	SpeechSynthesisProviderVoiceClassOnce sync.Once
)

func getSpeechSynthesisProviderVoiceClass() _SpeechSynthesisProviderVoiceClass {
	SpeechSynthesisProviderVoiceClassOnce.Do(func() {
		SpeechSynthesisProviderVoiceClass = _SpeechSynthesisProviderVoiceClass{objc.GetClass("AVSpeechSynthesisProviderVoice")}
	})
	return SpeechSynthesisProviderVoiceClass
}

type _SpeechSynthesisProviderVoiceClass struct {
	class objc.Class
}

// An interface definition for the [SpeechSynthesisProviderVoice] class.
type ISpeechSynthesisProviderVoice interface {
	objectivec.IObject
}

// An object that represents a voice that an audio unit provides to its host.
//
// This is a voice that an provides to the system, distinct from . Use to access the underlying in the voice quality .
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice
type SpeechSynthesisProviderVoice struct {
	objectivec.Object
}

// SpeechSynthesisProviderVoiceFrom constructs a [SpeechSynthesisProviderVoice] from an unsafe.Pointer.
//
// An object that represents a voice that an audio unit provides to its host.
func SpeechSynthesisProviderVoiceFrom(ptr unsafe.Pointer) SpeechSynthesisProviderVoice {
	return SpeechSynthesisProviderVoice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SpeechSynthesisProviderVoiceClass) Alloc() SpeechSynthesisProviderVoice {
	rv := objc.Send[SpeechSynthesisProviderVoice](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SpeechSynthesisProviderVoiceClass) New() SpeechSynthesisProviderVoice {
	rv := objc.Send[SpeechSynthesisProviderVoice](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpeechSynthesisProviderVoice) Init() SpeechSynthesisProviderVoice {
	rv := objc.Send[SpeechSynthesisProviderVoice](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpeechSynthesisProviderVoice) Autorelease() SpeechSynthesisProviderVoice {
	rv := objc.Send[SpeechSynthesisProviderVoice](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpeechSynthesisProviderVoice creates a new SpeechSynthesisProviderVoice instance.
func NewSpeechSynthesisProviderVoice() SpeechSynthesisProviderVoice {
	return getSpeechSynthesisProviderVoiceClass().New()
}


// Updates the voices your app provides to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/updateSpeechVoices()
func (sc _SpeechSynthesisProviderVoiceClass) UpdateSpeechVoices() {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("updateSpeechVoices"))
}



