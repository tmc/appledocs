// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AccessibilityCustomAction] class.
var (
	accessibilityCustomActionClass     _AccessibilityCustomActionClass
	accessibilityCustomActionClassOnce sync.Once
)

func getAccessibilityCustomActionClass() _AccessibilityCustomActionClass {
	accessibilityCustomActionClassOnce.Do(func() {
		accessibilityCustomActionClass = _AccessibilityCustomActionClass{objc.GetClass("NSAccessibilityCustomAction")}
	})
	return accessibilityCustomActionClass
}

type _AccessibilityCustomActionClass struct {
	class objc.Class
}

// An interface definition for the [AccessibilityCustomAction] class.
type IAccessibilityCustomAction interface {
	objectivec.IObject
}

// A custom action to perform on an accessible object.
//
// Apps that support custom actions can create instances of this class, specifying the user-readable name of the action, and either a handler closure or the object and selector to use when performing the action. Assistive apps display custom actions in response to specific user cues. For example, VoiceOver lets users access actions quickly using the Actions rotor. After creating an instance of this class, add it to the property of an appropriate accessible object.
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

// Alloc allocates a new instance without initialization.
func (ac _AccessibilityCustomActionClass) Alloc() AccessibilityCustomAction {
	rv := objc.Send[AccessibilityCustomAction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccessibilityCustomActionClass) New() AccessibilityCustomAction {
	rv := objc.Send[AccessibilityCustomAction](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccessibilityCustomAction) Init() AccessibilityCustomAction {
	rv := objc.Send[AccessibilityCustomAction](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccessibilityCustomAction) Autorelease() AccessibilityCustomAction {
	rv := objc.Send[AccessibilityCustomAction](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccessibilityCustomAction creates a new AccessibilityCustomAction instance.
func NewAccessibilityCustomAction() AccessibilityCustomAction {
	return getAccessibilityCustomActionClass().New()
}




