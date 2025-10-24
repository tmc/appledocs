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

// An interface definition for the [SFTranscriptionSegment] class.
type ISFTranscriptionSegment interface {
	objectivec.IObject
	// properties:
	FormattedString() objc.IObject /* cross-framework: NSString */
	SetFormattedString(value objc.IObject /* cross-framework: NSString */)
	Segments() ISFTranscriptionSegment
	SetSegments(value ISFTranscriptionSegment)
	AlternativeSubstrings() objc.IObject /* cross-framework: NSString */
	SetAlternativeSubstrings(value objc.IObject /* cross-framework: NSString */)
	Confidence() float32
	SetConfidence(value float32)
	Duration() float64
	SetDuration(value float64)
	Substring() objc.IObject /* cross-framework: NSString */
	SetSubstring(value objc.IObject /* cross-framework: NSString */)
	SubstringRange() objc.IObject /* cross-framework: Range */
	SetSubstringRange(value objc.IObject /* cross-framework: Range */)
	Timestamp() float64
	SetTimestamp(value float64)
	VoiceAnalytics() ISFVoiceAnalytics
	SetVoiceAnalytics(value ISFVoiceAnalytics)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (sc _SFTranscriptionSegmentClass) Alloc() SFTranscriptionSegment {
	rv := objc.Send[SFTranscriptionSegment](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The entire transcription of utterances, formatted into a single, user-displayable string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscription/formattedstring
func (s_ SFTranscriptionSegment) FormattedString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("formattedString"))
	return rv
}


// The entire transcription of utterances, formatted into a single, user-displayable string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscription/formattedstring
func (s_ SFTranscriptionSegment) SetFormattedString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFormattedString:"), value)
}


// An array of transcription segments that represent the parts of the transcription, as identified by the speech recognizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscription/segments
func (s_ SFTranscriptionSegment) Segments() ISFTranscriptionSegment {
	rv := objc.Send[SFTranscriptionSegment](s_.ID, objc.Sel("segments"))
	return rv
}


// An array of transcription segments that represent the parts of the transcription, as identified by the speech recognizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscription/segments
func (s_ SFTranscriptionSegment) SetSegments(value ISFTranscriptionSegment) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegments:"), value)
}


// An array of alternate interpretations of the utterance in the transcription segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/alternativesubstrings
func (s_ SFTranscriptionSegment) AlternativeSubstrings() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("alternativeSubstrings"))
	return rv
}


// An array of alternate interpretations of the utterance in the transcription segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/alternativesubstrings
func (s_ SFTranscriptionSegment) SetAlternativeSubstrings(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlternativeSubstrings:"), value)
}


// The level of confidence the speech recognizer has in its recognition of the speech transcribed for the segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/confidence
func (s_ SFTranscriptionSegment) Confidence() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("confidence"))
	return rv
}


// The level of confidence the speech recognizer has in its recognition of the speech transcribed for the segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/confidence
func (s_ SFTranscriptionSegment) SetConfidence(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setConfidence:"), value)
}


// The number of seconds it took for the user to speak the utterance represented by the segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/duration
func (s_ SFTranscriptionSegment) Duration() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("duration"))
	return rv
}


// The number of seconds it took for the user to speak the utterance represented by the segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/duration
func (s_ SFTranscriptionSegment) SetDuration(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDuration:"), value)
}


// The string representation of the utterance in the transcription segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/substring
func (s_ SFTranscriptionSegment) Substring() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("substring"))
	return rv
}


// The string representation of the utterance in the transcription segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/substring
func (s_ SFTranscriptionSegment) SetSubstring(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSubstring:"), value)
}


// The range information for the transcription segment’s substring, relative to the overall transcription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/substringrange
func (s_ SFTranscriptionSegment) SubstringRange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](s_.ID, objc.Sel("substringRange"))
	return rv
}


// The range information for the transcription segment’s substring, relative to the overall transcription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/substringrange
func (s_ SFTranscriptionSegment) SetSubstringRange(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSubstringRange:"), value)
}


// The start time of the segment in the processed audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/timestamp
func (s_ SFTranscriptionSegment) Timestamp() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("timestamp"))
	return rv
}


// The start time of the segment in the processed audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/timestamp
func (s_ SFTranscriptionSegment) SetTimestamp(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTimestamp:"), value)
}


// An analysis of the transcription segment’s vocal properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/voiceanalytics
func (s_ SFTranscriptionSegment) VoiceAnalytics() ISFVoiceAnalytics {
	rv := objc.Send[SFVoiceAnalytics](s_.ID, objc.Sel("voiceAnalytics"))
	return rv
}


// An analysis of the transcription segment’s vocal properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/voiceanalytics
func (s_ SFTranscriptionSegment) SetVoiceAnalytics(value ISFVoiceAnalytics) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoiceAnalytics:"), value)
}



