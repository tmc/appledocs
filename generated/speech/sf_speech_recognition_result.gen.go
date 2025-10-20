// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFSpeechRecognitionResult] class.
var (
	SFSpeechRecognitionResultClass     _SFSpeechRecognitionResultClass
	SFSpeechRecognitionResultClassOnce sync.Once
)

func getSFSpeechRecognitionResultClass() _SFSpeechRecognitionResultClass {
	SFSpeechRecognitionResultClassOnce.Do(func() {
		SFSpeechRecognitionResultClass = _SFSpeechRecognitionResultClass{objc.GetClass("SFSpeechRecognitionResult")}
	})
	return SFSpeechRecognitionResultClass
}

type _SFSpeechRecognitionResultClass struct {
	class objc.Class
}

// An interface definition for the [SFSpeechRecognitionResult] class.
type ISFSpeechRecognitionResult interface {
	objectivec.IObject
}

// An object that contains the partial or final results of a speech recognition request.
//
// Use an object to retrieve the results of a speech recognition request. You don’t create these objects directly. Instead, the Speech framework creates them and passes them to the handler block or delegate object you specified when starting your speech recognition task. A speech recognition result object contains one or more of the current utterance. Each transcription has a confidence rating indicating how likely it is to be correct. You can also get the transcription with the highest rating directly from the property. If you requested partial results from the speech recognizer, the transcriptions may represent only part of the total audio content. Use the property to determine if the request contains partial or final results.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionResult
type SFSpeechRecognitionResult struct {
	objectivec.Object
}

// SFSpeechRecognitionResultFrom constructs a [SFSpeechRecognitionResult] from an unsafe.Pointer.
//
// An object that contains the partial or final results of a speech recognition request.
func SFSpeechRecognitionResultFrom(ptr unsafe.Pointer) SFSpeechRecognitionResult {
	return SFSpeechRecognitionResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSpeechRecognitionResultClass) Alloc() SFSpeechRecognitionResult {
	rv := objc.Send[SFSpeechRecognitionResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSpeechRecognitionResultClass) New() SFSpeechRecognitionResult {
	rv := objc.Send[SFSpeechRecognitionResult](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSpeechRecognitionResult) Init() SFSpeechRecognitionResult {
	rv := objc.Send[SFSpeechRecognitionResult](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSpeechRecognitionResult) Autorelease() SFSpeechRecognitionResult {
	rv := objc.Send[SFSpeechRecognitionResult](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechRecognitionResult creates a new SFSpeechRecognitionResult instance.
func NewSFSpeechRecognitionResult() SFSpeechRecognitionResult {
	return getSFSpeechRecognitionResultClass().New()
}




