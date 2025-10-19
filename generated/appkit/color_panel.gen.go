// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ColorPanel] class.
var (
	colorPanelClass     _ColorPanelClass
	colorPanelClassOnce sync.Once
)

func getColorPanelClass() _ColorPanelClass {
	colorPanelClassOnce.Do(func() {
		colorPanelClass = _ColorPanelClass{objc.GetClass("NSColorPanel")}
	})
	return colorPanelClass
}

type _ColorPanelClass struct {
	class objc.Class
}

// An interface definition for the [ColorPanel] class.
type IColorPanel interface {
	IPanel
}

// A standard user interface for selecting color in an app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel
type ColorPanel struct {
	Panel
}

// ColorPanelFrom constructs a [ColorPanel] from an unsafe.Pointer.
//
// A standard user interface for selecting color in an app.
func ColorPanelFrom(ptr unsafe.Pointer) ColorPanel {
	return ColorPanel{
		Panel: PanelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ColorPanelClass) Alloc() ColorPanel {
	rv := objc.Send[ColorPanel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ColorPanelClass) New() ColorPanel {
	rv := objc.Send[ColorPanel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ColorPanel) Init() ColorPanel {
	rv := objc.Send[ColorPanel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ColorPanel) Autorelease() ColorPanel {
	rv := objc.Send[ColorPanel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewColorPanel creates a new ColorPanel instance.
func NewColorPanel() ColorPanel {
	return getColorPanelClass().New()
}




