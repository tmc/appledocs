// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNImageAlignmentObservation */


/* debug [class_header]: Header for VNImageAlignmentObservation */
// The class instance for the [ImageAlignmentObservation] class.
var (
	ImageAlignmentObservationClass     _ImageAlignmentObservationClass
	ImageAlignmentObservationClassOnce sync.Once
)

func getImageAlignmentObservationClass() _ImageAlignmentObservationClass {
	ImageAlignmentObservationClassOnce.Do(func() {
		ImageAlignmentObservationClass = _ImageAlignmentObservationClass{objc.GetClass("VNImageAlignmentObservation")}
	})
	return ImageAlignmentObservationClass
}

type _ImageAlignmentObservationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageAlignmentObservation */
// An interface definition for the [ImageAlignmentObservation] class.
type IImageAlignmentObservation interface {
	IObservation
	
/* debug [class_interface_properties]: Properties for ImageAlignmentObservation */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageAlignmentObservation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageAlignmentObservation */
// Alloc allocates a new instance without initialization.
func (ic _ImageAlignmentObservationClass) Alloc() ImageAlignmentObservation {
	rv := objc.Send[ImageAlignmentObservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageAlignmentObservationClass) New() ImageAlignmentObservation {
	rv := objc.Send[ImageAlignmentObservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageAlignmentObservation) Init() ImageAlignmentObservation {
	rv := objc.Send[ImageAlignmentObservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageAlignmentObservation) Autorelease() ImageAlignmentObservation {
	rv := objc.Send[ImageAlignmentObservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageAlignmentObservation creates a new ImageAlignmentObservation instance.
func NewImageAlignmentObservation() ImageAlignmentObservation {
	return getImageAlignmentObservationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageAlignmentObservation */
// The abstract superclass for image-analysis results that describe the relative alignment of two images.
//
// This abstract superclass forms the basis of image alignment or registration output. You receive its subclasses, such as and , by performing specific registration requests. Don’t create one of these classes yourself.


// The abstract superclass for image-analysis results that describe the relative alignment of two images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageAlignmentObservation
type ImageAlignmentObservation struct {
	Observation
}

// ImageAlignmentObservationFrom constructs a [ImageAlignmentObservation] from an unsafe.Pointer.
//
// The abstract superclass for image-analysis results that describe the relative alignment of two images.
func ImageAlignmentObservationFrom(ptr unsafe.Pointer) ImageAlignmentObservation {
	return ImageAlignmentObservation{
		Observation: ObservationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageAlignmentObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageAlignmentObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageAlignmentObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageAlignmentObservation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageAlignmentObservation */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNImageAlignmentObservation */



