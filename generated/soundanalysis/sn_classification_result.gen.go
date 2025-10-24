// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coremedia"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SNClassificationResult] class.
var (
	SNClassificationResultClass     _SNClassificationResultClass
	SNClassificationResultClassOnce sync.Once
)

func getSNClassificationResultClass() _SNClassificationResultClass {
	SNClassificationResultClassOnce.Do(func() {
		SNClassificationResultClass = _SNClassificationResultClass{objc.GetClass("SNClassificationResult")}
	})
	return SNClassificationResultClass
}

type _SNClassificationResultClass struct {
	class objc.Class
}

// An interface definition for the [SNClassificationResult] class.
type ISNClassificationResult interface {
	objectivec.IObject
	// properties:
	Classifications() []ISNClassification
	TimeRange() objc.IObject /* cross-framework: TimeRange */
	SetTimeRange(value objc.IObject /* cross-framework: TimeRange */)
	// methods:
}

// A result that contains the highest-ranking classifications in a time range.
//
// An represents the predictions that a sound classification model made for a time span in an audio file or stream. Each result contains one or more classification predictions and a time range within the audio data. An audio analyzer, such as and , produces an each time it recognizes a sound for any of its instances.


// A result that contains the highest-ranking classifications in a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassificationResult
type SNClassificationResult struct {
	objectivec.Object
}

// SNClassificationResultFrom constructs a [SNClassificationResult] from an unsafe.Pointer.
//
// A result that contains the highest-ranking classifications in a time range.
func SNClassificationResultFrom(ptr unsafe.Pointer) SNClassificationResult {
	return SNClassificationResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SNClassificationResultClass) Alloc() SNClassificationResult {
	rv := objc.Send[SNClassificationResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SNClassificationResultClass) New() SNClassificationResult {
	rv := objc.Send[SNClassificationResult](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SNClassificationResult) Init() SNClassificationResult {
	rv := objc.Send[SNClassificationResult](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SNClassificationResult) Autorelease() SNClassificationResult {
	rv := objc.Send[SNClassificationResult](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNClassificationResult creates a new SNClassificationResult instance.
func NewSNClassificationResult() SNClassificationResult {
	return getSNClassificationResultClass().New()
}



// A sorted array of the request’s top classification candidates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassificationResult/classifications
func (s_ SNClassificationResult) Classifications() []ISNClassification {
	rv := objc.Send[[]SNClassification](s_.ID, objc.Sel("classifications"))
	return rv
}


// The time span that corresponds to the result’s classifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassificationresult/timerange
func (s_ SNClassificationResult) TimeRange() objc.IObject /* cross-framework: TimeRange */ {
	rv := objc.Send[coremedia.TimeRange](s_.ID, objc.Sel("timeRange"))
	return rv
}


// The time span that corresponds to the result’s classifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassificationresult/timerange
func (s_ SNClassificationResult) SetTimeRange(value objc.IObject /* cross-framework: TimeRange */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTimeRange:"), value)
}



