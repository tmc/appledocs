// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	DeviceMotion() ICMDeviceMotion
	ConnectionStatusActive() bool
	DeviceMotionActive() bool
	DeviceMotionAvailable() bool
	IsConnectionStatusActive() bool
	SetIsConnectionStatusActive(value bool)
	IsDeviceMotionActive() bool
	SetIsDeviceMotionActive(value bool)
	IsDeviceMotionAvailable() bool
	SetIsDeviceMotionAvailable(value bool)
	StartConnectionStatusUpdates()
	StartDeviceMotionUpdates()
	StartDeviceMotionUpdatesToQueueWithHandler(queue foundation.OperationQueue, handler unsafe.Pointer)
	StopConnectionStatusUpdates()
	StopDeviceMotionUpdates()
}

// An object that starts and manages headphone motion services.
//
// This class delivers headphone motion updates to your app. Use an instance of the manager to determine if the device supports motion, and to start and stop updates. Adopt the protocol to receive and respond to motion updates. Before using this class, check to make sure the feature is available.


// An object that starts and manages headphone motion services.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/authorizationStatus()
func (hc _HeadphoneMotionManagerClass) AuthorizationStatus() CMAuthorizationStatus {
	rv := objc.Send[CMAuthorizationStatus](objc.ID(hc.class), objc.Sel("authorizationStatus"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/startConnectionStatusUpdates()
func (h_ HeadphoneMotionManager) StartConnectionStatusUpdates() {
	objc.Send[objc.ID](h_.ID, objc.Sel("startConnectionStatusUpdates"))
}


// Starts device-motion updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/startDeviceMotionUpdates()
func (h_ HeadphoneMotionManager) StartDeviceMotionUpdates() {
	objc.Send[objc.ID](h_.ID, objc.Sel("startDeviceMotionUpdates"))
}


// Starts device-motion updates with a handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/startDeviceMotionUpdates(to:withHandler:)
func (h_ HeadphoneMotionManager) StartDeviceMotionUpdatesToQueueWithHandler(queue foundation.OperationQueue, handler unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("startDeviceMotionUpdatesToQueue:withHandler:"), queue, handler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/stopConnectionStatusUpdates()
func (h_ HeadphoneMotionManager) StopConnectionStatusUpdates() {
	objc.Send[objc.ID](h_.ID, objc.Sel("stopConnectionStatusUpdates"))
}


// Stops device-motion updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/stopDeviceMotionUpdates()
func (h_ HeadphoneMotionManager) StopDeviceMotionUpdates() {
	objc.Send[objc.ID](h_.ID, objc.Sel("stopDeviceMotionUpdates"))
}


// The object that receives headphone motion manager events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/delegate
func (h_ HeadphoneMotionManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](h_.ID, objc.Sel("delegate"))
	return rv
}


// The object that receives headphone motion manager events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/delegate
func (h_ HeadphoneMotionManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDelegate:"), value)
}


// The latest device-motion data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/deviceMotion
func (h_ HeadphoneMotionManager) DeviceMotion() ICMDeviceMotion {
	rv := objc.Send[DeviceMotion](h_.ID, objc.Sel("deviceMotion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/isConnectionStatusActive
func (h_ HeadphoneMotionManager) ConnectionStatusActive() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("connectionStatusActive"))
	return rv
}


// A Boolean value that indicates whether the headphone motion manager is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/isDeviceMotionActive
func (h_ HeadphoneMotionManager) DeviceMotionActive() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("deviceMotionActive"))
	return rv
}


// A Boolean value that indicates whether the current device supports the headphone motion manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneMotionManager/isDeviceMotionAvailable
func (h_ HeadphoneMotionManager) DeviceMotionAvailable() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("deviceMotionAvailable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmheadphonemotionmanager/isconnectionstatusactive
func (h_ HeadphoneMotionManager) IsConnectionStatusActive() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isConnectionStatusActive"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmheadphonemotionmanager/isconnectionstatusactive
func (h_ HeadphoneMotionManager) SetIsConnectionStatusActive(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsConnectionStatusActive:"), value)
}


// A Boolean value that indicates whether the headphone motion manager is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmheadphonemotionmanager/isdevicemotionactive
func (h_ HeadphoneMotionManager) IsDeviceMotionActive() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isDeviceMotionActive"))
	return rv
}


// A Boolean value that indicates whether the headphone motion manager is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmheadphonemotionmanager/isdevicemotionactive
func (h_ HeadphoneMotionManager) SetIsDeviceMotionActive(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsDeviceMotionActive:"), value)
}


// A Boolean value that indicates whether the current device supports the headphone motion manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmheadphonemotionmanager/isdevicemotionavailable
func (h_ HeadphoneMotionManager) IsDeviceMotionAvailable() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isDeviceMotionAvailable"))
	return rv
}


// A Boolean value that indicates whether the current device supports the headphone motion manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmheadphonemotionmanager/isdevicemotionavailable
func (h_ HeadphoneMotionManager) SetIsDeviceMotionAvailable(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsDeviceMotionAvailable:"), value)
}



