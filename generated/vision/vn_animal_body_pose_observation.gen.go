// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNAnimalBodyPoseObservation */


/* debug [class_header]: Header for VNAnimalBodyPoseObservation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AnimalBodyPoseObservation */
// An interface definition for the [AnimalBodyPoseObservation] class.
type IAnimalBodyPoseObservation interface {
	IRecognizedPointsObservation
	
/* debug [class_interface_properties]: Properties for AnimalBodyPoseObservation */
	// properties:
	AvailableJointGroupNames() []string
	AvailableJointNames() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AnimalBodyPoseObservation */
	// methods:
	RecognizedPointForJointNameError(jointName AnimalBodyPoseObservationJointName /* typedef */, error_ objectivec.IObject) IRecognizedPoint
	RecognizedPointsForJointsGroupNameError(jointsGroupName AnimalBodyPoseObservationJointsGroupName /* typedef */, error_ objectivec.IObject) foundation.IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AnimalBodyPoseObservation */
// Alloc allocates a new instance without initialization.
func (ac _AnimalBodyPoseObservationClass) Alloc() AnimalBodyPoseObservation {
	rv := objc.Send[AnimalBodyPoseObservation](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AnimalBodyPoseObservation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AnimalBodyPoseObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AnimalBodyPoseObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AnimalBodyPoseObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AnimalBodyPoseObservation */

// Returns the point for a joint name the observation recognizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation/recognizedPoint(_:)
func (a_ AnimalBodyPoseObservation) RecognizedPointForJointNameError(jointName AnimalBodyPoseObservationJointName /* typedef */, error_ objectivec.IObject) IRecognizedPoint {
	rv := objc.Send[RecognizedPoint](a_.ID, objc.Sel("recognizedPointForJointName:error:"), jointName, error_)
	return rv
}/* debug [instance_methods/method]: RecognizedPointForJointNameError */


// Returns the points for a joint group name the observation recognizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation/recognizedPoints(_:)
func (a_ AnimalBodyPoseObservation) RecognizedPointsForJointsGroupNameError(jointsGroupName AnimalBodyPoseObservationJointsGroupName /* typedef */, error_ objectivec.IObject) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("recognizedPointsForJointsGroupName:error:"), jointsGroupName, error_)
	return rv
}/* debug [instance_methods/method]: RecognizedPointsForJointsGroupNameError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AnimalBodyPoseObservation */

// The available joint group names in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation/availableJointGroupNames
func (a_ AnimalBodyPoseObservation) AvailableJointGroupNames() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableJointGroupNames"))
	return rv
}/* debug [instance_properties/getter]: availableJointGroupNames */


// The names of the available joints in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNAnimalBodyPoseObservation/availableJointNames
func (a_ AnimalBodyPoseObservation) AvailableJointNames() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableJointNames"))
	return rv
}/* debug [instance_properties/getter]: availableJointNames */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNAnimalBodyPoseObservation */



