//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for MotionActivity


// iOS-only properties

// A Boolean indicating whether the device is in a bicycle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivity/cycling
func (m_ MotionActivity) Cycling() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("cycling"))
	return rv
}





