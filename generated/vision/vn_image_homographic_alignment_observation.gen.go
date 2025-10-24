// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNImageHomographicAlignmentObservation */


/* debug [class_header]: Header for VNImageHomographicAlignmentObservation */
// The class instance for the [ImageHomographicAlignmentObservation] class.
var (
	ImageHomographicAlignmentObservationClass     _ImageHomographicAlignmentObservationClass
	ImageHomographicAlignmentObservationClassOnce sync.Once
)

func getImageHomographicAlignmentObservationClass() _ImageHomographicAlignmentObservationClass {
	ImageHomographicAlignmentObservationClassOnce.Do(func() {
		ImageHomographicAlignmentObservationClass = _ImageHomographicAlignmentObservationClass{objc.GetClass("VNImageHomographicAlignmentObservation")}
	})
	return ImageHomographicAlignmentObservationClass
}

type _ImageHomographicAlignmentObservationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageHomographicAlignmentObservation */
// An interface definition for the [ImageHomographicAlignmentObservation] class.
type IImageHomographicAlignmentObservation interface {
	IImageAlignmentObservation
	
/* debug [class_interface_properties]: Properties for ImageHomographicAlignmentObservation */
	// properties:
	WarpTransform() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageHomographicAlignmentObservation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageHomographicAlignmentObservation */
// Alloc allocates a new instance without initialization.
func (ic _ImageHomographicAlignmentObservationClass) Alloc() ImageHomographicAlignmentObservation {
	rv := objc.Send[ImageHomographicAlignmentObservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageHomographicAlignmentObservationClass) New() ImageHomographicAlignmentObservation {
	rv := objc.Send[ImageHomographicAlignmentObservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageHomographicAlignmentObservation) Init() ImageHomographicAlignmentObservation {
	rv := objc.Send[ImageHomographicAlignmentObservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageHomographicAlignmentObservation) Autorelease() ImageHomographicAlignmentObservation {
	rv := objc.Send[ImageHomographicAlignmentObservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageHomographicAlignmentObservation creates a new ImageHomographicAlignmentObservation instance.
func NewImageHomographicAlignmentObservation() ImageHomographicAlignmentObservation {
	return getImageHomographicAlignmentObservationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageHomographicAlignmentObservation */
// An object that represents a perspective warp transformation.
//
// This type of observation results from a , informing the performed to align the input images.


// An object that represents a perspective warp transformation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageHomographicAlignmentObservation
type ImageHomographicAlignmentObservation struct {
	ImageAlignmentObservation
}

// ImageHomographicAlignmentObservationFrom constructs a [ImageHomographicAlignmentObservation] from an unsafe.Pointer.
//
// An object that represents a perspective warp transformation.
func ImageHomographicAlignmentObservationFrom(ptr unsafe.Pointer) ImageHomographicAlignmentObservation {
	return ImageHomographicAlignmentObservation{
		ImageAlignmentObservation: ImageAlignmentObservationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageHomographicAlignmentObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageHomographicAlignmentObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageHomographicAlignmentObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageHomographicAlignmentObservation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageHomographicAlignmentObservation */

// The warp transform matrix to morph the floating image into the reference image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageHomographicAlignmentObservation/warpTransform
func (i_ ImageHomographicAlignmentObservation) WarpTransform() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("warpTransform"))
	return rv
}/* debug [instance_properties/getter]: warpTransform */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNImageHomographicAlignmentObservation */



