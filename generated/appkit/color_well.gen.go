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

// A control that displays a color value and lets the user change that color value.
//
// An object lets people select colors from your interface. Incorporate this type of control if your app supports custom color selection. For example, a drawing app might include a color well to let someone choose the color to use when drawing. A color well control displays the currently selected color, and interactions with the color well display interfaces for selecting new colors. When you create a color well programmatically or in Interface Builder, specify the appearance and interaction style you want. The color well supports color selection using a color picker popover or the system object. When someone selects a new color in one of these interfaces, the color well updates its selected color to match. You can also provide your own color selection process using a custom action and update the color yourself.
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


// A Boolean value that determines whether the color well has a border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/isBordered
func (c_ ColorWell) Bordered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("bordered"))
	return rv
}

// SetBordered sets the value of the bordered property.
// A Boolean value that determines whether the color well has a border.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/isBordered
func (c_ ColorWell) SetBordered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBordered:"), value)
}
// The target object that defines the action you want to perform when someone interacts with the color well.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/pulldownTarget
func (c_ ColorWell) PulldownTarget() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("pulldownTarget"))
	return rv
}

// SetPulldownTarget sets the value of the pulldownTarget property.
// The target object that defines the action you want to perform when someone interacts with the color well.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/pulldownTarget
func (c_ ColorWell) SetPulldownTarget(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPulldownTarget:"), value)
}


