// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFTranscriptionSegment */


/* debug [class_header]: Header for SFTranscriptionSegment */
// The class instance for the [SFTranscriptionSegment] class.
var (
	SFTranscriptionSegmentClass     _SFTranscriptionSegmentClass
	SFTranscriptionSegmentClassOnce sync.Once
)

func getSFTranscriptionSegmentClass() _SFTranscriptionSegmentClass {
	SFTranscriptionSegmentClassOnce.Do(func() {
		SFTranscriptionSegmentClass = _SFTranscriptionSegmentClass{objc.GetClass("SFTranscriptionSegment")}
	})
	return SFTranscriptionSegmentClass
}

type _SFTranscriptionSegmentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFTranscriptionSegment */
// An interface definition for the [SFTranscriptionSegment] class.
type ISFTranscriptionSegment interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFTranscriptionSegment */
	// properties:
	AlternativeSubstrings() []string
	Confidence() float32
	Duration() float64
	Substring() objc.IObject /* cross-framework: NSString */
	SubstringRange() corefoundation.Range
	Timestamp() float64
	VoiceAnalytics() ISFVoiceAnalytics
	FormattedString() objc.IObject /* cross-framework: NSString */
	SetFormattedString(value objc.IObject /* cross-framework: NSString */)
	Segments() ISFTranscriptionSegment
	SetSegments(value ISFTranscriptionSegment)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFTranscriptionSegment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFTranscriptionSegment */
// Alloc allocates a new instance without initialization.
func (sc _SFTranscriptionSegmentClass) Alloc() SFTranscriptionSegment {
	rv := objc.Send[SFTranscriptionSegment](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFTranscriptionSegmentClass) New() SFTranscriptionSegment {
	rv := objc.Send[SFTranscriptionSegment](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFTranscriptionSegment) Init() SFTranscriptionSegment {
	rv := objc.Send[SFTranscriptionSegment](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFTranscriptionSegment) Autorelease() SFTranscriptionSegment {
	rv := objc.Send[SFTranscriptionSegment](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFTranscriptionSegment creates a new SFTranscriptionSegment instance.
func NewSFTranscriptionSegment() SFTranscriptionSegment {
	return getSFTranscriptionSegmentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFTranscriptionSegment */
// A discrete part of an entire transcription, as identified by the speech recognizer.
//
// Use to get details about a part of an overall . An represents an utterance, which is a vocalized word or group of words that represent a single meaning to the speech recognizer ( ). You don’t create transcription object segments directly. Instead, you access them from a transcription’s property. A transcription segment includes the following information: The text of the utterance, plus any alternative interpretations of the spoken word. The character range of the segment within the of its parent . A value, indicating how likely it is that the specified string matches the audible speech. A and value, indicating the position of the segment within the provided audio stream.


// A discrete part of an entire transcription, as identified by the speech recognizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment
type SFTranscriptionSegment struct {
	objectivec.Object
}

// SFTranscriptionSegmentFrom constructs a [SFTranscriptionSegment] from an unsafe.Pointer.
//
// A discrete part of an entire transcription, as identified by the speech recognizer.
func SFTranscriptionSegmentFrom(ptr unsafe.Pointer) SFTranscriptionSegment {
	return SFTranscriptionSegment{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFTranscriptionSegment *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFTranscriptionSegment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFTranscriptionSegment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFTranscriptionSegment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFTranscriptionSegment */

// An array of alternate interpretations of the utterance in the transcription segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment/alternativeSubstrings
func (s_ SFTranscriptionSegment) AlternativeSubstrings() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("alternativeSubstrings"))
	return rv
}/* debug [instance_properties/getter]: alternativeSubstrings */


// The level of confidence the speech recognizer has in its recognition of the speech transcribed for the segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment/confidence
func (s_ SFTranscriptionSegment) Confidence() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("confidence"))
	return rv
}/* debug [instance_properties/getter]: confidence */


// The number of seconds it took for the user to speak the utterance represented by the segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment/duration
func (s_ SFTranscriptionSegment) Duration() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The string representation of the utterance in the transcription segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment/substring
func (s_ SFTranscriptionSegment) Substring() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("substring"))
	return rv
}/* debug [instance_properties/getter]: substring */


// The range information for the transcription segment’s substring, relative to the overall transcription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment/substringRange
func (s_ SFTranscriptionSegment) SubstringRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](s_.ID, objc.Sel("substringRange"))
	return rv
}/* debug [instance_properties/getter]: substringRange */


// The start time of the segment in the processed audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment/timestamp
func (s_ SFTranscriptionSegment) Timestamp() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("timestamp"))
	return rv
}/* debug [instance_properties/getter]: timestamp */


// An analysis of the transcription segment’s vocal properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFTranscriptionSegment/voiceAnalytics
func (s_ SFTranscriptionSegment) VoiceAnalytics() ISFVoiceAnalytics {
	rv := objc.Send[SFVoiceAnalytics](s_.ID, objc.Sel("voiceAnalytics"))
	return rv
}/* debug [instance_properties/getter]: voiceAnalytics */


// The entire transcription of utterances, formatted into a single, user-displayable string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscription/formattedstring
func (s_ SFTranscriptionSegment) FormattedString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("formattedString"))
	return rv
}/* debug [instance_properties/getter]: formattedString */


// The entire transcription of utterances, formatted into a single, user-displayable string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscription/formattedstring
func (s_ SFTranscriptionSegment) SetFormattedString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFormattedString:"), value)
}/* debug [instance_properties/setter]: formattedString */


// An array of transcription segments that represent the parts of the transcription, as identified by the speech recognizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscription/segments
func (s_ SFTranscriptionSegment) Segments() ISFTranscriptionSegment {
	rv := objc.Send[SFTranscriptionSegment](s_.ID, objc.Sel("segments"))
	return rv
}/* debug [instance_properties/getter]: segments */


// An array of transcription segments that represent the parts of the transcription, as identified by the speech recognizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscription/segments
func (s_ SFTranscriptionSegment) SetSegments(value ISFTranscriptionSegment) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegments:"), value)
}/* debug [instance_properties/setter]: segments */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFTranscriptionSegment */



