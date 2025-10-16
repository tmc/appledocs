
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ButtonTouchBarItem] class.
var ButtonTouchBarItemClass _ButtonTouchBarItemClass

func init() {
	ButtonTouchBarItemClass = _ButtonTouchBarItemClass{objc.GetClass("NSButtonTouchBarItem")}
}

type _ButtonTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [ButtonTouchBarItem] class.
type IButtonTouchBarItem interface {
	ID() objc.ID
}

type ButtonTouchBarItem struct {
	id objc.ID
}

func ButtonTouchBarItemFrom(ptr unsafe.Pointer) ButtonTouchBarItem {
	return ButtonTouchBarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ ButtonTouchBarItem) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _ButtonTouchBarItemClass) Alloc() ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _ButtonTouchBarItemClass) New() ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewButtonTouchBarItem creates and returns a new initialized instance.
func NewButtonTouchBarItem() ButtonTouchBarItem {
	return ButtonTouchBarItemClass.New()
}

// Init initializes the instance.
func (b_ ButtonTouchBarItem) Init() ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](b_.ID(), selInit)
	return rv
}
