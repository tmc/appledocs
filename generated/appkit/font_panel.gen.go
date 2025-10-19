// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FontPanel] class.
var (
	fontPanelClass     _FontPanelClass
	fontPanelClassOnce sync.Once
)

func getFontPanelClass() _FontPanelClass {
	fontPanelClassOnce.Do(func() {
		fontPanelClass = _FontPanelClass{objc.GetClass("NSFontPanel")}
	})
	return fontPanelClass
}

type _FontPanelClass struct {
	class objc.Class
}

// An interface definition for the [FontPanel] class.
type IFontPanel interface {
	IPanel
}

// The Font panel—a user interface object that displays a list of available fonts, letting the user preview them and change the font used to display text.
//
// Actual changes to the font panel are made through conversion messages sent to the shared instance. There’s only one Font panel for each app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel
type FontPanel struct {
	Panel
}

// FontPanelFrom constructs a [FontPanel] from an unsafe.Pointer.
//
// The Font panel—a user interface object that displays a list of available fonts, letting the user preview them and change the font used to display text.
func FontPanelFrom(ptr unsafe.Pointer) FontPanel {
	return FontPanel{
		Panel: PanelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FontPanelClass) Alloc() FontPanel {
	rv := objc.Send[FontPanel](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FontPanelClass) New() FontPanel {
	rv := objc.Send[FontPanel](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FontPanel) Init() FontPanel {
	rv := objc.Send[FontPanel](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FontPanel) Autorelease() FontPanel {
	rv := objc.Send[FontPanel](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFontPanel creates a new FontPanel instance.
func NewFontPanel() FontPanel {
	return getFontPanelClass().New()
}




