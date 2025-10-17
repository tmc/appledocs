// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AccessibilityCustomAction] class.
var AccessibilityCustomActionClass objc.Class

func init() {
	AccessibilityCustomActionClass = objc.GetClass("NSAccessibilityCustomAction")
}

type AccessibilityCustomAction struct {
	objc.ID
}

func AccessibilityCustomActionFrom(ptr unsafe.Pointer) AccessibilityCustomAction {
	return AccessibilityCustomAction{
		ID: objc.ID(ptr),
	}
}



