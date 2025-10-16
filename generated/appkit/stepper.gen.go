
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Stepper] class.
var StepperClass _StepperClass

func init() {
	StepperClass = _StepperClass{objc.GetClass("NSStepper")}
}

type _StepperClass struct {
	objc.Class
}

// An interface definition for the [Stepper] class.
type IStepper interface {
	ID() objc.ID
}

type Stepper struct {
	id objc.ID
}

func StepperFrom(ptr unsafe.Pointer) Stepper {
	return Stepper{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ Stepper) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _StepperClass) Alloc() Stepper {
	rv := objc.Send[Stepper](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _StepperClass) New() Stepper {
	rv := objc.Send[Stepper](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewStepper creates and returns a new initialized instance.
func NewStepper() Stepper {
	return StepperClass.New()
}

// Init initializes the instance.
func (s_ Stepper) Init() Stepper {
	rv := objc.Send[Stepper](s_.ID(), selInit)
	return rv
}
