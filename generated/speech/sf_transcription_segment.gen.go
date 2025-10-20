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




