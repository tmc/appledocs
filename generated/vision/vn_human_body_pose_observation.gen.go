// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HumanBodyPoseObservation] class.
var (
	HumanBodyPoseObservationClass     _HumanBodyPoseObservationClass
	HumanBodyPoseObservationClassOnce sync.Once
)

func getHumanBodyPoseObservationClass() _HumanBodyPoseObservationClass {
	HumanBodyPoseObservationClassOnce.Do(func() {
		HumanBodyPoseObservationClass = _HumanBodyPoseObservationClass{objc.GetClass("VNHumanBodyPoseObservation")}
	})
	return HumanBodyPoseObservationClass
}

type _HumanBodyPoseObservationClass struct {
	class objc.Class
}

// An interface definition for the [HumanBodyPoseObservation] class.
type IHumanBodyPoseObservation interface {
	IRecognizedPointsObservation
	RecognizedPointForJointNameError(jointName IHumanBodyPoseObservationJointName, error_ unsafe.Pointer) RecognizedPoint
	RecognizedPointsForJointsGroupNameError(jointsGroupName IHumanBodyPoseObservationJointsGroupName, error_ unsafe.Pointer) unsafe.Pointer
}

// An observation that provides the body points the analysis recognized.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation
type HumanBodyPoseObservation struct {
	RecognizedPointsObservation
}

// HumanBodyPoseObservationFrom constructs a [HumanBodyPoseObservation] from an unsafe.Pointer.
//
// An observation that provides the body points the analysis recognized.
func HumanBodyPoseObservationFrom(ptr unsafe.Pointer) HumanBodyPoseObservation {
	return HumanBodyPoseObservation{
		RecognizedPointsObservation: RecognizedPointsObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HumanBodyPoseObservationClass) Alloc() HumanBodyPoseObservation {
	rv := objc.Send[HumanBodyPoseObservation](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HumanBodyPoseObservationClass) New() HumanBodyPoseObservation {
	rv := objc.Send[HumanBodyPoseObservation](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HumanBodyPoseObservation) Init() HumanBodyPoseObservation {
	rv := objc.Send[HumanBodyPoseObservation](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HumanBodyPoseObservation) Autorelease() HumanBodyPoseObservation {
	rv := objc.Send[HumanBodyPoseObservation](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHumanBodyPoseObservation creates a new HumanBodyPoseObservation instance.
func NewHumanBodyPoseObservation() HumanBodyPoseObservation {
	return getHumanBodyPoseObservationClass().New()
}


// Retrieves the recognized point for a joint name.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation/recognizedPoint(_:)
func (h_ HumanBodyPoseObservation) RecognizedPointForJointNameError(jointName IHumanBodyPoseObservationJointName, error_ unsafe.Pointer) RecognizedPoint {
	rv := objc.Send[RecognizedPoint](h_.ID, objc.Sel("recognizedPointForJointName:error:"), jointName, error_)
	return rv
}

// Retrieves the recognized points associated with the joint group name.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation/recognizedPoints(_:)
func (h_ HumanBodyPoseObservation) RecognizedPointsForJointsGroupNameError(jointsGroupName IHumanBodyPoseObservationJointsGroupName, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("recognizedPointsForJointsGroupName:error:"), jointsGroupName, error_)
	return rv
}

// The names of the available joints in the observation.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation/availableJointNames
func (h_ HumanBodyPoseObservation) AvailableJointNames() []string {
	rv := objc.Send[[]string](h_.ID, objc.Sel("availableJointNames"))
	return rv
}

// The available joint group names in the observation.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation/availableJointsGroupNames
func (h_ HumanBodyPoseObservation) AvailableJointsGroupNames() []string {
	rv := objc.Send[[]string](h_.ID, objc.Sel("availableJointsGroupNames"))
	return rv
}



