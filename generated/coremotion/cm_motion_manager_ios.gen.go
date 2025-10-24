//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MotionManager


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
func (m_ MotionManager) StartAccelerometerUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework: OperationQueue */, handler AccelerometerHandler /* not a class type */) {
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
func (m_ MotionManager) StartDeviceMotionUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework: OperationQueue */, handler DeviceMotionHandler /* not a class type */) {
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
func (m_ MotionManager) StartDeviceMotionUpdatesUsingReferenceFrameToQueueWithHandler(referenceFrame AttitudeReferenceFrame, queue objc.IObject /* cross-framework: OperationQueue */, handler DeviceMotionHandler /* not a class type */) {
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
func (m_ MotionManager) StartGyroUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework: OperationQueue */, handler GyroHandler /* not a class type */) {
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
func (m_ MotionManager) StartMagnetometerUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework: OperationQueue */, handler MagnetometerHandler /* not a class type */) {
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

// iOS-only properties

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
func (m_ MotionManager) AccelerometerUpdateInterval() float64 {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("accelerometerUpdateInterval"))
	return rv
}
func (m_ MotionManager) SetAccelerometerUpdateInterval(value float64) {
	m_.ID.Send(objc.RegisterName("setAccelerometerUpdateInterval:"), value)
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
func (m_ MotionManager) DeviceMotionUpdateInterval() float64 {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("deviceMotionUpdateInterval"))
	return rv
}
func (m_ MotionManager) SetDeviceMotionUpdateInterval(value float64) {
	m_.ID.Send(objc.RegisterName("setDeviceMotionUpdateInterval:"), value)
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
func (m_ MotionManager) GyroUpdateInterval() float64 {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("gyroUpdateInterval"))
	return rv
}
func (m_ MotionManager) SetGyroUpdateInterval(value float64) {
	m_.ID.Send(objc.RegisterName("setGyroUpdateInterval:"), value)
}

// A Boolean value that indicates whether accelerometer updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isAccelerometerActive
func (m_ MotionManager) AccelerometerActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("accelerometerActive"))
	return rv
}

// A Boolean value that indicates whether an accelerometer is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isAccelerometerAvailable
func (m_ MotionManager) AccelerometerAvailable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("accelerometerAvailable"))
	return rv
}

// A Boolean value that determines whether the app is receiving updates from the device-motion service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isDeviceMotionActive
func (m_ MotionManager) DeviceMotionActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("deviceMotionActive"))
	return rv
}

// A Boolean value that indicates whether the device-motion service is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isDeviceMotionAvailable
func (m_ MotionManager) DeviceMotionAvailable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("deviceMotionAvailable"))
	return rv
}

// A Boolean value that determines whether gyroscope updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isGyroActive
func (m_ MotionManager) GyroActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("gyroActive"))
	return rv
}

// A Boolean value that indicates whether a gyroscope is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isGyroAvailable
func (m_ MotionManager) GyroAvailable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("gyroAvailable"))
	return rv
}

// A Boolean value that determines whether magnetometer updates are currently happening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isMagnetometerActive
func (m_ MotionManager) MagnetometerActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("magnetometerActive"))
	return rv
}

// A Boolean value that indicates whether a magnetometer is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/isMagnetometerAvailable
func (m_ MotionManager) MagnetometerAvailable() bool {
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
func (m_ MotionManager) MagnetometerUpdateInterval() float64 {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("magnetometerUpdateInterval"))
	return rv
}
func (m_ MotionManager) SetMagnetometerUpdateInterval(value float64) {
	m_.ID.Send(objc.RegisterName("setMagnetometerUpdateInterval:"), value)
}

// Controls whether the device-movement display is shown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionManager/showsDeviceMovementDisplay
func (m_ MotionManager) ShowsDeviceMovementDisplay() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsDeviceMovementDisplay"))
	return rv
}
func (m_ MotionManager) SetShowsDeviceMovementDisplay(value bool) {
	m_.ID.Send(objc.RegisterName("setShowsDeviceMovementDisplay:"), value)
}





