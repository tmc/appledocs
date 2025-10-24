//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for WaterSubmersionManager


// iOS-only properties

// The object that receives updates about submersion data and events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionManager/delegate
func (w_ WaterSubmersionManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("delegate"))
	return rv
}
func (w_ WaterSubmersionManager) SetDelegate(value objc.ID) {
	w_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// The maximum depth supported by the water submersion manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionManager/maximumDepth
func (w_ WaterSubmersionManager) MaximumDepth() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("maximumDepth"))
	return rv
}





