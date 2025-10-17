
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ColorPickerTouchBarItem] class.
var ColorPickerTouchBarItemClass _ColorPickerTouchBarItemClass

func init() {
	ColorPickerTouchBarItemClass = _ColorPickerTouchBarItemClass{objc.GetClass("NSColorPickerTouchBarItem")}
}

type _ColorPickerTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [ColorPickerTouchBarItem] class.
type IColorPickerTouchBarItem interface {
	ID() objc.ID
}

type ColorPickerTouchBarItem struct {
	id objc.ID
}

func ColorPickerTouchBarItemFrom(ptr unsafe.Pointer) ColorPickerTouchBarItem {
	return ColorPickerTouchBarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ ColorPickerTouchBarItem) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ColorPickerTouchBarItemClass) Alloc() ColorPickerTouchBarItem {
	rv := objc.Send[ColorPickerTouchBarItem](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ColorPickerTouchBarItemClass) New() ColorPickerTouchBarItem {
	rv := objc.Send[ColorPickerTouchBarItem](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewColorPickerTouchBarItem creates and returns a new initialized instance.
func NewColorPickerTouchBarItem() ColorPickerTouchBarItem {
	return ColorPickerTouchBarItemClass.New()
}

// Init initializes the instance.
func (c_ ColorPickerTouchBarItem) Init() ColorPickerTouchBarItem {
	rv := objc.Send[ColorPickerTouchBarItem](c_.ID(), selInit)
	return rv
}
