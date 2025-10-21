// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [GCMotion] class.
type IGCMotion interface {
	objectivec.IObject
}

// A controller profile that supports orientation and motion.
//
// The motion controller profile provides attitude and rotation data, as well as acceleration and sensor information. Use this profile to get motion input from a controller that measures acceleration and rotation rate. If the controller’s property is a object, the controller supports motion. This illustration shows the direction of the x, y, and z axes of an iPhone when held upright.
//
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

// Alloc allocates a new instance without initialization.
func (gc _GCMotionClass) Alloc() GCMotion {
	rv := objc.Send[GCMotion](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A Boolean value that indicates whether the controller provides attitude and rotation data.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/hasAttitudeAndRotationRate
func (g_ GCMotion) HasAttitudeAndRotationRate() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hasAttitudeAndRotationRate"))
	return rv
}

// The rotation rate of the controller.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/rotationRate
func (g_ GCMotion) RotationRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("rotationRate"))
	return rv
}

// A Boolean value that indicates whether the sensors that compute the motion data require manual activation.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/sensorsRequireManualActivation
func (g_ GCMotion) SensorsRequireManualActivation() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("sensorsRequireManualActivation"))
	return rv
}

// The motion input profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCMotion) Motion() GCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}


// SetMotion sets the value of the motion property.
// The motion input profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCMotion) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}

// The total acceleration of the controller that includes gravity and the acceleration the user applies to the controller.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/acceleration
func (g_ GCMotion) Acceleration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("acceleration"))
	return rv
}


// SetAcceleration sets the value of the acceleration property.
// The total acceleration of the controller that includes gravity and the acceleration the user applies to the controller.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/acceleration
func (g_ GCMotion) SetAcceleration(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAcceleration:"), value)
}

// The attitude of the controller.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/attitude
func (g_ GCMotion) Attitude() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("attitude"))
	return rv
}


// SetAttitude sets the value of the attitude property.
// The attitude of the controller.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/attitude
func (g_ GCMotion) SetAttitude(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAttitude:"), value)
}

// The controller for the profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/controller
func (g_ GCMotion) Controller() GCController {
	rv := objc.Send[GCController](g_.ID, objc.Sel("controller"))
	return rv
}


// SetController sets the value of the controller property.
// The controller for the profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/controller
func (g_ GCMotion) SetController(value IGCController) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setController:"), value)
}

// The gravity acceleration vector from the controller’s reference frame.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/gravity
func (g_ GCMotion) Gravity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("gravity"))
	return rv
}


// SetGravity sets the value of the gravity property.
// The gravity acceleration vector from the controller’s reference frame.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/gravity
func (g_ GCMotion) SetGravity(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGravity:"), value)
}

// A Boolean value that indicates whether the controller provides attitude data.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/hasattitude
func (g_ GCMotion) HasAttitude() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hasAttitude"))
	return rv
}


// SetHasAttitude sets the value of the hasAttitude property.
// A Boolean value that indicates whether the controller provides attitude data.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/hasattitude
func (g_ GCMotion) SetHasAttitude(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setHasAttitude:"), value)
}

// A Boolean value that indicates whether the controller provides gravity and user acceleration data.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/hasgravityanduseracceleration
func (g_ GCMotion) HasGravityAndUserAcceleration() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hasGravityAndUserAcceleration"))
	return rv
}


// SetHasGravityAndUserAcceleration sets the value of the hasGravityAndUserAcceleration property.
// A Boolean value that indicates whether the controller provides gravity and user acceleration data.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/hasgravityanduseracceleration
func (g_ GCMotion) SetHasGravityAndUserAcceleration(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setHasGravityAndUserAcceleration:"), value)
}

// A Boolean value that indicates whether the controller provides rotation data.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/hasrotationrate
func (g_ GCMotion) HasRotationRate() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hasRotationRate"))
	return rv
}


// SetHasRotationRate sets the value of the hasRotationRate property.
// A Boolean value that indicates whether the controller provides rotation data.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/hasrotationrate
func (g_ GCMotion) SetHasRotationRate(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setHasRotationRate:"), value)
}

// A Boolean value that indicates whether the sensors that compute the motion data are active.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/sensorsactive
func (g_ GCMotion) SensorsActive() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("sensorsActive"))
	return rv
}


// SetSensorsActive sets the value of the sensorsActive property.
// A Boolean value that indicates whether the sensors that compute the motion data are active.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/sensorsactive
func (g_ GCMotion) SetSensorsActive(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSensorsActive:"), value)
}

// The acceleration that the user applies to the controller.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/useracceleration
func (g_ GCMotion) UserAcceleration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("userAcceleration"))
	return rv
}


// SetUserAcceleration sets the value of the userAcceleration property.
// The acceleration that the user applies to the controller.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/useracceleration
func (g_ GCMotion) SetUserAcceleration(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setUserAcceleration:"), value)
}

// The block that the profile calls when an element’s value changes.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/valuechangedhandler
func (g_ GCMotion) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}


// SetValueChangedHandler sets the value of the valueChangedHandler property.
// The block that the profile calls when an element’s value changes.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmotion/valuechangedhandler
func (g_ GCMotion) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}



