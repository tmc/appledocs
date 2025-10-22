// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RecognizedObjectObservation] class.
var (
	RecognizedObjectObservationClass     _RecognizedObjectObservationClass
	RecognizedObjectObservationClassOnce sync.Once
)

func getRecognizedObjectObservationClass() _RecognizedObjectObservationClass {
	RecognizedObjectObservationClassOnce.Do(func() {
		RecognizedObjectObservationClass = _RecognizedObjectObservationClass{objc.GetClass("VNRecognizedObjectObservation")}
	})
	return RecognizedObjectObservationClass
}

type _RecognizedObjectObservationClass struct {
	class objc.Class
}

// An interface definition for the [RecognizedObjectObservation] class.
type IRecognizedObjectObservation interface {
	IDetectedObjectObservation
	Labels() []ClassificationObservation
}

// A detected object observation with an array of classification labels that classify the recognized object.
//
// The confidence of the classifications sum up to Multiply the classification confidence with the confidence of this observation.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedObjectObservation
type RecognizedObjectObservation struct {
	DetectedObjectObservation
}

// RecognizedObjectObservationFrom constructs a [RecognizedObjectObservation] from an unsafe.Pointer.
//
// A detected object observation with an array of classification labels that classify the recognized object.
func RecognizedObjectObservationFrom(ptr unsafe.Pointer) RecognizedObjectObservation {
	return RecognizedObjectObservation{
		DetectedObjectObservation: DetectedObjectObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RecognizedObjectObservationClass) Alloc() RecognizedObjectObservation {
	rv := objc.Send[RecognizedObjectObservation](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RecognizedObjectObservationClass) New() RecognizedObjectObservation {
	rv := objc.Send[RecognizedObjectObservation](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizedObjectObservation) Init() RecognizedObjectObservation {
	rv := objc.Send[RecognizedObjectObservation](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizedObjectObservation) Autorelease() RecognizedObjectObservation {
	rv := objc.Send[RecognizedObjectObservation](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizedObjectObservation creates a new RecognizedObjectObservation instance.
func NewRecognizedObjectObservation() RecognizedObjectObservation {
	return getRecognizedObjectObservationClass().New()
}


// An array of observations that classify the recognized object.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedObjectObservation/labels
func (r_ RecognizedObjectObservation) Labels() []ClassificationObservation {
	rv := objc.Send[[]ClassificationObservation](r_.ID, objc.Sel("labels"))
	return rv
}



