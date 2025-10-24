// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNDetectedObjectObservation */


/* debug [class_header]: Header for VNDetectedObjectObservation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DetectedObjectObservation */
// An interface definition for the [DetectedObjectObservation] class.
type IDetectedObjectObservation interface {
	IObservation
	
/* debug [class_interface_properties]: Properties for DetectedObjectObservation */
	// properties:
	BoundingBox() corefoundation.CGRect
	GlobalSegmentationMask() IVNPixelBufferObservation
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DetectedObjectObservation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DetectedObjectObservation */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DetectedObjectObservation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DetectedObjectObservation */

// Creates an observation with a bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(boundingBox:)
func NewDetectedObjectObservationWithBoundingBox(boundingBox corefoundation.CGRect) DetectedObjectObservation {
	rv := objc.Send[DetectedObjectObservation](objc.ID(getDetectedObjectObservationClass().class), objc.Sel("observationWithBoundingBox:"), boundingBox)
	return rv
}/* debug [class_init_methods/constructor]: NewDetectedObjectObservationWithBoundingBox */


// Creates an observation with a revision number and bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(requestRevision:boundingBox:)
func NewDetectedObjectObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox corefoundation.CGRect) DetectedObjectObservation {
	rv := objc.Send[DetectedObjectObservation](objc.ID(getDetectedObjectObservationClass().class), objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return rv
}/* debug [class_init_methods/constructor]: NewDetectedObjectObservationWithRequestRevisionBoundingBox */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DetectedObjectObservation */

// Creates an observation with a bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(boundingBox:)
func (dc _DetectedObjectObservationClass) ObservationWithBoundingBox(boundingBox corefoundation.CGRect) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("observationWithBoundingBox:"), boundingBox)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ObservationWithBoundingBox) */


// Creates an observation with a revision number and bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(requestRevision:boundingBox:)
func (dc _DetectedObjectObservationClass) ObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox corefoundation.CGRect) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ObservationWithRequestRevisionBoundingBox) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DetectedObjectObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DetectedObjectObservation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DetectedObjectObservation */

// The bounding box of the object that the request detects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/boundingBox
func (d_ DetectedObjectObservation) BoundingBox() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](d_.ID, objc.Sel("boundingBox"))
	return rv
}/* debug [instance_properties/getter]: boundingBox */


// A resulting pixel buffer from a request to generate a segmentation mask for an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/globalSegmentationMask
func (d_ DetectedObjectObservation) GlobalSegmentationMask() IVNPixelBufferObservation {
	rv := objc.Send[PixelBufferObservation](d_.ID, objc.Sel("globalSegmentationMask"))
	return rv
}/* debug [instance_properties/getter]: globalSegmentationMask */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNDetectedObjectObservation */


