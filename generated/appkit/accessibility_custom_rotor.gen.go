// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AccessibilityCustomRotor] class.
var AccessibilityCustomRotorClass objc.Class

func init() {
	AccessibilityCustomRotorClass = objc.GetClass("NSAccessibilityCustomRotor")
}

type AccessibilityCustomRotor struct {
	objc.ID
}

func AccessibilityCustomRotorFrom(ptr unsafe.Pointer) AccessibilityCustomRotor {
	return AccessibilityCustomRotor{
		ID: objc.ID(ptr),
	}
}



