// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFVoiceAnalytics] class.
var (
	SFVoiceAnalyticsClass     _SFVoiceAnalyticsClass
	SFVoiceAnalyticsClassOnce sync.Once
)

func getSFVoiceAnalyticsClass() _SFVoiceAnalyticsClass {
	SFVoiceAnalyticsClassOnce.Do(func() {
		SFVoiceAnalyticsClass = _SFVoiceAnalyticsClass{objc.GetClass("SFVoiceAnalytics")}
	})
	return SFVoiceAnalyticsClass
}

type _SFVoiceAnalyticsClass struct {
	class objc.Class
}

// An interface definition for the [SFVoiceAnalytics] class.
type ISFVoiceAnalytics interface {
	objectivec.IObject
}

// A collection of vocal analysis metrics.
//
// Use an object to access the insights. Voice analytics include the following features: Use to measure how pitch varies in audio. Use to measure how amplitude varies in audio. Use to measure the highness and lowness of the tone. Use to identify voiced regions in speech. These results are part of the object and are available when the system sends the flag.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFVoiceAnalytics
type SFVoiceAnalytics struct {
	objectivec.Object
}

// SFVoiceAnalyticsFrom constructs a [SFVoiceAnalytics] from an unsafe.Pointer.
//
// A collection of vocal analysis metrics.
func SFVoiceAnalyticsFrom(ptr unsafe.Pointer) SFVoiceAnalytics {
	return SFVoiceAnalytics{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFVoiceAnalyticsClass) Alloc() SFVoiceAnalytics {
	rv := objc.Send[SFVoiceAnalytics](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFVoiceAnalyticsClass) New() SFVoiceAnalytics {
	rv := objc.Send[SFVoiceAnalytics](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFVoiceAnalytics) Init() SFVoiceAnalytics {
	rv := objc.Send[SFVoiceAnalytics](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFVoiceAnalytics) Autorelease() SFVoiceAnalytics {
	rv := objc.Send[SFVoiceAnalytics](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFVoiceAnalytics creates a new SFVoiceAnalytics instance.
func NewSFVoiceAnalytics() SFVoiceAnalytics {
	return getSFVoiceAnalyticsClass().New()
}


// A Boolean value that indicates whether speech recognition is complete and whether the transcriptions are final.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionresult/isfinal
func (s_ SFVoiceAnalytics) IsFinal() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isFinal"))
	return rv
}


// SetIsFinal sets the value of the isFinal property.
// A Boolean value that indicates whether speech recognition is complete and whether the transcriptions are final.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionresult/isfinal
func (s_ SFVoiceAnalytics) SetIsFinal(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsFinal:"), value)
}

// The variation in pitch in each frame of a transcription segment, expressed as a percentage of the frame’s fundamental frequency.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfvoiceanalytics/jitter
func (s_ SFVoiceAnalytics) Jitter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("jitter"))
	return rv
}


// SetJitter sets the value of the jitter property.
// The variation in pitch in each frame of a transcription segment, expressed as a percentage of the frame’s fundamental frequency.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfvoiceanalytics/jitter
func (s_ SFVoiceAnalytics) SetJitter(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setJitter:"), value)
}

// The highness or lowness of the tone (fundamental frequency) in each frame of a transcription segment, expressed as a logarithm.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfvoiceanalytics/pitch
func (s_ SFVoiceAnalytics) Pitch() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("pitch"))
	return rv
}


// SetPitch sets the value of the pitch property.
// The highness or lowness of the tone (fundamental frequency) in each frame of a transcription segment, expressed as a logarithm.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfvoiceanalytics/pitch
func (s_ SFVoiceAnalytics) SetPitch(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPitch:"), value)
}

// The variation in vocal volume stability (amplitude) in each frame of a transcription segment, expressed in decibels.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfvoiceanalytics/shimmer
func (s_ SFVoiceAnalytics) Shimmer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("shimmer"))
	return rv
}


// SetShimmer sets the value of the shimmer property.
// The variation in vocal volume stability (amplitude) in each frame of a transcription segment, expressed in decibels.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfvoiceanalytics/shimmer
func (s_ SFVoiceAnalytics) SetShimmer(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShimmer:"), value)
}

// The likelihood of a voice in each frame of a transcription segment.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfvoiceanalytics/voicing
func (s_ SFVoiceAnalytics) Voicing() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("voicing"))
	return rv
}


// SetVoicing sets the value of the voicing property.
// The likelihood of a voice in each frame of a transcription segment.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sfvoiceanalytics/voicing
func (s_ SFVoiceAnalytics) SetVoicing(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoicing:"), value)
}




