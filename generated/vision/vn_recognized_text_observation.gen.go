// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [RecognizedTextObservation] class.
var (
	RecognizedTextObservationClass     _RecognizedTextObservationClass
	RecognizedTextObservationClassOnce sync.Once
)

func getRecognizedTextObservationClass() _RecognizedTextObservationClass {
	RecognizedTextObservationClassOnce.Do(func() {
		RecognizedTextObservationClass = _RecognizedTextObservationClass{objc.GetClass("VNRecognizedTextObservation")}
	})
	return RecognizedTextObservationClass
}

type _RecognizedTextObservationClass struct {
	class objc.Class
}





// An interface definition for the [RecognizedTextObservation] class.
type IRecognizedTextObservation interface {
	IRectangleObservation
	

	// properties:


	

	// methods:
	TopCandidates(maxCandidateCount uint) []RecognizedText


}





// Alloc allocates a new instance without initialization.
func (rc _RecognizedTextObservationClass) Alloc() RecognizedTextObservation {
	rv := objc.Send[RecognizedTextObservation](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RecognizedTextObservationClass) New() RecognizedTextObservation {
	rv := objc.Send[RecognizedTextObservation](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizedTextObservation) Init() RecognizedTextObservation {
	rv := objc.Send[RecognizedTextObservation](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizedTextObservation) Autorelease() RecognizedTextObservation {
	rv := objc.Send[RecognizedTextObservation](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizedTextObservation creates a new RecognizedTextObservation instance.
func NewRecognizedTextObservation() RecognizedTextObservation {
	return getRecognizedTextObservationClass().New()
}





// A request that detects and recognizes regions of text in an image.
//
// This type of observation results from a . It contains information about both the location and content of text and glyphs that Vision recognized in the input image.


// A request that detects and recognizes regions of text in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedTextObservation
type RecognizedTextObservation struct {
	RectangleObservation
}

// RecognizedTextObservationFrom constructs a [RecognizedTextObservation] from an unsafe.Pointer.
//
// A request that detects and recognizes regions of text in an image.
func RecognizedTextObservationFrom(ptr unsafe.Pointer) RecognizedTextObservation {
	return RecognizedTextObservation{
		RectangleObservation: RectangleObservationFrom(ptr),
	}
}




















// Requests the top candidates for a recognized text string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedTextObservation/topCandidates(_:)
func (r_ RecognizedTextObservation) TopCandidates(maxCandidateCount uint) []RecognizedText {
	rv := objc.Send[[]RecognizedText](r_.ID, objc.Sel("topCandidates:"), maxCandidateCount)
	return rv
}













