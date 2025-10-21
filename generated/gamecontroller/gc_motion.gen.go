// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	SetAttitude(attitude unsafe.Pointer)
	SetUserAcceleration(userAcceleration unsafe.Pointer)
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


// Sets the controller’s attitude.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/setAttitude(_:)
func (g_ GCMotion) SetAttitude(attitude unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAttitude:"), attitude)
}

// Sets the acceleration the user applies to the controller.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMotion/setUserAcceleration(_:)
func (g_ GCMotion) SetUserAcceleration(userAcceleration unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setUserAcceleration:"), userAcceleration)
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



