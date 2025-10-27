// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [DetectedObjectObservation] class.
var (
	DetectedObjectObservationClass     _DetectedObjectObservationClass
	DetectedObjectObservationClassOnce sync.Once
)

func getDetectedObjectObservationClass() _DetectedObjectObservationClass {
	DetectedObjectObservationClassOnce.Do(func() {
		DetectedObjectObservationClass = _DetectedObjectObservationClass{objc.GetClass("VNDetectedObjectObservation")}
	})
	return DetectedObjectObservationClass
}

type _DetectedObjectObservationClass struct {
	class objc.Class
}





// An interface definition for the [DetectedObjectObservation] class.
type IDetectedObjectObservation interface {
	IObservation
	

	// properties:
	BoundingBox() corefoundation.CGRect
	GlobalSegmentationMask() IVNPixelBufferObservation


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (dc _DetectedObjectObservationClass) Alloc() DetectedObjectObservation {
	rv := objc.Send[DetectedObjectObservation](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DetectedObjectObservationClass) New() DetectedObjectObservation {
	rv := objc.Send[DetectedObjectObservation](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectedObjectObservation) Init() DetectedObjectObservation {
	rv := objc.Send[DetectedObjectObservation](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectedObjectObservation) Autorelease() DetectedObjectObservation {
	rv := objc.Send[DetectedObjectObservation](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectedObjectObservation creates a new DetectedObjectObservation instance.
func NewDetectedObjectObservation() DetectedObjectObservation {
	return getDetectedObjectObservationClass().New()
}





// An observation that provides the position and extent of an image feature that an image- analysis request detects.
//
// This class is the observation type that generates. It represents an object that the Vision request detects and tracks.


// An observation that provides the position and extent of an image feature that an image- analysis request detects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation
type DetectedObjectObservation struct {
	Observation
}

// DetectedObjectObservationFrom constructs a [DetectedObjectObservation] from an unsafe.Pointer.
//
// An observation that provides the position and extent of an image feature that an image- analysis request detects.
func DetectedObjectObservationFrom(ptr unsafe.Pointer) DetectedObjectObservation {
	return DetectedObjectObservation{
		Observation: ObservationFrom(ptr),
	}
}






// Creates an observation with a bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(boundingBox:)
func NewDetectedObjectObservationWithBoundingBox(boundingBox corefoundation.CGRect) DetectedObjectObservation {
	rv := objc.Send[DetectedObjectObservation](objc.ID(getDetectedObjectObservationClass().class), objc.Sel("observationWithBoundingBox:"), boundingBox)
	return rv
}


// Creates an observation with a revision number and bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(requestRevision:boundingBox:)
func NewDetectedObjectObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox corefoundation.CGRect) DetectedObjectObservation {
	rv := objc.Send[DetectedObjectObservation](objc.ID(getDetectedObjectObservationClass().class), objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return rv
}







// Creates an observation with a bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(boundingBox:)
func (dc _DetectedObjectObservationClass) ObservationWithBoundingBox(boundingBox corefoundation.CGRect) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("observationWithBoundingBox:"), boundingBox)
	return rv
}


// Creates an observation with a revision number and bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(requestRevision:boundingBox:)
func (dc _DetectedObjectObservationClass) ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox corefoundation.CGRect) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return rv
}

















// The bounding box of the object that the request detects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/boundingBox
func (d_ DetectedObjectObservation) BoundingBox() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](d_.ID, objc.Sel("boundingBox"))
	return rv
}


// A resulting pixel buffer from a request to generate a segmentation mask for an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/globalSegmentationMask
func (d_ DetectedObjectObservation) GlobalSegmentationMask() IVNPixelBufferObservation {
	rv := objc.Send[PixelBufferObservation](d_.ID, objc.Sel("globalSegmentationMask"))
	return rv
}







