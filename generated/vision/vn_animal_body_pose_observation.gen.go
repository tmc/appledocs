// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AnimalBodyPoseObservation] class.
var (
	AnimalBodyPoseObservationClass     _AnimalBodyPoseObservationClass
	AnimalBodyPoseObservationClassOnce sync.Once
)

func getAnimalBodyPoseObservationClass() _AnimalBodyPoseObservationClass {
	AnimalBodyPoseObservationClassOnce.Do(func() {
		AnimalBodyPoseObservationClass = _AnimalBodyPoseObservationClass{objc.GetClass("VNAnimalBodyPoseObservation")}
	})
	return AnimalBodyPoseObservationClass
}

type _AnimalBodyPoseObservationClass struct {
	class objc.Class
}

// An interface definition for the [AnimalBodyPoseObservation] class.
type IAnimalBodyPoseObservation interface {
	IRecognizedPointsObservation
	RecognizedPointForJointNameError(jointName unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	RecognizedPointsForJointsGroupNameError(jointsGroupName unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
}

// An observation that provides the animal body points the analysis recognizes.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation
type AnimalBodyPoseObservation struct {
	RecognizedPointsObservation
}

// AnimalBodyPoseObservationFrom constructs a [AnimalBodyPoseObservation] from an unsafe.Pointer.
//
// An observation that provides the animal body points the analysis recognizes.
func AnimalBodyPoseObservationFrom(ptr unsafe.Pointer) AnimalBodyPoseObservation {
	return AnimalBodyPoseObservation{
		RecognizedPointsObservation: RecognizedPointsObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AnimalBodyPoseObservationClass) Alloc() AnimalBodyPoseObservation {
	rv := objc.Send[AnimalBodyPoseObservation](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AnimalBodyPoseObservationClass) New() AnimalBodyPoseObservation {
	rv := objc.Send[AnimalBodyPoseObservation](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AnimalBodyPoseObservation) Init() AnimalBodyPoseObservation {
	rv := objc.Send[AnimalBodyPoseObservation](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AnimalBodyPoseObservation) Autorelease() AnimalBodyPoseObservation {
	rv := objc.Send[AnimalBodyPoseObservation](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAnimalBodyPoseObservation creates a new AnimalBodyPoseObservation instance.
func NewAnimalBodyPoseObservation() AnimalBodyPoseObservation {
	return getAnimalBodyPoseObservationClass().New()
}


// Returns the point for a joint name the observation recognizes.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation/recognizedPoint(_:)
func (a_ AnimalBodyPoseObservation) RecognizedPointForJointNameError(jointName unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("recognizedPointForJointName:error:"), jointName, error_)
	return rv
}

// Returns the points for a joint group name the observation recognizes.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation/recognizedPoints(_:)
func (a_ AnimalBodyPoseObservation) RecognizedPointsForJointsGroupNameError(jointsGroupName unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("recognizedPointsForJointsGroupName:error:"), jointsGroupName, error_)
	return rv
}



