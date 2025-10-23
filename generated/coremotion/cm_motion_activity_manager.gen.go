// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MotionActivityManager] class.
var (
	MotionActivityManagerClass     _MotionActivityManagerClass
	MotionActivityManagerClassOnce sync.Once
)

func getMotionActivityManagerClass() _MotionActivityManagerClass {
	MotionActivityManagerClassOnce.Do(func() {
		MotionActivityManagerClass = _MotionActivityManagerClass{objc.GetClass("CMMotionActivityManager")}
	})
	return MotionActivityManagerClass
}

type _MotionActivityManagerClass struct {
	class objc.Class
}

// An interface definition for the [MotionActivityManager] class.
type IMotionActivityManager interface {
	objectivec.IObject
	QueryActivityStartingFromDateToDateToQueueWithHandler(start foundation.IDate, end foundation.IDate, queue foundation.IOperationQueue, handler unsafe.Pointer)
	StartActivityUpdatesToQueueWithHandler(queue foundation.IOperationQueue, handler unsafe.Pointer)
	StopActivityUpdates()
}

// An object that manages access to the motion data stored by the device.
//
// Motion data reflects whether the user is walking, running, in a vehicle, or stationary for periods of time. Using this class, you can ask for notifications when the current type of motion changes or you can gather past motion change data. For example, a navigation app might look for changes in the current type of motion and offer different directions for each.


// An object that manages access to the motion data stored by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityManager
type MotionActivityManager struct {
	objectivec.Object
}

// MotionActivityManagerFrom constructs a [MotionActivityManager] from an unsafe.Pointer.
//
// An object that manages access to the motion data stored by the device.
func MotionActivityManagerFrom(ptr unsafe.Pointer) MotionActivityManager {
	return MotionActivityManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MotionActivityManagerClass) Alloc() MotionActivityManager {
	rv := objc.Send[MotionActivityManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MotionActivityManagerClass) New() MotionActivityManager {
	rv := objc.Send[MotionActivityManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MotionActivityManager) Init() MotionActivityManager {
	rv := objc.Send[MotionActivityManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MotionActivityManager) Autorelease() MotionActivityManager {
	rv := objc.Send[MotionActivityManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMotionActivityManager creates a new MotionActivityManager instance.
func NewMotionActivityManager() MotionActivityManager {
	return getMotionActivityManagerClass().New()
}



// Returns a value indicating whether the app is authorized to retrieve stored motion data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityManager/authorizationStatus()
func (mc _MotionActivityManagerClass) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](objc.ID(mc.class), objc.Sel("authorizationStatus"))
	return rv
}


// Returns a Boolean indicating whether motion data is available on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityManager/isActivityAvailable()
func (mc _MotionActivityManagerClass) IsActivityAvailable() bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("isActivityAvailable"))
	return rv
}


// Gathers and returns historical motion data for the specified time period
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityManager/queryActivityStarting(from:to:to:withHandler:)
func (m_ MotionActivityManager) QueryActivityStartingFromDateToDateToQueueWithHandler(start foundation.IDate, end foundation.IDate, queue foundation.IOperationQueue, handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("queryActivityStartingFromDate:toDate:toQueue:withHandler:"), start, end, queue, handler)
}


// Starts the delivery of current motion data updates to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityManager/startActivityUpdates(to:withHandler:)
func (m_ MotionActivityManager) StartActivityUpdatesToQueueWithHandler(queue foundation.IOperationQueue, handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startActivityUpdatesToQueue:withHandler:"), queue, handler)
}


// Stops the delivery of motion updates to your app
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityManager/stopActivityUpdates()
func (m_ MotionActivityManager) StopActivityUpdates() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopActivityUpdates"))
}



