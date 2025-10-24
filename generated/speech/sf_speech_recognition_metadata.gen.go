// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSpeechRecognitionMetadata */

/* debug [class_header]: Header for SFSpeechRecognitionMetadata */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for SFSpeechRecognitionMetadata */
// An interface definition for the [SFSpeechRecognitionMetadata] class.
type ISFSpeechRecognitionMetadata interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for SFSpeechRecognitionMetadata */
	// properties:
	AveragePauseDuration() float64
	SpeakingRate() float64
	SpeechDuration() float64
	SpeechStartTimestamp() float64
	VoiceAnalytics() ISFVoiceAnalytics
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for SFSpeechRecognitionMetadata */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for SFSpeechRecognitionMetadata */
// Alloc allocates a new instance without initialization.
func (sc _SFSpeechRecognitionMetadataClass) Alloc() SFSpeechRecognitionMetadata {
	rv := objc.Send[SFSpeechRecognitionMetadata](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for SFSpeechRecognitionMetadata */
// The metadata of speech in the audio of a speech recognition request.

// The metadata of speech in the audio of a speech recognition request.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for SFSpeechRecognitionMetadata */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for SFSpeechRecognitionMetadata */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for SFSpeechRecognitionMetadata */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for SFSpeechRecognitionMetadata */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for SFSpeechRecognitionMetadata */

// The average pause duration between words, measured in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionMetadata/averagePauseDuration
func (s_ SFSpeechRecognitionMetadata) AveragePauseDuration() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("averagePauseDuration"))
	return rv
} /* debug [instance_properties/getter]: averagePauseDuration */

// The number of words spoken per minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionMetadata/speakingRate
func (s_ SFSpeechRecognitionMetadata) SpeakingRate() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("speakingRate"))
	return rv
} /* debug [instance_properties/getter]: speakingRate */

// The duration in seconds of speech in the audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionMetadata/speechDuration
func (s_ SFSpeechRecognitionMetadata) SpeechDuration() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("speechDuration"))
	return rv
} /* debug [instance_properties/getter]: speechDuration */

// The start timestamp of speech in the audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionMetadata/speechStartTimestamp
func (s_ SFSpeechRecognitionMetadata) SpeechStartTimestamp() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("speechStartTimestamp"))
	return rv
} /* debug [instance_properties/getter]: speechStartTimestamp */

// An analysis of the transcription segment’s vocal properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionMetadata/voiceAnalytics
func (s_ SFSpeechRecognitionMetadata) VoiceAnalytics() ISFVoiceAnalytics {
	rv := objc.Send[SFVoiceAnalytics](s_.ID, objc.Sel("voiceAnalytics"))
	return rv
} /* debug [instance_properties/getter]: voiceAnalytics */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SFSpeechRecognitionMetadata */
