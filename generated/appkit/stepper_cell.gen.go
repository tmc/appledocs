
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StepperCell] class.
var StepperCellClass _StepperCellClass

func init() {
	StepperCellClass = _StepperCellClass{objc.GetClass("NSStepperCell")}
}

type _StepperCellClass struct {
	objc.Class
}

// An interface definition for the [StepperCell] class.
type IStepperCell interface {
	ID() objc.ID
}

type StepperCell struct {
	id objc.ID
}

func StepperCellFrom(ptr unsafe.Pointer) StepperCell {
	return StepperCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ StepperCell) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _StepperCellClass) Alloc() StepperCell {
	rv := objc.Send[StepperCell](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _StepperCellClass) New() StepperCell {
	rv := objc.Send[StepperCell](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewStepperCell creates and returns a new initialized instance.
func NewStepperCell() StepperCell {
	return StepperCellClass.New()
}

// Init initializes the instance.
func (s_ StepperCell) Init() StepperCell {
	rv := objc.Send[StepperCell](s_.ID(), selInit)
	return rv
}
