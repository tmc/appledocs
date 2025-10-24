// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNFaceLandmarkRegion */


/* debug [class_header]: Header for VNFaceLandmarkRegion */
// The class instance for the [FaceLandmarkRegion] class.
var (
	FaceLandmarkRegionClass     _FaceLandmarkRegionClass
	FaceLandmarkRegionClassOnce sync.Once
)

func getFaceLandmarkRegionClass() _FaceLandmarkRegionClass {
	FaceLandmarkRegionClassOnce.Do(func() {
		FaceLandmarkRegionClass = _FaceLandmarkRegionClass{objc.GetClass("VNFaceLandmarkRegion")}
	})
	return FaceLandmarkRegionClass
}

type _FaceLandmarkRegionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FaceLandmarkRegion */
// An interface definition for the [FaceLandmarkRegion] class.
type IFaceLandmarkRegion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FaceLandmarkRegion */
	// properties:
	PointCount() uint
	Landmarks() IVNFaceLandmarks2D
	SetLandmarks(value IVNFaceLandmarks2D)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FaceLandmarkRegion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FaceLandmarkRegion */
// Alloc allocates a new instance without initialization.
func (fc _FaceLandmarkRegionClass) Alloc() FaceLandmarkRegion {
	rv := objc.Send[FaceLandmarkRegion](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FaceLandmarkRegionClass) New() FaceLandmarkRegion {
	rv := objc.Send[FaceLandmarkRegion](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FaceLandmarkRegion) Init() FaceLandmarkRegion {
	rv := objc.Send[FaceLandmarkRegion](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FaceLandmarkRegion) Autorelease() FaceLandmarkRegion {
	rv := objc.Send[FaceLandmarkRegion](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFaceLandmarkRegion creates a new FaceLandmarkRegion instance.
func NewFaceLandmarkRegion() FaceLandmarkRegion {
	return getFaceLandmarkRegionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FaceLandmarkRegion */
// The abstract superclass for information about a specific face landmark.


// The abstract superclass for information about a specific face landmark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion
type FaceLandmarkRegion struct {
	objectivec.Object
}

// FaceLandmarkRegionFrom constructs a [FaceLandmarkRegion] from an unsafe.Pointer.
//
// The abstract superclass for information about a specific face landmark.
func FaceLandmarkRegionFrom(ptr unsafe.Pointer) FaceLandmarkRegion {
	return FaceLandmarkRegion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FaceLandmarkRegion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FaceLandmarkRegion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FaceLandmarkRegion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FaceLandmarkRegion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FaceLandmarkRegion */

// The number of points in the face region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion/pointCount
func (f_ FaceLandmarkRegion) PointCount() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("pointCount"))
	return rv
}/* debug [instance_properties/getter]: pointCount */


// The facial features of the detected face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/landmarks
func (f_ FaceLandmarkRegion) Landmarks() IVNFaceLandmarks2D {
	rv := objc.Send[FaceLandmarks2D](f_.ID, objc.Sel("landmarks"))
	return rv
}/* debug [instance_properties/getter]: landmarks */


// The facial features of the detected face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/landmarks
func (f_ FaceLandmarkRegion) SetLandmarks(value IVNFaceLandmarks2D) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLandmarks:"), value)
}/* debug [instance_properties/setter]: landmarks */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNFaceLandmarkRegion */



