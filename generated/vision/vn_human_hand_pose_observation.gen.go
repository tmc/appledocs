// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [HumanHandPoseObservation] class.
var (
	HumanHandPoseObservationClass     _HumanHandPoseObservationClass
	HumanHandPoseObservationClassOnce sync.Once
)

func getHumanHandPoseObservationClass() _HumanHandPoseObservationClass {
	HumanHandPoseObservationClassOnce.Do(func() {
		HumanHandPoseObservationClass = _HumanHandPoseObservationClass{objc.GetClass("VNHumanHandPoseObservation")}
	})
	return HumanHandPoseObservationClass
}

type _HumanHandPoseObservationClass struct {
	class objc.Class
}





// An interface definition for the [HumanHandPoseObservation] class.
type IHumanHandPoseObservation interface {
	IRecognizedPointsObservation
	

	// properties:
	AvailableJointNames() []string
	AvailableJointsGroupNames() []string
	Chirality() Chirality


	

	// methods:
	RecognizedPointForJointNameError(jointName HumanHandPoseObservationJointName /* typedef */, error_ objectivec.IObject) IRecognizedPoint
	RecognizedPointsForJointsGroupNameError(jointsGroupName HumanHandPoseObservationJointsGroupName /* typedef */, error_ objectivec.IObject) foundation.IDictionary


}





// Alloc allocates a new instance without initialization.
func (hc _HumanHandPoseObservationClass) Alloc() HumanHandPoseObservation {
	rv := objc.Send[HumanHandPoseObservation](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HumanHandPoseObservationClass) New() HumanHandPoseObservation {
	rv := objc.Send[HumanHandPoseObservation](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HumanHandPoseObservation) Init() HumanHandPoseObservation {
	rv := objc.Send[HumanHandPoseObservation](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HumanHandPoseObservation) Autorelease() HumanHandPoseObservation {
	rv := objc.Send[HumanHandPoseObservation](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHumanHandPoseObservation creates a new HumanHandPoseObservation instance.
func NewHumanHandPoseObservation() HumanHandPoseObservation {
	return getHumanHandPoseObservationClass().New()
}





// An observation that provides the hand points the analysis recognized.


// An observation that provides the hand points the analysis recognized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation
type HumanHandPoseObservation struct {
	RecognizedPointsObservation
}

// HumanHandPoseObservationFrom constructs a [HumanHandPoseObservation] from an unsafe.Pointer.
//
// An observation that provides the hand points the analysis recognized.
func HumanHandPoseObservationFrom(ptr unsafe.Pointer) HumanHandPoseObservation {
	return HumanHandPoseObservation{
		RecognizedPointsObservation: RecognizedPointsObservationFrom(ptr),
	}
}




















// Retrieves the recognized point for a joint name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation/recognizedPoint(_:)
func (h_ HumanHandPoseObservation) RecognizedPointForJointNameError(jointName HumanHandPoseObservationJointName /* typedef */, error_ objectivec.IObject) IRecognizedPoint {
	rv := objc.Send[RecognizedPoint](h_.ID, objc.Sel("recognizedPointForJointName:error:"), jointName, error_)
	return rv
}


// Retrieves the recognized points associated with the joint group name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation/recognizedPoints(_:)
func (h_ HumanHandPoseObservation) RecognizedPointsForJointsGroupNameError(jointsGroupName HumanHandPoseObservationJointsGroupName /* typedef */, error_ objectivec.IObject) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](h_.ID, objc.Sel("recognizedPointsForJointsGroupName:error:"), jointsGroupName, error_)
	return rv
}







// The names of the available joints in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation/availableJointNames
func (h_ HumanHandPoseObservation) AvailableJointNames() []string {
	rv := objc.Send[[]string](h_.ID, objc.Sel("availableJointNames"))
	return rv
}


// The joint group names available in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation/availableJointsGroupNames
func (h_ HumanHandPoseObservation) AvailableJointsGroupNames() []string {
	rv := objc.Send[[]string](h_.ID, objc.Sel("availableJointsGroupNames"))
	return rv
}


// The chirality, or handedness, of a pose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanHandPoseObservation/chirality
func (h_ HumanHandPoseObservation) Chirality() Chirality {
	rv := objc.Send[Chirality](h_.ID, objc.Sel("chirality"))
	return rv
}








