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
	

	// properties:
	Tab() IWindowTab
	SetTab(value IWindowTab)
	TabbingIdentifier() objectivec.IObject
	SetTabbingIdentifier(value objectivec.IObject)
	AccessoryView() IView
	SetAccessoryView(value IView)
	AttributedTitle() foundation.foundation.INSAttributedString
	SetAttributedTitle(value foundation.foundation.INSAttributedString)
	Title() foundation.foundation.INSString
	SetTitle(value foundation.foundation.INSString)
	ToolTip() foundation.foundation.INSString
	SetToolTip(value foundation.foundation.INSString)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (wc _WindowTabClass) Alloc() WindowTab {
	rv := objc.Send[WindowTab](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
func (w_ WindowTab) TabbingIdentifier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](w_.ID, objc.Sel("tabbingIdentifier"))
	return rv
}


// A value that allows a group of related windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/tabbingidentifier-swift.property
func (w_ WindowTab) SetTabbingIdentifier(value objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTabbingIdentifier:"), value)
}


// An optional accessory view for the tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/accessoryview
func (w_ WindowTab) AccessoryView() IView {
	rv := objc.Send[View](w_.ID, objc.Sel("accessoryView"))
	return rv
}


// An optional accessory view for the tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/accessoryview
func (w_ WindowTab) SetAccessoryView(value IView) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAccessoryView:"), value)
}


// The title for the window tab, specified as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/attributedtitle
func (w_ WindowTab) AttributedTitle() foundation.foundation.INSAttributedString {
	rv := objc.Send[foundation.NSAttributedString](w_.ID, objc.Sel("attributedTitle"))
	return rv
}


// The title for the window tab, specified as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/attributedtitle
func (w_ WindowTab) SetAttributedTitle(value foundation.foundation.INSAttributedString) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAttributedTitle:"), value)
}


// The title for the window tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/title
func (w_ WindowTab) Title() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("title"))
	return rv
}


// The title for the window tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/title
func (w_ WindowTab) SetTitle(value foundation.foundation.INSString) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitle:"), value)
}


// The tooltip for this window tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/tooltip
func (w_ WindowTab) ToolTip() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("toolTip"))
	return rv
}


// The tooltip for this window tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/tooltip
func (w_ WindowTab) SetToolTip(value foundation.foundation.INSString) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setToolTip:"), value)
}








