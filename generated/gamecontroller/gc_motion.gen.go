// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCMotion */


/* debug [class_header]: Header for GCMotion */
// The class instance for the [GCMotion] class.
var (
	GCMotionClass     _GCMotionClass
	GCMotionClassOnce sync.Once
)

func getGCMotionClass() _GCMotionClass {
	GCMotionClassOnce.Do(func() {
		GCMotionClass = _GCMotionClass{objc.GetClass("GCMotion")}
	})
	return GCMotionClass
}

type _GCMotionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCMotion */
// An interface definition for the [GCMotion] class.
type IGCMotion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCMotion */
	// properties:
	Acceleration() objc.IObject /* cross-framework: GCAcceleration */
	Attitude() objc.IObject /* cross-framework: GCQuaternion */
	Controller() IGCController
	Gravity() objc.IObject /* cross-framework: GCAcceleration */
	HasAttitude() bool
	HasAttitudeAndRotationRate() bool
	HasGravityAndUserAcceleration() bool
	HasRotationRate() bool
	RotationRate() objc.IObject /* cross-framework: GCRotationRate */
	SensorsActive() bool
	SetSensorsActive(value bool)
	SensorsRequireManualActivation() bool
	UserAcceleration() objc.IObject /* cross-framework: GCAcceleration */
	ValueChangedHandler() unsafe.Pointer
	SetValueChangedHandler(value unsafe.Pointer)
	Motion() IGCMotion
	SetMotion(value IGCMotion)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCMotion */
	// methods:
	SetStateFromMotion(motion IGCMotion)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCMotion */
// Alloc allocates a new instance without initialization.
func (gc _GCMotionClass) Alloc() GCMotion {
	rv := objc.Send[GCMotion](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCMotionClass) New() GCMotion {
	rv := objc.Send[GCMotion](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCMotion) Init() GCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCMotion) Autorelease() GCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCMotion creates a new GCMotion instance.
func NewGCMotion() GCMotion {
	return getGCMotionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCMotion */
// A controller profile that supports orientation and motion.
//
// The motion controller profile provides attitude and rotation data, as well as acceleration and sensor information. Use this profile to get motion input from a controller that measures acceleration and rotation rate. If the controller’s property is a object, the controller supports motion. This illustration shows the direction of the x, y, and z axes of an iPhone when held upright.


// A controller profile that supports orientation and motion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion
type GCMotion struct {
	objectivec.Object
}

// GCMotionFrom constructs a [GCMotion] from an unsafe.Pointer.
//
// A controller profile that supports orientation and motion.
func GCMotionFrom(ptr unsafe.Pointer) GCMotion {
	return GCMotion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCMotion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCMotion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCMotion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCMotion */

// Copies the input values from a specified motion profile to a snapshot of a motion profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/setStateFrom(_:)
func (g_ GCMotion) SetStateFromMotion(motion IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStateFromMotion:"), motion)
}/* debug [instance_methods/method]: SetStateFromMotion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCMotion */

// The total acceleration of the controller that includes gravity and the acceleration the user applies to the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/acceleration
func (g_ GCMotion) Acceleration() objc.IObject /* cross-framework: GCAcceleration */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("acceleration"))
	return rv
}/* debug [instance_properties/getter]: acceleration */


// The attitude of the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/attitude
func (g_ GCMotion) Attitude() objc.IObject /* cross-framework: GCQuaternion */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("attitude"))
	return rv
}/* debug [instance_properties/getter]: attitude */


// The controller for the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/controller
func (g_ GCMotion) Controller() IGCController {
	rv := objc.Send[GCController](g_.ID, objc.Sel("controller"))
	return rv
}/* debug [instance_properties/getter]: controller */


// The gravity acceleration vector from the controller’s reference frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/gravity
func (g_ GCMotion) Gravity() objc.IObject /* cross-framework: GCAcceleration */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("gravity"))
	return rv
}/* debug [instance_properties/getter]: gravity */


// A Boolean value that indicates whether the controller provides attitude data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/hasAttitude
func (g_ GCMotion) HasAttitude() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hasAttitude"))
	return rv
}/* debug [instance_properties/getter]: hasAttitude */


// A Boolean value that indicates whether the controller provides attitude and rotation data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/hasAttitudeAndRotationRate
func (g_ GCMotion) HasAttitudeAndRotationRate() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hasAttitudeAndRotationRate"))
	return rv
}/* debug [instance_properties/getter]: hasAttitudeAndRotationRate */


// A Boolean value that indicates whether the controller provides gravity and user acceleration data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/hasGravityAndUserAcceleration
func (g_ GCMotion) HasGravityAndUserAcceleration() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hasGravityAndUserAcceleration"))
	return rv
}/* debug [instance_properties/getter]: hasGravityAndUserAcceleration */


// A Boolean value that indicates whether the controller provides rotation data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/hasRotationRate
func (g_ GCMotion) HasRotationRate() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hasRotationRate"))
	return rv
}/* debug [instance_properties/getter]: hasRotationRate */


// The rotation rate of the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/rotationRate
func (g_ GCMotion) RotationRate() objc.IObject /* cross-framework: GCRotationRate */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("rotationRate"))
	return rv
}/* debug [instance_properties/getter]: rotationRate */


// A Boolean value that indicates whether the sensors that compute the motion data are active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/sensorsActive
func (g_ GCMotion) SensorsActive() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("sensorsActive"))
	return rv
}/* debug [instance_properties/getter]: sensorsActive */


// A Boolean value that indicates whether the sensors that compute the motion data are active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/sensorsActive
func (g_ GCMotion) SetSensorsActive(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSensorsActive:"), value)
}/* debug [instance_properties/setter]: sensorsActive */


// A Boolean value that indicates whether the sensors that compute the motion data require manual activation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/sensorsRequireManualActivation
func (g_ GCMotion) SensorsRequireManualActivation() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("sensorsRequireManualActivation"))
	return rv
}/* debug [instance_properties/getter]: sensorsRequireManualActivation */


// The acceleration that the user applies to the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/userAcceleration
func (g_ GCMotion) UserAcceleration() objc.IObject /* cross-framework: GCAcceleration */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("userAcceleration"))
	return rv
}/* debug [instance_properties/getter]: userAcceleration */


// The block that the profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/valueChangedHandler
func (g_ GCMotion) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}/* debug [instance_properties/getter]: valueChangedHandler */


// The block that the profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/valueChangedHandler
func (g_ GCMotion) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}/* debug [instance_properties/setter]: valueChangedHandler */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCMotion) Motion() IGCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}/* debug [instance_properties/getter]: motion */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCMotion) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}/* debug [instance_properties/setter]: motion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCMotion */



