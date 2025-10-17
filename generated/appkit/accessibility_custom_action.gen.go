// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AccessibilityCustomAction] class.
var accessibilityCustomActionClass = _AccessibilityCustomActionClass{objc.GetClass("NSAccessibilityCustomAction")}

type _AccessibilityCustomActionClass struct {
	class objc.Class
}

// A custom action to perform on an accessible object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction

type AccessibilityCustomAction struct {
	objectivec.Object
}

// AccessibilityCustomActionFrom constructs a [AccessibilityCustomAction] from an unsafe.Pointer.
//
// A custom action to perform on an accessible object.
func AccessibilityCustomActionFrom(ptr unsafe.Pointer) AccessibilityCustomAction {
	return AccessibilityCustomAction{objectivec.Object{objc.ID(ptr)}}
}



