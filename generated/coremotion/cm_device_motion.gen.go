// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [DeviceMotion] class.
type IDeviceMotion interface {
	ILogItem
	// properties:
	Attitude() ICMAttitude
	Gravity() Acceleration /* not a class type */
	Heading() float64 /* primitive/slice/pointer. */
	MagneticField() CalibratedMagneticField /* not a class type */
	RotationRate() RotationRate /* not a class type */
	SensorLocation() DeviceMotionSensorLocation
	UserAcceleration() Acceleration /* not a class type */
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (dc _DeviceMotionClass) Alloc() DeviceMotion {
	rv := objc.Send[DeviceMotion](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The attitude of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/attitude
func (d_ DeviceMotion) Attitude() ICMAttitude {
	rv := objc.Send[Attitude](d_.ID, objc.Sel("attitude"))
	return rv
}


// The gravity acceleration vector expressed in the device’s reference frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/gravity
func (d_ DeviceMotion) Gravity() Acceleration /* not a class type */ {
	rv := objc.Send[Acceleration](d_.ID, objc.Sel("gravity"))
	return rv
}


// The heading angle (measured in degrees) relative to the current reference frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/heading
func (d_ DeviceMotion) Heading() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](d_.ID, objc.Sel("heading"))
	return rv
}


// Returns the magnetic field vector with respect to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/magneticField
func (d_ DeviceMotion) MagneticField() CalibratedMagneticField /* not a class type */ {
	rv := objc.Send[CalibratedMagneticField](d_.ID, objc.Sel("magneticField"))
	return rv
}


// The rotation rate of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/rotationRate
func (d_ DeviceMotion) RotationRate() RotationRate /* not a class type */ {
	rv := objc.Send[RotationRate](d_.ID, objc.Sel("rotationRate"))
	return rv
}


// The location of the sensors that compute the device-motion data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/sensorLocation-swift.property
func (d_ DeviceMotion) SensorLocation() DeviceMotionSensorLocation {
	rv := objc.Send[DeviceMotionSensorLocation](d_.ID, objc.Sel("sensorLocation"))
	return rv
}


// The acceleration that the user is giving to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/userAcceleration
func (d_ DeviceMotion) UserAcceleration() Acceleration /* not a class type */ {
	rv := objc.Send[Acceleration](d_.ID, objc.Sel("userAcceleration"))
	return rv
}



