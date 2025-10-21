// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [SNAudioStreamAnalyzer] class.
type ISNAudioStreamAnalyzer interface {
	objectivec.IObject
}

// An object you create to analyze a stream of audio data and provide the results to your app.
//
// Run an on an audio stream by creating an . You can run the same sound analysis request on multiple stream analyzers, and each analyzer can process multiple requests. An audio file analyzer generates an each time any of its active requests recognizes a sound.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SNAudioStreamAnalyzerClass) Alloc() SNAudioStreamAnalyzer {
	rv := objc.Send[SNAudioStreamAnalyzer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a new audio stream analyzer.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer/init(format:)
func NewSNAudioStreamAnalyzerWithFormat(format unsafe.Pointer) SNAudioStreamAnalyzer {
	instance := getSNAudioStreamAnalyzerClass().Alloc()
	rv := objc.Send[SNAudioStreamAnalyzer](instance.ID, objc.Sel("initWithFormat:"), format)
	rv.Autorelease()
	return rv
}



