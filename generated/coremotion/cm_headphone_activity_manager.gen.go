// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HeadphoneActivityManager] class.
var (
	HeadphoneActivityManagerClass     _HeadphoneActivityManagerClass
	HeadphoneActivityManagerClassOnce sync.Once
)

func getHeadphoneActivityManagerClass() _HeadphoneActivityManagerClass {
	HeadphoneActivityManagerClassOnce.Do(func() {
		HeadphoneActivityManagerClass = _HeadphoneActivityManagerClass{objc.GetClass("CMHeadphoneActivityManager")}
	})
	return HeadphoneActivityManagerClass
}

type _HeadphoneActivityManagerClass struct {
	class objc.Class
}

// An interface definition for the [HeadphoneActivityManager] class.
type IHeadphoneActivityManager interface {
	objectivec.IObject
	// properties:
	ActivityActive() bool
	ActivityAvailable() bool
	StatusActive() bool
	StatusAvailable() bool
	IsActivityActive() bool
	SetIsActivityActive(value bool)
	IsActivityAvailable() bool
	SetIsActivityAvailable(value bool)
	IsStatusActive() bool
	SetIsStatusActive(value bool)
	IsStatusAvailable() bool
	SetIsStatusAvailable(value bool)
	// methods:
	StartActivityUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework: OperationQueue */, handler HeadphoneActivityHandler /* not a class type */)
	StartStatusUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework: OperationQueue */, handler HeadphoneActivityStatusHandler /* not a class type */)
	StopActivityUpdates()
	StopStatusUpdates()
}

// An object that starts and manages headphone activity services.
//
// This class delivers headphone activity updates to your app. Use an instance of the manager to determine if the device supports headphone activity updates, and to start and stop updates. Before using this class, check and to make sure the features are available. This class provides similar information to , except the activity information comes from headphone motion, rather than from device motion.


// An object that starts and manages headphone activity services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager
type HeadphoneActivityManager struct {
	objectivec.Object
}

// HeadphoneActivityManagerFrom constructs a [HeadphoneActivityManager] from an unsafe.Pointer.
//
// An object that starts and manages headphone activity services.
func HeadphoneActivityManagerFrom(ptr unsafe.Pointer) HeadphoneActivityManager {
	return HeadphoneActivityManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HeadphoneActivityManagerClass) Alloc() HeadphoneActivityManager {
	rv := objc.Send[HeadphoneActivityManager](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HeadphoneActivityManagerClass) New() HeadphoneActivityManager {
	rv := objc.Send[HeadphoneActivityManager](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HeadphoneActivityManager) Init() HeadphoneActivityManager {
	rv := objc.Send[HeadphoneActivityManager](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HeadphoneActivityManager) Autorelease() HeadphoneActivityManager {
	rv := objc.Send[HeadphoneActivityManager](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHeadphoneActivityManager creates a new HeadphoneActivityManager instance.
func NewHeadphoneActivityManager() HeadphoneActivityManager {
	return getHeadphoneActivityManagerClass().New()
}



// Returns the authorization status for monitoring headphone activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager/authorizationStatus()
func (hc _HeadphoneActivityManagerClass) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](objc.ID(hc.class), objc.Sel("authorizationStatus"))
	return rv
}


// Starts headphone activity updates, providing data to the given handler through the given queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager/startActivityUpdates(to:withHandler:)
func (h_ HeadphoneActivityManager) StartActivityUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework: OperationQueue */, handler HeadphoneActivityHandler /* not a class type */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("startActivityUpdatesToQueue:withHandler:"), queue, handler)
}


// Starts headphone status updates, providing data to the given handler through the given queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager/startStatusUpdates(to:withHandler:)
func (h_ HeadphoneActivityManager) StartStatusUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework: OperationQueue */, handler HeadphoneActivityStatusHandler /* not a class type */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("startStatusUpdatesToQueue:withHandler:"), queue, handler)
}


// Stops headphone activity updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager/stopActivityUpdates()
func (h_ HeadphoneActivityManager) StopActivityUpdates() {
	objc.Send[objc.ID](h_.ID, objc.Sel("stopActivityUpdates"))
}


// Stops headphone status updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager/stopStatusUpdates()
func (h_ HeadphoneActivityManager) StopStatusUpdates() {
	objc.Send[objc.ID](h_.ID, objc.Sel("stopStatusUpdates"))
}


// A Boolean value that indicates whether headphone motion activity is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager/isActivityActive
func (h_ HeadphoneActivityManager) ActivityActive() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("activityActive"))
	return rv
}


// A Boolean value that indicates whether the current device supports headphone activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager/isActivityAvailable
func (h_ HeadphoneActivityManager) ActivityAvailable() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("activityAvailable"))
	return rv
}


// A Boolean value that indicates whether headphone status is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager/isStatusActive
func (h_ HeadphoneActivityManager) StatusActive() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("statusActive"))
	return rv
}


// A Boolean value that indicates whether the current device supports headphone status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager/isStatusAvailable
func (h_ HeadphoneActivityManager) StatusAvailable() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("statusAvailable"))
	return rv
}


// A Boolean value that indicates whether headphone motion activity is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmheadphoneactivitymanager/isactivityactive
func (h_ HeadphoneActivityManager) IsActivityActive() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isActivityActive"))
	return rv
}


// A Boolean value that indicates whether headphone motion activity is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmheadphoneactivitymanager/isactivityactive
func (h_ HeadphoneActivityManager) SetIsActivityActive(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsActivityActive:"), value)
}


// A Boolean value that indicates whether the current device supports headphone activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmheadphoneactivitymanager/isactivityavailable
func (h_ HeadphoneActivityManager) IsActivityAvailable() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isActivityAvailable"))
	return rv
}


// A Boolean value that indicates whether the current device supports headphone activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmheadphoneactivitymanager/isactivityavailable
func (h_ HeadphoneActivityManager) SetIsActivityAvailable(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsActivityAvailable:"), value)
}


// A Boolean value that indicates whether headphone status is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmheadphoneactivitymanager/isstatusactive
func (h_ HeadphoneActivityManager) IsStatusActive() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isStatusActive"))
	return rv
}


// A Boolean value that indicates whether headphone status is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmheadphoneactivitymanager/isstatusactive
func (h_ HeadphoneActivityManager) SetIsStatusActive(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsStatusActive:"), value)
}


// A Boolean value that indicates whether the current device supports headphone status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmheadphoneactivitymanager/isstatusavailable
func (h_ HeadphoneActivityManager) IsStatusAvailable() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isStatusAvailable"))
	return rv
}


// A Boolean value that indicates whether the current device supports headphone status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmheadphoneactivitymanager/isstatusavailable
func (h_ HeadphoneActivityManager) SetIsStatusAvailable(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsStatusAvailable:"), value)
}



