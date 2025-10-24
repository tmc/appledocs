// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)

/* debug [class.gen.go]: Generating class VNFaceLandmarks2D */


/* debug [class_header]: Header for VNFaceLandmarks2D */
// The class instance for the [FaceLandmarks2D] class.
var (
	FaceLandmarks2DClass     _FaceLandmarks2DClass
	FaceLandmarks2DClassOnce sync.Once
)

func getFaceLandmarks2DClass() _FaceLandmarks2DClass {
	FaceLandmarks2DClassOnce.Do(func() {
		FaceLandmarks2DClass = _FaceLandmarks2DClass{objc.GetClass("VNFaceLandmarks2D")}
	})
	return FaceLandmarks2DClass
}

type _FaceLandmarks2DClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FaceLandmarks2D */
// An interface definition for the [FaceLandmarks2D] class.
type IFaceLandmarks2D interface {
	IFaceLandmarks
	
/* debug [class_interface_properties]: Properties for FaceLandmarks2D */
	// properties:
	AllPoints() IVNFaceLandmarkRegion2D
	FaceContour() IVNFaceLandmarkRegion2D
	InnerLips() IVNFaceLandmarkRegion2D
	LeftEye() IVNFaceLandmarkRegion2D
	LeftEyebrow() IVNFaceLandmarkRegion2D
	LeftPupil() IVNFaceLandmarkRegion2D
	MedianLine() IVNFaceLandmarkRegion2D
	Nose() IVNFaceLandmarkRegion2D
	NoseCrest() IVNFaceLandmarkRegion2D
	OuterLips() IVNFaceLandmarkRegion2D
	RightEye() IVNFaceLandmarkRegion2D
	RightEyebrow() IVNFaceLandmarkRegion2D
	RightPupil() IVNFaceLandmarkRegion2D
	BoundingBox() corefoundation.CGRect
	SetBoundingBox(value corefoundation.CGRect)
	Landmarks() IVNFaceLandmarks2D
	SetLandmarks(value IVNFaceLandmarks2D)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FaceLandmarks2D */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FaceLandmarks2D */
// Alloc allocates a new instance without initialization.
func (fc _FaceLandmarks2DClass) Alloc() FaceLandmarks2D {
	rv := objc.Send[FaceLandmarks2D](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FaceLandmarks2DClass) New() FaceLandmarks2D {
	rv := objc.Send[FaceLandmarks2D](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FaceLandmarks2D) Init() FaceLandmarks2D {
	rv := objc.Send[FaceLandmarks2D](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FaceLandmarks2D) Autorelease() FaceLandmarks2D {
	rv := objc.Send[FaceLandmarks2D](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFaceLandmarks2D creates a new FaceLandmarks2D instance.
func NewFaceLandmarks2D() FaceLandmarks2D {
	return getFaceLandmarks2DClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FaceLandmarks2D */
// A collection of facial features that a request detects.
//
// This class represents the set of all detectable 2D face landmarks and regions, exposed as properties. The coordinates of the face landmarks are normalized to the dimensions of the face observation’s , with the origin at the bounding box’s lower-left corner. Use the function to convert normalized face landmark points into absolute points within the image’s coordinate system.


// A collection of facial features that a request detects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D
type FaceLandmarks2D struct {
	FaceLandmarks
}

// FaceLandmarks2DFrom constructs a [FaceLandmarks2D] from an unsafe.Pointer.
//
// A collection of facial features that a request detects.
func FaceLandmarks2DFrom(ptr unsafe.Pointer) FaceLandmarks2D {
	return FaceLandmarks2D{
		FaceLandmarks: FaceLandmarksFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FaceLandmarks2D *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FaceLandmarks2D */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FaceLandmarks2D */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FaceLandmarks2D */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FaceLandmarks2D */

// The region containing all face landmark points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/allPoints
func (f_ FaceLandmarks2D) AllPoints() IVNFaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("allPoints"))
	return rv
}/* debug [instance_properties/getter]: allPoints */


// The region containing points that trace the face contour from the left cheek, over the chin, to the right cheek.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/faceContour
func (f_ FaceLandmarks2D) FaceContour() IVNFaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("faceContour"))
	return rv
}/* debug [instance_properties/getter]: faceContour */


// The region containing points that outline the space between the lips.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/innerLips
func (f_ FaceLandmarks2D) InnerLips() IVNFaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("innerLips"))
	return rv
}/* debug [instance_properties/getter]: innerLips */


// The region containing points that outline the left eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/leftEye
func (f_ FaceLandmarks2D) LeftEye() IVNFaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("leftEye"))
	return rv
}/* debug [instance_properties/getter]: leftEye */


// The region containing points that trace the left eyebrow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/leftEyebrow
func (f_ FaceLandmarks2D) LeftEyebrow() IVNFaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("leftEyebrow"))
	return rv
}/* debug [instance_properties/getter]: leftEyebrow */


// The region containing the point where the left pupil is located.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/leftPupil
func (f_ FaceLandmarks2D) LeftPupil() IVNFaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("leftPupil"))
	return rv
}/* debug [instance_properties/getter]: leftPupil */


// The region containing points that trace a vertical line down the center of the face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/medianLine
func (f_ FaceLandmarks2D) MedianLine() IVNFaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("medianLine"))
	return rv
}/* debug [instance_properties/getter]: medianLine */


// The region containing points that outline the nose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/nose
func (f_ FaceLandmarks2D) Nose() IVNFaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("nose"))
	return rv
}/* debug [instance_properties/getter]: nose */


// The region containing points that trace the center crest of the nose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/noseCrest
func (f_ FaceLandmarks2D) NoseCrest() IVNFaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("noseCrest"))
	return rv
}/* debug [instance_properties/getter]: noseCrest */


// The region containing points that outline the outside of the lips.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/outerLips
func (f_ FaceLandmarks2D) OuterLips() IVNFaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("outerLips"))
	return rv
}/* debug [instance_properties/getter]: outerLips */


// The region containing points that outline the right eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/rightEye
func (f_ FaceLandmarks2D) RightEye() IVNFaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("rightEye"))
	return rv
}/* debug [instance_properties/getter]: rightEye */


// The region containing points that trace the right eyebrow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/rightEyebrow
func (f_ FaceLandmarks2D) RightEyebrow() IVNFaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("rightEyebrow"))
	return rv
}/* debug [instance_properties/getter]: rightEyebrow */


// The region containing the point where the right pupil is located.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/rightPupil
func (f_ FaceLandmarks2D) RightPupil() IVNFaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("rightPupil"))
	return rv
}/* debug [instance_properties/getter]: rightPupil */


// The bounding box of the object that the request detects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/boundingbox
func (f_ FaceLandmarks2D) BoundingBox() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](f_.ID, objc.Sel("boundingBox"))
	return rv
}/* debug [instance_properties/getter]: boundingBox */


// The bounding box of the object that the request detects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/boundingbox
func (f_ FaceLandmarks2D) SetBoundingBox(value corefoundation.CGRect) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBoundingBox:"), value)
}/* debug [instance_properties/setter]: boundingBox */


// The facial features of the detected face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/landmarks
func (f_ FaceLandmarks2D) Landmarks() IVNFaceLandmarks2D {
	rv := objc.Send[FaceLandmarks2D](f_.ID, objc.Sel("landmarks"))
	return rv
}/* debug [instance_properties/getter]: landmarks */


// The facial features of the detected face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/landmarks
func (f_ FaceLandmarks2D) SetLandmarks(value IVNFaceLandmarks2D) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLandmarks:"), value)
}/* debug [instance_properties/setter]: landmarks */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNFaceLandmarks2D */



