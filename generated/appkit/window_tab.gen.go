
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [WindowTab] class.
var WindowTabClass _WindowTabClass

func init() {
	WindowTabClass = _WindowTabClass{objc.GetClass("NSWindowTab")}
}

type _WindowTabClass struct {
	objc.Class
}

// An interface definition for the [WindowTab] class.
type IWindowTab interface {
	ID() objc.ID
}

type WindowTab struct {
	id objc.ID
}

func WindowTabFrom(ptr unsafe.Pointer) WindowTab {
	return WindowTab{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ WindowTab) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _WindowTabClass) Alloc() WindowTab {
	rv := objc.Send[WindowTab](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _WindowTabClass) New() WindowTab {
	rv := objc.Send[WindowTab](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewWindowTab creates and returns a new initialized instance.
func NewWindowTab() WindowTab {
	return WindowTabClass.New()
}

// Init initializes the instance.
func (w_ WindowTab) Init() WindowTab {
	rv := objc.Send[WindowTab](w_.ID(), selInit)
	return rv
}
// An optional accessory view for the tab. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTab/accessoryView
func (w_ WindowTab) AccessoryView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("accessoryView"))
	return rv
}
// SetAccessoryView sets the value of the accessoryView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTab/accessoryView
func (w_ WindowTab) SetAccessoryView(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setAccessoryView:"), value)
}
// The title for the window tab, specified as an attributed string. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTab/attributedTitle
func (w_ WindowTab) AttributedTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("attributedTitle"))
	return rv
}
// SetAttributedTitle sets the value of the attributedTitle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTab/attributedTitle
func (w_ WindowTab) SetAttributedTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setAttributedTitle:"), value)
}
// The title for the window tab. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTab/title
func (w_ WindowTab) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("title"))
	return rv
}
// SetTitle sets the value of the title property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTab/title
func (w_ WindowTab) SetTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setTitle:"), value)
}
// The tooltip for this window tab. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTab/toolTip
func (w_ WindowTab) ToolTip() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("toolTip"))
	return rv
}
// SetToolTip sets the value of the toolTip property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTab/toolTip
func (w_ WindowTab) SetToolTip(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setToolTip:"), value)
}
