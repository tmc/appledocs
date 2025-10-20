// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WindowTab] class.
var (
	WindowTabClass     _WindowTabClass
	WindowTabClassOnce sync.Once
)

func getWindowTabClass() _WindowTabClass {
	WindowTabClassOnce.Do(func() {
		WindowTabClass = _WindowTabClass{objc.GetClass("NSWindowTab")}
	})
	return WindowTabClass
}

type _WindowTabClass struct {
	class objc.Class
}

// An interface definition for the [WindowTab] class.
type IWindowTab interface {
	objectivec.IObject
}

// A tab associated with a window that is part of a tabbing group.
//
// describes the way a window displays as part of a tabbed window group. The properties of are configurable at any time, but only take effect when the associated displays in a tab. AppKit automatically creates an instance of for each . You can access a window’s tab object using the property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab
type WindowTab struct {
	objectivec.Object
}

// WindowTabFrom constructs a [WindowTab] from an unsafe.Pointer.
//
// A tab associated with a window that is part of a tabbing group.
func WindowTabFrom(ptr unsafe.Pointer) WindowTab {
	return WindowTab{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WindowTabClass) Alloc() WindowTab {
	rv := objc.Send[WindowTab](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WindowTabClass) New() WindowTab {
	rv := objc.Send[WindowTab](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WindowTab) Init() WindowTab {
	rv := objc.Send[WindowTab](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WindowTab) Autorelease() WindowTab {
	rv := objc.Send[WindowTab](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWindowTab creates a new WindowTab instance.
func NewWindowTab() WindowTab {
	return getWindowTabClass().New()
}

// An optional accessory view for the tab.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/accessoryView
func (w_ WindowTab) AccessoryView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("accessoryView"))
	return rv
}

// SetAccessoryView sets the value of the accessoryView property.
// An optional accessory view for the tab.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/accessoryView
func (w_ WindowTab) SetAccessoryView(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAccessoryView:"), value)
}

// The title for the window tab, specified as an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/attributedTitle
func (w_ WindowTab) AttributedTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("attributedTitle"))
	return rv
}

// SetAttributedTitle sets the value of the attributedTitle property.
// The title for the window tab, specified as an attributed string.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/attributedTitle
func (w_ WindowTab) SetAttributedTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAttributedTitle:"), value)
}

// The title for the window tab.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/title
func (w_ WindowTab) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("title"))
	return rv
}

// SetTitle sets the value of the title property.
// The title for the window tab.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/title
func (w_ WindowTab) SetTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitle:"), value)
}

// The tooltip for this window tab.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/toolTip
func (w_ WindowTab) ToolTip() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("toolTip"))
	return rv
}

// SetToolTip sets the value of the toolTip property.
// The tooltip for this window tab.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/toolTip
func (w_ WindowTab) SetToolTip(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setToolTip:"), value)
}
