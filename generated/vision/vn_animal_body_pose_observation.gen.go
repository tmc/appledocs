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
	// properties:
	AvailableJointGroupNames() unsafe.Pointer
	SetAvailableJointGroupNames(value unsafe.Pointer)
	AvailableJointNames() unsafe.Pointer
	SetAvailableJointNames(value unsafe.Pointer)
	// methods:
}

// An observation that provides the animal body points the analysis recognizes.


// An observation that provides the animal body points the analysis recognizes.
//
// [Full Topic]
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



// The available joint group names in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnanimalbodyposeobservation/availablejointgroupnames
func (a_ AnimalBodyPoseObservation) AvailableJointGroupNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("availableJointGroupNames"))
	return rv
}


// The available joint group names in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnanimalbodyposeobservation/availablejointgroupnames
func (a_ AnimalBodyPoseObservation) SetAvailableJointGroupNames(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableJointGroupNames:"), value)
}


// The names of the available joints in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnanimalbodyposeobservation/availablejointnames
func (a_ AnimalBodyPoseObservation) AvailableJointNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("availableJointNames"))
	return rv
}


// The names of the available joints in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnanimalbodyposeobservation/availablejointnames
func (a_ AnimalBodyPoseObservation) SetAvailableJointNames(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableJointNames:"), value)
}



