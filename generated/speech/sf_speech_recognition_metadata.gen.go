// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFSpeechRecognitionMetadata] class.
var (
	SFSpeechRecognitionMetadataClass     _SFSpeechRecognitionMetadataClass
	SFSpeechRecognitionMetadataClassOnce sync.Once
)

func getSFSpeechRecognitionMetadataClass() _SFSpeechRecognitionMetadataClass {
	SFSpeechRecognitionMetadataClassOnce.Do(func() {
		SFSpeechRecognitionMetadataClass = _SFSpeechRecognitionMetadataClass{objc.GetClass("SFSpeechRecognitionMetadata")}
	})
	return SFSpeechRecognitionMetadataClass
}

type _SFSpeechRecognitionMetadataClass struct {
	class objc.Class
}

// An interface definition for the [SFSpeechRecognitionMetadata] class.
type ISFSpeechRecognitionMetadata interface {
	objectivec.IObject
}

// The metadata of speech in the audio of a speech recognition request.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionMetadata
type SFSpeechRecognitionMetadata struct {
	objectivec.Object
}

// SFSpeechRecognitionMetadataFrom constructs a [SFSpeechRecognitionMetadata] from an unsafe.Pointer.
//
// The metadata of speech in the audio of a speech recognition request.
func SFSpeechRecognitionMetadataFrom(ptr unsafe.Pointer) SFSpeechRecognitionMetadata {
	return SFSpeechRecognitionMetadata{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSpeechRecognitionMetadataClass) Alloc() SFSpeechRecognitionMetadata {
	rv := objc.Send[SFSpeechRecognitionMetadata](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSpeechRecognitionMetadataClass) New() SFSpeechRecognitionMetadata {
	rv := objc.Send[SFSpeechRecognitionMetadata](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSpeechRecognitionMetadata) Init() SFSpeechRecognitionMetadata {
	rv := objc.Send[SFSpeechRecognitionMetadata](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSpeechRecognitionMetadata) Autorelease() SFSpeechRecognitionMetadata {
	rv := objc.Send[SFSpeechRecognitionMetadata](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechRecognitionMetadata creates a new SFSpeechRecognitionMetadata instance.
func NewSFSpeechRecognitionMetadata() SFSpeechRecognitionMetadata {
	return getSFSpeechRecognitionMetadataClass().New()
}


// An analysis of the transcription segment’s vocal properties.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionmetadata/voiceanalytics
func (s_ SFSpeechRecognitionMetadata) VoiceAnalytics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("voiceAnalytics"))
	return rv
}


// SetVoiceAnalytics sets the value of the voiceAnalytics property.
// An analysis of the transcription segment’s vocal properties.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionmetadata/voiceanalytics
func (s_ SFSpeechRecognitionMetadata) SetVoiceAnalytics(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoiceAnalytics:"), value)
}

// The average pause duration between words, measured in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionmetadata/averagepauseduration
func (s_ SFSpeechRecognitionMetadata) AveragePauseDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("averagePauseDuration"))
	return rv
}


// SetAveragePauseDuration sets the value of the averagePauseDuration property.
// The average pause duration between words, measured in seconds.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionmetadata/averagepauseduration
func (s_ SFSpeechRecognitionMetadata) SetAveragePauseDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAveragePauseDuration:"), value)
}

// The start timestamp of speech in the audio.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionmetadata/speechstarttimestamp
func (s_ SFSpeechRecognitionMetadata) SpeechStartTimestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("speechStartTimestamp"))
	return rv
}


// SetSpeechStartTimestamp sets the value of the speechStartTimestamp property.
// The start timestamp of speech in the audio.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionmetadata/speechstarttimestamp
func (s_ SFSpeechRecognitionMetadata) SetSpeechStartTimestamp(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpeechStartTimestamp:"), value)
}

// The number of words spoken per minute.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionmetadata/speakingrate
func (s_ SFSpeechRecognitionMetadata) SpeakingRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("speakingRate"))
	return rv
}


// SetSpeakingRate sets the value of the speakingRate property.
// The number of words spoken per minute.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionmetadata/speakingrate
func (s_ SFSpeechRecognitionMetadata) SetSpeakingRate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpeakingRate:"), value)
}

// The duration in seconds of speech in the audio.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionmetadata/speechduration
func (s_ SFSpeechRecognitionMetadata) SpeechDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("speechDuration"))
	return rv
}


// SetSpeechDuration sets the value of the speechDuration property.
// The duration in seconds of speech in the audio.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionmetadata/speechduration
func (s_ SFSpeechRecognitionMetadata) SetSpeechDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpeechDuration:"), value)
}



