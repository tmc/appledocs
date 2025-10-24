// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFVoiceAnalytics */

/* debug [class_header]: Header for SFVoiceAnalytics */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for SFVoiceAnalytics */
// An interface definition for the [SFVoiceAnalytics] class.
type ISFVoiceAnalytics interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for SFVoiceAnalytics */
	// properties:
	Jitter() ISFAcousticFeature
	Pitch() ISFAcousticFeature
	Shimmer() ISFAcousticFeature
	Voicing() ISFAcousticFeature
	IsFinal() bool
	SetIsFinal(value bool)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for SFVoiceAnalytics */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for SFVoiceAnalytics */
// Alloc allocates a new instance without initialization.
func (sc _SFVoiceAnalyticsClass) Alloc() SFVoiceAnalytics {
	rv := objc.Send[SFVoiceAnalytics](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for SFVoiceAnalytics */
// A collection of vocal analysis metrics.
//
// Use an object to access the insights. Voice analytics include the following features: Use to measure how pitch varies in audio. Use to measure how amplitude varies in audio. Use to measure the highness and lowness of the tone. Use to identify voiced regions in speech. These results are part of the object and are available when the system sends the flag.

// A collection of vocal analysis metrics.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for SFVoiceAnalytics */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for SFVoiceAnalytics */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for SFVoiceAnalytics */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for SFVoiceAnalytics */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for SFVoiceAnalytics */

// The variation in pitch in each frame of a transcription segment, expressed as a percentage of the frame’s fundamental frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFVoiceAnalytics/jitter
func (s_ SFVoiceAnalytics) Jitter() ISFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](s_.ID, objc.Sel("jitter"))
	return rv
} /* debug [instance_properties/getter]: jitter */

// The highness or lowness of the tone (fundamental frequency) in each frame of a transcription segment, expressed as a logarithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFVoiceAnalytics/pitch
func (s_ SFVoiceAnalytics) Pitch() ISFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](s_.ID, objc.Sel("pitch"))
	return rv
} /* debug [instance_properties/getter]: pitch */

// The variation in vocal volume stability (amplitude) in each frame of a transcription segment, expressed in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFVoiceAnalytics/shimmer
func (s_ SFVoiceAnalytics) Shimmer() ISFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](s_.ID, objc.Sel("shimmer"))
	return rv
} /* debug [instance_properties/getter]: shimmer */

// The likelihood of a voice in each frame of a transcription segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFVoiceAnalytics/voicing
func (s_ SFVoiceAnalytics) Voicing() ISFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](s_.ID, objc.Sel("voicing"))
	return rv
} /* debug [instance_properties/getter]: voicing */

// A Boolean value that indicates whether speech recognition is complete and whether the transcriptions are final.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionresult/isfinal
func (s_ SFVoiceAnalytics) IsFinal() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isFinal"))
	return rv
} /* debug [instance_properties/getter]: isFinal */

// A Boolean value that indicates whether speech recognition is complete and whether the transcriptions are final.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitionresult/isfinal
func (s_ SFVoiceAnalytics) SetIsFinal(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsFinal:"), value)
} /* debug [instance_properties/setter]: isFinal */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SFVoiceAnalytics */
