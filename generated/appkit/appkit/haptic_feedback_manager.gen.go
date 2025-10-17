// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [HapticFeedbackManager] class.
var HapticFeedbackManagerClass objc.Class

func init() {
	HapticFeedbackManagerClass = objc.GetClass("NSHapticFeedbackManager")
}

type HapticFeedbackManager struct {
	objc.ID
}

func HapticFeedbackManagerFrom(ptr unsafe.Pointer) HapticFeedbackManager {
	return HapticFeedbackManager{
		ID: objc.ID(ptr),
	}
}




