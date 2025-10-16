
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SharingServicePickerToolbarItem] class.
var SharingServicePickerToolbarItemClass _SharingServicePickerToolbarItemClass

func init() {
	SharingServicePickerToolbarItemClass = _SharingServicePickerToolbarItemClass{objc.GetClass("NSSharingServicePickerToolbarItem")}
}

type _SharingServicePickerToolbarItemClass struct {
	objc.Class
}

// An interface definition for the [SharingServicePickerToolbarItem] class.
type ISharingServicePickerToolbarItem interface {
	ID() objc.ID
}

type SharingServicePickerToolbarItem struct {
	id objc.ID
}

func SharingServicePickerToolbarItemFrom(ptr unsafe.Pointer) SharingServicePickerToolbarItem {
	return SharingServicePickerToolbarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SharingServicePickerToolbarItem) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SharingServicePickerToolbarItemClass) Alloc() SharingServicePickerToolbarItem {
	rv := objc.Send[SharingServicePickerToolbarItem](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SharingServicePickerToolbarItemClass) New() SharingServicePickerToolbarItem {
	rv := objc.Send[SharingServicePickerToolbarItem](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSharingServicePickerToolbarItem creates and returns a new initialized instance.
func NewSharingServicePickerToolbarItem() SharingServicePickerToolbarItem {
	return SharingServicePickerToolbarItemClass.New()
}

// Init initializes the instance.
func (s_ SharingServicePickerToolbarItem) Init() SharingServicePickerToolbarItem {
	rv := objc.Send[SharingServicePickerToolbarItem](s_.ID(), selInit)
	return rv
}
