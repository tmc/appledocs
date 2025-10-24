// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AccessoryView() IView
	SetAccessoryView(value IView)
	AttributedTitle() objc.IObject /* cross-framework: AttributedString */
	SetAttributedTitle(value objc.IObject /* cross-framework: AttributedString */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	ToolTip() objc.IObject /* cross-framework: NSString */
	SetToolTip(value objc.IObject /* cross-framework: NSString */)
	Tab() IWindowTab
	SetTab(value IWindowTab)
	TabbingIdentifier() unsafe.Pointer
	SetTabbingIdentifier(value unsafe.Pointer)
	// methods:
}

// A tab associated with a window that is part of a tabbing group.
//
// describes the way a window displays as part of a tabbed window group. The properties of are configurable at any time, but only take effect when the associated displays in a tab. AppKit automatically creates an instance of for each . You can access a window’s tab object using the property.


// A tab associated with a window that is part of a tabbing group.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/accessoryView
func (w_ WindowTab) AccessoryView() IView {
	rv := objc.Send[View](w_.ID, objc.Sel("accessoryView"))
	return rv
}


// An optional accessory view for the tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/accessoryView
func (w_ WindowTab) SetAccessoryView(value IView) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAccessoryView:"), value)
}


// The title for the window tab, specified as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/attributedTitle
func (w_ WindowTab) AttributedTitle() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[foundation.AttributedString](w_.ID, objc.Sel("attributedTitle"))
	return rv
}


// The title for the window tab, specified as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/attributedTitle
func (w_ WindowTab) SetAttributedTitle(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAttributedTitle:"), value)
}


// The title for the window tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/title
func (w_ WindowTab) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("title"))
	return rv
}


// The title for the window tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/title
func (w_ WindowTab) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitle:"), value)
}


// The tooltip for this window tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/toolTip
func (w_ WindowTab) ToolTip() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("toolTip"))
	return rv
}


// The tooltip for this window tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab/toolTip
func (w_ WindowTab) SetToolTip(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setToolTip:"), value)
}


// An object that represents information about a window when it displays as a tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/tab
func (w_ WindowTab) Tab() IWindowTab {
	rv := objc.Send[WindowTab](w_.ID, objc.Sel("tab"))
	return rv
}


// An object that represents information about a window when it displays as a tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/tab
func (w_ WindowTab) SetTab(value IWindowTab) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTab:"), value)
}


// A value that allows a group of related windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/tabbingidentifier-swift.property
func (w_ WindowTab) TabbingIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("tabbingIdentifier"))
	return rv
}


// A value that allows a group of related windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/tabbingidentifier-swift.property
func (w_ WindowTab) SetTabbingIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTabbingIdentifier:"), value)
}



