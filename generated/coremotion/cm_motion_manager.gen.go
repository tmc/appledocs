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
	AccelerometerData() ICMAccelerometerData
	AccelerometerUpdateInterval() foundation.TimeInterval /* not a class type */
	SetAccelerometerUpdateInterval(value foundation.TimeInterval /* not a class type */)
	AttitudeReferenceFrame() AttitudeReferenceFrame
	DeviceMotion() ICMDeviceMotion
	DeviceMotionUpdateInterval() foundation.TimeInterval /* not a class type */
	SetDeviceMotionUpdateInterval(value foundation.TimeInterval /* not a class type */)
	GyroData() ICMGyroData
	GyroUpdateInterval() foundation.TimeInterval /* not a class type */
	SetGyroUpdateInterval(value foundation.TimeInterval /* not a class type */)
	AccelerometerActive() bool /* primitive/slice/pointer. */
	AccelerometerAvailable() bool /* primitive/slice/pointer. */
	DeviceMotionActive() bool /* primitive/slice/pointer. */
	DeviceMotionAvailable() bool /* primitive/slice/pointer. */
	GyroActive() bool /* primitive/slice/pointer. */
	GyroAvailable() bool /* primitive/slice/pointer. */
	MagnetometerActive() bool /* primitive/slice/pointer. */
	MagnetometerAvailable() bool /* primitive/slice/pointer. */
	MagnetometerData() ICMMagnetometerData
	MagnetometerUpdateInterval() foundation.TimeInterval /* not a class type */
	SetMagnetometerUpdateInterval(value foundation.TimeInterval /* not a class type */)
	ShowsDeviceMovementDisplay() bool /* primitive/slice/pointer. */
	SetShowsDeviceMovementDisplay(value bool /* primitive/slice/pointer. */)
	CMErrorDomain() string /* primitive/slice/pointer. */
	IsAccelerometerActive() bool /* primitive/slice/pointer. */
	SetIsAccelerometerActive(value bool /* primitive/slice/pointer. */)
	IsAccelerometerAvailable() bool /* primitive/slice/pointer. */
	SetIsAccelerometerAvailable(value bool /* primitive/slice/pointer. */)
	IsDeviceMotionActive() bool /* primitive/slice/pointer. */
	SetIsDeviceMotionActive(value bool /* primitive/slice/pointer. */)
	IsDeviceMotionAvailable() bool /* primitive/slice/pointer. */
	SetIsDeviceMotionAvailable(value bool /* primitive/slice/pointer. */)
	IsGyroActive() bool /* primitive/slice/pointer. */
	SetIsGyroActive(value bool /* primitive/slice/pointer. */)
	IsGyroAvailable() bool /* primitive/slice/pointer. */
	SetIsGyroAvailable(value bool /* primitive/slice/pointer. */)
	IsMagnetometerActive() bool /* primitive/slice/pointer. */
	SetIsMagnetometerActive(value bool /* primitive/slice/pointer. */)
	IsMagnetometerAvailable() bool /* primitive/slice/pointer. */
	SetIsMagnetometerAvailable(value bool /* primitive/slice/pointer. */)
	// methods:
	StartAccelerometerUpdates()
	StartAccelerometerUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework OperationQueue */, handler AccelerometerHandler /* not a class type */)
	StartDeviceMotionUpdates()
	StartDeviceMotionUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework OperationQueue */, handler DeviceMotionHandler /* not a class type */)
	StartDeviceMotionUpdatesUsingReferenceFrame(referenceFrame AttitudeReferenceFrame)
	StartDeviceMotionUpdatesUsingReferenceFrameToQueueWithHandler(referenceFrame AttitudeReferenceFrame, queue objc.IObject /* cross-framework OperationQueue */, handler DeviceMotionHandler /* not a class type */)
	StartGyroUpdates()
	StartGyroUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework OperationQueue */, handler GyroHandler /* not a class type */)
	StartMagnetometerUpdates()
	StartMagnetometerUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework OperationQueue */, handler MagnetometerHandler /* not a class type */)
	StopAccelerometerUpdates()
	StopDeviceMotionUpdates()
	StopGyroUpdates()
	StopMagnetometerUpdates()
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


// Starts accelerometer updates without a handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/startAccelerometerUpdates()
func (m_ MotionManager) StartAccelerometerUpdates() {
	objc.Send[objc.ID](m_.ID, objc.Sel("startAccelerometerUpdates"))
}


// Starts accelerometer updates on an operation queue and with a specified handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/startAccelerometerUpdates(to:withHandler:)
func (m_ MotionManager) StartAccelerometerUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework OperationQueue */, handler AccelerometerHandler /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startAccelerometerUpdatesToQueue:withHandler:"), queue, handler)
}


// Starts device-motion updates without a block handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/startDeviceMotionUpdates()
func (m_ MotionManager) StartDeviceMotionUpdates() {
	objc.Send[objc.ID](m_.ID, objc.Sel("startDeviceMotionUpdates"))
}


// Starts device-motion updates on an operation queue and using a specified block handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/startDeviceMotionUpdates(to:withHandler:)
func (m_ MotionManager) StartDeviceMotionUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework OperationQueue */, handler DeviceMotionHandler /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startDeviceMotionUpdatesToQueue:withHandler:"), queue, handler)
}


// Starts device-motion updates using a reference frame but without a block handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/startDeviceMotionUpdates(using:)
func (m_ MotionManager) StartDeviceMotionUpdatesUsingReferenceFrame(referenceFrame AttitudeReferenceFrame) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startDeviceMotionUpdatesUsingReferenceFrame:"), referenceFrame)
}


// Starts device-motion updates on an operation queue and using a specified reference frame and block handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/startDeviceMotionUpdates(using:to:withHandler:)
func (m_ MotionManager) StartDeviceMotionUpdatesUsingReferenceFrameToQueueWithHandler(referenceFrame AttitudeReferenceFrame, queue objc.IObject /* cross-framework OperationQueue */, handler DeviceMotionHandler /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startDeviceMotionUpdatesUsingReferenceFrame:toQueue:withHandler:"), referenceFrame, queue, handler)
}


// Starts gyroscope updates without a handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/startGyroUpdates()
func (m_ MotionManager) StartGyroUpdates() {
	objc.Send[objc.ID](m_.ID, objc.Sel("startGyroUpdates"))
}


// Starts gyroscope updates on an operation queue and with a specified handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/startGyroUpdates(to:withHandler:)
func (m_ MotionManager) StartGyroUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework OperationQueue */, handler GyroHandler /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startGyroUpdatesToQueue:withHandler:"), queue, handler)
}


// Starts magnetometer updates without a block handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/startMagnetometerUpdates()
func (m_ MotionManager) StartMagnetometerUpdates() {
	objc.Send[objc.ID](m_.ID, objc.Sel("startMagnetometerUpdates"))
}


// Starts magnetometer updates on an operation queue and with a specified handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/startMagnetometerUpdates(to:withHandler:)
func (m_ MotionManager) StartMagnetometerUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework OperationQueue */, handler MagnetometerHandler /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startMagnetometerUpdatesToQueue:withHandler:"), queue, handler)
}


// Stops accelerometer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/stopAccelerometerUpdates()
func (m_ MotionManager) StopAccelerometerUpdates() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopAccelerometerUpdates"))
}


// Stops device-motion updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/stopDeviceMotionUpdates()
func (m_ MotionManager) StopDeviceMotionUpdates() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopDeviceMotionUpdates"))
}


// Stops gyroscope updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/stopGyroUpdates()
func (m_ MotionManager) StopGyroUpdates() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopGyroUpdates"))
}


// Stops magnetometer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/stopMagnetometerUpdates()
func (m_ MotionManager) StopMagnetometerUpdates() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopMagnetometerUpdates"))
}


// The latest sample of accelerometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/accelerometerData
func (m_ MotionManager) AccelerometerData() ICMAccelerometerData {
	rv := objc.Send[AccelerometerData](m_.ID, objc.Sel("accelerometerData"))
	return rv
}


// The interval, in seconds, for providing accelerometer updates to the block handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/accelerometerUpdateInterval
func (m_ MotionManager) AccelerometerUpdateInterval() foundation.TimeInterval /* not a class type */ {
	rv := objc.Send[foundation.TimeInterval](m_.ID, objc.Sel("accelerometerUpdateInterval"))
	return rv
}


// The interval, in seconds, for providing accelerometer updates to the block handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/accelerometerUpdateInterval
func (m_ MotionManager) SetAccelerometerUpdateInterval(value foundation.TimeInterval /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAccelerometerUpdateInterval:"), value)
}


// Returns either the reference frame currently being used or the default attitude reference frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/attitudeReferenceFrame
func (m_ MotionManager) AttitudeReferenceFrame() AttitudeReferenceFrame {
	rv := objc.Send[AttitudeReferenceFrame](m_.ID, objc.Sel("attitudeReferenceFrame"))
	return rv
}


// The latest sample of device-motion data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/deviceMotion
func (m_ MotionManager) DeviceMotion() ICMDeviceMotion {
	rv := objc.Send[DeviceMotion](m_.ID, objc.Sel("deviceMotion"))
	return rv
}


// The interval, in seconds, for providing device-motion updates to the block handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/deviceMotionUpdateInterval
func (m_ MotionManager) DeviceMotionUpdateInterval() foundation.TimeInterval /* not a class type */ {
	rv := objc.Send[foundation.TimeInterval](m_.ID, objc.Sel("deviceMotionUpdateInterval"))
	return rv
}


// The interval, in seconds, for providing device-motion updates to the block handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/deviceMotionUpdateInterval
func (m_ MotionManager) SetDeviceMotionUpdateInterval(value foundation.TimeInterval /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceMotionUpdateInterval:"), value)
}


// The latest sample of gyroscope data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/gyroData
func (m_ MotionManager) GyroData() ICMGyroData {
	rv := objc.Send[GyroData](m_.ID, objc.Sel("gyroData"))
	return rv
}


// The interval, in seconds, for providing gyroscope updates to the block handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/gyroUpdateInterval
func (m_ MotionManager) GyroUpdateInterval() foundation.TimeInterval /* not a class type */ {
	rv := objc.Send[foundation.TimeInterval](m_.ID, objc.Sel("gyroUpdateInterval"))
	return rv
}


// The interval, in seconds, for providing gyroscope updates to the block handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/gyroUpdateInterval
func (m_ MotionManager) SetGyroUpdateInterval(value foundation.TimeInterval /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGyroUpdateInterval:"), value)
}


// A Boolean value that indicates whether accelerometer updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isAccelerometerActive
func (m_ MotionManager) AccelerometerActive() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("accelerometerActive"))
	return rv
}


// A Boolean value that indicates whether an accelerometer is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isAccelerometerAvailable
func (m_ MotionManager) AccelerometerAvailable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("accelerometerAvailable"))
	return rv
}


// A Boolean value that determines whether the app is receiving updates from the device-motion service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isDeviceMotionActive
func (m_ MotionManager) DeviceMotionActive() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("deviceMotionActive"))
	return rv
}


// A Boolean value that indicates whether the device-motion service is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isDeviceMotionAvailable
func (m_ MotionManager) DeviceMotionAvailable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("deviceMotionAvailable"))
	return rv
}


// A Boolean value that determines whether gyroscope updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isGyroActive
func (m_ MotionManager) GyroActive() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("gyroActive"))
	return rv
}


// A Boolean value that indicates whether a gyroscope is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isGyroAvailable
func (m_ MotionManager) GyroAvailable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("gyroAvailable"))
	return rv
}


// A Boolean value that determines whether magnetometer updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isMagnetometerActive
func (m_ MotionManager) MagnetometerActive() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("magnetometerActive"))
	return rv
}


// A Boolean value that indicates whether a magnetometer is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isMagnetometerAvailable
func (m_ MotionManager) MagnetometerAvailable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("magnetometerAvailable"))
	return rv
}


// The latest sample of magnetometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/magnetometerData
func (m_ MotionManager) MagnetometerData() ICMMagnetometerData {
	rv := objc.Send[MagnetometerData](m_.ID, objc.Sel("magnetometerData"))
	return rv
}


// The interval, in seconds, at which the system delivers magnetometer data to the block handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/magnetometerUpdateInterval
func (m_ MotionManager) MagnetometerUpdateInterval() foundation.TimeInterval /* not a class type */ {
	rv := objc.Send[foundation.TimeInterval](m_.ID, objc.Sel("magnetometerUpdateInterval"))
	return rv
}


// The interval, in seconds, at which the system delivers magnetometer data to the block handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/magnetometerUpdateInterval
func (m_ MotionManager) SetMagnetometerUpdateInterval(value foundation.TimeInterval /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMagnetometerUpdateInterval:"), value)
}


// Controls whether the device-movement display is shown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/showsDeviceMovementDisplay
func (m_ MotionManager) ShowsDeviceMovementDisplay() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsDeviceMovementDisplay"))
	return rv
}


// Controls whether the device-movement display is shown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/showsDeviceMovementDisplay
func (m_ MotionManager) SetShowsDeviceMovementDisplay(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsDeviceMovementDisplay:"), value)
}


// The error domain for Core Motion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmerrordomain
func (m_ MotionManager) CMErrorDomain() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("CMErrorDomain"))
	return rv
}


// A Boolean value that indicates whether accelerometer updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isaccelerometeractive
func (m_ MotionManager) IsAccelerometerActive() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isAccelerometerActive"))
	return rv
}


// A Boolean value that indicates whether accelerometer updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isaccelerometeractive
func (m_ MotionManager) SetIsAccelerometerActive(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsAccelerometerActive:"), value)
}


// A Boolean value that indicates whether an accelerometer is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isaccelerometeravailable
func (m_ MotionManager) IsAccelerometerAvailable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isAccelerometerAvailable"))
	return rv
}


// A Boolean value that indicates whether an accelerometer is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isaccelerometeravailable
func (m_ MotionManager) SetIsAccelerometerAvailable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsAccelerometerAvailable:"), value)
}


// A Boolean value that determines whether the app is receiving updates from the device-motion service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isdevicemotionactive
func (m_ MotionManager) IsDeviceMotionActive() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isDeviceMotionActive"))
	return rv
}


// A Boolean value that determines whether the app is receiving updates from the device-motion service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isdevicemotionactive
func (m_ MotionManager) SetIsDeviceMotionActive(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsDeviceMotionActive:"), value)
}


// A Boolean value that indicates whether the device-motion service is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isdevicemotionavailable
func (m_ MotionManager) IsDeviceMotionAvailable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isDeviceMotionAvailable"))
	return rv
}


// A Boolean value that indicates whether the device-motion service is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isdevicemotionavailable
func (m_ MotionManager) SetIsDeviceMotionAvailable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsDeviceMotionAvailable:"), value)
}


// A Boolean value that determines whether gyroscope updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isgyroactive
func (m_ MotionManager) IsGyroActive() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isGyroActive"))
	return rv
}


// A Boolean value that determines whether gyroscope updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isgyroactive
func (m_ MotionManager) SetIsGyroActive(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsGyroActive:"), value)
}


// A Boolean value that indicates whether a gyroscope is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isgyroavailable
func (m_ MotionManager) IsGyroAvailable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isGyroAvailable"))
	return rv
}


// A Boolean value that indicates whether a gyroscope is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/isgyroavailable
func (m_ MotionManager) SetIsGyroAvailable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsGyroAvailable:"), value)
}


// A Boolean value that determines whether magnetometer updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/ismagnetometeractive
func (m_ MotionManager) IsMagnetometerActive() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isMagnetometerActive"))
	return rv
}


// A Boolean value that determines whether magnetometer updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/ismagnetometeractive
func (m_ MotionManager) SetIsMagnetometerActive(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsMagnetometerActive:"), value)
}


// A Boolean value that indicates whether a magnetometer is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/ismagnetometeravailable
func (m_ MotionManager) IsMagnetometerAvailable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isMagnetometerAvailable"))
	return rv
}


// A Boolean value that indicates whether a magnetometer is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/ismagnetometeravailable
func (m_ MotionManager) SetIsMagnetometerAvailable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsMagnetometerAvailable:"), value)
}



