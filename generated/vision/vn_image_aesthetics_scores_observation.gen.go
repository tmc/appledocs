// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNImageAestheticsScoresObservation */


/* debug [class_header]: Header for VNImageAestheticsScoresObservation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageAestheticsScoresObservation */
// An interface definition for the [ImageAestheticsScoresObservation] class.
type IImageAestheticsScoresObservation interface {
	IObservation
	
/* debug [class_interface_properties]: Properties for ImageAestheticsScoresObservation */
	// properties:
	IsUtility() bool
	OverallScore() float32
	Results() IVNImageAestheticsScoresObservation
	SetResults(value IVNImageAestheticsScoresObservation)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageAestheticsScoresObservation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageAestheticsScoresObservation */
// Alloc allocates a new instance without initialization.
func (ic _ImageAestheticsScoresObservationClass) Alloc() ImageAestheticsScoresObservation {
	rv := objc.Send[ImageAestheticsScoresObservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageAestheticsScoresObservation */
// An object that represents the overall score of aesthetic attributes for an image.


// An object that represents the overall score of aesthetic attributes for an image.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageAestheticsScoresObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageAestheticsScoresObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageAestheticsScoresObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageAestheticsScoresObservation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageAestheticsScoresObservation */

// A Boolean value that represents images that are not necessarily of poor image quality, but may not have memorable or exciting content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageAestheticsScoresObservation/isUtility
func (i_ ImageAestheticsScoresObservation) IsUtility() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isUtility"))
	return rv
}/* debug [instance_properties/getter]: isUtility */


// A score which incorporates aesthetic score, failure score, and utility labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageAestheticsScoresObservation/overallScore
func (i_ ImageAestheticsScoresObservation) OverallScore() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("overallScore"))
	return rv
}/* debug [instance_properties/getter]: overallScore */


// The results of the aesthetics request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncalculateimageaestheticsscoresrequest/results
func (i_ ImageAestheticsScoresObservation) Results() IVNImageAestheticsScoresObservation {
	rv := objc.Send[ImageAestheticsScoresObservation](i_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// The results of the aesthetics request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncalculateimageaestheticsscoresrequest/results
func (i_ ImageAestheticsScoresObservation) SetResults(value IVNImageAestheticsScoresObservation) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setResults:"), value)
}/* debug [instance_properties/setter]: results */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNImageAestheticsScoresObservation */



