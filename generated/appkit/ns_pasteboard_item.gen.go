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
	PasteboardItemClass     _PasteboardItemClass
	PasteboardItemClassOnce sync.Once
)

func getPasteboardItemClass() _PasteboardItemClass {
	PasteboardItemClassOnce.Do(func() {
		PasteboardItemClass = _PasteboardItemClass{objc.GetClass("NSPasteboardItem")}
	})
	return PasteboardItemClass
}

type _PasteboardItemClass struct {
	class objc.Class
}

// An interface definition for the [PasteboardItem] class.
type IPasteboardItem interface {
	objectivec.IObject
	StringForType(type_ unsafe.Pointer) unsafe.Pointer
}

// An item on a pasteboard.
//
// There are three main uses for an object: Providing data on the pasteboard. You can create one or more pasteboard items, set data or data providers for types, and write them to the pasteboard. Customizing data already on the pasteboard. As a delegate or subclass, you can retrieve the pasteboard items currently on the pasteboard, read the existing types and data, and set new data and data providers for types as necessary. Retrieving data from the pasteboard. You can retrieve pasteboard items from the pasteboard and then read the data for types you’re interested in. A pasteboard item can be associated with a single pasteboard. When you create an item, you can write it to any pasteboard. When you pass an item to a pasteboard in , that item becomes bound to the pasteboard it writes to. When you retrieve items from a pasteboard using or , the returned items are associated with the messaged pasteboard. Passing an item that is already associated with a pasteboard into causes an exception. Use pasteboard items during a single pasteboard interaction, rather than retaining and reusing them. A pasteboard item is only valid until the owner of the pasteboard changes.
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

// Returns the value for the specified type as a string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/string(forType:)
func (p_ PasteboardItem) StringForType(type_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("stringForType:"), type_)
	return rv
}

// A model object you use for conveying data during a collaboration.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/collaborationMetadata
func (p_ PasteboardItem) CollaborationMetadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("collaborationMetadata"))
	return rv
}

// SetCollaborationMetadata sets the value of the collaborationMetadata property.
// A model object you use for conveying data during a collaboration.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/collaborationMetadata
func (p_ PasteboardItem) SetCollaborationMetadata(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollaborationMetadata:"), value)
}
