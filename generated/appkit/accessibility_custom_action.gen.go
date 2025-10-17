
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AccessibilityCustomAction] class.
var AccessibilityCustomActionClass _AccessibilityCustomActionClass

func init() {
	AccessibilityCustomActionClass = _AccessibilityCustomActionClass{objc.GetClass("NSAccessibilityCustomAction")}
}

type _AccessibilityCustomActionClass struct {
	objc.Class
}

// An interface definition for the [AccessibilityCustomAction] class.
type IAccessibilityCustomAction interface {
	ID() objc.ID
}

type AccessibilityCustomAction struct {
	id objc.ID
}

func AccessibilityCustomActionFrom(ptr unsafe.Pointer) AccessibilityCustomAction {
	return AccessibilityCustomAction{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ AccessibilityCustomAction) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _AccessibilityCustomActionClass) Alloc() AccessibilityCustomAction {
	rv := objc.Send[AccessibilityCustomAction](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _AccessibilityCustomActionClass) New() AccessibilityCustomAction {
	rv := objc.Send[AccessibilityCustomAction](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewAccessibilityCustomAction creates and returns a new initialized instance.
func NewAccessibilityCustomAction() AccessibilityCustomAction {
	return AccessibilityCustomActionClass.New()
}

// Init initializes the instance.
func (a_ AccessibilityCustomAction) Init() AccessibilityCustomAction {
	rv := objc.Send[AccessibilityCustomAction](a_.ID(), selInit)
	return rv
}
