// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HumanBodyPose3DObservation] class.
var (
	HumanBodyPose3DObservationClass     _HumanBodyPose3DObservationClass
	HumanBodyPose3DObservationClassOnce sync.Once
)

func getHumanBodyPose3DObservationClass() _HumanBodyPose3DObservationClass {
	HumanBodyPose3DObservationClassOnce.Do(func() {
		HumanBodyPose3DObservationClass = _HumanBodyPose3DObservationClass{objc.GetClass("VNHumanBodyPose3DObservation")}
	})
	return HumanBodyPose3DObservationClass
}

type _HumanBodyPose3DObservationClass struct {
	class objc.Class
}

// An interface definition for the [HumanBodyPose3DObservation] class.
type IHumanBodyPose3DObservation interface {
	IRecognizedPoints3DObservation
	// properties:
	AvailableJointNames() unsafe.Pointer
	SetAvailableJointNames(value unsafe.Pointer)
	AvailableJointsGroupNames() unsafe.Pointer
	SetAvailableJointsGroupNames(value unsafe.Pointer)
	BodyHeight() float32
	SetBodyHeight(value float32)
	CameraOriginMatrix() unsafe.Pointer
	SetCameraOriginMatrix(value unsafe.Pointer)
	HeightEstimation() unsafe.Pointer
	SetHeightEstimation(value unsafe.Pointer)
	// methods:
}

// An observation that provides the 3D body points the request recognizes.


// An observation that provides the 3D body points the request recognizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation
type HumanBodyPose3DObservation struct {
	RecognizedPoints3DObservation
}

// HumanBodyPose3DObservationFrom constructs a [HumanBodyPose3DObservation] from an unsafe.Pointer.
//
// An observation that provides the 3D body points the request recognizes.
func HumanBodyPose3DObservationFrom(ptr unsafe.Pointer) HumanBodyPose3DObservation {
	return HumanBodyPose3DObservation{
		RecognizedPoints3DObservation: RecognizedPoints3DObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HumanBodyPose3DObservationClass) Alloc() HumanBodyPose3DObservation {
	rv := objc.Send[HumanBodyPose3DObservation](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HumanBodyPose3DObservationClass) New() HumanBodyPose3DObservation {
	rv := objc.Send[HumanBodyPose3DObservation](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HumanBodyPose3DObservation) Init() HumanBodyPose3DObservation {
	rv := objc.Send[HumanBodyPose3DObservation](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HumanBodyPose3DObservation) Autorelease() HumanBodyPose3DObservation {
	rv := objc.Send[HumanBodyPose3DObservation](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHumanBodyPose3DObservation creates a new HumanBodyPose3DObservation instance.
func NewHumanBodyPose3DObservation() HumanBodyPose3DObservation {
	return getHumanBodyPose3DObservationClass().New()
}



// The names of the available joints in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodypose3dobservation/availablejointnames
func (h_ HumanBodyPose3DObservation) AvailableJointNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("availableJointNames"))
	return rv
}


// The names of the available joints in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodypose3dobservation/availablejointnames
func (h_ HumanBodyPose3DObservation) SetAvailableJointNames(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAvailableJointNames:"), value)
}


// The available joint group names in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodypose3dobservation/availablejointsgroupnames
func (h_ HumanBodyPose3DObservation) AvailableJointsGroupNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("availableJointsGroupNames"))
	return rv
}


// The available joint group names in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodypose3dobservation/availablejointsgroupnames
func (h_ HumanBodyPose3DObservation) SetAvailableJointsGroupNames(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAvailableJointsGroupNames:"), value)
}


// The estimated human body height, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodypose3dobservation/bodyheight
func (h_ HumanBodyPose3DObservation) BodyHeight() float32 {
	rv := objc.Send[float32](h_.ID, objc.Sel("bodyHeight"))
	return rv
}


// The estimated human body height, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodypose3dobservation/bodyheight
func (h_ HumanBodyPose3DObservation) SetBodyHeight(value float32) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setBodyHeight:"), value)
}


// A transform from the skeleton hip to the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodypose3dobservation/cameraoriginmatrix
func (h_ HumanBodyPose3DObservation) CameraOriginMatrix() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("cameraOriginMatrix"))
	return rv
}


// A transform from the skeleton hip to the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodypose3dobservation/cameraoriginmatrix
func (h_ HumanBodyPose3DObservation) SetCameraOriginMatrix(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCameraOriginMatrix:"), value)
}


// The technique the framework uses to estimate body height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodypose3dobservation/heightestimation-swift.property
func (h_ HumanBodyPose3DObservation) HeightEstimation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("heightEstimation"))
	return rv
}


// The technique the framework uses to estimate body height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodypose3dobservation/heightestimation-swift.property
func (h_ HumanBodyPose3DObservation) SetHeightEstimation(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setHeightEstimation:"), value)
}



