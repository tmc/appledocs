
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CustomTouchBarItem] class.
var CustomTouchBarItemClass _CustomTouchBarItemClass

func init() {
	CustomTouchBarItemClass = _CustomTouchBarItemClass{objc.GetClass("NSCustomTouchBarItem")}
}

type _CustomTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [CustomTouchBarItem] class.
type ICustomTouchBarItem interface {
	ID() objc.ID
}

type CustomTouchBarItem struct {
	id objc.ID
}

func CustomTouchBarItemFrom(ptr unsafe.Pointer) CustomTouchBarItem {
	return CustomTouchBarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ CustomTouchBarItem) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _CustomTouchBarItemClass) Alloc() CustomTouchBarItem {
	rv := objc.Send[CustomTouchBarItem](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _CustomTouchBarItemClass) New() CustomTouchBarItem {
	rv := objc.Send[CustomTouchBarItem](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewCustomTouchBarItem creates and returns a new initialized instance.
func NewCustomTouchBarItem() CustomTouchBarItem {
	return CustomTouchBarItemClass.New()
}

// Init initializes the instance.
func (c_ CustomTouchBarItem) Init() CustomTouchBarItem {
	rv := objc.Send[CustomTouchBarItem](c_.ID(), selInit)
	return rv
}
