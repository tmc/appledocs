//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for WaterSubmersionEvent


// iOS-only properties

// The time and date of the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionEvent/date
func (w_ WaterSubmersionEvent) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](w_.ID, objc.Sel("date"))
	return rv
}

// The new submersion state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionEvent/state-swift.property
func (w_ WaterSubmersionEvent) State() WaterSubmersionState {
	rv := objc.Send[WaterSubmersionState](w_.ID, objc.Sel("state"))
	return rv
}





