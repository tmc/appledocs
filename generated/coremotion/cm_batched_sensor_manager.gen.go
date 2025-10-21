// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [BatchedSensorManager] class.
var (
	BatchedSensorManagerClass     _BatchedSensorManagerClass
	BatchedSensorManagerClassOnce sync.Once
)

func getBatchedSensorManagerClass() _BatchedSensorManagerClass {
	BatchedSensorManagerClassOnce.Do(func() {
		BatchedSensorManagerClass = _BatchedSensorManagerClass{objc.GetClass("CMBatchedSensorManager")}
	})
	return BatchedSensorManagerClass
}

type _BatchedSensorManagerClass struct {
	class objc.Class
}

// An interface definition for the [BatchedSensorManager] class.
type IBatchedSensorManager interface {
	objectivec.IObject
	StartAccelerometerUpdates()
	StartAccelerometerUpdatesWithHandler(handler unsafe.Pointer)
	StartDeviceMotionUpdates()
	StartDeviceMotionUpdatesWithHandler(handler unsafe.Pointer)
	StopAccelerometerUpdates()
	StopDeviceMotionUpdates()
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager
type BatchedSensorManager struct {
	objectivec.Object
}

// BatchedSensorManagerFrom constructs a [BatchedSensorManager] from an unsafe.Pointer.
func BatchedSensorManagerFrom(ptr unsafe.Pointer) BatchedSensorManager {
	return BatchedSensorManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BatchedSensorManagerClass) Alloc() BatchedSensorManager {
	rv := objc.Send[BatchedSensorManager](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BatchedSensorManagerClass) New() BatchedSensorManager {
	rv := objc.Send[BatchedSensorManager](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BatchedSensorManager) Init() BatchedSensorManager {
	rv := objc.Send[BatchedSensorManager](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BatchedSensorManager) Autorelease() BatchedSensorManager {
	rv := objc.Send[BatchedSensorManager](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBatchedSensorManager creates a new BatchedSensorManager instance.
func NewBatchedSensorManager() BatchedSensorManager {
	return getBatchedSensorManagerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/authorizationStatus
func (bc _BatchedSensorManagerClass) AuthorizationStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("authorizationStatus"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/isAccelerometerSupported
func (bc _BatchedSensorManagerClass) AccelerometerSupported() bool {
	rv := objc.Send[bool](objc.ID(bc.class), objc.Sel("accelerometerSupported"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/isDeviceMotionSupported
func (bc _BatchedSensorManagerClass) DeviceMotionSupported() bool {
	rv := objc.Send[bool](objc.ID(bc.class), objc.Sel("deviceMotionSupported"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/startAccelerometerUpdates()
func (b_ BatchedSensorManager) StartAccelerometerUpdates() {
	objc.Send[objc.ID](b_.ID, objc.Sel("startAccelerometerUpdates"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/startAccelerometerUpdates(handler:)
func (b_ BatchedSensorManager) StartAccelerometerUpdatesWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("startAccelerometerUpdatesWithHandler:"), handler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/startDeviceMotionUpdates()
func (b_ BatchedSensorManager) StartDeviceMotionUpdates() {
	objc.Send[objc.ID](b_.ID, objc.Sel("startDeviceMotionUpdates"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/startDeviceMotionUpdates(handler:)
func (b_ BatchedSensorManager) StartDeviceMotionUpdatesWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("startDeviceMotionUpdatesWithHandler:"), handler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/stopAccelerometerUpdates()
func (b_ BatchedSensorManager) StopAccelerometerUpdates() {
	objc.Send[objc.ID](b_.ID, objc.Sel("stopAccelerometerUpdates"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/stopDeviceMotionUpdates()
func (b_ BatchedSensorManager) StopDeviceMotionUpdates() {
	objc.Send[objc.ID](b_.ID, objc.Sel("stopDeviceMotionUpdates"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/accelerometerBatch
func (b_ BatchedSensorManager) AccelerometerBatch() []AccelerometerData {
	rv := objc.Send[[]AccelerometerData](b_.ID, objc.Sel("accelerometerBatch"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/accelerometerDataFrequency
func (b_ BatchedSensorManager) AccelerometerDataFrequency() int {
	rv := objc.Send[int](b_.ID, objc.Sel("accelerometerDataFrequency"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/authorizationStatus
func (b_ BatchedSensorManager) AuthorizationStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("authorizationStatus"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/deviceMotionBatch
func (b_ BatchedSensorManager) DeviceMotionBatch() []DeviceMotion {
	rv := objc.Send[[]DeviceMotion](b_.ID, objc.Sel("deviceMotionBatch"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/deviceMotionDataFrequency
func (b_ BatchedSensorManager) DeviceMotionDataFrequency() int {
	rv := objc.Send[int](b_.ID, objc.Sel("deviceMotionDataFrequency"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/isAccelerometerActive
func (b_ BatchedSensorManager) AccelerometerActive() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("accelerometerActive"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/isAccelerometerSupported
func (b_ BatchedSensorManager) AccelerometerSupported() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("accelerometerSupported"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/isDeviceMotionActive
func (b_ BatchedSensorManager) DeviceMotionActive() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("deviceMotionActive"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMBatchedSensorManager/isDeviceMotionSupported
func (b_ BatchedSensorManager) DeviceMotionSupported() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("deviceMotionSupported"))
	return rv
}



