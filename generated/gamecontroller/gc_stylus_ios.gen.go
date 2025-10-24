//go:build darwin && ios

// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for GCStylus


// iOS-only properties

// Gets the haptics profile for the stylus, if supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCStylus/haptics
func (g_ GCStylus) Haptics() IGCDeviceHaptics {
	rv := objc.Send[GCDeviceHaptics](g_.ID, objc.Sel("haptics"))
	return rv
}





