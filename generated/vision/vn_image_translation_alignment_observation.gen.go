// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)

/* debug [class.gen.go]: Generating class VNImageTranslationAlignmentObservation */


/* debug [class_header]: Header for VNImageTranslationAlignmentObservation */
// The class instance for the [ImageTranslationAlignmentObservation] class.
var (
	ImageTranslationAlignmentObservationClass     _ImageTranslationAlignmentObservationClass
	ImageTranslationAlignmentObservationClassOnce sync.Once
)

func getImageTranslationAlignmentObservationClass() _ImageTranslationAlignmentObservationClass {
	ImageTranslationAlignmentObservationClassOnce.Do(func() {
		ImageTranslationAlignmentObservationClass = _ImageTranslationAlignmentObservationClass{objc.GetClass("VNImageTranslationAlignmentObservation")}
	})
	return ImageTranslationAlignmentObservationClass
}

type _ImageTranslationAlignmentObservationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageTranslationAlignmentObservation */
// An interface definition for the [ImageTranslationAlignmentObservation] class.
type IImageTranslationAlignmentObservation interface {
	IImageAlignmentObservation
	
/* debug [class_interface_properties]: Properties for ImageTranslationAlignmentObservation */
	// properties:
	AlignmentTransform() corefoundation.CGAffineTransform
	VNTranslationalImageRegistrationRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageTranslationAlignmentObservation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageTranslationAlignmentObservation */
// Alloc allocates a new instance without initialization.
func (ic _ImageTranslationAlignmentObservationClass) Alloc() ImageTranslationAlignmentObservation {
	rv := objc.Send[ImageTranslationAlignmentObservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageTranslationAlignmentObservationClass) New() ImageTranslationAlignmentObservation {
	rv := objc.Send[ImageTranslationAlignmentObservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageTranslationAlignmentObservation) Init() ImageTranslationAlignmentObservation {
	rv := objc.Send[ImageTranslationAlignmentObservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageTranslationAlignmentObservation) Autorelease() ImageTranslationAlignmentObservation {
	rv := objc.Send[ImageTranslationAlignmentObservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageTranslationAlignmentObservation creates a new ImageTranslationAlignmentObservation instance.
func NewImageTranslationAlignmentObservation() ImageTranslationAlignmentObservation {
	return getImageTranslationAlignmentObservationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageTranslationAlignmentObservation */
// Affine transform information that an image-alignment request produces.
//
// This type of observation results from a , informing the performed to align the input images.


// Affine transform information that an image-alignment request produces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageTranslationAlignmentObservation
type ImageTranslationAlignmentObservation struct {
	ImageAlignmentObservation
}

// ImageTranslationAlignmentObservationFrom constructs a [ImageTranslationAlignmentObservation] from an unsafe.Pointer.
//
// Affine transform information that an image-alignment request produces.
func ImageTranslationAlignmentObservationFrom(ptr unsafe.Pointer) ImageTranslationAlignmentObservation {
	return ImageTranslationAlignmentObservation{
		ImageAlignmentObservation: ImageAlignmentObservationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageTranslationAlignmentObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageTranslationAlignmentObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageTranslationAlignmentObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageTranslationAlignmentObservation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageTranslationAlignmentObservation */

// The alignment transform to align the floating image with the reference image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageTranslationAlignmentObservation/alignmentTransform
func (i_ ImageTranslationAlignmentObservation) AlignmentTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](i_.ID, objc.Sel("alignmentTransform"))
	return rv
}/* debug [instance_properties/getter]: alignmentTransform */


// A constant for specifying revision 1 of the translational image registration request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntranslationalimageregistrationrequestrevision1
func (i_ ImageTranslationAlignmentObservation) VNTranslationalImageRegistrationRequestRevision1() int {
	rv := objc.Send[int](i_.ID, objc.Sel("VNTranslationalImageRegistrationRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNTranslationalImageRegistrationRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNImageTranslationAlignmentObservation */



