// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SFTranscription] class.
var (
	SFTranscriptionClass     _SFTranscriptionClass
	SFTranscriptionClassOnce sync.Once
)

func getSFTranscriptionClass() _SFTranscriptionClass {
	SFTranscriptionClassOnce.Do(func() {
		SFTranscriptionClass = _SFTranscriptionClass{objc.GetClass("SFTranscription")}
	})
	return SFTranscriptionClass
}

type _SFTranscriptionClass struct {
	class objc.Class
}

// An interface definition for the [SFTranscription] class.
type ISFTranscription interface {
	objectivec.IObject
}

// A textual representation of the specified speech in its entirety, as recognized by the speech recognizer.
//
// Use to obtain all the recognized utterances from your audio content. An is a vocalized word or group of words that represent a single meaning to the speech recognizer ( ). Use the property to retrieve the entire transcription of utterances, or use the property to retrieve an individual utterance ( ). You don’t create an directly. Instead, you retrieve it from an instance. The speech recognizer sends a speech recognition result to your app in one of two ways, depending on how your app started a speech recognition task. You can start a speech recognition task by using the speech recognizer’s method. When the task is complete, the speech recognizer sends an instance to your closure. Alternatively, you can use the speech recognizer’s method to start a speech recognition task. When the task is complete, the speech recognizer uses your to send an by using the delegate’s method. An represents only a potential version of the speech. It might not be an accurate representation of the utterances.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFTranscription
type SFTranscription struct {
	objectivec.Object
}

// SFTranscriptionFrom constructs a [SFTranscription] from an unsafe.Pointer.
//
// A textual representation of the specified speech in its entirety, as recognized by the speech recognizer.
func SFTranscriptionFrom(ptr unsafe.Pointer) SFTranscription {
	return SFTranscription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFTranscriptionClass) Alloc() SFTranscription {
	rv := objc.Send[SFTranscription](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFTranscriptionClass) New() SFTranscription {
	rv := objc.Send[SFTranscription](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFTranscription) Init() SFTranscription {
	rv := objc.Send[SFTranscription](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFTranscription) Autorelease() SFTranscription {
	rv := objc.Send[SFTranscription](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFTranscription creates a new SFTranscription instance.
func NewSFTranscription() SFTranscription {
	return getSFTranscriptionClass().New()
}




