
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SliderAccessory] class.
var SliderAccessoryClass _SliderAccessoryClass

func init() {
	SliderAccessoryClass = _SliderAccessoryClass{objc.GetClass("NSSliderAccessory")}
}

type _SliderAccessoryClass struct {
	objc.Class
}

// An interface definition for the [SliderAccessory] class.
type ISliderAccessory interface {
	ID() objc.ID
}

type SliderAccessory struct {
	id objc.ID
}

func SliderAccessoryFrom(ptr unsafe.Pointer) SliderAccessory {
	return SliderAccessory{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SliderAccessory) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SliderAccessoryClass) Alloc() SliderAccessory {
	rv := objc.Send[SliderAccessory](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SliderAccessoryClass) New() SliderAccessory {
	rv := objc.Send[SliderAccessory](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSliderAccessory creates and returns a new initialized instance.
func NewSliderAccessory() SliderAccessory {
	return SliderAccessoryClass.New()
}

// Init initializes the instance.
func (s_ SliderAccessory) Init() SliderAccessory {
	rv := objc.Send[SliderAccessory](s_.ID(), selInit)
	return rv
}
