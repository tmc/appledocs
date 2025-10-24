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
	// properties:
	AvailableJointNames() unsafe.Pointer
	SetAvailableJointNames(value unsafe.Pointer)
	AvailableJointsGroupNames() unsafe.Pointer
	SetAvailableJointsGroupNames(value unsafe.Pointer)
	// methods:
}

// An observation that provides the body points the analysis recognized.


// An observation that provides the body points the analysis recognized.
//
// [Full Topic]
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



// The names of the available joints in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservation/availablejointnames
func (h_ HumanBodyPoseObservation) AvailableJointNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("availableJointNames"))
	return rv
}


// The names of the available joints in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservation/availablejointnames
func (h_ HumanBodyPoseObservation) SetAvailableJointNames(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAvailableJointNames:"), value)
}


// The available joint group names in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservation/availablejointsgroupnames
func (h_ HumanBodyPoseObservation) AvailableJointsGroupNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("availableJointsGroupNames"))
	return rv
}


// The available joint group names in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyposeobservation/availablejointsgroupnames
func (h_ HumanBodyPoseObservation) SetAvailableJointsGroupNames(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAvailableJointsGroupNames:"), value)
}



