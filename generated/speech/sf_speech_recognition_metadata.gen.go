// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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




