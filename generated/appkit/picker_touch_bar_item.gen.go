
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PickerTouchBarItem] class.
var PickerTouchBarItemClass _PickerTouchBarItemClass

func init() {
	PickerTouchBarItemClass = _PickerTouchBarItemClass{objc.GetClass("NSPickerTouchBarItem")}
}

type _PickerTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [PickerTouchBarItem] class.
type IPickerTouchBarItem interface {
	ID() objc.ID
}

type PickerTouchBarItem struct {
	id objc.ID
}

func PickerTouchBarItemFrom(ptr unsafe.Pointer) PickerTouchBarItem {
	return PickerTouchBarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PickerTouchBarItem) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PickerTouchBarItemClass) Alloc() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PickerTouchBarItemClass) New() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPickerTouchBarItem creates and returns a new initialized instance.
func NewPickerTouchBarItem() PickerTouchBarItem {
	return PickerTouchBarItemClass.New()
}

// Init initializes the instance.
func (p_ PickerTouchBarItem) Init() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](p_.ID(), selInit)
	return rv
}
