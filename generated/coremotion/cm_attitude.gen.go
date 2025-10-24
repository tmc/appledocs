// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMAttitude */


/* debug [class_header]: Header for CMAttitude */
// The class instance for the [Attitude] class.
var (
	AttitudeClass     _AttitudeClass
	AttitudeClassOnce sync.Once
)

func getAttitudeClass() _AttitudeClass {
	AttitudeClassOnce.Do(func() {
		AttitudeClass = _AttitudeClass{objc.GetClass("CMAttitude")}
	})
	return AttitudeClass
}

type _AttitudeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Attitude */
// An interface definition for the [Attitude] class.
type IAttitude interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Attitude */
	// properties:
	Pitch() float64
	Quaternion() objc.IObject /* cross-framework: CMQuaternion */
	Roll() float64
	RotationMatrix() objc.IObject /* cross-framework: CMRotationMatrix */
	Yaw() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Attitude */
	// methods:
	MultiplyByInverseOfAttitude(attitude ICMAttitude)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Attitude */
// Alloc allocates a new instance without initialization.
func (ac _AttitudeClass) Alloc() Attitude {
	rv := objc.Send[Attitude](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AttitudeClass) New() Attitude {
	rv := objc.Send[Attitude](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Attitude) Init() Attitude {
	rv := objc.Send[Attitude](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Attitude) Autorelease() Attitude {
	rv := objc.Send[Attitude](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAttitude creates a new Attitude instance.
func NewAttitude() Attitude {
	return getAttitudeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Attitude */
// The device’s orientation relative to a known frame of reference at a point in time.
//
// The class offers three different mathematical representations of attitude: a rotation matrix, a quaternion, and Euler angles (roll, pitch, and yaw values). You access objects through the attitude property of each objects passed to an application. An application starts receiving these device-motion objects as a result of calling the method, the method, the method, or the method of the class.


// The device’s orientation relative to a known frame of reference at a point in time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitude
type Attitude struct {
	objectivec.Object
}

// AttitudeFrom constructs a [Attitude] from an unsafe.Pointer.
//
// The device’s orientation relative to a known frame of reference at a point in time.
func AttitudeFrom(ptr unsafe.Pointer) Attitude {
	return Attitude{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Attitude *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Attitude */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Attitude */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Attitude */

// Yields the change in attitude given a specific attitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitude/multiply(byInverseOf:)
func (a_ Attitude) MultiplyByInverseOfAttitude(attitude ICMAttitude) {
	objc.Send[objc.ID](a_.ID, objc.Sel("multiplyByInverseOfAttitude:"), attitude)
}/* debug [instance_methods/method]: MultiplyByInverseOfAttitude */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Attitude */

// The pitch of the device, in radians.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitude/pitch
func (a_ Attitude) Pitch() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("pitch"))
	return rv
}/* debug [instance_properties/getter]: pitch */


// Returns a quaternion representing the device’s attitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitude/quaternion
func (a_ Attitude) Quaternion() objc.IObject /* cross-framework: CMQuaternion */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("quaternion"))
	return rv
}/* debug [instance_properties/getter]: quaternion */


// The roll of the device, in radians.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitude/roll
func (a_ Attitude) Roll() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("roll"))
	return rv
}/* debug [instance_properties/getter]: roll */


// Returns a rotation matrix representing the device’s attitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitude/rotationMatrix
func (a_ Attitude) RotationMatrix() objc.IObject /* cross-framework: CMRotationMatrix */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("rotationMatrix"))
	return rv
}/* debug [instance_properties/getter]: rotationMatrix */


// The yaw of the device, in radians.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitude/yaw
func (a_ Attitude) Yaw() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("yaw"))
	return rv
}/* debug [instance_properties/getter]: yaw */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMAttitude */



