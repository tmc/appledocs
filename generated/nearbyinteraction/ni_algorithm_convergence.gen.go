// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NIAlgorithmConvergence */


/* debug [class_header]: Header for NIAlgorithmConvergence */
// The class instance for the [NIAlgorithmConvergence] class.
var (
	NIAlgorithmConvergenceClass     _NIAlgorithmConvergenceClass
	NIAlgorithmConvergenceClassOnce sync.Once
)

func getNIAlgorithmConvergenceClass() _NIAlgorithmConvergenceClass {
	NIAlgorithmConvergenceClassOnce.Do(func() {
		NIAlgorithmConvergenceClass = _NIAlgorithmConvergenceClass{objc.GetClass("NIAlgorithmConvergence")}
	})
	return NIAlgorithmConvergenceClass
}

type _NIAlgorithmConvergenceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NIAlgorithmConvergence */
// An interface definition for the [NIAlgorithmConvergence] class.
type INIAlgorithmConvergence interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NIAlgorithmConvergence */
	// properties:
	IsCameraAssistanceEnabled() bool
	SetIsCameraAssistanceEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NIAlgorithmConvergence */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NIAlgorithmConvergence */
// Alloc allocates a new instance without initialization.
func (nc _NIAlgorithmConvergenceClass) Alloc() NIAlgorithmConvergence {
	rv := objc.Send[NIAlgorithmConvergence](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NIAlgorithmConvergenceClass) New() NIAlgorithmConvergence {
	rv := objc.Send[NIAlgorithmConvergence](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NIAlgorithmConvergence) Init() NIAlgorithmConvergence {
	rv := objc.Send[NIAlgorithmConvergence](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NIAlgorithmConvergence) Autorelease() NIAlgorithmConvergence {
	rv := objc.Send[NIAlgorithmConvergence](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNIAlgorithmConvergence creates a new NIAlgorithmConvergence instance.
func NewNIAlgorithmConvergence() NIAlgorithmConvergence {
	return getNIAlgorithmConvergenceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NIAlgorithmConvergence */
// An object that provides the state and reason for user coaching recommendations.
//
// This class conveys the current state of the framework’s Camera Assistance feature when you turn on . When the status indicates that user action is required to achieve the highest-quality results, instances of this class identify specific actions the user can do to help. To improve the status, the app needs to coach the user such as by presenting instructional text. The information you provide tells the user, for example, where and at what speed to pan the device around the environment. To listen for the convergence status, implement .


// An object that provides the state and reason for user coaching recommendations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIAlgorithmConvergence
type NIAlgorithmConvergence struct {
	objectivec.Object
}

// NIAlgorithmConvergenceFrom constructs a [NIAlgorithmConvergence] from an unsafe.Pointer.
//
// An object that provides the state and reason for user coaching recommendations.
func NIAlgorithmConvergenceFrom(ptr unsafe.Pointer) NIAlgorithmConvergence {
	return NIAlgorithmConvergence{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NIAlgorithmConvergence *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NIAlgorithmConvergence */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NIAlgorithmConvergence */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NIAlgorithmConvergence */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NIAlgorithmConvergence */

// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbypeerconfiguration/iscameraassistanceenabled
func (n_ NIAlgorithmConvergence) IsCameraAssistanceEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isCameraAssistanceEnabled"))
	return rv
}/* debug [instance_properties/getter]: isCameraAssistanceEnabled */


// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbypeerconfiguration/iscameraassistanceenabled
func (n_ NIAlgorithmConvergence) SetIsCameraAssistanceEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsCameraAssistanceEnabled:"), value)
}/* debug [instance_properties/setter]: isCameraAssistanceEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NIAlgorithmConvergence */


