// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [FaceObservation] class.
var (
	FaceObservationClass     _FaceObservationClass
	FaceObservationClassOnce sync.Once
)

func getFaceObservationClass() _FaceObservationClass {
	FaceObservationClassOnce.Do(func() {
		FaceObservationClass = _FaceObservationClass{objc.GetClass("VNFaceObservation")}
	})
	return FaceObservationClass
}

type _FaceObservationClass struct {
	class objc.Class
}

// An interface definition for the [FaceObservation] class.
type IFaceObservation interface {
	IDetectedObjectObservation
}

// Face or facial-feature information that an image analysis request detects.
//
// This type of observation results from a . It contains information about facial landmarks and regions it finds in the image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceObservation
type FaceObservation struct {
	DetectedObjectObservation
}

// FaceObservationFrom constructs a [FaceObservation] from an unsafe.Pointer.
//
// Face or facial-feature information that an image analysis request detects.
func FaceObservationFrom(ptr unsafe.Pointer) FaceObservation {
	return FaceObservation{
		DetectedObjectObservation: DetectedObjectObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FaceObservationClass) Alloc() FaceObservation {
	rv := objc.Send[FaceObservation](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FaceObservationClass) New() FaceObservation {
	rv := objc.Send[FaceObservation](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FaceObservation) Init() FaceObservation {
	rv := objc.Send[FaceObservation](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FaceObservation) Autorelease() FaceObservation {
	rv := objc.Send[FaceObservation](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFaceObservation creates a new FaceObservation instance.
func NewFaceObservation() FaceObservation {
	return getFaceObservationClass().New()
}




// Creates an observation that contains the roll and yaw of the face.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceObservation/init(requestRevision:boundingBox:roll:yaw:)
func NewFaceObservationWithRequestRevisionBoundingBoxRollYaw(requestRevision uint, boundingBox coregraphics.CGRect, roll unsafe.Pointer, yaw unsafe.Pointer) FaceObservation {
	rv := objc.Send[FaceObservation](objc.ID(getFaceObservationClass().class), objc.Sel("faceObservationWithRequestRevision:boundingBox:roll:yaw:"), requestRevision, boundingBox, roll, yaw)
	return rv
}



// Creates an observation that contains the roll, yaw, and pitch of the face.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceObservation/init(requestRevision:boundingBox:roll:yaw:pitch:)
func NewFaceObservationWithRequestRevisionBoundingBoxRollYawPitch(requestRevision uint, boundingBox coregraphics.CGRect, roll unsafe.Pointer, yaw unsafe.Pointer, pitch unsafe.Pointer) FaceObservation {
	rv := objc.Send[FaceObservation](objc.ID(getFaceObservationClass().class), objc.Sel("faceObservationWithRequestRevision:boundingBox:roll:yaw:pitch:"), requestRevision, boundingBox, roll, yaw, pitch)
	return rv
}


// Creates an observation that contains the roll and yaw of the face.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceObservation/init(requestRevision:boundingBox:roll:yaw:)
func (fc _FaceObservationClass) FaceObservationWithRequestRevisionBoundingBoxRollYaw(requestRevision uint, boundingBox coregraphics.CGRect, roll unsafe.Pointer, yaw unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("faceObservationWithRequestRevision:boundingBox:roll:yaw:"), requestRevision, boundingBox, roll, yaw)
	return rv
}

// Creates an observation that contains the roll, yaw, and pitch of the face.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceObservation/init(requestRevision:boundingBox:roll:yaw:pitch:)
func (fc _FaceObservationClass) FaceObservationWithRequestRevisionBoundingBoxRollYawPitch(requestRevision uint, boundingBox coregraphics.CGRect, roll unsafe.Pointer, yaw unsafe.Pointer, pitch unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("faceObservationWithRequestRevision:boundingBox:roll:yaw:pitch:"), requestRevision, boundingBox, roll, yaw, pitch)
	return rv
}

// A value that indicates the quality of the face capture.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceObservation/faceCaptureQuality-2o4xv
func (f_ FaceObservation) FaceCaptureQuality() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("faceCaptureQuality"))
	return rv
}

// The facial features of the detected face.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFaceObservation/landmarks
func (f_ FaceObservation) Landmarks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("landmarks"))
	return rv
}


