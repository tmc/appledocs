//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MotionActivityManager


// Gathers and returns historical motion data for the specified time period
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityManager/queryActivityStarting(from:to:to:withHandler:)
func (m_ MotionActivityManager) QueryActivityStartingFromDateToDateToQueueWithHandler(start objc.IObject /* cross-framework: NSDate */, end objc.IObject /* cross-framework: NSDate */, queue objc.IObject /* cross-framework: OperationQueue */, handler MotionActivityQueryHandler /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("queryActivityStartingFromDate:toDate:toQueue:withHandler:"), start, end, queue, handler)
}

// Starts the delivery of current motion data updates to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityManager/startActivityUpdates(to:withHandler:)
func (m_ MotionActivityManager) StartActivityUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework: OperationQueue */, handler MotionActivityHandler /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startActivityUpdatesToQueue:withHandler:"), queue, handler)
}

// Stops the delivery of motion updates to your app
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityManager/stopActivityUpdates()
func (m_ MotionActivityManager) StopActivityUpdates() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopActivityUpdates"))
}

// iOS-only properties





