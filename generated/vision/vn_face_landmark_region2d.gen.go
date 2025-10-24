// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [FaceLandmarkRegion2D] class.
var (
	FaceLandmarkRegion2DClass     _FaceLandmarkRegion2DClass
	FaceLandmarkRegion2DClassOnce sync.Once
)

func getFaceLandmarkRegion2DClass() _FaceLandmarkRegion2DClass {
	FaceLandmarkRegion2DClassOnce.Do(func() {
		FaceLandmarkRegion2DClass = _FaceLandmarkRegion2DClass{objc.GetClass("VNFaceLandmarkRegion2D")}
	})
	return FaceLandmarkRegion2DClass
}

type _FaceLandmarkRegion2DClass struct {
	class objc.Class
}





// An interface definition for the [FaceLandmarkRegion2D] class.
type IFaceLandmarkRegion2D interface {
	IFaceLandmarkRegion
	

	// properties:
	NormalizedPoints() corefoundation.CGPoint
	PointsClassification() PointsClassification
	PrecisionEstimatesPerPoint() []foundation.Number
	Landmarks() IVNFaceLandmarks2D
	SetLandmarks(value IVNFaceLandmarks2D)


	

	// methods:
	PointsInImageOfSize(imageSize corefoundation.CGSize) corefoundation.CGPoint


}





// Alloc allocates a new instance without initialization.
func (fc _FaceLandmarkRegion2DClass) Alloc() FaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FaceLandmarkRegion2DClass) New() FaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FaceLandmarkRegion2D) Init() FaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FaceLandmarkRegion2D) Autorelease() FaceLandmarkRegion2D {
	rv := objc.Send[FaceLandmarkRegion2D](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFaceLandmarkRegion2D creates a new FaceLandmarkRegion2D instance.
func NewFaceLandmarkRegion2D() FaceLandmarkRegion2D {
	return getFaceLandmarkRegion2DClass().New()
}





// 2D geometry information for a specific facial feature.
//
// This class represents the set of all facial landmark regions in 2D, exposed as properties.


// 2D geometry information for a specific facial feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion2D
type FaceLandmarkRegion2D struct {
	FaceLandmarkRegion
}

// FaceLandmarkRegion2DFrom constructs a [FaceLandmarkRegion2D] from an unsafe.Pointer.
//
// 2D geometry information for a specific facial feature.
func FaceLandmarkRegion2DFrom(ptr unsafe.Pointer) FaceLandmarkRegion2D {
	return FaceLandmarkRegion2D{
		FaceLandmarkRegion: FaceLandmarkRegionFrom(ptr),
	}
}




















// A buffer in memory containing landmark points in the coordinate space of the specified image size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion2D/pointsInImageOfSize:
func (f_ FaceLandmarkRegion2D) PointsInImageOfSize(imageSize corefoundation.CGSize) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](f_.ID, objc.Sel("pointsInImageOfSize:"), imageSize)
	return rv
}







// A buffer in memory containing normalized landmark points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion2D/normalizedPoints-1o38f
func (f_ FaceLandmarkRegion2D) NormalizedPoints() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](f_.ID, objc.Sel("normalizedPoints"))
	return rv
}


// An enumeration that describes how to interpret the points the region provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion2D/pointsClassification
func (f_ FaceLandmarkRegion2D) PointsClassification() PointsClassification {
	rv := objc.Send[PointsClassification](f_.ID, objc.Sel("pointsClassification"))
	return rv
}


// An array of precision estimates for each landmark point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarkRegion2D/precisionEstimatesPerPoint-3kx5a
func (f_ FaceLandmarkRegion2D) PrecisionEstimatesPerPoint() []foundation.Number {
	rv := objc.Send[[]foundation.Number](f_.ID, objc.Sel("precisionEstimatesPerPoint"))
	return rv
}


// The facial features of the detected face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/landmarks
func (f_ FaceLandmarkRegion2D) Landmarks() IVNFaceLandmarks2D {
	rv := objc.Send[FaceLandmarks2D](f_.ID, objc.Sel("landmarks"))
	return rv
}


// The facial features of the detected face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/landmarks
func (f_ FaceLandmarkRegion2D) SetLandmarks(value IVNFaceLandmarks2D) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLandmarks:"), value)
}








