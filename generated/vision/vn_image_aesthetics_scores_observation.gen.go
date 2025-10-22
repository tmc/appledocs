// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageAestheticsScoresObservation] class.
var (
	ImageAestheticsScoresObservationClass     _ImageAestheticsScoresObservationClass
	ImageAestheticsScoresObservationClassOnce sync.Once
)

func getImageAestheticsScoresObservationClass() _ImageAestheticsScoresObservationClass {
	ImageAestheticsScoresObservationClassOnce.Do(func() {
		ImageAestheticsScoresObservationClass = _ImageAestheticsScoresObservationClass{objc.GetClass("VNImageAestheticsScoresObservation")}
	})
	return ImageAestheticsScoresObservationClass
}

type _ImageAestheticsScoresObservationClass struct {
	class objc.Class
}

// An interface definition for the [ImageAestheticsScoresObservation] class.
type IImageAestheticsScoresObservation interface {
	IObservation
	OverallScore() float32
	Results() VNImageAestheticsScoresObservation
	SetResults(value IVNImageAestheticsScoresObservation)
	IsUtility() bool
	SetIsUtility(value bool)
}

// An object that represents the overall score of aesthetic attributes for an image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageAestheticsScoresObservation
type ImageAestheticsScoresObservation struct {
	Observation
}

// ImageAestheticsScoresObservationFrom constructs a [ImageAestheticsScoresObservation] from an unsafe.Pointer.
//
// An object that represents the overall score of aesthetic attributes for an image.
func ImageAestheticsScoresObservationFrom(ptr unsafe.Pointer) ImageAestheticsScoresObservation {
	return ImageAestheticsScoresObservation{
		Observation: ObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageAestheticsScoresObservationClass) Alloc() ImageAestheticsScoresObservation {
	rv := objc.Send[ImageAestheticsScoresObservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageAestheticsScoresObservationClass) New() ImageAestheticsScoresObservation {
	rv := objc.Send[ImageAestheticsScoresObservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageAestheticsScoresObservation) Init() ImageAestheticsScoresObservation {
	rv := objc.Send[ImageAestheticsScoresObservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageAestheticsScoresObservation) Autorelease() ImageAestheticsScoresObservation {
	rv := objc.Send[ImageAestheticsScoresObservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageAestheticsScoresObservation creates a new ImageAestheticsScoresObservation instance.
func NewImageAestheticsScoresObservation() ImageAestheticsScoresObservation {
	return getImageAestheticsScoresObservationClass().New()
}


// A score which incorporates aesthetic score, failure score, and utility labels.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageAestheticsScoresObservation/overallScore
func (i_ ImageAestheticsScoresObservation) OverallScore() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("overallScore"))
	return rv
}

// The results of the aesthetics request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncalculateimageaestheticsscoresrequest/results
func (i_ ImageAestheticsScoresObservation) Results() VNImageAestheticsScoresObservation {
	rv := objc.Send[VNImageAestheticsScoresObservation](i_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The results of the aesthetics request.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncalculateimageaestheticsscoresrequest/results
func (i_ ImageAestheticsScoresObservation) SetResults(value IVNImageAestheticsScoresObservation) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setResults:"), value)
}

// A Boolean value that represents images that are not necessarily of poor image quality, but may not have memorable or exciting content.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimageaestheticsscoresobservation/isutility
func (i_ ImageAestheticsScoresObservation) IsUtility() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isUtility"))
	return rv
}


// SetIsUtility sets the value of the isUtility property.
// A Boolean value that represents images that are not necessarily of poor image quality, but may not have memorable or exciting content.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimageaestheticsscoresobservation/isutility
func (i_ ImageAestheticsScoresObservation) SetIsUtility(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsUtility:"), value)
}



