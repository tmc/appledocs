
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TouchBarItem] class.
var TouchBarItemClass _TouchBarItemClass

func init() {
	TouchBarItemClass = _TouchBarItemClass{objc.GetClass("NSTouchBarItem")}
}

type _TouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [TouchBarItem] class.
type ITouchBarItem interface {
	ID() objc.ID
}

type TouchBarItem struct {
	id objc.ID
}

func TouchBarItemFrom(ptr unsafe.Pointer) TouchBarItem {
	return TouchBarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TouchBarItem) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TouchBarItemClass) Alloc() TouchBarItem {
	rv := objc.Send[TouchBarItem](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TouchBarItemClass) New() TouchBarItem {
	rv := objc.Send[TouchBarItem](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTouchBarItem creates and returns a new initialized instance.
func NewTouchBarItem() TouchBarItem {
	return TouchBarItemClass.New()
}

// Init initializes the instance.
func (t_ TouchBarItem) Init() TouchBarItem {
	rv := objc.Send[TouchBarItem](t_.ID(), selInit)
	return rv
}
// The user-visible string identifying this item during bar customization. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBarItem/customizationLabel
func (t_ TouchBarItem) CustomizationLabel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("customizationLabel"))
	return rv
}
