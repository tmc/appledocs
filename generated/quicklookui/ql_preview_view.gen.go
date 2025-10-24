// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [PreviewView] class.
var (
	PreviewViewClass     _PreviewViewClass
	PreviewViewClassOnce sync.Once
)

func getPreviewViewClass() _PreviewViewClass {
	PreviewViewClassOnce.Do(func() {
		PreviewViewClass = _PreviewViewClass{objc.GetClass("QLPreviewView")}
	})
	return PreviewViewClass
}

type _PreviewViewClass struct {
	class objc.Class
}

// An interface definition for the [PreviewView] class.
type IPreviewView interface {
	appkit.IView
	// properties:
	Autostarts() bool
	SetAutostarts(value bool)
	DisplayState() unsafe.Pointer
	SetDisplayState(value unsafe.Pointer)
	PreviewItem() PreviewItem /* not a class type */
	SetPreviewItem(value PreviewItem /* not a class type */)
	ShouldCloseWithWindow() bool
	SetShouldCloseWithWindow(value bool)
	// methods:
}

// A Quick Look preview of an item that you can embed into your view hierarchy.


// A Quick Look preview of an item that you can embed into your view hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView
type PreviewView struct {
	appkit.View
}

// PreviewViewFrom constructs a [PreviewView] from an unsafe.Pointer.
//
// A Quick Look preview of an item that you can embed into your view hierarchy.
func PreviewViewFrom(ptr unsafe.Pointer) PreviewView {
	return PreviewView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PreviewViewClass) Alloc() PreviewView {
	rv := objc.Send[PreviewView](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PreviewViewClass) New() PreviewView {
	rv := objc.Send[PreviewView](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewView) Init() PreviewView {
	rv := objc.Send[PreviewView](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewView) Autorelease() PreviewView {
	rv := objc.Send[PreviewView](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewView creates a new PreviewView instance.
func NewPreviewView() PreviewView {
	return getPreviewViewClass().New()
}



// A Boolean value that determines whether the preview starts automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView/autostarts
func (p_ PreviewView) Autostarts() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("autostarts"))
	return rv
}


// A Boolean value that determines whether the preview starts automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewView/autostarts
func (p_ PreviewView) SetAutostarts(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutostarts:"), value)
}


// The current display state of the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewview/displaystate
func (p_ PreviewView) DisplayState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("displayState"))
	return rv
}


// The current display state of the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewview/displaystate
func (p_ PreviewView) SetDisplayState(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplayState:"), value)
}


// The item to preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewview/previewitem
func (p_ PreviewView) PreviewItem() PreviewItem /* not a class type */ {
	rv := objc.Send[PreviewItem](p_.ID, objc.Sel("previewItem"))
	return rv
}


// The item to preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewview/previewitem
func (p_ PreviewView) SetPreviewItem(value PreviewItem /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreviewItem:"), value)
}


// A Boolean value that determines whether the preview should close when its window closes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewview/shouldclosewithwindow
func (p_ PreviewView) ShouldCloseWithWindow() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldCloseWithWindow"))
	return rv
}


// A Boolean value that determines whether the preview should close when its window closes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewview/shouldclosewithwindow
func (p_ PreviewView) SetShouldCloseWithWindow(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldCloseWithWindow:"), value)
}




