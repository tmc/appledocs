// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [FaceLandmarks] class.
var (
	FaceLandmarksClass     _FaceLandmarksClass
	FaceLandmarksClassOnce sync.Once
)

func getFaceLandmarksClass() _FaceLandmarksClass {
	FaceLandmarksClassOnce.Do(func() {
		FaceLandmarksClass = _FaceLandmarksClass{objc.GetClass("VNFaceLandmarks")}
	})
	return FaceLandmarksClass
}

type _FaceLandmarksClass struct {
	class objc.Class
}





// An interface definition for the [FaceLandmarks] class.
type IFaceLandmarks interface {
	objectivec.IObject
	

	// properties:
	Confidence() Confidence /* typedef */
	Landmarks() IVNFaceLandmarks2D
	SetLandmarks(value IVNFaceLandmarks2D)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (fc _FaceLandmarksClass) Alloc() FaceLandmarks {
	rv := objc.Send[FaceLandmarks](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FaceLandmarksClass) New() FaceLandmarks {
	rv := objc.Send[FaceLandmarks](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FaceLandmarks) Init() FaceLandmarks {
	rv := objc.Send[FaceLandmarks](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FaceLandmarks) Autorelease() FaceLandmarks {
	rv := objc.Send[FaceLandmarks](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFaceLandmarks creates a new FaceLandmarks instance.
func NewFaceLandmarks() FaceLandmarks {
	return getFaceLandmarksClass().New()
}





// The abstract superclass for containers of face landmark information.
//
// This class represents the set of all detectable facial landmarks and regions, exposed as properties.


// The abstract superclass for containers of face landmark information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks
type FaceLandmarks struct {
	objectivec.Object
}

// FaceLandmarksFrom constructs a [FaceLandmarks] from an unsafe.Pointer.
//
// The abstract superclass for containers of face landmark information.
func FaceLandmarksFrom(ptr unsafe.Pointer) FaceLandmarks {
	return FaceLandmarks{objectivec.Object{objc.ID(ptr)}}
}

























// A confidence estimate for the detected landmarks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceLandmarks/confidence
func (f_ FaceLandmarks) Confidence() Confidence /* typedef */ {
	rv := objc.Send[float32](f_.ID, objc.Sel("confidence"))
	return rv
}


// The facial features of the detected face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/landmarks
func (f_ FaceLandmarks) Landmarks() IVNFaceLandmarks2D {
	rv := objc.Send[FaceLandmarks2D](f_.ID, objc.Sel("landmarks"))
	return rv
}


// The facial features of the detected face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/landmarks
func (f_ FaceLandmarks) SetLandmarks(value IVNFaceLandmarks2D) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLandmarks:"), value)
}








