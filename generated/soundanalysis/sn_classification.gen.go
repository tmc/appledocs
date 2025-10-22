// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SNClassification] class.
var (
	SNClassificationClass     _SNClassificationClass
	SNClassificationClassOnce sync.Once
)

func getSNClassificationClass() _SNClassificationClass {
	SNClassificationClassOnce.Do(func() {
		SNClassificationClass = _SNClassificationClass{objc.GetClass("SNClassification")}
	})
	return SNClassificationClass
}

type _SNClassificationClass struct {
	class objc.Class
}

// An interface definition for the [SNClassification] class.
type ISNClassification interface {
	objectivec.IObject
	Identifier() string
	Confidence() float64
	SetConfidence(value float64)
	Classifications() SNClassification
	SetClassifications(value ISNClassification)
	TimeRange() unsafe.Pointer
	SetTimeRange(value unsafe.Pointer)
}

// A type that pairs a sound classifier’s prediction with its confidence in that prediction.
//
// An represents a single sound classification prediction, and the sound classifier model’s confidence in that prediction.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassification
type SNClassification struct {
	objectivec.Object
}

// SNClassificationFrom constructs a [SNClassification] from an unsafe.Pointer.
//
// A type that pairs a sound classifier’s prediction with its confidence in that prediction.
func SNClassificationFrom(ptr unsafe.Pointer) SNClassification {
	return SNClassification{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SNClassificationClass) Alloc() SNClassification {
	rv := objc.Send[SNClassification](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SNClassificationClass) New() SNClassification {
	rv := objc.Send[SNClassification](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SNClassification) Init() SNClassification {
	rv := objc.Send[SNClassification](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SNClassification) Autorelease() SNClassification {
	rv := objc.Send[SNClassification](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNClassification creates a new SNClassification instance.
func NewSNClassification() SNClassification {
	return getSNClassificationClass().New()
}


// A prediction label that’s one of the classifications a sound classifier’s underlying model defines.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassification/identifier
func (s_ SNClassification) Identifier() string {
	rv := objc.Send[string](s_.ID, objc.Sel("identifier"))
	return rv
}

// The confidence value the model has in its prediction.
//
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassification/confidence
func (s_ SNClassification) Confidence() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("confidence"))
	return rv
}


// SetConfidence sets the value of the confidence property.
// The confidence value the model has in its prediction.

//
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassification/confidence
func (s_ SNClassification) SetConfidence(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setConfidence:"), value)
}

// A sorted array of the request’s top classification candidates.
//
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassificationresult/classifications
func (s_ SNClassification) Classifications() SNClassification {
	rv := objc.Send[SNClassification](s_.ID, objc.Sel("classifications"))
	return rv
}


// SetClassifications sets the value of the classifications property.
// A sorted array of the request’s top classification candidates.

//
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassificationresult/classifications
func (s_ SNClassification) SetClassifications(value ISNClassification) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setClassifications:"), value)
}

// The time span that corresponds to the result’s classifications.
//
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassificationresult/timerange
func (s_ SNClassification) TimeRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("timeRange"))
	return rv
}


// SetTimeRange sets the value of the timeRange property.
// The time span that corresponds to the result’s classifications.

//
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassificationresult/timerange
func (s_ SNClassification) SetTimeRange(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTimeRange:"), value)
}



