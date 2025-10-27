//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureSystemPressureState


// iOS-only properties

// The set of underlying causes for the system pressure level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/SystemPressureState-swift.class/factors-swift.property
func (c_ CaptureSystemPressureState) Factors() CaptureSystemPressureFactors {
	rv := objc.Send[CaptureSystemPressureFactors](c_.ID, objc.Sel("factors"))
	return rv
}

// The overall level of performance constraints on the capture system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/SystemPressureState-swift.class/level-swift.property
func (c_ CaptureSystemPressureState) Level() CaptureSystemPressureLevel {
	rv := objc.Send[CaptureSystemPressureLevel](c_.ID, objc.Sel("level"))
	return rv
}





