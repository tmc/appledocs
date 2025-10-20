// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SNAudioFileAnalyzer] class.
var (
	SNAudioFileAnalyzerClass     _SNAudioFileAnalyzerClass
	SNAudioFileAnalyzerClassOnce sync.Once
)

func getSNAudioFileAnalyzerClass() _SNAudioFileAnalyzerClass {
	SNAudioFileAnalyzerClassOnce.Do(func() {
		SNAudioFileAnalyzerClass = _SNAudioFileAnalyzerClass{objc.GetClass("SNAudioFileAnalyzer")}
	})
	return SNAudioFileAnalyzerClass
}

type _SNAudioFileAnalyzerClass struct {
	class objc.Class
}

// An interface definition for the [SNAudioFileAnalyzer] class.
type ISNAudioFileAnalyzer interface {
	objectivec.IObject
	AddRequestWithObserverError(request objc.ID, observer objc.ID, error_ unsafe.Pointer) bool
	Analyze()
	CancelAnalysis()
	RemoveRequest(request objc.ID)
	RemoveAllRequests()
}

// An analyzer that runs sound classification requests on an audio file.
//
// Run an on an audio file by creating an . You can run the same sound analysis request on multiple file analyzers, and each analyzer can process multiple requests. An audio file analyzer generates an each time any of its active requests recognizes a sound.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer
type SNAudioFileAnalyzer struct {
	objectivec.Object
}

// SNAudioFileAnalyzerFrom constructs a [SNAudioFileAnalyzer] from an unsafe.Pointer.
//
// An analyzer that runs sound classification requests on an audio file.
func SNAudioFileAnalyzerFrom(ptr unsafe.Pointer) SNAudioFileAnalyzer {
	return SNAudioFileAnalyzer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SNAudioFileAnalyzerClass) Alloc() SNAudioFileAnalyzer {
	rv := objc.Send[SNAudioFileAnalyzer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SNAudioFileAnalyzerClass) New() SNAudioFileAnalyzer {
	rv := objc.Send[SNAudioFileAnalyzer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SNAudioFileAnalyzer) Init() SNAudioFileAnalyzer {
	rv := objc.Send[SNAudioFileAnalyzer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SNAudioFileAnalyzer) Autorelease() SNAudioFileAnalyzer {
	rv := objc.Send[SNAudioFileAnalyzer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNAudioFileAnalyzer creates a new SNAudioFileAnalyzer instance.
func NewSNAudioFileAnalyzer() SNAudioFileAnalyzer {
	return getSNAudioFileAnalyzerClass().New()
}


// Adds a new analysis request to the audio file analyzer.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/add(_:withObserver:)
func (s_ SNAudioFileAnalyzer) AddRequestWithObserverError(request objc.ID, observer objc.ID, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("addRequest:withObserver:error:"), request, observer, error_)
	return rv
}

// Analyzes the audio file synchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/analyze()
func (s_ SNAudioFileAnalyzer) Analyze() {
	objc.Send[objc.ID](s_.ID, objc.Sel("analyze"))
}

// Cancels all the asynchronous sound analysis requests the analyzer is currently processing.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/cancelAnalysis()
func (s_ SNAudioFileAnalyzer) CancelAnalysis() {
	objc.Send[objc.ID](s_.ID, objc.Sel("cancelAnalysis"))
}

// Removes an existing request from the audio file analyzer.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/remove(_:)
func (s_ SNAudioFileAnalyzer) RemoveRequest(request objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeRequest:"), request)
}

// Removes all the sound analysis requests from the audio file analyzer.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/removeAllRequests()
func (s_ SNAudioFileAnalyzer) RemoveAllRequests() {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeAllRequests"))
}



