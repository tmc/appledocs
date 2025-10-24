//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Altimeter


// Starts the delivery of absolute altitude data to the specified handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter/startAbsoluteAltitudeUpdates(to:withHandler:)
func (a_ Altimeter) StartAbsoluteAltitudeUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework: OperationQueue */, handler AbsoluteAltitudeHandler /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("startAbsoluteAltitudeUpdatesToQueue:withHandler:"), queue, handler)
}

// Starts the delivery of relative altitude data to the specified handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter/startRelativeAltitudeUpdates(to:withHandler:)
func (a_ Altimeter) StartRelativeAltitudeUpdatesToQueueWithHandler(queue objc.IObject /* cross-framework: OperationQueue */, handler AltitudeHandler /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("startRelativeAltitudeUpdatesToQueue:withHandler:"), queue, handler)
}

// Stops the delivery of absolute altitude data for this altimeter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter/stopAbsoluteAltitudeUpdates()
func (a_ Altimeter) StopAbsoluteAltitudeUpdates() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopAbsoluteAltitudeUpdates"))
}

// Stops the delivery of relative altitude data for the altimeter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter/stopRelativeAltitudeUpdates()
func (a_ Altimeter) StopRelativeAltitudeUpdates() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopRelativeAltitudeUpdates"))
}

// iOS-only properties





