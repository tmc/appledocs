// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SNAudioFileAnalyzer */


/* debug [class_header]: Header for SNAudioFileAnalyzer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SNAudioFileAnalyzer */
// An interface definition for the [SNAudioFileAnalyzer] class.
type ISNAudioFileAnalyzer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SNAudioFileAnalyzer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SNAudioFileAnalyzer */
	// methods:
	AddRequestWithObserverError(request unsafe.Pointer, observer unsafe.Pointer, error_ unsafe.Pointer) bool
	Analyze()
	AnalyzeWithCompletionHandler(completionHandler unsafe.Pointer)
	CancelAnalysis()
	RemoveRequest(request unsafe.Pointer)
	RemoveAllRequests()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SNAudioFileAnalyzer */
// Alloc allocates a new instance without initialization.
func (sc _SNAudioFileAnalyzerClass) Alloc() SNAudioFileAnalyzer {
	rv := objc.Send[SNAudioFileAnalyzer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SNAudioFileAnalyzer */
// An analyzer that runs sound classification requests on an audio file.
//
// Run an on an audio file by creating an . You can run the same sound analysis request on multiple file analyzers, and each analyzer can process multiple requests. An audio file analyzer generates an each time any of its active requests recognizes a sound.


// An analyzer that runs sound classification requests on an audio file.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SNAudioFileAnalyzer */

// Creates a new audio file analyzer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/init(url:)
func NewSNAudioFileAnalyzerWithURLError(url objc.IObject /* cross-framework: NSURL */, error_ unsafe.Pointer) SNAudioFileAnalyzer {
	instance := getSNAudioFileAnalyzerClass().Alloc()
	rv := objc.Send[SNAudioFileAnalyzer](instance.ID, objc.Sel("initWithURL:error:"), url, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSNAudioFileAnalyzerWithURLError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SNAudioFileAnalyzer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SNAudioFileAnalyzer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SNAudioFileAnalyzer */

// Adds a new analysis request to the audio file analyzer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/add(_:withObserver:)
func (s_ SNAudioFileAnalyzer) AddRequestWithObserverError(request unsafe.Pointer, observer unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("addRequest:withObserver:error:"), request, observer, error_)
	return rv
}/* debug [instance_methods/method]: AddRequestWithObserverError */


// Analyzes the audio file synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/analyze()
func (s_ SNAudioFileAnalyzer) Analyze() {
	objc.Send[objc.ID](s_.ID, objc.Sel("analyze"))
}/* debug [instance_methods/method]: Analyze */


// Analyzes the audio file asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/analyze(completionHandler:)
func (s_ SNAudioFileAnalyzer) AnalyzeWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("analyzeWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: AnalyzeWithCompletionHandler */


// Cancels all the asynchronous sound analysis requests the analyzer is currently processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/cancelAnalysis()
func (s_ SNAudioFileAnalyzer) CancelAnalysis() {
	objc.Send[objc.ID](s_.ID, objc.Sel("cancelAnalysis"))
}/* debug [instance_methods/method]: CancelAnalysis */


// Removes an existing request from the audio file analyzer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/remove(_:)
func (s_ SNAudioFileAnalyzer) RemoveRequest(request unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeRequest:"), request)
}/* debug [instance_methods/method]: RemoveRequest */


// Removes all the sound analysis requests from the audio file analyzer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioFileAnalyzer/removeAllRequests()
func (s_ SNAudioFileAnalyzer) RemoveAllRequests() {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeAllRequests"))
}/* debug [instance_methods/method]: RemoveAllRequests */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SNAudioFileAnalyzer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SNAudioFileAnalyzer */


