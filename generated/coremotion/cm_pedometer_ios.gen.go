//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Pedometer


// Starts the delivery of pedometer events to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/startEventUpdates(handler:)
func (p_ Pedometer) StartPedometerEventUpdatesWithHandler(handler PedometerEventHandler /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("startPedometerEventUpdatesWithHandler:"), handler)
}

// Stops the delivery of pedometer events to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometer/stopEventUpdates()
func (p_ Pedometer) StopPedometerEventUpdates() {
	objc.Send[objc.ID](p_.ID, objc.Sel("stopPedometerEventUpdates"))
}

// iOS-only properties





