
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ColorPicker] class.
var ColorPickerClass _ColorPickerClass

func init() {
	ColorPickerClass = _ColorPickerClass{objc.GetClass("NSColorPicker")}
}

type _ColorPickerClass struct {
	objc.Class
}

// An interface definition for the [ColorPicker] class.
type IColorPicker interface {
	ID() objc.ID
}

type ColorPicker struct {
	id objc.ID
}

func ColorPickerFrom(ptr unsafe.Pointer) ColorPicker {
	return ColorPicker{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ ColorPicker) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ColorPickerClass) Alloc() ColorPicker {
	rv := objc.Send[ColorPicker](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ColorPickerClass) New() ColorPicker {
	rv := objc.Send[ColorPicker](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewColorPicker creates and returns a new initialized instance.
func NewColorPicker() ColorPicker {
	return ColorPickerClass.New()
}

// Init initializes the instance.
func (c_ ColorPicker) Init() ColorPicker {
	rv := objc.Send[ColorPicker](c_.ID(), selInit)
	return rv
}
