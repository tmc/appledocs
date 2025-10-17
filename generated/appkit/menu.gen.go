
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Menu] class.
var MenuClass _MenuClass

func init() {
	MenuClass = _MenuClass{objc.GetClass("NSMenu")}
}

type _MenuClass struct {
	objc.Class
}

// An interface definition for the [Menu] class.
type IMenu interface {
	ID() objc.ID
}

type Menu struct {
	id objc.ID
}

func MenuFrom(ptr unsafe.Pointer) Menu {
	return Menu{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ Menu) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _MenuClass) Alloc() Menu {
	rv := objc.Send[Menu](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _MenuClass) New() Menu {
	rv := objc.Send[Menu](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewMenu creates and returns a new initialized instance.
func NewMenu() Menu {
	return MenuClass.New()
}

// Init initializes the instance.
func (m_ Menu) Init() Menu {
	rv := objc.Send[Menu](m_.ID(), selInit)
	return rv
}
// Displays a contextual menu over a view for an event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSMenu/popUpContextMenu(_:with:for:)
func (mc _MenuClass) PopUpContextMenuWithEventForView(menu unsafe.Pointer, event unsafe.Pointer, view unsafe.Pointer)  {
	objc.Send[objc.ID](objc.ID(mc.Class), objc.RegisterName("popUpContextMenu:withEvent:forView:"), menu, event, view)
}
