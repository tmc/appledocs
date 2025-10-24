// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNHumanBodyPoseObservation */


/* debug [class_header]: Header for VNHumanBodyPoseObservation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HumanBodyPoseObservation */
// An interface definition for the [HumanBodyPoseObservation] class.
type IHumanBodyPoseObservation interface {
	IRecognizedPointsObservation
	
/* debug [class_interface_properties]: Properties for HumanBodyPoseObservation */
	// properties:
	AvailableJointNames() []string
	AvailableJointsGroupNames() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HumanBodyPoseObservation */
	// methods:
	RecognizedPointForJointNameError(jointName HumanBodyPoseObservationJointName /* typedef */, error_ objectivec.IObject) IRecognizedPoint
	RecognizedPointsForJointsGroupNameError(jointsGroupName HumanBodyPoseObservationJointsGroupName /* typedef */, error_ objectivec.IObject) foundation.IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HumanBodyPoseObservation */
// Alloc allocates a new instance without initialization.
func (hc _HumanBodyPoseObservationClass) Alloc() HumanBodyPoseObservation {
	rv := objc.Send[HumanBodyPoseObservation](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HumanBodyPoseObservation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HumanBodyPoseObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HumanBodyPoseObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HumanBodyPoseObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HumanBodyPoseObservation */

// Retrieves the recognized point for a joint name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation/recognizedPoint(_:)
func (h_ HumanBodyPoseObservation) RecognizedPointForJointNameError(jointName HumanBodyPoseObservationJointName /* typedef */, error_ objectivec.IObject) IRecognizedPoint {
	rv := objc.Send[RecognizedPoint](h_.ID, objc.Sel("recognizedPointForJointName:error:"), jointName, error_)
	return rv
}/* debug [instance_methods/method]: RecognizedPointForJointNameError */


// Retrieves the recognized points associated with the joint group name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation/recognizedPoints(_:)
func (h_ HumanBodyPoseObservation) RecognizedPointsForJointsGroupNameError(jointsGroupName HumanBodyPoseObservationJointsGroupName /* typedef */, error_ objectivec.IObject) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](h_.ID, objc.Sel("recognizedPointsForJointsGroupName:error:"), jointsGroupName, error_)
	return rv
}/* debug [instance_methods/method]: RecognizedPointsForJointsGroupNameError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HumanBodyPoseObservation */

// The names of the available joints in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation/availableJointNames
func (h_ HumanBodyPoseObservation) AvailableJointNames() []string {
	rv := objc.Send[[]string](h_.ID, objc.Sel("availableJointNames"))
	return rv
}/* debug [instance_properties/getter]: availableJointNames */


// The available joint group names in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPoseObservation/availableJointsGroupNames
func (h_ HumanBodyPoseObservation) AvailableJointsGroupNames() []string {
	rv := objc.Send[[]string](h_.ID, objc.Sel("availableJointsGroupNames"))
	return rv
}/* debug [instance_properties/getter]: availableJointsGroupNames */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNHumanBodyPoseObservation */



