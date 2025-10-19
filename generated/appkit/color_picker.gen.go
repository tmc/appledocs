// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ColorPicker] class.
var (
	colorPickerClass     _ColorPickerClass
	colorPickerClassOnce sync.Once
)

func getColorPickerClass() _ColorPickerClass {
	colorPickerClassOnce.Do(func() {
		colorPickerClass = _ColorPickerClass{objc.GetClass("NSColorPicker")}
	})
	return colorPickerClass
}

type _ColorPickerClass struct {
	class objc.Class
}

// An interface definition for the [ColorPicker] class.
type IColorPicker interface {
	objectivec.IObject
}

// An abstract superclass that implements the default color picking protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPicker
type ColorPicker struct {
	objectivec.Object
}

// ColorPickerFrom constructs a [ColorPicker] from an unsafe.Pointer.
//
// An abstract superclass that implements the default color picking protocol.
func ColorPickerFrom(ptr unsafe.Pointer) ColorPicker {
	return ColorPicker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ColorPickerClass) Alloc() ColorPicker {
	rv := objc.Send[ColorPicker](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ColorPickerClass) New() ColorPicker {
	rv := objc.Send[ColorPicker](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ColorPicker) Init() ColorPicker {
	rv := objc.Send[ColorPicker](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ColorPicker) Autorelease() ColorPicker {
	rv := objc.Send[ColorPicker](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewColorPicker creates a new ColorPicker instance.
func NewColorPicker() ColorPicker {
	return getColorPickerClass().New()
}




