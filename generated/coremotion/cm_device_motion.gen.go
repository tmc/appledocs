// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CMDeviceMotion */


/* debug [class_header]: Header for CMDeviceMotion */
// The class instance for the [DeviceMotion] class.
var (
	DeviceMotionClass     _DeviceMotionClass
	DeviceMotionClassOnce sync.Once
)

func getDeviceMotionClass() _DeviceMotionClass {
	DeviceMotionClassOnce.Do(func() {
		DeviceMotionClass = _DeviceMotionClass{objc.GetClass("CMDeviceMotion")}
	})
	return DeviceMotionClass
}

type _DeviceMotionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DeviceMotion */
// An interface definition for the [DeviceMotion] class.
type IDeviceMotion interface {
	ILogItem
	
/* debug [class_interface_properties]: Properties for DeviceMotion */
	// properties:
	Attitude() ICMAttitude
	Gravity() objc.IObject /* cross-framework: CMAcceleration */
	Heading() float64
	MagneticField() objc.IObject /* cross-framework: CMCalibratedMagneticField */
	RotationRate() objc.IObject /* cross-framework: CMRotationRate */
	SensorLocation() DeviceMotionSensorLocation
	UserAcceleration() objc.IObject /* cross-framework: CMAcceleration */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DeviceMotion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DeviceMotion */
// Alloc allocates a new instance without initialization.
func (dc _DeviceMotionClass) Alloc() DeviceMotion {
	rv := objc.Send[DeviceMotion](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DeviceMotionClass) New() DeviceMotion {
	rv := objc.Send[DeviceMotion](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DeviceMotion) Init() DeviceMotion {
	rv := objc.Send[DeviceMotion](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DeviceMotion) Autorelease() DeviceMotion {
	rv := objc.Send[DeviceMotion](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDeviceMotion creates a new DeviceMotion instance.
func NewDeviceMotion() DeviceMotion {
	return getDeviceMotionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DeviceMotion */
// Encapsulated measurements of the attitude, rotation rate, and acceleration of a device.
//
// An application receives or samples objects at regular intervals after calling the method, the method, the method, or the method of the class. The accelerometer measures the sum of two acceleration vectors: gravity and user acceleration. User acceleration is the acceleration that the user imparts to the device. Because Core Motion is able to track a device’s attitude using both the gyroscope and the accelerometer, it can differentiate between gravity and user acceleration. A object provides both measurements in the and properties.


// Encapsulated measurements of the attitude, rotation rate, and acceleration of a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion
type DeviceMotion struct {
	LogItem
}

// DeviceMotionFrom constructs a [DeviceMotion] from an unsafe.Pointer.
//
// Encapsulated measurements of the attitude, rotation rate, and acceleration of a device.
func DeviceMotionFrom(ptr unsafe.Pointer) DeviceMotion {
	return DeviceMotion{
		LogItem: LogItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DeviceMotion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DeviceMotion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DeviceMotion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DeviceMotion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DeviceMotion */

// The attitude of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/attitude
func (d_ DeviceMotion) Attitude() ICMAttitude {
	rv := objc.Send[Attitude](d_.ID, objc.Sel("attitude"))
	return rv
}/* debug [instance_properties/getter]: attitude */


// The gravity acceleration vector expressed in the device’s reference frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/gravity
func (d_ DeviceMotion) Gravity() objc.IObject /* cross-framework: CMAcceleration */ {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("gravity"))
	return rv
}/* debug [instance_properties/getter]: gravity */


// The heading angle (measured in degrees) relative to the current reference frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/heading
func (d_ DeviceMotion) Heading() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("heading"))
	return rv
}/* debug [instance_properties/getter]: heading */


// Returns the magnetic field vector with respect to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/magneticField
func (d_ DeviceMotion) MagneticField() objc.IObject /* cross-framework: CMCalibratedMagneticField */ {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("magneticField"))
	return rv
}/* debug [instance_properties/getter]: magneticField */


// The rotation rate of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/rotationRate
func (d_ DeviceMotion) RotationRate() objc.IObject /* cross-framework: CMRotationRate */ {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("rotationRate"))
	return rv
}/* debug [instance_properties/getter]: rotationRate */


// The location of the sensors that compute the device-motion data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/sensorLocation-swift.property
func (d_ DeviceMotion) SensorLocation() DeviceMotionSensorLocation {
	rv := objc.Send[DeviceMotionSensorLocation](d_.ID, objc.Sel("sensorLocation"))
	return rv
}/* debug [instance_properties/getter]: sensorLocation */


// The acceleration that the user is giving to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/userAcceleration
func (d_ DeviceMotion) UserAcceleration() objc.IObject /* cross-framework: CMAcceleration */ {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("userAcceleration"))
	return rv
}/* debug [instance_properties/getter]: userAcceleration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMDeviceMotion */



