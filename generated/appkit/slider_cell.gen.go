
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SliderCell] class.
var SliderCellClass _SliderCellClass

func init() {
	SliderCellClass = _SliderCellClass{objc.GetClass("NSSliderCell")}
}

type _SliderCellClass struct {
	objc.Class
}

// An interface definition for the [SliderCell] class.
type ISliderCell interface {
	ID() objc.ID
}

type SliderCell struct {
	id objc.ID
}

func SliderCellFrom(ptr unsafe.Pointer) SliderCell {
	return SliderCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SliderCell) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SliderCellClass) Alloc() SliderCell {
	rv := objc.Send[SliderCell](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SliderCellClass) New() SliderCell {
	rv := objc.Send[SliderCell](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSliderCell creates and returns a new initialized instance.
func NewSliderCell() SliderCell {
	return SliderCellClass.New()
}

// Init initializes the instance.
func (s_ SliderCell) Init() SliderCell {
	rv := objc.Send[SliderCell](s_.ID(), selInit)
	return rv
}
