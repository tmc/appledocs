// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MotionManager] class.
var (
	MotionManagerClass     _MotionManagerClass
	MotionManagerClassOnce sync.Once
)

func getMotionManagerClass() _MotionManagerClass {
	MotionManagerClassOnce.Do(func() {
		MotionManagerClass = _MotionManagerClass{objc.GetClass("CMMotionManager")}
	})
	return MotionManagerClass
}

type _MotionManagerClass struct {
	class objc.Class
}

// An interface definition for the [MotionManager] class.
type IMotionManager interface {
	objectivec.IObject
	// properties:
	CMErrorDomain() objc.IObject /* cross-framework: NSString */
	IsAccelerometerActive() bool
	SetIsAccelerometerActive(value bool)
	IsAccelerometerAvailable() bool
	SetIsAccelerometerAvailable(value bool)
	IsDeviceMotionActive() bool
	SetIsDeviceMotionActive(value bool)
	IsDeviceMotionAvailable() bool
	SetIsDeviceMotionAvailable(value bool)
	IsGyroActive() bool
	SetIsGyroActive(value bool)
	IsGyroAvailable() bool
	SetIsGyroAvailable(value bool)
	IsMagnetometerActive() bool
	SetIsMagnetometerActive(value bool)
	IsMagnetometerAvailable() bool
	SetIsMagnetometerAvailable(value bool)
	// methods:
}

// The object for starting and managing motion services.
//
// Use a object to start the services that report movement detected by the device’s onboard sensors. Use this object to receive four types of motion data: , indicating the instantaneous acceleration of the device in three dimensional space. , indicating the instantaneous rotation around the device’s three primary axes. , indicating the device’s orientation relative to Earth’s magnetic field. , indicating key motion-related attributes such as the device’s user-initiated acceleration, its attitude, rotation rates, orientation relative to calibrated magnetic fields, and orientation relative to gravity. Core Motion’s sensor fusion algorithms provide this data. The processed device-motion data gives the device’s attitude, rotation rate, calibrated magnetic fields, the direction of gravity, and the amount of acceleration the user contributes to the device. You can receive live sensor data at a specified update interval, or you can let the sensors collect data and store it for retrieval later. With both of these approaches, call the appropriate stop method ( , , , and ) when you no longer need the data.


// The object for starting and managing motion services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager
type MotionManager struct {
	objectivec.Object
}

// MotionManagerFrom constructs a [MotionManager] from an unsafe.Pointer.
//
// The object for starting and managing motion services.
func MotionManagerFrom(ptr unsafe.Pointer) MotionManager {
	return MotionManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MotionManagerClass) Alloc() MotionManager {
	rv := objc.Send[MotionManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MotionManagerClass) New() MotionManager {
	rv := objc.Send[MotionManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MotionManager) Init() MotionManager {
	rv := objc.Send[MotionManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MotionManager) Autorelease() MotionManager {
	rv := objc.Send[MotionManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMotionManager creates a new MotionManager instance.
func NewMotionManager() MotionManager {
	return getMotionManagerClass().New()
}



// Returns a bitmask of the available reference frames for reporting the attitude of the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/availableAttitudeReferenceFrames()
func (mc _MotionManagerClass) AvailableAttitudeReferenceFrames() AttitudeReferenceFrame {
	rv := objc.Send[AttitudeReferenceFrame](objc.ID(mc.class), objc.Sel("availableAttitudeReferenceFrames"))
	return rv
}


// The error domain for Core Motion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmerrordomain
func (m_ MotionManager) CMErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("CMErrorDomain"))
	return rv
}


// A Boolean value that indicates whether accelerometer updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isaccelerometeractive
func (m_ MotionManager) IsAccelerometerActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isAccelerometerActive"))
	return rv
}


// A Boolean value that indicates whether accelerometer updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isaccelerometeractive
func (m_ MotionManager) SetIsAccelerometerActive(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsAccelerometerActive:"), value)
}


// A Boolean value that indicates whether an accelerometer is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isaccelerometeravailable
func (m_ MotionManager) IsAccelerometerAvailable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isAccelerometerAvailable"))
	return rv
}


// A Boolean value that indicates whether an accelerometer is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isaccelerometeravailable
func (m_ MotionManager) SetIsAccelerometerAvailable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsAccelerometerAvailable:"), value)
}


// A Boolean value that determines whether the app is receiving updates from the device-motion service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isdevicemotionactive
func (m_ MotionManager) IsDeviceMotionActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isDeviceMotionActive"))
	return rv
}


// A Boolean value that determines whether the app is receiving updates from the device-motion service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isdevicemotionactive
func (m_ MotionManager) SetIsDeviceMotionActive(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsDeviceMotionActive:"), value)
}


// A Boolean value that indicates whether the device-motion service is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isdevicemotionavailable
func (m_ MotionManager) IsDeviceMotionAvailable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isDeviceMotionAvailable"))
	return rv
}


// A Boolean value that indicates whether the device-motion service is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isdevicemotionavailable
func (m_ MotionManager) SetIsDeviceMotionAvailable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsDeviceMotionAvailable:"), value)
}


// A Boolean value that determines whether gyroscope updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isgyroactive
func (m_ MotionManager) IsGyroActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isGyroActive"))
	return rv
}


// A Boolean value that determines whether gyroscope updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isgyroactive
func (m_ MotionManager) SetIsGyroActive(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsGyroActive:"), value)
}


// A Boolean value that indicates whether a gyroscope is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isgyroavailable
func (m_ MotionManager) IsGyroAvailable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isGyroAvailable"))
	return rv
}


// A Boolean value that indicates whether a gyroscope is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isgyroavailable
func (m_ MotionManager) SetIsGyroAvailable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsGyroAvailable:"), value)
}


// A Boolean value that determines whether magnetometer updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/ismagnetometeractive
func (m_ MotionManager) IsMagnetometerActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isMagnetometerActive"))
	return rv
}


// A Boolean value that determines whether magnetometer updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/ismagnetometeractive
func (m_ MotionManager) SetIsMagnetometerActive(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsMagnetometerActive:"), value)
}


// A Boolean value that indicates whether a magnetometer is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/ismagnetometeravailable
func (m_ MotionManager) IsMagnetometerAvailable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isMagnetometerAvailable"))
	return rv
}


// A Boolean value that indicates whether a magnetometer is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/ismagnetometeravailable
func (m_ MotionManager) SetIsMagnetometerAvailable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsMagnetometerAvailable:"), value)
}


