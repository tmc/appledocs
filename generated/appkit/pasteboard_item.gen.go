// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PasteboardItem] class.
var (
	pasteboardItemClass     _PasteboardItemClass
	pasteboardItemClassOnce sync.Once
)

func getPasteboardItemClass() _PasteboardItemClass {
	pasteboardItemClassOnce.Do(func() {
		pasteboardItemClass = _PasteboardItemClass{objc.GetClass("NSPasteboardItem")}
	})
	return pasteboardItemClass
}

type _PasteboardItemClass struct {
	class objc.Class
}

// An interface definition for the [PasteboardItem] class.
type IPasteboardItem interface {
	objectivec.IObject
}

// An item on a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem

type PasteboardItem struct {
	objectivec.Object
}

// PasteboardItemFrom constructs a [PasteboardItem] from an unsafe.Pointer.
//
// An item on a pasteboard.
func PasteboardItemFrom(ptr unsafe.Pointer) PasteboardItem {
	return PasteboardItem{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (pc _PasteboardItemClass) Alloc() PasteboardItem {
	rv := objc.Send[PasteboardItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PasteboardItemClass) New() PasteboardItem {
	rv := objc.Send[PasteboardItem](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasteboardItem) Init() PasteboardItem {
	rv := objc.Send[PasteboardItem](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasteboardItem) Autorelease() PasteboardItem {
	rv := objc.Send[PasteboardItem](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasteboardItem creates a new PasteboardItem instance.
func NewPasteboardItem() PasteboardItem {
	return getPasteboardItemClass().New()
}




