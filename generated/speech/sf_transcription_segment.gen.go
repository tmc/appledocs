// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A discrete part of an entire transcription, as identified by the speech recognizer.
//
// Use to get details about a part of an overall . An represents an utterance, which is a vocalized word or group of words that represent a single meaning to the speech recognizer ( ). You don’t create transcription object segments directly. Instead, you access them from a transcription’s property. A transcription segment includes the following information: The text of the utterance, plus any alternative interpretations of the spoken word. The character range of the segment within the of its parent . A value, indicating how likely it is that the specified string matches the audible speech. A and value, indicating the position of the segment within the provided audio stream.
//
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


// The number of seconds it took for the user to speak the utterance represented by the segment.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/duration
func (s_ SFTranscriptionSegment) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
// The number of seconds it took for the user to speak the utterance represented by the segment.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/duration
func (s_ SFTranscriptionSegment) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDuration:"), value)
}

// An array of alternate interpretations of the utterance in the transcription segment.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/alternativesubstrings
func (s_ SFTranscriptionSegment) AlternativeSubstrings() string {
	rv := objc.Send[string](s_.ID, objc.Sel("alternativeSubstrings"))
	return rv
}


// SetAlternativeSubstrings sets the value of the alternativeSubstrings property.
// An array of alternate interpretations of the utterance in the transcription segment.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/alternativesubstrings
func (s_ SFTranscriptionSegment) SetAlternativeSubstrings(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlternativeSubstrings:"), objc.String(value))
}

// The start time of the segment in the processed audio stream.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/timestamp
func (s_ SFTranscriptionSegment) Timestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("timestamp"))
	return rv
}


// SetTimestamp sets the value of the timestamp property.
// The start time of the segment in the processed audio stream.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/timestamp
func (s_ SFTranscriptionSegment) SetTimestamp(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTimestamp:"), value)
}

// An analysis of the transcription segment’s vocal properties.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/voiceanalytics
func (s_ SFTranscriptionSegment) VoiceAnalytics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("voiceAnalytics"))
	return rv
}


// SetVoiceAnalytics sets the value of the voiceAnalytics property.
// An analysis of the transcription segment’s vocal properties.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/voiceanalytics
func (s_ SFTranscriptionSegment) SetVoiceAnalytics(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVoiceAnalytics:"), value)
}

// The string representation of the utterance in the transcription segment.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/substring
func (s_ SFTranscriptionSegment) Substring() string {
	rv := objc.Send[string](s_.ID, objc.Sel("substring"))
	return rv
}


// SetSubstring sets the value of the substring property.
// The string representation of the utterance in the transcription segment.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/substring
func (s_ SFTranscriptionSegment) SetSubstring(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSubstring:"), objc.String(value))
}

// The level of confidence the speech recognizer has in its recognition of the speech transcribed for the segment.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/confidence
func (s_ SFTranscriptionSegment) Confidence() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("confidence"))
	return rv
}


// SetConfidence sets the value of the confidence property.
// The level of confidence the speech recognizer has in its recognition of the speech transcribed for the segment.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/confidence
func (s_ SFTranscriptionSegment) SetConfidence(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setConfidence:"), value)
}

// The range information for the transcription segment’s substring, relative to the overall transcription.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/substringrange
func (s_ SFTranscriptionSegment) SubstringRange() Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("substringRange"))
	return rv
}


// SetSubstringRange sets the value of the substringRange property.
// The range information for the transcription segment’s substring, relative to the overall transcription.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscriptionsegment/substringrange
func (s_ SFTranscriptionSegment) SetSubstringRange(value Range) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSubstringRange:"), value)
}

// An array of transcription segments that represent the parts of the transcription, as identified by the speech recognizer.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscription/segments
func (s_ SFTranscriptionSegment) Segments() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("segments"))
	return rv
}


// SetSegments sets the value of the segments property.
// An array of transcription segments that represent the parts of the transcription, as identified by the speech recognizer.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscription/segments
func (s_ SFTranscriptionSegment) SetSegments(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegments:"), value)
}

// The entire transcription of utterances, formatted into a single, user-displayable string.
//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscription/formattedstring
func (s_ SFTranscriptionSegment) FormattedString() string {
	rv := objc.Send[string](s_.ID, objc.Sel("formattedString"))
	return rv
}


// SetFormattedString sets the value of the formattedString property.
// The entire transcription of utterances, formatted into a single, user-displayable string.

//
// [Full Topic]: https://developer.apple.com/documentation/speech/sftranscription/formattedstring
func (s_ SFTranscriptionSegment) SetFormattedString(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFormattedString:"), objc.String(value))
}



