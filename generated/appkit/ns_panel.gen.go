// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Panel] class.
var (
	PanelClass     _PanelClass
	PanelClassOnce sync.Once
)

func getPanelClass() _PanelClass {
	PanelClassOnce.Do(func() {
		PanelClass = _PanelClass{objc.GetClass("NSPanel")}
	})
	return PanelClass
}

type _PanelClass struct {
	class objc.Class
}

// An interface definition for the [Panel] class.
type IPanel interface {
	IWindow
}

// A special kind of window that typically performs a function that is auxiliary to the main window.
//
// For details about how panels work (especially to find out how their behavior differs from window behavior), see .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A Boolean value that indicates whether the receiver becomes the key window only when needed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel/becomesKeyOnlyIfNeeded
func (p_ Panel) BecomesKeyOnlyIfNeeded() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("becomesKeyOnlyIfNeeded"))
	return rv
}


// SetBecomesKeyOnlyIfNeeded sets the value of the becomesKeyOnlyIfNeeded property.
// A Boolean value that indicates whether the receiver becomes the key window only when needed.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel/becomesKeyOnlyIfNeeded
func (p_ Panel) SetBecomesKeyOnlyIfNeeded(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBecomesKeyOnlyIfNeeded:"), value)
}

// A Boolean value that indicates whether the receiver is a floating panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel/isFloatingPanel
func (p_ Panel) FloatingPanel() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("floatingPanel"))
	return rv
}


// SetFloatingPanel sets the value of the floatingPanel property.
// A Boolean value that indicates whether the receiver is a floating panel.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel/isFloatingPanel
func (p_ Panel) SetFloatingPanel(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFloatingPanel:"), value)
}

// A Boolean value that indicates whether the panel receives keyboard and mouse events even when some other window is being run modally.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel/worksWhenModal
func (p_ Panel) WorksWhenModal() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("worksWhenModal"))
	return rv
}


// SetWorksWhenModal sets the value of the worksWhenModal property.
// A Boolean value that indicates whether the panel receives keyboard and mouse events even when some other window is being run modally.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel/worksWhenModal
func (p_ Panel) SetWorksWhenModal(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWorksWhenModal:"), value)
}



