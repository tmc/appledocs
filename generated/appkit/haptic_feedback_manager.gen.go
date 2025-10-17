// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HapticFeedbackManager] class.
var hapticFeedbackManagerClass = _HapticFeedbackManagerClass{objc.GetClass("NSHapticFeedbackManager")}

type _HapticFeedbackManagerClass struct {
	class objc.Class
}

// An interface definition for the [HapticFeedbackManager] class.
type IHapticFeedbackManager interface {
	objectivec.IObject
}

// An object that provides access to the haptic feedback management attributes on a system with a Force Touch trackpad. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager

type HapticFeedbackManager struct {
	objectivec.Object
}

// HapticFeedbackManagerFrom constructs a [HapticFeedbackManager] from an unsafe.Pointer.
//
// An object that provides access to the haptic feedback management attributes on a system with a Force Touch trackpad.
func HapticFeedbackManagerFrom(ptr unsafe.Pointer) HapticFeedbackManager {
	return HapticFeedbackManager{objectivec.Object{objc.ID(ptr)}}
}



