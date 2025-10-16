
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AccessibilityCustomRotor] class.
var AccessibilityCustomRotorClass _AccessibilityCustomRotorClass

func init() {
	AccessibilityCustomRotorClass = _AccessibilityCustomRotorClass{objc.GetClass("NSAccessibilityCustomRotor")}
}

type _AccessibilityCustomRotorClass struct {
	objc.Class
}

// An interface definition for the [AccessibilityCustomRotor] class.
type IAccessibilityCustomRotor interface {
	ID() objc.ID
}

type AccessibilityCustomRotor struct {
	id objc.ID
}

func AccessibilityCustomRotorFrom(ptr unsafe.Pointer) AccessibilityCustomRotor {
	return AccessibilityCustomRotor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ AccessibilityCustomRotor) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _AccessibilityCustomRotorClass) Alloc() AccessibilityCustomRotor {
	rv := objc.Send[AccessibilityCustomRotor](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _AccessibilityCustomRotorClass) New() AccessibilityCustomRotor {
	rv := objc.Send[AccessibilityCustomRotor](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewAccessibilityCustomRotor creates and returns a new initialized instance.
func NewAccessibilityCustomRotor() AccessibilityCustomRotor {
	return AccessibilityCustomRotorClass.New()
}

// Init initializes the instance.
func (a_ AccessibilityCustomRotor) Init() AccessibilityCustomRotor {
	rv := objc.Send[AccessibilityCustomRotor](a_.ID(), selInit)
	return rv
}
