//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for BatchedSensorManager


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/startAccelerometerUpdates()
func (b_ BatchedSensorManager) StartAccelerometerUpdates() {
	objc.Send[objc.ID](b_.ID, objc.Sel("startAccelerometerUpdates"))
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/startAccelerometerUpdates(handler:)
func (b_ BatchedSensorManager) StartAccelerometerUpdatesWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("startAccelerometerUpdatesWithHandler:"), handler)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/startDeviceMotionUpdates()
func (b_ BatchedSensorManager) StartDeviceMotionUpdates() {
	objc.Send[objc.ID](b_.ID, objc.Sel("startDeviceMotionUpdates"))
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/startDeviceMotionUpdates(handler:)
func (b_ BatchedSensorManager) StartDeviceMotionUpdatesWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("startDeviceMotionUpdatesWithHandler:"), handler)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/stopAccelerometerUpdates()
func (b_ BatchedSensorManager) StopAccelerometerUpdates() {
	objc.Send[objc.ID](b_.ID, objc.Sel("stopAccelerometerUpdates"))
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/stopDeviceMotionUpdates()
func (b_ BatchedSensorManager) StopDeviceMotionUpdates() {
	objc.Send[objc.ID](b_.ID, objc.Sel("stopDeviceMotionUpdates"))
}

// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/accelerometerBatch
func (b_ BatchedSensorManager) AccelerometerBatch() []IAccelerometerData {
	rv := objc.Send[[]AccelerometerData](b_.ID, objc.Sel("accelerometerBatch"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/accelerometerDataFrequency
func (b_ BatchedSensorManager) AccelerometerDataFrequency() int {
	rv := objc.Send[int](b_.ID, objc.Sel("accelerometerDataFrequency"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/deviceMotionBatch
func (b_ BatchedSensorManager) DeviceMotionBatch() []IDeviceMotion {
	rv := objc.Send[[]DeviceMotion](b_.ID, objc.Sel("deviceMotionBatch"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/deviceMotionDataFrequency
func (b_ BatchedSensorManager) DeviceMotionDataFrequency() int {
	rv := objc.Send[int](b_.ID, objc.Sel("deviceMotionDataFrequency"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/isAccelerometerActive
func (b_ BatchedSensorManager) AccelerometerActive() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("accelerometerActive"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/isDeviceMotionActive
func (b_ BatchedSensorManager) DeviceMotionActive() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("deviceMotionActive"))
	return rv
}





