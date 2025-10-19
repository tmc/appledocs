// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Panel] class.
var (
	panelClass     _PanelClass
	panelClassOnce sync.Once
)

func getPanelClass() _PanelClass {
	panelClassOnce.Do(func() {
		panelClass = _PanelClass{objc.GetClass("NSPanel")}
	})
	return panelClass
}

type _PanelClass struct {
	class objc.Class
}

// An interface definition for the [Panel] class.
type IPanel interface {
	IWindow
}

// A special kind of window that typically performs a function that is auxiliary to the main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel

type Panel struct {
	Window
}

// PanelFrom constructs a [Panel] from an unsafe.Pointer.
//
// A special kind of window that typically performs a function that is auxiliary to the main window.
func PanelFrom(ptr unsafe.Pointer) Panel {
	return Panel{
		Window: WindowFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (pc _PanelClass) Alloc() Panel {
	rv := objc.Send[Panel](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PanelClass) New() Panel {
	rv := objc.Send[Panel](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Panel) Init() Panel {
	rv := objc.Send[Panel](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Panel) Autorelease() Panel {
	rv := objc.Send[Panel](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPanel creates a new Panel instance.
func NewPanel() Panel {
	return getPanelClass().New()
}




