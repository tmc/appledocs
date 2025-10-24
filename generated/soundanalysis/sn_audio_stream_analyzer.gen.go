// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SNAudioStreamAnalyzer */


/* debug [class_header]: Header for SNAudioStreamAnalyzer */
// The class instance for the [SNAudioStreamAnalyzer] class.
var (
	SNAudioStreamAnalyzerClass     _SNAudioStreamAnalyzerClass
	SNAudioStreamAnalyzerClassOnce sync.Once
)

func getSNAudioStreamAnalyzerClass() _SNAudioStreamAnalyzerClass {
	SNAudioStreamAnalyzerClassOnce.Do(func() {
		SNAudioStreamAnalyzerClass = _SNAudioStreamAnalyzerClass{objc.GetClass("SNAudioStreamAnalyzer")}
	})
	return SNAudioStreamAnalyzerClass
}

type _SNAudioStreamAnalyzerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SNAudioStreamAnalyzer */
// An interface definition for the [SNAudioStreamAnalyzer] class.
type ISNAudioStreamAnalyzer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SNAudioStreamAnalyzer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SNAudioStreamAnalyzer */
	// methods:
	AddRequestWithObserverError(request unsafe.Pointer, observer unsafe.Pointer, error_ unsafe.Pointer) bool
	AnalyzeAudioBufferAtAudioFramePosition(audioBuffer avfaudio.AudioBuffer, audioFramePosition AudioFramePosition /* not a class type */)
	CompleteAnalysis()
	RemoveRequest(request unsafe.Pointer)
	RemoveAllRequests()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SNAudioStreamAnalyzer */
// Alloc allocates a new instance without initialization.
func (sc _SNAudioStreamAnalyzerClass) Alloc() SNAudioStreamAnalyzer {
	rv := objc.Send[SNAudioStreamAnalyzer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SNAudioStreamAnalyzerClass) New() SNAudioStreamAnalyzer {
	rv := objc.Send[SNAudioStreamAnalyzer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SNAudioStreamAnalyzer) Init() SNAudioStreamAnalyzer {
	rv := objc.Send[SNAudioStreamAnalyzer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SNAudioStreamAnalyzer) Autorelease() SNAudioStreamAnalyzer {
	rv := objc.Send[SNAudioStreamAnalyzer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNAudioStreamAnalyzer creates a new SNAudioStreamAnalyzer instance.
func NewSNAudioStreamAnalyzer() SNAudioStreamAnalyzer {
	return getSNAudioStreamAnalyzerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SNAudioStreamAnalyzer */
// An object you create to analyze a stream of audio data and provide the results to your app.
//
// Run an on an audio stream by creating an . You can run the same sound analysis request on multiple stream analyzers, and each analyzer can process multiple requests. An audio file analyzer generates an each time any of its active requests recognizes a sound.


// An object you create to analyze a stream of audio data and provide the results to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer
type SNAudioStreamAnalyzer struct {
	objectivec.Object
}

// SNAudioStreamAnalyzerFrom constructs a [SNAudioStreamAnalyzer] from an unsafe.Pointer.
//
// An object you create to analyze a stream of audio data and provide the results to your app.
func SNAudioStreamAnalyzerFrom(ptr unsafe.Pointer) SNAudioStreamAnalyzer {
	return SNAudioStreamAnalyzer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SNAudioStreamAnalyzer */

// Creates a new audio stream analyzer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer/init(format:)
func NewSNAudioStreamAnalyzerWithFormat(format avfaudio.AudioFormat) SNAudioStreamAnalyzer {
	instance := getSNAudioStreamAnalyzerClass().Alloc()
	rv := objc.Send[SNAudioStreamAnalyzer](instance.ID, objc.Sel("initWithFormat:"), format)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSNAudioStreamAnalyzerWithFormat */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SNAudioStreamAnalyzer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SNAudioStreamAnalyzer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SNAudioStreamAnalyzer */

// Adds a new analysis request to the audio stream analyzer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer/add(_:withObserver:)
func (s_ SNAudioStreamAnalyzer) AddRequestWithObserverError(request unsafe.Pointer, observer unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("addRequest:withObserver:error:"), request, observer, error_)
	return rv
}/* debug [instance_methods/method]: AddRequestWithObserverError */


// Adds a new audio buffer to the analyzer’s larger stream buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer/analyze(_:atAudioFramePosition:)
func (s_ SNAudioStreamAnalyzer) AnalyzeAudioBufferAtAudioFramePosition(audioBuffer avfaudio.AudioBuffer, audioFramePosition AudioFramePosition /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("analyzeAudioBuffer:atAudioFramePosition:"), audioBuffer, audioFramePosition)
}/* debug [instance_methods/method]: AnalyzeAudioBufferAtAudioFramePosition */


// Notifies the analyzer when it receives the final audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer/completeAnalysis()
func (s_ SNAudioStreamAnalyzer) CompleteAnalysis() {
	objc.Send[objc.ID](s_.ID, objc.Sel("completeAnalysis"))
}/* debug [instance_methods/method]: CompleteAnalysis */


// Removes an existing request from the audio stream analyzer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer/remove(_:)
func (s_ SNAudioStreamAnalyzer) RemoveRequest(request unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeRequest:"), request)
}/* debug [instance_methods/method]: RemoveRequest */


// Removes all the sound analysis requests from the audio stream analyzer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer/removeAllRequests()
func (s_ SNAudioStreamAnalyzer) RemoveAllRequests() {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeAllRequests"))
}/* debug [instance_methods/method]: RemoveAllRequests */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SNAudioStreamAnalyzer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SNAudioStreamAnalyzer */


