// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ColorWell] class.
var (
	colorWellClass     _ColorWellClass
	colorWellClassOnce sync.Once
)

func getColorWellClass() _ColorWellClass {
	colorWellClassOnce.Do(func() {
		colorWellClass = _ColorWellClass{objc.GetClass("NSColorWell")}
	})
	return colorWellClass
}

type _ColorWellClass struct {
	class objc.Class
}

// An interface definition for the [ColorWell] class.
type IColorWell interface {
	IControl
}

// A control that displays a color value and lets the user change that color value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell
type ColorWell struct {
	Control
}

// ColorWellFrom constructs a [ColorWell] from an unsafe.Pointer.
//
// A control that displays a color value and lets the user change that color value.
func ColorWellFrom(ptr unsafe.Pointer) ColorWell {
	return ColorWell{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ColorWellClass) Alloc() ColorWell {
	rv := objc.Send[ColorWell](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ColorWellClass) New() ColorWell {
	rv := objc.Send[ColorWell](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ColorWell) Init() ColorWell {
	rv := objc.Send[ColorWell](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ColorWell) Autorelease() ColorWell {
	rv := objc.Send[ColorWell](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewColorWell creates a new ColorWell instance.
func NewColorWell() ColorWell {
	return getColorWellClass().New()
}




