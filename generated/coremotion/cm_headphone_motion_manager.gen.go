// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HeadphoneMotionManager] class.
var (
	HeadphoneMotionManagerClass     _HeadphoneMotionManagerClass
	HeadphoneMotionManagerClassOnce sync.Once
)

func getHeadphoneMotionManagerClass() _HeadphoneMotionManagerClass {
	HeadphoneMotionManagerClassOnce.Do(func() {
		HeadphoneMotionManagerClass = _HeadphoneMotionManagerClass{objc.GetClass("CMHeadphoneMotionManager")}
	})
	return HeadphoneMotionManagerClass
}

type _HeadphoneMotionManagerClass struct {
	class objc.Class
}

// An interface definition for the [HeadphoneMotionManager] class.
type IHeadphoneMotionManager interface {
	objectivec.IObject
	StartConnectionStatusUpdates()
	StartDeviceMotionUpdates()
	StartDeviceMotionUpdatesToQueueWithHandler(queue unsafe.Pointer, handler unsafe.Pointer)
	StopConnectionStatusUpdates()
	StopDeviceMotionUpdates()
}

// An object that starts and manages headphone motion services.
//
// This class delivers headphone motion updates to your app. Use an instance of the manager to determine if the device supports motion, and to start and stop updates. Adopt the protocol to receive and respond to motion updates. Before using this class, check to make sure the feature is available.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager
type HeadphoneMotionManager struct {
	objectivec.Object
}

// HeadphoneMotionManagerFrom constructs a [HeadphoneMotionManager] from an unsafe.Pointer.
//
// An object that starts and manages headphone motion services.
func HeadphoneMotionManagerFrom(ptr unsafe.Pointer) HeadphoneMotionManager {
	return HeadphoneMotionManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HeadphoneMotionManagerClass) Alloc() HeadphoneMotionManager {
	rv := objc.Send[HeadphoneMotionManager](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HeadphoneMotionManagerClass) New() HeadphoneMotionManager {
	rv := objc.Send[HeadphoneMotionManager](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HeadphoneMotionManager) Init() HeadphoneMotionManager {
	rv := objc.Send[HeadphoneMotionManager](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HeadphoneMotionManager) Autorelease() HeadphoneMotionManager {
	rv := objc.Send[HeadphoneMotionManager](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHeadphoneMotionManager creates a new HeadphoneMotionManager instance.
func NewHeadphoneMotionManager() HeadphoneMotionManager {
	return getHeadphoneMotionManagerClass().New()
}


// Returns the authorization status for monitoring headphone motion.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/authorizationStatus()
func (hc _HeadphoneMotionManagerClass) AuthorizationStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("authorizationStatus"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/startConnectionStatusUpdates()
func (h_ HeadphoneMotionManager) StartConnectionStatusUpdates() {
	objc.Send[objc.ID](h_.ID, objc.Sel("startConnectionStatusUpdates"))
}

// Starts device-motion updates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/startDeviceMotionUpdates()
func (h_ HeadphoneMotionManager) StartDeviceMotionUpdates() {
	objc.Send[objc.ID](h_.ID, objc.Sel("startDeviceMotionUpdates"))
}

// Starts device-motion updates with a handler.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/startDeviceMotionUpdates(to:withHandler:)
func (h_ HeadphoneMotionManager) StartDeviceMotionUpdatesToQueueWithHandler(queue unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("startDeviceMotionUpdatesToQueue:withHandler:"), queue, handler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/stopConnectionStatusUpdates()
func (h_ HeadphoneMotionManager) StopConnectionStatusUpdates() {
	objc.Send[objc.ID](h_.ID, objc.Sel("stopConnectionStatusUpdates"))
}

// Stops device-motion updates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/stopDeviceMotionUpdates()
func (h_ HeadphoneMotionManager) StopDeviceMotionUpdates() {
	objc.Send[objc.ID](h_.ID, objc.Sel("stopDeviceMotionUpdates"))
}

// The object that receives headphone motion manager events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/delegate
func (h_ HeadphoneMotionManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](h_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The object that receives headphone motion manager events.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/delegate
func (h_ HeadphoneMotionManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDelegate:"), value)
}
// The latest device-motion data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/deviceMotion
func (h_ HeadphoneMotionManager) DeviceMotion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("deviceMotion"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/isConnectionStatusActive
func (h_ HeadphoneMotionManager) ConnectionStatusActive() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("connectionStatusActive"))
	return rv
}

// A Boolean value that indicates whether the headphone motion manager is active.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/isDeviceMotionActive
func (h_ HeadphoneMotionManager) DeviceMotionActive() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("deviceMotionActive"))
	return rv
}

// A Boolean value that indicates whether the current device supports the headphone motion manager.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/isDeviceMotionAvailable
func (h_ HeadphoneMotionManager) DeviceMotionAvailable() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("deviceMotionAvailable"))
	return rv
}



