// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AccessibilityCustomRotor] class.
var accessibilityCustomRotorClass = _AccessibilityCustomRotorClass{objc.GetClass("NSAccessibilityCustomRotor")}

type _AccessibilityCustomRotorClass struct {
	class objc.Class
}

// An interface definition for the [AccessibilityCustomRotor] class.
type IAccessibilityCustomRotor interface {
	objectivec.IObject
}

// A context-sensitive function that helps VoiceOver users find the next instance of a related accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor

type AccessibilityCustomRotor struct {
	objectivec.Object
}

// AccessibilityCustomRotorFrom constructs a [AccessibilityCustomRotor] from an unsafe.Pointer.
//
// A context-sensitive function that helps VoiceOver users find the next instance of a related accessibility element.
func AccessibilityCustomRotorFrom(ptr unsafe.Pointer) AccessibilityCustomRotor {
	return AccessibilityCustomRotor{objectivec.Object{objc.ID(ptr)}}
}



