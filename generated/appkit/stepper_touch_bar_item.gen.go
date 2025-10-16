
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StepperTouchBarItem] class.
var StepperTouchBarItemClass _StepperTouchBarItemClass

func init() {
	StepperTouchBarItemClass = _StepperTouchBarItemClass{objc.GetClass("NSStepperTouchBarItem")}
}

type _StepperTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [StepperTouchBarItem] class.
type IStepperTouchBarItem interface {
	ID() objc.ID
}

type StepperTouchBarItem struct {
	id objc.ID
}

func StepperTouchBarItemFrom(ptr unsafe.Pointer) StepperTouchBarItem {
	return StepperTouchBarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ StepperTouchBarItem) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _StepperTouchBarItemClass) Alloc() StepperTouchBarItem {
	rv := objc.Send[StepperTouchBarItem](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _StepperTouchBarItemClass) New() StepperTouchBarItem {
	rv := objc.Send[StepperTouchBarItem](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewStepperTouchBarItem creates and returns a new initialized instance.
func NewStepperTouchBarItem() StepperTouchBarItem {
	return StepperTouchBarItemClass.New()
}

// Init initializes the instance.
func (s_ StepperTouchBarItem) Init() StepperTouchBarItem {
	rv := objc.Send[StepperTouchBarItem](s_.ID(), selInit)
	return rv
}
