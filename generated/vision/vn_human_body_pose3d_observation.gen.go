// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	AvailableJointNames() []string
	AvailableJointsGroupNames() []string
	BodyHeight() float32
	CameraOriginMatrix() objectivec.IObject
	HeightEstimation() HumanBodyPose3DObservationHeightEstimation


	

	// methods:
	GetCameraRelativePositionForJointNameError(modelPositionOut objectivec.IObject, jointName HumanBodyPose3DObservationJointName /* typedef */, error_ objectivec.IObject) bool
	ParentJointNameForJointName(jointName HumanBodyPose3DObservationJointName /* typedef */) HumanBodyPose3DObservationJointName /* typedef */
	PointInImageForJointNameError(jointName HumanBodyPose3DObservationJointName /* typedef */, error_ objectivec.IObject) IPoint
	RecognizedPointForJointNameError(jointName HumanBodyPose3DObservationJointName /* typedef */, error_ objectivec.IObject) IHumanBodyRecognizedPoint3D
	RecognizedPointsForJointsGroupNameError(jointsGroupName HumanBodyPose3DObservationJointsGroupName /* typedef */, error_ objectivec.IObject) foundation.IDictionary


}





// Alloc allocates a new instance without initialization.
func (hc _HumanBodyPose3DObservationClass) Alloc() HumanBodyPose3DObservation {
	rv := objc.Send[HumanBodyPose3DObservation](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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




















// Gets a position relative to the camera for the body joint you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/getCameraRelativePosition:forJointName:error:
func (h_ HumanBodyPose3DObservation) GetCameraRelativePositionForJointNameError(modelPositionOut objectivec.IObject, jointName HumanBodyPose3DObservationJointName /* typedef */, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("getCameraRelativePosition:forJointName:error:"), modelPositionOut, jointName, error_)
	return rv
}


// Returns the parent joint of the joint name you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/parentJointName(_:)
func (h_ HumanBodyPose3DObservation) ParentJointNameForJointName(jointName HumanBodyPose3DObservationJointName /* typedef */) HumanBodyPose3DObservationJointName /* typedef */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("parentJointNameForJointName:"), jointName)
	return rv
}


// Returns a 2D point for the joint name you specify, relative to the input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/pointInImage(_:)
func (h_ HumanBodyPose3DObservation) PointInImageForJointNameError(jointName HumanBodyPose3DObservationJointName /* typedef */, error_ objectivec.IObject) IPoint {
	rv := objc.Send[Point](h_.ID, objc.Sel("pointInImageForJointName:error:"), jointName, error_)
	return rv
}


// Returns the point for a joint name that the observation recognizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/recognizedPoint(_:)
func (h_ HumanBodyPose3DObservation) RecognizedPointForJointNameError(jointName HumanBodyPose3DObservationJointName /* typedef */, error_ objectivec.IObject) IHumanBodyRecognizedPoint3D {
	rv := objc.Send[HumanBodyRecognizedPoint3D](h_.ID, objc.Sel("recognizedPointForJointName:error:"), jointName, error_)
	return rv
}


// Returns a collection of points for the group name you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/recognizedPoints(_:)
func (h_ HumanBodyPose3DObservation) RecognizedPointsForJointsGroupNameError(jointsGroupName HumanBodyPose3DObservationJointsGroupName /* typedef */, error_ objectivec.IObject) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](h_.ID, objc.Sel("recognizedPointsForJointsGroupName:error:"), jointsGroupName, error_)
	return rv
}







// The names of the available joints in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/availableJointNames
func (h_ HumanBodyPose3DObservation) AvailableJointNames() []string {
	rv := objc.Send[[]string](h_.ID, objc.Sel("availableJointNames"))
	return rv
}


// The available joint group names in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/availableJointsGroupNames
func (h_ HumanBodyPose3DObservation) AvailableJointsGroupNames() []string {
	rv := objc.Send[[]string](h_.ID, objc.Sel("availableJointsGroupNames"))
	return rv
}


// The estimated human body height, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/bodyHeight
func (h_ HumanBodyPose3DObservation) BodyHeight() float32 {
	rv := objc.Send[float32](h_.ID, objc.Sel("bodyHeight"))
	return rv
}


// A transform from the skeleton hip to the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/cameraOriginMatrix
func (h_ HumanBodyPose3DObservation) CameraOriginMatrix() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](h_.ID, objc.Sel("cameraOriginMatrix"))
	return rv
}


// The technique the framework uses to estimate body height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/heightEstimation-swift.property
func (h_ HumanBodyPose3DObservation) HeightEstimation() HumanBodyPose3DObservationHeightEstimation {
	rv := objc.Send[HumanBodyPose3DObservationHeightEstimation](h_.ID, objc.Sel("heightEstimation"))
	return rv
}








